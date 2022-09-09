// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// GType values.
var (
	GTypeLicense = coreglib.Type(C.gtk_license_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeLicense, F: marshalLicense},
	})
}

// License: type of license for an application.
//
// This enumeration can be expanded at later date.
type License C.gint

const (
	// LicenseUnknown: no license specified.
	LicenseUnknown License = iota
	// LicenseCustom: license text is going to be specified by the developer.
	LicenseCustom
	// LicenseGPL20: GNU General Public License, version 2.0 or later.
	LicenseGPL20
	// LicenseGPL30: GNU General Public License, version 3.0 or later.
	LicenseGPL30
	// LicenseLGPL21: GNU Lesser General Public License, version 2.1 or later.
	LicenseLGPL21
	// LicenseLGPL30: GNU Lesser General Public License, version 3.0 or later.
	LicenseLGPL30
	// LicenseBSD: BSD standard license.
	LicenseBSD
	// LicenseMITX11: MIT/X11 standard license.
	LicenseMITX11
	// LicenseArtistic: artistic License, version 2.0.
	LicenseArtistic
	// LicenseGPL20_Only: GNU General Public License, version 2.0 only. Since
	// 3.12.
	LicenseGPL20_Only
	// LicenseGPL30_Only: GNU General Public License, version 3.0 only. Since
	// 3.12.
	LicenseGPL30_Only
	// LicenseLGPL21_Only: GNU Lesser General Public License, version 2.1 only.
	// Since 3.12.
	LicenseLGPL21_Only
	// LicenseLGPL30_Only: GNU Lesser General Public License, version 3.0 only.
	// Since 3.12.
	LicenseLGPL30_Only
	// LicenseAGPL30: GNU Affero General Public License, version 3.0 or later.
	// Since: 3.22.
	LicenseAGPL30
	// LicenseAGPL30_Only: GNU Affero General Public License, version 3.0 only.
	// Since: 3.22.27.
	LicenseAGPL30_Only
	// LicenseBSD3: 3-clause BSD licence. Since: 3.24.20.
	LicenseBSD3
	// LicenseApache20: apache License, version 2.0. Since: 3.24.20.
	LicenseApache20
	// LicenseMPL20: mozilla Public License, version 2.0. Since: 3.24.20.
	LicenseMPL20
)

func marshalLicense(p uintptr) (interface{}, error) {
	return License(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for License.
func (l License) String() string {
	switch l {
	case LicenseUnknown:
		return "Unknown"
	case LicenseCustom:
		return "Custom"
	case LicenseGPL20:
		return "GPL20"
	case LicenseGPL30:
		return "GPL30"
	case LicenseLGPL21:
		return "LGPL21"
	case LicenseLGPL30:
		return "LGPL30"
	case LicenseBSD:
		return "BSD"
	case LicenseMITX11:
		return "MITX11"
	case LicenseArtistic:
		return "Artistic"
	case LicenseGPL20_Only:
		return "GPL20_Only"
	case LicenseGPL30_Only:
		return "GPL30_Only"
	case LicenseLGPL21_Only:
		return "LGPL21_Only"
	case LicenseLGPL30_Only:
		return "LGPL30_Only"
	case LicenseAGPL30:
		return "AGPL30"
	case LicenseAGPL30_Only:
		return "AGPL30_Only"
	case LicenseBSD3:
		return "BSD3"
	case LicenseApache20:
		return "Apache20"
	case LicenseMPL20:
		return "MPL20"
	default:
		return fmt.Sprintf("License(%d)", l)
	}
}
