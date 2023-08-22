// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [FragmentedMovieTrack] class.
var FragmentedMovieTrackClass = _FragmentedMovieTrackClass{objc.GetClass("AVFragmentedMovieTrack")}

type _FragmentedMovieTrackClass struct {
	objc.Class
}

// An interface definition for the [FragmentedMovieTrack] class.
type IFragmentedMovieTrack interface {
	IMovieTrack
}

// An object that represents a track in a fragmented movie. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avfragmentedmovietrack?language=objc
type FragmentedMovieTrack struct {
	MovieTrack
}

func FragmentedMovieTrackFrom(ptr unsafe.Pointer) FragmentedMovieTrack {
	return FragmentedMovieTrack{
		MovieTrack: MovieTrackFrom(ptr),
	}
}

func (fc _FragmentedMovieTrackClass) Alloc() FragmentedMovieTrack {
	rv := objc.Call[FragmentedMovieTrack](fc, objc.Sel("alloc"))
	return rv
}

func FragmentedMovieTrack_Alloc() FragmentedMovieTrack {
	return FragmentedMovieTrackClass.Alloc()
}

func (fc _FragmentedMovieTrackClass) New() FragmentedMovieTrack {
	rv := objc.Call[FragmentedMovieTrack](fc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewFragmentedMovieTrack() FragmentedMovieTrack {
	return FragmentedMovieTrackClass.New()
}

func (f_ FragmentedMovieTrack) Init() FragmentedMovieTrack {
	rv := objc.Call[FragmentedMovieTrack](f_, objc.Sel("init"))
	return rv
}
