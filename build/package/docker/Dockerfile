ARG APP_NAME=platform
# Step 1: Modules caching
FROM golang:1.22.5-alpine as modules
RUN apk add --no-cache ca-certificates git
COPY go.mod go.sum /modules/
WORKDIR /modules
ENV GO111MODULE="on"
ENV GOPROXY="https://goproxy.cn,https://gocenter.io,direct"
RUN go mod download

# Step 2: Builder
FROM golang:1.22.5-alpine as builder
ARG APP_NAME
COPY --from=modules /go/pkg /go/pkg
COPY . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o /bin/app ./cmd/${APP_NAME}/service

# Step 3: Final
FROM alpine
RUN #sed -i -e 's/http:/https:/' /etc/apk/repositories
RUN apk --no-cache add tzdata
# you can fix it to your TZ
ENV TZ=Asia/Shanghai
COPY --from=builder /bin/app /app
CMD ["/app"]