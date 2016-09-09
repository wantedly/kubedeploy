# kubedeploy

This tool enables to display and update Deployment's image in Kubernetes.

## How to build and run

### Run as standalone binary

Go 1.6 or above is required.

```
$ go get -d https://github.com/wantedly/kubedeploy
$ cd $GOPATH/src/github.com/wantedly/kubedeploy
$ make
$ bin/kubedeploy COMMAND [OPTION]
```

### Use From Docker

```
$ git clone https://github.com/wantedly/kubedeploy
$ docker build .
$ docker run kubedeploy COMMAND [OPTION]
```

## Usage

```
$ kubedeploy get [-n namespace]
$ kubedeploy replace -p pod -i image -n namespace
$ kubedeploy deploy -s service -n namespace
$ kubedeploy list -i image
```
