FROM golang:1.17-alpine3.15 as builder

COPY . /goapp
WORKDIR /goapp

RUN apk add git build-base && go mod download && go build -ldflags="-linkmode external -extldflags -static" -o main && chmod +x main

FROM scratch

COPY --from=builder /goapp/main /main
WORKDIR /workdir

ENTRYPOINT ["/main"]