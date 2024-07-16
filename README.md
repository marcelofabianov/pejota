# Pejota

Pejota é um projeto de estudo de GoLang, onde estou desenvolvendo um serviço gRPC para gerenciamento de trabalhos de uma pessoa como PJ.

## Arquitetura

A arquitetura do projeto é baseada no Clean Architecture e Ports and Adapters.

## Tecnologias

- [Golang](https://golang.org/)
- [gRPC](https://grpc.io/)
- [protocol buffers](https://developers.google.com/protocol-buffers)
- [PostgreSQL](https://www.postgresql.org/)
- [Viper](https://pkg.go.dev/github.com/spf13/viper)
- [Dig](https://pkg.go.dev/go.uber.org/dig)
- [Zap](https://pkg.go.dev/go.uber.org/zap)
- [Goose](https://pkg.go.dev/github.com/pressly/goose)
- [Argo2 com Crypto](golang.org/x/crypto)
- [RabbitMQ](https://www.rabbitmq.com/)
- [Keycloak](https://www.keycloak.org/)
- [Docker](https://www.docker.com/)
- [Prometheus](https://prometheus.io/)
- [Elasticsearch](https://www.elastic.co/)
- [Kibana](https://www.elastic.co/)
- [Jaeger](https://www.jaegertracing.io/)
- [GitHub Actions](https://docs.github.com/pt/actions)
- [Redis (?)](https://redis.io/)

## Requisitos

- [GoLand](https://golang.org/)
- [Docker](https://www.docker.com/)
- [Git](https://git-scm.com/)

## Estrutura

Até o momento, a estrutura do projeto está da seguinte forma:

```bash
.
├── LICENSE
├── README.md
├── _scripts
│   └── generate_protos.sh
├── bootstrap
│   ├── config.yaml
│   └── viper.go
├── cmd
│   └── main.go
├── docker-compose.yml
├── go.mod
├── go.sum
├── internal
│   └── user
│       ├── adapter
│       │   ├── grpc
│       │   │   ├── generated
│       │   │   │   ├── user.pb.go
│       │   │   │   └── user_grpc.pb.go
│       │   │   └── user_service_server.go
│       │   └── pgx
│       │       └── user_repository.go
│       ├── application
│       │   └── service.go
│       ├── container.go
│       ├── domain
│       │   ├── event.go
│       │   ├── usecase
│       │   │   ├── create_user.go
│       │   │   ├── delete_user.go
│       │   │   ├── get_user.go
│       │   │   ├── get_users.go
│       │   │   └── update_user.go
│       │   └── user.go
│       └── port
│           ├── create_user.go
│           ├── delete_user.go
│           ├── get_user.go
│           ├── get_users.go
│           ├── update_user.go
│           └── user.go
├── migrations
│   └── 20240714111257_create_user_table.sql
├── pkg
│   ├── hasher
│   │   └── hasher.go
│   └── postgres
│       └── db.go
└── proto
    └── user.proto
```

## Requisitos Funcionais

**Users**

- [x] Cadastrar usuário
- [x] Listar usuários com paginação, ordenação e filtro
- [x] Buscar usuário
- [x] Atualizar usuário
- [x] Deletar usuário
- [x] Desativar login de usuário
- [ ] Ativar login de usuário
- [ ] Alterar senha

**Customers**

- [ ] Cadastrar cliente
- [ ] Listar clientes com paginação, ordenação e filtro
- [ ] Buscar cliente
- [ ] Atualizar cliente
- [ ] Deletar cliente

**Projects**

- [ ] Cadastrar projeto
- [ ] Listar projetos com paginação, ordenação e filtro
- [ ] Buscar projeto
- [ ] Atualizar projeto
- [ ] Deletar projeto
- [ ] Atualizar status do projeto

**Tasks**

- [ ] Cadastrar tarefa
- [ ] Listar tarefas com paginação, ordenação e filtro
- [ ] Buscar tarefa
- [ ] Atualizar tarefa
- [ ] Deletar tarefa
- [ ] Atualizar status da tarefa

**Registro de horas**

- [ ] Cadastrar registro de horas
- [ ] Listar registros de horas com paginação, ordenação e filtro
- [ ] Buscar registro de horas
- [ ] Atualizar registro de horas
- [ ] Deletar registro de horas
- [ ] Atualizar status do registro de horas

**Contratos**

- [ ] Cadastrar contrato
- [ ] Listar contratos com paginação, ordenação e filtro
- [ ] Buscar contrato
- [ ] Atualizar contrato
- [ ] Deletar contrato
- [ ] Atualizar status do contrato

**Relatórios**

- [ ] Gerar relatório de horas por período e por projeto e/ou cliente

**Notificações**

...

**Autenticação**

...

**Autorização**

...

## Requisitos Não Funcionais

- [ ] Autenticação e autorização com Keycloak
- [ ] Log com Zap
- [ ] Monitoramento com Prometheus e Elasticstack/Kibana
- [ ] Tracing com Jaeger
- [ ] Eventos com RabbitMQ
- [ ] Envio de e-mails a definir
- [ ] CI/CD com GitHub Actions
- [ ] Cache a definir
- [x] gRPC com protocol buffers
- [x] Basear na arquitetura Clean Architecture e Ports and Adapters

## Executando o projeto

Para executar o projeto, siga os passos abaixo:

1. Clone o repositório:

```bash
git clone https://github.com/marcelofabianov/pejota.git
```

2. Suba o banco de dados e o RabbitMQ:

```bash
docker-compose up -d
```

3. Execute o comando para aplicar as migrations:

**Obs.:** Instale a ferramenta Goose para executar as migrations consultando o [Site](https://pressly.github.io/goose/installation/)

```bash
export DATABASE_URL="postgresql://username:password@localhost:5432/pejota_db"
```
Após configurar a variável de ambiente, execute o comando abaixo:

```bash
goose -dir migrations postgres "$DATABASE_URL" up
```

4. Executar o projeto:

```bash
go run cmd/main.go
```

5. Em caso de atualização do arquivo proto, execute o script para gerar os arquivos necessários:

```bash
./_scripts/generate_protos.sh
```

## Autor

- [Marcelo Fabiano](https://www.linkedin.com/in/marcelofabianov/)
