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
// extern GList* _gotk4_gtk4_CellLayoutIface_get_cells(GtkCellLayout*);
// extern GtkCellArea* _gotk4_gtk4_CellLayoutIface_get_area(GtkCellLayout*);
// extern void _gotk4_gtk4_CellLayoutIface_clear(GtkCellLayout*);
// extern void _gotk4_gtk4_CellLayoutIface_clear_attributes(GtkCellLayout*, GtkCellRenderer*);
// extern void _gotk4_gtk4_CellLayoutIface_pack_end(GtkCellLayout*, GtkCellRenderer*, gboolean);
// extern void _gotk4_gtk4_CellLayoutIface_pack_start(GtkCellLayout*, GtkCellRenderer*, gboolean);
import "C"

// glib.Type values for gtkcelllayout.go.
var GTypeCellLayout = coreglib.Type(C.gtk_cell_layout_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeCellLayout, F: marshalCellLayout},
	})
}

// CellLayoutDataFunc: function which should set the value of cell_layout’s cell
// renderer(s) as appropriate.
type CellLayoutDataFunc func(cellLayout CellLayouter, cell CellRendererer, treeModel TreeModeller, iter *TreeIter)

//export _gotk4_gtk4_CellLayoutDataFunc
func _gotk4_gtk4_CellLayoutDataFunc(arg1 *C.GtkCellLayout, arg2 *C.GtkCellRenderer, arg3 *C.GtkTreeModel, arg4 *C.GtkTreeIter, arg5 C.gpointer) {
	var fn CellLayoutDataFunc
	{
		v := gbox.Get(uintptr(arg5))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(CellLayoutDataFunc)
	}

	var _cellLayout CellLayouter // out
	var _cell CellRendererer     // out
	var _treeModel TreeModeller  // out
	var _iter *TreeIter          // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gtk.CellLayouter is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(CellLayouter)
			return ok
		})
		rv, ok := casted.(CellLayouter)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.CellLayouter")
		}
		_cellLayout = rv
	}
	{
		objptr := unsafe.Pointer(arg2)
		if objptr == nil {
			panic("object of type gtk.CellRendererer is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(CellRendererer)
			return ok
		})
		rv, ok := casted.(CellRendererer)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.CellRendererer")
		}
		_cell = rv
	}
	{
		objptr := unsafe.Pointer(arg3)
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
		_treeModel = rv
	}
	_iter = (*TreeIter)(gextras.NewStructNative(unsafe.Pointer(arg4)))

	fn(_cellLayout, _cell, _treeModel, _iter)
}

// CellLayoutOverrider contains methods that are overridable.
type CellLayoutOverrider interface {
	// Clear unsets all the mappings on all renderers on cell_layout and removes
	// all renderers from cell_layout.
	Clear()
	// ClearAttributes clears all existing attributes previously set with
	// gtk_cell_layout_set_attributes().
	//
	// The function takes the following parameters:
	//
	//    - cell to clear the attribute mapping on.
	//
	ClearAttributes(cell CellRendererer)
	// Area returns the underlying CellArea which might be cell_layout if called
	// on a CellArea or might be NULL if no CellArea is used by cell_layout.
	//
	// The function returns the following values:
	//
	//    - cellArea (optional): cell area used by cell_layout, or NULL in case
	//      no cell area is used.
	//
	Area() CellAreaer
	// Cells returns the cell renderers which have been added to cell_layout.
	//
	// The function returns the following values:
	//
	//    - list: a list of cell renderers. The list, but not the renderers has
	//      been newly allocated and should be freed with g_list_free() when no
	//      longer needed.
	//
	Cells() []CellRendererer
	// PackEnd adds the cell to the end of cell_layout. If expand is FALSE, then
	// the cell is allocated no more space than it needs. Any unused space is
	// divided evenly between cells for which expand is TRUE.
	//
	// Note that reusing the same cell renderer is not supported.
	//
	// The function takes the following parameters:
	//
	//    - cell: CellRenderer.
	//    - expand: TRUE if cell is to be given extra space allocated to
	//      cell_layout.
	//
	PackEnd(cell CellRendererer, expand bool)
	// PackStart packs the cell into the beginning of cell_layout. If expand is
	// FALSE, then the cell is allocated no more space than it needs. Any unused
	// space is divided evenly between cells for which expand is TRUE.
	//
	// Note that reusing the same cell renderer is not supported.
	//
	// The function takes the following parameters:
	//
	//    - cell: CellRenderer.
	//    - expand: TRUE if cell is to be given extra space allocated to
	//      cell_layout.
	//
	PackStart(cell CellRendererer, expand bool)
}

// CellLayout: interface for packing cells
//
// CellLayout is an interface to be implemented by all objects which want to
// provide a TreeViewColumn like API for packing cells, setting attributes and
// data funcs.
//
// One of the notable features provided by implementations of GtkCellLayout are
// attributes. Attributes let you set the properties in flexible ways. They can
// just be set to constant values like regular properties. But they can also be
// mapped to a column of the underlying tree model with
// gtk_cell_layout_set_attributes(), which means that the value of the attribute
// can change from cell to cell as they are rendered by the cell renderer.
// Finally, it is possible to specify a function with
// gtk_cell_layout_set_cell_data_func() that is called to determine the value of
// the attribute for each cell that is rendered.
//
//
// GtkCellLayouts as GtkBuildable
//
// Implementations of GtkCellLayout which also implement the GtkBuildable
// interface (CellView, IconView, ComboBox, EntryCompletion, TreeViewColumn)
// accept GtkCellRenderer objects as <child> elements in UI definitions. They
// support a custom <attributes> element for their children, which can contain
// multiple <attribute> elements. Each <attribute> element has a name attribute
// which specifies a property of the cell renderer; the content of the element
// is the attribute value.
//
// This is an example of a UI definition fragment specifying attributes:
//
//    <object class="GtkCellView">
//      <child>
//        <object class="GtkCellRendererText"/>
//        <attributes>
//          <attribute name="text">0</attribute>
//        </attributes>
//      </child>"
//    </object>
//
// Furthermore for implementations of GtkCellLayout that use a CellArea to lay
// out cells (all GtkCellLayouts in GTK use a GtkCellArea) [cell
// properties][cell-properties] can also be defined in the format by specifying
// the custom <cell-packing> attribute which can contain multiple <property>
// elements defined in the normal way.
//
// Here is a UI definition fragment specifying cell properties:
//
//    <object class="GtkTreeViewColumn">
//      <child>
//        <object class="GtkCellRendererText"/>
//        <cell-packing>
//          <property name="align">True</property>
//          <property name="expand">False</property>
//        </cell-packing>
//      </child>"
//    </object>
//
//
// Subclassing GtkCellLayout implementations
//
// When subclassing a widget that implements CellLayout like IconView or
// ComboBox, there are some considerations related to the fact that these
// widgets internally use a CellArea. The cell area is exposed as a
// construct-only property by these widgets. This means that it is possible to
// e.g. do
//
//    static void
//    my_combo_box_init (MyComboBox *b)
//    {
//      GtkCellRenderer *cell;
//
//      cell = gtk_cell_renderer_pixbuf_new ();
//      // The following call causes the default cell area for combo boxes,
//      // a GtkCellAreaBox, to be instantiated
//      gtk_cell_layout_pack_start (GTK_CELL_LAYOUT (b), cell, FALSE);
//      ...
//    }
//
//    GtkWidget *
//    my_combo_box_new (GtkCellArea *area)
//    {
//      // This call is going to cause a warning about area being ignored
//      return g_object_new (MY_TYPE_COMBO_BOX, "cell-area", area, NULL);
//    }
//
// If supporting alternative cell areas with your derived widget is not
// important, then this does not have to concern you. If you want to support
// alternative cell areas, you can do so by moving the problematic calls out of
// init() and into a constructor() for your class.
//
// CellLayout wraps an interface. This means the user can get the
// underlying type by calling Cast().
type CellLayout struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*CellLayout)(nil)
)

// CellLayouter describes CellLayout's interface methods.
type CellLayouter interface {
	coreglib.Objector

	// Clear unsets all the mappings on all renderers on cell_layout and removes
	// all renderers from cell_layout.
	Clear()
	// ClearAttributes clears all existing attributes previously set with
	// gtk_cell_layout_set_attributes().
	ClearAttributes(cell CellRendererer)
	// Area returns the underlying CellArea which might be cell_layout if called
	// on a CellArea or might be NULL if no CellArea is used by cell_layout.
	Area() CellAreaer
	// Cells returns the cell renderers which have been added to cell_layout.
	Cells() []CellRendererer
	// PackEnd adds the cell to the end of cell_layout.
	PackEnd(cell CellRendererer, expand bool)
	// PackStart packs the cell into the beginning of cell_layout.
	PackStart(cell CellRendererer, expand bool)
}

var _ CellLayouter = (*CellLayout)(nil)

func ifaceInitCellLayouter(gifacePtr, data C.gpointer) {
	iface := (*C.GtkCellLayoutIface)(unsafe.Pointer(gifacePtr))
	iface.clear = (*[0]byte)(C._gotk4_gtk4_CellLayoutIface_clear)
	iface.clear_attributes = (*[0]byte)(C._gotk4_gtk4_CellLayoutIface_clear_attributes)
	iface.get_area = (*[0]byte)(C._gotk4_gtk4_CellLayoutIface_get_area)
	iface.get_cells = (*[0]byte)(C._gotk4_gtk4_CellLayoutIface_get_cells)
	iface.pack_end = (*[0]byte)(C._gotk4_gtk4_CellLayoutIface_pack_end)
	iface.pack_start = (*[0]byte)(C._gotk4_gtk4_CellLayoutIface_pack_start)
}

//export _gotk4_gtk4_CellLayoutIface_clear
func _gotk4_gtk4_CellLayoutIface_clear(arg0 *C.GtkCellLayout) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(CellLayoutOverrider)

	iface.Clear()
}

//export _gotk4_gtk4_CellLayoutIface_clear_attributes
func _gotk4_gtk4_CellLayoutIface_clear_attributes(arg0 *C.GtkCellLayout, arg1 *C.GtkCellRenderer) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(CellLayoutOverrider)

	var _cell CellRendererer // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gtk.CellRendererer is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(CellRendererer)
			return ok
		})
		rv, ok := casted.(CellRendererer)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.CellRendererer")
		}
		_cell = rv
	}

	iface.ClearAttributes(_cell)
}

//export _gotk4_gtk4_CellLayoutIface_get_area
func _gotk4_gtk4_CellLayoutIface_get_area(arg0 *C.GtkCellLayout) (cret *C.GtkCellArea) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(CellLayoutOverrider)

	cellArea := iface.Area()

	if cellArea != nil {
		cret = (*C.void)(unsafe.Pointer(coreglib.InternObject(cellArea).Native()))
	}

	return cret
}

//export _gotk4_gtk4_CellLayoutIface_get_cells
func _gotk4_gtk4_CellLayoutIface_get_cells(arg0 *C.GtkCellLayout) (cret *C.GList) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(CellLayoutOverrider)

	list := iface.Cells()

	for i := len(list) - 1; i >= 0; i-- {
		src := list[i]
		var dst *C.void // out
		dst = (*C.void)(unsafe.Pointer(coreglib.InternObject(src).Native()))
		cret = C.g_list_prepend(cret, C.gpointer(unsafe.Pointer(dst)))
	}

	return cret
}

//export _gotk4_gtk4_CellLayoutIface_pack_end
func _gotk4_gtk4_CellLayoutIface_pack_end(arg0 *C.GtkCellLayout, arg1 *C.GtkCellRenderer, arg2 C.gboolean) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(CellLayoutOverrider)

	var _cell CellRendererer // out
	var _expand bool         // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gtk.CellRendererer is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(CellRendererer)
			return ok
		})
		rv, ok := casted.(CellRendererer)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.CellRendererer")
		}
		_cell = rv
	}
	if arg2 != 0 {
		_expand = true
	}

	iface.PackEnd(_cell, _expand)
}

//export _gotk4_gtk4_CellLayoutIface_pack_start
func _gotk4_gtk4_CellLayoutIface_pack_start(arg0 *C.GtkCellLayout, arg1 *C.GtkCellRenderer, arg2 C.gboolean) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(CellLayoutOverrider)

	var _cell CellRendererer // out
	var _expand bool         // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gtk.CellRendererer is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(CellRendererer)
			return ok
		})
		rv, ok := casted.(CellRendererer)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.CellRendererer")
		}
		_cell = rv
	}
	if arg2 != 0 {
		_expand = true
	}

	iface.PackStart(_cell, _expand)
}

func wrapCellLayout(obj *coreglib.Object) *CellLayout {
	return &CellLayout{
		Object: obj,
	}
}

func marshalCellLayout(p uintptr) (interface{}, error) {
	return wrapCellLayout(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Clear unsets all the mappings on all renderers on cell_layout and removes all
// renderers from cell_layout.
func (cellLayout *CellLayout) Clear() {
	var args [1]girepository.Argument
	var _arg0 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(cellLayout).Native()))
	*(**CellLayout)(unsafe.Pointer(&args[0])) = _arg0

	runtime.KeepAlive(cellLayout)
}

// ClearAttributes clears all existing attributes previously set with
// gtk_cell_layout_set_attributes().
//
// The function takes the following parameters:
//
//    - cell to clear the attribute mapping on.
//
func (cellLayout *CellLayout) ClearAttributes(cell CellRendererer) {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(cellLayout).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(cell).Native()))
	*(**CellLayout)(unsafe.Pointer(&args[1])) = _arg1

	runtime.KeepAlive(cellLayout)
	runtime.KeepAlive(cell)
}

// Area returns the underlying CellArea which might be cell_layout if called on
// a CellArea or might be NULL if no CellArea is used by cell_layout.
//
// The function returns the following values:
//
//    - cellArea (optional): cell area used by cell_layout, or NULL in case no
//      cell area is used.
//
func (cellLayout *CellLayout) Area() CellAreaer {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(cellLayout).Native()))
	*(**CellLayout)(unsafe.Pointer(&args[0])) = _arg0

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(cellLayout)

	var _cellArea CellAreaer // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(CellAreaer)
				return ok
			})
			rv, ok := casted.(CellAreaer)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.CellAreaer")
			}
			_cellArea = rv
		}
	}

	return _cellArea
}

// Cells returns the cell renderers which have been added to cell_layout.
//
// The function returns the following values:
//
//    - list: a list of cell renderers. The list, but not the renderers has been
//      newly allocated and should be freed with g_list_free() when no longer
//      needed.
//
func (cellLayout *CellLayout) Cells() []CellRendererer {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(cellLayout).Native()))
	*(**CellLayout)(unsafe.Pointer(&args[0])) = _arg0

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(cellLayout)

	var _list []CellRendererer // out

	_list = make([]CellRendererer, 0, gextras.ListSize(unsafe.Pointer(_cret)))
	gextras.MoveList(unsafe.Pointer(_cret), true, func(v unsafe.Pointer) {
		src := (*C.void)(v)
		var dst CellRendererer // out
		{
			objptr := unsafe.Pointer(src)
			if objptr == nil {
				panic("object of type gtk.CellRendererer is nil")
			}

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(CellRendererer)
				return ok
			})
			rv, ok := casted.(CellRendererer)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.CellRendererer")
			}
			dst = rv
		}
		_list = append(_list, dst)
	})

	return _list
}

// PackEnd adds the cell to the end of cell_layout. If expand is FALSE, then the
// cell is allocated no more space than it needs. Any unused space is divided
// evenly between cells for which expand is TRUE.
//
// Note that reusing the same cell renderer is not supported.
//
// The function takes the following parameters:
//
//    - cell: CellRenderer.
//    - expand: TRUE if cell is to be given extra space allocated to cell_layout.
//
func (cellLayout *CellLayout) PackEnd(cell CellRendererer, expand bool) {
	var args [3]girepository.Argument
	var _arg0 *C.void    // out
	var _arg1 *C.void    // out
	var _arg2 C.gboolean // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(cellLayout).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(cell).Native()))
	if expand {
		_arg2 = C.TRUE
	}
	*(**CellLayout)(unsafe.Pointer(&args[1])) = _arg1
	*(*CellRendererer)(unsafe.Pointer(&args[2])) = _arg2

	runtime.KeepAlive(cellLayout)
	runtime.KeepAlive(cell)
	runtime.KeepAlive(expand)
}

// PackStart packs the cell into the beginning of cell_layout. If expand is
// FALSE, then the cell is allocated no more space than it needs. Any unused
// space is divided evenly between cells for which expand is TRUE.
//
// Note that reusing the same cell renderer is not supported.
//
// The function takes the following parameters:
//
//    - cell: CellRenderer.
//    - expand: TRUE if cell is to be given extra space allocated to cell_layout.
//
func (cellLayout *CellLayout) PackStart(cell CellRendererer, expand bool) {
	var args [3]girepository.Argument
	var _arg0 *C.void    // out
	var _arg1 *C.void    // out
	var _arg2 C.gboolean // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(cellLayout).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(cell).Native()))
	if expand {
		_arg2 = C.TRUE
	}
	*(**CellLayout)(unsafe.Pointer(&args[1])) = _arg1
	*(*CellRendererer)(unsafe.Pointer(&args[2])) = _arg2

	runtime.KeepAlive(cellLayout)
	runtime.KeepAlive(cell)
	runtime.KeepAlive(expand)
}
