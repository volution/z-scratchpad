

package zscratchpad


import "bytes"
import "encoding/json"
import "fmt"
import "net"
import "os"
import "sort"


import "github.com/jessevdk/go-flags"




type GlobalFlags struct {
	Help *bool `long:"help" short:"h"`
}

type LibraryFlags struct {
	Path *string `long:"library-path" value-name:"{library-path}"`
	UseFileNameAsIdentifier *bool `long:"library-use-file-name"`
	UseFileExtensionAsFormat *bool `long:"library-use-file-ext"`
}

type ServerFlags struct {
	EndpointIp *string `long:"server-ip" value-name:"{ip}"`
	EndpointPort *uint16 `long:"server-port" value-name:"{port}"`
}

type ListFlags struct {
	Type *string `long:"type" short:"t" choice:"libraries" choice:"documents"`
	What *string `long:"what" short:"w" choice:"identifiers" choice:"titles" choice:"names" choice:"paths"`
	Format *string `long:"format" short:"f" choice:"text" choice:"text-0" choice:"json"`
}

type DumpFlags struct {}

type MainFlags struct {
	Global *GlobalFlags `group:"Global options"`
	Library *LibraryFlags `group:"Library options"`
	List *ListFlags `command:"list"`
	Server *ServerFlags `command:"server"`
	Dump *DumpFlags `command:"dump"`
}




func Main (_executable string, _arguments []string, _environment map[string]string) (*Error) {
	
	_flags := & MainFlags {
			Global : & GlobalFlags {},
			Library : & LibraryFlags {},
			List : & ListFlags {},
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
		_buffer := bytes.NewBuffer (nil)
		_parser.WriteHelp (_buffer)
		if _, _error := _buffer.WriteTo (os.Stdout); _error != nil {
			return errorw (0xf4170873, _error)
		}
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
		
		case "list" :
			return MainList (_flags.List, _globals, _index)
		
		case "server" :
			return MainServer (_flags.Server, _globals, _index, _editor)
		
		case "dump" :
			return MainDump (_flags.Dump, _globals, _index)
		
		default :
			return errorw (0xaca17bb9, nil)
	}
}




func MainList (_flags *ListFlags, _globals *Globals, _index *Index) (*Error) {
	
	_type := flagStringOrDefault (_flags.Type, "documents")
	_what := flagStringOrDefault (_flags.What, "identifiers")
	_format := flagStringOrDefault (_flags.Format, "text")
	
	_list := make ([]string, 0, 1024)
	
	switch _type {
		
		case "libraries", "library" :
			_libraries, _error := IndexLibrariesSelectAll (_index)
			if _error != nil {
				return _error
			}
			for _, _library := range _libraries {
				_value := ""
				switch _what {
					case "identifiers", "identifier" :
						_value = _library.Identifier
					case "titles", "names", "title", "name" :
						_value = _library.Name
					case "paths", "path" :
						_value = _library.Path
					default :
						return errorw (0x4fab7acb, nil)
				}
				if _value != "" {
					_list = append (_list, _value)
				}
			}
		
		case "documents", "document" :
			_documents, _error := IndexDocumentsSelectAll (_index)
			if _error != nil {
				return _error
			}
			for _, _document := range _documents {
				_value := ""
				switch _what {
					case "identifiers", "identifier" :
						_value = _document.Identifier
					case "titles", "names", "title", "name" :
						_value = _document.Title
					case "paths", "path" :
						_value = _document.Path
					default :
						return errorw (0x2f341212, nil)
				}
				if _value != "" {
					_list = append (_list, _value)
				}
			}
		
		default :
			return errorw (0x2c37fb9c, nil)
	}
	
	sort.Strings (_list)
	
	_buffer := bytes.NewBuffer (nil)
	
	switch _format {
		
		case "text", "text-0" :
			_separator := byte ('\n')
			if _format == "text-0" {
				_separator = 0
			}
			for _, _value := range _list {
				_buffer.WriteString (_value)
				_buffer.WriteByte (_separator)
			}
		
		case "json" :
			_encoder := json.NewEncoder (_buffer)
			if _error := _encoder.Encode (_list); _error != nil {
				return errorw (0xc65a050c, _error)
			}
		
		default :
			return errorw (0x4def007c, nil)
	}
	
	if _, _error := _buffer.WriteTo (os.Stdout); _error != nil {
		return errorw (0xcf76965f, _error)
	}
	
	return nil
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

