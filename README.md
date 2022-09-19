# Linking `cgo` library with static Swift library

This library is a demo for linking a `cgo` compiled library with a static swift library compiled using swift package manager.

## Instructions 

1. Clone repo

```bash
$ git clone <repo>
$ cd example-go-ios
```

2. Build swift library

```bash
$ cd ios
$ swift build
```

3. Run `go` module
```bash
# From project root
$ LD_DEAD_STRIP=1 go run cmd/libfoo/main.go
HELLO FROM A C LIBRARY FUNCTION
hello from a c library function
```

The go library passes a `C` string (`"HELLO FROM A C LIBRARY FUNCTION
hello from a c library function"`) to the `ios` library which then converts it into lower case.

## Changes from linking dynamic library

1. **Set library search paths to swift library folders**

```
#cgo -L/Applications/Xcode.app/Contents/Developer/Toolchains/XcodeDefault.xctoolchain/usr/lib/swift/macosx/ ... -Wl,-rpath,/Applications/Xcode.app/Contents/Developer/Toolchains/XcodeDefault.xctoolchain/usr/lib/swift/macosx/
```

This is repeated for all `swift` library folders in the `/Applications/Xcode.app/Contents/Developer/Toolchains/XcodeDefault.xctoolchain/usr/lib/` directory (`swift`, `swift-5.0` and `swift-5.5`):

This fixes all erros when linking with the messages:

```
ld: warning: Could not find or use auto-linked library 'swift_Concurrency'
ld: warning: Could not find or use auto-linked library 'swiftCompatibility51'
ld: warning: Could not find or use auto-linked library 'swiftSwiftOnoneSupport'
ld: warning: Could not find or use auto-linked library 'swiftCompatibilityConcurrency'
ld: warning: Could not find or use auto-linked library 'swiftCompatibility50'
ld: warning: Could not find or use auto-linked library 'swiftCompatibilityDynamicReplacements'
...
```

References:

https://forums.swift.org/t/could-not-find-or-use-auto-linked-library-swiftcompatibility50/54351

For linking to swift toolchain

2. **Set `LD_DEAD_STRIP=1` when using `go run`**

This fixes the error:

```
/usr/local/go/pkg/tool/darwin_amd64/link: running clang failed: exit status 1
Undefined symbols for architecture x86_64:
  "_swift_getFunctionReplacement", referenced from:
      _swift_getFunctionReplacement50 in libswiftCompatibilityDynamicReplacements.a(DynamicReplaceable.cpp.o)
     (maybe you meant: _swift_getFunctionReplacement50)
  "_swift_getOrigOfReplaceable", referenced from:
      _swift_getOrigOfReplaceable50 in libswiftCompatibilityDynamicReplacements.a(DynamicReplaceable.cpp.o)
     (maybe you meant: _swift_getOrigOfReplaceable50)
ld: symbol(s) not found for architecture x86_64
clang: error: linker command failed with exit code 1 (use -v to see invocation)
```

that appears after appying the fix in no. 1

References:

https://github.com/react-native-community/upgrade-support/issues/25

For fixing the issue in Xcode.

https://opensource.apple.com/source/cctools/cctools-622.5.1/RelNotes/CompilerTools.html?txt

For references to `LD_DEAD_STRIP=1` flag that is the `clang` equivalent to the Xcode `Dead Code Stripping` build setting.

3. Add `/usr/lib/swift/` to run time search paths in linker flags

```
#cgo LDFLAGS: ... -Wl,-rpath,/usr/lib/swift/
```

Fixes the error when running `go run`:

```
This copy of libswiftCore.dylib requires an OS version prior to 10.14.4.
```

that appears after adding no.1 and no.2 fix

References:

https://stackoverflow.com/questions/55361057/this-copy-of-libswiftcore-dylib-requires-an-os-version-prior-to-12-2-0

For adding `/user/lib/swift` to runtime search paths for linker if `iOS >= 12.2`

## Disclaimers and TODO

1. This uses `go run` instead of `go build` to make iteration quickly during development and save time.

2. Only `x86_64` architecture is used. Not tested on `arm64`

3. Built using Intel Macbook macOS BigSur 11.6.7 and Xcode version 13.2.1

4. go version go1.19.1 darwin/amd64

### TODO

1. Test using `go build` and link to an iOS project

2. Test for `arm64` architecture on real device

3. Test for `Edge-SDK` 

4. Clean `cgo` flag directives (specially `LDFLAGS`) by using environment variables and conditionally using correct architecture (`ios` instead of `macosx`)