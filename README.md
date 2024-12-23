<h3 align="center">Task Manager</h3>

<div align="center">

[![Status](https://img.shields.io/badge/status-active-success.svg)]()
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](/LICENSE)

</div>

---

<p align="center">
    Task manager is written using 🐹 Golang Fiber and 🚀 React.js
    <br>
    ! Currently in development..
</p>

## 📝 Table of Contents

- [Getting Started](#getting_started)
- [Usage](#usage)
- [Built Using](#built_using)
- [Authors](#authors)

## 🏁 Getting Started <a name = "getting_started"></a>

### Prerequisites

```
- Golang 1.23+
- Node.js 20+
- Postman/any other HTTP client (optional)
```

### Installing

1. Clone the repository:

```bash
1 - $ git clone https://github.com/st0pcha/task-manager.git
2 - $ cd task-manager
```

2. Go to backend folder:

```bash
1 - $ cd backend
```

3. Install backend dependencies:

```bash
1 - $ go run tidy | with clear go
  - $ godo tidy   | with st0pcha/godo
  - $ make tidy   | with make
```

4. Configure .env file: copy/rename `.env.example` to `.env`

```
MODE="PROD" # "DEV" | "PROD"

SERVER_HOST="0.0.0.0"
SERVER_PORT="8080"
SERVER_ALLOWED_ORIGINS="*"

JWT_SECRET_KEY="changemepls"

POSTGRES_HOST="127.0.0.1"
POSTGRES_PORT="5432"
POSTGRES_USER="user"
POSTGRES_PASS="changemepls"
POSTGRES_NAME="taskmanager"
POSTGRES_SCHEMA="public"
POSTGRES_DSN="host=${POSTGRES_HOST} user=${POSTGRES_USER} password=${POSTGRES_PASS} dbname=${POSTGRES_NAME} port=${POSTGRES_PORT} sslmode=disable search_path=${POSTGRES_SCHEMA}"
```

5. Build backend:

```bash
1 - $ go build -o api ./cmd                      | with clear go
  - $ godo build (or godo build-win for Windows) | with st0pcha/godo
  - $ make build (or make build-win for Windows) | with make
```

6. Go to frontend folder:

```bash
1 - $ cd frontend
```

7. Install frontend dependencies:

```
1 - $ npm install  | with npm
  - $ godo install | with st0pcha/godo
  - $ make install | with make
```

8. Build frontend:

```bash
1 - $ npm build  | with npm
  - $ godo build | with st0pcha/godo
  - $ make build | with make
```

## 🎈 Usage <a name="usage"></a>

Run backend and frontend server.
<br>
Site will available at https://localhost:3000/

## ⛏️ Built Using <a name = "built_using"></a>

#### **Backend**

- [Go](https://go.dev/) — Programming Language
- [JWT](https://jwt.io) - Authentication and Authorization
- [Go Validator](https://github.com/go-playground/validator) — Data Validation

#### **Database**

- [PostgreSQL](https://postgresql.org/) - Database

#### **Frontend**

- ...
- ...
- ...

#### **Testing and API Debugging**

- [Postman](https://www.postman.com/) — API Testing and Debugging Tool

## ✍️ Authors <a name = "authors"></a>

- [@st0pcha](https://github.com/st0pcha) - Idea and all work
