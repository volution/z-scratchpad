

package zscratchpad


import "unicode/utf8"

import "github.com/volution/z-scratchpad/extensions/gemini"




func parseAndRenderGeminiToHtml (_sourceLines []string) (string, *Error) {
	
	_outputBuffer := BytesBufferNewSize (128 * 1024)
	defer BytesBufferRelease (_outputBuffer)
	
	_state := gemini.NewHtmlState (_outputBuffer)
	_render := func (_line gemini.Line) () {
			_line.Html (_state)
		}
	
	_error := gemini.ParseLines (_sourceLines, _render)
	if _error != nil {
		return "", errorw (0x378cf702, _error)
	}
	
	_state.Flush ()
	
	_outputBytes := _outputBuffer.Bytes ()
	if ! utf8.Valid (_outputBytes) {
		return "", errorw (0xb67ee18b, nil)
	}
	
	_output := string (_outputBytes)
	
	return _output, nil
}

