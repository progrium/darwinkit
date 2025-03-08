package speech

import (
	"testing"
)

// TestSpeechValid tests that the Speech framework is properly imported
func TestSpeechValid(t *testing.T) {
	// Verify that Speech enums are properly defined
	if SFSpeechRecognizerAuthorizationStatusAuthorized != 3 {
		t.Errorf("Expected SFSpeechRecognizerAuthorizationStatusAuthorized to be 3, got %d", SFSpeechRecognizerAuthorizationStatusAuthorized)
	}

	if SFSpeechRecognitionTaskStateCompleted != 3 {
		t.Errorf("Expected SFSpeechRecognitionTaskStateCompleted to be 3, got %d", SFSpeechRecognitionTaskStateCompleted)
	}
}