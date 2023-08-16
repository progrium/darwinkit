// AUTO-GENERATED CODE, DO NOT MODIFY

package corespotlight

// Error codes that describe indexing-specific errors. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/csindexerrorcode?language=objc
type IndexErrorCode int

const (
	IndexErrorCodeIndexUnavailableError   IndexErrorCode = -1000
	IndexErrorCodeIndexingUnsupported     IndexErrorCode = -1005
	IndexErrorCodeInvalidClientStateError IndexErrorCode = -1002
	IndexErrorCodeInvalidItemError        IndexErrorCode = -1001
	IndexErrorCodeQuotaExceeded           IndexErrorCode = -1004
	IndexErrorCodeRemoteConnectionError   IndexErrorCode = -1003
	IndexErrorCodeUnknownError            IndexErrorCode = -1
)

// Error codes that describe reasons a query might fail. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchqueryerrorcode?language=objc
type SearchQueryErrorCode int

const (
	SearchQueryErrorCodeCancelled        SearchQueryErrorCode = -2003
	SearchQueryErrorCodeIndexUnreachable SearchQueryErrorCode = -2001
	SearchQueryErrorCodeInvalidQuery     SearchQueryErrorCode = -2002
	SearchQueryErrorCodeUnknown          SearchQueryErrorCode = -2000
)
