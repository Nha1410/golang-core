# research/docker/Dockerfile

FROM research/devbox:latest

WORKDIR /app

# Sao chép các file module trước để cache dependencies
COPY go.mod /app/go.mod
COPY go.sum /app/go.sum

RUN go mod download

# Sao chép toàn bộ mã nguồn vào /app
COPY . /app

# Sao chép file .env để cấu hình port
COPY .env /app/.env

# Tải biến môi trường từ file .env
ENV PORT=${PORT}

# Sao chép cấu hình Air từ thư mục .air vào container
COPY .air/.air.toml /app/.air/.air.toml

# Lệnh để chạy ứng dụng với Air và cấu hình port từ .env
CMD ["air", "-c", ".air/.air.toml"]
