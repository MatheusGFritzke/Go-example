# Go To-Do List API

Este projeto é uma simples API RESTful escrita em Go usando o framework Gin para gerenciar uma lista de tarefas (to-do list). A API suporta operações CRUD (Create, Read, Update, Delete).

## Pré-requisitos

- Go 1.16 ou superior

## Instalação

1. Clone o repositório:

    ```sh
    git clone https://github.com/MatheusGFritzke/Go-example.git
    cd Go-example
    ```

2. Inicialize o módulo Go:

    ```sh
    go mod init Go-example
    ```

3. Instale as dependências:

    ```sh
    go get -u github.com/gin-gonic/gin
    ```

## Uso

1. Execute a aplicação:

    ```sh
    go run main.go
    ```

2. Utilize os seguintes endpoints para interagir com a API:

### Endpoints e Exemplos de Uso no Terminal

#### Listar todas as tarefas

```http
GET /todos
#### Criar uma nova tarefa

```http
POST /todos
```

#### Atualizar uma tarefa existente

```http
PUT /todos/:id
```

#### Excluir uma tarefa

```http
DELETE /todos/:id
