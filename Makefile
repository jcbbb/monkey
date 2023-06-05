build:
	@go build -o bin/monkey

run:
	@./bin/monkey server -addr=:8010
