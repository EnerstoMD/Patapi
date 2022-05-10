FROM golang:1.16-alpine as builder

ENV GO111MODULE=on
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GOPROXY=https://goproxy.io,direct

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . ./
RUN go build -o /patapi

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /patapi ./

CMD [ "./patapi" ]