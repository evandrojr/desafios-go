# Use uma imagem base do Golang
FROM golang:1.23.4

# Defina o diretório de trabalho dentro do contêiner
WORKDIR /app

# Copie os arquivos para o contêiner
COPY . .

# Baixe as dependências
RUN go mod tidy

# Compile o aplicativo
RUN go build -o main .

# Especifique o comando para rodar o aplicativo
RUN chmod +x /run.sh
CMD ["./run.sh"]
