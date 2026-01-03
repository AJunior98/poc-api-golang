# 📦 POC — API em Golang com Gin

Este projeto é uma **Prova de Conceito (POC)** desenvolvida em **Golang** utilizando o framework **Gin**, com o objetivo de estudar organização de camadas, conexão com banco de dados PostgreSQL e boas práticas básicas para construção de APIs REST.

A API expõe operações simples relacionadas a **Produtos**, permitindo cadastro e consulta.

---

## 🚀 Tecnologias utilizadas

- Golang
- Gin Web Framework
- PostgreSQL
- database/sql + lib/pq
- Docker & Docker Compose
- Postman (Collection anexada)

---

## 📁 Estrutura do Projeto

```bash
app/
 ├── cmd/
 │   └── main.go
 ├── config/
 │   └── conn.go
 ├── controller/
 │   └── product_controller.go
 ├── model/
 │   ├── model.go
 │   └── response.go
 ├── repository/
 │   └── product_repository.go
 ├── usecase/
 │   └── product_usecase.go
 ├── Dockerfile
 ├── docker-compose.yml
 ├── go.mod
 ├── go.sum
api_golang.postman_collection.json
```

A API foi organizada em **camadas**, separando responsabilidades:

| Camada | Descrição |
|--------|----------|
| controller | Recebe e trata as requisições HTTP |
| usecase | Contém as regras de negócio |
| repository | Comunicação com o banco de dados |
| model | Estruturas da entidade e resposta |
| config | Configurações gerais (DB etc.) |

---

## 🗄️ Banco de Dados

A API utiliza **PostgreSQL**.

Tabela esperada:

```sql
CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    product_name VARCHAR(255) NOT NULL,
    price NUMERIC(10,2) NOT NULL
);
```

---

## ▶️ Como Executar o Projeto

### 🐳 Opção 1 — Executar com Docker (recomendado)

Pré-requisitos:

- Docker
- Docker Compose

Dentro da pasta `app/` execute:

```bash
docker-compose up --build
```

Serão iniciados:

| Serviço | Porta |
|--------|------|
| API Go | 8000 |
| PostgreSQL | 5432 |

A API ficará disponível em:

```
http://localhost:8000
```

---

### 💻 Opção 2 — Executar localmente sem Docker

Pré-requisitos:

- Go 1.21+
- PostgreSQL rodando localmente

Configure a conexão em:

```
app/config/conn.go
```

Execute a aplicação:

```bash
go run cmd/main.go
```

---

## 📡 Endpoints Disponíveis

### 🔍 Listar todos os produtos
```
GET /products
```

**Resposta exemplo**
```json
[
  {
    "id_product": 1,
    "name": "Notebook",
    "price": 3500
  }
]
```

---

### 🔍 Buscar produto por ID
```
GET /product/{id}
```

**Resposta**
```json
{
  "id_product": 1,
  "name": "Notebook",
  "price": 3500
}
```

Caso o produto não exista:

```json
{
  "message": "product not found"
}
```

---

### ➕ Criar um produto
```
POST /products
```

**Body**
```json
{
  "name": "Teclado",
  "price": 199.90
}
```

**Resposta**
```json
{
  "id_product": 5,
  "name": "Teclado",
  "price": 199.9
}
```

---

## 🧪 Collection Postman

Este projeto inclui uma **collection do Postman**:

```
api_golang.postman_collection.json
```

Basta importar e executar as requisições ✨

---

## 🧱 Arquitetura — visão geral

Este projeto foi estruturado com inspiração em uma arquitetura limpa:

✔ Separação de camadas  
✔ Código organizado  
✔ Facilidade de manutenção e evolução  
✔ Reaproveitamento de regras de negócio

---

## ⚙️ Configurações do Banco no Docker

| Variável | Valor |
|---------|------|
| POSTGRES_USER | postgres |
| POSTGRES_PASSWORD | 1234 |
| POSTGRES_DB | postgres |
| HOST | go_db |

---

## 🙋‍♂️ Autor

Projeto criado para estudos por **Amauri**.

---

## 📌 Possíveis evoluções

- Validação de dados
- Testes automatizados
- Swagger/OpenAPI
- Migrations
- Melhorias de logs e tratamento de erros
- Evolução para uma Clean Architecture mais completa

---
