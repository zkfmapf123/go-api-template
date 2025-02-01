example-run:
	docker build -f Dockerfile.example -t ex . && docker run -p 3000:3000 ex

example-swagger-run:
	docker build -f Dockerfile.example.swagger -t ex . && docker run -p 3000:3000 ex