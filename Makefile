setup-devbox:
	docker build -t research/devbox -f docker/devbox/Dockerfile .

build:
	docker build -t research/app -f docker/Dockerfile .

run:
	docker-compose up
