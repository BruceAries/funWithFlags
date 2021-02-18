FROM golang:1.15.8
WORKDIR /app
COPY . /app/
COPY $GOPATH/pig /app/
CMD ["sleep", "infinity"]
