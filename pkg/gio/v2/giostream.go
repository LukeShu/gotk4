// Code generated by girgen. DO NOT EDIT.

package gio

// #include <stdlib.h>
// #include <gio/gio.h>
import "C"

// IOStreamClass: instance of this type is always passed by reference.
type IOStreamClass struct {
	*ioStreamClass
}

// ioStreamClass is the struct that's finalized.
type ioStreamClass struct {
	native *C.GIOStreamClass
}
