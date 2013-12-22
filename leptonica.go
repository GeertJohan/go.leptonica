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
	"unsafe"
)

type Pix struct {
	cPix *C.PIX // exported C.PIX so it can be used with other cgo wrap packages
}

func (p *Pix) CPIX() *C.PIX {
	return p.cPix
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
	return &Pix{CPIX}, nil
}

// NewPixReadMem creates a new Pix instance from a byte array
func NewPixReadMem(image *[]byte) (*Pix, error) {
	//ptr := (*C.l_uint8)(*C.uchar)(unsafe.Pointer(&(*image)[0]))
	ptr := C.uglycast(unsafe.Pointer(&(*image)[0]))
	CPIX := C.pixReadMem(ptr, C.size_t(len(*image)))
	if CPIX == nil {
		return nil, errors.New("Cannot create PIX from given image data")
	}
	return &Pix{CPIX}, nil
}
