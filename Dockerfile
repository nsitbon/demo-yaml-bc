FROM golang:1.13.4-alpine3.10 as builder
WORKDIR /project
COPY . .
RUN CGO_ENABLED=0 go build -mod vendor -ldflags "-s -w" -o demo-cmd .

FROM alpine:3.10.3
COPY --from=builder /project/demo-cmd  /
