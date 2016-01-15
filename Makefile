build:
	GOOS=linux GOARCH=amd64 go build -tags release -o metro_delay_linux conf_loader.go const.go notifier.go watcher.go metro_delay.go
