package leptonica

/*
#cgo LDFLAGS: -llept
#include "leptonica/allheaders.h"
#include <stdlib.h>

l_uint8* uglycast(void* value) { return (l_uint8*)value; }

*/
import "C"
import (
	"errors"
	"strconv"
	"sync"
	"unsafe"
)

type ImageType int32

const (
	UNKNOWN ImageType = iota
	BMP
	JFIF_JPEG
	PNG
	TIFF
	TIFF_PACKBITS
	TIFF_RLE
	TIFF_G3
	TIFF_G4
	TIFF_LZW
	TIFF_ZIP
	PNM
	PS
	GIF
	JP2
	WEBP
	LPDF
	DEFAULT
	SPIX
)

type Pix struct {
	cPix   *C.PIX // exported C.PIX so it can be used with other cgo wrap packages
	closed bool
	lock   sync.Mutex
}

func (p *Pix) CPIX() *C.PIX {
	return p.cPix
}

// GetDimensions returns the dimensions in Width, Height, Depth, Error format
func (p *Pix) GetDimensions() (int32, int32, int32, error) {
	p.lock.Lock()
	defer p.lock.Unlock()
	var w, h, d int32
	cW := C.l_int32(w)
	cH := C.l_int32(h)
	cD := C.l_int32(d)
	if !p.closed {
		code := C.pixGetDimensions(p.cPix, &cW, &cH, &cD)
		if code != 0 {
			return 0, 0, 0, errors.New("could not get dimensions")
		}

	}
	return int32(cW), int32(cH), int32(cD), nil
}

func (p *Pix) Close() {
	p.lock.Lock()
	defer p.lock.Unlock()
	if !p.closed {
		// LEPT_DLL extern void pixDestroy ( PIX **ppix );
		C.pixDestroy(&p.cPix)
		C.free(unsafe.Pointer(p.cPix))
		p.closed = true
	}
}

// WriteFile saves to disk the current pix in the given format
func (p *Pix) WriteFile(filename string, format ImageType) error {
	cFilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cFilename))
	cFormat := C.l_int32(format)

	// create new PIX
	code := C.pixWrite(cFilename, p.cPix, cFormat)
	if code != 0 {
		return errors.New("could not write PIX to given filename: " + filename + " (format: " + strconv.Itoa(int(format)) + ")")
	}

	return nil
}

// EncodedBytes will return a byte array holding the data from PIX in the given format
func (p *Pix) EncodedBytes(format ImageType) ([]byte, error) {
	var memory []byte
	memPtr := C.uglycast(unsafe.Pointer(&(memory)))
	var i int64
	sizePtr := C.size_t(i)
	cFormat := C.l_int32(format)
	code := C.pixWriteMem(&memPtr, &sizePtr, p.cPix, cFormat)
	if code != 0 {
		return nil, errors.New("Cannot write type to given memory.  WriteMem returned: " + strconv.Itoa(int(code)))
	}
	data := C.GoBytes(unsafe.Pointer(memPtr), C.int(sizePtr))
	C.free(unsafe.Pointer(memPtr))
	return data, nil
}

// LEPT_DLL extern PIX * pixRead ( const char *filename );
// NewPixFromFile creates a new Pix from given filename
func NewPixFromFile(filename string) (*Pix, error) {
	cFilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cFilename))

	// create new PIX
	CPIX := C.pixRead(cFilename)
	if CPIX == nil {
		return nil, errors.New("could not create PIX from given filename")
	}

	// all done
	pix := &Pix{
		cPix: CPIX,
	}
	return pix, nil
}

// NewPixReadMem creates a new Pix instance from a byte array
func NewPixReadMem(image *[]byte) (*Pix, error) {
	//ptr := (*C.l_uint8)(*C.uchar)(unsafe.Pointer(&(*image)[0]))
	ptr := C.uglycast(unsafe.Pointer(&(*image)[0]))
	CPIX := C.pixReadMem(ptr, C.size_t(len(*image)))
	if CPIX == nil {
		return nil, errors.New("Cannot create PIX from given image data")
	}
	pix := &Pix{
		cPix: CPIX,
	}
	return pix, nil
}
