.PHONY: build clean

build: clean
	go mod tidy
	GOOS="linux" GOARCH=arm GOARM=7 go build -o bin/echo ./cmd/

deploy:
	scp bin/echo ${remote_user}@${remote_host}:~/
	scp deployments/http-echo.service ${remote_user}@${remote_host}:~/

clean:
	rm -f bin/*
