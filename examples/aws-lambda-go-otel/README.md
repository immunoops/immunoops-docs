# AWS Lambda Go + OpenTelemetry + ImmunoOps

Exemplo compilável do Quickstart publicado em `docs.immunoops.io`. Envia traces,
métricas e logs por OTLP/HTTP diretamente ao ImmunoOps.

## ZIP

```bash
go mod download
GOOS=linux GOARCH=arm64 CGO_ENABLED=0 \
  go build -trimpath -tags lambda.norpc -ldflags="-s -w" -o bootstrap .
zip function.zip bootstrap
```

Crie uma Lambda `provided.al2023`, arquitetura `arm64`, handler `bootstrap`,
timeout de 15 segundos e configure as variáveis descritas no Quickstart.

## AWS SAM

```bash
sam build
sam deploy --guided \
  --parameter-overrides ImmunoOpsToken=SUA_CHAVE
```

Não salve a chave em código, commits ou arquivos versionados. Em produção,
prefira resolver o segredo por um mecanismo seguro da sua esteira.

## Testes

Invocação normal:

```bash
aws lambda invoke \
  --function-name immunoops-otel-quickstart \
  --cli-binary-format raw-in-base64-out \
  --payload '{}' response.json
```

Invocação que registra erro no span:

```bash
aws lambda invoke \
  --function-name immunoops-otel-quickstart \
  --cli-binary-format raw-in-base64-out \
  --payload '{"fail":true}' response-error.json
```

O código mantém os providers entre warm starts e executa `ForceFlush` depois de
finalizar o span de cada invocação.
