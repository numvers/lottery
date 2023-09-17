FROM golang:1.21-alpine3.17 as builder
RUN apk update && apk upgrade && \
    apk --update add git make bash build-base
WORKDIR /build
COPY . .
RUN make build

FROM builder as tester
RUN make test

FROM alpine:latest
RUN apk update && apk upgrade && \
    apk --update --no-cache add tzdata && \
    mkdir /app 

FROM alpine:latest
RUN apk update && apk upgrade && \
    apk --update --no-cache add tzdata
WORKDIR /
COPY --from=builder /build/cmd/http/app /app
COPY --from=builder /build/lottery.db /lottery.db
EXPOSE 8080
ENTRYPOINT ["/app"]
CMD ["--db_path=/lottery.db"]