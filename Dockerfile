FROM golang:1.25.3-alpine

WORKDIR /app

COPY . .

#RUN go mod download

RUN go build -o elevator .

CMD ["./elevator"]
