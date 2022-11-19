# Build
FROM golang:1.19-alpine as builder

RUN apk --no-cache add ca-certificates git

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o ./app main.go \
    && chmod +x ./app
##

# App
FROM alpine:3.16.3 as app
RUN apk add bash
RUN mkdir -p /operator/storage
COPY --from=builder /app/app /operator/
EXPOSE 80

COPY deployment/common/boot.sh /boot/boot.sh
RUN chmod +x /boot/boot.sh
CMD ["/boot/boot.sh"]
##