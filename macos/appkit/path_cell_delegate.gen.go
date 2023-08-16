// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/objc"
)

// A set of methods that enable the delegate of a path cell object to customize the Open panel or pop-up menu of a path whose style is set to [appkit/nspathstyle/nspathstylepopup]. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcelldelegate?language=objc
type PPathCellDelegate interface {
	// optional
	PathCellWillPopUpMenu(pathCell PathCell, menu Menu)
	HasPathCellWillPopUpMenu() bool
}

// A delegate implementation builder for the [PPathCellDelegate] protocol.
type PathCellDelegate struct {
	_PathCellWillPopUpMenu func(pathCell PathCell, menu Menu)
}

func (di *PathCellDelegate) HasPathCellWillPopUpMenu() bool {
	return di._PathCellWillPopUpMenu != nil
}

// Implement this method to customize the menu of a pop-up–style path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcelldelegate/1525005-pathcell?language=objc
func (di *PathCellDelegate) SetPathCellWillPopUpMenu(f func(pathCell PathCell, menu Menu)) {
	di._PathCellWillPopUpMenu = f
}

// Implement this method to customize the menu of a pop-up–style path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcelldelegate/1525005-pathcell?language=objc
func (di *PathCellDelegate) PathCellWillPopUpMenu(pathCell PathCell, menu Menu) {
	di._PathCellWillPopUpMenu(pathCell, menu)
}

// A concrete type wrapper for the [PPathCellDelegate] protocol.
type PathCellDelegateWrapper struct {
	objc.Object
}

func (p_ PathCellDelegateWrapper) HasPathCellWillPopUpMenu() bool {
	return p_.RespondsToSelector(objc.Sel("pathCell:willPopUpMenu:"))
}

// Implement this method to customize the menu of a pop-up–style path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcelldelegate/1525005-pathcell?language=objc
func (p_ PathCellDelegateWrapper) PathCellWillPopUpMenu(pathCell IPathCell, menu IMenu) {
	objc.Call[objc.Void](p_, objc.Sel("pathCell:willPopUpMenu:"), objc.Ptr(pathCell), objc.Ptr(menu))
}
