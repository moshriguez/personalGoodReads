FROM golang:1.21

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

ENV GO111MODULE="on"
ENV GOOS="linux"
ENV CGO_ENABLED=0

COPY . .
RUN go build -v -o /usr/local/bin/app ./...

EXPOSE 3000

CMD ["app"]