build:
	docker build -t proxy-server .

up:
	docker-compose up -d

down:
	docker-compose down

restart: down up

.PHONY: build up down restart
