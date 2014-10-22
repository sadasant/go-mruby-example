package mruby

// #cgo CFLAGS: -Imruby/include
// #cgo LDFLAGS: -L mruby/build/host/lib -lmruby -lm
// #include "gomruby.h"
import "C"

import (
	"fmt"
	"unsafe"
)

func Compile(code string) ([]byte, error) {
	// Setting up the mrb_state and the mrbc_context
	mrb := C.mrb_open()
	cxt := C.mrbc_context_new(mrb)

	defer C.mrb_close(mrb)

	// We don't want to run this code
	C.__cxt_no_exec(cxt)

	// This is necessary to dump bytecode correctly in MRuby 1.0.0
	c_filename := C.CString("")
	defer C.free(unsafe.Pointer(c_filename))
	C.__cxt_flename(mrb, cxt, c_filename)

	// Loading the source code
	c_code := C.CString(code)
	defer C.free(unsafe.Pointer(c_code))
	defer recover()
	ptr := C.mrb_load_string_cxt(mrb, c_code, cxt)

	// Getting the exception
	exc, err := C.__mrb_exc_cstr(mrb)
	if err == nil {
		return nil, fmt.Errorf(C.GoString(exc))
	}

	// Dumping the bytecode
	C.mrbc_context_free(mrb, cxt)

	var c_bin *C.uint8_t
	var c_bin_size C.size_t

	irep := C.__ptr_irep(ptr)
	if status := C.mrb_dump_irep(mrb, irep, 1, &c_bin, &c_bin_size); status < 0 {
		return nil, fmt.Errorf("MRuby dump_irep failed with status %d", status)
	}

	bin := C.GoBytes(unsafe.Pointer(c_bin), C.int(c_bin_size))

	return bin, nil
}

func RunSource(code string) error {
	// Setting up the mrb_state and the mrbc_context
	mrb := C.mrb_open()
	cxt := C.mrbc_context_new(mrb)

	defer C.mrb_close(mrb)

	// Loading the source code into the mrb_state
	c_code := C.CString(code)
	defer C.free(unsafe.Pointer(c_code))
	defer recover()
	C.mrb_load_string_cxt(mrb, c_code, cxt)

	// Getting the exception
	exc, err := C.__mrb_exc_cstr(mrb)
	if err == nil {
		return fmt.Errorf(C.GoString(exc))
	}

	return nil
}
