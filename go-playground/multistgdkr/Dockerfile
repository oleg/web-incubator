FROM golang:1.15-alpine AS build

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

FROM alpine

USER nobody:nobody
COPY --from=build /go/bin/app /app

CMD [ "/app" ]