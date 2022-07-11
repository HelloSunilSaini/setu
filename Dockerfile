FROM golang:1.16-alpine
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /splitwise
EXPOSE 7776
CMD [ "/splitwise" ]