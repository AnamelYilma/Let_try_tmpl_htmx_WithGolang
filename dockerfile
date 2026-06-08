FROM golang:1.25.1-alpine AS builder

WORKDIR /app

RUN go install github.com/a-h/templ/cmd/templ@v0.3.1020

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN templ generate
RUN CGO_ENABLED=0 GOOS=linux go build -o /out/app .

FROM alpine:3.22

WORKDIR /app

RUN adduser -D -g '' appuser

COPY --from=builder /out/app /app/app
COPY --from=builder /app/public /app/public

EXPOSE 4000

USER appuser

CMD ["/app/app"]