FROM golang:1.20
WORKDIR /app

COPY . .
RUN go build -o /app/main .

RUN apt-get update && apt-get install -y unoconv

EXPOSE 8000
CMD ["/app/main"]