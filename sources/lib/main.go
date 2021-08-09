

package zscratchpad


import "bytes"
import "encoding/json"
import "fmt"
import "net"
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
	Library *string `long:"library" short:"l" value-name:"{identifier}"`
	Type *string `long:"type" short:"t" choice:"library" choice:"document"`
	What *string `long:"what" short:"w" choice:"identifier" choice:"title" choice:"name" choice:"path"`
	Format *string `long:"format" short:"f" choice:"text" choice:"text-0" choice:"json"`
}

type SelectFlags struct {
	Library *string `long:"library" short:"l" value-name:"{identifier}"`
	Type *string `long:"type" short:"t" choice:"library" choice:"document"`
	What *string `long:"what" short:"w" choice:"identifier" choice:"title" choice:"name" choice:"path"`
	How *string `long:"how" short:"W" choice:"identifier" choice:"title" choice:"name" choice:"pah"`
	Format *string `long:"format" short:"f" choice:"text" choice:"text-0" choice:"json"`
}

type ExportFlags struct {
	Document *string `long:"document" short:"d" required:"-" value-name:"{identifier}"`
	Format *string `long:"format" short:"f" choice:"source" choice:"text" choice:"html"`
}

type EditFlags struct {
	Library *string `long:"library" short:"l" value-name:"{identifier}"`
	Document *string `long:"document" short:"d" value-name:"{identifier}"`
	Select *bool `long:"select" short:"s"`
}

type CreateFlags struct {
	Library *string `long:"library" short:"l" value-name:"{identifier}"`
	Document *string `long:"document" short:"d" value-name:"{identifier}"`
	Select *bool `long:"select" short:"s"`
}

type DumpFlags struct {}

type MainFlags struct {
	Global *GlobalFlags `group:"Global options"`
	Library *LibraryFlags `group:"Library options"`
	List *ListFlags `command:"list"`
	Select *SelectFlags `command:"select"`
	Export *ExportFlags `command:"export"`
	Edit *EditFlags `command:"edit"`
	Create *CreateFlags `command:"create"`
	Server *ServerFlags `command:"server"`
	Dump *DumpFlags `command:"dump"`
}




func Main (_executable string, _arguments []string, _environment map[string]string) (*Error) {
	
	_globals, _error := GlobalsNew (_executable, _environment)
	if _error != nil {
		return _error
	}
	
	_flags := & MainFlags {
			Global : & GlobalFlags {},
			Library : & LibraryFlags {},
			List : & ListFlags {},
			Select : & SelectFlags {},
			Export : & ExportFlags {},
			Edit : & EditFlags {},
			Create : & CreateFlags {},
			Server : & ServerFlags {},
			Dump : & DumpFlags {},
		}
	
	_parser := flags.NewNamedParser ("z-scratchpad", flags.PassDoubleDash)
	_parser.SubcommandsOptional = true
	if _, _error := _parser.AddGroup ("", "", _flags); _error != nil {
		return errorw (0x5b48e356, _error)
	}
	
	_help := func (_log bool, _error *Error) (*Error) {
		_buffer := bytes.NewBuffer (nil)
		_parser.WriteHelp (_buffer)
		if _log {
			if _globals.StdioIsTty && _globals.TerminalEnabled {
				logf ('`', 0xa725b4bc, "\n%s\n", _buffer.String ())
			}
		} else {
			if _, _error := _buffer.WriteTo (_globals.Stdout); _error != nil {
				return errorw (0xf4170873, _error)
			}
		}
		return _error
	}
	
	// FIXME:  The parser always uses the actual environment variables and not `_environment`!
	if _argumentsRest, _error := _parser.ParseArgs (_arguments); _error != nil {
		if flagBoolOrDefault (_flags.Global.Help, false) {
			return _help (false, nil)
		} else {
			return _help (true, errorw (0xa198fbfd, _error))
		}
	} else if len (_argumentsRest) != 0 {
		return _help (true, errorw (0x3c7b6224, nil))
	}
	
	if flagBoolOrDefault (_flags.Global.Help, false) {
		return _help (false, nil)
	}
	
	if _parser.Active == nil {
		return _help (true, errorw (0x4cae2ee5, nil))
	}
	
	return MainWithFlags (_parser.Active.Name, _flags, _globals)
}




func MainWithFlags (_command string, _flags *MainFlags, _globals *Globals) (*Error) {
	
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
		
		case "select" :
			return MainSelect (_flags.Select, _globals, _index, _editor)
		
		case "export" :
			return MainExport (_flags.Export, _globals, _index)
		
		case "edit" :
			return MainEdit (_flags.Edit, _globals, _index, _editor)
		
		case "create" :
			return MainCreate (_flags.Create, _globals, _index, _editor)
		
		case "server" :
			return MainServer (_flags.Server, _globals, _index, _editor)
		
		case "dump" :
			return MainDump (_flags.Dump, _globals, _index)
		
		default :
			return errorw (0xaca17bb9, nil)
	}
}




func MainExport (_flags *ExportFlags, _globals *Globals, _index *Index) (*Error) {
	
	if _flags.Document == nil {
		return errorw (0x1826914a, nil)
	}
	_document, _error := WorkflowDocumentResolve (*_flags.Document, _index)
	if _error != nil {
		return _error
	}
	
	_format := flagStringOrDefault (_flags.Format, "source")
	
	_buffer := (*bytes.Buffer) (nil)
	switch _format {
		
		case "source" :
			if _output, _error := DocumentRenderToSource (_document); _error == nil {
				_buffer = bytes.NewBufferString (_output)
			} else {
				return _error
			}
		
		case "html" :
			if _output, _error := DocumentRenderToHtml (_document); _error == nil {
				_buffer = bytes.NewBufferString (_output)
			} else {
				return _error
			}
		
		case "text" :
			if _output, _error := DocumentRenderToText (_document); _error == nil {
				_buffer = bytes.NewBufferString (_output)
			} else {
				return _error
			}
		
		default :
			return errorw (0x326240d3, nil)
	}
	
	if _, _error := _buffer.WriteTo (_globals.Stdout); _error != nil {
		return errorw (0xa797b17f, _error)
	}
	
	return nil
}




func MainEdit (_flags *EditFlags, _globals *Globals, _index *Index, _editor *Editor) (*Error) {
	
	_flagSelect := flagBoolOrDefault (_flags.Select, false)
	if _flagSelect && (_flags.Document != nil) {
		return errorw (0x17114913, nil)
	}
	
	_identifier := ""
	if _flagSelect {
		
		_libraryIdentifier := flagStringOrDefault (_flags.Library, "")
		_options, _error := mainListOptionsAndSelect (_libraryIdentifier, "document", "identifier", "title", _index, _editor)
		if _error != nil {
			return _error
		}
		switch len (_options) {
			case 0 :
				return errorw (0x29abcd02, nil)
			case 1 :
				_identifier = _options[0][1]
			default :
				return errorw (0x22d4ddbe, nil)
		}
		
	} else {
		
		if _library, _document, _error := mainMergeLibraryAndDocumentIdentifiers (_flags.Library, _flags.Document); _error != nil {
			return _error
		} else if _document != "" {
			_identifier = _document
		} else if _library != "" {
			return errorw (0xdbe83c6c, nil)
		} else {
			return errorw (0xa0fc749c, nil)
		}
	}
	
	return WorkflowDocumentEdit (_identifier, _index, _editor, true)
}




func MainCreate (_flags *CreateFlags, _globals *Globals, _index *Index, _editor *Editor) (*Error) {
	
	_flagSelect := flagBoolOrDefault (_flags.Select, false)
	if _flagSelect && (_flags.Document != nil) {
		return errorw (0x2a0a4328, nil)
	}
	if _flagSelect && (_flags.Library != nil) {
		return errorw (0x4d3444df, nil)
	}
	
	_identifier := ""
	if _flagSelect {
		
		_options, _error := mainListOptionsAndSelect ("", "library", "identifier", "title", _index, _editor)
		if _error != nil {
			return _error
		}
		switch len (_options) {
			case 0 :
				return errorw (0x29abcd02, nil)
			case 1 :
				_identifier = _options[0][1]
			default :
				return errorw (0x22d4ddbe, nil)
		}
		
	} else {
		
		if _library, _document, _error := mainMergeLibraryAndDocumentIdentifiers (_flags.Library, _flags.Document); _error != nil {
			return _error
		} else if _document != "" {
			_identifier = _document
		} else if _library != "" {
			_identifier = _library
		} else {
			return errorw (0x22cc7dea, nil)
		}
	}
	
	return WorkflowDocumentCreate (_identifier, _index, _editor, true)
}




func MainList (_flags *ListFlags, _globals *Globals, _index *Index) (*Error) {
	
	_libraryIdentifier := flagStringOrDefault (_flags.Library, "")
	_type := flagStringOrDefault (_flags.Type, "documents")
	_what := flagStringOrDefault (_flags.What, "identifiers")
	_format := flagStringOrDefault (_flags.Format, "text")
	
	_options, _error := mainListOptions (_libraryIdentifier, _type, _what, "identifiers", _index)
	if _error != nil {
		return _error
	}
	
	return mainListOutput (_options, _format, _globals)
}


func MainSelect (_flags *SelectFlags, _globals *Globals, _index *Index, _editor *Editor) (*Error) {
	
	_libraryIdentifier := flagStringOrDefault (_flags.Library, "")
	_type := flagStringOrDefault (_flags.Type, "documents")
	_what := flagStringOrDefault (_flags.What, "identifiers")
	_how := flagStringOrDefault (_flags.How, "titles")
	_format := flagStringOrDefault (_flags.Format, "text")
	
	_options, _error := mainListOptionsAndSelect (_libraryIdentifier, _type, _what, _how, _index, _editor)
	if _error != nil {
		return _error
	}
	
	return mainListOutput (_options, _format, _globals)
}


func mainListOptionsAndSelect (_libraryIdentifier string, _type string, _what string, _how string, _index *Index, _editor *Editor) ([][2]string, *Error) {
	
	_options, _error := mainListOptions (_libraryIdentifier, _type, _what, _how, _index)
	if _error != nil {
		return nil, _error
	}
	
	_selection, _error := mainListSelect (_options, _editor)
	if _error != nil {
		return nil, _error
	}
	
	return _selection, nil
}


func mainListOptions (_libraryIdentifier string, _type string, _what string, _how string, _index *Index) ([][2]string, *Error) {
	
	_library := (*Library) (nil)
	if _libraryIdentifier != "" {
		if _library_0, _error := WorkflowLibraryResolve (_libraryIdentifier, _index); _error == nil {
			_library = _library_0
		} else {
			return nil, errorw (0x5a3e46e1, nil)
		}
	}
	
	_options := make ([][2]string, 0, 1024)
	
	switch _type {
		
		case "libraries", "library" :
			_libraries := []*Library (nil)
			if _library != nil {
				_libraries = []*Library { _library }
			} else {
				if _libraries_0, _error := IndexLibrariesSelectAll (_index); _error == nil {
					_libraries = _libraries_0
				} else {
					return nil, _error
				}
			}
			for _, _library := range _libraries {
				_value := ""
				switch _what {
					case "identifiers", "identifier" :
						_value = _library.Identifier
					case "titles", "names", "title", "name" :
						_value = _library.Name
						if _value == "" {
							_value = "[" + _library.Identifier + "]"
						}
					case "paths", "path" :
						_value = _library.Path
					default :
						return nil, errorw (0x4fab7acb, nil)
				}
				_label := ""
				switch _how {
					case "identifiers", "identifier" :
						_label = _library.Identifier
					case "titles", "names", "title", "name" :
						_label = _library.Name
					case "paths", "path" :
						_label = _library.Path
					default :
						return nil, errorw (0xf0f17afb, nil)
				}
				if _label == "" {
					_label = "[" + _library.Identifier + "]"
				}
				if (_label != "") && (_value != "") {
					_options = append (_options, [2]string { _label, _value })
				}
			}
		
		case "documents", "document" :
			_documents := []*Document (nil)
			if _library != nil {
				if _documents_0, _error := IndexDocumentsSelectInLibrary (_index, _library.Identifier); _error == nil {
					_documents = _documents_0
				} else {
					return nil, _error
				}
			} else {
				if _documents_0, _error := IndexDocumentsSelectAll (_index); _error == nil {
					_documents = _documents_0
				} else {
					return nil, _error
				}
			}
			for _, _document := range _documents {
				_value := ""
				switch _what {
					case "identifiers", "identifier" :
						_value = _document.Identifier
					case "titles", "names", "title", "name" :
						_value = _document.Title
						if _value == "" {
							_value = "[" + _document.Identifier + "]"
						}
					case "paths", "path" :
						_value = _document.Path
					default :
						return nil, errorw (0x2f341212, nil)
				}
				_label := ""
				switch _how {
					case "identifiers", "identifier" :
						_label = _document.Identifier
					case "titles", "names", "title", "name" :
						_label = _document.Title
					case "paths", "path" :
						_label = _document.Path
					default :
						return nil, errorw (0x9f3c1037, nil)
				}
				if _label == "" {
					_label = "[" + _document.Identifier + "]"
				}
				if (_label != "") && (_value != "") {
					_options = append (_options, [2]string { _label, _value })
				}
			}
		
		default :
			return nil, errorw (0x2c37fb9c, nil)
	}
	
	return _options, nil
}


func mainListSelect (_options [][2]string, _editor *Editor) ([][2]string, *Error) {
	
	_labels := make ([]string, 0, len (_options))
	_values := make (map[string][]string, len (_options))
	for _, _option := range _options {
		_label := _option[0]
		_value := _option[1]
		_labels = append (_labels, _label)
		if _, _exists := _values[_label]; _exists {
			// FIXME:  How should we handle duplicate labels?
			_values[_label] = append (_values[_label], _value)
		} else {
			_values[_label] = []string { _value }
		}
	}
	
	sort.Strings (_labels)
	
	_selection_0, _error := EditorSelect (_editor, _labels)
	if _error != nil {
		return nil, _error
	}
	
	_selection := make ([][2]string, 0, 16)
	for _, _label := range _selection_0 {
		if _values_0, _exists := _values[_label]; _exists {
			for _, _value := range _values_0 {
				_selection = append (_selection, [2]string { _label, _value})
			}
		} else {
			return nil, errorw (0xdbff774c, nil)
		}
	}
	
	return _selection, nil
}


func mainListOutput (_options [][2]string, _format string, _globals *Globals) (*Error) {
	
	_list := make ([]string, 0, len (_options))
	for _, _option := range _options {
		_value := _option[1]
		_list = append (_list, _value)
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
	
	if _, _error := _buffer.WriteTo (_globals.Stdout); _error != nil {
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
	
	_globals.TerminalEnabled = false
	
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
	
	if _, _error := _buffer.WriteTo (_globals.Stdout); _error != nil {
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




func mainMergeLibraryAndDocumentIdentifiers (_library *string, _document *string) (string, string, *Error) {
	
	if _library != nil {
		
		if _document != nil {
			if _identifier, _error := DocumentFormatIdentifier (*_library, *_document); _error == nil {
				return "", _identifier, nil
			} else {
				return "", "", _error
			}
		} else {
			if _identifier, _error := LibraryParseIdentifier (*_library); _error == nil {
				return _identifier, "", nil
			} else {
				return "", "", _error
			}
		}
		
	} else if _document != nil {
		
		if _identifier, _, _, _error := DocumentParseIdentifier (*_document); _error == nil {
			return "", _identifier, nil
		} else {
			return "", "", _error
		}
		
	} else {
		
		return "", "", nil
	}
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

