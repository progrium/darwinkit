package main

import (
	"fmt"
	"time"

	"github.com/progrium/darwinkit/macos/appkit"
	"github.com/progrium/darwinkit/macos/foundation"
	"github.com/progrium/darwinkit/macos/intents"
	"github.com/progrium/darwinkit/objc"
)

func main() {
	objc.WithAutoreleasePool(func() {
		// Create the application
		app := appkit.SharedApplication()
		
		fmt.Println("Intents Framework Example")
		fmt.Println("-------------------------")
		
		// Create a person handle for a sender
		personHandle := intents.NewPersonHandle(foundation.StringWithString("johndoe@example.com"), 1) // Email type
		
		// Create name components for the person
		nameComponents := intents.INPersonNameComponents{
			GivenName:  "John",
			FamilyName: "Doe",
		}
		
		// Create a person object
		person := intents.NewPerson(personHandle, nameComponents, foundation.StringWithString("John Doe"))
		
		// Add person to an array for recipients
		peopleArray := foundation.ArrayWithObjects([]objc.Object{person})
		
		// Create a message content
		messageContent := foundation.StringWithString("Hello! This is a test message from DarwinKit Intents example.")
		
		// Create a group name
		groupName := foundation.StringWithString("DarwinKit Test Group")
		
		// Create a conversation identifier
		conversationID := foundation.StringWithString("darwinkit-convo-123")
		
		// Create a send message intent
		sendMessageIntent := intents.NewSendMessageIntent(peopleArray, messageContent, groupName, conversationID)
		
		// Get the content of the message
		fmt.Println("Created SendMessageIntent with content:", sendMessageIntent.Content().String())
		
		// Create a response for the intent
		response := intents.NewIntentResponse(intents.INIntentResponseCodeSuccess)
		
		// Create an interaction with the intent and response
		interaction := intents.NewInteraction(sendMessageIntent, response)
		
		// Show the direction of the interaction
		direction := interaction.Direction()
		directionText := "Unknown"
		switch direction {
		case intents.INInteractionDirectionOutgoing:
			directionText = "Outgoing"
		case intents.INInteractionDirectionIncoming:
			directionText = "Incoming"
		}
		fmt.Println("Interaction direction:", directionText)
		
		// Now create a search messages intent
		// Create search terms
		searchTerms := foundation.StringWithString("meeting")
		
		// Create a date range for the search (last 7 days)
		now := foundation.DateWithTimeIntervalSinceNow(0)
		sevenDaysAgo := foundation.DateWithTimeIntervalSinceNow(-7 * 24 * 60 * 60)
		dateRange := foundation.DateIntervalWithStartDateEndDate(sevenDaysAgo, now)
		
		// Create empty arrays for the parameters we're not using
		emptyArray := foundation.ArrayWithCapacity(0)
		
		// Create the search intent
		searchIntent := intents.NewSearchForMessagesIntent(
			emptyArray, // recipients
			peopleArray, // senders
			searchTerms,
			emptyArray, // attributes
			dateRange,
			emptyArray, // identifiers
		)
		
		// Show the search terms
		fmt.Println("Created SearchForMessagesIntent with search terms:", searchIntent.SearchTerms().String())
		
		// Create a new intent definition
		intentDefinition := intents.NewIntentDefinition()
		intentDefinition.SetIdentifier(foundation.StringWithString("com.example.SendMessage"))
		intentDefinition.SetCategoryName(foundation.StringWithString("messaging"))
		
		// Get vocabulary instance
		vocabulary := intents.SharedVocabulary()
		fmt.Println("Got shared vocabulary instance")
		
		// Example of intent error codes
		fmt.Println("\nIntent Error Codes:")
		fmt.Println("- Request Timed Out:", intents.INIntentErrorCodeRequestTimedOut)
		fmt.Println("- Network Unavailable:", intents.INIntentErrorCodeNetworkUnavailable)
		fmt.Println("- No App Available:", intents.INIntentErrorCodeNoAppAvailable)
		
		// Example of intent response codes
		fmt.Println("\nIntent Response Codes:")
		fmt.Println("- Success:", intents.INIntentResponseCodeSuccess)
		fmt.Println("- Failure:", intents.INIntentResponseCodeFailure)
		fmt.Println("- In Progress:", intents.INIntentResponseCodeInProgress)
		
		// Example of speech recognition results
		recognitionResult := intents.INSpeechRecognitionResult{
			TranscriptionString: "Call John at work",
			Confidence:          0.95,
		}
		fmt.Printf("\nSpeech Recognition Result: %s (Confidence: %.2f)\n", 
			recognitionResult.TranscriptionString, 
			recognitionResult.Confidence)
		
		// Wait a moment before exiting
		time.Sleep(1 * time.Second)
		fmt.Println("\nExample completed.")
	})
}