build:
	docker build . -t gdg-kutaisi-123

run:
	docker run --rm -p 8888:8080 gdg-kutaisi-123