build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o feedback .
run:build
	docker-compose up --build
