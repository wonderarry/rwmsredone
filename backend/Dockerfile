FROM golang:1.23 as builder

WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download

RUN go install github.com/swaggo/swag/cmd/swag@latest

COPY . .
RUN swag init --generalInfo main.go --output docs --parseDependency

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /out/rwms main.go

FROM gcr.io/distroless/base-debian12

WORKDIR /app
COPY --from=builder /out/rwms /app/rwms
COPY infra/db/templates /app/templates
COPY --from=builder /src/docs /app/docs

ENV RWMS_HTTP__ADDR=":8080"
ENV RWMS_TEMPLATES_DIR="/app/templates"

EXPOSE 8080
ENTRYPOINT ["/app/rwms"]
