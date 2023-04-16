FROM golang:latest
workdir /app
COPY go.mod ./
COPY *.go ./
RUN go build -o main .
CMD ["./main"]
EXPOSE 80