// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/box"
)

// #cgo pkg-config: glib-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <glib.h>
import "C"

// LogWriterOutput: return values from WriterFuncs to indicate whether the given
// log entry was successfully handled by the writer, or whether there was an
// error in handling it (and hence a fallback writer should be used).
//
// If a WriterFunc ignores a log entry, it should return G_LOG_WRITER_HANDLED.
type LogWriterOutput int

const (
	// Handled: log writer has handled the log entry.
	LogWriterOutputHandled LogWriterOutput = 1
	// Unhandled: log writer could not handle the log entry.
	LogWriterOutputUnhandled LogWriterOutput = 0
)

// LogLevelFlags flags specifying the level of log messages.
//
// It is possible to change how GLib treats messages of the various levels using
// g_log_set_handler() and g_log_set_fatal_mask().
type LogLevelFlags int

const (
	// LogLevelFlagsFlagRecursion: internal flag
	LogLevelFlagsFlagRecursion LogLevelFlags = 0b1
	// LogLevelFlagsFlagFatal: internal flag
	LogLevelFlagsFlagFatal LogLevelFlags = 0b10
	// LogLevelFlagsLevelError: log level for errors, see g_error(). This level
	// is also used for messages produced by g_assert().
	LogLevelFlagsLevelError LogLevelFlags = 0b100
	// LogLevelFlagsLevelCritical: log level for critical warning messages, see
	// g_critical(). This level is also used for messages produced by
	// g_return_if_fail() and g_return_val_if_fail().
	LogLevelFlagsLevelCritical LogLevelFlags = 0b1000
	// LogLevelFlagsLevelWarning: log level for warnings, see g_warning()
	LogLevelFlagsLevelWarning LogLevelFlags = 0b10000
	// LogLevelFlagsLevelMessage: log level for messages, see g_message()
	LogLevelFlagsLevelMessage LogLevelFlags = 0b100000
	// LogLevelFlagsLevelInfo: log level for informational messages, see
	// g_info()
	LogLevelFlagsLevelInfo LogLevelFlags = 0b1000000
	// LogLevelFlagsLevelDebug: log level for debug messages, see g_debug()
	LogLevelFlagsLevelDebug LogLevelFlags = 0b10000000
	// LogLevelFlagsLevelMask: mask including all log levels
	LogLevelFlagsLevelMask LogLevelFlags = -4
)

// LogFunc specifies the prototype of log handler functions.
//
// The default log handler, g_log_default_handler(), automatically appends a
// new-line character to @message when printing it. It is advised that any
// custom log handler functions behave similarly, so that logging calls in user
// code do not need modifying to add a new-line character to the message if the
// log handler is changed.
//
// This is not used if structured logging is enabled; see [Using Structured
// Logging][using-structured-logging].
type LogFunc func(logDomain string, logLevel LogLevelFlags, message string)

//export gotk4_LogFunc
func gotk4_LogFunc(arg0 *C.gchar, arg1 C.GLogLevelFlags, arg2 *C.gchar, arg3 C.gpointer) {
	v := box.Get(uintptr(arg3))
	if v == nil {
		panic(`callback not found`)
	}

	var logDomain string       // out
	var logLevel LogLevelFlags // out
	var message string         // out

	logDomain = C.GoString(arg0)
	logLevel = LogLevelFlags(arg1)
	message = C.GoString(arg2)

	fn := v.(LogFunc)
	fn(logDomain, logLevel, message)
}

func AssertWarning(logDomain string, file string, line int, prettyFunction string, expression string) {
	var _arg1 *C.char // out
	var _arg2 *C.char // out
	var _arg3 C.int   // out
	var _arg4 *C.char // out
	var _arg5 *C.char // out

	_arg1 = (*C.char)(C.CString(logDomain))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.char)(C.CString(file))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = C.int(line)
	_arg4 = (*C.char)(C.CString(prettyFunction))
	defer C.free(unsafe.Pointer(_arg4))
	_arg5 = (*C.char)(C.CString(expression))
	defer C.free(unsafe.Pointer(_arg5))

	C.g_assert_warning(_arg1, _arg2, _arg3, _arg4, _arg5)
}

// LogDefaultHandler: the default log handler set up by GLib;
// g_log_set_default_handler() allows to install an alternate default log
// handler. This is used if no log handler has been set for the particular log
// domain and log level combination. It outputs the message to stderr or stdout
// and if the log level is fatal it calls G_BREAKPOINT(). It automatically
// prints a new-line character after the message, so one does not need to be
// manually included in @message.
//
// The behavior of this log handler can be influenced by a number of environment
// variables:
//
// - `G_MESSAGES_PREFIXED`: A :-separated list of log levels for which messages
// should be prefixed by the program name and PID of the application.
//
// - `G_MESSAGES_DEBUG`: A space-separated list of log domains for which debug
// and informational messages are printed. By default these messages are not
// printed.
//
// stderr is used for levels G_LOG_LEVEL_ERROR, G_LOG_LEVEL_CRITICAL,
// G_LOG_LEVEL_WARNING and G_LOG_LEVEL_MESSAGE. stdout is used for the rest,
// unless stderr was requested by g_log_writer_default_set_use_stderr().
//
// This has no effect if structured logging is enabled; see [Using Structured
// Logging][using-structured-logging].
func LogDefaultHandler(logDomain string, logLevel LogLevelFlags, message string, unusedData interface{}) {
	var _arg1 *C.gchar         // out
	var _arg2 C.GLogLevelFlags // out
	var _arg3 *C.gchar         // out
	var _arg4 C.gpointer       // out

	_arg1 = (*C.gchar)(C.CString(logDomain))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.GLogLevelFlags(logLevel)
	_arg3 = (*C.gchar)(C.CString(message))
	defer C.free(unsafe.Pointer(_arg3))
	_arg4 = (C.gpointer)(box.Assign(unusedData))

	C.g_log_default_handler(_arg1, _arg2, _arg3, _arg4)
}

// LogRemoveHandler removes the log handler.
//
// This has no effect if structured logging is enabled; see [Using Structured
// Logging][using-structured-logging].
func LogRemoveHandler(logDomain string, handlerId uint) {
	var _arg1 *C.gchar // out
	var _arg2 C.guint  // out

	_arg1 = (*C.gchar)(C.CString(logDomain))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.guint(handlerId)

	C.g_log_remove_handler(_arg1, _arg2)
}

// LogSetAlwaysFatal sets the message levels which are always fatal, in any log
// domain. When a message with any of these levels is logged the program
// terminates. You can only set the levels defined by GLib to be fatal.
// G_LOG_LEVEL_ERROR is always fatal.
//
// You can also make some message levels fatal at runtime by setting the
// `G_DEBUG` environment variable (see Running GLib Applications
// (glib-running.html)).
//
// Libraries should not call this function, as it affects all messages logged by
// a process, including those from other libraries.
//
// Structured log messages (using g_log_structured() and
// g_log_structured_array()) are fatal only if the default log writer is used;
// otherwise it is up to the writer function to determine which log messages are
// fatal. See [Using Structured Logging][using-structured-logging].
func LogSetAlwaysFatal(fatalMask LogLevelFlags) LogLevelFlags {
	var _arg1 C.GLogLevelFlags // out
	var _cret C.GLogLevelFlags // in

	_arg1 = C.GLogLevelFlags(fatalMask)

	_cret = C.g_log_set_always_fatal(_arg1)

	var _logLevelFlags LogLevelFlags // out

	_logLevelFlags = LogLevelFlags(_cret)

	return _logLevelFlags
}

// LogSetFatalMask sets the log levels which are fatal in the given domain.
// G_LOG_LEVEL_ERROR is always fatal.
//
// This has no effect on structured log messages (using g_log_structured() or
// g_log_structured_array()). To change the fatal behaviour for specific log
// messages, programs must install a custom log writer function using
// g_log_set_writer_func(). See [Using Structured
// Logging][using-structured-logging].
//
// This function is mostly intended to be used with G_LOG_LEVEL_CRITICAL. You
// should typically not set G_LOG_LEVEL_WARNING, G_LOG_LEVEL_MESSAGE,
// G_LOG_LEVEL_INFO or G_LOG_LEVEL_DEBUG as fatal except inside of test
// programs.
func LogSetFatalMask(logDomain string, fatalMask LogLevelFlags) LogLevelFlags {
	var _arg1 *C.gchar         // out
	var _arg2 C.GLogLevelFlags // out
	var _cret C.GLogLevelFlags // in

	_arg1 = (*C.gchar)(C.CString(logDomain))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.GLogLevelFlags(fatalMask)

	_cret = C.g_log_set_fatal_mask(_arg1, _arg2)

	var _logLevelFlags LogLevelFlags // out

	_logLevelFlags = LogLevelFlags(_cret)

	return _logLevelFlags
}

// LogVariant: log a message with structured data, accepting the data within a
// #GVariant. This version is especially useful for use in other languages, via
// introspection.
//
// The only mandatory item in the @fields dictionary is the "MESSAGE" which must
// contain the text shown to the user.
//
// The values in the @fields dictionary are likely to be of type String
// (VARIANT_TYPE_STRING). Array of bytes (VARIANT_TYPE_BYTESTRING) is also
// supported. In this case the message is handled as binary and will be
// forwarded to the log writer as such. The size of the array should not be
// higher than G_MAXSSIZE. Otherwise it will be truncated to this size. For
// other types g_variant_print() will be used to convert the value into a
// string.
//
// For more details on its usage and about the parameters, see
// g_log_structured().
func LogVariant(logDomain string, logLevel LogLevelFlags, fields *Variant) {
	var _arg1 *C.gchar         // out
	var _arg2 C.GLogLevelFlags // out
	var _arg3 *C.GVariant      // out

	_arg1 = (*C.gchar)(C.CString(logDomain))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.GLogLevelFlags(logLevel)
	_arg3 = (*C.GVariant)(unsafe.Pointer(fields))

	C.g_log_variant(_arg1, _arg2, _arg3)
}

// LogWriterDefaultSetUseStderr: configure whether the built-in log functions
// (g_log_default_handler() for the old-style API, and both
// g_log_writer_default() and g_log_writer_standard_streams() for the structured
// API) will output all log messages to `stderr`.
//
// By default, log messages of levels G_LOG_LEVEL_INFO and G_LOG_LEVEL_DEBUG are
// sent to `stdout`, and other log messages are sent to `stderr`. This is
// problematic for applications that intend to reserve `stdout` for structured
// output such as JSON or XML.
//
// This function sets global state. It is not thread-aware, and should be called
// at the very start of a program, before creating any other threads or creating
// objects that could create worker threads of their own.
func LogWriterDefaultSetUseStderr(useStderr bool) {
	var _arg1 C.gboolean // out

	if useStderr {
		_arg1 = C.TRUE
	}

	C.g_log_writer_default_set_use_stderr(_arg1)
}

// LogWriterDefaultWouldDrop: check whether g_log_writer_default() and
// g_log_default_handler() would ignore a message with the given domain and
// level.
//
// As with g_log_default_handler(), this function drops debug and informational
// messages unless their log domain (or `all`) is listed in the space-separated
// `G_MESSAGES_DEBUG` environment variable.
//
// This can be used when implementing log writers with the same filtering
// behaviour as the default, but a different destination or output format:
//
//      if (!g_log_writer_default_would_drop (G_LOG_LEVEL_DEBUG, G_LOG_DOMAIN))
//        {
//          gchar *result = expensive_computation (my_object);
//
//          g_debug ("my_object result: s", result);
//          g_free (result);
//        }
func LogWriterDefaultWouldDrop(logLevel LogLevelFlags, logDomain string) bool {
	var _arg1 C.GLogLevelFlags // out
	var _arg2 *C.char          // out
	var _cret C.gboolean       // in

	_arg1 = C.GLogLevelFlags(logLevel)
	_arg2 = (*C.char)(C.CString(logDomain))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.g_log_writer_default_would_drop(_arg1, _arg2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// LogWriterIsJournald: check whether the given @output_fd file descriptor is a
// connection to the systemd journal, or something else (like a log file or
// `stdout` or `stderr`).
//
// Invalid file descriptors are accepted and return false, which allows for the
// following construct without needing any additional error handling:
//
//    is_journald = g_log_writer_is_journald (fileno (stderr));
func LogWriterIsJournald(outputFd int) bool {
	var _arg1 C.gint     // out
	var _cret C.gboolean // in

	_arg1 = C.gint(outputFd)

	_cret = C.g_log_writer_is_journald(_arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// LogWriterSupportsColor: check whether the given @output_fd file descriptor
// supports ANSI color escape sequences. If so, they can safely be used when
// formatting log messages.
func LogWriterSupportsColor(outputFd int) bool {
	var _arg1 C.gint     // out
	var _cret C.gboolean // in

	_arg1 = C.gint(outputFd)

	_cret = C.g_log_writer_supports_color(_arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// LogField: structure representing a single field in a structured log entry.
// See g_log_structured() for details.
//
// Log fields may contain arbitrary values, including binary with embedded nul
// bytes. If the field contains a string, the string must be UTF-8 encoded and
// have a trailing nul byte. Otherwise, @length must be set to a non-negative
// value.
type LogField struct {
	native C.GLogField
}

// WrapLogField wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapLogField(ptr unsafe.Pointer) *LogField {
	return (*LogField)(ptr)
}

// Native returns the underlying C source pointer.
func (l *LogField) Native() unsafe.Pointer {
	return unsafe.Pointer(&l.native)
}

// Key: field name (UTF-8 string)
func (l *LogField) Key() string {
	var v string // out
	v = C.GoString(l.native.key)
	return v
}

// Value: field value (arbitrary bytes)
func (l *LogField) Value() interface{} {
	var v interface{} // out
	v = box.Get(uintptr(l.native.value))
	return v
}

// Length: length of @value, in bytes, or -1 if it is nul-terminated
func (l *LogField) Length() int {
	var v int // out
	v = int(l.native.length)
	return v
}
