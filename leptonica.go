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
	"sync"
	"unsafe"
)

type Pix struct {
	cPix   *C.PIX // exported C.PIX so it can be used with other cgo wrap packages
	closed bool
	lock   sync.Mutex
}

func (p *Pix) CPIX() *C.PIX {
	return p.cPix
}

func (p *Pix) GetDimensions() (int32, int32, int32) {
	p.lock.Lock()
	defer p.lock.Unlock()
	var w, h, d int32
	cW := C.l_int32(w)
	cH := C.l_int32(h)
	cD := C.l_int32(d)
	if !p.closed {
		C.pixGetDimensions(p.cPix, &cW, &cH, &cD)
	}
	return int32(cW), int32(cH), int32(cD)
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
