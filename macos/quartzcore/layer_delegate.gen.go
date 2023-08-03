// auto-generated code, do not modify
package quartzcore

import (
	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

type ILayerDelegate interface {
	ImplementsDisplayLayer() bool
	// optional
	DisplayLayer(layer Layer)
	ImplementsDrawLayerInContext() bool
	// optional
	DrawLayerInContext(layer Layer, ctx coregraphics.ContextRef)
	ImplementsLayerWillDraw() bool
	// optional
	LayerWillDraw(layer Layer)
	ImplementsLayoutSublayersOfLayer() bool
	// optional
	LayoutSublayersOfLayer(layer Layer)
	ImplementsActionForLayerForKey() bool
	// optional
	ActionForLayerForKey(layer Layer, event string) IAction
}

type LayerDelegate struct {
	_DisplayLayer           func(layer Layer)
	_DrawLayerInContext     func(layer Layer, ctx coregraphics.ContextRef)
	_LayerWillDraw          func(layer Layer)
	_LayoutSublayersOfLayer func(layer Layer)
	_ActionForLayerForKey   func(layer Layer, event string) IAction
}

func (di *LayerDelegate) ImplementsDisplayLayer() bool {
	return di._DisplayLayer != nil
}

func (di *LayerDelegate) SetDisplayLayer(f func(layer Layer)) {
	di._DisplayLayer = f
}

func (di *LayerDelegate) DisplayLayer(layer Layer) {
	di._DisplayLayer(layer)
}
func (di *LayerDelegate) ImplementsDrawLayerInContext() bool {
	return di._DrawLayerInContext != nil
}

func (di *LayerDelegate) SetDrawLayerInContext(f func(layer Layer, ctx coregraphics.ContextRef)) {
	di._DrawLayerInContext = f
}

func (di *LayerDelegate) DrawLayerInContext(layer Layer, ctx coregraphics.ContextRef) {
	di._DrawLayerInContext(layer, ctx)
}
func (di *LayerDelegate) ImplementsLayerWillDraw() bool {
	return di._LayerWillDraw != nil
}

func (di *LayerDelegate) SetLayerWillDraw(f func(layer Layer)) {
	di._LayerWillDraw = f
}

func (di *LayerDelegate) LayerWillDraw(layer Layer) {
	di._LayerWillDraw(layer)
}
func (di *LayerDelegate) ImplementsLayoutSublayersOfLayer() bool {
	return di._LayoutSublayersOfLayer != nil
}

func (di *LayerDelegate) SetLayoutSublayersOfLayer(f func(layer Layer)) {
	di._LayoutSublayersOfLayer = f
}

func (di *LayerDelegate) LayoutSublayersOfLayer(layer Layer) {
	di._LayoutSublayersOfLayer(layer)
}
func (di *LayerDelegate) ImplementsActionForLayerForKey() bool {
	return di._ActionForLayerForKey != nil
}

func (di *LayerDelegate) SetActionForLayerForKey(f func(layer Layer, event string) IAction) {
	di._ActionForLayerForKey = f
}

func (di *LayerDelegate) ActionForLayerForKey(layer Layer, event string) IAction {
	return di._ActionForLayerForKey(layer, event)
}

type LayerDelegateWrapper struct {
	objc.Object
}

func (l_ LayerDelegateWrapper) ImplementsDisplayLayer() bool {
	return l_.RespondsToSelector(objc.GetSelector("displayLayer:"))
}

func (l_ LayerDelegateWrapper) DisplayLayer(layer ILayer) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("displayLayer:"), objc.ExtractPtr(layer))
}

func (l_ LayerDelegateWrapper) ImplementsDrawLayerInContext() bool {
	return l_.RespondsToSelector(objc.GetSelector("drawLayer:inContext:"))
}

func (l_ LayerDelegateWrapper) DrawLayerInContext(layer ILayer, ctx coregraphics.ContextRef) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("drawLayer:inContext:"), objc.ExtractPtr(layer), ctx)
}

func (l_ LayerDelegateWrapper) ImplementsLayerWillDraw() bool {
	return l_.RespondsToSelector(objc.GetSelector("layerWillDraw:"))
}

func (l_ LayerDelegateWrapper) LayerWillDraw(layer ILayer) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("layerWillDraw:"), objc.ExtractPtr(layer))
}

func (l_ LayerDelegateWrapper) ImplementsLayoutSublayersOfLayer() bool {
	return l_.RespondsToSelector(objc.GetSelector("layoutSublayersOfLayer:"))
}

func (l_ LayerDelegateWrapper) LayoutSublayersOfLayer(layer ILayer) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("layoutSublayersOfLayer:"), objc.ExtractPtr(layer))
}

func (l_ LayerDelegateWrapper) ImplementsActionForLayerForKey() bool {
	return l_.RespondsToSelector(objc.GetSelector("actionForLayer:forKey:"))
}

func (l_ LayerDelegateWrapper) ActionForLayerForKey(layer ILayer, event string) ActionWrapper {
	rv := objc.CallMethod[ActionWrapper](l_, objc.GetSelector("actionForLayer:forKey:"), objc.ExtractPtr(layer), event)
	return rv
}
