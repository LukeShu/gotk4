// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime/cgo"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_cell_layout_get_type()), F: marshalCellLayouter},
	})
}

// CellLayoutDataFunc: function which should set the value of @cell_layout’s
// cell renderer(s) as appropriate.
type CellLayoutDataFunc func(cellLayout *CellLayout, cell *CellRenderer, treeModel *TreeModel, iter *TreeIter, data cgo.Handle)

//export gotk4_CellLayoutDataFunc
func gotk4_CellLayoutDataFunc(arg0 *C.GtkCellLayout, arg1 *C.GtkCellRenderer, arg2 *C.GtkTreeModel, arg3 *C.GtkTreeIter, arg4 C.gpointer) {
	v := gbox.Get(uintptr(arg4))
	if v == nil {
		panic(`callback not found`)
	}

	var cellLayout *CellLayout // out
	var cell *CellRenderer     // out
	var treeModel *TreeModel   // out
	var iter *TreeIter         // out
	var data cgo.Handle        // out

	cellLayout = (gextras.CastObject(externglib.Take(unsafe.Pointer(arg0)))).(*CellLayout)
	cell = (gextras.CastObject(externglib.Take(unsafe.Pointer(arg1)))).(*CellRenderer)
	treeModel = (gextras.CastObject(externglib.Take(unsafe.Pointer(arg2)))).(*TreeModel)
	iter = (*TreeIter)(unsafe.Pointer(arg3))
	data = (cgo.Handle)(unsafe.Pointer(arg4))

	fn := v.(CellLayoutDataFunc)
	fn(cellLayout, cell, treeModel, iter, data)
}

// CellLayoutOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type CellLayoutOverrider interface {
	// AddAttribute adds an attribute mapping to the list in @cell_layout.
	//
	// The @column is the column of the model to get a value from, and the
	// @attribute is the parameter on @cell to be set from the value. So for
	// example if column 2 of the model contains strings, you could have the
	// “text” attribute of a CellRendererText get its values from column 2.
	AddAttribute(cell CellRendererer, attribute string, column int)
	// Clear unsets all the mappings on all renderers on @cell_layout and
	// removes all renderers from @cell_layout.
	Clear()
	// ClearAttributes clears all existing attributes previously set with
	// gtk_cell_layout_set_attributes().
	ClearAttributes(cell CellRendererer)
	// Area returns the underlying CellArea which might be @cell_layout if
	// called on a CellArea or might be nil if no CellArea is used by
	// @cell_layout.
	Area() *CellArea
	// PackEnd adds the @cell to the end of @cell_layout. If @expand is false,
	// then the @cell is allocated no more space than it needs. Any unused space
	// is divided evenly between cells for which @expand is true.
	//
	// Note that reusing the same cell renderer is not supported.
	PackEnd(cell CellRendererer, expand bool)
	// PackStart packs the @cell into the beginning of @cell_layout. If @expand
	// is false, then the @cell is allocated no more space than it needs. Any
	// unused space is divided evenly between cells for which @expand is true.
	//
	// Note that reusing the same cell renderer is not supported.
	PackStart(cell CellRendererer, expand bool)
	// Reorder re-inserts @cell at @position.
	//
	// Note that @cell has already to be packed into @cell_layout for this to
	// function properly.
	Reorder(cell CellRendererer, position int)
}

// CellLayouter describes CellLayout's methods.
type CellLayouter interface {
	// AddAttribute adds an attribute mapping to the list in @cell_layout.
	AddAttribute(cell CellRendererer, attribute string, column int)
	// Clear unsets all the mappings on all renderers on @cell_layout and
	// removes all renderers from @cell_layout.
	Clear()
	// ClearAttributes clears all existing attributes previously set with
	// gtk_cell_layout_set_attributes().
	ClearAttributes(cell CellRendererer)
	// Area returns the underlying CellArea which might be @cell_layout if
	// called on a CellArea or might be nil if no CellArea is used by
	// @cell_layout.
	Area() *CellArea
	// PackEnd adds the @cell to the end of @cell_layout.
	PackEnd(cell CellRendererer, expand bool)
	// PackStart packs the @cell into the beginning of @cell_layout.
	PackStart(cell CellRendererer, expand bool)
	// Reorder re-inserts @cell at @position.
	Reorder(cell CellRendererer, position int)
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
type CellLayout struct {
	*externglib.Object
}

var (
	_ CellLayouter    = (*CellLayout)(nil)
	_ gextras.Nativer = (*CellLayout)(nil)
)

func wrapCellLayout(obj *externglib.Object) CellLayouter {
	return &CellLayout{
		Object: obj,
	}
}

func marshalCellLayouter(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapCellLayout(obj), nil
}

// AddAttribute adds an attribute mapping to the list in @cell_layout.
//
// The @column is the column of the model to get a value from, and the
// @attribute is the parameter on @cell to be set from the value. So for example
// if column 2 of the model contains strings, you could have the “text”
// attribute of a CellRendererText get its values from column 2.
func (cellLayout *CellLayout) AddAttribute(cell CellRendererer, attribute string, column int) {
	var _arg0 *C.GtkCellLayout   // out
	var _arg1 *C.GtkCellRenderer // out
	var _arg2 *C.char            // out
	var _arg3 C.int              // out

	_arg0 = (*C.GtkCellLayout)(unsafe.Pointer(cellLayout.Native()))
	_arg1 = (*C.GtkCellRenderer)(unsafe.Pointer((cell).(gextras.Nativer).Native()))
	_arg2 = (*C.char)(unsafe.Pointer(C.CString(attribute)))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = C.int(column)

	C.gtk_cell_layout_add_attribute(_arg0, _arg1, _arg2, _arg3)
}

// Clear unsets all the mappings on all renderers on @cell_layout and removes
// all renderers from @cell_layout.
func (cellLayout *CellLayout) Clear() {
	var _arg0 *C.GtkCellLayout // out

	_arg0 = (*C.GtkCellLayout)(unsafe.Pointer(cellLayout.Native()))

	C.gtk_cell_layout_clear(_arg0)
}

// ClearAttributes clears all existing attributes previously set with
// gtk_cell_layout_set_attributes().
func (cellLayout *CellLayout) ClearAttributes(cell CellRendererer) {
	var _arg0 *C.GtkCellLayout   // out
	var _arg1 *C.GtkCellRenderer // out

	_arg0 = (*C.GtkCellLayout)(unsafe.Pointer(cellLayout.Native()))
	_arg1 = (*C.GtkCellRenderer)(unsafe.Pointer((cell).(gextras.Nativer).Native()))

	C.gtk_cell_layout_clear_attributes(_arg0, _arg1)
}

// Area returns the underlying CellArea which might be @cell_layout if called on
// a CellArea or might be nil if no CellArea is used by @cell_layout.
func (cellLayout *CellLayout) Area() *CellArea {
	var _arg0 *C.GtkCellLayout // out
	var _cret *C.GtkCellArea   // in

	_arg0 = (*C.GtkCellLayout)(unsafe.Pointer(cellLayout.Native()))

	_cret = C.gtk_cell_layout_get_area(_arg0)

	var _cellArea *CellArea // out

	_cellArea = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*CellArea)

	return _cellArea
}

// PackEnd adds the @cell to the end of @cell_layout. If @expand is false, then
// the @cell is allocated no more space than it needs. Any unused space is
// divided evenly between cells for which @expand is true.
//
// Note that reusing the same cell renderer is not supported.
func (cellLayout *CellLayout) PackEnd(cell CellRendererer, expand bool) {
	var _arg0 *C.GtkCellLayout   // out
	var _arg1 *C.GtkCellRenderer // out
	var _arg2 C.gboolean         // out

	_arg0 = (*C.GtkCellLayout)(unsafe.Pointer(cellLayout.Native()))
	_arg1 = (*C.GtkCellRenderer)(unsafe.Pointer((cell).(gextras.Nativer).Native()))
	if expand {
		_arg2 = C.TRUE
	}

	C.gtk_cell_layout_pack_end(_arg0, _arg1, _arg2)
}

// PackStart packs the @cell into the beginning of @cell_layout. If @expand is
// false, then the @cell is allocated no more space than it needs. Any unused
// space is divided evenly between cells for which @expand is true.
//
// Note that reusing the same cell renderer is not supported.
func (cellLayout *CellLayout) PackStart(cell CellRendererer, expand bool) {
	var _arg0 *C.GtkCellLayout   // out
	var _arg1 *C.GtkCellRenderer // out
	var _arg2 C.gboolean         // out

	_arg0 = (*C.GtkCellLayout)(unsafe.Pointer(cellLayout.Native()))
	_arg1 = (*C.GtkCellRenderer)(unsafe.Pointer((cell).(gextras.Nativer).Native()))
	if expand {
		_arg2 = C.TRUE
	}

	C.gtk_cell_layout_pack_start(_arg0, _arg1, _arg2)
}

// Reorder re-inserts @cell at @position.
//
// Note that @cell has already to be packed into @cell_layout for this to
// function properly.
func (cellLayout *CellLayout) Reorder(cell CellRendererer, position int) {
	var _arg0 *C.GtkCellLayout   // out
	var _arg1 *C.GtkCellRenderer // out
	var _arg2 C.int              // out

	_arg0 = (*C.GtkCellLayout)(unsafe.Pointer(cellLayout.Native()))
	_arg1 = (*C.GtkCellRenderer)(unsafe.Pointer((cell).(gextras.Nativer).Native()))
	_arg2 = C.int(position)

	C.gtk_cell_layout_reorder(_arg0, _arg1, _arg2)
}
