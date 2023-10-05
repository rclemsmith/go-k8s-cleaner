FROM golang:1.20-alpine AS BuildStage
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN go build -o /go-demo
CMD [ "/go-demo" ]

FROM alpine:latest
WORKDIR /
COPY --from=BuildStage /go-demo /go-demo
RUN chmod +x /go-demo
ENTRYPOINT ["/go-demo"]