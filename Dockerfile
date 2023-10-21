FROM docker.io/library/golang:1.21.3-alpine3.18 as builder

RUN apk add --update-cache \
    g++=12.2.1_git20220924-r10 \
    && rm -rf /var/cache/apk/*

WORKDIR /build

# Cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy code
COPY . .

RUN go build -a --ldflags '-linkmode external -extldflags "-static"' .

# ----------

FROM docker.io/library/alpine:3.18

RUN apk add --update-cache \
    ca-certificates=20230506-r0 \
    && rm -rf /var/cache/apk/* \
    && addgroup -S loginsrv \
    && adduser -S -g loginsrv loginsrv
USER loginsrv

ENV LOGINSRV_HOST=0.0.0.0 LOGINSRV_PORT=8080
ENTRYPOINT ["/loginsrv"]
EXPOSE 8080

COPY --from=builder /build/loginsrv /
