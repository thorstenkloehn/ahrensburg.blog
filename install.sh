sudo apt-get update
sudo apt-get upgrade
sudo apt-get install gnupg2
sudo apt-get install pandoc
sudo apt-get install cmake gcc clang gdb build-essential git-core -y
sudo apt install python3 python3-dev git curl python-is-python3 -y
sudo apt-get install php php-cli php-mbstring php-xml php-fpm -y
sudo apt-get install sqlite3 -y
sudo apt-get install composer
 sudo composer global require daux/daux.io
 if [ -f /usr/bin/node ] ; then
  echo "node ist Vorhanden"
else
 echo "node wird Installiert"
curl -fsSL https://deb.nodesource.com/setup_16.x | sudo -E bash -
sudo apt-get install -y nodejs
fi

if [ -f /usr/bin/dotnet ] ; then
  echo "Dotnet ist Vorhanden"
else
echo "Dotnet wird Installiert"
wget https://packages.microsoft.com/config/ubuntu/20.04/packages-microsoft-prod.deb -O packages-microsoft-prod.deb
sudo dpkg -i packages-microsoft-prod.deb
sudo apt-get update; \
sudo apt-get install -y apt-transport-https && \
sudo apt-get update && \
sudo apt-get install -y dotnet-sdk-5.0
  fi

 cd /root
if [ -d /usr/local/go/ ] ; then
 echo "Go ist Vorhanden"
 else
 wget https://golang.org/dl/go1.16.4.linux-amd64.tar.gz
 sudo tar -C /usr/local -xzf  go1.16.4.linux-amd64.tar.gz
 echo "Go wird Installiert"
 fi
export GOROOT=/usr/local/go
export GOPATH=/go
export PATH=$GOPATH/bin:$HOME/.cargo/bin:$GOROOT/bin:$PATH
 cd /Server
 git submodule update --init --recursive
 sudo cp -u ahrensburg.service /etc/systemd/system/ahrensburg.service
 sudo /root/.composer/vendor/daux/daux.io/bin/daux generate
go build
sudo  systemctl enable ahrensburg.service
sudo  systemctl start ahrensburg.service