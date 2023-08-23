// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [FragmentedMovieMinder] class.
var FragmentedMovieMinderClass = _FragmentedMovieMinderClass{objc.GetClass("AVFragmentedMovieMinder")}

type _FragmentedMovieMinderClass struct {
	objc.Class
}

// An interface definition for the [FragmentedMovieMinder] class.
type IFragmentedMovieMinder interface {
	IFragmentedAssetMinder
	RemoveFragmentedMovie(movie IFragmentedMovie)
	AddFragmentedMovie(movie IFragmentedMovie)
	Movies() []FragmentedMovie
}

// An object that checks whether a fragmented movie appends additional movie fragments. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avfragmentedmovieminder?language=objc
type FragmentedMovieMinder struct {
	FragmentedAssetMinder
}

func FragmentedMovieMinderFrom(ptr unsafe.Pointer) FragmentedMovieMinder {
	return FragmentedMovieMinder{
		FragmentedAssetMinder: FragmentedAssetMinderFrom(ptr),
	}
}

func (f_ FragmentedMovieMinder) InitWithMovieMindingInterval(movie IFragmentedMovie, mindingInterval foundation.TimeInterval) FragmentedMovieMinder {
	rv := objc.Call[FragmentedMovieMinder](f_, objc.Sel("initWithMovie:mindingInterval:"), objc.Ptr(movie), mindingInterval)
	return rv
}

// Creates a movie minder and adds a movie with a minding interval. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avfragmentedmovieminder/1390294-initwithmovie?language=objc
func NewFragmentedMovieMinderWithMovieMindingInterval(movie IFragmentedMovie, mindingInterval foundation.TimeInterval) FragmentedMovieMinder {
	instance := FragmentedMovieMinderClass.Alloc().InitWithMovieMindingInterval(movie, mindingInterval)
	instance.Autorelease()
	return instance
}

func (fc _FragmentedMovieMinderClass) FragmentedMovieMinderWithMovieMindingInterval(movie IFragmentedMovie, mindingInterval foundation.TimeInterval) FragmentedMovieMinder {
	rv := objc.Call[FragmentedMovieMinder](fc, objc.Sel("fragmentedMovieMinderWithMovie:mindingInterval:"), objc.Ptr(movie), mindingInterval)
	return rv
}

// Creates a movie minder and adds a movie with a minding interval. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avfragmentedmovieminder/1458262-fragmentedmovieminderwithmovie?language=objc
func FragmentedMovieMinder_FragmentedMovieMinderWithMovieMindingInterval(movie IFragmentedMovie, mindingInterval foundation.TimeInterval) FragmentedMovieMinder {
	return FragmentedMovieMinderClass.FragmentedMovieMinderWithMovieMindingInterval(movie, mindingInterval)
}

func (fc _FragmentedMovieMinderClass) Alloc() FragmentedMovieMinder {
	rv := objc.Call[FragmentedMovieMinder](fc, objc.Sel("alloc"))
	return rv
}

func FragmentedMovieMinder_Alloc() FragmentedMovieMinder {
	return FragmentedMovieMinderClass.Alloc()
}

func (fc _FragmentedMovieMinderClass) New() FragmentedMovieMinder {
	rv := objc.Call[FragmentedMovieMinder](fc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewFragmentedMovieMinder() FragmentedMovieMinder {
	return FragmentedMovieMinderClass.New()
}

func (f_ FragmentedMovieMinder) Init() FragmentedMovieMinder {
	rv := objc.Call[FragmentedMovieMinder](f_, objc.Sel("init"))
	return rv
}

func (fc _FragmentedMovieMinderClass) FragmentedAssetMinderWithAssetMindingInterval(asset IAsset, mindingInterval foundation.TimeInterval) FragmentedMovieMinder {
	rv := objc.Call[FragmentedMovieMinder](fc, objc.Sel("fragmentedAssetMinderWithAsset:mindingInterval:"), objc.Ptr(asset), mindingInterval)
	return rv
}

// Creates a fragmented asset minder containing the specified asset and minding interval. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avfragmentedassetminder/1387182-fragmentedassetminderwithasset?language=objc
func FragmentedMovieMinder_FragmentedAssetMinderWithAssetMindingInterval(asset IAsset, mindingInterval foundation.TimeInterval) FragmentedMovieMinder {
	return FragmentedMovieMinderClass.FragmentedAssetMinderWithAssetMindingInterval(asset, mindingInterval)
}

func (f_ FragmentedMovieMinder) InitWithAssetMindingInterval(asset IAsset, mindingInterval foundation.TimeInterval) FragmentedMovieMinder {
	rv := objc.Call[FragmentedMovieMinder](f_, objc.Sel("initWithAsset:mindingInterval:"), objc.Ptr(asset), mindingInterval)
	return rv
}

// Creates a fragmented asset minder that monitors the specified asset at the indicated minding interval. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avfragmentedassetminder/2966508-initwithasset?language=objc
func NewFragmentedMovieMinderWithAssetMindingInterval(asset IAsset, mindingInterval foundation.TimeInterval) FragmentedMovieMinder {
	instance := FragmentedMovieMinderClass.Alloc().InitWithAssetMindingInterval(asset, mindingInterval)
	instance.Autorelease()
	return instance
}

// Removes a fragmented movie from the array of movies being minded. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avfragmentedmovieminder/1389794-removefragmentedmovie?language=objc
func (f_ FragmentedMovieMinder) RemoveFragmentedMovie(movie IFragmentedMovie) {
	objc.Call[objc.Void](f_, objc.Sel("removeFragmentedMovie:"), objc.Ptr(movie))
}

// Adds a fragmented movie to the array of movies being minded. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avfragmentedmovieminder/1386171-addfragmentedmovie?language=objc
func (f_ FragmentedMovieMinder) AddFragmentedMovie(movie IFragmentedMovie) {
	objc.Call[objc.Void](f_, objc.Sel("addFragmentedMovie:"), objc.Ptr(movie))
}

// An array containing the fragmented movie objects being minded. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avfragmentedmovieminder/1388707-movies?language=objc
func (f_ FragmentedMovieMinder) Movies() []FragmentedMovie {
	rv := objc.Call[[]FragmentedMovie](f_, objc.Sel("movies"))
	return rv
}
