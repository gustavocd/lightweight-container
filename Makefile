build:
	docker build -t goapp . --no-cache

up:
	docker-compose up --remove-orphans

down:
	docker-compose down --remove-orphans
