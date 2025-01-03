# Desafio 3

## Rodar:

```
Renomeie:
cmd/ordersystem/.env.example para cmd/ordersystem/.env 

Execute:
docker-compose up

Já vai deve subir o banco, migrations, compilar e subir o app

Portas:
REST web server: 8000
gRPC server: 50051
GraphQL server: 8080
```


## Rest:

Instale a extensão Rest Client no VSCODE:

https://marketplace.visualstudio.com/items?itemName=humao.rest-client

Use as requisições GET e POST da pasta API

## Evans:

```
evans -r repl

127.0.0.1:50051> package pb

pb@127.0.0.1:50051> service OrderService

pb.OrderService@127.0.0.1:50051> call ListOrders
```

## GraphQL

```
mutation MyCreateCategory {
  createOrder(input: {id: "Presunto", Price: 15.0, Tax: 10.0}) {
    id
    Price
    Tax
  }
}


query OrderQuery {
  orders{
    id
    Price
    Tax
    FinalPrice
  }
}
```


## Requisitos

```
Olá devs!
Agora é a hora de botar a mão na massa. Para este desafio, você precisará criar o usecase de listagem das orders.
Esta listagem precisa ser feita com:
- Endpoint REST (GET /order)
- Service ListOrders com GRPC
- Query ListOrders GraphQL
Não esqueça de criar as migrações necessárias e o arquivo api.http com a request para criar e listar as orders.

Para a criação do banco de dados, utilize o Docker (Dockerfile / docker-compose.yaml), com isso ao rodar o comando docker compose up tudo deverá subir, preparando o banco de dados.
Inclua um README.md com os passos a serem executados no desafio e a porta em que a aplicação deverá responder em cada serviço.
```