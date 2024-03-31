# Builder stage
FROM golang:latest AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . /app
RUN CGO_ENABLED=0 go build -o bin/ ./...

# Release stage
FROM gcr.io/distroless/static-debian11
COPY --from=build /app/bin/server /server
COPY --from=build /app/cert /certs
CMD ["/server"]
