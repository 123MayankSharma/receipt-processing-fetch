from golang:1.23.2-alpine

WORKDIR /challenge/

COPY go.mod go.sum ./

RUN go mod download

COPY . /challenge/


CMD ["go" ,"run","cmd/app/main.go"]

