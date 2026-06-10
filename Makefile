.PHONY: dev build

dev:
	cd frontend && npm run dev:app

build:
	cd frontend && npm run build
	go build -ldflags="-s -w -H windowsgui" -o .build/tasks.exe .