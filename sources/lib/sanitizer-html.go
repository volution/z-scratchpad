

package zscratchpad


import "bytes"
// import "html"

import "github.com/microcosm-cc/bluemonday"



func DocumentSanitizeHtml (_document *Document, _unsafe string) (string, *Error) {
	
	_parser := bluemonday.UGCPolicy()
	
	_parser.RequireParseableURLs (true)
	_parser.RequireNoFollowOnLinks (true)
	_parser.RequireNoReferrerOnLinks (true)
	_parser.RequireCrossOriginAnonymous (true)
	
	_unsafeBuffer := bytes.NewBufferString (_unsafe)
	
	_sanitizedBuffer := _parser.SanitizeReader (_unsafeBuffer)
	
	_sanitized := _sanitizedBuffer.String ()
	
	return _sanitized, nil
}

