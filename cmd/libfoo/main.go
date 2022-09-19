package main

/*
#cgo CFLAGS: -I./include
#cgo LDFLAGS: -L${SRCDIR}/../../ios/.build/debug/ -lios -L/Applications/Xcode.app/Contents/Developer/Toolchains/XcodeDefault.xctoolchain/usr/lib/swift/macosx/ -L/Applications/Xcode.app/Contents/Developer/Toolchains/XcodeDefault.xctoolchain/usr/lib/swift-5.0/macosx/ -L/Applications/Xcode.app/Contents/Developer/Toolchains/XcodeDefault.xctoolchain/usr/lib/swift-5.5/macosx/ -Wl,-rpath,/usr/lib/swift/ -Wl,-rpath,${SRCDIR}/../../ios/.build/debug/ -Wl,-rpath,/Applications/Xcode.app/Contents/Developer/Toolchains/XcodeDefault.xctoolchain/usr/lib/swift/macosx/ -Wl,-rpath,/Applications/Xcode.app/Contents/Developer/Toolchains/XcodeDefault.xctoolchain/usr/lib/swift-5.0/macosx/ -Wl,-rpath,/Applications/Xcode.app/Contents/Developer/Toolchains/XcodeDefault.xctoolchain/usr/lib/swift-5.5/macosx/
#include <stdlib.h>
#include "ios.h"
*/
import "C"
import (
  "fmt"
  "unsafe"
)

func main() { 
  fmt.Println("HELLO FROM A C LIBRARY FUNCTION")

  mystr := C.CString("HELLO FROM A C LIBRARY FUNCTION")
	lowercaseStr := C.GoString(C.lowercaseString(mystr))
  fmt.Println(lowercaseStr)
	C.free(unsafe.Pointer(mystr))

}
