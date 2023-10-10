# syntax = docker/dockerfile:experimental
# Build Container
FROM golang:1.19 as builder

ENV GO111MODULE on
ENV GOPRIVATE=github.com/latonaio
WORKDIR /go/src/github.com/latonaio

COPY . .
RUN go mod download
RUN go build -o convert-to-sap-product-master-from-dpfm-product-master ./

# Runtime Container
FROM alpine:3.14
RUN apk add --no-cache libc6-compat
ENV SERVICE=convert-to-sap-product-master-from-dpfm-product-master \
    APP_DIR="${AION_HOME}/${POSITION}/${SERVICE}"

WORKDIR ${AION_HOME}

COPY --from=builder /go/src/github.com/latonaio/convert-to-sap-product-master-from-dpfm-product-master .

CMD ["./convert-to-sap-product-master-from-dpfm-product-master"]
