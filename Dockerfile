FROM golang:1.22.2

WORKDIR /app

ENTRYPOINT [ "tail", "-f", "/dev/null" ]