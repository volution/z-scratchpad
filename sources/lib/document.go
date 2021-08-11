

package zscratchpad


import "bytes"
import "fmt"
import "io"
import "os"
import "path"
import "regexp"
import "sort"
import "strings"
import "time"
import "unicode/utf8"




type Document struct {
	
	Identifier string
	Library string
	Path string
	PathInLibrary string
	
	Title string
	TitleAlternatives []string
	
	SourceFingerprint string
	
	Format string
	
	BodyLines []string
	BodyEmpty bool
	BodyFingerprint string
	
	RenderHtml string
	RenderText string
	
	EditEnabled bool
	Timestamp time.Time
}




func DocumentInitializeIdentifier (_document *Document, _library *Library) (*Error) {
	if (_library != nil) && (_document.Library != _library.Identifier) {
		return errorw (0x767046ec, nil)
	}
	_useLibraryPrefix := false
	_usePathInLibrary := false
	_useFileName := false
	_usePathFingerprint := true
	if _library != nil {
		_useLibraryPrefix = _library.UseLibraryAsIdentifierPrefix
		_usePathInLibrary = _library.UsePathInLibraryAsIdentifier
		_useFileName = _library.UseFileNameAsIdentifier
		_usePathFingerprint = _library.UsePathFingerprintAsIdentifier
	}
	return DocumentInitializeIdentifier_0 (_document, _useLibraryPrefix, _usePathInLibrary, _useFileName, _usePathFingerprint)
}


func DocumentInitializeIdentifier_0 (_document *Document, _useLibraryPrefix bool, _usePathInLibrary bool, _useFileName bool, _usePathFingerprint bool) (*Error) {
	
	_libraryIdentifier := ""
	if (_document.Library != "") && _useLibraryPrefix {
		_libraryIdentifier = _document.Library
	}
	
	_documentName := ""
	
	if (_document.Identifier != "") {
		_documentName = _document.Identifier
		goto _resolve
	}
	
	if (_document.PathInLibrary != "") && _usePathInLibrary {
		_folderPath, _fileName := path.Split (_document.PathInLibrary)
		if _documentName_0, _, _error := pathSplitFileNameAndExtension (_fileName); _error == nil {
			if _folderPath != "" {
				_folderPath = _folderPath[: len(_folderPath) - 1]
			}
			if _folderPath != "" {
				_documentName = strings.ReplaceAll (_folderPath, "/", "~~") + "~~" + _documentName_0
			} else {
				_documentName = _documentName_0
			}
			goto _resolve
		} else {
			return _error
		}
	}
	
	if (_document.Path != "") && _useFileName {
		if _documentName_0, _, _error := pathSplitFileNameAndExtension (_document.Path); _error == nil {
			_documentName = _documentName_0
			goto _resolve
		} else {
			return _error
		}
	}
	
	if (_document.Path != "") && _usePathFingerprint {
		_fingerprint := fingerprintString (_document.Path)
		_documentName = _fingerprint[:32]
		goto _resolve
	}
	
//	logf ('d', 0xadfa2993, "%s", _document.Path)
	
	return errorf (0x1c58da80, "identifier unresolvable")
	
	_resolve :
	
	if _identifier, _error := DocumentFormatIdentifier (_libraryIdentifier, _documentName); _error == nil {
		_document.Identifier = _identifier
		return nil
	} else {
		return _error
	}
}




func DocumentInitializeFormat (_document *Document, _library *Library) (*Error) {
	if (_library != nil) && (_document.Library != _library.Identifier) {
		return errorw (0xc8a19353, nil)
	}
	_useFileExtension := true
	if _library != nil {
		_useFileExtension = _library.UseFileExtensionAsFormat
	}
	return DocumentInitializeFormat_0 (_document, _useFileExtension)
}


func DocumentInitializeFormat_0 (_document *Document, _useFileExtension bool) (*Error) {
	
	if _document.Format != "" {
		return nil
	}
	
	if (_document.Path != "") && _useFileExtension {
		if _, _extension, _error := pathSplitFileNameAndExtension (_document.Path); _error == nil {
			_format := ""
			switch _extension {
				case "md" :
					_format = "commonmark"
				case "txt" :
					_format = "text"
			}
			if _format != "" {
				_document.Format = _format
				return nil
			}
		} else {
			return _error
		}
	}
	
//	logf ('d', 0xff65fe47, "%s", _document.Path)
	
	return errorf (0xe5e1dd0f, "format unresolvable")
}




func DocumentValidateIdentifier (_identifier string) (*Error) {
	if ! DocumentIdentifierRegex.MatchString (_identifier) {
		return errorw (0x55874ebf, nil)
	}
	return nil
}

func DocumentParseIdentifier (_identifier string) (string, string, string, *Error) {
	if _error := DocumentValidateIdentifier (_identifier); _error != nil {
		return "", "", "", _error
	}
	if _splitIndex := strings.IndexByte (_identifier, ':'); _splitIndex != -1 {
		_libraryIdentifier := _identifier[:_splitIndex]
		_documentName := _identifier[_splitIndex + 1:]
		return _identifier, _libraryIdentifier, _documentName, nil
	} else {
		return _identifier, "", "", nil
	}
}

func DocumentFormatIdentifier (_libraryIdentifier string, _documentName string) (string, *Error) {
	if ! DocumentIdentifierWithoutLibraryRegex.MatchString (_documentName) {
		return "", errorw (0x9f777d70, nil)
	}
	if _libraryIdentifier != "" {
		if ! LibraryIdentifierRegex.MatchString (_libraryIdentifier) {
			return "", errorw (0xfc88cf9f, nil)
		}
		_identifier := _libraryIdentifier + ":" + _documentName
		return _identifier, nil
	} else {
		return _documentName, nil
	}
}


var DocumentIdentifierWithoutLibraryRegexToken string = `(?:(?:[a-z0-9]+)(?:(?:[_-]{1,2}|~~)[a-z0-9]+)*)`
var DocumentIdentifierWithoutLibraryRegex *regexp.Regexp = regexp.MustCompile (`^` + DocumentIdentifierWithoutLibraryRegexToken + `$`)
var DocumentIdentifierWithLibraryRegexToken string = `(?:` + LibraryIdentifierRegexToken + `:` + DocumentIdentifierWithoutLibraryRegexToken + `)`
var DocumentIdentifierRegexToken string = `(?:` + DocumentIdentifierWithoutLibraryRegexToken + `|` + DocumentIdentifierWithLibraryRegexToken + `)`
var DocumentIdentifierRegex *regexp.Regexp = regexp.MustCompile (`^` + DocumentIdentifierRegexToken + `$`)




func DocumentLoadFromPath (_path string) (*Document, *Error) {
	
	var _file *os.File
	if _file_0, _error := os.OpenFile (_path, os.O_RDONLY, 0); _error == nil {
		_file = _file_0
	} else {
		return nil, errorw (0xc1e080d9, _error)
	}
	
	var _stat os.FileInfo
	if _stat_0, _error := _file.Stat (); _error == nil {
		_stat = _stat_0
	} else {
		return nil, errorw (0xe18c3be5, _error)
	}
	
	_timestamp := _stat.ModTime ()
	
	var _sourceBytes []byte
	if _bytes, _error := os.ReadFile (_path); _error == nil {
		_sourceBytes = _bytes
	} else {
		return nil, errorw (0x483c6b27, _error)
	}
	
	if ! utf8.Valid (_sourceBytes) {
//		logf ('d', 0x742720c2, "%s", _path)
		return nil, errorf (0xa24965ce, "invalid UTF-8 source")
	}
	_source := string (_sourceBytes)
	
	var _document *Document
	if _document_0, _error := DocumentLoadFromBuffer (_source); _error == nil {
		_document = _document_0
	} else {
		return nil, _error
	}
	
	_document.Path = _path
	_document.Timestamp = _timestamp
	
	return _document, nil
}



func DocumentLoadFromBuffer (_source string) (*Document, *Error) {
	
	var _identifier string
	var _library string
	var _format string
	var _title string
	var _titles []string
	
	_body := _source
	for {
		
		if _body == "" {
			break
		}
		_header, _rest, _ok := stringSplitLine (_body)
		if !_ok {
			break
		}
		
		if _header == "" {
			_body = _rest
			break
		}
		
		if ! strings.HasPrefix (_header, "## ") {
			break
		}
		
		_body = _rest
		
		_header = _header[3:]
		_header = stringTrimSpaces (_header)
		
		if _header == "" {
			return nil, errorf (0x8d4a068d, "header empty")
		}
		
		if strings.HasPrefix (_header, "-- ") {
			
			_header = _header[3:]
			_header = stringTrimSpaces (_header)
			
			if strings.HasPrefix (_header, "identifier:") {
				_identifier_0 := _header[11:]
				_identifier_0 = stringTrimSpaces (_identifier_0)
				_identifier = _identifier_0
			} else if strings.HasPrefix (_header, "library:") {
				_library_0 := _header[8:]
				_library_0 = stringTrimSpaces (_library_0)
				_library = _library_0
			} else if strings.HasPrefix (_header, "format:") {
				_format_0 := _header[7:]
				_format_0 = stringTrimSpaces (_format_0)
				_format = _format_0
			} else if strings.HasPrefix (_header, "timestamp:") {
				// NOTE:  Ignore timestamps from file.
			} else {
				return nil, errorf (0xc5ccdc9e, "metadata invalid `%s`", _header)
			}
			
		} else {
			
			_titles = append (_titles, _header)
			if _title == "" {
				_title = _header
			}
		}
	}
	
	_bodyLines_0, _ := stringSplitLines (_body)
	_bodyLines := make ([]string, 0, len (_bodyLines_0))
	_bodyLinesEmpty := 0
	_bodyEmpty := true
	for _, _line := range _bodyLines_0 {
		_line = stringTrimSpacesRight (_line)
		if _line != "" {
			_bodyEmpty = false
			_bodyLinesEmpty = 0
		} else {
			if len (_bodyLines) == 0 {
				continue
			} else {
				_bodyLinesEmpty += 1
			}
		}
		_bodyLines = append (_bodyLines, _line)
	}
	if _bodyLinesEmpty > 0 {
		_bodyLines = _bodyLines[: len (_bodyLines) - _bodyLinesEmpty]
	}
	
	
	if _identifier != "" {
		if ! DocumentIdentifierWithoutLibraryRegex.MatchString (_identifier) {
			return nil, errorw (0x31e50aa1, nil)
		}
	}
	if _library != "" {
		if _error := LibraryValidateIdentifier (_identifier); _error != nil {
			return nil, _error
		}
	}
	if _format != "" {
		switch _format {
			case "commonmark", "snippets", "text" :
				// NOP
			case "markdown" :
				_format = "commonmark"
			default :
				return nil, errorf (0x32158fbf, "format invalid `%s`", _format)
		}
	}
	
	sort.Strings (_titles)
	
	_sourceFingerprint := fingerprintString (_source)
	_bodyFingerprint := fingerprintStringLines (_bodyLines)
	
	_document := & Document {
			Title : _title,
			TitleAlternatives : _titles,
			Identifier : _identifier,
			Library : _library,
			Format : _format,
			SourceFingerprint : _sourceFingerprint,
			BodyLines : _bodyLines,
			BodyEmpty : _bodyEmpty,
			BodyFingerprint : _bodyFingerprint,
		}
	
	return _document, nil
}




func DocumentDump (_stream io.Writer, _document *Document, _includeIdentifiers bool, _includeBody bool, _includeRender bool) (*Error) {
	
	_buffer := bytes.NewBuffer (nil)
	
	if _document.Title != "" {
		fmt.Fprintf (_buffer, "-- title (primary): `%s`\n", _document.Title)
	}
	for _, _title := range _document.TitleAlternatives {
		if _title == _document.Title {
			continue
		}
		fmt.Fprintf (_buffer, "-- title (alternative): `%s`\n", _title)
	}
	
	if _includeIdentifiers {
		if _document.Identifier != "" {
			fmt.Fprintf (_buffer, "-- identifier: `%s`\n", _document.Identifier)
		}
		if _document.Library != "" {
			fmt.Fprintf (_buffer, "-- library: `%s`\n", _document.Library)
		}
		if _document.Format != "" {
			fmt.Fprintf (_buffer, "-- format: `%s`\n", _document.Format)
		}
		if _document.Path != "" {
			fmt.Fprintf (_buffer, "-- path: `%s`\n", _document.Path)
		}
		if ! _document.Timestamp.IsZero () {
			fmt.Fprintf (_buffer, "-- timestamp: `%s`\n", _document.Timestamp.Format ("2006-01-02 15:04:05"))
		}
		if _document.SourceFingerprint != "" {
			fmt.Fprintf (_buffer, "-- source fingerprint: `%s`\n", _document.SourceFingerprint)
		}
		if _document.BodyFingerprint != "" {
			fmt.Fprintf (_buffer, "-- body fingerprint: `%s`\n", _document.BodyFingerprint)
		}
	}
	
	if _document.BodyEmpty {
		fmt.Fprintf (_buffer, "-- body: empty\n")
	} else if _includeBody {
		fmt.Fprintf (_buffer, "-- body:\n")
		fmt.Fprintf (_buffer, "~~~~~~~~\n")
		for _, _line := range _document.BodyLines {
			fmt.Fprintf (_buffer, "    %s\n", _line)
		}
		fmt.Fprintf (_buffer, "~~~~~~~~\n")
	}
	
	if _includeRender && (_document.RenderText != "") {
		fmt.Fprintf (_buffer, "-- render text:\n")
		fmt.Fprintf (_buffer, "~~~~~~~~\n")
		_lines, _ := stringSplitLines (_document.RenderText)
		for _, _line := range _lines {
			fmt.Fprintf (_buffer, "    %s\n", _line)
		}
		fmt.Fprintf (_buffer, "~~~~~~~~\n")
	}
	
	if _includeRender && (_document.RenderHtml != "") {
		fmt.Fprintf (_buffer, "-- render HTML:\n")
		fmt.Fprintf (_buffer, "~~~~~~~~\n")
		_lines, _ := stringSplitLines (_document.RenderHtml)
		for _, _line := range _lines {
			fmt.Fprintf (_buffer, "    %s\n", _line)
		}
		fmt.Fprintf (_buffer, "~~~~~~~~\n")
	}
	
	if _, _error := _buffer.WriteTo (_stream); _error != nil {
		return errorw (0x9fe2e5af, _error)
	}
	
	return nil
}

