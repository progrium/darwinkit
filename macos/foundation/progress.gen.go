// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Progress] class.
var ProgressClass = _ProgressClass{objc.GetClass("NSProgress")}

type _ProgressClass struct {
	objc.Class
}

// An interface definition for the [Progress] class.
type IProgress interface {
	objc.IObject
	Unpublish()
	BecomeCurrentWithPendingUnitCount(unitCount int64)
	Publish()
	AddChildWithPendingUnitCount(child IProgress, inUnitCount int64)
	ResignCurrent()
	SetUserInfoObjectForKey(objectOrNil objc.IObject, key ProgressUserInfoKey)
	Pause()
	Resume()
	Cancel()
	PerformAsCurrentWithPendingUnitCountUsingBlock(unitCount int64, work func())
	IsCancelled() bool
	LocalizedAdditionalDescription() string
	SetLocalizedAdditionalDescription(value string)
	LocalizedDescription() string
	SetLocalizedDescription(value string)
	FileOperationKind() ProgressFileOperationKind
	SetFileOperationKind(value ProgressFileOperationKind)
	ResumingHandler() func()
	SetResumingHandler(value func())
	EstimatedTimeRemaining() Number
	SetEstimatedTimeRemaining(value INumber)
	FileTotalCount() Number
	SetFileTotalCount(value INumber)
	IsOld() bool
	FileURL() URL
	SetFileURL(value IURL)
	Throughput() Number
	SetThroughput(value INumber)
	FractionCompleted() float64
	TotalUnitCount() int64
	SetTotalUnitCount(value int64)
	UserInfo() map[ProgressUserInfoKey]objc.Object
	PausingHandler() func()
	SetPausingHandler(value func())
	IsCancellable() bool
	SetCancellable(value bool)
	CompletedUnitCount() int64
	SetCompletedUnitCount(value int64)
	FileCompletedCount() Number
	SetFileCompletedCount(value INumber)
	CancellationHandler() func()
	SetCancellationHandler(value func())
	Kind() ProgressKind
	SetKind(value ProgressKind)
	IsPausable() bool
	SetPausable(value bool)
	IsPaused() bool
	IsFinished() bool
	IsIndeterminate() bool
}

// An object that conveys ongoing progress to the user for a specified task. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprogress?language=objc
type Progress struct {
	objc.Object
}

func ProgressFrom(ptr unsafe.Pointer) Progress {
	return Progress{
		Object: objc.ObjectFrom(ptr),
	}
}

func (p_ Progress) InitWithParentUserInfo(parentProgressOrNil IProgress, userInfoOrNil map[ProgressUserInfoKey]objc.IObject) Progress {
	rv := objc.Call[Progress](p_, objc.Sel("initWithParent:userInfo:"), objc.Ptr(parentProgressOrNil), userInfoOrNil)
	return rv
}

// Creates a new progress instance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprogress/1409133-initwithparent?language=objc
func NewProgressWithParentUserInfo(parentProgressOrNil IProgress, userInfoOrNil map[ProgressUserInfoKey]objc.IObject) Progress {
	instance := ProgressClass.Alloc().InitWithParentUserInfo(parentProgressOrNil, userInfoOrNil)
	instance.Autorelease()
	return instance
}

func (pc _ProgressClass) Alloc() Progress {
	rv := objc.Call[Progress](pc, objc.Sel("alloc"))
	return rv
}

func Progress_Alloc() Progress {
	return ProgressClass.Alloc()
}

func (pc _ProgressClass) New() Progress {
	rv := objc.Call[Progress](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewProgress() Progress {
	return ProgressClass.New()
}

func (p_ Progress) Init() Progress {
	rv := objc.Call[Progress](p_, objc.Sel("init"))
	return rv
}

// Removes a progress object from publication, making it unobservable by other processes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprogress/1413268-unpublish?language=objc
func (p_ Progress) Unpublish() {
	objc.Call[objc.Void](p_, objc.Sel("unpublish"))
}

// Returns the progress instance, if any. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprogress/1412499-currentprogress?language=objc
func (pc _ProgressClass) CurrentProgress() Progress {
	rv := objc.Call[Progress](pc, objc.Sel("currentProgress"))
	return rv
}

// Returns the progress instance, if any. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprogress/1412499-currentprogress?language=objc
func Progress_CurrentProgress() Progress {
	return ProgressClass.CurrentProgress()
}

// Sets the progress object as the current object of the current thread, and assigns the amount of work for the next suboperation progress object to perform. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprogress/1410103-becomecurrentwithpendingunitcoun?language=objc
func (p_ Progress) BecomeCurrentWithPendingUnitCount(unitCount int64) {
	objc.Call[objc.Void](p_, objc.Sel("becomeCurrentWithPendingUnitCount:"), unitCount)
}

// Creates and returns a progress instance with the specified unit count that isn’t part of any existing progress tree. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprogress/1410951-discreteprogresswithtotalunitcou?language=objc
func (pc _ProgressClass) DiscreteProgressWithTotalUnitCount(unitCount int64) Progress {
	rv := objc.Call[Progress](pc, objc.Sel("discreteProgressWithTotalUnitCount:"), unitCount)
	return rv
}

// Creates and returns a progress instance with the specified unit count that isn’t part of any existing progress tree. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprogress/1410951-discreteprogresswithtotalunitcou?language=objc
func Progress_DiscreteProgressWithTotalUnitCount(unitCount int64) Progress {
	return ProgressClass.DiscreteProgressWithTotalUnitCount(unitCount)
}

// Creates and returns a progress instance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprogress/1415509-progresswithtotalunitcount?language=objc
func (pc _ProgressClass) ProgressWithTotalUnitCount(unitCount int64) Progress {
	rv := objc.Call[Progress](pc, objc.Sel("progressWithTotalUnitCount:"), unitCount)
	return rv
}

// Creates and returns a progress instance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprogress/1415509-progresswithtotalunitcount?language=objc
func Progress_ProgressWithTotalUnitCount(unitCount int64) Progress {
	return ProgressClass.ProgressWithTotalUnitCount(unitCount)
}

// Registers a file URL to hear about the progress of a file operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprogress/1418475-addsubscriberforfileurl?language=objc
func (pc _ProgressClass) AddSubscriberForFileURLWithPublishingHandler(url IURL, publishingHandler ProgressPublishingHandler) objc.Object {
	rv := objc.Call[objc.Object](pc, objc.Sel("addSubscriberForFileURL:withPublishingHandler:"), objc.Ptr(url), publishingHandler)
	return rv
}

// Registers a file URL to hear about the progress of a file operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprogress/1418475-addsubscriberforfileurl?language=objc
func Progress_AddSubscriberForFileURLWithPublishingHandler(url IURL, publishingHandler ProgressPublishingHandler) objc.Object {
	return ProgressClass.AddSubscriberForFileURLWithPublishingHandler(url, publishingHandler)
}

// Publishes the progress object for other processes to observe it. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprogress/1416782-publish?language=objc
func (p_ Progress) Publish() {
	objc.Call[objc.Void](p_, objc.Sel("publish"))
}

// Removes a proxy progress object that the add subscriber method returns. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprogress/1410457-removesubscriber?language=objc
func (pc _ProgressClass) RemoveSubscriber(subscriber objc.IObject) {
	objc.Call[objc.Void](pc, objc.Sel("removeSubscriber:"), subscriber)
}

// Removes a proxy progress object that the add subscriber method returns. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprogress/1410457-removesubscriber?language=objc
func Progress_RemoveSubscriber(subscriber objc.IObject) {
	ProgressClass.RemoveSubscriber(subscriber)
}

// Adds a process object as a suboperation of a progress tree. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprogress/1417260-addchild?language=objc
func (p_ Progress) AddChildWithPendingUnitCount(child IProgress, inUnitCount int64) {
	objc.Call[objc.Void](p_, objc.Sel("addChild:withPendingUnitCount:"), objc.Ptr(child), inUnitCount)
}

// Restores the previous progress object to become the current progress object on the thread. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprogress/1407180-resigncurrent?language=objc
func (p_ Progress) ResignCurrent() {
	objc.Call[objc.Void](p_, objc.Sel("resignCurrent"))
}

// Sets a value in the user info dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprogress/1407537-setuserinfoobject?language=objc
func (p_ Progress) SetUserInfoObjectForKey(objectOrNil objc.IObject, key ProgressUserInfoKey) {
	objc.Call[objc.Void](p_, objc.Sel("setUserInfoObject:forKey:"), objectOrNil, key)
}

// Pauses progress tracking. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprogress/1412377-pause?language=objc
func (p_ Progress) Pause() {
	objc.Call[objc.Void](p_, objc.Sel("pause"))
}

// Resumes progress tracking. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprogress/1413616-resume?language=objc
func (p_ Progress) Resume() {
	objc.Call[objc.Void](p_, objc.Sel("resume"))
}

// Cancels progress tracking. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprogress/1413832-cancel?language=objc
func (p_ Progress) Cancel() {
	objc.Call[objc.Void](p_, objc.Sel("cancel"))
}

// Retrieves the current thread’s progress object, executes the specified block, and increments the progress object by the specified units of work. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprogress/2865587-performascurrentwithpendingunitc?language=objc
func (p_ Progress) PerformAsCurrentWithPendingUnitCountUsingBlock(unitCount int64, work func()) {
	objc.Call[objc.Void](p_, objc.Sel("performAsCurrentWithPendingUnitCount:usingBlock:"), unitCount, work)
}

// A Boolean value that Indicates whether the receiver is tracking canceled work. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprogress/1414454-cancelled?language=objc
func (p_ Progress) IsCancelled() bool {
	rv := objc.Call[bool](p_, objc.Sel("isCancelled"))
	return rv
}

// A more specific localized description of tracked progress for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprogress/1412455-localizedadditionaldescription?language=objc
func (p_ Progress) LocalizedAdditionalDescription() string {
	rv := objc.Call[string](p_, objc.Sel("localizedAdditionalDescription"))
	return rv
}

// A more specific localized description of tracked progress for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprogress/1412455-localizedadditionaldescription?language=objc
func (p_ Progress) SetLocalizedAdditionalDescription(value string) {
	objc.Call[objc.Void](p_, objc.Sel("setLocalizedAdditionalDescription:"), value)
}

// A localized description of tracked progress for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprogress/1417251-localizeddescription?language=objc
func (p_ Progress) LocalizedDescription() string {
	rv := objc.Call[string](p_, objc.Sel("localizedDescription"))
	return rv
}

// A localized description of tracked progress for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprogress/1417251-localizeddescription?language=objc
func (p_ Progress) SetLocalizedDescription(value string) {
	objc.Call[objc.Void](p_, objc.Sel("setLocalizedDescription:"), value)
}

// The kind of file operation for the progress object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprogress/2865625-fileoperationkind?language=objc
func (p_ Progress) FileOperationKind() ProgressFileOperationKind {
	rv := objc.Call[ProgressFileOperationKind](p_, objc.Sel("fileOperationKind"))
	return rv
}

// The kind of file operation for the progress object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprogress/2865625-fileoperationkind?language=objc
func (p_ Progress) SetFileOperationKind(value ProgressFileOperationKind) {
	objc.Call[objc.Void](p_, objc.Sel("setFileOperationKind:"), value)
}

// The block to invoke when progress resumes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprogress/1410158-resuminghandler?language=objc
func (p_ Progress) ResumingHandler() func() {
	rv := objc.Call[func()](p_, objc.Sel("resumingHandler"))
	return rv
}

// The block to invoke when progress resumes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprogress/1410158-resuminghandler?language=objc
func (p_ Progress) SetResumingHandler(value func()) {
	objc.Call[objc.Void](p_, objc.Sel("setResumingHandler:"), value)
}

// A value that indicates the estimated amount of time remaining to complete the progress. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprogress/2868041-estimatedtimeremaining?language=objc
func (p_ Progress) EstimatedTimeRemaining() Number {
	rv := objc.Call[Number](p_, objc.Sel("estimatedTimeRemaining"))
	return rv
}

// A value that indicates the estimated amount of time remaining to complete the progress. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprogress/2868041-estimatedtimeremaining?language=objc
func (p_ Progress) SetEstimatedTimeRemaining(value INumber) {
	objc.Call[objc.Void](p_, objc.Sel("setEstimatedTimeRemaining:"), objc.Ptr(value))
}

// The total number of files for a file progress object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprogress/2868045-filetotalcount?language=objc
func (p_ Progress) FileTotalCount() Number {
	rv := objc.Call[Number](p_, objc.Sel("fileTotalCount"))
	return rv
}

// The total number of files for a file progress object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprogress/2868045-filetotalcount?language=objc
func (p_ Progress) SetFileTotalCount(value INumber) {
	objc.Call[objc.Void](p_, objc.Sel("setFileTotalCount:"), objc.Ptr(value))
}

// A Boolean value that indicates when the observed progress object invokes the publish method before you subscribe to it. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprogress/1407931-old?language=objc
func (p_ Progress) IsOld() bool {
	rv := objc.Call[bool](p_, objc.Sel("isOld"))
	return rv
}

// A URL that represents the file for the current progress object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprogress/2865663-fileurl?language=objc
func (p_ Progress) FileURL() URL {
	rv := objc.Call[URL](p_, objc.Sel("fileURL"))
	return rv
}

// A URL that represents the file for the current progress object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprogress/2865663-fileurl?language=objc
func (p_ Progress) SetFileURL(value IURL) {
	objc.Call[objc.Void](p_, objc.Sel("setFileURL:"), objc.Ptr(value))
}

// A value that represents the speed of data processing, in bytes per second. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprogress/2868052-throughput?language=objc
func (p_ Progress) Throughput() Number {
	rv := objc.Call[Number](p_, objc.Sel("throughput"))
	return rv
}

// A value that represents the speed of data processing, in bytes per second. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprogress/2868052-throughput?language=objc
func (p_ Progress) SetThroughput(value INumber) {
	objc.Call[objc.Void](p_, objc.Sel("setThroughput:"), objc.Ptr(value))
}

// The fraction of the overall work that the progress object completes, including work from its suboperations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprogress/1408579-fractioncompleted?language=objc
func (p_ Progress) FractionCompleted() float64 {
	rv := objc.Call[float64](p_, objc.Sel("fractionCompleted"))
	return rv
}

// The total number of tracked units of work for the current progress. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprogress/1410940-totalunitcount?language=objc
func (p_ Progress) TotalUnitCount() int64 {
	rv := objc.Call[int64](p_, objc.Sel("totalUnitCount"))
	return rv
}

// The total number of tracked units of work for the current progress. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprogress/1410940-totalunitcount?language=objc
func (p_ Progress) SetTotalUnitCount(value int64) {
	objc.Call[objc.Void](p_, objc.Sel("setTotalUnitCount:"), value)
}

// A dictionary of arbitrary values for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprogress/1413314-userinfo?language=objc
func (p_ Progress) UserInfo() map[ProgressUserInfoKey]objc.Object {
	rv := objc.Call[map[ProgressUserInfoKey]objc.Object](p_, objc.Sel("userInfo"))
	return rv
}

// The block to invoke when pausing progress. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprogress/1412873-pausinghandler?language=objc
func (p_ Progress) PausingHandler() func() {
	rv := objc.Call[func()](p_, objc.Sel("pausingHandler"))
	return rv
}

// The block to invoke when pausing progress. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprogress/1412873-pausinghandler?language=objc
func (p_ Progress) SetPausingHandler(value func()) {
	objc.Call[objc.Void](p_, objc.Sel("setPausingHandler:"), value)
}

// A Boolean value that indicates whether the receiver is tracking work that you can cancel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprogress/1409348-cancellable?language=objc
func (p_ Progress) IsCancellable() bool {
	rv := objc.Call[bool](p_, objc.Sel("isCancellable"))
	return rv
}

// A Boolean value that indicates whether the receiver is tracking work that you can cancel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprogress/1409348-cancellable?language=objc
func (p_ Progress) SetCancellable(value bool) {
	objc.Call[objc.Void](p_, objc.Sel("setCancellable:"), value)
}

// The number of completed units of work for the current job. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprogress/1407934-completedunitcount?language=objc
func (p_ Progress) CompletedUnitCount() int64 {
	rv := objc.Call[int64](p_, objc.Sel("completedUnitCount"))
	return rv
}

// The number of completed units of work for the current job. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprogress/1407934-completedunitcount?language=objc
func (p_ Progress) SetCompletedUnitCount(value int64) {
	objc.Call[objc.Void](p_, objc.Sel("setCompletedUnitCount:"), value)
}

// The number of completed files for a file progress object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprogress/2868050-filecompletedcount?language=objc
func (p_ Progress) FileCompletedCount() Number {
	rv := objc.Call[Number](p_, objc.Sel("fileCompletedCount"))
	return rv
}

// The number of completed files for a file progress object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprogress/2868050-filecompletedcount?language=objc
func (p_ Progress) SetFileCompletedCount(value INumber) {
	objc.Call[objc.Void](p_, objc.Sel("setFileCompletedCount:"), objc.Ptr(value))
}

// The block to invoke when canceling progress. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprogress/1408913-cancellationhandler?language=objc
func (p_ Progress) CancellationHandler() func() {
	rv := objc.Call[func()](p_, objc.Sel("cancellationHandler"))
	return rv
}

// The block to invoke when canceling progress. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprogress/1408913-cancellationhandler?language=objc
func (p_ Progress) SetCancellationHandler(value func()) {
	objc.Call[objc.Void](p_, objc.Sel("setCancellationHandler:"), value)
}

// An object that represents the kind of progress for the progress object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprogress/1416139-kind?language=objc
func (p_ Progress) Kind() ProgressKind {
	rv := objc.Call[ProgressKind](p_, objc.Sel("kind"))
	return rv
}

// An object that represents the kind of progress for the progress object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprogress/1416139-kind?language=objc
func (p_ Progress) SetKind(value ProgressKind) {
	objc.Call[objc.Void](p_, objc.Sel("setKind:"), value)
}

// A Boolean value that indicates whether the receiver is tracking work that you can pause. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprogress/1417421-pausable?language=objc
func (p_ Progress) IsPausable() bool {
	rv := objc.Call[bool](p_, objc.Sel("isPausable"))
	return rv
}

// A Boolean value that indicates whether the receiver is tracking work that you can pause. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprogress/1417421-pausable?language=objc
func (p_ Progress) SetPausable(value bool) {
	objc.Call[objc.Void](p_, objc.Sel("setPausable:"), value)
}

// A Boolean value that indicates whether the receiver is tracking paused work. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprogress/1415495-paused?language=objc
func (p_ Progress) IsPaused() bool {
	rv := objc.Call[bool](p_, objc.Sel("isPaused"))
	return rv
}

// A Boolean value that indicates the progress object is complete. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprogress/2865603-finished?language=objc
func (p_ Progress) IsFinished() bool {
	rv := objc.Call[bool](p_, objc.Sel("isFinished"))
	return rv
}

// A Boolean value that indicates whether the tracked progress is indeterminate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprogress/1412871-indeterminate?language=objc
func (p_ Progress) IsIndeterminate() bool {
	rv := objc.Call[bool](p_, objc.Sel("isIndeterminate"))
	return rv
}
