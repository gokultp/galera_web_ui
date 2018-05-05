# galera_web_ui

install go 1.8 +

sudo add-apt-repository -y ppa:longsleep/golang-backports
sudo apt-get update -y
sudo apt-get install -y golang-1.8-go     <------ should be golang-go
and then, when you run the correct command later

$ vagrant ssh -c "sudo apt-get install -y golang-go"




go get 

rm -r docker/docker/vendor/github.com/docker/go-connections
go get github.com/go-errors/errors
go get github.com/pkg/errors
go get golang.org/x/net/proxy


#install docker 
https://docs.docker.com/install/linux/docker-ce/ubuntu/