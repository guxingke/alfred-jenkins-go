# Alfred-jenkins-go
No dep, pure go.

## USE
1. Simple
1.1 download alfred-jenkins-go.workflow
1.2 open it
1.3 go to alfred workflow preferences , replace this workflow cfg
```
./alfred-jenkins-go https://{username}:{password}@{jenkins.server} $1
=>
e.g
./alfred-jenkins-go https://gxk:xxx@ci.duitang.net $1
```

### 2. For dev
2.1 clone or go get
2.2 ln -sn `pwd`/alfred-jenkins-go {Alfred Workflow Path}/user.workflow.alfred-jenkins-go
e.g
```
ln -sn $HOME/toy/go/src/github.com/guxingke/alfred-jenkins-go $HOME/Dropbox/Alfred/Alfred.alfredpreferences/workflows/user.workflow.alfred-jenkins-go
```
2.3 ref 1.3

## Dev
1. go
2. clone it
3. go build
4. ref USE#2

