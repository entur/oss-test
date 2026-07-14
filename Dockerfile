FROM golang:1.26-alpine as builder

WORKDIR /go/src/app
COPY . .

RUN go mod download && CGO_ENABLED=0 go build -o /go/bin/app

FROM gcr.io/distroless/static-debian13:nonroot
COPY --from=builder /go/bin/app /

EXPOSE 8080
USER 10001:10001
CMD ["/app"]

