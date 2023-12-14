FROM golang:latest as builder

WORKDIR /app

RUN export GO111MODULE=on

COPY go.mod ./
COPY go.sum ./
RUN go mod download 
COPY . .

RUN CGO_ENABLED=0 go build -v -o ./redditclone ./06_databases/99_hw/redditclone/cmd/redditclone

FROM alpine:latest

COPY --from=builder /app/redditclone .

EXPOSE 8080

ENTRYPOINT ["./redditclone"]
