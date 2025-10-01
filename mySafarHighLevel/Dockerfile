FROM golang:1.22 AS build
WORKDIR /app
COPY . .
RUN go build -o req ./cmd/req

FROM debian:bookworm-slim
WORKDIR /app
COPY --from=build /app/req .
CMD ["./req"]
