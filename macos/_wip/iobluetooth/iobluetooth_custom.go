// WIP
//
// TODO:
// * Issue where there is both OBEXSession and IOBluetoothOBEXSession, which inherits from OBEXSession, except the prefix is trimmed so both cannot exist
package iobluetooth

import "github.com/progrium/darwinkit/objc"

// this guy was generated with Autorelease on OBEXError (an int) because of starting with Copy... todo: fix?

// Copy a remote file to a local path [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/iobluetooth/obexfiletransferservices/1434277-copyremotefile?language=objc
func (o_ OBEXFileTransferServices) CopyRemoteFileToLocalPath(inRemoteFileName string, inLocalPathAndName string) OBEXError {
	rv := objc.Call[OBEXError](o_, objc.Sel("copyRemoteFile:toLocalPath:"), inRemoteFileName, inLocalPathAndName)
	return rv
}
