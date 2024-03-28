run:
	@go run . serve

dev:
	@echo "[INFO]: Moving assets"
	@rm -rf ./res/views/*
	@rm -rf ./res/static/*
	@cp -r ./frontend/index.html ./res/views
	@cp -r ./frontend/assets/* ./res/static/
	@go run . serve

init:
	@go run . init
	@go run . migrate

clear:
	@rm -rf ./res

test:
	@go test ./...
