

package zscratchpad


import "time"




func WorkflowDocumentCreate (_identifierUnsafe string, _index *Index, _editor *Editor, _synchronous bool) (*Error) {
	
	_timestamp := time.Now ()
	
	_libraryIdentifier := ""
	_documentName := ""
	if _libraryIdentifier == "" {
		if _identifierUnsafe == "" {
			if _editor.DefaultCreateLibrary != "" {
				_libraryIdentifier = _editor.DefaultCreateLibrary
			} else {
				return errorw (0x19f48aa6, nil)
			}
		}
	}
	if _libraryIdentifier == "" {
		if _libraryIdentifier_0, _error := LibraryParseIdentifier (_identifierUnsafe); _error == nil {
			_libraryIdentifier = _libraryIdentifier_0
		}
	}
	if _libraryIdentifier == "" {
		if _, _libraryIdentifier_0, _documentName_0, _error := DocumentParseIdentifier (_identifierUnsafe); _error == nil {
			_libraryIdentifier = _libraryIdentifier_0
			_documentName = _documentName_0
		}
	}
	if _libraryIdentifier == "" {
		return errorw (0x4f21b7fb, nil)
	}
	
	_library, _error := IndexLibraryResolve (_index, _libraryIdentifier)
	if _error != nil {
		return _error
	}
	if _library == nil {
		return errorw (0x5e581595, nil)
	}
	
	if _documentName == "" {
		if _library.CreateNameTimestampLength > 0 {
			_format := ""
			switch _library.CreateNameTimestampLength {
				case 1 :
					_format = "2006"
				case 2 :
					_format = "2006-01"
				case 3 :
					_format = "2006-01-02"
				case 4 :
					_format = "2006-01-02-15"
				case 5 :
					_format = "2006-01-02-15-04"
				case 6 :
					_format = "2006-01-02-15-04-05"
				default :
					return errorw (0x770836aa, nil)
			}
			_token := _timestamp.Format (_format)
			if _documentName == "" {
				_documentName = _token
			} else {
				_documentName = _documentName + "--" + _token
			}
		}
		if _library.CreateNameRandomLength > 0 {
			_token := generateRandomToken ()
			_token = _token[: _library.CreateNameRandomLength]
			if _documentName == "" {
				_documentName = _token
			} else {
				_documentName = _documentName + "--" + _token
			}
		}
	}
	
	_identifier, _error := DocumentFormatIdentifier (_libraryIdentifier, _documentName)
	if _error != nil {
		return _error
	}
	
	_documentExisting, _error := IndexDocumentResolve (_index, _identifier)
	if _error != nil {
		return _error
	}
	if _documentExisting != nil {
		return errorw (0x538cfbae, nil)
	}
	
	
	return EditorDocumentCreate (_editor, _library, _documentName, _synchronous)
}




func WorkflowDocumentEdit (_identifierUnsafe string, _index *Index, _editor *Editor, _synchronous bool) (*Error) {
	
	_document, _library, _error := WorkflowDocumentAndLibraryResolve (_identifierUnsafe, _index)
	if _error != nil {
		return _error
	}
	
	return EditorDocumentEdit (_editor, _library, _document, _synchronous)
}




func WorkflowDocumentBrowse (_identifierUnsafe string, _index *Index, _browser *Browser, _synchronous bool) (*Error) {
	
	_document, _library, _error := WorkflowDocumentAndLibraryResolve (_identifierUnsafe, _index)
	if _error != nil {
		return _error
	}
	
	return BrowserDocumentOpen (_browser, _library, _document, _synchronous)
}


func WorkflowLibraryBrowse (_identifierUnsafe string, _index *Index, _browser *Browser, _synchronous bool) (*Error) {
	
	_library, _error := WorkflowLibraryResolve (_identifierUnsafe, _index)
	if _error != nil {
		return _error
	}
	
	return BrowserLibraryOpen (_browser, _library, _synchronous)
}




func WorkflowLibraryResolve (_identifierUnsafe string, _index *Index) (*Library, *Error) {
	if _identifierUnsafe == "" {
		return nil, errorw (0xbef72625, nil)
	}
	_identifier, _error := LibraryParseIdentifier (_identifierUnsafe)
	if _error != nil {
		return nil, _error
	}
	_library, _error := IndexLibraryResolve (_index, _identifier)
	if _error != nil {
		return nil, _error
	}
	if _library == nil {
		return nil, errorw (0xb1852bf9, nil)
	}
	return _library, nil
}


func WorkflowDocumentResolve (_identifierUnsafe string, _index *Index) (*Document, *Error) {
	if _identifierUnsafe == "" {
		return nil, errorw (0xc7f50900, nil)
	}
	_identifier, _, _, _error := DocumentParseIdentifier (_identifierUnsafe)
	if _error != nil {
		return nil, _error
	}
	_document, _error := IndexDocumentResolve (_index, _identifier)
	if _error != nil {
		return nil, _error
	}
	if _document == nil {
		return nil, errorw (0x054e7a60, nil)
	}
	return _document, nil
}


func WorkflowDocumentAndLibraryResolve (_identifierUnsafe string, _index *Index) (*Document, *Library, *Error) {
	_document, _error := WorkflowDocumentResolve (_identifierUnsafe, _index)
	if _error != nil {
		return nil, nil, _error
	}
	if _document.Library == "" {
		return _document, nil, nil
	}
	_library, _error := WorkflowLibraryResolve (_document.Library, _index)
	if _error != nil {
		return nil, nil, _error
	}
	return _document, _library, nil
}

