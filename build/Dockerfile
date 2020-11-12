FROM golang:1.15-alpine as builder

RUN apk add --no-cache git

ENV APP_USER gotask
ENV APP_HOME /go/src/gotask

RUN adduser -D -u 10000 $APP_USER

WORKDIR $APP_HOME

COPY . .

RUN chown -R $APP_USER $APP_HOME
USER $APP_USER

RUN go mod download 
RUN go mod verify

RUN cd /go/src/gotask/cmd/app && \
    CGO_ENABLED=0 go build -o gotask 

USER root
RUN mv /go/src/gotask/cmd/app/gotask /usr/bin/gotask

FROM alpine

ENV APP_USER gotask
RUN adduser -D -u 10000 $APP_USER

USER $APP_USER

WORKDIR /go/src/gotask

COPY --chown=0:0 --from=builder /usr/bin/gotask /usr/local/bin/gotask
COPY --chown=0:0 --from=builder /go/src/gotask/build/database.env .

CMD ["gotask"]