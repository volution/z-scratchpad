

package zscratchpad


import "bytes"
import "html"
import "strings"
import "unicode"




type SnippetBlock interface {
	RenderHtmlInto (*bytes.Buffer) (*Error)
}

type SnippetTextBlock struct {
	Lines []string
}

type SnippetBreakBlock struct {
	Lines uint
}

type SnippetEmptyBlock struct {
	Lines uint
}




func (_block *SnippetTextBlock) RenderHtmlInto (_buffer *bytes.Buffer) (*Error) {
	
	_buffer.WriteString ("<pre>\n")
	for _, _line := range _block.Lines {
		_line = html.EscapeString (_line)
		_buffer.WriteString (_line)
		_buffer.WriteString ("\n")
	}
	_buffer.WriteString ("</pre>\n")
	
	return nil
}


func (_block *SnippetBreakBlock) RenderHtmlInto (_buffer *bytes.Buffer) (*Error) {
	
	for _lines := _block.Lines; _lines >= 1; _lines -= 1 {
		_buffer.WriteString ("<hr/>\n")
	}
	
	return nil
}


func (_block *SnippetEmptyBlock) RenderHtmlInto (_buffer *bytes.Buffer) (*Error) {
	
	_lines := 0
	if _block.Lines <= 3 {
		_lines = 0
	} else if _block.Lines <= 4 {
		_lines = 1
	} else {
		_lines = 2
	}
	
	for _lines := _lines; _lines >= 1; _lines -= 1 {
		_buffer.WriteString ("<hr/>\n")
	}
	
	return nil
}




func parseAndRenderSnippetsToHtml (_source string) (string, *Error) {
	
	var _blocks []SnippetBlock
	if _blocks_0, _error := parseSnippets (_source); _error == nil {
		_blocks = _blocks_0
	} else {
		return "", _error
	}
	
	_buffer := bytes.NewBuffer (nil)
	
	for _, _block := range _blocks {
		if _error := _block.RenderHtmlInto (_buffer); _error != nil {
			return "", _error
		}
	}
	
	_output := string (_buffer.Bytes ())
	
	return _output, nil
}



func parseSnippets (_source string) ([]SnippetBlock, *Error) {
	
	_lines, _ := stringSplitLines (_source)
	
	_blocks := []SnippetBlock (nil)
	
	_blockLines := []string (nil)
	_breakLines := 0
	_emptyLines := 0
	_atStart := true
	
	_blockLinesPush := func () () {
		if len (_blockLines) == 0 {
			return
		}
		_block := & SnippetTextBlock {
				Lines : _blockLines,
			}
		_blockLines = nil
		_blocks = append (_blocks, _block)
	}
	
	for _, _line := range _lines {
		
		_lineTrimmed := stringTrimSpaces (_line)
		
		_isBreakLine := false
		_isEmptyLine := false
		
		for _, _breakPrefix := range []string {"########", "========", "--------", "%%%%%%%%"} {
			if strings.HasPrefix (_lineTrimmed, _breakPrefix) {
				_breakRune := rune (_breakPrefix[0])
				if strings.IndexFunc (_lineTrimmed, func (_rune rune) (bool) { return (_rune != _breakRune) && ! unicode.IsSpace (_rune) }) == -1 {
					_isBreakLine = true
					break
				}
			}
		}
		if _lineTrimmed == "" {
			_isEmptyLine = true
		}
		
//		logf ('d', 0xfd669a22, "%d %v/%d %v/%d -- `%s`", len (_blockLines), _isBreakLine, _breakLines, _isEmptyLine, _emptyLines, _line)
		
		if _isBreakLine || _isEmptyLine {
			if _atStart {
				continue
			}
			if _isBreakLine {
				_blockLinesPush ()
				_breakLines += 1
				_emptyLines = 0
			}
			if _isEmptyLine {
				if (_breakLines == 0) || (len (_blockLines) > 0) {
					_emptyLines += 1
				}
			}
			continue
		}
		
		if _breakLines > 0 {
			_block := & SnippetBreakBlock {
					Lines : uint (_breakLines),
				}
			_blocks = append (_blocks, _block)
			_breakLines = 0
			_emptyLines = 0
		}
		
		if _emptyLines > 0 {
			if _emptyLines == 1 {
				_blockLines = append (_blockLines, "")
			} else {
				_blockLinesPush ()
				_block := & SnippetEmptyBlock {
						Lines : uint (_emptyLines),
					}
				_blocks = append (_blocks, _block)
			}
			_emptyLines = 0
		}
		
		_blockLines = append (_blockLines, _line)
		
		_atStart = false
	}
	
	_blockLinesPush ()
	
	return _blocks, nil
}

