FROM golang:1.23-alpine as builder
RUN apk --update add build-base

WORKDIR /src/app
ADD go.mod .
RUN go mod download

ADD . .
RUN go build -o bin/ittyBitty ./cmd/main.go

FROM alpine
RUN apk add --no-cache tzdata ca-certificates
WORKDIR /

# Copying binaries
COPY --from=builder /src/app/bin/ittyBitty /bin/

# Copying views
COPY --from=builder /src/app/internal/views app/internal/

EXPOSE 8080

CMD /bin/ittyBitty