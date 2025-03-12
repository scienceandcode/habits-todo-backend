<p align="center"><img src="https://github.com/user-attachments/assets/73eb58e4-5270-4c00-a262-0dab049cfc48"></img></p>

<h1 align="center">HabitsTodo - Backend</h1>

O HabitsTodo é um aplicativo assistente de hábitos que auxilia no monitoramento e motivação de seus hábitos e tarefas.

Este repositório contém o back-end com todas funcionalidades (API first) do projeto. O repositório com o front-end (aplicativo KMM) está [neste repositório](https://github.com/scienceandcode/habits-todo-front-end).

## Funcionalidades

- [ ] Registro e Autenticação (confirmação de e-mail)
- [ ] Cadastro de hábitos (com frequência)
- [ ] Cadastro de motivadores:
  - [ ] v1 - Cadastro 100% manual;
  - [ ] v2 - IA sugere motivadores no cadastro;
  - [ ] v3 - IA gera motivadores e coletamos feedback do usuário.
- [ ] Gerar tarefas baseado no hábito
- [ ] Histórico de tarefas
- [ ] Gerar métricas de adesão nos hábitos (Percentual de adesão de hábitos cumpridos)
- [ ] Notificar o usuário das tarefas que precisam ser realizadas com base num horário marcado:
  - [ ] v1 - vincular google agenda.
- [ ] Receber sugestões de funcionalidades do usuário.

## Modelo de entidades e relacionamentos

- https://app.eraser.io/workspace/ae5Yr05DqVKMCY5fjlKV

## Tecnologias Utilizadas

- Back-end:
  - Go lang
    - GIN Framework
    - AWS lambda (infraestrutura)
  - CockroachDB (postgres)
- Front-end:
  - KMM (Kotlin Multiplatform Mobile): Android and IOS

Para maiores informações sobre a estrutura e documentações do projeto, consulte a [Wiki](https://github.com/scienceandcode/habits-todo-backend/wiki).

## Setup

### Dependências
- [Go](https://go.dev/) (v1.23.0 ou superior) - Linguagem de programação de código aberto que facilita a construção de software simples, confiável e eficiente.
- [Docker](https://www.docker.com/) - Plataforma para desenvolvimento, envio e execução de aplicativos em contêineres.
- [Docker Compose](https://docs.docker.com/compose/) - Ferramenta para definir e executar aplicativos Docker com múltiplos contêineres usando um arquivo YAML.


### Executando o projeto

1. Para criar o .env e arquivos usados nos volumes dos containers docker:
```
make setup
```

> Este comando irá configurar a variável GOPATH, criar o diretório de dados do PostgreSQL, copiar o arquivo .env.example para .env (se não existir) e instalar o Wire.

> Realize eventuais ajustes nas portas do `.env` caso alguma porta já esteja ocupada em seu ambiente local.

> O GOPATH é o diretório onde o Go instala as dependências e o diretório onde o projeto será clonado. O diretório padrão é `$HOME/go`, mas pode ser alterado conforme a necessidade.

2. Para subir o container docker:
```
docker compose up
```

3. Para resolver as dependências, na raiz do projeto, execute:

```
go mod tidy
```

4. Para iniciar o projeto, execute:

```
go run cmd/habits-todo-backend/main.go
```

A mensagem de log `[GIN-debug] Listening and serving HTTP on :<porta>` deve aparecer e será possível acessar a rota `localhost:<porta>/api/health` para o health check.
