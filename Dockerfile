FROM golang:1.20 AS builder

MAINTAINER shouxin

WORKDIR /build

ENV GOPROXY https://goproxy.cn,direct

COPY . .
RUN CGO_ENABLED=0 go build -o main cmd/main.go


FROM alpine AS runner

WORKDIR /build
COPY  --from=builder /build/main /build/main
COPY  --from=builder /build/etc/* /build/etc/


CMD ["./main"]