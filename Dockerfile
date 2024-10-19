# Usa uma imagem base do Go com Alpine (leve e eficiente)
FROM golang:1.22-alpine

# Define o diretório de trabalho dentro do contêiner
WORKDIR /app

# Copia o arquivo de dependências
COPY go.mod ./

# Baixa as dependências
RUN go mod download

# Copia o código fonte para o contêiner
COPY . .

# Compila o binário da aplicação Go
RUN go build -o /ott-metadata-api

# Expõe a porta 8080 para comunicação HTTP
EXPOSE 8080

# Comando de execução do contêiner
CMD [ "/ott-metadata-api" ]