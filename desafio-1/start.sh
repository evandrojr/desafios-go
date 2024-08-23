#!/usr/bin/env zsh



# Verifica se há algum processo rodando na porta 8080
PID=$(sudo lsof -t -i:8080)

if [ -n "$PID" ]; then
  echo "Processo encontrado na porta 8080 com PID: $PID. Matando o processo..."
  sudo kill -9 $PID
  echo "Processo $PID terminado."
else
  echo "Nenhum processo encontrado rodando na porta 8080."
fi

# Cria uma nova sessão tmux chamada 'sessao-desafio' e a primeira janela chamada 'Janela1'
tmux new-session -d -s sessao-desafio -n 'Janela1'

# Divide a janela 'Janela1' em dois painéis horizontalmente
tmux split-window -h -t sessao-desafio:Janela1

# Executa o primeiro script no painel 0 da 'Janela1'
tmux send-keys -t sessao-desafio:Janela1.0 'echo "Executando o primeiro script no Painel 0 da Janela 1"' C-m
tmux send-keys -t sessao-desafio:Janela1.0 'cd server && air' C-m
# tmux send-keys -t sessao-desafio:Janela1.0 'go run server/server.go' C-m

sleep 5

# Executa o segundo script no painel 1 da 'Janela1'
tmux send-keys -t sessao-desafio:Janela1.1 'echo "Executando o segundo script no Painel 1 da Janela 1"' C-m
tmux send-keys -t sessao-desafio:Janela1.1 'go run client/client.go' C-m

# Anexa à sessão
tmux attach -t sessao-desafio

