FROM golang:1.13.8
WORKDIR /go/src/github.com/aapi-rp/geo-velocity/

COPY main.go .

COPY go.mod .
COPY go.sum .
RUN go mod download


RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main main.go


FROM alpine:latest

RUN apk update \
    && apk --no-cache add ca-certificates


WORKDIR /root/
COPY --from=0 /go/src/github.com/aapi-rp/geo-velocity/main .
CMD ["./main"]



