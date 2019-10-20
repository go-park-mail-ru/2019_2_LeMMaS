FROM golang:1.13

COPY . /home/app
WORKDIR /home/app

CMD ["go", "run", "."]