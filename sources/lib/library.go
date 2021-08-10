

package zscratchpad


import "errors"
import "io/fs"
import "os"
import "path/filepath"
import "regexp"
import "strings"


import "github.com/gobwas/glob"




type Library struct {
	
	Identifier string
	Name string
	
	Paths []string
	
	EditEnabled bool
	
	CreateEnabled bool
	CreatePath string
	CreateExtension string
	
	SnapshotEnabled bool
	SnapshotExtension string
	
	IncludeGlobPatterns []string
	ExcludeGlobPatterns []string
	
	IncludeRegexPatterns []string
	ExcludeRegexPatterns []string
	
	UseFileNameAsIdentifier bool
	UseFileExtensionAsFormat bool
	
	includeGlobMatchers []glob.Glob
	excludeGlobMatchers []glob.Glob
	
	includeRegexMatchers []*regexp.Regexp
	excludeRegexMatchers []*regexp.Regexp
}




func LibraryInitialize (_library *Library) (*Error) {
	
	if _library.Identifier == "" {
		return errorw (0x94465013, nil)
	}
	
	for _index, _path := range _library.Paths {
		if _path_0, _error := filepath.Abs (_path); _error == nil {
			_path = _path_0
		} else {
			return errorw (0xe0ece239, _error)
		}
		if _stat, _error := os.Stat (_path); _error == nil {
			if ! _stat.IsDir () {
				return errorw (0x410a4abd, nil)
			}
		} else {
			return errorw (0x1513652d, _error)
		}
		_library.Paths[_index] = _path
	}
	
	if _library.CreateEnabled {
		if _library.CreatePath == "" {
			if len (_library.Paths) == 1 {
				_library.CreatePath = _library.Paths[0]
			} else {
				return errorw (0xd76cef62, nil)
			}
		}
		if _path_0, _error := filepath.Abs (_library.CreatePath); _error == nil {
			_library.CreatePath = _path_0
		} else {
			return errorw (0xea573d9c, _error)
		}
		if _stat, _error := os.Stat (_library.CreatePath); _error == nil {
			if ! _stat.IsDir () {
				return errorw (0x1ad922a3, nil)
			}
		} else {
			return errorw (0x98ade3fc, _error)
		}
		if _library.CreateExtension == "" {
			_library.CreateExtension = "txt"
		}
	} else {
		if _library.CreatePath != "" {
			return errorw (0x5b55e852, nil)
		}
		if _library.CreateExtension != "" {
			return errorw (0x2ffc3bf4, nil)
		}
	}
	
	if _library.SnapshotEnabled {
		if _library.SnapshotExtension == "" {
			_library.SnapshotExtension = "snapshot"
		}
	} else {
		if _library.SnapshotExtension != "" {
			return errorw (0x3ede0dc5, nil)
		}
	}
	
	_library.includeGlobMatchers = make ([]glob.Glob, 0, len (_library.IncludeGlobPatterns))
	for _, _pattern := range _library.IncludeGlobPatterns {
		if _matcher, _error := glob.Compile (_pattern); _error == nil {
			_library.includeGlobMatchers = append (_library.includeGlobMatchers, _matcher)
		} else {
			return errorw (0x674d8ba9, _error)
		}
	}
	
	_library.excludeGlobMatchers = make ([]glob.Glob, 0, len (_library.IncludeGlobPatterns))
	for _, _pattern := range _library.IncludeGlobPatterns {
		if _matcher, _error := glob.Compile (_pattern); _error == nil {
			_library.excludeGlobMatchers = append (_library.excludeGlobMatchers, _matcher)
		} else {
			return errorw (0x5d547147, _error)
		}
	}
	
	_library.includeRegexMatchers = make ([]*regexp.Regexp, 0, len (_library.IncludeRegexPatterns))
	for _, _pattern := range _library.IncludeRegexPatterns {
		if _matcher, _error := regexp.Compile (_pattern); _error == nil {
			_library.includeRegexMatchers = append (_library.includeRegexMatchers, _matcher)
		} else {
			return errorw (0x3515908f, _error)
		}
	}
	
	_library.excludeRegexMatchers = make ([]*regexp.Regexp, 0, len (_library.IncludeRegexPatterns))
	for _, _pattern := range _library.IncludeRegexPatterns {
		if _matcher, _error := regexp.Compile (_pattern); _error == nil {
			_library.excludeRegexMatchers = append (_library.excludeRegexMatchers, _matcher)
		} else {
			return errorw (0xe3938785, _error)
		}
	}
	
	return nil
}




func libraryDocumentsLoad (_library *Library, _documentPaths []string) ([]*Document, *Error) {
	
	_documents := make ([]*Document, 0, len (_documentPaths))
	
	for _, _documentPath := range _documentPaths {
		if _document, _error := DocumentLoadFromPath (_documentPath); _error == nil {
			_documents = append (_documents, _document)
		} else {
			return nil, _error
		}
	}
	
	return _documents, nil
}




func libraryDocumentsWalk (_library *Library) ([]string, *Error) {
	
	_documentPaths := []string (nil)
	for _, _libraryPath := range _library.Paths {
		if _documentPaths_0, _error := libraryDocumentsWalkPath (_library, _libraryPath); _error == nil {
			if _documentPaths == nil {
				_documentPaths = _documentPaths_0
			} else {
				_documentPaths = append (_documentPaths, _documentPaths_0 ...)
			}
		} else {
			return nil, _error
		}
	}
	
	return _documentPaths, nil
}


func libraryDocumentsWalkPath (_library *Library, _libraryPath string) ([]string, *Error) {
	
	if _libraryPath == "" {
		return nil, errorw (0x83afc399, nil)
	}
	
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
			_documentPath := filepath.Join (_libraryPath, _path)
			_documentPaths = append (_documentPaths, _documentPath)
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

