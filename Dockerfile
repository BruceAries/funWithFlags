FROM golang:1.15.8
WORKDIR /app
COPY . /app/
RUN chmod 777 ./pig
CMD ["./pig"]