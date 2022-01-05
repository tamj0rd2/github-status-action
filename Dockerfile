FROM golang:1.17.2 as build

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /repo

COPY . .
RUN go mod vendor

RUN go build -mod vendor -o githubstatus ./cmd/githubstatus/*.go

FROM scratch

WORKDIR /root/
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=build /repo/githubstatus ./githubstatus

ENTRYPOINT ["./githubstatus"]
