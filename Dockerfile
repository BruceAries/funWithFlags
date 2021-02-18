FROM golang:1.15.8
WORKDIR /app
COPY . /app/
COPY ~/pig /app/
CMD ["sleep", "infinity"]
