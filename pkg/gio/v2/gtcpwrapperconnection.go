// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
)

// #include <stdlib.h>
// #include <gio/gio.h>
import "C"

// TCPWrapperConnectionClass: instance of this type is always passed by
// reference.
type TCPWrapperConnectionClass struct {
	*tcpWrapperConnectionClass
}

// tcpWrapperConnectionClass is the struct that's finalized.
type tcpWrapperConnectionClass struct {
	native *C.GTcpWrapperConnectionClass
}

func (t *TCPWrapperConnectionClass) ParentClass() *TCPConnectionClass {
	valptr := &t.native.parent_class
	var _v *TCPConnectionClass // out
	_v = (*TCPConnectionClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
