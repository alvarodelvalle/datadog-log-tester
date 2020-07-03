FROM golang
WORKDIR /build
ADD . /build
RUN GOOS=linux go build

FROM golang
WORKDIR /app
COPY --from=0 /build/datadog-log-tester .
ENTRYPOINT ["/app/datadog-log-tester"]