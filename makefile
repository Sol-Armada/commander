clean-proto:
	@rm -rf ./src/proto/*

proto-gen-api: clean-proto
	@protoc -I=./server/internal/proto \
	--js_out=import_style=es6:./src/proto \
	--grpc-web_out=import_style=commonjs,mode=grpcwebtext:./src/proto \
	./server/internal/proto/commander/api/v1/*.proto

clean-build:
	@echo "Cleaning build..."
	@rm -rf ./server/dist
	@rm -rf ./server/bin

build-staging: clean-build
	@echo "Buidling web for staging..."
	@yarn build-staging
	@echo "Building service for staging..."
	@cd ./server && go build -o ./bin/commander-staging .

deploy-staging: build-staging
	@echo "Copying environment to server..."
	@scp -r ./.env.staging $(USER)@$(SERVER):/etc/commander/staging/.env
	@echo "Copying service config to server..."
	@scp -r ./server/commander-staging.service $(USER)@$(SERVER):/etc/systemd/system/commander-staging.service
	@echo "Stopping service on server..."
	@ssh $(USER)@$(SERVER) "sudo systemctl stop commander-staging"
	@echo "Copying service to server..."
	@scp ./server/bin/commander-staging $(USER)@$(SERVER):/opt
	@echo "Starting service on server..."
	@ssh $(USER)@$(SERVER) "sudo systemctl start commander-staging"
	@echo "Staging deployment complete."
