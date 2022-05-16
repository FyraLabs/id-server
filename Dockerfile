FROM golang:alpine as builder

ENV APP_HOME /go/src/id-server

WORKDIR "$APP_HOME"
COPY . .

RUN apk add gcc musl-dev
RUN wget https://github.com/FyraLabs/geolite2/releases/latest/download/GeoLite2-City.mmdb -O GeoLite2-City.mmdb

RUN go mod download
RUN go mod verify
RUN go build -o id-server

FROM golang:alpine

ENV APP_HOME /go/src/id-server
WORKDIR "$APP_HOME"

COPY --from=builder "$APP_HOME"/id-server $APP_HOME
COPY --from=builder "$APP_HOME"/GeoLite2-City.mmdb $APP_HOME

ENV GEO_LITE_2_CITY_PATH=$APP_HOME/GeoLite2-City.mmdb
EXPOSE 8080
CMD ["./id-server"]
