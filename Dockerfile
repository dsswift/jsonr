FROM golang:1.23
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO111MODULE=on GO_ENABLED=0 GOOS=linux go build -o /jsonr main.go
RUN chmod +x /jsonr
ENTRYPOINT ["/jsonr"]