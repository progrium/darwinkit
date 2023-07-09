package cocoa

import (
	"testing"

	"github.com/progrium/macdriver/core"
	"github.com/stretchr/testify/assert"
)

func TestNSImageSize(t *testing.T) {
	img := NSImage_Alloc().InitWithSize(core.Size(100, 200))
	size := img.Size()
	assert.EqualValues(t, 100, size.Width)
	assert.EqualValues(t, 200, size.Height)
}
