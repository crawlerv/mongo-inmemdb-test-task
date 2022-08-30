FROM golang:1.17 as builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o app ./*.go

FROM scratch
ARG CONTAINER_VERSION=dev
ENV ACCOUNTING_SENTRY_CONTAINER_VERSION=$CONTAINER_VERSION

COPY --from=builder /app/app .
COPY --from=builder /app/config/config.yml ./config/config.yml

COPY --from=builder /tmp /tmp

ENTRYPOINT ["/app"]
