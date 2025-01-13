#!/usr/bin/env bash

echo "Aguardando banco e migrações..."
until ./migrate -path=./internal/infra/database/migrations -database "mysql://root:root@tcp(mysqldb:3306)/orders" -verbose up; do
  sleep 2
done

echo "Banco de dados pronto. Iniciando aplicação..."
cd cmd/ordersystem && go run main.go wire_gen.go
