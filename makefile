build:
	@go build -o north-star .
	
run: build
	@./north-star
