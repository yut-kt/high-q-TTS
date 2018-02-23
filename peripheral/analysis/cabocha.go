package analysis

// #cgo LDFLAGS: -lcabocha
// #include <stdio.h>
// #include "/Users/yutakato/Dev/install/include/cabocha.h"
import "C"
import (
	"regexp"
	"strings"
)

type Cabocha struct {
	ptr *C.struct_cabocha_t
}

func NewCabocha() Cabocha {
	ptr := C.cabocha_new2(C.CString(""))
	return Cabocha{
		ptr: ptr,
	}
}

func (c *Cabocha) ParseToString(str string) string {
	return C.GoString(C.cabocha_sparse_tostr(c.ptr, C.CString(str)))
}

func (c *Cabocha) ParseToWakati(str string) string {
	tree := c.ParseToString(str)
	str = regexp.MustCompile(`[-D| ]|EOS`).ReplaceAllString(tree, "")
	str = regexp.MustCompile(`\n`).ReplaceAllString(str, " ")
	return strings.TrimSpace(str)
}
