

package zscratchpad


import "strings"
import "unicode"
import "path"




func stringSplitLine (_input string) (string, string, bool) {
	_splitIndex := strings.IndexByte (_input, '\n')
	if _splitIndex == -1 {
		return "", _input, false
	}
	_line := _input[:_splitIndex]
	_rest := _input[_splitIndex + 1:]
	if (_line != "") && (_line[len (_line) - 1] == '\r') {
		_line = _line[: len (_line) - 1]
	}
	if (_rest != "") && (_rest[0] == '\r') {
		_rest = _rest[1:]
	}
	return _line, _rest, true
}


func stringSplitLines (_input string) ([]string, bool) {
	_lines := make ([]string, 0, 128)
	for {
		if _input == "" {
			break
		}
		_line, _rest, _ok := stringSplitLine (_input)
		if !_ok {
			_lines = append (_lines, _rest)
			return _lines, false
		} else {
			_lines = append (_lines, _line)
			_input = _rest
		}
	}
	return _lines, true
}




func stringTrimSpaces (_input string) (string) {
	return strings.TrimFunc (_input, unicode.IsSpace)
}

func stringTrimSpacesLeft (_input string) (string) {
	return strings.TrimLeftFunc (_input, unicode.IsSpace)
}

func stringTrimSpacesRight (_input string) (string) {
	return strings.TrimRightFunc (_input, unicode.IsSpace)
}




func pathSplitFileNameAndExtension (_path string) (string, string, *Error) {
	if _path == "" {
		return "", "", errorf (0xd2d47410, "path empty")
	}
	_base := path.Base (_path)
	if _base == "." {
		return "", "", errorf (0x640581f3, "path empty")
	}
	if _base == "/" {
		return "", "", errorf (0xb01bd167, "path empty")
	}
	if _base[len (_base) - 1] == '.' {
		return "", "", errorf (0x9f585e8b, "dot suffix")
	}
	_extension := path.Ext (_base)
	if _extension == "" {
		return _base, "", nil
	} else {
		_base = _base[: len (_base) - len (_extension)]
		_extension = _extension[1:]
		return _base, _extension, nil
	}
}

