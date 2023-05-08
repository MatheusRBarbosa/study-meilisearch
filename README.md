## 10x - Boilerplate Go

Este boilerplate contém as seguintes libs e sua respectiva utilidades:

#### Cobra (CLI)

https://github.com/spf13/cobra

Esta lib é usada para fazer gestão de comandos CLI. Por exemplo:

- `> go run main.go api` - Inicia servidor HTTP (api)
- `> go run main.go create migration [nome]` - Cria um migration
- `> go run main.go migrate [up|down]` - Aplica ou faz rollback de migration
- `> go run main.go worker` - TODO
- `> go run main.go cron` - TODO

#### Fiber (API)

https://gofiber.io/

Esta lib resolve toda parte http do projeto. Com isso algumas rotas e middlewares já estão configurados. Além disso, usa a lib [go-json](https://github.com/goccy/go-json) para fazer `JSON.Marshal` e `Json.Unmarshal`

**Rotas existes:**

- Signup `/users`
- Login `/login`
- Forgot Password `/forgot-password`
- Validate code `/validate-code`
- Change password `/change-password`

Também já existe um validator de body de request, localizado em `api/requests/validator.go`

**Middlewares disponíveis:**

- **auth_middleware**: Verifica se header de jwt está válido
- **error_middleware**: Handler global de error. Usado para traduzir erros de domínio (`domain/exceptions`)
- **is_admin_middleware**: Verifica se usuário autenticado possui role de admin
- **is_sadmin_middleware**: Verifica se usuário autenticado possui role de super admin
- **is_user_middleware**: Verifica se usuário autenticado possui role de user

#### godotenv (.env)

https://github.com/joho/godotenv

Lib para carregar as variáveis de ambiente apartir do arquivo `.env`

#### Gorm (ORM)

https://gorm.io/

Lib para fazer interface do banco (ORM). Atualmente toda conexão do banco pressupoe que seja Postgres. Padrão do projeto:

- Nome de tabela sneak case em plural. Ex: `users`, `confirmations`, `user_accounts`
- Nome de colunas em camelCase. Ex: `roleId`, `userId`.
  - As unicas exceções são para as colunas auto-gerenciadas pelo Grom que atualmente são: `created_at`, `updated_at` e `deleted_at`

#### Goose (migrations control)

https://github.com/pressly/goose

Lib para fazer controle de migrations. Interage junto ao Cobra.

Algumas migrations para criação de algumas entidades bases já foram criadas, como `users`, `roles` e `confirmations`.

#### Zap (logger)

https://github.com/uber-go/zap

Lib para melhorar o logger padrao

#### go-gomail (Email)

https://github.com/go-gomail/gomail

Lib usada para fazer envio de emails. Atualmente existe uma configuração de um serviço de notificação para usar o serviço de email, que é capaz de carregar templates e arquivos para serem enviados por email.

Preferenciamente, use o serviço de notificação ao inves de usar o de email diretamente.

#### Worker

TODO

#### Cron

TODO

#### Outras

Esse projeto usa como gerenciador de scripts e algumas lib de desenvolvimento o `npm`. Então sempre execute `npm install` depois de baixar o projeto.

Esse boilerplate também já vem configurado com um `docker-compose.yml` na raiz do p
rojeto com as dependências de infra-estrutura.
