.PHONY: init
init:
	go install github.com/cloudwego/kitex/tool/cmd/kitex@latest
	go install github.com/cloudwego/thriftgo@latest
	go install github.com/google/wire/cmd/wire@latest
	go install github.com/swaggo/swag/cmd/swag@latest

.PHONY: wire
wire:
	wire
.PHONY: swag
swag:
	swag init  --parseDependency