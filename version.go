package libversion

/*
#cgo LDFLAGS: -ldl
#include <dlfcn.h>

typedef int (*version_compare2_fn)(char* a, char* b);

int invoke_fn_2_args(void* fn, char* a, char* b) {
	return ((version_compare2_fn) fn)(a, b);
}

typedef int (*version_compare4_fn)(char* a, char* b, int v1_flags, int v2_flags);

int invoke_fn_4_args(void* fn, char* a, char* b, int v1_flags, int v2_flags) {
	return ((version_compare4_fn) fn)(a, b, v1_flags, v2_flags);
}
*/
import "C"
import (
	"fmt"
	"os"
)

var Compare func(lhsVersion, rhsVersion string) int
var CompareWithFlags func(lhsVersion, rhsVersion string, lhsFlags, rhsFlags int) int

const (
	FLAG_P_IS_PATCH   = 0x1
	FLAG_ANY_IS_PATCH = 0x2
	FLAG_LOWER_BOUND  = 0x4
	FLAG_UPPER_BOUND  = 0x8
)

var libPaths = []string{
	".",
	"./lib",
	"/usr/local/lib",
	"/usr/lib",
	"/lib",
	"/usr/local/lib64",
	"/usr/lib64",
	"/lib64",
	"~/.local/lib",
	"~/.local/lib64",
}

func getLibVersionPath() (string, error) {
	const LIB_VERSION_FILE_NAME = "libversion.so"

	for _, path := range libPaths {
		if _, err := os.Stat(path + "/" + LIB_VERSION_FILE_NAME); err == nil {
			return path + "/" + LIB_VERSION_FILE_NAME, nil
		}
	}

	return "", fmt.Errorf("libversion.so not found")
}

func init() {
	libVersionPath, err := getLibVersionPath()
	if err != nil {
		panic(err)
	}

	handle := C.dlopen(C.CString(libVersionPath), C.RTLD_LAZY)
	if handle == nil {
		panic("dlopen failed: " + libVersionPath)
	}

	versionCompare2Handle := C.dlsym(handle, C.CString("version_compare2"))
	if versionCompare2Handle == nil {
		panic("dlsym failed: version_compare2")
	}

	Compare = func(lhsVersion, rhsVersion string) int {
		return int(C.invoke_fn_2_args(versionCompare2Handle, C.CString(lhsVersion), C.CString(rhsVersion)))
	}

	versionCompare4Handle := C.dlsym(handle, C.CString("version_compare4"))
	if versionCompare4Handle == nil {
		panic("dlsym failed: version_compare4")
	}

	CompareWithFlags = func(lhsVersion, rhsVersion string, lhsFlags, rhsFlags int) int {
		return int(C.invoke_fn_4_args(versionCompare4Handle, C.CString(lhsVersion), C.CString(rhsVersion), C.int(lhsFlags), C.int(rhsFlags)))
	}
}
