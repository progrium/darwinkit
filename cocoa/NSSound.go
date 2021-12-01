package cocoa

import (
	"github.com/progrium/macdriver/core"
)

type NSSound struct {
	gen_NSSound
}

func NSSound_InitWithData(data core.NSDataRef) NSSound {
	return NSSound_alloc().InitWithData__asNSSound(data)
}

func (sound NSSound) NSSound_Play() {
	sound.Play()
}
