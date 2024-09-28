# research/Makefile

# Build Devbox image chỉ với Air
setup-devbox:
	docker build -t research/devbox -f docker/devbox/Dockerfile .

# Build Docker image chính của ứng dụng
build:
	docker build -t research/app -f docker/Dockerfile .

# Chạy ứng dụng với Docker Compose (nếu có Docker Compose)
run:
	docker-compose up
