
npm install


# aparently we need openssl and the current image dont have it so run:

apt update && apt upgrade -y

apt install build-essential checkinstall zlib1g-dev -y

cd /usr/local/src/

//no wget either

apt install wget

wget https://www.openssl.org/source/openssl-3.0.2.tar.gz

tar -xvf openssl-3.0.2.tar.gz

./config --prefix=/usr/local/ssl --openssldir=/usr/local/ssl shared zlib

apt update && apt install make

make

make test

sleep 4

make install

clear

openssl version -a

npm run start:dev