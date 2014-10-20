## go.leptonica
go.leptonica wraps the leptonica library for "efficient image processing and image analysis operations".

### Installation
go.leptonica requires the leptonica library and development headers to compile.
Install following instructions below, then use the gopkg.in versioned release:

`go get gopkg.in/GeertJohan/go.leptonica.v1`

#### Debian Wheezy 7 (or later)
`sudo apt-get install libleptonica-dev`

#### Manual installation
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
