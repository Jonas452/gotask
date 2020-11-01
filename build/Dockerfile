FROM golang:1.15-alpine as builder

WORKDIR /go/src/gotask

COPY go.mod go.sum ./

RUN go mod download 

COPY . .

RUN cd /go/src/gotask/cmd/app && \
    go build -o gotask && \
    mv ./gotask /usr/bin/gotask

FROM alpine

WORKDIR /go/src/gotask

RUN apk add --no-cache ca-certificates && update-ca-certificates

COPY --from=builder /usr/bin/gotask /usr/local/bin/gotask
COPY ./build/database.env /go/src/gotask

CMD ["gotask"]