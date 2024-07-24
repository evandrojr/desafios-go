# Pós Go Expert 2024 - Desafio 1


Usando o tmux para dividir a tela em dois painéis para ficar visível a execução do servidor e cliente simultaneamente.


## Instalação em GNU/Linux compatíveis com pacotes Debian

```
sudo apt install tmux -y
go install github.com/air-verse/air@latest
```


## Execução

```
./start.sh
```

No painel do client para ficar pegando novas requisições:
```
go run client/client.go
```


## Finalizar

```
./stop.sh
```