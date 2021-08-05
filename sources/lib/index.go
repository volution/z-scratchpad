

package zscratchpad


import "sort"




type Index struct {
	globals *Globals
	documents map[string]*Document
	libraries map[string]*Library
	libraryDocuments map[string][]string
}




func IndexNew (_globals *Globals) (*Index, *Error) {
	_index := & Index {
			globals : _globals,
			documents : make (map[string]*Document, 1024),
			libraries : make (map[string]*Library, 128),
			libraryDocuments : make (map[string][]string, 128),
		}
	return _index, nil
}




func IndexLibraryInclude (_index *Index, _library *Library) (*Error) {
	if _library.Identifier == "" {
		return errorw (0xcba1a612, nil)
	}
	if _, _exists := _index.libraries[_library.Identifier]; _exists {
		return errorw (0xb8990674, nil)
	}
	_index.libraries[_library.Identifier] = _library
	_index.libraryDocuments[_library.Identifier] = make ([]string, 0, 1024)
	return nil
}




func IndexDocumentInclude (_index *Index, _document *Document) (*Error) {
	if _document.Identifier == "" {
		return errorw (0x5567bd56, nil)
	}
	if _document.Library == "" {
		return errorw (0xf45e4633, nil)
	}
	if _, _exists := _index.libraries[_document.Library]; ! _exists {
		return errorw (0x4c68b8a9, nil)
	}
	if _, _exists := _index.documents[_document.Identifier]; _exists {
		return errorw (0x9c9f9c42, nil)
	}
	_index.documents[_document.Identifier] = _document
	_index.libraryDocuments[_document.Library] = append (_index.libraryDocuments[_document.Library], _document.Identifier)
	return nil
}




func IndexLibrariesSelectAll (_index *Index) ([]*Library, *Error) {
	_libraries := make ([]*Library, 0, len (_index.libraries))
	for _, _library := range _index.libraries {
		_libraries = append (_libraries, _library)
	}
	LibrariesSort (_libraries)
	return _libraries, nil
}

func IndexDocumentsSelectAll (_index *Index) ([]*Document, *Error) {
	_documents := make ([]*Document, 0, len (_index.documents))
	for _, _document := range _index.documents {
		_documents = append (_documents, _document)
	}
	DocumentsSort (_documents)
	return _documents, nil
}

func IndexDocumentsSelectInLibrary (_index *Index, _libraryIdentifier string) ([]*Document, *Error) {
	_documents := make ([]*Document, 0, len (_index.documents))
	_libraryDocuments, _libraryExists := _index.libraryDocuments[_libraryIdentifier]
	if !_libraryExists {
		return nil, errorw (0xb14719d6, nil)
	}
	for _, _documentIdentifier := range _libraryDocuments {
		_document, _documentExists := _index.documents[_documentIdentifier]
		if !_documentExists {
			return nil, errorw (0x842e9a51, nil)
		}
		_documents = append (_documents, _document)
	}
	DocumentsSort (_documents)
	return _documents, nil
}




func IndexLibraryResolve (_index *Index, _identifier string) (*Library, *Error) {
	if _identifier == "" {
		return nil, errorw (0xabe00a27, nil)
	}
	_library, _ := _index.libraries[_identifier]
	return _library, nil
}

func IndexDocumentResolve (_index *Index, _identifier string) (*Document, *Error) {
	if _identifier == "" {
		return nil, errorw (0x25e42042, nil)
	}
	_document, _ := _index.documents[_identifier]
	return _document, nil
}




func LibrariesSort (_libraries []*Library) () {
	sort.Slice (_libraries, func (_leftIndex, _rightIndex int) (bool) {
			_left := _libraries[_leftIndex]
			_right := _libraries[_rightIndex]
			return compareByNameOrIdentifier (_left.Name, _left.Identifier, _right.Name, _right.Identifier)
		})
}

func DocumentsSort (_documents []*Document) () {
	sort.Slice (_documents, func (_leftIndex, _rightIndex int) (bool) {
			_left := _documents[_leftIndex]
			_right := _documents[_rightIndex]
			return compareByNameOrIdentifier (_left.Title, _left.Identifier, _right.Title, _right.Identifier)
		})
}


func compareByNameOrIdentifier (_leftTitle, _leftIdentifier, _rightTitle, _rightIdentifier string) (bool) {
	if (_leftTitle != "") && (_rightTitle != "") {
		if _leftTitle == _rightTitle {
			return _leftIdentifier < _rightIdentifier
		}
		return _leftTitle < _rightTitle
	} else if (_leftTitle == "") && (_rightTitle == "") {
		return _leftIdentifier < _rightIdentifier
	} else if (_leftTitle != "") {
		return true
	} else if (_rightTitle != "") {
		return false
	} else {
		return false
	}
}

