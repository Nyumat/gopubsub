test:
	@cd internal/domain && \
	go test -v . | sed '/PASS/s//$(shell printf "\033[32mPASS\033[0m")/' | sed '/FAIL/s//$(shell printf "\033[31mFAIL\033[0m")/' && \
	cd ..
build:
	@go build ./... \
	&& echo "\n\nBuild success\n\n"
run:
	@go run ./... \
	&& echo "\n\nRun success\n\n"
install:
	@go install ./... \
	&& npm install \
	&& go mod tidy \
	&& echo "\n\nInstall success\n\n"
help:
	@echo "make test - Run tests"
	@echo "make build - Build the project"
	@echo "make run - Run the project"
	@echo "make install - Install the project"
	@echo "make help - Show this help message"
	@echo "make all - Run all commands"