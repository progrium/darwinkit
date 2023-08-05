// AUTO-GENERATED CODE, DO NOT MODIFY
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/macos/quartzcore"
	"github.com/progrium/macdriver/objc"
)

var AnimationContextClass = _AnimationContextClass{objc.GetClass("NSAnimationContext")}

type _AnimationContextClass struct {
	objc.Class
}

type IAnimationContext interface {
	objc.IObject
	CompletionHandler() func()
	SetCompletionHandler(value func())
	Duration() foundation.TimeInterval
	SetDuration(value foundation.TimeInterval)
	TimingFunction() quartzcore.MediaTimingFunction
	SetTimingFunction(value quartzcore.IMediaTimingFunction)
	AllowsImplicitAnimation() bool
	SetAllowsImplicitAnimation(value bool)
}

type AnimationContext struct {
	objc.Object
}

func MakeAnimationContext(ptr unsafe.Pointer) AnimationContext {
	return AnimationContext{
		Object: objc.MakeObject(ptr),
	}
}

func (ac _AnimationContextClass) Alloc() AnimationContext {
	rv := objc.CallMethod[AnimationContext](ac, objc.GetSelector("alloc"))
	return rv
}

func AnimationContext_Alloc() AnimationContext {
	return AnimationContextClass.Alloc()
}

func (ac _AnimationContextClass) New() AnimationContext {
	rv := objc.CallMethod[AnimationContext](ac, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewAnimationContext() AnimationContext {
	return AnimationContextClass.New()
}

func AnimationContext_New() AnimationContext {
	return AnimationContextClass.New()
}

func (a_ AnimationContext) Init() AnimationContext {
	rv := objc.CallMethod[AnimationContext](a_, objc.GetSelector("init"))
	return rv
}

func AnimationContext_Init() AnimationContext {
	return AnimationContextClass.Alloc().Init()
}

func (ac _AnimationContextClass) BeginGrouping() {
	objc.CallMethod[objc.Void](ac, objc.GetSelector("beginGrouping"))
}

func AnimationContext_BeginGrouping() {
	AnimationContextClass.BeginGrouping()
}

func (ac _AnimationContextClass) EndGrouping() {
	objc.CallMethod[objc.Void](ac, objc.GetSelector("endGrouping"))
}

func AnimationContext_EndGrouping() {
	AnimationContextClass.EndGrouping()
}

func (ac _AnimationContextClass) RunAnimationGroupCompletionHandler(changes func(context AnimationContext), completionHandler func()) {
	objc.CallMethod[objc.Void](ac, objc.GetSelector("runAnimationGroup:completionHandler:"), changes, completionHandler)
}

func AnimationContext_RunAnimationGroupCompletionHandler(changes func(context AnimationContext), completionHandler func()) {
	AnimationContextClass.RunAnimationGroupCompletionHandler(changes, completionHandler)
}

func (ac _AnimationContextClass) RunAnimationGroup(changes func(context AnimationContext)) {
	objc.CallMethod[objc.Void](ac, objc.GetSelector("runAnimationGroup:"), changes)
}

func AnimationContext_RunAnimationGroup(changes func(context AnimationContext)) {
	AnimationContextClass.RunAnimationGroup(changes)
}

func (ac _AnimationContextClass) CurrentContext() AnimationContext {
	rv := objc.CallMethod[AnimationContext](ac, objc.GetSelector("currentContext"))
	return rv
}

func AnimationContext_CurrentContext() AnimationContext {
	return AnimationContextClass.CurrentContext()
}

func (a_ AnimationContext) CompletionHandler() func() {
	rv := objc.CallMethod[func()](a_, objc.GetSelector("completionHandler"))
	return rv
}

func (a_ AnimationContext) SetCompletionHandler(value func()) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("setCompletionHandler:"), value)
}

func (a_ AnimationContext) Duration() foundation.TimeInterval {
	rv := objc.CallMethod[foundation.TimeInterval](a_, objc.GetSelector("duration"))
	return rv
}

func (a_ AnimationContext) SetDuration(value foundation.TimeInterval) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("setDuration:"), value)
}

func (a_ AnimationContext) TimingFunction() quartzcore.MediaTimingFunction {
	rv := objc.CallMethod[quartzcore.MediaTimingFunction](a_, objc.GetSelector("timingFunction"))
	return rv
}

func (a_ AnimationContext) SetTimingFunction(value quartzcore.IMediaTimingFunction) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("setTimingFunction:"), objc.ExtractPtr(value))
}

func (a_ AnimationContext) AllowsImplicitAnimation() bool {
	rv := objc.CallMethod[bool](a_, objc.GetSelector("allowsImplicitAnimation"))
	return rv
}

func (a_ AnimationContext) SetAllowsImplicitAnimation(value bool) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("setAllowsImplicitAnimation:"), value)
}
