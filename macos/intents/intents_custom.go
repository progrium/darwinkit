package intents

import (
	"github.com/progrium/darwinkit/macos/foundation"
	"github.com/progrium/darwinkit/objc"
)

// Intent represents an intention or action to be performed
type Intent struct {
	objc.Object
}

// IntentClass is the class object for Intent
var IntentClass = objc.GetClass("INIntent")

// NewIntent creates a new Intent instance
func NewIntent() Intent {
	alloc := objc.Call[objc.Object](IntentClass, objc.Sel("alloc"))
	obj := objc.Call[objc.Object](alloc, objc.Sel("init"))
	return Intent{obj}
}

// Identifier returns the identifier of the intent
func (i Intent) Identifier() foundation.String {
	return objc.Call[foundation.String](i, objc.Sel("identifier"))
}

// IntentResponse represents a response to an intent
type IntentResponse struct {
	objc.Object
}

// IntentResponseClass is the class object for IntentResponse
var IntentResponseClass = objc.GetClass("INIntentResponse")

// NewIntentResponse creates a new IntentResponse instance
func NewIntentResponse(code INIntentResponseCode) IntentResponse {
	alloc := objc.Call[objc.Object](IntentResponseClass, objc.Sel("alloc"))
	obj := objc.Call[objc.Object](alloc, objc.Sel("initWithCode:"), code)
	return IntentResponse{obj}
}

// Code returns the response code
func (r IntentResponse) Code() INIntentResponseCode {
	return INIntentResponseCode(objc.Call[int](r, objc.Sel("code")))
}

// UserActivity returns the user activity associated with the response
func (r IntentResponse) UserActivity() foundation.UserActivity {
	return objc.Call[foundation.UserActivity](r, objc.Sel("userActivity"))
}

// IntentHandling represents a protocol for handling intents
type IntentHandling struct {
	objc.Object
}

// IntentHandlingClass is the class object for IntentHandling
var IntentHandlingClass = objc.GetClass("INIntentHandling")

// HandleIntent handles an intent and returns a response
func (h IntentHandling) HandleIntent(intent Intent, completion objc.Object) {
	objc.Call[objc.Void](h, objc.Sel("handleIntent:completion:"), intent, completion)
}

// Interaction represents an interaction between the user and the app
type Interaction struct {
	objc.Object
}

// InteractionClass is the class object for Interaction
var InteractionClass = objc.GetClass("INInteraction")

// NewInteraction creates a new Interaction instance
func NewInteraction(intent Intent, response IntentResponse) Interaction {
	alloc := objc.Call[objc.Object](InteractionClass, objc.Sel("alloc"))
	obj := objc.Call[objc.Object](alloc, objc.Sel("initWithIntent:response:"), intent, response)
	return Interaction{obj}
}

// Intent returns the intent associated with the interaction
func (i Interaction) Intent() Intent {
	return Intent{objc.Call[objc.Object](i, objc.Sel("intent"))}
}

// IntentResponse returns the response associated with the interaction
func (i Interaction) IntentResponse() IntentResponse {
	return IntentResponse{objc.Call[objc.Object](i, objc.Sel("intentResponse"))}
}

// Direction returns the direction of the interaction
func (i Interaction) Direction() INInteractionDirection {
	return INInteractionDirection(objc.Call[int](i, objc.Sel("direction")))
}

// DateInterval returns the date interval of the interaction
func (i Interaction) DateInterval() foundation.DateInterval {
	return objc.Call[foundation.DateInterval](i, objc.Sel("dateInterval"))
}

// DonateInteractionWithCompletion donates the interaction to the system
func (i Interaction) DonateInteractionWithCompletion(completion objc.Object) {
	objc.Call[objc.Void](i, objc.Sel("donateInteractionWithCompletion:"), completion)
}

// DeleteAllInteractionsWithCompletion deletes all interactions
func DeleteAllInteractionsWithCompletion(completion objc.Object) {
	objc.Call[objc.Void](InteractionClass, objc.Sel("deleteAllInteractionsWithCompletion:"), completion)
}

// Person represents a person in the Intents framework
type Person struct {
	objc.Object
}

// PersonClass is the class object for Person
var PersonClass = objc.GetClass("INPerson")

// NewPerson creates a new Person instance
func NewPerson(personHandle PersonHandle, nameComponents INPersonNameComponents, displayName foundation.String) Person {
	alloc := objc.Call[objc.Object](PersonClass, objc.Sel("alloc"))
	obj := objc.Call[objc.Object](alloc, objc.Sel("initWithPersonHandle:nameComponents:displayName:"), personHandle, nameComponents, displayName)
	return Person{obj}
}

// PersonHandle represents a handle for a person
type PersonHandle struct {
	objc.Object
}

// PersonHandleClass is the class object for PersonHandle
var PersonHandleClass = objc.GetClass("INPersonHandle")

// NewPersonHandle creates a new PersonHandle instance
func NewPersonHandle(value foundation.String, typ int) PersonHandle {
	alloc := objc.Call[objc.Object](PersonHandleClass, objc.Sel("alloc"))
	obj := objc.Call[objc.Object](alloc, objc.Sel("initWithValue:type:"), value, typ)
	return PersonHandle{obj}
}

// Value returns the value of the person handle
func (h PersonHandle) Value() foundation.String {
	return objc.Call[foundation.String](h, objc.Sel("value"))
}

// Type returns the type of the person handle
func (h PersonHandle) Type() int {
	return objc.Call[int](h, objc.Sel("type"))
}

// Vocabulary represents a vocabulary used for intents
type Vocabulary struct {
	objc.Object
}

// VocabularyClass is the class object for Vocabulary
var VocabularyClass = objc.GetClass("INVocabulary")

// SharedVocabulary returns the shared vocabulary instance
func SharedVocabulary() Vocabulary {
	return Vocabulary{objc.Call[objc.Object](VocabularyClass, objc.Sel("sharedVocabulary"))}
}

// RemoveAllVocabularyStrings removes all vocabulary strings
func (v Vocabulary) RemoveAllVocabularyStrings() {
	objc.Call[objc.Void](v, objc.Sel("removeAllVocabularyStrings"))
}

// IntentDefinition represents a definition of an intent
type IntentDefinition struct {
	objc.Object
}

// IntentDefinitionClass is the class object for IntentDefinition
var IntentDefinitionClass = objc.GetClass("INIntentDefinition")

// NewIntentDefinition creates a new IntentDefinition instance
func NewIntentDefinition() IntentDefinition {
	alloc := objc.Call[objc.Object](IntentDefinitionClass, objc.Sel("alloc"))
	obj := objc.Call[objc.Object](alloc, objc.Sel("init"))
	return IntentDefinition{obj}
}

// SetIdentifier sets the identifier of the intent definition
func (d IntentDefinition) SetIdentifier(identifier foundation.String) {
	objc.Call[objc.Void](d, objc.Sel("setIdentifier:"), identifier)
}

// SetCategoryName sets the category name of the intent definition
func (d IntentDefinition) SetCategoryName(categoryName foundation.String) {
	objc.Call[objc.Void](d, objc.Sel("setCategoryName:"), categoryName)
}

// Parameters returns the parameters of the intent definition
func (d IntentDefinition) Parameters() foundation.Array {
	return objc.Call[foundation.Array](d, objc.Sel("parameters"))
}

// SetParameters sets the parameters of the intent definition
func (d IntentDefinition) SetParameters(parameters foundation.Array) {
	objc.Call[objc.Void](d, objc.Sel("setParameters:"), parameters)
}

// SendMessageIntent represents an intent to send a message
type SendMessageIntent struct {
	Intent
}

// SendMessageIntentClass is the class object for SendMessageIntent
var SendMessageIntentClass = objc.GetClass("INSendMessageIntent")

// NewSendMessageIntent creates a new SendMessageIntent instance
func NewSendMessageIntent(recipients foundation.Array, content foundation.String, speakableGroupName foundation.String, conversationIdentifier foundation.String) SendMessageIntent {
	alloc := objc.Call[objc.Object](SendMessageIntentClass, objc.Sel("alloc"))
	obj := objc.Call[objc.Object](alloc, objc.Sel("initWithRecipients:content:speakableGroupName:conversationIdentifier:"), recipients, content, speakableGroupName, conversationIdentifier)
	return SendMessageIntent{Intent{obj}}
}

// Recipients returns the recipients of the message
func (i SendMessageIntent) Recipients() foundation.Array {
	return objc.Call[foundation.Array](i, objc.Sel("recipients"))
}

// Content returns the content of the message
func (i SendMessageIntent) Content() foundation.String {
	return objc.Call[foundation.String](i, objc.Sel("content"))
}

// SearchForMessagesIntent represents an intent to search for messages
type SearchForMessagesIntent struct {
	Intent
}

// SearchForMessagesIntentClass is the class object for SearchForMessagesIntent
var SearchForMessagesIntentClass = objc.GetClass("INSearchForMessagesIntent")

// NewSearchForMessagesIntent creates a new SearchForMessagesIntent instance
func NewSearchForMessagesIntent(recipients foundation.Array, senders foundation.Array, searchTerms foundation.String, attributes foundation.Array, dateTimeRange foundation.DateInterval, identifiers foundation.Array) SearchForMessagesIntent {
	alloc := objc.Call[objc.Object](SearchForMessagesIntentClass, objc.Sel("alloc"))
	obj := objc.Call[objc.Object](alloc, objc.Sel("initWithRecipients:senders:searchTerms:attributes:dateTimeRange:identifiers:"), recipients, senders, searchTerms, attributes, dateTimeRange, identifiers)
	return SearchForMessagesIntent{Intent{obj}}
}

// Recipients returns the recipients of the messages to search for
func (i SearchForMessagesIntent) Recipients() foundation.Array {
	return objc.Call[foundation.Array](i, objc.Sel("recipients"))
}

// Senders returns the senders of the messages to search for
func (i SearchForMessagesIntent) Senders() foundation.Array {
	return objc.Call[foundation.Array](i, objc.Sel("senders"))
}

// SetSenders sets the senders of the messages to search for
func (i SearchForMessagesIntent) SetSenders(senders foundation.Array) {
	objc.Call[objc.Void](i, objc.Sel("setSenders:"), senders)
}

// SearchTerms returns the search terms for the messages
func (i SearchForMessagesIntent) SearchTerms() foundation.String {
	return objc.Call[foundation.String](i, objc.Sel("searchTerms"))
}

// SetSearchTerms sets the search terms for the messages
func (i SearchForMessagesIntent) SetSearchTerms(searchTerms foundation.String) {
	objc.Call[objc.Void](i, objc.Sel("setSearchTerms:"), searchTerms)
}