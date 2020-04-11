FROM golang:1.13.8-alpine
WORKDIR /go/src/github.com/aapi-rp/geo-velocity


COPY config config
COPY data data
COPY db db
COPY location_api location_api
COPY logger logger
COPY messages messages
COPY model_struct model_struct
COPY models models
COPY security security
COPY test test
COPY utils utils
COPY cert.pem .
COPY go.mod .
COPY go.sum .
COPY key.pem .
COPY main.go .


RUN apk update  \
       && apk add build-base \
       && apk add g++ \
       && apk add git

RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o main main.go

RUN ls db
RUN ls .


FROM alpine:latest

RUN apk update \
    && apk add sqlite \
    && apk --no-cache add ca-certificates




WORKDIR /root/
COPY --from=0 /go/src/github.com/aapi-rp/geo-velocity/main .
COPY --from=0 /go/src/github.com/aapi-rp/geo-velocity/data data
CMD ["./main"]






