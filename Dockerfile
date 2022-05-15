FROM golang:alpine as builder

ENV APP_HOME /go/src/id-server

WORKDIR "$APP_HOME"
COPY . .

RUN go mod download
RUN go mod verify
RUN go build -o /id-server

FROM golang:alpine

ENV APP_HOME /go/src/id-server
RUN mkdir -p "$APP_HOME"
WORKDIR "$APP_HOME"

COPY --from=builder "$APP_HOME"/id-server $APP_HOME

EXPOSE 8080
CMD ["./id-server"]
