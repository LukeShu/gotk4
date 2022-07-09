// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// extern gboolean _gotk4_gtk4_TreeSortableIface_get_sort_column_id(void*, void*, void*);
// extern gboolean _gotk4_gtk4_TreeSortableIface_has_default_sort_func(void*);
// extern int _gotk4_gtk4_TreeIterCompareFunc(void*, void*, void*, gpointer);
// extern void _gotk4_gtk4_TreeSortableIface_sort_column_changed(void*);
// extern void _gotk4_gtk4_TreeSortable_ConnectSortColumnChanged(gpointer, guintptr);
// extern void callbackDelete(gpointer);
import "C"

// GTypeTreeSortable returns the GType for the type TreeSortable.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeTreeSortable() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "TreeSortable").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalTreeSortable)
	return gtype
}

// TREE_SORTABLE_DEFAULT_SORT_COLUMN_ID uses the default sort function in a
// gtk.TreeSortable.
//
// See also: gtk.TreeSortable.SetSortColumnID().
const TREE_SORTABLE_DEFAULT_SORT_COLUMN_ID = -1

// TREE_SORTABLE_UNSORTED_SORT_COLUMN_ID disables sorting in a gtk.TreeSortable.
//
// See also: gtk.TreeSortable.SetSortColumnID().
const TREE_SORTABLE_UNSORTED_SORT_COLUMN_ID = -2

// TreeIterCompareFunc should return a negative integer, zero, or a positive
// integer if a sorts before b, a sorts with b, or a sorts after b respectively.
//
// If two iters compare as equal, their order in the sorted model is undefined.
// In order to ensure that the TreeSortable behaves as expected, the
// GtkTreeIterCompareFunc must define a partial order on the model, i.e. it must
// be reflexive, antisymmetric and transitive.
//
// For example, if model is a product catalogue, then a compare function for the
// “price” column could be one which returns price_of(a) - price_of(b).
type TreeIterCompareFunc func(model TreeModeller, a, b *TreeIter) (gint int32)

//export _gotk4_gtk4_TreeIterCompareFunc
func _gotk4_gtk4_TreeIterCompareFunc(arg1 *C.void, arg2 *C.void, arg3 *C.void, arg4 C.gpointer) (cret C.int) {
	var fn TreeIterCompareFunc
	{
		v := gbox.Get(uintptr(arg4))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(TreeIterCompareFunc)
	}

	var _model TreeModeller // out
	var _a *TreeIter        // out
	var _b *TreeIter        // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gtk.TreeModeller is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(TreeModeller)
			return ok
		})
		rv, ok := casted.(TreeModeller)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.TreeModeller")
		}
		_model = rv
	}
	_a = (*TreeIter)(gextras.NewStructNative(unsafe.Pointer(arg2)))
	_b = (*TreeIter)(gextras.NewStructNative(unsafe.Pointer(arg3)))

	gint := fn(_model, _a, _b)

	cret = C.int(gint)

	return cret
}

// TreeSortable: interface for sortable models used by GtkTreeView
//
// TreeSortable is an interface to be implemented by tree models which support
// sorting. The TreeView uses the methods provided by this interface to sort the
// model.
//
// TreeSortable wraps an interface. This means the user can get the
// underlying type by calling Cast().
type TreeSortable struct {
	_ [0]func() // equal guard
	TreeModel
}

var ()

// TreeSortabler describes TreeSortable's interface methods.
type TreeSortabler interface {
	coreglib.Objector

	// SortColumnID fills in sort_column_id and order with the current sort
	// column and the order.
	SortColumnID() (int32, SortType, bool)
	// HasDefaultSortFunc returns TRUE if the model has a default sort function.
	HasDefaultSortFunc() bool
	// SetDefaultSortFunc sets the default comparison function used when sorting
	// to be sort_func.
	SetDefaultSortFunc(sortFunc TreeIterCompareFunc)
	// SetSortFunc sets the comparison function used when sorting to be
	// sort_func.
	SetSortFunc(sortColumnId int32, sortFunc TreeIterCompareFunc)
	// SortColumnChanged emits a TreeSortable::sort-column-changed signal on
	// sortable.
	SortColumnChanged()

	// Sort-column-changed signal is emitted when the sort column or sort order
	// of sortable is changed.
	ConnectSortColumnChanged(func()) coreglib.SignalHandle
}

var _ TreeSortabler = (*TreeSortable)(nil)

func wrapTreeSortable(obj *coreglib.Object) *TreeSortable {
	return &TreeSortable{
		TreeModel: TreeModel{
			Object: obj,
		},
	}
}

func marshalTreeSortable(p uintptr) (interface{}, error) {
	return wrapTreeSortable(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk4_TreeSortable_ConnectSortColumnChanged
func _gotk4_gtk4_TreeSortable_ConnectSortColumnChanged(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectSortColumnChanged signal is emitted when the sort column or sort order
// of sortable is changed. The signal is emitted before the contents of sortable
// are resorted.
func (sortable *TreeSortable) ConnectSortColumnChanged(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(sortable, "sort-column-changed", false, unsafe.Pointer(C._gotk4_gtk4_TreeSortable_ConnectSortColumnChanged), f)
}

// SortColumnID fills in sort_column_id and order with the current sort column
// and the order. It returns TRUE unless the sort_column_id is
// GTK_TREE_SORTABLE_DEFAULT_SORT_COLUMN_ID or
// GTK_TREE_SORTABLE_UNSORTED_SORT_COLUMN_ID.
//
// The function returns the following values:
//
//    - sortColumnId: sort column id to be filled in.
//    - order to be filled in.
//    - ok: TRUE if the sort column is not one of the special sort column ids.
//
func (sortable *TreeSortable) SortColumnID() (int32, SortType, bool) {
	var _args [1]girepository.Argument
	var _outs [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(sortable).Native()))

	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(sortable)

	var _sortColumnId int32 // out
	var _order SortType     // out
	var _ok bool            // out

	_sortColumnId = *(*int32)(unsafe.Pointer(_outs[0]))
	_order = *(*SortType)(unsafe.Pointer(_outs[1]))
	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _sortColumnId, _order, _ok
}

// HasDefaultSortFunc returns TRUE if the model has a default sort function.
// This is used primarily by GtkTreeViewColumns in order to determine if a model
// can go back to the default state, or not.
//
// The function returns the following values:
//
//    - ok: TRUE, if the model has a default sort function.
//
func (sortable *TreeSortable) HasDefaultSortFunc() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(sortable).Native()))

	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(sortable)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// SetDefaultSortFunc sets the default comparison function used when sorting to
// be sort_func. If the current sort column id of sortable is
// GTK_TREE_SORTABLE_DEFAULT_SORT_COLUMN_ID, then the model will sort using this
// function.
//
// If sort_func is NULL, then there will be no default comparison function. This
// means that once the model has been sorted, it can’t go back to the default
// state. In this case, when the current sort column id of sortable is
// GTK_TREE_SORTABLE_DEFAULT_SORT_COLUMN_ID, the model will be unsorted.
//
// The function takes the following parameters:
//
//    - sortFunc: comparison function.
//
func (sortable *TreeSortable) SetDefaultSortFunc(sortFunc TreeIterCompareFunc) {
	var _args [4]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(sortable).Native()))
	*(*C.gpointer)(unsafe.Pointer(&_args[1])) = (*[0]byte)(C._gotk4_gtk4_TreeIterCompareFunc)
	_args[2] = C.gpointer(gbox.Assign(sortFunc))
	_args[3] = (C.GDestroyNotify)((*[0]byte)(C.callbackDelete))

	runtime.KeepAlive(sortable)
	runtime.KeepAlive(sortFunc)
}

// SetSortFunc sets the comparison function used when sorting to be sort_func.
// If the current sort column id of sortable is the same as sort_column_id, then
// the model will sort using this function.
//
// The function takes the following parameters:
//
//    - sortColumnId: sort column id to set the function for.
//    - sortFunc: comparison function.
//
func (sortable *TreeSortable) SetSortFunc(sortColumnId int32, sortFunc TreeIterCompareFunc) {
	var _args [5]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(sortable).Native()))
	*(*C.int)(unsafe.Pointer(&_args[1])) = C.int(sortColumnId)
	*(*C.gpointer)(unsafe.Pointer(&_args[2])) = (*[0]byte)(C._gotk4_gtk4_TreeIterCompareFunc)
	_args[3] = C.gpointer(gbox.Assign(sortFunc))
	_args[4] = (C.GDestroyNotify)((*[0]byte)(C.callbackDelete))

	runtime.KeepAlive(sortable)
	runtime.KeepAlive(sortColumnId)
	runtime.KeepAlive(sortFunc)
}

// SortColumnChanged emits a TreeSortable::sort-column-changed signal on
// sortable.
func (sortable *TreeSortable) SortColumnChanged() {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(sortable).Native()))

	runtime.KeepAlive(sortable)
}
