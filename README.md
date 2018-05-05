# Galera Web UI

##  install go 1.8 +

for ubuntu 16.04 LTS
```
sudo add-apt-repository -y ppa:longsleep/golang-backports
sudo apt-get update -y
sudo apt-get install -y golang-1.8-go     <------ should be golang-go
```
and then, when you run the correct command later
```
vagrant ssh -c "sudo apt-get install -y golang-go"
```



## Download package
for getting the package
```
go get github.com/gokultp/galera_web_ui
```

there are a few dependancy related bugs with docker package  
do the following to resolve those  

```
rm -r docker/docker/vendor/github.com/docker/go-connections
go get github.com/go-errors/errors
go get github.com/pkg/errors
go get golang.org/x/net/proxy
```


## install docker 

Refer this documentation to install docker  
https://docs.docker.com/install/linux/docker-ce/ubuntu/