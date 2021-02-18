FROM golang:1.15.8
WORKDIR /app
COPY . /app/
COPY $HOME/pig /app/
CMD ["sleep", "infinity"]
