export GOROOT=/usr/local/go
export GOPATH=/go
export PATH=$GOPATH/bin:$HOME/.cargo/bin:$GOROOT/bin:$PATH
cd /Server
git pull
git submodule update --recursive --remote
go build
 sudo /root/.composer/vendor/daux/daux.io/bin/daux generate

sudo  systemctl stop ahrensburg.service
sudo  systemctl start ahrensburg.service