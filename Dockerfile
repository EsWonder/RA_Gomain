# Usa una imagen base de Go
FROM golang:1.16

# Crea un directorio de trabajo
WORKDIR /app

# Copia el archivo de la aplicación
COPY main.go .

# Compila la aplicación
RUN go build -o main main.go

# Exponer el puerto en el que el servidor escucha
EXPOSE 8080

# Ejecuta el servidor
CMD ["./main"]
