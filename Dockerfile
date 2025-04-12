# Go бейнесін пайдалану
FROM golang:1.23-alpine

# Жұмыс директориясын құру
WORKDIR /app

# Go модульдерін жүктеу үшін go.mod файлын көшіру
COPY go.mod go.sum ./

# Тәуелділіктерді орнату
RUN go mod download

# Жоба файлының барлық файлдарын көшіру
COPY . .

# Жобаны құру
RUN go build -o main cmd/main.go

# Контейнер жұмыс істейтін портты ашу
EXPOSE 8080

# Қолданбаны іске қосу
CMD ["./main"]
