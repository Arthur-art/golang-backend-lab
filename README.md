# golang-backend-lab
Uma API RESTful em Golang desenvolvida como parte dos meus estudos na linguagem. Este projeto inclui:

## Principais recursos:
- ✅ CRUD completo (Create, Read, Update, Delete)
- ✅ Conexão com banco de dados (PostgreSQL/MySQL/SQLite)
- ✅ Autenticação JWT (opcional)
- ✅ Gerenciamento de rotas com Gorilla Mux ou Gin
- ✅ Documentação com Swagger (opcional)

Objetivo:
Aprender os fundamentos de Go (structs, handlers, concorrência) e boas práticas em APIs REST.

## Inicialização do projeto
1. dentro da pasta cmd, execute o comando:
```bash
go run main.go
```
2. no terminal execute para inciar a imagem postgres no docker:
```bash
docker compose up -d go_db
```