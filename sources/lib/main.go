

package zscratchpad


import "bytes"
import "net"
import "os"




func Main (_executable string, _arguments []string, _environment map[string]string) (*Error) {
	
	if len (_arguments) != 0 {
		return errorw (0x79150e1f, nil)
	}
	
	_serverEndpoint := "127.79.75.28:8080"
	
	_globals, _error := GlobalsNew ()
	if _error != nil {
		return _error
	}
	
	_index, _error := IndexNew (_globals)
	if _error != nil {
		return _error
	}
	
	_editor, _error := EditorNew (_globals, _index)
	if _error != nil {
		return _error
	}
	
	_libraries := []*Library {
			{
				Identifier : "inbox",
				Name : "Inbox",
				Path : "./examples/inbox",
				UseFileNameAsIdentifier : true,
				UseFileExtensionAsFormat : true,
			},
			{
				Identifier : "tests",
				Name : "Tests",
				Path : "./examples/tests",
				UseFileNameAsIdentifier : true,
				UseFileExtensionAsFormat : true,
			},
		}
	
	for _, _library := range _libraries {
		
		_error = IndexLibraryInclude (_index, _library)
		if _error != nil {
			return _error
		}
		
		_documentPaths, _error := libraryDocumentsWalk (_library.Path)
		if _error != nil {
			return _error
		}
		
		_documents, _error := libraryDocumentsLoad (_library.Path, _documentPaths)
		if _error != nil {
			return _error
		}
		
		for _, _document := range _documents {
			
			if _document.Library == "" {
				_document.Library = _library.Identifier
			}
			
			_error = DocumentResolveIdentifier (_document, _library.UseFileNameAsIdentifier)
			if _error != nil {
				return _error
			}
			
			_error = DocumentResolveFormat (_document, _library.UseFileExtensionAsFormat)
			if _error != nil {
				return _error
			}
			
			_error = IndexDocumentInclude (_index, _document)
			if _error != nil {
				return _error
			}
		}
	}
	
	if true {
		_documents, _error := IndexDocumentsSelectAll (_index)
		if _error != nil {
			return _error
		}
		for _, _document := range _documents {
			_, _error = DocumentRenderToText (_document)
			if _error != nil {
				return _error
			}
			_, _error = DocumentRenderToHtml (_document)
			if _error != nil {
				return _error
			}
		}
	}
	if true {
		_documents, _error := IndexDocumentsSelectAll (_index)
		if _error != nil {
			return _error
		}
		_buffer := bytes.NewBuffer (nil)
		for _, _document := range _documents {
			_error = DocumentDump (_buffer, _document, true, false, false)
			if _error != nil {
				return _error
			}
			_buffer.WriteString ("\n")
		}
		if false {
			if _, _error := _buffer.WriteTo (os.Stdout); _error != nil {
				return errorw (0xbf6a449c, _error)
			}
		}
	}
	
	if _serverEndpoint != "" {
		_serverListener, _error_0 := net.Listen ("tcp", _serverEndpoint)
		if _error_0 != nil {
			return errorw (0xedeea766, _error_0)
		}
		_server, _error := ServerNew (_globals, _index, _editor, _serverListener)
		if _error != nil {
			return _error
		}
		_error = ServerRun (_server)
		if _error != nil {
			return _error
		}
	}
	
	return nil
}

