
## go.leptonica
go.leptonica wraps the leptonica library for "efficient image processing and image analysis operations".

### Installation

#### On Debian Wheezy 7 (or later)

apt-get install libleptonica-dev

It's already at 1.69, so forget about the rest of this document.

#### Elsewhere

Install dependencies.
```
sudo apt-get install autoconf automake libtool libpng12-dev libjpeg62-dev libtiff4-dev zlib1g-dev
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