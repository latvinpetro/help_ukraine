#/bin/bash

apt update -y

apt install docker.io golang-go -y

git clone https://github.com/latvinpetro/help_ukraine.git /var/help

cd /var/help
