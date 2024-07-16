# Pejota

Projeto de estudo de GoLang.

Pejota é um projeto de estudo de GoLang, onde estou desenvolvendo um serviço gRPC para gerenciamento de trabalhos de uma pessoa como PJ.

## Tecnologias

- (GoLang)[https://golang.org/]
- (gRPC)[https://grpc.io/]
- (protocol buffers)[https://developers.google.com/protocol-buffers]
- (PostgreSQL)[https://www.postgresql.org/]
- (Viper)[https://pkg.go.dev/github.com/spf13/viper@v1.19.0]
- (Dig)[https://pkg.go.dev/go.uber.org/dig@v1.17.1]
- (Argo2 com Crypto)[golang.org/x/crypto]
- (RabbitMQ)[https://www.rabbitmq.com/]

## Requisitos

- (GoLang)[https://golang.org/]
- (Docker)[https://www.docker.com/]
- (Git)[https://git-scm.com/]

## Estrutura

Até o momento, a estrutura do projeto está da seguinte forma:

```bash
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

## Funcionalidades

Users

- [x] Cadastrar usuário
- [x] Listar usuários com paginação, ordenação e filtro
- [x] Buscar usuário
- [x] Atualizar usuário
- [x] Deletar usuário
- [ ] Desativar login de usuário
- [ ] Ativar login de usuário

Customers

- [ ] Cadastrar cliente
- [ ] Listar clientes com paginação, ordenação e filtro
- [ ] Buscar cliente
- [ ] Atualizar cliente
- [ ] Deletar cliente

Projects

- [ ] Cadastrar projeto
- [ ] Listar projetos com paginação, ordenação e filtro
- [ ] Buscar projeto
- [ ] Atualizar projeto
- [ ] Deletar projeto
- [ ] Atualizar status do projeto

Tasks

- [ ] Cadastrar tarefa
- [ ] Listar tarefas com paginação, ordenação e filtro
- [ ] Buscar tarefa
- [ ] Atualizar tarefa
- [ ] Deletar tarefa
- [ ] Atualizar status da tarefa

Registro de horas

- [ ] Cadastrar registro de horas
- [ ] Listar registros de horas com paginação, ordenação e filtro
- [ ] Buscar registro de horas
- [ ] Atualizar registro de horas
- [ ] Deletar registro de horas
- [ ] Atualizar status do registro de horas

Contratos

- [ ] Cadastrar contrato
- [ ] Listar contratos com paginação, ordenação e filtro
- [ ] Buscar contrato
- [ ] Atualizar contrato
- [ ] Deletar contrato
- [ ] Atualizar status do contrato

Notificações

...
