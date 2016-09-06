# kubedeploy

## Usage

```
$ kubedeploy get [-n namespace]
$ kubedeploy replace -p pod -i image -n namespace
$ kubedeploy deploy -s service
$ kubedeploy list -i image
```


### Example

#### deploy

```
$ kubedeploy replace -p Hello-World-xxxxx -i carumisu9/xxxyyyzzz
$ kubedeploy deploy -s rails
```
