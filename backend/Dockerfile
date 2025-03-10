# Etapa de construção: usando a imagem oficial do Go 1.21.5
FROM golang:1.21.5-alpine AS build

# Definindo o diretório de trabalho dentro do contêiner
WORKDIR /app

# Copiando o código fonte para dentro do contêiner
COPY . .

# Baixando dependências e compilando o binário
RUN go mod tidy
RUN go build -o backend ./cmd

# Etapa final: imagem mais leve para rodar o binário
FROM alpine:latest

# Definindo o diretório de trabalho dentro do contêiner
WORKDIR /root/

# Copiando o binário compilado da etapa de build
COPY --from=build /app/backend .

# Expondo a porta que o backend irá rodar
EXPOSE 8080

# Comando para rodar o binário
CMD ["./backend"]
