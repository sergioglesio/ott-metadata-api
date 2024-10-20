FROM golang:1.22 AS builder
WORKDIR /app

COPY go.mod go.sum ./

# Baixe as dependências
RUN go mod download

# Copie o restante do código fonte para o diretório de trabalho
COPY . .

# Compile o aplicativo Go
RUN go build -o ott-metadata-api cmd/main.go

# Etapa final usando uma imagem leve
FROM alpine:latest

# Defina o diretório de trabalho
WORKDIR /root/

# Copie o binário da etapa de construção
COPY --from=builder /app/ott-metadata-api .

# Exponha a porta que a aplicação irá usar
EXPOSE 8080

# Comando para executar o aplicativo
CMD ["./ott-metadata-api"]