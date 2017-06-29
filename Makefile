install:
	glide install

run:
	go run app.go

build: 
	go build app.go -o carbon

clean:
	rm carbon