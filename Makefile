BINARY := kubedeploy
LDFLAGS := -ldflags="-s -w"

bin/$(BINARY):
	GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o bin/$(BINARY)

clean:
	rm -rf bin/*
