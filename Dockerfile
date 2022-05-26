FROM golang:1.18.2-alpine

ENV PORT=8080

WORKDIR /template-go-server
COPY . /template-go-server

RUN go mod download
RUN go build

EXPOSE $PORT

CMD [ "./template-go-server" ]