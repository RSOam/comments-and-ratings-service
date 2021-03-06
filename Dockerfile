FROM golang:1.17.3-alpine3.14 as builder
RUN apk add git
RUN mkdir /comments-and-ratings-service
ADD . /comments-and-ratings-service
WORKDIR /comments-and-ratings-service

COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest

RUN apk --no-cache add ca-certificates curl

RUN mkdir /comments-and-ratings-service

WORKDIR /comments-and-ratings-service/

COPY --from=builder /comments-and-ratings-service/main .

ARG DBpw_arg=default_value 
ENV DBpw=$DBpw_arg

EXPOSE 8080

CMD ["./main"]