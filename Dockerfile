FROM pangpanglabs/golang:builder AS builder

RUN go get github.com/fatih/structs \
    && go get github.com/relax-space/go-kit/...

ADD . /go/src/sample-go-api
WORKDIR /go/src/sample-go-api
ENV CGO_ENABLED=0
RUN go build -o sample-go-api

FROM alpine
RUN apk --no-cache add ca-certificates
# FROM scratch
WORKDIR /go/src/sample-go-api
COPY --from=builder /go/src/sample-go-api/*.yml /go/src/sample-go-api/
COPY --from=builder /go/src/sample-go-api/sample-go-api /go/src/sample-go-api/
COPY --from=builder /swagger-ui/ /go/src/sample-go-api/swagger-ui/
COPY --from=builder /go/src/sample-go-api/index.html /go/src/sample-go-api/swagger-ui/


EXPOSE 8080

CMD ["./sample-go-api"]
