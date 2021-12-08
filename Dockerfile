FROM hub.deepin.com/library/golang:latest AS builder
WORKDIR /src
COPY . .
RUN CGO_ENABLED=0 go build -mod vendor -o link main.go

FROM hub.deepin.com/library/alpine:latest
WORKDIR /service
COPY --from=builder /src/link /link-api
CMD ["/link-api"]