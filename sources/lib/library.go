

package zscratchpad


import "errors"
import "io/fs"
import "path/filepath"
import "regexp"
import "strings"




type Library struct {
	
	Identifier string
	Name string
	Path string
	
	UseFileNameAsIdentifier bool
	UseFileExtensionAsFormat bool
}




func libraryDocumentsLoad (_libraryPath string, _documentPaths []string) ([]*Document, *Error) {
	
	_documents := make ([]*Document, 0, len (_documentPaths))
	
	for _, _documentPath := range _documentPaths {
		_path := filepath.Join (_libraryPath, _documentPath)
		if _document, _error := DocumentLoadFromPath (_path); _error == nil {
			_documents = append (_documents, _document)
		} else {
			return nil, _error
		}
	}
	
	return _documents, nil
}




func libraryDocumentsWalk (_libraryPath string) ([]string, *Error) {
	
	_error_1 := (*Error) (nil)
	
	_documentPaths := make ([]string, 0, 1024)
	
	_walkQuit := errors.New ("")
	_walk := func (_path string, _entry fs.DirEntry, _error error) (error) {
		if _error != nil {
			_error_1 = errorw (0x534f844f, _error)
			return _walkQuit
		}
		if _path_0, _error_0 := filepath.Rel (_libraryPath, _path); _error_0 == nil {
			_path = _path_0
		} else {
			_error_1 = errorw (0xacc84f2b, _error_0)
			return _walkQuit
		}
		_name := _entry.Name ()
		_mode := _entry.Type ()
		if strings.HasPrefix (_name, ".") {
			if _mode.IsDir () {
				return filepath.SkipDir
			} else {
				return nil
			}
		}
		if _mode.IsRegular () {
			_documentPaths = append (_documentPaths, _path)
		} else if _mode.IsDir () {
			// NOP
		} else {
			_error_1 = errorf (0xb0cc4319, "invalid entry `%s`", _path)
			return _walkQuit
		}
		return nil
	}
	
	if _error_2 := filepath.WalkDir (_libraryPath, _walk); (_error_2 != nil) && (_error_2 != _walkQuit) {
		return nil, errorw (0xdc8ea9dd, _error_2)
	}
	if _error_1 != nil {
		return nil, _error_1
	}
	
	return _documentPaths, nil
}




func LibraryValidateIdentifier (_identifier string) (*Error) {
	if ! LibraryIdentifierRegex.MatchString (_identifier) {
		return errorw (0x2d8a1040, nil)
	}
	return nil
}

func LibraryParseIdentifier (_identifier string) (string, *Error) {
	if _error := LibraryValidateIdentifier (_identifier); _error != nil {
		return "", _error
	}
	return _identifier, nil
}

var LibraryIdentifierRegexToken string = `(?:(?:[a-z0-9]+)(?:[_-]+[a-z0-9]+)*)`
var LibraryIdentifierRegex *regexp.Regexp = regexp.MustCompile (`^` + LibraryIdentifierRegexToken + `$`)

