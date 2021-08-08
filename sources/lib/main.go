

package zscratchpad


import "bytes"
import "fmt"
import "net"
import "os"


import "github.com/jessevdk/go-flags"




type GlobalFlags struct {
	Help *bool `long:"help" short:"h"`
}

type LibraryFlags struct {
	Path *string `long:"library-path" value-name:"{library-path}"`
	UseFileNameAsIdentifier *bool `long:"library-use-file-name"`
	UseFileExtensionAsFormat *bool `long:"library-use-file-ext"`
}

type DumpFlags struct {}

type ServerFlags struct {
	EndpointIp *string `long:"server-ip" value-name:"{ip}"`
	EndpointPort *uint16 `long:"server-port" value-name:"{port}"`
}

type MainFlags struct {
	Global *GlobalFlags `group:"Global options"`
	Library *LibraryFlags `group:"Library options"`
	Server *ServerFlags `command:"server"`
	Dump *DumpFlags `command:"dump"`
}




func Main (_executable string, _arguments []string, _environment map[string]string) (*Error) {
	
	_flags := & MainFlags {
			Global : & GlobalFlags {},
			Library : & LibraryFlags {},
			Server : & ServerFlags {},
			Dump : & DumpFlags {},
		}
	
	_parser := flags.NewNamedParser ("z-scratchpad", flags.PassDoubleDash)
	_parser.SubcommandsOptional = true
	if _, _error := _parser.AddGroup ("", "", _flags); _error != nil {
		return errorw (0x5b48e356, _error)
	}
	
	// FIXME:  The parser always uses the actual environment variables and not `_environment`!
	if _argumentsRest, _error := _parser.ParseArgs (_arguments); _error != nil {
		return errorw (0xa198fbfd, _error)
	} else if len (_argumentsRest) != 0 {
		return errorw (0x3c7b6224, nil)
	}
	
	if flagBoolOrDefault (_flags.Global.Help, false) {
		_parser.WriteHelp (os.Stdout)
		return nil
	}
	
	if _parser.Active == nil {
		_buffer := bytes.NewBuffer (nil)
		_parser.WriteHelp (_buffer)
		logf ('`', 0xa725b4bc, "\n%s\n", _buffer.String ())
		return errorw (0x4cae2ee5, nil)
	}
	
	return MainWithFlags (_parser.Active.Name, _flags)
}




func MainWithFlags (_command string, _flags *MainFlags) (*Error) {
	
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
	
	_error = MainLoadLibraries (_flags.Library, _globals, _index)
	if _error != nil {
		return _error
	}
	
	switch _command {
		
		case "server" :
			return MainServer (_flags.Server, _globals, _index, _editor)
		
		case "dump" :
			return MainDump (_flags.Dump, _globals, _index)
		
		default :
			return errorw (0xaca17bb9, nil)
	}
}




func MainServer (_flags *ServerFlags, _globals *Globals, _index *Index, _editor *Editor) (*Error) {
	
	_endpointIp := flagStringOrDefault (_flags.EndpointIp, "127.13.160.195")
	_endpointPort := flagUint16OrDefault (_flags.EndpointPort, 8080)
	
	_endpoint := fmt.Sprintf ("%s:%d", _endpointIp, _endpointPort)
	
	logf ('i', 0x210494be, "[server]  listening on `%s`...", _endpoint)
	
	_listener, _error_0 := net.Listen ("tcp", _endpoint)
	if _error_0 != nil {
		return errorw (0xedeea766, _error_0)
	}
	
	_server, _error := ServerNew (_globals, _index, _editor, _listener)
	if _error != nil {
		return _error
	}
	
	_error = ServerRun (_server)
	if _error != nil {
		return _error
	}
	
	return nil
}




func MainDump (_flags *DumpFlags, _globals *Globals, _index *Index) (*Error) {
	
	_documents, _error := IndexDocumentsSelectAll (_index)
	if _error != nil {
		return _error
	}
	
	_buffer := bytes.NewBuffer (nil)
	for _, _document := range _documents {
		_buffer.WriteString ("\n")
		_error = DocumentDump (_buffer, _document, true, false, false)
		if _error != nil {
			return _error
		}
		_buffer.WriteString ("\n")
	}
	
	if _, _error := _buffer.WriteTo (os.Stdout); _error != nil {
		return errorw (0xbf6a449c, _error)
	}
	
	return nil
}




func MainLoadLibraries (_flags *LibraryFlags, _globals *Globals, _index *Index) (*Error) {
	
	_libraries := []*Library (nil)
	
	if _flags.Path != nil {
		_library := & Library {
				Identifier : "library",
				Name : "Library",
				Path : *_flags.Path,
				UseFileNameAsIdentifier : flagBoolOrDefault (_flags.UseFileNameAsIdentifier, false),
				UseFileExtensionAsFormat : flagBoolOrDefault (_flags.UseFileExtensionAsFormat, false),
			}
		_libraries = []*Library { _library }
	}
	
	if _libraries == nil {
		_libraries = []*Library {
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
	}
	
	for _, _library := range _libraries {
		
		_error := IndexLibraryInclude (_index, _library)
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
	
	return nil
}




func flagBoolOrDefault (_value *bool, _default bool) (bool) {
	if _value == nil {
		return _default
	}
	return *_value
}

func flagUint16OrDefault (_value *uint16, _default uint16) (uint16) {
	if _value == nil {
		return _default
	}
	return *_value
}

func flagStringOrDefault (_value *string, _default string) (string) {
	if _value == nil {
		return _default
	}
	return *_value
}

