FROM golang:1.13.8-alpine
WORKDIR /go/src/github.com/aapi-rp/geo-velocity/

COPY main.go .

COPY go.mod .
COPY go.sum .
RUN go mod download
RUN apk update  \
       && apk add build-base \
       && apk add g++

RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o main main.go


FROM alpine:latest

RUN apk update \
    && apk add sqlite \
    && apk --no-cache add ca-certificates


WORKDIR /root/
COPY --from=0 /go/src/github.com/aapi-rp/geo-velocity/main .
CMD ["./main"]



