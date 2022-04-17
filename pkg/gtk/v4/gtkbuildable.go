// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"runtime/cgo"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
// extern GObject* _gotk4_gtk4_BuildableIface_get_internal_child(GtkBuildable*, GtkBuilder*, char*);
// extern char* _gotk4_gtk4_BuildableIface_get_id(GtkBuildable*);
// extern gboolean _gotk4_gtk4_BuildableIface_custom_tag_start(GtkBuildable*, GtkBuilder*, GObject*, char*, GtkBuildableParser*, gpointer*);
// extern void _gotk4_gtk4_BuildableIface_add_child(GtkBuildable*, GtkBuilder*, GObject*, char*);
// extern void _gotk4_gtk4_BuildableIface_custom_finished(GtkBuildable*, GtkBuilder*, GObject*, char*, gpointer);
// extern void _gotk4_gtk4_BuildableIface_custom_tag_end(GtkBuildable*, GtkBuilder*, GObject*, char*, gpointer);
// extern void _gotk4_gtk4_BuildableIface_parser_finished(GtkBuildable*, GtkBuilder*);
// extern void _gotk4_gtk4_BuildableIface_set_buildable_property(GtkBuildable*, GtkBuilder*, char*, GValue*);
// extern void _gotk4_gtk4_BuildableIface_set_id(GtkBuildable*, char*);
import "C"

// glib.Type values for gtkbuildable.go.
var GTypeBuildable = externglib.Type(C.gtk_buildable_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeBuildable, F: marshalBuildable},
	})
}

// BuildableOverrider contains methods that are overridable.
type BuildableOverrider interface {
	externglib.Objector
	// AddChild adds a child to buildable. type is an optional string describing
	// how the child should be added.
	//
	// The function takes the following parameters:
	//
	//    - builder: Builder.
	//    - child to add.
	//    - typ (optional): kind of child or NULL.
	//
	AddChild(builder *Builder, child *externglib.Object, typ string)
	// CustomFinished: similar to gtk_buildable_parser_finished() but is called
	// once for each custom tag handled by the buildable.
	//
	// The function takes the following parameters:
	//
	//    - builder: Builder.
	//    - child (optional) object or NULL for non-child tags.
	//    - tagname: name of the tag.
	//    - data (optional): user data created in custom_tag_start.
	//
	CustomFinished(builder *Builder, child *externglib.Object, tagname string, data cgo.Handle)
	// CustomTagEnd: called at the end of each custom element handled by the
	// buildable.
	//
	// The function takes the following parameters:
	//
	//    - builder used to construct this object.
	//    - child (optional) object or NULL for non-child tags.
	//    - tagname: name of tag.
	//    - data (optional): user data that will be passed in to parser
	//      functions.
	//
	CustomTagEnd(builder *Builder, child *externglib.Object, tagname string, data cgo.Handle)
	// CustomTagStart: called for each unknown element under <child>.
	//
	// The function takes the following parameters:
	//
	//    - builder used to construct this object.
	//    - child (optional) object or NULL for non-child tags.
	//    - tagname: name of tag.
	//
	// The function returns the following values:
	//
	//    - parser to fill in.
	//    - data (optional): return location for user data that will be passed in
	//      to parser functions.
	//    - ok: TRUE if an object has a custom implementation, FALSE if it
	//      doesn't.
	//
	CustomTagStart(builder *Builder, child *externglib.Object, tagname string) (*BuildableParser, cgo.Handle, bool)
	// The function returns the following values:
	//
	ID() string
	// InternalChild retrieves the internal child called childname of the
	// buildable object.
	//
	// The function takes the following parameters:
	//
	//    - builder: Builder.
	//    - childname: name of child.
	//
	// The function returns the following values:
	//
	//    - object: internal child of the buildable object.
	//
	InternalChild(builder *Builder, childname string) *externglib.Object
	// The function takes the following parameters:
	//
	ParserFinished(builder *Builder)
	// The function takes the following parameters:
	//
	//    - builder
	//    - name
	//    - value
	//
	SetBuildableProperty(builder *Builder, name string, value *externglib.Value)
	// The function takes the following parameters:
	//
	SetID(id string)
}

// WrapBuildableOverrider wraps the BuildableOverrider
// interface implementation to access the instance methods.
func WrapBuildableOverrider(obj BuildableOverrider) *Buildable {
	return wrapBuildable(externglib.BaseObject(obj))
}

// Buildable: GtkBuildable allows objects to extend and customize their
// deserialization from ui files.
//
// The interface includes methods for setting names and properties of objects,
// parsing custom tags and constructing child objects.
//
// The GtkBuildable interface is implemented by all widgets and many of the
// non-widget objects that are provided by GTK. The main user of this interface
// is gtk.Builder. There should be very little need for applications to call any
// of these functions directly.
//
// An object only needs to implement this interface if it needs to extend the
// GtkBuilder XML format or run any extra routines at deserialization time.
type Buildable struct {
	_ [0]func() // equal guard
	*externglib.Object
}

var (
	_ externglib.Objector = (*Buildable)(nil)
)

// Buildabler describes Buildable's interface methods.
type Buildabler interface {
	externglib.Objector

	// BuildableID gets the ID of the buildable object.
	BuildableID() string
}

var _ Buildabler = (*Buildable)(nil)

func ifaceInitBuildabler(gifacePtr, data C.gpointer) {
	iface := (*C.GtkBuildableIface)(unsafe.Pointer(gifacePtr))
	iface.add_child = (*[0]byte)(C._gotk4_gtk4_BuildableIface_add_child)
	iface.custom_finished = (*[0]byte)(C._gotk4_gtk4_BuildableIface_custom_finished)
	iface.custom_tag_end = (*[0]byte)(C._gotk4_gtk4_BuildableIface_custom_tag_end)
	iface.custom_tag_start = (*[0]byte)(C._gotk4_gtk4_BuildableIface_custom_tag_start)
	iface.get_id = (*[0]byte)(C._gotk4_gtk4_BuildableIface_get_id)
	iface.get_internal_child = (*[0]byte)(C._gotk4_gtk4_BuildableIface_get_internal_child)
	iface.parser_finished = (*[0]byte)(C._gotk4_gtk4_BuildableIface_parser_finished)
	iface.set_buildable_property = (*[0]byte)(C._gotk4_gtk4_BuildableIface_set_buildable_property)
	iface.set_id = (*[0]byte)(C._gotk4_gtk4_BuildableIface_set_id)
}

//export _gotk4_gtk4_BuildableIface_add_child
func _gotk4_gtk4_BuildableIface_add_child(arg0 *C.GtkBuildable, arg1 *C.GtkBuilder, arg2 *C.GObject, arg3 *C.char) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(BuildableOverrider)

	var _builder *Builder         // out
	var _child *externglib.Object // out
	var _typ string               // out

	_builder = wrapBuilder(externglib.Take(unsafe.Pointer(arg1)))
	_child = externglib.Take(unsafe.Pointer(arg2))
	if arg3 != nil {
		_typ = C.GoString((*C.gchar)(unsafe.Pointer(arg3)))
	}

	iface.AddChild(_builder, _child, _typ)
}

//export _gotk4_gtk4_BuildableIface_custom_finished
func _gotk4_gtk4_BuildableIface_custom_finished(arg0 *C.GtkBuildable, arg1 *C.GtkBuilder, arg2 *C.GObject, arg3 *C.char, arg4 C.gpointer) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(BuildableOverrider)

	var _builder *Builder         // out
	var _child *externglib.Object // out
	var _tagname string           // out
	var _data cgo.Handle          // out

	_builder = wrapBuilder(externglib.Take(unsafe.Pointer(arg1)))
	if arg2 != nil {
		_child = externglib.Take(unsafe.Pointer(arg2))
	}
	_tagname = C.GoString((*C.gchar)(unsafe.Pointer(arg3)))
	_data = (cgo.Handle)(unsafe.Pointer(arg4))

	iface.CustomFinished(_builder, _child, _tagname, _data)
}

//export _gotk4_gtk4_BuildableIface_custom_tag_end
func _gotk4_gtk4_BuildableIface_custom_tag_end(arg0 *C.GtkBuildable, arg1 *C.GtkBuilder, arg2 *C.GObject, arg3 *C.char, arg4 C.gpointer) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(BuildableOverrider)

	var _builder *Builder         // out
	var _child *externglib.Object // out
	var _tagname string           // out
	var _data cgo.Handle          // out

	_builder = wrapBuilder(externglib.Take(unsafe.Pointer(arg1)))
	if arg2 != nil {
		_child = externglib.Take(unsafe.Pointer(arg2))
	}
	_tagname = C.GoString((*C.gchar)(unsafe.Pointer(arg3)))
	_data = (cgo.Handle)(unsafe.Pointer(arg4))

	iface.CustomTagEnd(_builder, _child, _tagname, _data)
}

//export _gotk4_gtk4_BuildableIface_custom_tag_start
func _gotk4_gtk4_BuildableIface_custom_tag_start(arg0 *C.GtkBuildable, arg1 *C.GtkBuilder, arg2 *C.GObject, arg3 *C.char, arg4 *C.GtkBuildableParser, arg5 *C.gpointer) (cret C.gboolean) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(BuildableOverrider)

	var _builder *Builder         // out
	var _child *externglib.Object // out
	var _tagname string           // out

	_builder = wrapBuilder(externglib.Take(unsafe.Pointer(arg1)))
	if arg2 != nil {
		_child = externglib.Take(unsafe.Pointer(arg2))
	}
	_tagname = C.GoString((*C.gchar)(unsafe.Pointer(arg3)))

	parser, data, ok := iface.CustomTagStart(_builder, _child, _tagname)

	*arg4 = *(*C.GtkBuildableParser)(gextras.StructNative(unsafe.Pointer(parser)))
	*arg5 = (C.gpointer)(unsafe.Pointer(data))
	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gtk4_BuildableIface_get_id
func _gotk4_gtk4_BuildableIface_get_id(arg0 *C.GtkBuildable) (cret *C.char) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(BuildableOverrider)

	utf8 := iface.ID()

	cret = (*C.char)(unsafe.Pointer(C.CString(utf8)))
	defer C.free(unsafe.Pointer(cret))

	return cret
}

//export _gotk4_gtk4_BuildableIface_get_internal_child
func _gotk4_gtk4_BuildableIface_get_internal_child(arg0 *C.GtkBuildable, arg1 *C.GtkBuilder, arg2 *C.char) (cret *C.GObject) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(BuildableOverrider)

	var _builder *Builder // out
	var _childname string // out

	_builder = wrapBuilder(externglib.Take(unsafe.Pointer(arg1)))
	_childname = C.GoString((*C.gchar)(unsafe.Pointer(arg2)))

	object := iface.InternalChild(_builder, _childname)

	cret = (*C.GObject)(unsafe.Pointer(object.Native()))

	return cret
}

//export _gotk4_gtk4_BuildableIface_parser_finished
func _gotk4_gtk4_BuildableIface_parser_finished(arg0 *C.GtkBuildable, arg1 *C.GtkBuilder) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(BuildableOverrider)

	var _builder *Builder // out

	_builder = wrapBuilder(externglib.Take(unsafe.Pointer(arg1)))

	iface.ParserFinished(_builder)
}

//export _gotk4_gtk4_BuildableIface_set_buildable_property
func _gotk4_gtk4_BuildableIface_set_buildable_property(arg0 *C.GtkBuildable, arg1 *C.GtkBuilder, arg2 *C.char, arg3 *C.GValue) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(BuildableOverrider)

	var _builder *Builder        // out
	var _name string             // out
	var _value *externglib.Value // out

	_builder = wrapBuilder(externglib.Take(unsafe.Pointer(arg1)))
	_name = C.GoString((*C.gchar)(unsafe.Pointer(arg2)))
	_value = externglib.ValueFromNative(unsafe.Pointer(arg3))

	iface.SetBuildableProperty(_builder, _name, _value)
}

//export _gotk4_gtk4_BuildableIface_set_id
func _gotk4_gtk4_BuildableIface_set_id(arg0 *C.GtkBuildable, arg1 *C.char) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(BuildableOverrider)

	var _id string // out

	_id = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))

	iface.SetID(_id)
}

func wrapBuildable(obj *externglib.Object) *Buildable {
	return &Buildable{
		Object: obj,
	}
}

func marshalBuildable(p uintptr) (interface{}, error) {
	return wrapBuildable(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// BuildableID gets the ID of the buildable object.
//
// GtkBuilder sets the name based on the ID attribute of the <object> tag used
// to construct the buildable.
//
// The function returns the following values:
//
//    - utf8: ID of the buildable object.
//
func (buildable *Buildable) BuildableID() string {
	var _arg0 *C.GtkBuildable // out
	var _cret *C.char         // in

	_arg0 = (*C.GtkBuildable)(unsafe.Pointer(externglib.InternObject(buildable).Native()))

	_cret = C.gtk_buildable_get_buildable_id(_arg0)
	runtime.KeepAlive(buildable)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// BuildableParser: sub-parser for GtkBuildable implementations.
//
// An instance of this type is always passed by reference.
type BuildableParser struct {
	*buildableParser
}

// buildableParser is the struct that's finalized.
type buildableParser struct {
	native *C.GtkBuildableParser
}
