FRONTEND_DIR=./frontend
GO_DIR=.
SERVER_FILES=main.go files.go database.go

build-frontend:
	cd $(FRONTEND_DIR) && npm install && npm run build

run-server:
	cd $(GO_DIR) && go build $(SERVER_FILES) && ./main

all: build-frontend run-server
