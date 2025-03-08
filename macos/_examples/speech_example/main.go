package main

import (
	"fmt"
	"time"

	"github.com/progrium/darwinkit/macos/appkit"
	"github.com/progrium/darwinkit/macos/avfoundation"
	"github.com/progrium/darwinkit/macos/foundation"
	"github.com/progrium/darwinkit/macos/speech"
	"github.com/progrium/darwinkit/objc"
)

func main() {
	objc.WithAutoreleasePool(func() {
		// Create the application
		app := appkit.SharedApplication()
		
		fmt.Println("Speech Framework Example")
		fmt.Println("-----------------------")
		
		// Check authorization status
		authStatus := speech.AuthorizationStatus()
		var authStatusStr string
		
		switch authStatus {
		case speech.SFSpeechRecognizerAuthorizationStatusAuthorized:
			authStatusStr = "Authorized"
		case speech.SFSpeechRecognizerAuthorizationStatusDenied:
			authStatusStr = "Denied"
		case speech.SFSpeechRecognizerAuthorizationStatusRestricted:
			authStatusStr = "Restricted"
		case speech.SFSpeechRecognizerAuthorizationStatusNotDetermined:
			authStatusStr = "Not Determined"
		}
		
		fmt.Println("Speech recognition authorization status:", authStatusStr)
		
		// Request authorization (this will show a prompt if needed)
		fmt.Println("Requesting speech recognition authorization...")
		
		// Create an authorization handler
		authHandler := objc.NewBlock(func(status speech.SFSpeechRecognizerAuthorizationStatus) {
			fmt.Println("Authorization handler called with status:", int(status))
		})
		
		// Request authorization
		speech.RequestAuthorization(authHandler)
		
		// Check for supported locales
		supportedLocales := speech.SupportedLocales()
		if !supportedLocales.IsNil() {
			count := supportedLocales.Count()
			fmt.Printf("There are %d supported locales for speech recognition\n", count)
			
			// Get the current language code
			currentLang := speech.CurrentLanguageCode()
			if !currentLang.IsNil() {
				fmt.Println("Current language code:", currentLang.String())
			}
		}
		
		// Create a speech recognizer (without specifying locale to use default)
		recognizer := speech.NewSpeechRecognizer()
		
		if !recognizer.IsNil() {
			available := recognizer.IsAvailable()
			fmt.Println("Speech recognizer available:", available)
			
			// Create a URL recognition request (simulated, file doesn't exist)
			audioFileURL := foundation.URLWithString(foundation.StringWithString("file:///path/to/audio/file.m4a"))
			urlRequest := speech.NewSpeechURLRecognitionRequest(audioFileURL)
			
			// Configure the request
			urlRequest.SetShouldReportPartialResults(true)
			urlRequest.SetRequiresOnDeviceRecognition(false)
			
			fmt.Println("Created URL recognition request")
			fmt.Println("- Should report partial results:", urlRequest.ShouldReportPartialResults())
			fmt.Println("- Requires on-device recognition:", urlRequest.RequiresOnDeviceRecognition())
			
			// Create an audio buffer recognition request
			bufferRequest := speech.NewSpeechAudioBufferRecognitionRequest()
			
			// Configure the request
			bufferRequest.SetShouldReportPartialResults(true)
			
			fmt.Println("Created audio buffer recognition request")
			
			// Demonstrate task states without actually creating a task
			fmt.Println("\nSpeech recognition task states:")
			fmt.Println("- Starting:", speech.SFSpeechRecognitionTaskStateStarting)
			fmt.Println("- Running:", speech.SFSpeechRecognitionTaskStateRunning)
			fmt.Println("- Finishing:", speech.SFSpeechRecognitionTaskStateFinishing)
			fmt.Println("- Completed:", speech.SFSpeechRecognitionTaskStateCompleted)
			fmt.Println("- Canceling:", speech.SFSpeechRecognitionTaskStateCanceling)
			fmt.Println("- Cancelled:", speech.SFSpeechRecognitionTaskStateCancelled)
			
			// Demonstrate a speech recognition result
			recognitionResult := speech.SFSpeechRecognitionResult{
				BestTranscription: "This is a sample transcription",
				IsFinal: true,
				Confidence: 0.92,
			}
			
			fmt.Println("\nSample speech recognition result:")
			fmt.Printf("- Transcription: %s\n", recognitionResult.BestTranscription)
			fmt.Printf("- Is Final: %t\n", recognitionResult.IsFinal)
			fmt.Printf("- Confidence: %.2f\n", recognitionResult.Confidence)
			
			// Demonstrate voice analysis features
			fmt.Println("\nVoice analysis features:")
			fmt.Println("- Voicing:", speech.SFAcousticFeatureVoicing)
			fmt.Println("- Voicelessness:", speech.SFAcousticFeatureVoicelessness)
			fmt.Println("- Pitch:", speech.SFAcousticFeaturePitch)
			fmt.Println("- Jitter:", speech.SFAcousticFeatureJitter)
			fmt.Println("- Shimmer:", speech.SFAcousticFeatureShimmer)
		} else {
			fmt.Println("Speech recognizer could not be created")
		}
		
		// Wait a moment before exiting
		time.Sleep(1 * time.Second)
		fmt.Println("\nExample completed.")
	})
}