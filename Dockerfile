FROM docker.io/library/golang:1.21.3-alpine3.18 as builder
RUN apk --update --no-cache add g++

WORKDIR /build

# Cache dependencies
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

# Copy code
COPY . .

RUN go build -a --ldflags '-linkmode external -extldflags "-static"' .

# ----------

FROM docker.io/library/alpine:3.18
RUN apk --update --no-cache add ca-certificates \
    && addgroup -S loginsrv && adduser -S -g loginsrv loginsrv
USER loginsrv

ENV LOGINSRV_HOST=0.0.0.0 LOGINSRV_PORT=8080
ENTRYPOINT ["/loginsrv"]
EXPOSE 8080

COPY --from=builder /build/loginsrv /