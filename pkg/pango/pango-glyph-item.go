// Code generated by girgen. DO NOT EDIT.

package pango

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: pango
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <pango/pango.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.pango_glyph_item_get_type()), F: marshalGlyphItem},
		{T: externglib.Type(C.pango_glyph_item_iter_get_type()), F: marshalGlyphItemIter},
	})
}

// GlyphItem: PangoGlyphItem is a pair of a PangoItem and the glyphs resulting
// from shaping the items text.
//
// As an example of the usage of PangoGlyphItem, the results of shaping text
// with PangoLayout is a list of PangoLayoutLine, each of which contains a list
// of PangoGlyphItem.
type GlyphItem struct {
	nocopy gextras.NoCopy
	native *C.PangoGlyphItem
}

func marshalGlyphItem(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return &GlyphItem{native: (*C.PangoGlyphItem)(unsafe.Pointer(b))}, nil
}

// Item: corresponding PangoItem
func (g *GlyphItem) Item() *Item {
	var v *Item // out
	v = (*Item)(gextras.NewStructNative(unsafe.Pointer(g.native.item)))
	return v
}

// Glyphs: corresponding PangoGlyphString
func (g *GlyphItem) Glyphs() *GlyphString {
	var v *GlyphString // out
	v = (*GlyphString)(gextras.NewStructNative(unsafe.Pointer(g.native.glyphs)))
	return v
}

// Copy: make a deep copy of an existing PangoGlyphItem structure.
func (orig *GlyphItem) Copy() *GlyphItem {
	var _arg0 *C.PangoGlyphItem // out
	var _cret *C.PangoGlyphItem // in

	if orig != nil {
		_arg0 = (*C.PangoGlyphItem)(gextras.StructNative(unsafe.Pointer(orig)))
	}

	_cret = C.pango_glyph_item_copy(_arg0)

	runtime.KeepAlive(orig)

	var _glyphItem *GlyphItem // out

	if _cret != nil {
		_glyphItem = (*GlyphItem)(gextras.NewStructNative(unsafe.Pointer(_cret)))
		runtime.SetFinalizer(_glyphItem, func(v *GlyphItem) {
			C.pango_glyph_item_free((*C.PangoGlyphItem)(gextras.StructNative(unsafe.Pointer(v))))
		})
	}

	return _glyphItem
}

// Split modifies orig to cover only the text after split_index, and returns a
// new item that covers the text before split_index that used to be in orig.
//
// You can think of split_index as the length of the returned item. split_index
// may not be 0, and it may not be greater than or equal to the length of orig
// (that is, there must be at least one byte assigned to each item, you can't
// create a zero-length item).
//
// This function is similar in function to pango_item_split() (and uses it
// internally.)
func (orig *GlyphItem) Split(text string, splitIndex int) *GlyphItem {
	var _arg0 *C.PangoGlyphItem // out
	var _arg1 *C.char           // out
	var _arg2 C.int             // out
	var _cret *C.PangoGlyphItem // in

	_arg0 = (*C.PangoGlyphItem)(gextras.StructNative(unsafe.Pointer(orig)))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(text)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.int(splitIndex)

	_cret = C.pango_glyph_item_split(_arg0, _arg1, _arg2)

	runtime.KeepAlive(orig)
	runtime.KeepAlive(text)
	runtime.KeepAlive(splitIndex)

	var _glyphItem *GlyphItem // out

	_glyphItem = (*GlyphItem)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(_glyphItem, func(v *GlyphItem) {
		C.pango_glyph_item_free((*C.PangoGlyphItem)(gextras.StructNative(unsafe.Pointer(v))))
	})

	return _glyphItem
}

// GlyphItemIter: PangoGlyphItemIter is an iterator over the clusters in a
// PangoGlyphItem.
//
// The *forward direction* of the iterator is the logical direction of text.
// That is, with increasing start_index and start_char values. If glyph_item is
// right-to-left (that is, if glyph_item->item->analysis.level is odd), then
// start_glyph decreases as the iterator moves forward. Moreover, in
// right-to-left cases, start_glyph is greater than end_glyph.
//
// An iterator should be initialized using either
// pango_glyph_item_iter_init_start() or pango_glyph_item_iter_init_end(), for
// forward and backward iteration respectively, and walked over using any
// desired mixture of pango_glyph_item_iter_next_cluster() and
// pango_glyph_item_iter_prev_cluster().
//
// A common idiom for doing a forward iteration over the clusters is:
//
//    PangoGlyphItemIter cluster_iter;
//    gboolean have_cluster;
//
//    for (have_cluster = pango_glyph_item_iter_init_start (&cluster_iter,
//                                                          glyph_item, text);
//         have_cluster;
//         have_cluster = pango_glyph_item_iter_next_cluster (&cluster_iter))
//    {
//      ...
//    }
//
//
// Note that text is the start of the text for layout, which is then indexed by
// glyph_item->item->offset to get to the text of glyph_item. The start_index
// and end_index values can directly index into text. The start_glyph,
// end_glyph, start_char, and end_char values however are zero-based for the
// glyph_item. For each cluster, the item pointed at by the start variables is
// included in the cluster while the one pointed at by end variables is not.
//
// None of the members of a PangoGlyphItemIter should be modified manually.
type GlyphItemIter struct {
	nocopy gextras.NoCopy
	native *C.PangoGlyphItemIter
}

func marshalGlyphItemIter(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return &GlyphItemIter{native: (*C.PangoGlyphItemIter)(unsafe.Pointer(b))}, nil
}

func (g *GlyphItemIter) GlyphItem() *GlyphItem {
	var v *GlyphItem // out
	v = (*GlyphItem)(gextras.NewStructNative(unsafe.Pointer(g.native.glyph_item)))
	return v
}

func (g *GlyphItemIter) Text() string {
	var v string // out
	v = C.GoString((*C.gchar)(unsafe.Pointer(g.native.text)))
	return v
}

func (g *GlyphItemIter) StartGlyph() int {
	var v int // out
	v = int(g.native.start_glyph)
	return v
}

func (g *GlyphItemIter) StartIndex() int {
	var v int // out
	v = int(g.native.start_index)
	return v
}

func (g *GlyphItemIter) StartChar() int {
	var v int // out
	v = int(g.native.start_char)
	return v
}

func (g *GlyphItemIter) EndGlyph() int {
	var v int // out
	v = int(g.native.end_glyph)
	return v
}

func (g *GlyphItemIter) EndIndex() int {
	var v int // out
	v = int(g.native.end_index)
	return v
}

func (g *GlyphItemIter) EndChar() int {
	var v int // out
	v = int(g.native.end_char)
	return v
}

// Copy: make a shallow copy of an existing PangoGlyphItemIter structure.
func (orig *GlyphItemIter) Copy() *GlyphItemIter {
	var _arg0 *C.PangoGlyphItemIter // out
	var _cret *C.PangoGlyphItemIter // in

	if orig != nil {
		_arg0 = (*C.PangoGlyphItemIter)(gextras.StructNative(unsafe.Pointer(orig)))
	}

	_cret = C.pango_glyph_item_iter_copy(_arg0)

	runtime.KeepAlive(orig)

	var _glyphItemIter *GlyphItemIter // out

	if _cret != nil {
		_glyphItemIter = (*GlyphItemIter)(gextras.NewStructNative(unsafe.Pointer(_cret)))
		runtime.SetFinalizer(_glyphItemIter, func(v *GlyphItemIter) {
			C.pango_glyph_item_iter_free((*C.PangoGlyphItemIter)(gextras.StructNative(unsafe.Pointer(v))))
		})
	}

	return _glyphItemIter
}

// InitEnd initializes a PangoGlyphItemIter structure to point to the last
// cluster in a glyph item.
//
// See PangoGlyphItemIter for details of cluster orders.
func (iter *GlyphItemIter) InitEnd(glyphItem *GlyphItem, text string) bool {
	var _arg0 *C.PangoGlyphItemIter // out
	var _arg1 *C.PangoGlyphItem     // out
	var _arg2 *C.char               // out
	var _cret C.gboolean            // in

	_arg0 = (*C.PangoGlyphItemIter)(gextras.StructNative(unsafe.Pointer(iter)))
	_arg1 = (*C.PangoGlyphItem)(gextras.StructNative(unsafe.Pointer(glyphItem)))
	_arg2 = (*C.char)(unsafe.Pointer(C.CString(text)))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.pango_glyph_item_iter_init_end(_arg0, _arg1, _arg2)

	runtime.KeepAlive(iter)
	runtime.KeepAlive(glyphItem)
	runtime.KeepAlive(text)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// InitStart initializes a PangoGlyphItemIter structure to point to the first
// cluster in a glyph item.
//
// See PangoGlyphItemIter for details of cluster orders.
func (iter *GlyphItemIter) InitStart(glyphItem *GlyphItem, text string) bool {
	var _arg0 *C.PangoGlyphItemIter // out
	var _arg1 *C.PangoGlyphItem     // out
	var _arg2 *C.char               // out
	var _cret C.gboolean            // in

	_arg0 = (*C.PangoGlyphItemIter)(gextras.StructNative(unsafe.Pointer(iter)))
	_arg1 = (*C.PangoGlyphItem)(gextras.StructNative(unsafe.Pointer(glyphItem)))
	_arg2 = (*C.char)(unsafe.Pointer(C.CString(text)))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.pango_glyph_item_iter_init_start(_arg0, _arg1, _arg2)

	runtime.KeepAlive(iter)
	runtime.KeepAlive(glyphItem)
	runtime.KeepAlive(text)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// NextCluster advances the iterator to the next cluster in the glyph item.
//
// See PangoGlyphItemIter for details of cluster orders.
func (iter *GlyphItemIter) NextCluster() bool {
	var _arg0 *C.PangoGlyphItemIter // out
	var _cret C.gboolean            // in

	_arg0 = (*C.PangoGlyphItemIter)(gextras.StructNative(unsafe.Pointer(iter)))

	_cret = C.pango_glyph_item_iter_next_cluster(_arg0)

	runtime.KeepAlive(iter)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// PrevCluster moves the iterator to the preceding cluster in the glyph item.
// See PangoGlyphItemIter for details of cluster orders.
func (iter *GlyphItemIter) PrevCluster() bool {
	var _arg0 *C.PangoGlyphItemIter // out
	var _cret C.gboolean            // in

	_arg0 = (*C.PangoGlyphItemIter)(gextras.StructNative(unsafe.Pointer(iter)))

	_cret = C.pango_glyph_item_iter_prev_cluster(_arg0)

	runtime.KeepAlive(iter)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
