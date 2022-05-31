// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// glib.Type values for gtkpapersize.go.
var GTypePaperSize = coreglib.Type(C.gtk_paper_size_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypePaperSize, F: marshalPaperSize},
	})
}

// PAPER_NAME_A3: name for the A3 paper size.
const PAPER_NAME_A3 = "iso_a3"

// PAPER_NAME_A4: name for the A4 paper size.
const PAPER_NAME_A4 = "iso_a4"

// PAPER_NAME_A5: name for the A5 paper size.
const PAPER_NAME_A5 = "iso_a5"

// PAPER_NAME_B5: name for the B5 paper size.
const PAPER_NAME_B5 = "iso_b5"

// PAPER_NAME_EXECUTIVE: name for the Executive paper size.
const PAPER_NAME_EXECUTIVE = "na_executive"

// PAPER_NAME_LEGAL: name for the Legal paper size.
const PAPER_NAME_LEGAL = "na_legal"

// PAPER_NAME_LETTER: name for the Letter paper size.
const PAPER_NAME_LETTER = "na_letter"

// PaperSize: GtkPaperSize handles paper sizes.
//
// It uses the standard called PWG 5101.1-2002 PWG: Standard for Media
// Standardized Names (http://www.pwg.org/standards.html) to name the paper
// sizes (and to get the data for the page sizes). In addition to standard paper
// sizes, GtkPaperSize allows to construct custom paper sizes with arbitrary
// dimensions.
//
// The GtkPaperSize object stores not only the dimensions (width and height) of
// a paper size and its name, it also provides default print margins.
//
// An instance of this type is always passed by reference.
type PaperSize struct {
	*paperSize
}

// paperSize is the struct that's finalized.
type paperSize struct {
	native *C.GtkPaperSize
}

func marshalPaperSize(p uintptr) (interface{}, error) {
	b := coreglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &PaperSize{&paperSize{(*C.GtkPaperSize)(b)}}, nil
}

// NewPaperSize constructs a struct PaperSize.
func NewPaperSize(name string) *PaperSize {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	if name != "" {
		_arg0 = (*C.void)(unsafe.Pointer(C.CString(name)))
		defer C.free(unsafe.Pointer(_arg0))
	}
	*(*string)(unsafe.Pointer(&args[0])) = _arg0

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(name)

	var _paperSize *PaperSize // out

	_paperSize = (*PaperSize)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_paperSize)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.gtk_paper_size_free((*C.GtkPaperSize)(intern.C))
		},
	)

	return _paperSize
}

// NewPaperSizeFromGVariant constructs a struct PaperSize.
func NewPaperSizeFromGVariant(variant *glib.Variant) *PaperSize {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(gextras.StructNative(unsafe.Pointer(variant)))
	*(**glib.Variant)(unsafe.Pointer(&args[0])) = _arg0

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(variant)

	var _paperSize *PaperSize // out

	_paperSize = (*PaperSize)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_paperSize)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.gtk_paper_size_free((*C.GtkPaperSize)(intern.C))
		},
	)

	return _paperSize
}

// NewPaperSizeFromIPP constructs a struct PaperSize.
func NewPaperSizeFromIPP(ippName string, width float64, height float64) *PaperSize {
	var args [3]girepository.Argument
	var _arg0 *C.void  // out
	var _arg1 C.double // out
	var _arg2 C.double // out
	var _cret *C.void  // in

	_arg0 = (*C.void)(unsafe.Pointer(C.CString(ippName)))
	defer C.free(unsafe.Pointer(_arg0))
	_arg1 = C.double(width)
	_arg2 = C.double(height)
	*(*string)(unsafe.Pointer(&args[0])) = _arg0
	*(*float64)(unsafe.Pointer(&args[1])) = _arg1
	*(*float64)(unsafe.Pointer(&args[2])) = _arg2

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(ippName)
	runtime.KeepAlive(width)
	runtime.KeepAlive(height)

	var _paperSize *PaperSize // out

	_paperSize = (*PaperSize)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_paperSize)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.gtk_paper_size_free((*C.GtkPaperSize)(intern.C))
		},
	)

	return _paperSize
}

// NewPaperSizeFromKeyFile constructs a struct PaperSize.
func NewPaperSizeFromKeyFile(keyFile *glib.KeyFile, groupName string) (*PaperSize, error) {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out
	var _cret *C.void // in
	var _cerr *C.void // in

	_arg0 = (*C.void)(gextras.StructNative(unsafe.Pointer(keyFile)))
	if groupName != "" {
		_arg1 = (*C.void)(unsafe.Pointer(C.CString(groupName)))
		defer C.free(unsafe.Pointer(_arg1))
	}
	*(**glib.KeyFile)(unsafe.Pointer(&args[0])) = _arg0
	*(*string)(unsafe.Pointer(&args[1])) = _arg1

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(keyFile)
	runtime.KeepAlive(groupName)

	var _paperSize *PaperSize // out
	var _goerr error          // out

	_paperSize = (*PaperSize)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_paperSize)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.gtk_paper_size_free((*C.GtkPaperSize)(intern.C))
		},
	)
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _paperSize, _goerr
}

// NewPaperSizeFromPPD constructs a struct PaperSize.
func NewPaperSizeFromPPD(ppdName string, ppdDisplayName string, width float64, height float64) *PaperSize {
	var args [4]girepository.Argument
	var _arg0 *C.void  // out
	var _arg1 *C.void  // out
	var _arg2 C.double // out
	var _arg3 C.double // out
	var _cret *C.void  // in

	_arg0 = (*C.void)(unsafe.Pointer(C.CString(ppdName)))
	defer C.free(unsafe.Pointer(_arg0))
	_arg1 = (*C.void)(unsafe.Pointer(C.CString(ppdDisplayName)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.double(width)
	_arg3 = C.double(height)
	*(*string)(unsafe.Pointer(&args[0])) = _arg0
	*(*string)(unsafe.Pointer(&args[1])) = _arg1
	*(*float64)(unsafe.Pointer(&args[2])) = _arg2
	*(*float64)(unsafe.Pointer(&args[3])) = _arg3

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(ppdName)
	runtime.KeepAlive(ppdDisplayName)
	runtime.KeepAlive(width)
	runtime.KeepAlive(height)

	var _paperSize *PaperSize // out

	_paperSize = (*PaperSize)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_paperSize)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.gtk_paper_size_free((*C.GtkPaperSize)(intern.C))
		},
	)

	return _paperSize
}

// Copy copies an existing GtkPaperSize.
//
// The function returns the following values:
//
//    - paperSize: copy of other.
//
func (other *PaperSize) Copy() *PaperSize {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(gextras.StructNative(unsafe.Pointer(other)))
	*(**PaperSize)(unsafe.Pointer(&args[0])) = _arg0

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(other)

	var _paperSize *PaperSize // out

	_paperSize = (*PaperSize)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_paperSize)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.gtk_paper_size_free((*C.GtkPaperSize)(intern.C))
		},
	)

	return _paperSize
}

// DisplayName gets the human-readable name of the GtkPaperSize.
//
// The function returns the following values:
//
//    - utf8: human-readable name of size.
//
func (size *PaperSize) DisplayName() string {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(gextras.StructNative(unsafe.Pointer(size)))
	*(**PaperSize)(unsafe.Pointer(&args[0])) = _arg0

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(size)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Name gets the name of the GtkPaperSize.
//
// The function returns the following values:
//
//    - utf8: name of size.
//
func (size *PaperSize) Name() string {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(gextras.StructNative(unsafe.Pointer(size)))
	*(**PaperSize)(unsafe.Pointer(&args[0])) = _arg0

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(size)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// PPDName gets the PPD name of the GtkPaperSize, which may be NULL.
//
// The function returns the following values:
//
//    - utf8: PPD name of size.
//
func (size *PaperSize) PPDName() string {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(gextras.StructNative(unsafe.Pointer(size)))
	*(**PaperSize)(unsafe.Pointer(&args[0])) = _arg0

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(size)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// IsCustom returns TRUE if size is not a standard paper size.
//
// The function returns the following values:
//
//    - ok: whether size is a custom paper size.
//
func (size *PaperSize) IsCustom() bool {
	var args [1]girepository.Argument
	var _arg0 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(gextras.StructNative(unsafe.Pointer(size)))
	*(**PaperSize)(unsafe.Pointer(&args[0])) = _arg0

	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(size)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsEqual compares two GtkPaperSize objects.
//
// The function takes the following parameters:
//
//    - size2: another GtkPaperSize object.
//
// The function returns the following values:
//
//    - ok: TRUE, if size1 and size2 represent the same paper size.
//
func (size1 *PaperSize) IsEqual(size2 *PaperSize) bool {
	var args [2]girepository.Argument
	var _arg0 *C.void    // out
	var _arg1 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(gextras.StructNative(unsafe.Pointer(size1)))
	_arg1 = (*C.void)(gextras.StructNative(unsafe.Pointer(size2)))
	*(**PaperSize)(unsafe.Pointer(&args[1])) = _arg1

	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(size1)
	runtime.KeepAlive(size2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsIPP returns TRUE if size is an IPP standard paper size.
//
// The function returns the following values:
//
//    - ok: whether size is not an IPP custom paper size.
//
func (size *PaperSize) IsIPP() bool {
	var args [1]girepository.Argument
	var _arg0 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(gextras.StructNative(unsafe.Pointer(size)))
	*(**PaperSize)(unsafe.Pointer(&args[0])) = _arg0

	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(size)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ToGVariant: serialize a paper size to an a{sv} variant.
//
// The function returns the following values:
//
//    - variant: new, floating, GVariant.
//
func (paperSize *PaperSize) ToGVariant() *glib.Variant {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(gextras.StructNative(unsafe.Pointer(paperSize)))
	*(**PaperSize)(unsafe.Pointer(&args[0])) = _arg0

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(paperSize)

	var _variant *glib.Variant // out

	_variant = (*glib.Variant)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	C.g_variant_ref(_cret)
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_variant)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.g_variant_unref((*C.GVariant)(intern.C))
		},
	)

	return _variant
}

// ToKeyFile: this function adds the paper size from size to key_file.
//
// The function takes the following parameters:
//
//    - keyFile: GKeyFile to save the paper size to.
//    - groupName: group to add the settings to in key_file.
//
func (size *PaperSize) ToKeyFile(keyFile *glib.KeyFile, groupName string) {
	var args [3]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out
	var _arg2 *C.void // out

	_arg0 = (*C.void)(gextras.StructNative(unsafe.Pointer(size)))
	_arg1 = (*C.void)(gextras.StructNative(unsafe.Pointer(keyFile)))
	_arg2 = (*C.void)(unsafe.Pointer(C.CString(groupName)))
	defer C.free(unsafe.Pointer(_arg2))
	*(**PaperSize)(unsafe.Pointer(&args[1])) = _arg1
	*(**glib.KeyFile)(unsafe.Pointer(&args[2])) = _arg2

	runtime.KeepAlive(size)
	runtime.KeepAlive(keyFile)
	runtime.KeepAlive(groupName)
}

// PaperSizeGetDefault returns the name of the default paper size, which depends
// on the current locale.
//
// The function returns the following values:
//
//    - utf8: name of the default paper size. The string is owned by GTK and
//      should not be modified.
//
func PaperSizeGetDefault() string {
	var _cret *C.void // in

	_gret := girepository.MustFind("Gtk", "get_default").Invoke(nil, nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// PaperSizeGetPaperSizes creates a list of known paper sizes.
//
// The function takes the following parameters:
//
//    - includeCustom: whether to include custom paper sizes as defined in the
//      page setup dialog.
//
// The function returns the following values:
//
//    - list: newly allocated list of newly allocated GtkPaperSize objects.
//
func PaperSizeGetPaperSizes(includeCustom bool) []*PaperSize {
	var args [1]girepository.Argument
	var _arg0 C.gboolean // out
	var _cret *C.void    // in

	if includeCustom {
		_arg0 = C.TRUE
	}
	*(*bool)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "get_paper_sizes").Invoke(args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(includeCustom)

	var _list []*PaperSize // out

	_list = make([]*PaperSize, 0, gextras.ListSize(unsafe.Pointer(_cret)))
	gextras.MoveList(unsafe.Pointer(_cret), true, func(v unsafe.Pointer) {
		src := (*C.void)(v)
		var dst *PaperSize // out
		dst = (*PaperSize)(gextras.NewStructNative(unsafe.Pointer(src)))
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(dst)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.gtk_paper_size_free((*C.GtkPaperSize)(intern.C))
			},
		)
		_list = append(_list, dst)
	})

	return _list
}
