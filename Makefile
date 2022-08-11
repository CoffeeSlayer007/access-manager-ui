build-ui:
	npm --prefix ui install
	npm --prefix ui run build
	rm -rf ./internal/server/ui && mkdir ./internal/server/ui && mv ./ui/build/* ./internal/server/ui

build-api:
	go build ./cmd/access-manager-ui

run:
	./access-manager-ui start