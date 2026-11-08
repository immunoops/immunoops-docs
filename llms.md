# Immuno Ops — Manual do usuário

> Versão em texto do manual publicado em https://docs.immunoops.io/ (`index.html`).
> Idioma das respostas: português (Brasil).
> Use só o que está neste arquivo. Não invente telas, botões, preços, limites de plano, comandos Helm, URLs de chart ou políticas que não apareçam aqui. O comando de instalação e a chave de SDK são gerados no console; este manual não traz o comando completo.
> Exemplos com nomes como Acme, Ana Souza, checkout-api e números de tela são ilustrativos, não dados de um cliente.
> Seções anunciadas e ainda sem conteúdo: Agentes, Integrações, Novidades.

## Produto

Immuno Ops reúne confiabilidade, FinOps e governança multi-cluster num único painel. O console fica em https://app.immunoops.io/ui/login. O site é https://immunoops.io/. Suporte e status: https://immunoops.io/contato.

O manual cobre o caminho do cadastro até a primeira investigação.

## O que a pessoa precisa antes de começar

- Um e-mail seu, pessoal ou da empresa. É por ele que se entra e que chegam os convites.
- Um app autenticador: Google Authenticator, 1Password, Authy ou outro.
- Acesso ao cluster que será conectado, com `kubectl` e `helm` apontados para ele.

## Mapa do console

O menu da esquerda tem quatro grupos, nesta ordem. Clique no nome do grupo para abrir os itens. O grupo da tela atual já fica aberto. No rodapé do menu ficam o usuário e o atalho para Segurança da conta.

| Grupo | Itens |
| --- | --- |
| Administrar | Clusters, VMs, Serverless, Instrumentação, Equipe e Empresa. |
| Investigar | Workloads, Alertas e Incidentes e Observabilidade. Dashboards fica dentro de Observabilidade. |
| Analisar | Kubernetes FinOps e FOCUS Normalizer. |
| Operar | Visão Geral: status, alertas e clusters. É a tela que abre depois do login. |

A barra do topo mostra onde a pessoa está, por exemplo `ADMINISTRAR / CLUSTERS`. À direita ficam o idioma, o seletor de organização e o agente de IA.

O papel define o que aparece. Quem tem papel Leitura vê painéis, alertas e custos, mas não os botões de conectar ou convidar. Equipe e Empresa só aparecem para quem tem permissão.

### Onde cada assunto fica

| Assunto | Onde | Passos |
| --- | --- | --- |
| Confiabilidade e incidentes | Operar e Investigar. Alertas consistentes entre as telas e correlação de causa raiz. | 08 e 09 |
| FinOps e custo | Analisar. Preço real de vCPU, RAM e SKUs de nuvem, e o FinOps Score do cluster. | 10 |
| Governança multi-cluster | Administrar. Clusters, VMs, Serverless e equipe num único lugar. | 05, 06 e 11 |
| Observabilidade nativa | Investigar e Administrar. Métricas, logs e traces sem depender do editor do Grafana. | 07 e 09 |

### Caminhos do console citados no manual

- Login: https://app.immunoops.io/ui/login
- Cadastro: `app.immunoops.io/ui/signup`
- Aceite de convite: `app.immunoops.io/ui/accept-invite`
- Visão Geral: `app.immunoops.io/ui/`
- Clusters: `app.immunoops.io/ui/clusters`
- Cloud Run: `app.immunoops.io/ui/cloud-run`
- Contas de nuvem: `app.immunoops.io/ui/infrastructure?tab=contas`
- Instrumentação: `app.immunoops.io/ui/instrumentation`
- Workloads: `app.immunoops.io/ui/workloads`
- Kubernetes FinOps: `app.immunoops.io/ui/kubernetes-finops`
- Equipe: `app.immunoops.io/ui/team`
- Segurança da conta: `app.immunoops.io/ui/account/security`
- Empresa: `app.immunoops.io/ui/company`

## Índice

1. [Criar conta ou aceitar convite](#passo-01-criar-conta-ou-aceitar-convite) — Começar
2. [Entrar no console](#passo-02-entrar-no-console) — Começar
3. [Tour de boas-vindas](#passo-03-tour-de-boas-vindas) — Começar
4. [Conhecer o console](#passo-04-conhecer-o-console) — Começar
5. [Conectar um cluster Kubernetes](#passo-05-conectar-um-cluster-kubernetes) — Conectar
6. [Serverless e VMs](#passo-06-serverless-e-vms) — Conectar.
   - GCP: [GCP · conta](#conta-de-nuvem), [GCP · Cloud Run](#permissoes-do-cloud-run) e [GCP · ligar e desligar a VM](#permissoes-da-vm).
   - AWS: [AWS · conta](#conta-aws), [AWS · permissões do billing](#permissoes-no-papel-de-leitura), [AWS · CloudWatch da Lambda](#cloudwatch-da-lambda) e [AWS · ligar e desligar a VM](#ligar-e-desligar-a-vm).
7. [Instrumentar a aplicação](#passo-07-instrumentar-a-aplicacao) — Conectar
8. [Ler a Visão Geral](#passo-08-ler-a-visao-geral) — Operar
9. [Investigar um incidente](#passo-09-investigar-um-incidente) — Operar
10. [Custo e FinOps Score](#passo-10-custo-e-finops-score) — Operar
11. [Convidar a equipe](#passo-11-convidar-a-equipe) — Organizar
12. [Proteger sua conta](#passo-12-proteger-sua-conta) — Organizar
13. [Falar com vendas](#passo-13-falar-com-vendas) — Organizar
14. [Perguntas frequentes](#perguntas-frequentes) — Ajuda

Atalhos da capa: criar conta, conectar um cluster, instrumentar a aplicação, ler a Visão Geral, convidar a equipe, falar com vendas.

## Passo 01 — Criar conta ou aceitar convite

- Grupo: Começar
- Tempo: cerca de 3 minutos
- Quem: quem cria a organização ou quem foi convidado
- Palavras-chave: cadastro, cadastrar, criar conta, registrar, signup, inscrição, código de verificação, OTP, convite, convidado, aceitar convite, link do convite, senha nova, primeiro acesso, onboarding

Se a empresa ainda não usa o Immuno Ops, crie a conta: quem cria vira administrador da organização. Se alguém da empresa já usa, peça um convite. Não crie uma segunda conta para a mesma empresa.

### Criar uma conta nova

1. Na tela de login, clique em **Criar conta →**.
2. Informe o e-mail e clique em **Enviar código →**.
3. Digite o código de 6 dígitos que chegou no e-mail. Se não chegou, espere o contador e use **Reenviar**.
4. Preencha nome completo e senha e clique em **Criar minha conta →**. A senha precisa de 8 ou mais caracteres, com maiúscula, minúscula, número e símbolo.
5. Responda as perguntas rápidas: WhatsApp, empresa e cargo, tamanho do cluster e o que dói mais hoje. As opções mostradas para “o que dói mais hoje” são: custo de nuvem fora de controle; alertas contraditórios ou ruído; falta de visão multi-cluster; incidentes sem causa raiz clara.
6. Confira os dados e clique em **Concluir cadastro**. A pessoa entra direto no console.

### Aceitar um convite

1. Abra o link do e-mail de convite.
2. A tela mostra quem convidou, a organização e o papel.
3. Defina nome e senha e clique em **Aceitar e entrar**.

A tela pede e-mail, nome completo e senha.

**Já tem conta?** Se o e-mail já usa o Immuno Ops em outra organização, clique em **Entrar para aceitar**. A nova organização entra na conta que já existe, sem outra senha.

**Convite expirado ou revogado.** Convites têm prazo. Se a tela disser que expirou, foi revogado ou já foi usado, peça ao administrador para reenviar.

## Passo 02 — Entrar no console

- Grupo: Começar
- Tempo: cerca de 1 minuto
- Quem: todo acesso
- Palavras-chave: login, entrar, acessar, acesso, logar, senha, esqueci a senha, recuperar senha, redefinir senha, código do app, organização padrão, trocar de organização, tenant, app.immunoops.io

1. Acesse https://app.immunoops.io/ui/login.
2. Informe e-mail e senha e clique em **Entrar →**.
3. Se a verificação em duas etapas estiver ativa, digite o código de 6 dígitos do app autenticador e clique em **Verificar código**.
4. Se o e-mail tem acesso a mais de uma organização, escolha a organização padrão. Ela abre sempre que a pessoa entra. Dá para trocar pelo seletor de organização, no canto superior direito.

**Esqueceu a senha?** Clique em **Esqueceu sua senha?**, informe o e-mail e siga o link que chegar.

A tela de login também oferece **Criar conta →** para quem ainda não tem conta.

## Passo 03 — Tour de boas-vindas

- Grupo: Começar
- Tempo: cerca de 2 minutos
- Quando: abre sozinho no primeiro acesso
- Palavras-chave: tour, boas-vindas, onboarding, guia inicial, primeiro acesso, pular tour, tutorial, passo a passo

Logo depois do cadastro, e enquanto nenhum serviço estiver conectado, um tour guiado abre por cima da Visão Geral. Ele ilumina cada item do menu e mostra onde conectar os serviços.

| Passo do tour | O que mostra |
| --- | --- |
| 1 | Boas-vindas |
| 2 e 3 | Clusters: onde ficam e como conectar |
| 4 e 5 | Cloud Run: onde fica e como registrar |
| 6 | Instrumentar a aplicação com OpenTelemetry |
| 7 | VMs e a aba Suspensão |

O passo que ilumina Clusters diz: “Seus serviços ficam em Clusters. É aqui que você conecta e acompanha tudo o que o Immuno Ops monitora.” Os botões do tour são **Pular**, **Voltar** e, nesse passo, **Ir para Clusters**.

**Pode pular.** **Pular** ou `Esc` fecham o tour. Ele volta no próximo login enquanto nada estiver conectado. Os passos 05 a 07 deste manual cobrem o mesmo caminho com mais detalhe.

## Passo 04 — Conhecer o console

- Grupo: Começar
- Tempo: cerca de 2 minutos
- Quem: leitura
- Palavras-chave: menu, navegação, menu lateral, Administrar, Investigar, Analisar, Operar, barra do topo, idioma, seletor de organização, agente de IA, papel, permissão, onde fica

O menu da esquerda tem quatro grupos, nesta ordem: **Administrar**, **Investigar**, **Analisar** e **Operar**. Clique no nome do grupo para abrir os itens. O grupo da tela em que a pessoa está já fica aberto. No rodapé do menu ficam o usuário e o atalho para **Segurança da conta**.

A composição dos grupos e o efeito do papel estão em [Mapa do console](#mapa-do-console).

### Agente de IA

O botão ✦ no topo (atalho ⌘J) abre o **Agente IA**, um especialista em SRE, DevOps e FinOps que responde com os dados reais da sua organização. Antes de responder, ele consulta sozinho o que a pergunta precisa:

| Pergunte sobre | O agente consulta |
| --- | --- |
| Saúde, "o que está quebrado", um workload | Workloads: status, taxa de erro, p95, reinícios, réplicas, CPU e memória contra o pedido |
| Custo, "onde gasto mais", tendência | Custo do mês por provedor, serviço e namespace, gasto diário e custo de um workload |
| Incidentes e alertas | Incidentes abertos ou resolvidos e alertas que dispararam |
| "Por que não chegam dados" | VMs, Cloud Run, Lambda, apps e contas de nuvem, com o status de cada um |

Exemplos: "Onde estou gastando mais este mês?", "Tem incidente aberto?", "Como está o checkout-api?", "Minha Lambda não manda dados, por quê?". A resposta vem no idioma da pergunta, com hipótese e próximos passos quando houver.

O agente só vê os dados da organização selecionada e não executa ações: ele lê e explica. Confira a evidência antes de agir.

## Passo 05 — Conectar um cluster Kubernetes

- Grupo: Conectar
- Tempo: cerca de 5 minutos
- Quem: Administrador ou Operador
- Requisito: `helm` no cluster
- Palavras-chave: cluster, kubernetes, k8s, conectar cluster, agente, agent, instalar agente, helm, kubectl, chart, token, GKE, EKS, AKS, pendente, conectado, desconectado, heartbeat, status do cluster, namespace

É o passo mais importante. Com o agent instalado, métricas, logs, traces e o inventário de workloads começam a chegar sozinhos.

1. No menu, vá em **Administrar › Clusters** e clique em **Conectar cluster**.
2. Em **Escolha a nuvem**, selecione o provedor.
3. Dê um nome ao cluster, escolha o ambiente (produção, staging ou outro) e clique em **Registrar cluster**.
4. A janela **Instale o agent** mostra um comando `helm` já com o token de conexão. Clique em **Copiar comando** e rode num terminal com acesso ao cluster.
5. Clique em **Concluir**. Em alguns minutos o status muda de Pendente para Conectado.

O comando exibido na janela segue a forma `helm upgrade --install` do agent, com `connectionToken` preenchido. Copie o comando do console. Não use um comando montado fora dessa janela: o token e o endereço do chart vêm de lá.

**Guarde o comando com cuidado.** Ele carrega um token de conexão da organização. Não cole em chats públicos nem em repositórios.

### Como ler o status do cluster

| Status | O que significa | O que fazer |
| --- | --- | --- |
| Pendente | Cluster registrado, mas o agent nunca mandou dado. | Confira se o comando rodou sem erro e se os pods do agent estão `Running`. |
| Conectado | O agent está mandando dados. | Nada. Os dados já aparecem em Investigar. |
| Desconectado | O último sinal tem mais de 5 minutos. | Verifique o agent no cluster: pods, rede de saída e token. |

Comparado com os outros tipos de recurso: no cluster se instala um agent com Helm, e ele traz workloads, métricas, logs e traces. Serverless e VM estão no passo 06.

## Passo 06 — Serverless e VMs

- Grupo: Conectar
- Tempo: cerca de 3 minutos cada
- Quem: quem usa esses serviços
- Palavras-chave: serverless, Cloud Run, AWS Lambda, Lambda, função, VM, máquina virtual, EC2, Compute Engine, GCE, Azure VM, on-prem, multicloud, GCP, AWS, Azure, Google Cloud, projeto GCP, conta AWS, assinatura Azure, região, zona, docker, agente na VM, IMDS, IMDSv2, metadados, suspensão, desligar VM, horário, conta de nuvem, billing export, BigQuery, dataViewer, permissão, monitoring.viewer, logging.viewer, compute.instances, immunoOpsVmPower, domínio, allowedPolicyMemberDomains, s3:ListBucket, s3:GetObject, immunoops-billing-reader, AssumeRoleWithWebIdentity, oaud

O Immuno Ops é multicloud: VMs em GCP, AWS, Azure ou on-prem, e serviços serverless no GCP (Cloud Run) e na AWS (Lambda), na mesma tela.

VM e Serverless exigem ao menos uma conta de nuvem já cadastrada. Cloud Run exige uma conta GCP; Lambda exige uma conta AWS. Sem essa conta, **Conectar VM** e **Conectar serviço Serverless** abrem **Adicionar conta** em vez do formulário do recurso. O cadastro da conta está em [Conta de nuvem](#conta-de-nuvem). Os papéis no projeto do recurso estão em [Permissões de VM e Cloud Run](#permissoes-de-vm-e-cloud-run).

| Tipo | Nuvem | O que instalar | O que chega |
| --- | --- | --- | --- |
| Cluster | GKE, EKS, AKS ou outro Kubernetes | Agent via Helm no cluster. | Workloads, métricas, logs e traces. |
| Cloud Run | GCP | Nada. | Métricas e logs do Cloud Monitoring e do Cloud Logging do projeto. |
| Lambda | AWS | O SDK na função (passo 07). | Métricas, logs e eventos que a função envia. |
| VM | GCP, AWS, Azure ou on-prem | Agente no host (container Docker ou binário). | CPU, memória, logs do host e dos containers, tipo de máquina e sistema. |

### Serverless: Cloud Run (GCP)

1. Vá em **Infraestrutura › Serverless** e clique em **Conectar serviço Serverless**.
2. Em **Provedor**, escolha **GCP**.
3. Informe o nome do serviço (exatamente como no Cloud Run), o ID do projeto GCP e a região. Clique em **Registrar**.

### Serverless: Lambda (AWS)

1. Vá em **Infraestrutura › Serverless** e clique em **Conectar serviço Serverless**.
2. Em **Provedor**, escolha **AWS**.
3. Informe o nome da função, a região (por exemplo `us-east-1`) e o ID da conta AWS, com 12 dígitos. O ID vem preenchido com a conta AWS da organização. Clique em **Registrar**.
4. Instrumente a função com o SDK, como no [passo 07](#passo-07-instrumentar-a-aplicacao), e chame `flush()` antes de a função retornar. Veja [Lambda com o SDK](#lambda-com-o-sdk).

O SDK, se a função o usar, envia os dados para `app.immunoops.io` com a chave de conexão e não pede papel na função. Numa VPC sem internet, a função precisa de um NAT Gateway para essa saída.

As métricas da função (invocações, duração e execuções simultâneas) e os logs do grupo `/aws/lambda/<função>` vêm do CloudWatch, lidos pelo worker com o papel `immunoops-billing-reader`. Esse papel é criado com um clique em [Conta AWS](#conta-aws) (**Conectar à AWS**). Azure Functions aparece como **em breve**.

**"Falta conectar a conta AWS".** É o que a aba de logs da função mostra quando a conta AWS dela ainda não tem o papel. Clique em **Conectar à AWS** ali mesmo. Depois de criada a pilha no CloudFormation, o Immuno Ops tenta de novo em até 15 minutos e os dados começam a chegar.

### VMs

O cadastro pede só o nome. O agente descobre sozinho onde a VM roda, lendo o serviço de metadados da nuvem: GCP (Compute Engine), AWS (EC2, IMDSv2) ou Azure. Fora de uma nuvem, a VM aparece como **On-prem**.

1. Vá em **Infraestrutura › VMs** e clique em **Conectar VM**.
2. Dê um nome à VM, do jeito que você quer vê-la no console. Clique em **Registrar**.
3. Copie o comando em **Instalar o agente**, rode na VM e clique em **Concluir**.

Quando o agente manda o primeiro dado, o cartão da VM mostra o provedor (GCP, AWS, Azure ou On-prem), a zona ou região, a conta (projeto GCP, conta AWS ou assinatura Azure) e o tipo de máquina.

| Nuvem | O que o agente detecta | Cuidado |
| --- | --- | --- |
| GCP | Projeto, zona, nome da instância e tipo de máquina | Nada. |
| AWS (EC2) | Conta, região, zona, ID da instância e tipo de instância | Com o agente em Docker, o IMDSv2 precisa de `HttpPutResponseHopLimit` 2 (o padrão do EC2 é 1 e bloqueia o container). |
| Azure | Assinatura, região, zona, nome e tamanho da VM | Nada. |
| On-prem | Nada de nuvem: só host, sistema e uso | Nada. |

Para ajustar o hop limit numa EC2 existente:

```
aws ec2 modify-instance-metadata-options --instance-id ID_DA_INSTANCIA --http-tokens required --http-put-response-hop-limit 2
```

**Economize com a aba Suspensão.** Em **VMs › Suspensão**, use **Nova política** para ligar e desligar VMs por horário, no Compute Engine ou na EC2. No GCP, o acesso está em [Permissões da VM](#permissoes-da-vm). Na EC2, a nuvem da política é AWS e as ações estão em [Ligar e desligar a VM](#ligar-e-desligar-a-vm).

### Permissões de VM e Cloud Run

Quem libera no Google Cloud é alguém que altera o IAM do projeto do recurso. A conta é a do worker, `aegis-worker@immunoops-prod.iam.gserviceaccount.com`, a mesma do billing em [Conta de nuvem](#conta-de-nuvem). Se a organização bloqueia contas de fora, inclua o customer ID **C011tigzz** neste projeto também, do mesmo jeito descrito em [Se a organização bloqueia contas de fora](#se-a-organizacao-bloqueia-contas-de-fora).

#### Permissões do Cloud Run

O worker lê o Cloud Monitoring e o Cloud Logging do projeto. Nada é instalado no container. No projeto GCP do serviço, a conta do worker precisa de dois papéis de leitura:

| Papel | ID | Onde | Para quê |
| --- | --- | --- | --- |
| Monitoring Viewer | `roles/monitoring.viewer` | projeto do serviço | Requisições, latência, CPU, memória e instâncias |
| Logs Viewer | `roles/logging.viewer` | projeto do serviço | Logs de stdout, stderr e requisições |

```
gcloud projects add-iam-policy-binding PROJETO --member="serviceAccount:aegis-worker@immunoops-prod.iam.gserviceaccount.com" --role="roles/monitoring.viewer"
gcloud projects add-iam-policy-binding PROJETO --member="serviceAccount:aegis-worker@immunoops-prod.iam.gserviceaccount.com" --role="roles/logging.viewer"
```

Troque `PROJETO` pelo ID do projeto do serviço. No exemplo da tela, é `acme-prod-123`. Rode os dois num terminal autenticado no Google Cloud.

Sem o papel de Monitoring, os números do serviço ficam sem leitura. Sem o papel de Logging, a tela de logs cita `roles/logging.viewer` e o projeto. Se a tela disser que a coleta de logs ainda não está habilitada para a conta, peça ao suporte para habilitá-la: o papel prepara o acesso, e a coleta só começa depois dessa liberação.

#### Permissões da VM

Métricas e logs da VM saem do agente instalado no host. Quem instala precisa rodar o comando na máquina, e a VM precisa de saída HTTPS para `app.immunoops.io`.

A aba **Suspensão** liga e desliga a instância no Compute Engine. Para isso, a conta do worker precisa listar, iniciar e parar instâncias nesse projeto. Crie um papel só com essas três permissões e conceda à conta:

```
gcloud iam roles create immunoOpsVmPower --project=PROJETO --title="Immuno Ops VM power" --permissions=compute.instances.list,compute.instances.start,compute.instances.stop --stage=GA
gcloud projects add-iam-policy-binding PROJETO --member="serviceAccount:aegis-worker@immunoops-prod.iam.gserviceaccount.com" --role="projects/PROJETO/roles/immunoOpsVmPower"
```

Troque `PROJETO` pelo ID do projeto da VM. O papel cobre listar, iniciar e parar. A coleta do agente segue com a instalação no host.

### CloudWatch da Lambda

Quem libera na AWS altera o papel `immunoops-billing-reader`, o mesmo que o worker já assume para o billing. A função não precisa de cluster nem de EC2. Nada é instalado na função. O worker lê o CloudWatch.

| Ação | Onde | Para quê |
| --- | --- | --- |
| `cloudwatch:GetMetricData` | namespace `AWS/Lambda` | Invocações, p95 da duração e pico de execuções simultâneas |
| `cloudwatch:ListMetrics` | namespace `AWS/Lambda` | Encontrar as métricas da função |
| `logs:FilterLogEvents` | grupo `/aws/lambda/*` | Ler as linhas da função na aba de logs |

O worker lê o grupo `/aws/lambda/` seguido do nome da função. As linhas aparecem na aba de logs. A leitura do export continua no mesmo papel, em [Permissões no papel de leitura](#permissoes-no-papel-de-leitura).

```
aws iam put-role-policy --role-name immunoops-billing-reader --policy-name immunoops-lambda-cloudwatch --policy-document '{"Version":"2012-10-17","Statement":[{"Effect":"Allow","Action":["cloudwatch:GetMetricData","cloudwatch:ListMetrics"],"Resource":"*","Condition":{"StringEquals":{"cloudwatch:namespace":"AWS/Lambda"}}},{"Effect":"Allow","Action":["logs:FilterLogEvents"],"Resource":["arn:aws:logs:*:*:log-group:/aws/lambda/*","arn:aws:logs:*:*:log-group:/aws/lambda/*:*"]}]}'
```

Sem as ações de métrica, a função fica cadastrada e os números do CloudWatch não chegam. Sem `logs:FilterLogEvents`, a aba de logs fica sem as linhas da função.

### Ligar e desligar a VM

Métricas e logs da EC2 saem do agente instalado no host. A conta não precisa de cluster. Ligar e desligar usa um papel **separado e opcional**, `immunoops-power`: quem só quer custo, métricas e logs não dá permissão de escrita ao Immuno Ops.

| Papel | Ação | Onde | Para quê |
| --- | --- | --- | --- |
| `immunoops-billing-reader` | `ec2:DescribeInstances` | conta | Listar a instância |
| `immunoops-power` | `ec2:StartInstances` | instâncias da conta | Ligar |
| `immunoops-power` | `ec2:StopInstances` | instâncias da conta | Desligar |

O `immunoops-power` tem a mesma confiança do papel de leitura: a conta `aegis-worker@immunoops-prod.iam.gserviceaccount.com` no Google e a audiência da sua organização (a mesma `AUDIENCIA` de [Conta AWS](#conta-aws)). No GCP, as permissões equivalentes estão em [Permissões da VM](#permissoes-da-vm).

```
aws iam put-role-policy --role-name immunoops-billing-reader --policy-name immunoops-ec2-describe --policy-document '{"Version":"2012-10-17","Statement":[{"Effect":"Allow","Action":"ec2:DescribeInstances","Resource":"*"}]}'

aws iam create-role --role-name immunoops-power --assume-role-policy-document '{"Version":"2012-10-17","Statement":[{"Effect":"Allow","Principal":{"Federated":"accounts.google.com"},"Action":"sts:AssumeRoleWithWebIdentity","Condition":{"StringEquals":{"accounts.google.com:sub":"101814243271068167435","accounts.google.com:aud":"101814243271068167435","accounts.google.com:oaud":"'$AUDIENCIA'"}}}]}'
aws iam put-role-policy --role-name immunoops-power --policy-name immunoops-vm-power --policy-document '{"Version":"2012-10-17","Statement":[{"Effect":"Allow","Action":["ec2:StartInstances","ec2:StopInstances"],"Resource":"arn:aws:ec2:*:*:instance/*"}]}'
```

Sem o papel `immunoops-power`, a política de Suspensão não liga nem desliga a EC2; o resto continua funcionando. No CloudTrail, cada ação aparece com a sessão `immunoops-<ID da organização>`.

### Conta de nuvem

- Quem cadastra no Immuno Ops: Administrador
- Quem libera no Google Cloud: alguém que altera o IAM do dataset de billing e, se a organização restringe domínios, a política do projeto
- Onde: **Infraestrutura**, aba **Contas** (`app.immunoops.io/ui/infrastructure?tab=contas`)
- Palavras-chave: conta de nuvem, adicionar conta, billing export, BigQuery, tabela, projeto.dataset.tabela, dataViewer, aegis-worker, customer ID, C011tigzz, domínio, Sem tabela, testar conexão

A conta é somente leitura do billing. Nada é criado nem alterado nos recursos da nuvem. Uma conta GCP traz o custo do projeto. Uma conta AWS ou Azure conta como “ao menos uma conta” para uma VM, mas não substitui a conta GCP do Cloud Run.

#### Ativar o export no Google Cloud

1. No Google Cloud, abra a conta de faturamento que paga os projetos e vá em **Exportação de faturamento › Exportação para o BigQuery**.
2. Ative o export de **custo de uso** no modo **detalhado** (ou padrão com recursos: `gcp_billing_export_resource_v1_…`). A tabela do Immuno Ops é essa, com as colunas de custo do Cloud Billing. Não use uma tabela no formato FOCUS.
3. Escolha o projeto e o dataset. O Google cria uma tabela com nome no formato `gcp_billing_export_resource_v1_…`.
4. Anote o identificador completo, no formato `projeto.dataset.tabela`. Exemplo ilustrativo: `acme-billing.billing_export.gcp_billing_export_resource_v1_012345_ABCDE1_234567`.

A primeira linha do export pode levar algumas horas para aparecer. O teste de conexão aceita a tabela vazia; o status só fica em dia quando chega custo recente.

#### Cadastrar no Immuno Ops

1. Vá em **Infraestrutura**, aba **Contas**, e clique em **Adicionar conta**.
2. Escolha **GCP**.
3. Informe o **ID do projeto** e a **Tabela do billing export** (`projeto.dataset.tabela`). O nome é opcional.
4. Clique em **Adicionar conta**.

**Conta GCP antiga, sem tabela.** O cartão mostra **Sem tabela**. Abra a conta, informe a tabela e clique em **Salvar tabela**. A partir daí valem os mesmos comandos e o mesmo teste.

#### Permissões no dataset

No dataset do export, a conta do worker precisa do papel **BigQuery Data Viewer** (`roles/bigquery.dataViewer`). É leitura. Não conceda papéis de edição, de administração do projeto nem de faturamento.

| Conta de serviço | Papel | Onde |
| --- | --- | --- |
| `aegis-worker@immunoops-prod.iam.gserviceaccount.com` | `roles/bigquery.dataViewer` | dataset do export |

Depois de salvar a tabela, a conta no console mostra o comando já com o dataset preenchido. Copie de lá. A forma, com `PROJETO:DATASET` no lugar do dataset (os dois primeiros trechos de `projeto.dataset.tabela`, separados por `:`), é:

```
bq add-iam-policy-binding --member="serviceAccount:aegis-worker@immunoops-prod.iam.gserviceaccount.com" --role="roles/bigquery.dataViewer" PROJETO:DATASET
```

Rode o comando num terminal autenticado no Google Cloud, com permissão para alterar o IAM desse dataset. No exemplo ilustrativo, `PROJETO:DATASET` é `acme-billing:billing_export`.

#### Permissão para executar queries

A conta do worker também precisa submeter queries ao BigQuery no projeto de plataforma (`immunoops-prod`). Ela precisa do papel **BigQuery Job User** (`roles/bigquery.jobUser`) nesse projeto. Sem isso, o teste de conexão falha com `export_query_failed`.

```
gcloud projects add-iam-policy-binding immunoops-prod --member="serviceAccount:aegis-worker@immunoops-prod.iam.gserviceaccount.com" --role="roles/bigquery.jobUser" --condition=None
```

Rode o comando com uma conta que tenha permissão para alterar o IAM do projeto `immunoops-prod`.

#### Se a organização bloqueia contas de fora

Algumas organizações restringem quem pode receber papel no IAM (`iam.allowedPolicyMemberDomains`, “Domain restricted sharing”). O comando acima falha com a mensagem de que o membro não pertence a um cliente permitido.

Inclua o customer ID **C011tigzz** como exceção dessa política. A organização é **immunoops.io**, id **424787708642**. Faça a exceção só no projeto do billing, não na organização inteira.

1. No Google Cloud, abra **IAM e administrador › Políticas da organização** no projeto do billing.
2. Abra a restrição **Domain restricted sharing** (`constraints/iam.allowedPolicyMemberDomains`).
3. Adicione **C011tigzz** aos valores permitidos, sem remover os que já existem.
4. Salve e rode de novo o comando `bq`.

#### Testar

1. Volte na conta, em **Infraestrutura › Contas**, e clique em **Testar conexão**.
2. O teste lê uma linha da tabela. Sem erro de permissão, a conta fica pendente até chegar custo, ou em dia quando o custo recente aparece.

| O que a tela diz | O que fazer |
| --- | --- |
| Sem tabela | Informe `projeto.dataset.tabela` e salve. |
| A organização bloqueia contas de fora | Inclua o customer ID C011tigzz e teste de novo. |
| Não foi possível ler o export | Confira se o comando `bq` rodou no dataset certo, e se `roles/bigquery.jobUser` foi concedido no projeto immunoops-prod. Teste de novo. |
| Custo chegou, mas logs mostram erro de permissão | Conceda `roles/logging.viewer` ao dataset ou ao projeto, e teste de novo. |

**A última conta fica.** Enquanto houver VM ou Cloud Run, a última conta de nuvem não pode ser removida. Enquanto houver Cloud Run, a última conta GCP não pode ser removida. Cadastre a conta que continua valendo e só então remova a outra.

### Conta AWS

- Quem cadastra no Immuno Ops: Administrador
- Quem faz na AWS: alguém com permissão de IAM e de Billing na conta (no console, **Billing and Cost Management**)
- Onde: **Infraestrutura**, aba **Contas** (`app.immunoops.io/ui/infrastructure?tab=contas`)
- Palavras-chave: Conectar à AWS, um clique, CloudFormation, pilha, stack, immunoops-integration, role_missing, Falta conectar a conta AWS, AWS, conta AWS, Amazon, EKS, EC2, Lambda, ID da conta, 12 dígitos, billing AWS, Data Exports, FOCUS, S3, CloudFormation, AWS Organizations, conta pagadora, conta-membro, s3:ListBucket, s3:GetObject, immunoops-billing-reader, AssumeRoleWithWebIdentity, oaud, chave de acesso

A conta AWS é o que libera, para a sua organização, os recursos da AWS no Immuno Ops. O que já funciona e o que ainda vem:

| O que | Situação |
| --- | --- |
| Cluster EKS | Disponível. Agent via Helm, como no [passo 05](#passo-05-conectar-um-cluster-kubernetes). |
| VM EC2 | Disponível. O agente detecta conta, região e tipo de instância ([VMs](#vms)). |
| Lambda | Disponível. Métricas e logs pelo CloudWatch depois de **Conectar à AWS**; ou pelo SDK, sem nada na conta. |
| Custo da AWS (billing) | O export FOCUS é criado no **Conectar à AWS**. A leitura do custo no console está **em breve**; o histórico já começa a contar. |

#### Passo 1 — Descobrir o ID da conta

No console da AWS, clique no nome da conta no canto superior direito e copie o **Account ID**, com 12 dígitos. Pelo terminal:

```
aws sts get-caller-identity --query Account --output text
```

#### Passo 2 — Cadastrar no Immuno Ops

1. Vá em **Infraestrutura**, aba **Contas**, e clique em **Adicionar conta**.
2. Escolha **AWS**.
3. Informe o **ID da conta (12 dígitos)**. O nome é opcional.
4. Clique em **Adicionar conta**.

A conta aparece como pendente: o custo da AWS ainda não é lido. Isso não impede EKS, EC2 e Lambda.

#### Passo 2b — Conectar à AWS (um clique)

É como nas outras ferramentas de observabilidade: você não escreve IAM.

1. No cartão da conta AWS, em **Infraestrutura › Contas**, clique em **Conectar à AWS**.
2. O console da AWS abre no **CloudFormation**, na região `us-east-1`, com a pilha `immunoops-integration` já preenchida para a sua organização. Entre com um usuário que possa criar papéis IAM.
3. Marque a confirmação de criação de recursos IAM e clique em **Criar pilha**. Leva cerca de 2 minutos.

A pilha cria:

| Recurso | Para quê |
| --- | --- |
| Papel `immunoops-billing-reader` | Leitura: CloudWatch e logs da Lambda, listar EC2 e funções, ler o export de custo. Sem chave de acesso: confia no Immuno Ops só com a audiência da sua organização |
| Bucket `immunoops-focus-CONTA` e export `immunoops-focus` | Custo diário no formato FOCUS 1.2. O primeiro arquivo chega em até 24 horas |
| Papel `immunoops-power` (opcional, `EnablePower=true`) | Ligar e desligar EC2 nas políticas de Suspensão |

Para conferir antes de criar, abra **Detalhes da conexão** no cartão: aparecem a audiência e o ARN do papel. Para remover o acesso, apague a pilha no CloudFormation; o bucket do export fica guardado.

Os passos 4 e as seções seguintes descrevem a mesma configuração feita à mão, para quem não pode usar o CloudFormation.

#### Passo 3 — Conectar o que roda nessa conta

- **EKS:** **Infraestrutura › Clusters › Conectar cluster**, escolha **AWS** e rode o comando `helm` no cluster ([passo 05](#passo-05-conectar-um-cluster-kubernetes)).
- **EC2:** **Infraestrutura › VMs › Conectar VM**, dê um nome e rode o comando na instância. Com o agente em Docker, ajuste o hop limit do IMDSv2 para 2 ([VMs](#vms)).
- **Lambda:** **Infraestrutura › Serverless › Conectar serviço Serverless**, provedor **AWS**. Métricas e logs chegam pelo CloudWatch depois do passo 2b; o SDK ([Lambda com o SDK](#lambda-com-o-sdk)) é opcional, para métricas de negócio.

Os passos de EKS, EC2 e SDK não dão ao Immuno Ops acesso à conta: os dados saem do agente ou do SDK. O acesso de leitura à conta vem só da pilha do passo 2b.

#### Passo 4 — Preparar o billing à mão (sem CloudFormation)

O custo da AWS vai ser lido de um **export de faturamento no formato FOCUS** que você cria na sua conta, num bucket seu. A AWS só gera esse export daqui para a frente: criar agora faz o histórico já estar lá quando a leitura for liberada. Nada é compartilhado com o Immuno Ops neste passo.

Rode num terminal autenticado na conta, trocando `CONTA` pelo ID de 12 dígitos:

```
CONTA=123456789012
BUCKET=immunoops-focus-$CONTA

aws s3api create-bucket --bucket $BUCKET --region us-east-1
aws s3api put-public-access-block --bucket $BUCKET --public-access-block-configuration BlockPublicAcls=true,IgnorePublicAcls=true,BlockPublicPolicy=true,RestrictPublicBuckets=true

aws s3api put-bucket-policy --bucket $BUCKET --policy '{"Version":"2012-10-17","Statement":[{"Effect":"Allow","Principal":{"Service":["billingreports.amazonaws.com","bcm-data-exports.amazonaws.com"]},"Action":["s3:PutObject","s3:GetBucketPolicy"],"Resource":["arn:aws:s3:::'$BUCKET'","arn:aws:s3:::'$BUCKET'/*"],"Condition":{"StringLike":{"aws:SourceAccount":"'$CONTA'"}}}]}'
```

Depois crie o export no console da AWS:

1. Abra **Billing and Cost Management › Data Exports** e clique em **Create**.
2. Escolha **FOCUS 1.2 with AWS columns** e o nome `immunoops-focus`.
3. Em **Data export delivery options**, formato **Parquet**, atualização diária e **Overwrite existing data export file**.
4. Em **Data export storage settings**, o bucket `immunoops-focus-CONTA` e o prefixo `focus`.
5. Clique em **Create**. O primeiro arquivo pode levar até 24 horas.

Use exatamente essa política do bucket. Uma variação com `aws:SourceArn`, que aparece em exemplos da AWS, faz a criação do export falhar com `S3 bucket permission validation failed`.

**Conta dentro de uma AWS Organization.** Uma conta-membro pode criar o próprio export, com o custo só dela. Para várias contas de uma vez, o export vai na conta pagadora (management); a seleção de contas por organização no Immuno Ops ainda não está disponível.

#### Permissões no papel de leitura

Quem libera na AWS é alguém que altera o IAM da conta. O papel é somente leitura do export. Nada é criado nem alterado nos recursos da nuvem. Não crie usuário IAM e não gere chave de acesso.

| Ação | Onde | Para quê |
| --- | --- | --- |
| `s3:ListBucket` | bucket do export, só o prefixo `focus/` | Listar os arquivos do export |
| `s3:GetObject` | objetos `focus/*` nesse bucket | Ler o manifesto e os arquivos Parquet |

A confiança do papel aponta para o Google (`accounts.google.com`), para a conta `aegis-worker@immunoops-prod.iam.gserviceaccount.com`. São três condições, todas obrigatórias:

| Condição | Valor |
| --- | --- |
| `accounts.google.com:sub` | `101814243271068167435` |
| `accounts.google.com:aud` | `101814243271068167435` |
| `accounts.google.com:oaud` | `immunoops-tenant-` + id da organização |

O número é o ID único dessa conta de serviço. A condição `oaud` é exclusiva da organização: com outra audiência, a AWS responde `AccessDenied`.

Não conceda `s3:PutObject`, `s3:*`, `AdministratorAccess` nem papel de faturamento. Essas duas ações cobrem a leitura.

Rode num terminal autenticado na conta. `CONTA` é o ID de 12 dígitos. `BUCKET` é o bucket do passo 4. Troque `AUDIENCIA` por `immunoops-tenant-` seguido do id da organização. A audiência aparece em **Detalhes da conexão**, no cartão da conta AWS.

```
CONTA=123456789012
BUCKET=immunoops-focus-$CONTA
AUDIENCIA=immunoops-tenant-ID_DA_ORGANIZACAO

aws iam create-role --role-name immunoops-billing-reader --assume-role-policy-document '{"Version":"2012-10-17","Statement":[{"Effect":"Allow","Principal":{"Federated":"accounts.google.com"},"Action":"sts:AssumeRoleWithWebIdentity","Condition":{"StringEquals":{"accounts.google.com:sub":"101814243271068167435","accounts.google.com:aud":"101814243271068167435","accounts.google.com:oaud":"'$AUDIENCIA'"}}}]}'

aws iam put-role-policy --role-name immunoops-billing-reader --policy-name immunoops-billing-read --policy-document '{"Version":"2012-10-17","Statement":[{"Effect":"Allow","Action":"s3:ListBucket","Resource":"arn:aws:s3:::'$BUCKET'","Condition":{"StringLike":{"s3:prefix":["focus/","focus/*"]}}},{"Effect":"Allow","Action":"s3:GetObject","Resource":"arn:aws:s3:::'$BUCKET'/focus/*"}]}'
```

O papel fica `arn:aws:iam::CONTA:role/immunoops-billing-reader`. Guarde esse ARN.

| O que acontece | O que fazer |
| --- | --- |
| A criação do export falha com `S3 bucket permission validation failed` | A política do bucket tem `aws:SourceArn`. Use só `aws:SourceAccount`, como no passo 4. |
| `AccessDenied` ao assumir o papel | Confira `sub`, `aud` e `oaud`. O `oaud` tem que ser o da sua organização. |
| Pedido de chave de acesso ou de usuário IAM | Não crie. O acesso é só por esse papel. |

## Passo 07 — Instrumentar a aplicação

- Grupo: Conectar
- Tempo: cerca de 10 minutos
- Quem: time de desenvolvimento
- Opcional
- Palavras-chave: OpenTelemetry Collector, OTLP, OTLP/HTTP, OTEL_EXPORTER_OTLP_ENDPOINT, OTEL_EXPORTER_OTLP_HEADERS, Go, Golang, Python, Java, .NET, traces, spans, @immunoops/sdk, npm install, instrumentação, instrumentar, SDK, aplicação, app, código, chave de conexão, apiKey, API key, OpenTelemetry, OTel, Node.js, npm, métricas da aplicação, eventos de negócio, conectada, aguardando primeiro dado, sem dados recentes

O agent já mostra a saúde da infraestrutura. Para ver métricas de negócio, logs, traces e eventos de dentro do código, conecte a aplicação. Há dois caminhos, e os dois usam a mesma chave de conexão:

| Caminho | Para quem | O que chega |
| --- | --- | --- |
| [SDK Node.js](#sdk-nodejs) (`@immunoops/sdk`) | Apps em Node.js e TypeScript | Métricas, logs e eventos de negócio, com mascaramento de dados pessoais |
| [OpenTelemetry](#opentelemetry) | Qualquer linguagem (Go, Python, Java, .NET, Ruby, PHP…) ou um OpenTelemetry Collector | Traces, métricas e logs no padrão OTLP |

### Gerar a chave de conexão

1. Vá em **Infraestrutura › Instrumentação** e clique em **Conectar via SDK**.
2. Dê um nome à aplicação e clique em **Gerar chave**.
3. Copie a chave: ela só aparece inteira agora. Se perder, gere outra nos detalhes da conexão; a anterior deixa de valer.

O status muda sozinho de **Aguardando primeiro dado** para **Conectada** quando o primeiro dado chega. **Sem dados recentes** quer dizer que a aplicação parou de enviar.

Guarde a chave numa variável de ambiente ou num cofre de segredos, nunca no código.

### SDK Node.js

O pacote é público no npm: [`@immunoops/sdk`](https://www.npmjs.com/package/@immunoops/sdk). Node.js 18 ou mais novo.

```
npm install @immunoops/sdk
```

```js
import immuno from '@immunoops/sdk';

immuno.init({
  serviceName: 'checkout-api',
  apiKey: process.env.IMMUNOOPS_API_KEY,
  endpoint: 'https://app.immunoops.io/agent/ingest',
});

immuno.event('app_started', {});
```

Funciona igual com `require`: `const immuno = require('@immunoops/sdk');`. As chamadas principais:

| Chamada | Para quê |
| --- | --- |
| `immuno.increment('pedidos', 1, { status: 'ok' })` | Contador |
| `immuno.metric('fila.tamanho', 42)` | Valor medido (gauge) |
| `await immuno.time('checkout.duration_ms', () => processar())` | Mede a duração de uma função |
| `immuno.info('mensagem', { contexto })` / `warn` / `error` | Logs da aplicação |
| `immuno.event('pedido.pago', { total: 49.9 })` | Evento de negócio |
| `await immuno.flush()` | Envia o que está na fila agora (obrigatório antes de encerrar o processo e na Lambda) |

O SDK manda os dados em lotes a cada 5 segundos e mascara CPF, número de cartão, senhas e tokens antes de enviar. Para ver no terminal o que foi enviado, use `debug: true` no `init`.

### OpenTelemetry

Qualquer aplicação instrumentada com OpenTelemetry, em qualquer linguagem, envia direto para o Immuno Ops pelo protocolo OTLP/HTTP. Não precisa do SDK do Immuno Ops nem de agent. Use a mesma chave de conexão.

| Configuração | Valor |
| --- | --- |
| Endpoint OTLP/HTTP | `https://app.immunoops.io/agent/ingest` |
| Caminhos | `/v1/traces`, `/v1/metrics` e `/v1/logs` (o exporter acrescenta sozinho) |
| Cabeçalho | `Authorization: Bearer <chave de conexão>` |
| Protocolo | `http/protobuf` (padrão) ou `http/json`. gRPC não é aceito |

O jeito mais simples é configurar pelas variáveis de ambiente padrão do OpenTelemetry, que todas as linguagens leem:

```
export OTEL_SERVICE_NAME=checkout-api
export OTEL_EXPORTER_OTLP_ENDPOINT=https://app.immunoops.io/agent/ingest
export OTEL_EXPORTER_OTLP_PROTOCOL=http/protobuf
export OTEL_EXPORTER_OTLP_HEADERS="Authorization=Bearer SUA_CHAVE"
```

#### Exemplo em Go

Com as variáveis acima, o exporter HTTP do OpenTelemetry já sabe para onde enviar e com qual chave:

```go
import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

func setupOTel(ctx context.Context) (func(context.Context) error, error) {
	traceExp, err := otlptracehttp.New(ctx) // lê OTEL_EXPORTER_OTLP_*
	if err != nil {
		return nil, err
	}
	metricExp, err := otlpmetrichttp.New(ctx)
	if err != nil {
		return nil, err
	}
	tp := sdktrace.NewTracerProvider(sdktrace.WithBatcher(traceExp))
	mp := sdkmetric.NewMeterProvider(sdkmetric.WithReader(sdkmetric.NewPeriodicReader(metricExp)))
	otel.SetTracerProvider(tp)
	otel.SetMeterProvider(mp)
	return func(ctx context.Context) error { // chame ao encerrar
		_ = tp.Shutdown(ctx)
		return mp.Shutdown(ctx)
	}, nil
}
```

Logs seguem o mesmo padrão com `otlploghttp`. Para HTTP e gRPC do próprio Go, os pacotes `otelhttp` e `otelgrpc` criam os spans de cada requisição.

#### Outras linguagens

Com as mesmas variáveis de ambiente:

- **Python:** `pip install opentelemetry-distro opentelemetry-exporter-otlp && opentelemetry-bootstrap -a install`, depois rode a aplicação com `opentelemetry-instrument python app.py`.
- **Java:** baixe o `opentelemetry-javaagent.jar` e rode com `java -javaagent:opentelemetry-javaagent.jar -jar app.jar`.
- **Node.js com OpenTelemetry** (em vez do SDK do Immuno Ops): `npm install @opentelemetry/auto-instrumentations-node` e rode com `node --require @opentelemetry/auto-instrumentations-node/register app.js`.
- **.NET, Ruby, PHP:** o SDK de OpenTelemetry da linguagem com o exporter OTLP/HTTP.

#### Pelo OpenTelemetry Collector

Se você já tem um Collector, acrescente um exporter OTLP/HTTP e coloque-o nos pipelines. Nas versões 0.165 em diante o nome é `otlp_http`; nas anteriores, `otlphttp`:

```yaml
exporters:
  otlphttp/immunoops:
    endpoint: https://app.immunoops.io/agent/ingest
    headers:
      Authorization: "Bearer ${env:IMMUNOOPS_TOKEN}"

service:
  pipelines:
    traces:  { receivers: [otlp], processors: [batch], exporters: [otlphttp/immunoops] }
    metrics: { receivers: [otlp], processors: [batch], exporters: [otlphttp/immunoops] }
    logs:    { receivers: [otlp], processors: [batch], exporters: [otlphttp/immunoops] }
```

Rode o Collector com a variável `IMMUNOOPS_TOKEN` com a chave de conexão. O exporter comprime em gzip por padrão; o Immuno Ops aceita.

#### Se não chega nada

| O que aparece | O que fazer |
| --- | --- |
| `401` no exporter | A chave está errada ou foi trocada por uma nova. Gere outra nos detalhes da conexão e use a última |
| Endpoint com `/v1/traces` duas vezes | Use `https://app.immunoops.io/agent/ingest` sem caminho em `OTEL_EXPORTER_OTLP_ENDPOINT`; o exporter acrescenta |
| Exporter gRPC (porta 4317) | Troque para OTLP/HTTP (`http/protobuf`) |
| Status continua **Aguardando primeiro dado** | Confira se a aplicação tem saída HTTPS para `app.immunoops.io` |

**Atalho para Cloud Run e VMs.** Na linha de um serviço já registrado, **Instrumentar minha aplicação** gera a chave e o snippet prontos para colar.

### Lambda com o SDK

A Lambda congela o processo assim que a função retorna. Por isso, envie os dados antes de retornar:

```js
const immuno = require("@immunoops/sdk");

immuno.init({
  apiKey: process.env.IMMUNOOPS_API_KEY,
  endpoint: "https://app.immunoops.io/agent/ingest",
  serviceName: "checkout",
  environment: "production",
});

exports.handler = async (event) => {
  await immuno.time("checkout.duration_ms", () => processar(event));
  immuno.increment("checkout.orders", 1, { status: "ok" });
  immuno.event("checkout.completed", { total: 49.9 });
  await immuno.flush(); // obrigatório na Lambda
  return { ok: true };
};
```

Guarde a chave numa variável de ambiente da função (ou no Secrets Manager), nunca no código.

## Passo 08 — Ler a Visão Geral

- Grupo: Operar
- Quando: todo dia. É a primeira tela depois do login.
- Palavras-chave: visão geral, overview, dashboard inicial, status, saudável, atenção, crítico, alertas disparando, heartbeat, tela vazia, primeiro recurso, home, início

A Visão Geral responde uma pergunta: está tudo bem agora? Os números são os mesmos em todas as telas: um alerta crítico aqui é o mesmo alerta em Alertas e Incidentes.

Menu: **Operar › Visão Geral**. Há o botão **Atualizar status**.

1. **Faixa de status:** Saudável, Atenção ou Crítico, com a frase que explica o porquê.
2. **Números do topo:** alertas disparando, clusters conectados, workloads monitorados e o último sinal recebido (heartbeat).
3. **Alertas disparando:** cada linha traz o problema e leva à regra em Alertas e Incidentes.
4. **Clusters:** a conexão do agent de cada cluster.

**Tela vazia?** **Conecte seu primeiro recurso** quer dizer que nenhum cluster, VM ou serviço Serverless mandou dado ainda. Volte ao passo 05. A tela não mostra número inventado enquanto não há dado.

## Passo 09 — Investigar um incidente

- Grupo: Operar
- Quando: quando algo sai do normal
- Palavras-chave: investigar, incidente, alerta, alertas e incidentes, workloads, pods, CrashLoop, restart, observabilidade, logs, traces, métricas, OpenTelemetry, dashboards, Grafana, causa raiz, troubleshooting, latência, erro 5xx

Resposta a incidente com contexto, não alerta solto. Roteiro do alerta à causa:

1. **Alertas e Incidentes — triagem.** Veja a linha do tempo ao vivo e reconheça o alerta, para o time saber que alguém está olhando.
2. **Workloads — qual serviço.** Filtre por namespace, ambiente e severidade. Cada workload traz a evidência e a ação recomendada. A tela se chama Workloads operacionais e cobre saúde de runtime, risco de confiabilidade, capacidade e custo. Filtros citados: namespace, ambiente e severidade. A tabela mostra workload, restarts, ação recomendada e severidade. Exemplos de ação recomendada na tela ilustrativa: investigar CrashLoop e validar probes; ajustar requests de CPU e memória.
3. **Observabilidade — por quê.** Métricas, logs e traces da última hora cruzados. Sinais correlacionados junta pico de erro, logs de erro e spans lentos.
4. **Dashboards — acompanhar.** Dashboards do time com painéis reais, sem depender do editor do Grafana. Dashboards fica dentro de Observabilidade.

Menu de Investigar: Workloads, Alertas e Incidentes, Observabilidade, e Dashboards como subitem.

**Logs e traces vazios?** Eles só chegam de workloads instrumentados com OpenTelemetry (passo 07). Os demais continuam visíveis em **Entidades do Cluster**.

## Passo 10 — Custo e FinOps Score

- Grupo: Operar
- Quando: a pergunta é “onde estou gastando à toa?”
- Palavras-chave: custo, custos, FinOps, FinOps Score, nota, desperdício, economia, billing, fatura, FOCUS, FOCUS Normalizer, rightsizing, requests, CPU, memória, storage, rede, nuvem, SKU, preço

O cluster tem uma nota. O FinOps Score vai de 0 a 100 e é calculado com os dados reais do cluster, nunca com amostra.

1. Vá em **Analisar › Kubernetes FinOps** e escolha o cluster. A tela descreve capacidade, desperdício e recomendações por cluster.
2. Leia o **FinOps Score**: 100 menos o desperdício ponderado por custo. As barras (Capacidade, Requests, Autoscaling, Storage, Rede e Dono) mostram onde está o desperdício. A mais baixa é onde começar.
3. Veja o custo de CPU e memória, o desperdício e as opções de rightsizing.
4. Use **Como ler estes dados** para saber de onde vem cada número e o que ainda não é coletado.

**Mais de uma nuvem?** **Analisar › FOCUS Normalizer** mostra o billing das nuvens no mesmo formato (padrão FOCUS), para comparar SKUs lado a lado. O custo da conta GCP entra por aqui depois do export descrito em [Conta de nuvem](#conta-de-nuvem).

## Passo 11 — Convidar a equipe

- Grupo: Organizar
- Tempo: cerca de 1 minuto por pessoa
- Quem: Administrador
- Palavras-chave: equipe, time, usuários, pessoas, convidar, convite, papel, papéis, permissões, administrador, admin, operador, leitura, viewer, suspender, remover, reenviar convite, RBAC

1. Vá em **Administrar › Equipe** e clique em **Convidar funcionário**.
2. Informe o e-mail. O nome é opcional. Escolha o papel.
3. Clique em **Convidar**. A pessoa aparece como **Convite pendente** até aceitar.
4. No menu **Ações** de cada linha dá para reenviar ou revogar convites, alterar o papel, suspender ou remover alguém da organização.

**O e-mail não saiu?** A tela mostra o link do convite. Copie e envie direto para a pessoa.

| Papel | O que pode fazer | Indicado para |
| --- | --- | --- |
| Administrador | Gerencia pessoas, clusters e políticas. | Head de plataforma, SRE lead. |
| Operador | Aplica correções e cuida de alertas. Não gerencia pessoas. | SREs e quem está de plantão. |
| Leitura | Só vê painéis, alertas e custos. | Gestores, financeiro e devs consultando. |

Quem está com Leitura não vê os botões de conectar ou convidar. Peça a um administrador para alterar o papel em Equipe.

## Passo 12 — Proteger sua conta

- Grupo: Organizar
- Tempo: cerca de 2 minutos
- Quem: todo usuário
- Quando: recomendado no primeiro dia
- Palavras-chave: segurança, 2FA, MFA, TOTP, verificação em duas etapas, autenticação em dois fatores, autenticador, Google Authenticator, Authy, 1Password, QR Code, senha, trocar senha, proteger conta

A tela Segurança da conta pede uma senha forte e uma segunda etapa no login.

1. Clique no ícone de **Segurança da conta** no rodapé do menu, ao lado do nome.
2. Em **Verificação em duas etapas**, clique em **Ativar verificação em duas etapas**.
3. Leia o QR Code com o app autenticador. Sem câmera, use **Copiar chave** e cole no app.
4. Digite o código de 6 dígitos e clique em **Confirmar e ativar**.

**A partir daqui.** Todo login pede e-mail, senha e o código do app. Na mesma tela também se troca a senha.

Apps citados no manual: Google Authenticator, 1Password, Authy ou outro.

## Passo 13 — Falar com vendas

- Grupo: Organizar
- Tempo: cerca de 2 minutos
- Quem: Administrador
- Palavras-chave: vendas, comercial, contratar, plano, preço, limites, créditos, mais créditos, Enterprise, contrato, SSO, SLA, faturamento, proposta, protocolo, WhatsApp, e-mail, falar com vendas, upgrade

Para ampliar os limites ou contratar com SSO e SLA, fale com o time direto pelo console. O uso já vai junto no pedido.

1. Vá em **Administrar › Empresa** e, no painel **Plano**, clique em **Falar com vendas**.
2. Escolha o assunto:
   - **Trocar de plano** — ver outro plano que caiba melhor no uso.
   - **Aumentar volume de créditos** — mais créditos por mês para o uso da plataforma.
   - **Enterprise ou contrato personalizado** — SSO, SLA, volume alto, faturamento anual.
   - **Outra dúvida** — qualquer outra coisa sobre plano ou cobrança.
3. Se quiser, conte mais: por exemplo, quantos clusters pretende conectar.
4. Escolha como prefere ser contatado: **e-mail** (resposta em até 1 dia útil) ou **WhatsApp** (mais rápido no horário comercial).
5. Envie. A pessoa recebe um protocolo, e alguém do time técnico responde direto, sem fila de qualificação automática.

O manual não lista preços, nomes de planos nem cotas. Para isso, use este fluxo.

## Perguntas frequentes

- Grupo: Ajuda
- Quando: antes de chamar o time
- Palavras-chave: perguntas frequentes, FAQ, dúvidas, problemas, erro, ajuda, suporte, não recebi o código, pendente, workloads vazio, sem organização, não vejo botões, acesso de escrita, números reais, monitoring.viewer, logging.viewer, Cloud Run sem métricas, suspensão, VM não desliga, AWS, s3:GetObject, s3:ListBucket, immunoops-billing-reader, chave de acesso, CloudWatch, cloudwatch:GetMetricData, logs:FilterLogEvents, ec2:StartInstances, ec2:StopInstances, Lambda sem métricas

### Dados e acesso

**Preciso dar acesso de escrita na minha conta de nuvem?**
Não. A conta é somente leitura do billing. No GCP, o papel é **BigQuery Data Viewer** (`roles/bigquery.dataViewer`) no dataset do export, para `aegis-worker@immunoops-prod.iam.gserviceaccount.com`. O passo a passo está em [Conta de nuvem](#conta-de-nuvem). Na AWS, o papel `immunoops-billing-reader` só tem `s3:ListBucket` no prefixo `focus/` e `s3:GetObject` em `focus/*`, para a identidade do Google do Immuno Ops. Não há usuário IAM nem chave de acesso. O passo a passo está em [Permissões no papel de leitura](#permissoes-no-papel-de-leitura). Métricas da Lambda usam leitura do CloudWatch no papel `immunoops-billing-reader`, em [CloudWatch da Lambda](#cloudwatch-da-lambda). Métricas e logs do Cloud Run, e ligar ou desligar VM no GCP, usam a conta do worker, em [Permissões de VM e Cloud Run](#permissoes-de-vm-e-cloud-run). Na EC2, listar, iniciar e parar estão no papel `immunoops-billing-reader`, em [Ligar e desligar a VM](#ligar-e-desligar-a-vm).

**O Cloud Run não mostra métricas ou logs.**
No projeto do serviço, conceda `roles/monitoring.viewer` e `roles/logging.viewer` para `aegis-worker@immunoops-prod.iam.gserviceaccount.com`. Os comandos estão em [Permissões do Cloud Run](#permissoes-do-cloud-run). Se a tela disser que a coleta de logs ainda não está habilitada para a conta, fale com o suporte: o papel prepara o acesso, e a coleta começa depois dessa liberação.

**A política de suspensão não desliga a VM.**
No GCP, a conta do worker `aegis-worker@immunoops-prod.iam.gserviceaccount.com` precisa listar, iniciar e parar instâncias no projeto da VM. Os comandos do papel `immunoOpsVmPower` estão em [Permissões da VM](#permissoes-da-vm). A VM também precisa do agente instalado e de saída HTTPS para `app.immunoops.io` para enviar métricas e logs. Na EC2, listar fica no papel `immunoops-billing-reader` e ligar e desligar no papel `immunoops-power`, em [Ligar e desligar a VM](#ligar-e-desligar-a-vm).

**A Lambda não mostra invocações, duração ou concorrência.**
O worker lê isso no CloudWatch. No papel `immunoops-billing-reader`, conceda `cloudwatch:GetMetricData` e `cloudwatch:ListMetrics` no namespace `AWS/Lambda`. Sem `logs:FilterLogEvents` no grupo `/aws/lambda/*`, a aba de logs fica sem as linhas da função. Os comandos estão em [CloudWatch da Lambda](#cloudwatch-da-lambda).

**A organização bloqueia a conta do Immuno Ops no IAM.**
A política `iam.allowedPolicyMemberDomains` está recusando um membro de fora. Inclua o customer ID **C011tigzz** (organização immunoops.io, id 424787708642) como exceção, de preferência só no projeto do billing, e rode de novo os comandos de leitura.

**A conta GCP está “Sem tabela”.**
Abra a conta, informe a tabela no formato `projeto.dataset.tabela` e clique em **Salvar tabela**. Sem essa tabela a conta não entra no billing.

**Os números são reais?**
Sim. Todo dado vem do cluster e da conta de nuvem. Quando ainda não há dado, a tela diz isso em vez de mostrar um número inventado.

### Problemas comuns

**Meu cluster está como “Pendente” há muito tempo.**
O agent nunca mandou dado. Confira se o comando `helm` terminou sem erro, se os pods do agent estão `Running` e se o cluster tem saída de rede para o Immuno Ops.

**O cluster está conectado, mas Workloads está vazio.**
O inventário pode levar alguns minutos para chegar depois da conexão. Se a tela disser “Ainda não recebemos workloads de…”, teste a conexão e confira os logs do agent.

**Não recebi o código de verificação no cadastro.**
Olhe a caixa de spam e espere o contador para usar **Reenviar**. Confira o e-mail em **← Trocar e-mail**.

**“Não foi possível criar a conta”.**
Normalmente a empresa já tem cadastro. Peça um convite ao administrador da organização.

**“Sua conta não tem organização ativa”.**
O acesso foi suspenso ou removido em todas as organizações. Peça a um administrador para reativar ou convidar de novo.

**Não vejo os botões de conectar ou convidar.**
O papel é **Leitura**. Peça a um administrador para alterar o papel em **Equipe**.
