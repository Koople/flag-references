FROM golang:alpine as builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOTRACEBACK=all \
    GOARCH=amd64

WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go test ./...
RUN go build -o bin/kpl ./cmd

WORKDIR /dist
RUN cp /build/bin/kpl .


FROM scratch
COPY --from=builder /dist/kpl /
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

WORKDIR /application

ENTRYPOINT ["/kpl"]