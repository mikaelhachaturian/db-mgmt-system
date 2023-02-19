FROM golang:1.19-alpine

RUN apk add build-base

WORKDIR /app

COPY . ./

RUN go mod download
RUN go build -o /backend
EXPOSE 8081
CMD [ "/backend" ]