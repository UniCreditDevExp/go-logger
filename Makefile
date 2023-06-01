lint:
	golangci-lint run

install_lint_mac:
	brew tap "golangci/tap"
	brew install "golangci/tap/golangci-lint"

