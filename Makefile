build-docker:
	docker build -t go-template:0.1 .

run-docker:
	docker run --name go-template -it --rm -p 3000:8080 go-template:0.1

dev-docker:
	$(MAKE) build
	$(MAKE) run

dev:
	go build .
	go run .