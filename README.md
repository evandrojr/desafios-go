# Pós Go Expert 2024 - Desafio 1

Usando o tmux para dividir a tela em dois painéis para ficar visível a execução do servidor e cliente simultaneamente.

## Instalação em GNU/Linux compatíveis com pacotes Debian

```sh
sudo apt install tmux -y
go install github.com/air-verse/air@latest
```

## Execução

```sh
./start.sh
```

No painel do client para ficar pegando novas requisições:
```sh
go run client/client.go
```

## Finalizar

```sh
./stop.sh
```

## A Fazer:  

~~Compartilhar o código da struct cotação pra o cliente e servidor~~
