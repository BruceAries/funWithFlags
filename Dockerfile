FROM golang:1.15.8
WORKDIR /app
COPY . /app/
CMD ["sleep", "infinity"]
