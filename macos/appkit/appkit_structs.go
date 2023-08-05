package appkit

import (
	"github.com/progrium/macdriver/macos/coregraphics"
)

// stuct define should be synced with <AppKit/NSCollectionViewCompositionalLayout.h>

type DirectionalEdgeInsets struct {
	Top      coregraphics.Float
	Leading  coregraphics.Float
	Bottom   coregraphics.Float
	Trailing coregraphics.Float
}
