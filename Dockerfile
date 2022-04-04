FROM golang:1.16-alpine

ENV GO111MODULE=on
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GOPROXY=https://goproxy.io,direct
ENV PORT=4545

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . ./
RUN go build -o /patapi

EXPOSE 4545

CMD [ "/patapi" ]