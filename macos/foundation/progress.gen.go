// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var ProgressClass = _ProgressClass{objc.GetClass("NSProgress")}

type _ProgressClass struct {
	objc.Class
}

type IProgress interface {
	objc.IObject
	BecomeCurrentWithPendingUnitCount(unitCount int64)
	AddChildWithPendingUnitCount(child IProgress, inUnitCount int64)
	PerformAsCurrentWithPendingUnitCountUsingBlock(unitCount int64, work func())
	ResignCurrent()
	Cancel()
	Pause()
	Resume()
	SetUserInfoObjectForKey(objectOrNil objc.IObject, key ProgressUserInfoKey)
	Publish()
	Unpublish()
	TotalUnitCount() int64
	SetTotalUnitCount(value int64)
	CompletedUnitCount() int64
	SetCompletedUnitCount(value int64)
	LocalizedDescription() string
	SetLocalizedDescription(value string)
	LocalizedAdditionalDescription() string
	SetLocalizedAdditionalDescription(value string)
	IsCancellable() bool
	SetCancellable(value bool)
	IsCancelled() bool
	CancellationHandler() func()
	SetCancellationHandler(value func())
	IsPausable() bool
	SetPausable(value bool)
	IsPaused() bool
	PausingHandler() func()
	SetPausingHandler(value func())
	IsIndeterminate() bool
	FractionCompleted() float64
	IsFinished() bool
	ResumingHandler() func()
	SetResumingHandler(value func())
	Kind() ProgressKind
	SetKind(value ProgressKind)
	EstimatedTimeRemaining() Number
	SetEstimatedTimeRemaining(value INumber)
	Throughput() Number
	SetThroughput(value INumber)
	UserInfo() map[ProgressUserInfoKey]objc.Object
	FileOperationKind() ProgressFileOperationKind
	SetFileOperationKind(value ProgressFileOperationKind)
	FileURL() URL
	SetFileURL(value IURL)
	FileTotalCount() Number
	SetFileTotalCount(value INumber)
	FileCompletedCount() Number
	SetFileCompletedCount(value INumber)
	IsOld() bool
}

type Progress struct {
	objc.Object
}

func MakeProgress(ptr unsafe.Pointer) Progress {
	return Progress{
		Object: objc.MakeObject(ptr),
	}
}

func (p_ Progress) InitWithParentUserInfo(parentProgressOrNil IProgress, userInfoOrNil map[ProgressUserInfoKey]objc.IObject) Progress {
	rv := objc.CallMethod[Progress](p_, objc.GetSelector("initWithParent:userInfo:"), objc.ExtractPtr(parentProgressOrNil), userInfoOrNil)
	return rv
}

func Progress_InitWithParentUserInfo(parentProgressOrNil IProgress, userInfoOrNil map[ProgressUserInfoKey]objc.IObject) Progress {
	return ProgressClass.Alloc().InitWithParentUserInfo(parentProgressOrNil, userInfoOrNil)
}

func (pc _ProgressClass) Alloc() Progress {
	rv := objc.CallMethod[Progress](pc, objc.GetSelector("alloc"))
	return rv
}

func Progress_Alloc() Progress {
	return ProgressClass.Alloc()
}

func (pc _ProgressClass) New() Progress {
	rv := objc.CallMethod[Progress](pc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewProgress() Progress {
	return ProgressClass.New()
}

func Progress_New() Progress {
	return ProgressClass.New()
}

func (p_ Progress) Init() Progress {
	rv := objc.CallMethod[Progress](p_, objc.GetSelector("init"))
	return rv
}

func Progress_Init() Progress {
	return ProgressClass.Alloc().Init()
}

func (pc _ProgressClass) DiscreteProgressWithTotalUnitCount(unitCount int64) Progress {
	rv := objc.CallMethod[Progress](pc, objc.GetSelector("discreteProgressWithTotalUnitCount:"), unitCount)
	return rv
}

func Progress_DiscreteProgressWithTotalUnitCount(unitCount int64) Progress {
	return ProgressClass.DiscreteProgressWithTotalUnitCount(unitCount)
}

func (pc _ProgressClass) ProgressWithTotalUnitCount(unitCount int64) Progress {
	rv := objc.CallMethod[Progress](pc, objc.GetSelector("progressWithTotalUnitCount:"), unitCount)
	return rv
}

func Progress_ProgressWithTotalUnitCount(unitCount int64) Progress {
	return ProgressClass.ProgressWithTotalUnitCount(unitCount)
}

func (pc _ProgressClass) ProgressWithTotalUnitCountParentPendingUnitCount(unitCount int64, parent IProgress, portionOfParentTotalUnitCount int64) Progress {
	rv := objc.CallMethod[Progress](pc, objc.GetSelector("progressWithTotalUnitCount:parent:pendingUnitCount:"), unitCount, objc.ExtractPtr(parent), portionOfParentTotalUnitCount)
	return rv
}

func Progress_ProgressWithTotalUnitCountParentPendingUnitCount(unitCount int64, parent IProgress, portionOfParentTotalUnitCount int64) Progress {
	return ProgressClass.ProgressWithTotalUnitCountParentPendingUnitCount(unitCount, parent, portionOfParentTotalUnitCount)
}

func (pc _ProgressClass) CurrentProgress() Progress {
	rv := objc.CallMethod[Progress](pc, objc.GetSelector("currentProgress"))
	return rv
}

func Progress_CurrentProgress() Progress {
	return ProgressClass.CurrentProgress()
}

func (p_ Progress) BecomeCurrentWithPendingUnitCount(unitCount int64) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("becomeCurrentWithPendingUnitCount:"), unitCount)
}

func (p_ Progress) AddChildWithPendingUnitCount(child IProgress, inUnitCount int64) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("addChild:withPendingUnitCount:"), objc.ExtractPtr(child), inUnitCount)
}

func (p_ Progress) PerformAsCurrentWithPendingUnitCountUsingBlock(unitCount int64, work func()) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("performAsCurrentWithPendingUnitCount:usingBlock:"), unitCount, work)
}

func (p_ Progress) ResignCurrent() {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("resignCurrent"))
}

func (p_ Progress) Cancel() {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("cancel"))
}

func (p_ Progress) Pause() {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("pause"))
}

func (p_ Progress) Resume() {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("resume"))
}

func (p_ Progress) SetUserInfoObjectForKey(objectOrNil objc.IObject, key ProgressUserInfoKey) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setUserInfoObject:forKey:"), objc.ExtractPtr(objectOrNil), key)
}

func (p_ Progress) Publish() {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("publish"))
}

func (p_ Progress) Unpublish() {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("unpublish"))
}

func (pc _ProgressClass) RemoveSubscriber(subscriber objc.IObject) {
	objc.CallMethod[objc.Void](pc, objc.GetSelector("removeSubscriber:"), objc.ExtractPtr(subscriber))
}

func Progress_RemoveSubscriber(subscriber objc.IObject) {
	ProgressClass.RemoveSubscriber(subscriber)
}

func (p_ Progress) TotalUnitCount() int64 {
	rv := objc.CallMethod[int64](p_, objc.GetSelector("totalUnitCount"))
	return rv
}

func (p_ Progress) SetTotalUnitCount(value int64) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setTotalUnitCount:"), value)
}

func (p_ Progress) CompletedUnitCount() int64 {
	rv := objc.CallMethod[int64](p_, objc.GetSelector("completedUnitCount"))
	return rv
}

func (p_ Progress) SetCompletedUnitCount(value int64) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setCompletedUnitCount:"), value)
}

func (p_ Progress) LocalizedDescription() string {
	rv := objc.CallMethod[string](p_, objc.GetSelector("localizedDescription"))
	return rv
}

func (p_ Progress) SetLocalizedDescription(value string) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setLocalizedDescription:"), value)
}

func (p_ Progress) LocalizedAdditionalDescription() string {
	rv := objc.CallMethod[string](p_, objc.GetSelector("localizedAdditionalDescription"))
	return rv
}

func (p_ Progress) SetLocalizedAdditionalDescription(value string) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setLocalizedAdditionalDescription:"), value)
}

func (p_ Progress) IsCancellable() bool {
	rv := objc.CallMethod[bool](p_, objc.GetSelector("isCancellable"))
	return rv
}

func (p_ Progress) SetCancellable(value bool) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setCancellable:"), value)
}

func (p_ Progress) IsCancelled() bool {
	rv := objc.CallMethod[bool](p_, objc.GetSelector("isCancelled"))
	return rv
}

func (p_ Progress) CancellationHandler() func() {
	rv := objc.CallMethod[func()](p_, objc.GetSelector("cancellationHandler"))
	return rv
}

func (p_ Progress) SetCancellationHandler(value func()) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setCancellationHandler:"), value)
}

func (p_ Progress) IsPausable() bool {
	rv := objc.CallMethod[bool](p_, objc.GetSelector("isPausable"))
	return rv
}

func (p_ Progress) SetPausable(value bool) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setPausable:"), value)
}

func (p_ Progress) IsPaused() bool {
	rv := objc.CallMethod[bool](p_, objc.GetSelector("isPaused"))
	return rv
}

func (p_ Progress) PausingHandler() func() {
	rv := objc.CallMethod[func()](p_, objc.GetSelector("pausingHandler"))
	return rv
}

func (p_ Progress) SetPausingHandler(value func()) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setPausingHandler:"), value)
}

func (p_ Progress) IsIndeterminate() bool {
	rv := objc.CallMethod[bool](p_, objc.GetSelector("isIndeterminate"))
	return rv
}

func (p_ Progress) FractionCompleted() float64 {
	rv := objc.CallMethod[float64](p_, objc.GetSelector("fractionCompleted"))
	return rv
}

func (p_ Progress) IsFinished() bool {
	rv := objc.CallMethod[bool](p_, objc.GetSelector("isFinished"))
	return rv
}

func (p_ Progress) ResumingHandler() func() {
	rv := objc.CallMethod[func()](p_, objc.GetSelector("resumingHandler"))
	return rv
}

func (p_ Progress) SetResumingHandler(value func()) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setResumingHandler:"), value)
}

func (p_ Progress) Kind() ProgressKind {
	rv := objc.CallMethod[ProgressKind](p_, objc.GetSelector("kind"))
	return rv
}

func (p_ Progress) SetKind(value ProgressKind) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setKind:"), value)
}

func (p_ Progress) EstimatedTimeRemaining() Number {
	rv := objc.CallMethod[Number](p_, objc.GetSelector("estimatedTimeRemaining"))
	return rv
}

func (p_ Progress) SetEstimatedTimeRemaining(value INumber) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setEstimatedTimeRemaining:"), objc.ExtractPtr(value))
}

func (p_ Progress) Throughput() Number {
	rv := objc.CallMethod[Number](p_, objc.GetSelector("throughput"))
	return rv
}

func (p_ Progress) SetThroughput(value INumber) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setThroughput:"), objc.ExtractPtr(value))
}

func (p_ Progress) UserInfo() map[ProgressUserInfoKey]objc.Object {
	rv := objc.CallMethod[map[ProgressUserInfoKey]objc.Object](p_, objc.GetSelector("userInfo"))
	return rv
}

func (p_ Progress) FileOperationKind() ProgressFileOperationKind {
	rv := objc.CallMethod[ProgressFileOperationKind](p_, objc.GetSelector("fileOperationKind"))
	return rv
}

func (p_ Progress) SetFileOperationKind(value ProgressFileOperationKind) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setFileOperationKind:"), value)
}

func (p_ Progress) FileURL() URL {
	rv := objc.CallMethod[URL](p_, objc.GetSelector("fileURL"))
	return rv
}

func (p_ Progress) SetFileURL(value IURL) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setFileURL:"), objc.ExtractPtr(value))
}

func (p_ Progress) FileTotalCount() Number {
	rv := objc.CallMethod[Number](p_, objc.GetSelector("fileTotalCount"))
	return rv
}

func (p_ Progress) SetFileTotalCount(value INumber) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setFileTotalCount:"), objc.ExtractPtr(value))
}

func (p_ Progress) FileCompletedCount() Number {
	rv := objc.CallMethod[Number](p_, objc.GetSelector("fileCompletedCount"))
	return rv
}

func (p_ Progress) SetFileCompletedCount(value INumber) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setFileCompletedCount:"), objc.ExtractPtr(value))
}

func (p_ Progress) IsOld() bool {
	rv := objc.CallMethod[bool](p_, objc.GetSelector("isOld"))
	return rv
}
