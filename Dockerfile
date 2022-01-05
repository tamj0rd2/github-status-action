FROM golang:1.17.2

WORKDIR /repo

COPY . .
RUN go mod vendor

# Code file to execute when the docker container starts up (`entrypoint.sh`)
ENTRYPOINT ["go", "run", "./cmd/githubstatus"]
