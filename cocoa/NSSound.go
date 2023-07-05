package cocoa

import (
	"github.com/progrium/macdriver/core"
)

type NSSound struct {
	gen_NSSound
}

func NSSound_InitWithData(data core.NSDataRef) NSSound {
	return NSSound_alloc().InitWithData_asNSSound(data)
}

func (sound NSSound) Play() {
	sound.Play()
}

func (sound NSSound) Pause() {
	sound.Pause()
}

func (sound NSSound) Resume() {
	sound.Resume()
}
