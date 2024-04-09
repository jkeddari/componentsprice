templ:
	templ generate -watch -proxy=http://localhost:3000

dev:
	templ &
	air &

build:
	make templ-generate && go build -o ./build/$(APP_NAME) ./cmd/$(APP_NAME)/main.go


docker:
	docker build -t componentsprice .
