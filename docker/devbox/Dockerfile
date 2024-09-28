# research/devbox/Dockerfile

FROM golang:1.23

# Đặt thư mục làm việc mặc định
WORKDIR /go/src

# Cài đặt Air cho hot reload
RUN go install github.com/air-verse/air@latest

# Thiết lập môi trường phát triển (nếu cần)
ENV DEV_ENV=development

# Đặt thư mục làm việc
WORKDIR /app

# Lệnh mặc định cho container
CMD ["air"]
