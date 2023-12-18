FROM golang:latest

WORKDIR /go/src/simple-go-app-linker

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN GOOS=linux CGO_ENABLED=0 \
    go build -a \
    -ldflags="-w -s" \
    -o md5me \
    .

CMD ["/bin/bash"]
