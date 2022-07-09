// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// extern GObject* _gotk4_gtk4_BuildableIface_get_internal_child(void*, void*, void*);
// extern char* _gotk4_gtk4_BuildableIface_get_id(void*);
// extern gboolean _gotk4_gtk4_BuildableIface_custom_tag_start(void*, void*, void*, void*, void*, void*);
// extern void _gotk4_gtk4_BuildableIface_add_child(void*, void*, void*, void*);
// extern void _gotk4_gtk4_BuildableIface_custom_finished(void*, void*, void*, void*, gpointer);
// extern void _gotk4_gtk4_BuildableIface_custom_tag_end(void*, void*, void*, void*, gpointer);
// extern void _gotk4_gtk4_BuildableIface_parser_finished(void*, void*);
// extern void _gotk4_gtk4_BuildableIface_set_buildable_property(void*, void*, void*, void*);
// extern void _gotk4_gtk4_BuildableIface_set_id(void*, void*);
import "C"

// GTypeBuildable returns the GType for the type Buildable.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeBuildable() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "Buildable").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalBuildable)
	return gtype
}

// BuildableOverrider contains methods that are overridable.
type BuildableOverrider interface {
	// AddChild adds a child to buildable. type is an optional string describing
	// how the child should be added.
	//
	// The function takes the following parameters:
	//
	//    - builder: Builder.
	//    - child to add.
	//    - typ (optional): kind of child or NULL.
	//
	AddChild(builder *Builder, child *coreglib.Object, typ string)
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
	CustomFinished(builder *Builder, child *coreglib.Object, tagname string, data unsafe.Pointer)
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
	CustomTagEnd(builder *Builder, child *coreglib.Object, tagname string, data unsafe.Pointer)
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
	CustomTagStart(builder *Builder, child *coreglib.Object, tagname string) (*BuildableParser, unsafe.Pointer, bool)
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
	InternalChild(builder *Builder, childname string) *coreglib.Object
	// The function takes the following parameters:
	//
	ParserFinished(builder *Builder)
	// The function takes the following parameters:
	//
	//    - builder
	//    - name
	//    - value
	//
	SetBuildableProperty(builder *Builder, name string, value *coreglib.Value)
	// The function takes the following parameters:
	//
	SetID(id string)
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
//
// Buildable wraps an interface. This means the user can get the
// underlying type by calling Cast().
type Buildable struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*Buildable)(nil)
)

// Buildabler describes Buildable's interface methods.
type Buildabler interface {
	coreglib.Objector

	// BuildableID gets the ID of the buildable object.
	BuildableID() string
}

var _ Buildabler = (*Buildable)(nil)

func ifaceInitBuildabler(gifacePtr, data C.gpointer) {
	iface := girepository.MustFind("Gtk", "BuildableIface")
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gifacePtr), pclass.StructFieldOffset("add_child"))) = unsafe.Pointer(C._gotk4_gtk4_BuildableIface_add_child)
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gifacePtr), pclass.StructFieldOffset("custom_finished"))) = unsafe.Pointer(C._gotk4_gtk4_BuildableIface_custom_finished)
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gifacePtr), pclass.StructFieldOffset("custom_tag_end"))) = unsafe.Pointer(C._gotk4_gtk4_BuildableIface_custom_tag_end)
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gifacePtr), pclass.StructFieldOffset("custom_tag_start"))) = unsafe.Pointer(C._gotk4_gtk4_BuildableIface_custom_tag_start)
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gifacePtr), pclass.StructFieldOffset("get_id"))) = unsafe.Pointer(C._gotk4_gtk4_BuildableIface_get_id)
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gifacePtr), pclass.StructFieldOffset("get_internal_child"))) = unsafe.Pointer(C._gotk4_gtk4_BuildableIface_get_internal_child)
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gifacePtr), pclass.StructFieldOffset("parser_finished"))) = unsafe.Pointer(C._gotk4_gtk4_BuildableIface_parser_finished)
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gifacePtr), pclass.StructFieldOffset("set_buildable_property"))) = unsafe.Pointer(C._gotk4_gtk4_BuildableIface_set_buildable_property)
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gifacePtr), pclass.StructFieldOffset("set_id"))) = unsafe.Pointer(C._gotk4_gtk4_BuildableIface_set_id)
}

//export _gotk4_gtk4_BuildableIface_add_child
func _gotk4_gtk4_BuildableIface_add_child(arg0 *C.void, arg1 *C.void, arg2 *C.void, arg3 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(BuildableOverrider)

	var _builder *Builder       // out
	var _child *coreglib.Object // out
	var _typ string             // out

	_builder = wrapBuilder(coreglib.Take(unsafe.Pointer(arg1)))
	_child = coreglib.Take(unsafe.Pointer(arg2))
	if arg3 != nil {
		_typ = C.GoString((*C.gchar)(unsafe.Pointer(arg3)))
	}

	iface.AddChild(_builder, _child, _typ)
}

//export _gotk4_gtk4_BuildableIface_custom_finished
func _gotk4_gtk4_BuildableIface_custom_finished(arg0 *C.void, arg1 *C.void, arg2 *C.void, arg3 *C.void, arg4 C.gpointer) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(BuildableOverrider)

	var _builder *Builder       // out
	var _child *coreglib.Object // out
	var _tagname string         // out
	var _data unsafe.Pointer    // out

	_builder = wrapBuilder(coreglib.Take(unsafe.Pointer(arg1)))
	if arg2 != nil {
		_child = coreglib.Take(unsafe.Pointer(arg2))
	}
	_tagname = C.GoString((*C.gchar)(unsafe.Pointer(arg3)))
	_data = (unsafe.Pointer)(unsafe.Pointer(arg4))

	iface.CustomFinished(_builder, _child, _tagname, _data)
}

//export _gotk4_gtk4_BuildableIface_custom_tag_end
func _gotk4_gtk4_BuildableIface_custom_tag_end(arg0 *C.void, arg1 *C.void, arg2 *C.void, arg3 *C.void, arg4 C.gpointer) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(BuildableOverrider)

	var _builder *Builder       // out
	var _child *coreglib.Object // out
	var _tagname string         // out
	var _data unsafe.Pointer    // out

	_builder = wrapBuilder(coreglib.Take(unsafe.Pointer(arg1)))
	if arg2 != nil {
		_child = coreglib.Take(unsafe.Pointer(arg2))
	}
	_tagname = C.GoString((*C.gchar)(unsafe.Pointer(arg3)))
	_data = (unsafe.Pointer)(unsafe.Pointer(arg4))

	iface.CustomTagEnd(_builder, _child, _tagname, _data)
}

//export _gotk4_gtk4_BuildableIface_custom_tag_start
func _gotk4_gtk4_BuildableIface_custom_tag_start(arg0 *C.void, arg1 *C.void, arg2 *C.void, arg3 *C.void, arg4 *C.void, arg5 *C.void) (cret C.gboolean) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(BuildableOverrider)

	var _builder *Builder       // out
	var _child *coreglib.Object // out
	var _tagname string         // out

	_builder = wrapBuilder(coreglib.Take(unsafe.Pointer(arg1)))
	if arg2 != nil {
		_child = coreglib.Take(unsafe.Pointer(arg2))
	}
	_tagname = C.GoString((*C.gchar)(unsafe.Pointer(arg3)))

	parser, data, ok := iface.CustomTagStart(_builder, _child, _tagname)

	*arg4 = (*C.void)(gextras.StructNative(unsafe.Pointer(parser)))
	*arg5 = (*C.void)(unsafe.Pointer(data))
	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gtk4_BuildableIface_get_id
func _gotk4_gtk4_BuildableIface_get_id(arg0 *C.void) (cret *C.char) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(BuildableOverrider)

	utf8 := iface.ID()

	cret = (*C.void)(unsafe.Pointer(C.CString(utf8)))
	defer C.free(unsafe.Pointer(cret))

	return cret
}

//export _gotk4_gtk4_BuildableIface_get_internal_child
func _gotk4_gtk4_BuildableIface_get_internal_child(arg0 *C.void, arg1 *C.void, arg2 *C.void) (cret *C.GObject) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(BuildableOverrider)

	var _builder *Builder // out
	var _childname string // out

	_builder = wrapBuilder(coreglib.Take(unsafe.Pointer(arg1)))
	_childname = C.GoString((*C.gchar)(unsafe.Pointer(arg2)))

	object := iface.InternalChild(_builder, _childname)

	cret = (*C.void)(unsafe.Pointer(object.Native()))

	return cret
}

//export _gotk4_gtk4_BuildableIface_parser_finished
func _gotk4_gtk4_BuildableIface_parser_finished(arg0 *C.void, arg1 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(BuildableOverrider)

	var _builder *Builder // out

	_builder = wrapBuilder(coreglib.Take(unsafe.Pointer(arg1)))

	iface.ParserFinished(_builder)
}

//export _gotk4_gtk4_BuildableIface_set_buildable_property
func _gotk4_gtk4_BuildableIface_set_buildable_property(arg0 *C.void, arg1 *C.void, arg2 *C.void, arg3 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(BuildableOverrider)

	var _builder *Builder      // out
	var _name string           // out
	var _value *coreglib.Value // out

	_builder = wrapBuilder(coreglib.Take(unsafe.Pointer(arg1)))
	_name = C.GoString((*C.gchar)(unsafe.Pointer(arg2)))
	_value = coreglib.ValueFromNative(unsafe.Pointer(arg3))

	iface.SetBuildableProperty(_builder, _name, _value)
}

//export _gotk4_gtk4_BuildableIface_set_id
func _gotk4_gtk4_BuildableIface_set_id(arg0 *C.void, arg1 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(BuildableOverrider)

	var _id string // out

	_id = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))

	iface.SetID(_id)
}

func wrapBuildable(obj *coreglib.Object) *Buildable {
	return &Buildable{
		Object: obj,
	}
}

func marshalBuildable(p uintptr) (interface{}, error) {
	return wrapBuildable(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
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
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(buildable).Native()))

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

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
	native unsafe.Pointer
}
