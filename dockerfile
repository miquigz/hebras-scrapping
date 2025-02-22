FROM golang:1.22 AS builder

WORKDIR /app

COPY . .

RUN go mod download

#Compilamos binarios completamente estáticos
RUN CGO_ENABLED=0 go build -o main -ldflags '-extldflags "-static"' .

# Usamos una img mas ligera para la app
FROM alpine:3.20

WORKDIR /app

# Copiamos el binario compilado
COPY --from=builder /app/main .

# Exponemos el puerto de la app
EXPOSE 8080

# ejecutamos la aplicación
CMD ["./main"]