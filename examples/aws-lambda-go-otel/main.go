package main

import (
	"context"
	"errors"
	"log"
	"sync/atomic"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploghttp"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	otellog "go.opentelemetry.io/otel/log"
	"go.opentelemetry.io/otel/metric"
	sdklog "go.opentelemetry.io/otel/sdk/log"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

var (
	tp *sdktrace.TracerProvider
	mp *sdkmetric.MeterProvider
	lp *sdklog.LoggerProvider

	invocations metric.Int64Counter
	errorsTotal metric.Int64Counter
	coldStarts  metric.Int64Counter
	isColdStart atomic.Bool
)

type event struct {
	Fail bool `json:"fail"`
}

func setupOTel(ctx context.Context) error {
	res, err := resource.New(ctx,
		resource.WithFromEnv(),
		resource.WithTelemetrySDK(),
	)
	if err != nil {
		return err
	}

	traceExp, err := otlptracehttp.New(ctx)
	if err != nil {
		return err
	}
	metricExp, err := otlpmetrichttp.New(ctx)
	if err != nil {
		return err
	}
	logExp, err := otlploghttp.New(ctx)
	if err != nil {
		return err
	}

	tp = sdktrace.NewTracerProvider(
		sdktrace.WithResource(res),
		sdktrace.WithBatcher(traceExp),
	)
	mp = sdkmetric.NewMeterProvider(
		sdkmetric.WithResource(res),
		sdkmetric.WithReader(sdkmetric.NewPeriodicReader(metricExp)),
	)
	lp = sdklog.NewLoggerProvider(
		sdklog.WithResource(res),
		sdklog.WithProcessor(sdklog.NewBatchProcessor(logExp)),
	)

	otel.SetTracerProvider(tp)
	otel.SetMeterProvider(mp)

	meter := mp.Meter("lambda-checkout")
	invocations, err = meter.Int64Counter("checkout.invocations")
	if err != nil {
		return err
	}
	errorsTotal, err = meter.Int64Counter("checkout.errors")
	if err != nil {
		return err
	}
	coldStarts, err = meter.Int64Counter("checkout.cold_starts")
	return err
}

func handler(ctx context.Context, input event) (string, error) {
	cold := isColdStart.Swap(false)
	ctx, span := otel.Tracer("lambda-checkout").Start(ctx, "checkout.process")
	span.SetAttributes(attribute.Bool("faas.coldstart", cold))
	defer flushTelemetry()
	defer span.End()

	invocations.Add(ctx, 1)
	if cold {
		coldStarts.Add(ctx, 1)
	}

	if input.Fail {
		err := errors.New("falha de teste no checkout")
		errorsTotal.Add(ctx, 1)
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		emitLog(ctx, otellog.SeverityError, "Checkout falhou", cold)
		return "", err
	}

	emitLog(ctx, otellog.SeverityInfo, "Checkout processado", cold)
	return "ImmunoOps conectado!", nil
}

func emitLog(ctx context.Context, severity otellog.Severity, message string, cold bool) {
	var record otellog.Record
	record.SetTimestamp(time.Now())
	record.SetBody(attribute.StringValue(message))
	record.SetSeverity(severity)
	record.AddAttributes(
		attribute.String("event.name", "checkout.processed"),
		attribute.Bool("faas.coldstart", cold),
	)
	lp.Logger("lambda-checkout").Emit(ctx, record)
}

func flushTelemetry() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := tp.ForceFlush(ctx); err != nil {
		log.Printf("trace export: %v", err)
	}
	if err := mp.ForceFlush(ctx); err != nil {
		log.Printf("metric export: %v", err)
	}
	if err := lp.ForceFlush(ctx); err != nil {
		log.Printf("log export: %v", err)
	}
}

func main() {
	isColdStart.Store(true)
	if err := setupOTel(context.Background()); err != nil {
		log.Fatal(err)
	}
	lambda.Start(handler)
}
