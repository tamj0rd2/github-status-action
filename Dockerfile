FROM golang:1.17.2 as build

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /repo

COPY . .
RUN go mod vendor

ENTRYPOINT ["go", "run", "./cmd/githubstatus/"]
