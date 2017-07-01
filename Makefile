install:
	docker-compose build

run:
	docker-compose up

stop:
	docker-compose stop

down:
	docker-compose down

build: 
	go build 

clean:
	rm carbon
	docker-compose down

