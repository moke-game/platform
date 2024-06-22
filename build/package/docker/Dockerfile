ARG APP_NAME=platform
# Step 1: Modules caching
FROM atomhub.openatom.cn/library/golang:1-alpine3.17 as modules
RUN set -eux && sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories
MAINTAINER gstones
RUN apk add --no-cache ca-certificates git
COPY go.mod go.sum /modules/
WORKDIR /modules
ENV GO111MODULE="on"
ENV GOPROXY="https://goproxy.cn,https://gocenter.io,direct"
RUN go mod download

# Step 2: Builder
FROM atomhub.openatom.cn/library/golang:1-alpine3.17 as builder
ARG APP_NAME
COPY --from=modules /go/pkg /go/pkg
COPY . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o /bin/app ./cmd/${APP_NAME}/service

# Step 3: Final
FROM alpine
RUN #sed -i -e 's/http:/https:/' /etc/apk/repositories
RUN set -eux && sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories
RUN #apk --no-cache add tzdata
#ENV TZ=Asia/Shanghai
COPY --from=builder /bin/app /app
CMD ["/app"]