build:
	docker build -t go-template:0.1 .

run:
	docker run --name go-template -it --rm -p 3000:8080 go-template:0.1

dev:
	$(MAKE) build
	$(MAKE) run