# Rey David
FROM golang:latest
WORKDIR /app
COPY . /app
RUN go mod download
RUN go build -o main .
EXPOSE 8088
CMD ["/app/main"]
