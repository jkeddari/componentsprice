templ:
	templ generate -watch -proxy=http://localhost:3000

dev:
	make templ &
	go build -o ./tmp/$(APP_NAME) ./cmd/$(APP_NAME)/main.go && air

build:
	make templ-generate && go build -o ./bin/$(APP_NAME) ./cmd/$(APP_NAME)/main.go


docker:
	docker build -t componentsprice .
