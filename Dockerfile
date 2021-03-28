FROM registry.cn-shenzhen.aliyuncs.com/pangxieke/portrays-builder:latest as builder
WORKDIR /go/src/reverse

COPY . .
ENV GOPROXY https://goproxy.io
ENV GO111MODULE on
RUN go mod download

#RUN go test ./... -coverprofile .testCoverage.txt \
#    && go tool cover -func=.testCoverage.txt
RUN CGO_ENABLED=0 go build -o app_d ./cmd/main.go
#     && CGO_ENABLED=0 go build ./cmd/migrate

FROM alpine:3.8
RUN apk --no-cache add ca-certificates
LABEL \
    SERVICE_80_NAME=re_http \
    SERVICE_NAME=reverse \
    description="reverse" \
    maintainer="pangxieke@126.com"

EXPOSE 3000
COPY --from=builder /go/src/reverse/app_d /bin/app
#COPY --from=builder /go/src/reverse/migrate /bin/migrate
ENTRYPOINT ["app"]
