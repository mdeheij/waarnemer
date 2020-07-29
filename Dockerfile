# Builder image
FROM golang:alpine AS build_base

RUN apk add ca-certificates git
WORKDIR /app
ENV GO111MODULE=on

COPY go.mod .
COPY go.sum .

RUN go mod download

# Compilation image
FROM build_base AS server_builder
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go install -v -a .

# Application image without Go to reduce image size
FROM alpine AS kernreactor
RUN apk add ca-certificates
WORKDIR /app

# Finally copy statically compiled Go binary.
COPY --from=server_builder /app /app
COPY --from=server_builder /go/bin/waarnemer /bin/waarnemer

ENTRYPOINT ["/bin/waarnemer"]