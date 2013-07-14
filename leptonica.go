package leptonica

// #cgo LDFLAGS: -llept
// #include "leptonica/allheaders.h"
// #include <stdlib.h>
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
