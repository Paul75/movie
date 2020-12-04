FROM golang:1.15 AS builder
WORKDIR ./src
RUN pwd
COPY . .
RUN go mod download -x
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o api_movie .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder ./go/src/api_movie .
COPY --from=builder ./go/src/config.yaml .
CMD ["./api_movie"]  