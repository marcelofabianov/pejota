# Pejota

- [x] Criar estrutura inicial de configuracoes.
- [x] Criar conexao com banco de dados postgres.
- [-] Feature/Create-User
  - [x] Criar interfaces para definir portas GetUser
  - [-] Criar interfaces para definir portas Macro User
  - [-] Implementar UserRepository
  - [x] Implementar GetUserUseCase
  - [-] Implementar UserService
    - [x] Implementar GetUser sobre UserService
    - [x] Implementar CreateUser sobre UserService
  - [-] Implementar UserServiceServer do gRPC
    - [x] Implementar GetUser sobre UserServiceServer
  - [-] Implementar server gRPC, Reflection e HealthCheck
    - [-] Implementar Reflection
    - [ ] Implementar HealthCheck
    - [-] Registrar serviços no server UserServiceServer

