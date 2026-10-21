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
| Administrar | Clusters, VMs, Cloud Run, Instrumentação, Equipe e Empresa. |
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
| Governança multi-cluster | Administrar. Clusters, VMs, Cloud Run e equipe num único lugar. | 05, 06 e 11 |
| Observabilidade nativa | Investigar e Administrar. Métricas, logs e traces sem depender do editor do Grafana. | 07 e 09 |

### Caminhos do console citados no manual

- Login: https://app.immunoops.io/ui/login
- Cadastro: `app.immunoops.io/ui/signup`
- Aceite de convite: `app.immunoops.io/ui/accept-invite`
- Visão Geral: `app.immunoops.io/ui/`
- Clusters: `app.immunoops.io/ui/clusters`
- Cloud Run: `app.immunoops.io/ui/cloud-run`
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
6. [Cloud Run e VMs](#passo-06-cloud-run-e-vms) — Conectar
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

Comparado com os outros tipos de recurso: no cluster se instala um agent com Helm, e ele traz workloads, métricas, logs e traces. Cloud Run e VM estão no passo 06.

## Passo 06 — Cloud Run e VMs

- Grupo: Conectar
- Tempo: cerca de 3 minutos cada
- Quem: quem usa esses serviços
- Palavras-chave: Cloud Run, VM, máquina virtual, Compute Engine, GCE, GCP, Google Cloud, projeto GCP, região, zona, serverless, docker, agente na VM, suspensão, desligar VM, horário

| Tipo | O que instalar | O que chega |
| --- | --- | --- |
| Cluster | Agent via Helm no cluster. | Workloads, métricas, logs e traces. |
| Cloud Run | Nada. | As métricas vêm do Cloud Monitoring do projeto GCP. |
| VM | Agente no host, com comando de instalação na VM do Compute Engine. | Dados do host. |

### Cloud Run

1. Vá em **Administrar › Cloud Run** e clique em **Conectar Cloud Run**.
2. Informe nome, ID do projeto GCP e região. Clique em **Registrar**.

### VMs

1. Vá em **Administrar › VMs** e clique em **Conectar VM**.
2. Informe nome, ID do projeto GCP e zona. Clique em **Registrar**.
3. Copie o comando em **Instalar o agente**, rode na VM e clique em **Concluir**.

**Economize com a aba Suspensão.** Em **VMs › Suspensão**, use **Nova política** para ligar e desligar VMs por horário, por exemplo desligar staging fora do expediente.

## Passo 07 — Instrumentar a aplicação

- Grupo: Conectar
- Tempo: cerca de 10 minutos
- Quem: time de desenvolvimento
- Opcional
- Palavras-chave: instrumentação, instrumentar, SDK, aplicação, app, código, chave de conexão, apiKey, API key, OpenTelemetry, OTel, Node.js, npm, métricas da aplicação, eventos de negócio, conectada, aguardando primeiro dado, sem dados recentes

O agent já mostra a saúde da infraestrutura. Para ver métricas de negócio, logs e eventos de dentro do código, conecte a aplicação pelo SDK.

1. Vá em **Administrar › Instrumentação** e clique em **Conectar via SDK**.
2. Dê um nome à aplicação e clique em **Gerar chave**.
3. Instale o pacote do SDK com o comando mostrado na tela.
4. Inicialize o client colando a chave de conexão em `apiKey`.
5. Envie o primeiro dado. O status muda sozinho para Conectada.

A lista de aplicações mostra, entre outros, estes estados: **Conectada** (com o tempo desde o último dado), **Aguardando primeiro dado** e **Sem dados recentes**.

**Atalho para Cloud Run e VMs.** Na linha de um serviço já registrado, **Instrumentar minha aplicação** gera token e snippet OpenTelemetry prontos para colar.

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

**Tela vazia?** **Conecte seu primeiro recurso** quer dizer que nenhum cluster, VM ou Cloud Run mandou dado ainda. Volte ao passo 05. A tela não mostra número inventado enquanto não há dado.

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

**Mais de uma nuvem?** **Analisar › FOCUS Normalizer** mostra o billing das nuvens no mesmo formato (padrão FOCUS), para comparar SKUs lado a lado.

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
- Palavras-chave: perguntas frequentes, FAQ, dúvidas, problemas, erro, ajuda, suporte, não recebi o código, pendente, workloads vazio, sem organização, não vejo botões, acesso de escrita, números reais

### Dados e acesso

**Preciso dar acesso de escrita na minha conta de nuvem?**
Não para visualização e FinOps Score: isso é leitura. O escopo nas contas de nuvem conectadas é de leitura por padrão.

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
