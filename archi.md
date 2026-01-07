backend-design/
├─ cmd/
│  └─ api/
│     └─ main.go                         # entrypoint da API
│
├─ internal/
│  ├─ common/                            # infra & utilidades compartilhadas
│  │  ├─ config/
│  │  ├─ logger/
│  │  └─ database/                       # pool, tx helper, migrations helpers, etc.
│  │
│  ├─ domain/
│  │  └─ {domain}/                       # ex: user, billing, exampleDomain
│  │     ├─ ports.go                     # todas as ports (interfaces) do domínio
│  │     │                               # ex: UserRepository, NotificationGateway, etc.
│  │     │
│  │     ├─ service/                     # casos de uso / aplicação
│  │     │  └─ {service}/                # ex: account, invoice, example
│  │     │     └─ service.go             # orquestra regras + ports do domínio
│  │     │
│  │     ├─ dto/                         # DTOs por service
│  │     │  └─ {service}/
│  │     │     └─ dto.go                 # pode quebrar em request.go/response.go se precisar
│  │     │
│  │     ├─ handler/                     # handlers por service
│  │     │  └─ {service}/
│  │     │     └─ http.go                # ex: HTTP handlers; pode ter grpc.go etc.
│  │     │
│  │     └─ storage/                     # adapters concretos (infra p/ o domínio)
│  │        └─ postgres/                 # implementação em Postgres
│  │           └─ {service}/
│  │              └─ repository.go       # implementa as interfaces de ports.go
│  │
│  └─ http/
│     ├─ router/
│     │  └─ router.go                    # monta rotas e faz o wiring dos handlers
│     └─ middleware/
│        └─ ...
│
├─ dto/                                  # DTOs realmente cross-domain (se existirem)
│  └─ shared.go
│
├─ ARCHITECTURE.md
└─ go.mod