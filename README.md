
## go.leptonica
go.leptonica wraps the leptonica library for "efficient image processing and image analysis operations".

### Installation

Install dependencies.
```
sudo apt-get install autoconf automake libtool
sudo apt-get install libpng12-dev
sudo apt-get install libjpeg62-dev
sudo apt-get install libtiff4-dev
sudo apt-get install zlib1g-dev
```

Download, configure, make and install
```
wget http://leptonica.googlecode.com/files/leptonica-1.69.tar.gz
tar zxvf leptonica-1.69.tar.gz
cd leptonica-1.69
./configure
make
sudo make install
```