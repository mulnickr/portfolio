FROM golang:1.21.6-bookworm

WORKDIR /backend

COPY . .

RUN go mod download

RUN go build -o /bin/backend

EXPOSE 5000

CMD [ "/bin/backend" ]