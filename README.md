# kubedeploy

This tool enables to display and update Deployment's image in Kubernetes.

## Setup

### Use From Local PC

```
$ git clone https://github.com/wantedly/kubedeploy
```

Set $GOPATH this repository.

```
$ go build
$ ./kubedeploy COMMAND [OPTION]
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
