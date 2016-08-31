BINARY := kubedeploy
LDFLAGS := -ldflags="-s -w"

bin/kubenetes-slack:
	GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o bin/kubedeploy

clean:
	rm -rf bin/*
