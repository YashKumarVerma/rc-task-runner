pack_unix:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o build/task-runner internal/main.go

pack_win:
	GOOS=windows GOARCH=386 go build -o build/task-runner.exe internal/main.go

run:
	go run internal/main.go

clean:
	rm -rf build/*
	rm -rf output/*
	rm -rf codes/*
