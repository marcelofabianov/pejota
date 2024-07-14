# Pejota

- [x] Criar estrutura inicial de configuracoes.
- [x] Criar conexao com banco de dados postgres.
- [-] Feature/Create-User
  - [x] Implementar CreateUserUseCase
  - [x] Implementar GetUserUseCase
  - [x] Implementar DeleteUserUseCase
  - [-] Implementar UserServiceApplication
    - [x] Implementar GetUser
    - [x] Implementar CreateUser
    - [x] Implementar DeleteUser
  - [-] Implementar UserServiceServer do gRPC
    - [x] Implementar GetUser sobre UserServiceServer
    - [x] Implementar CreateUser sobre UserServiceServer
    - [x] Implementar DeleteUser sobre UserServiceServer
  - [-] Implementar server gRPC, Reflection e HealthCheck
    - [-] Implementar Reflection
    - [ ] Implementar HealthCheck
    - [-] Registrar serviços no server UserServiceServer

