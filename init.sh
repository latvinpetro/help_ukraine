#/bin/bash

apt update -y

apt install docker.io golang-go git -y

git clone https://github.com/latvinpetro/help_ukraine.git /var/help

