run:
	docker-compose up

stop:
	docker-compose stop

down:
	docker-compose down

build:
	docker-compose build go

clean:
	rm carbon
	docker-compose down
