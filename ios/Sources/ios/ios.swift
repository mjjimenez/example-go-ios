import Foundation

@_cdecl("lowercaseString") // export to C as `sayHello`
public func lowercaseString(strPtr: UnsafePointer<CChar>?) -> UnsafePointer<CChar> {
    // Creates a new string by copying the null-terminated UTF-8 data (C String)
    // referenced by the given pointer.
    let original = String(cString: strPtr!)
    let lowercase = original.lowercased()
    return UnsafePointer<CChar>(lowercase)!
}
