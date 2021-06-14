// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"runtime"
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 glib-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <glib.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_option_group_get_type()), F: marshalOptionGroup},
	})
}

// OptionArg: the Arg enum values determine which type of extra argument the
// options expect to find. If an option expects an extra argument, it can be
// specified in several ways; with a short option: `-x arg`, with a long option:
// `--name arg` or combined in a single argument: `--name=arg`.
type OptionArg int

const (
	// OptionArgNone: no extra argument. This is useful for simple flags.
	OptionArgNone OptionArg = 0
	// OptionArgString: the option takes a UTF-8 string argument.
	OptionArgString OptionArg = 1
	// OptionArgInt: the option takes an integer argument.
	OptionArgInt OptionArg = 2
	// OptionArgCallback: the option provides a callback (of type ArgFunc) to
	// parse the extra argument.
	OptionArgCallback OptionArg = 3
	// OptionArgFilename: the option takes a filename as argument, which will be
	// in the GLib filename encoding rather than UTF-8.
	OptionArgFilename OptionArg = 4
	// OptionArgStringArray: the option takes a string argument, multiple uses
	// of the option are collected into an array of strings.
	OptionArgStringArray OptionArg = 5
	// OptionArgFilenameArray: the option takes a filename as argument, multiple
	// uses of the option are collected into an array of strings.
	OptionArgFilenameArray OptionArg = 6
	// OptionArgDouble: the option takes a double argument. The argument can be
	// formatted either for the user's locale or for the "C" locale. Since 2.12
	OptionArgDouble OptionArg = 7
	// OptionArgInt64: the option takes a 64-bit integer. Like G_OPTION_ARG_INT
	// but for larger numbers. The number can be in decimal base, or in
	// hexadecimal (when prefixed with `0x`, for example, `0xffffffff`). Since
	// 2.12
	OptionArgInt64 OptionArg = 8
)

// OptionError: error codes returned by option parsing.
type OptionError int

const (
	// OptionErrorUnknownOption: an option was not known to the parser. This
	// error will only be reported, if the parser hasn't been instructed to
	// ignore unknown options, see
	// g_option_context_set_ignore_unknown_options().
	OptionErrorUnknownOption OptionError = 0
	// OptionErrorBadValue: a value couldn't be parsed.
	OptionErrorBadValue OptionError = 1
	// OptionErrorFailed: a ArgFunc callback failed.
	OptionErrorFailed OptionError = 2
)

// OptionFlags flags which modify individual options.
type OptionFlags int

const (
	// OptionFlagsNone: no flags. Since: 2.42.
	OptionFlagsNone OptionFlags = 0
	// OptionFlagsHidden: the option doesn't appear in `--help` output.
	OptionFlagsHidden OptionFlags = 1
	// OptionFlagsInMain: the option appears in the main section of the `--help`
	// output, even if it is defined in a group.
	OptionFlagsInMain OptionFlags = 2
	// OptionFlagsReverse: for options of the G_OPTION_ARG_NONE kind, this flag
	// indicates that the sense of the option is reversed.
	OptionFlagsReverse OptionFlags = 4
	// OptionFlagsNoArg: for options of the G_OPTION_ARG_CALLBACK kind, this
	// flag indicates that the callback does not take any argument (like a
	// G_OPTION_ARG_NONE option). Since 2.8
	OptionFlagsNoArg OptionFlags = 8
	// OptionFlagsFilename: for options of the G_OPTION_ARG_CALLBACK kind, this
	// flag indicates that the argument should be passed to the callback in the
	// GLib filename encoding rather than UTF-8. Since 2.8
	OptionFlagsFilename OptionFlags = 16
	// OptionFlagsOptionalArg: for options of the G_OPTION_ARG_CALLBACK kind,
	// this flag indicates that the argument supply is optional. If no argument
	// is given then data of GOptionParseFunc will be set to NULL. Since 2.8
	OptionFlagsOptionalArg OptionFlags = 32
	// OptionFlagsNoalias: this flag turns off the automatic conflict resolution
	// which prefixes long option names with `groupname-` if there is a
	// conflict. This option should only be used in situations where aliasing is
	// necessary to model some legacy commandline interface. It is not safe to
	// use this option, unless all option groups are under your direct control.
	// Since 2.8.
	OptionFlagsNoalias OptionFlags = 64
)

// OptionContext: a `GOptionContext` struct defines which options are accepted
// by the commandline option parser. The struct has only private fields and
// should not be directly accessed.
type OptionContext struct {
	native C.GOptionContext
}

// WrapOptionContext wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapOptionContext(ptr unsafe.Pointer) *OptionContext {
	if ptr == nil {
		return nil
	}

	return (*OptionContext)(ptr)
}

// Native returns the underlying C source pointer.
func (o *OptionContext) Native() unsafe.Pointer {
	return unsafe.Pointer(&o.native)
}

// AddGroup adds a Group to the @context, so that parsing with @context will
// recognize the options in the group. Note that this will take ownership of the
// @group and thus the @group should not be freed.
func (c *OptionContext) AddGroup(group *OptionGroup) {
	var _arg0 *C.GOptionContext // out
	var _arg1 *C.GOptionGroup   // out

	_arg0 = (*C.GOptionContext)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GOptionGroup)(unsafe.Pointer(group.Native()))

	C.g_option_context_add_group(_arg0, _arg1)
}

// AddMainEntries: a convenience function which creates a main group if it
// doesn't exist, adds the @entries to it and sets the translation domain.
func (c *OptionContext) AddMainEntries(entries []OptionEntry, translationDomain string) {
	var _arg0 *C.GOptionContext // out
	var _arg1 *C.GOptionEntry
	var _arg2 *C.gchar // out

	_arg0 = (*C.GOptionContext)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GOptionEntry)(C.malloc(C.ulong((len(entries) + 1)) * C.ulong(C.sizeof_GOptionEntry)))
	defer C.free(unsafe.Pointer(_arg1))

	{
		out := unsafe.Slice(_arg1, len(entries))
		for i := range entries {
			out[i] = (C.GOptionEntry)(unsafe.Pointer(entries[i].Native()))
		}
	}
	_arg2 = (*C.gchar)(C.CString(translationDomain))
	defer C.free(unsafe.Pointer(_arg2))

	C.g_option_context_add_main_entries(_arg0, _arg1, _arg2)
}

// Free frees context and all the groups which have been added to it.
//
// Please note that parsed arguments need to be freed separately (see Entry).
func (c *OptionContext) Free() {
	var _arg0 *C.GOptionContext // out

	_arg0 = (*C.GOptionContext)(unsafe.Pointer(c.Native()))

	C.g_option_context_free(_arg0)
}

// Description returns the description. See g_option_context_set_description().
func (c *OptionContext) Description() string {
	var _arg0 *C.GOptionContext // out

	_arg0 = (*C.GOptionContext)(unsafe.Pointer(c.Native()))

	var _cret *C.gchar // in

	_cret = C.g_option_context_get_description(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// Help returns a formatted, translated help text for the given context. To
// obtain the text produced by `--help`, call `g_option_context_get_help
// (context, TRUE, NULL)`. To obtain the text produced by `--help-all`, call
// `g_option_context_get_help (context, FALSE, NULL)`. To obtain the help text
// for an option group, call `g_option_context_get_help (context, FALSE,
// group)`.
func (c *OptionContext) Help(mainHelp bool, group *OptionGroup) string {
	var _arg0 *C.GOptionContext // out
	var _arg1 C.gboolean        // out
	var _arg2 *C.GOptionGroup   // out

	_arg0 = (*C.GOptionContext)(unsafe.Pointer(c.Native()))
	if mainHelp {
		_arg1 = C.TRUE
	}
	_arg2 = (*C.GOptionGroup)(unsafe.Pointer(group.Native()))

	var _cret *C.gchar // in

	_cret = C.g_option_context_get_help(_arg0, _arg1, _arg2)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// HelpEnabled returns whether automatic `--help` generation is turned on for
// @context. See g_option_context_set_help_enabled().
func (c *OptionContext) HelpEnabled() bool {
	var _arg0 *C.GOptionContext // out

	_arg0 = (*C.GOptionContext)(unsafe.Pointer(c.Native()))

	var _cret C.gboolean // in

	_cret = C.g_option_context_get_help_enabled(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IgnoreUnknownOptions returns whether unknown options are ignored or not. See
// g_option_context_set_ignore_unknown_options().
func (c *OptionContext) IgnoreUnknownOptions() bool {
	var _arg0 *C.GOptionContext // out

	_arg0 = (*C.GOptionContext)(unsafe.Pointer(c.Native()))

	var _cret C.gboolean // in

	_cret = C.g_option_context_get_ignore_unknown_options(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// MainGroup returns a pointer to the main group of @context.
func (c *OptionContext) MainGroup() *OptionGroup {
	var _arg0 *C.GOptionContext // out

	_arg0 = (*C.GOptionContext)(unsafe.Pointer(c.Native()))

	var _cret *C.GOptionGroup // in

	_cret = C.g_option_context_get_main_group(_arg0)

	var _optionGroup *OptionGroup // out

	_optionGroup = WrapOptionGroup(unsafe.Pointer(_cret))

	return _optionGroup
}

// StrictPosix returns whether strict POSIX code is enabled.
//
// See g_option_context_set_strict_posix() for more information.
func (c *OptionContext) StrictPosix() bool {
	var _arg0 *C.GOptionContext // out

	_arg0 = (*C.GOptionContext)(unsafe.Pointer(c.Native()))

	var _cret C.gboolean // in

	_cret = C.g_option_context_get_strict_posix(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Summary returns the summary. See g_option_context_set_summary().
func (c *OptionContext) Summary() string {
	var _arg0 *C.GOptionContext // out

	_arg0 = (*C.GOptionContext)(unsafe.Pointer(c.Native()))

	var _cret *C.gchar // in

	_cret = C.g_option_context_get_summary(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// SetDescription adds a string to be displayed in `--help` output after the
// list of options. This text often includes a bug reporting address.
//
// Note that the summary is translated (see
// g_option_context_set_translate_func()).
func (c *OptionContext) SetDescription(description string) {
	var _arg0 *C.GOptionContext // out
	var _arg1 *C.gchar          // out

	_arg0 = (*C.GOptionContext)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.gchar)(C.CString(description))
	defer C.free(unsafe.Pointer(_arg1))

	C.g_option_context_set_description(_arg0, _arg1)
}

// SetHelpEnabled enables or disables automatic generation of `--help` output.
// By default, g_option_context_parse() recognizes `--help`, `-h`, `-?`,
// `--help-all` and `--help-groupname` and creates suitable output to stdout.
func (c *OptionContext) SetHelpEnabled(helpEnabled bool) {
	var _arg0 *C.GOptionContext // out
	var _arg1 C.gboolean        // out

	_arg0 = (*C.GOptionContext)(unsafe.Pointer(c.Native()))
	if helpEnabled {
		_arg1 = C.TRUE
	}

	C.g_option_context_set_help_enabled(_arg0, _arg1)
}

// SetIgnoreUnknownOptions sets whether to ignore unknown options or not. If an
// argument is ignored, it is left in the @argv array after parsing. By default,
// g_option_context_parse() treats unknown options as error.
//
// This setting does not affect non-option arguments (i.e. arguments which don't
// start with a dash). But note that GOption cannot reliably determine whether a
// non-option belongs to a preceding unknown option.
func (c *OptionContext) SetIgnoreUnknownOptions(ignoreUnknown bool) {
	var _arg0 *C.GOptionContext // out
	var _arg1 C.gboolean        // out

	_arg0 = (*C.GOptionContext)(unsafe.Pointer(c.Native()))
	if ignoreUnknown {
		_arg1 = C.TRUE
	}

	C.g_option_context_set_ignore_unknown_options(_arg0, _arg1)
}

// SetMainGroup sets a Group as main group of the @context. This has the same
// effect as calling g_option_context_add_group(), the only difference is that
// the options in the main group are treated differently when generating
// `--help` output.
func (c *OptionContext) SetMainGroup(group *OptionGroup) {
	var _arg0 *C.GOptionContext // out
	var _arg1 *C.GOptionGroup   // out

	_arg0 = (*C.GOptionContext)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GOptionGroup)(unsafe.Pointer(group.Native()))

	C.g_option_context_set_main_group(_arg0, _arg1)
}

// SetStrictPosix sets strict POSIX mode.
//
// By default, this mode is disabled.
//
// In strict POSIX mode, the first non-argument parameter encountered (eg:
// filename) terminates argument processing. Remaining arguments are treated as
// non-options and are not attempted to be parsed.
//
// If strict POSIX mode is disabled then parsing is done in the GNU way where
// option arguments can be freely mixed with non-options.
//
// As an example, consider "ls foo -l". With GNU style parsing, this will list
// "foo" in long mode. In strict POSIX style, this will list the files named
// "foo" and "-l".
//
// It may be useful to force strict POSIX mode when creating "verb style"
// command line tools. For example, the "gsettings" command line tool supports
// the global option "--schemadir" as well as many subcommands ("get", "set",
// etc.) which each have their own set of arguments. Using strict POSIX mode
// will allow parsing the global options up to the verb name while leaving the
// remaining options to be parsed by the relevant subcommand (which can be
// determined by examining the verb name, which should be present in argv[1]
// after parsing).
func (c *OptionContext) SetStrictPosix(strictPosix bool) {
	var _arg0 *C.GOptionContext // out
	var _arg1 C.gboolean        // out

	_arg0 = (*C.GOptionContext)(unsafe.Pointer(c.Native()))
	if strictPosix {
		_arg1 = C.TRUE
	}

	C.g_option_context_set_strict_posix(_arg0, _arg1)
}

// SetSummary adds a string to be displayed in `--help` output before the list
// of options. This is typically a summary of the program functionality.
//
// Note that the summary is translated (see
// g_option_context_set_translate_func() and
// g_option_context_set_translation_domain()).
func (c *OptionContext) SetSummary(summary string) {
	var _arg0 *C.GOptionContext // out
	var _arg1 *C.gchar          // out

	_arg0 = (*C.GOptionContext)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.gchar)(C.CString(summary))
	defer C.free(unsafe.Pointer(_arg1))

	C.g_option_context_set_summary(_arg0, _arg1)
}

// SetTranslationDomain: a convenience function to use gettext() for translating
// user-visible strings.
func (c *OptionContext) SetTranslationDomain(domain string) {
	var _arg0 *C.GOptionContext // out
	var _arg1 *C.gchar          // out

	_arg0 = (*C.GOptionContext)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.gchar)(C.CString(domain))
	defer C.free(unsafe.Pointer(_arg1))

	C.g_option_context_set_translation_domain(_arg0, _arg1)
}

// OptionEntry: a GOptionEntry struct defines a single option. To have an
// effect, they must be added to a Group with
// g_option_context_add_main_entries() or g_option_group_add_entries().
type OptionEntry struct {
	native C.GOptionEntry
}

// WrapOptionEntry wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapOptionEntry(ptr unsafe.Pointer) *OptionEntry {
	if ptr == nil {
		return nil
	}

	return (*OptionEntry)(ptr)
}

// Native returns the underlying C source pointer.
func (o *OptionEntry) Native() unsafe.Pointer {
	return unsafe.Pointer(&o.native)
}

// LongName gets the field inside the struct.
func (o *OptionEntry) LongName() string {
	var v string // out
	v = C.GoString(o.native.long_name)
	return v
}

// ShortName gets the field inside the struct.
func (o *OptionEntry) ShortName() byte {
	var v byte // out
	v = (byte)(o.native.short_name)
	return v
}

// Flags gets the field inside the struct.
func (o *OptionEntry) Flags() int {
	var v int // out
	v = (int)(o.native.flags)
	return v
}

// Arg gets the field inside the struct.
func (o *OptionEntry) Arg() OptionArg {
	var v OptionArg // out
	v = OptionArg(o.native.arg)
	return v
}

// ArgData gets the field inside the struct.
func (o *OptionEntry) ArgData() interface{} {
	var v interface{} // out
	v = (interface{})(o.native.arg_data)
	return v
}

// Description gets the field inside the struct.
func (o *OptionEntry) Description() string {
	var v string // out
	v = C.GoString(o.native.description)
	return v
}

// ArgDescription gets the field inside the struct.
func (o *OptionEntry) ArgDescription() string {
	var v string // out
	v = C.GoString(o.native.arg_description)
	return v
}

// OptionGroup: a `GOptionGroup` struct defines the options in a single group.
// The struct has only private fields and should not be directly accessed.
//
// All options in a group share the same translation function. Libraries which
// need to parse commandline options are expected to provide a function for
// getting a `GOptionGroup` holding their options, which the application can
// then add to its Context.
type OptionGroup struct {
	native C.GOptionGroup
}

// WrapOptionGroup wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapOptionGroup(ptr unsafe.Pointer) *OptionGroup {
	if ptr == nil {
		return nil
	}

	return (*OptionGroup)(ptr)
}

func marshalOptionGroup(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapOptionGroup(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (o *OptionGroup) Native() unsafe.Pointer {
	return unsafe.Pointer(&o.native)
}

// AddEntries adds the options specified in @entries to @group.
func (g *OptionGroup) AddEntries(entries []OptionEntry) {
	var _arg0 *C.GOptionGroup // out
	var _arg1 *C.GOptionEntry

	_arg0 = (*C.GOptionGroup)(unsafe.Pointer(g.Native()))
	_arg1 = (*C.GOptionEntry)(C.malloc(C.ulong((len(entries) + 1)) * C.ulong(C.sizeof_GOptionEntry)))
	defer C.free(unsafe.Pointer(_arg1))

	{
		out := unsafe.Slice(_arg1, len(entries))
		for i := range entries {
			out[i] = (C.GOptionEntry)(unsafe.Pointer(entries[i].Native()))
		}
	}

	C.g_option_group_add_entries(_arg0, _arg1)
}

// Free frees a Group. Note that you must not free groups which have been added
// to a Context.
func (g *OptionGroup) Free() {
	var _arg0 *C.GOptionGroup // out

	_arg0 = (*C.GOptionGroup)(unsafe.Pointer(g.Native()))

	C.g_option_group_free(_arg0)
}

// Ref increments the reference count of @group by one.
func (g *OptionGroup) Ref() *OptionGroup {
	var _arg0 *C.GOptionGroup // out

	_arg0 = (*C.GOptionGroup)(unsafe.Pointer(g.Native()))

	var _cret *C.GOptionGroup // in

	_cret = C.g_option_group_ref(_arg0)

	var _optionGroup *OptionGroup // out

	_optionGroup = WrapOptionGroup(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_optionGroup, func(v *OptionGroup) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return _optionGroup
}

// SetTranslationDomain: a convenience function to use gettext() for translating
// user-visible strings.
func (g *OptionGroup) SetTranslationDomain(domain string) {
	var _arg0 *C.GOptionGroup // out
	var _arg1 *C.gchar        // out

	_arg0 = (*C.GOptionGroup)(unsafe.Pointer(g.Native()))
	_arg1 = (*C.gchar)(C.CString(domain))
	defer C.free(unsafe.Pointer(_arg1))

	C.g_option_group_set_translation_domain(_arg0, _arg1)
}

// Unref decrements the reference count of @group by one. If the reference count
// drops to 0, the @group will be freed. and all memory allocated by the @group
// is released.
func (g *OptionGroup) Unref() {
	var _arg0 *C.GOptionGroup // out

	_arg0 = (*C.GOptionGroup)(unsafe.Pointer(g.Native()))

	C.g_option_group_unref(_arg0)
}
