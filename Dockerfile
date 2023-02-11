FROM golang:alpine
ENV LANG en_US.UTF-8

COPY . .
RUN go build -o main .

CMD ["./main"]