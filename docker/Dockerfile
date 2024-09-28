FROM research/devbox:latest

WORKDIR /app

COPY go.mod /app/go.mod
COPY go.sum /app/go.sum

RUN go mod download

COPY . /app

COPY .env /app/.env

ENV PORT=${PORT}

COPY .air/.air.toml /app/.air/.air.toml

CMD ["air", "-c", ".air/.air.toml"]
