

package zscratchpad


import "bufio"
import "bytes"
import "io"
import "os"
import "os/exec"
import "path"
import "sync"




type Editor struct {
	globals *Globals
	index *Index
}

type editSession struct {
	globals *Globals
	editor *Editor
	library *Library
	documentOld *Document
	documentNew *Document
	path string
	file *os.File
	command *exec.Cmd
	synchronous bool
	error *Error
}




func EditorNew (_globals *Globals, _index *Index) (*Editor, *Error) {
	_editor := & Editor {
			globals : _globals,
			index : _index,
		}
	return _editor, nil
}




func EditorDocumentEdit (_editor *Editor, _library *Library, _document *Document, _synchronous bool) (*Error) {
	
	_globals := _editor.globals
	
	if !_globals.TerminalEnabled && !_globals.XorgEnabled {
		return errorw (0xa302fef3, nil)
	}
	
	if !_library.EditEnabled {
		return errorw (0xfefc63a0, nil)
	}
	if !_document.EditEnabled {
		return errorw (0xaa64d776, nil)
	}
	
	_path := _document.Path
	if _path == "" {
		return errorw (0xbdb59e67, nil)
	}
	
	logf ('d', 0x226a3cbd, "[editor-session]  opening file for `%s`...", _path)
	
	_file, _error := os.OpenFile (_path, os.O_RDWR, 0)
	if _error != nil {
		return errorw (0xa51cdc41, _error)
	}
	
	_session := & editSession {
			globals : _globals,
			editor : _editor,
			library : _library,
			documentOld : _document,
			path : _path,
			file : _file,
			synchronous : _synchronous,
		}
	
	if !_session.synchronous {
		go editSessionRun (_session)
		return nil
	}
	
	return editSessionRun (_session)
}


func EditorDocumentCreate (_editor *Editor, _library *Library, _documentName string, _synchronous bool) (*Error) {
	
	_globals := _editor.globals
	
	if !_globals.TerminalEnabled && !_globals.XorgEnabled {
		return errorw (0x0175c9ec, nil)
	}
	
	if !_library.CreateEnabled {
		return errorw (0x2752e1cc, nil)
	}
	_path := path.Join (_library.CreatePath, _documentName)
	if _library.CreateExtension != "" {
		_path = _path + "." + _library.CreateExtension
	}
	
	logf ('d', 0x6292b948, "[editor-session]  creating file for `%s`...", _path)
	
	_file, _error := os.OpenFile (_path, os.O_RDWR | os.O_CREATE | os.O_EXCL, 0o640)
	if _error != nil {
		return errorw (0x5d8b586a, _error)
	}
	
	_session := & editSession {
			globals : _globals,
			editor : _editor,
			library : _library,
			path : _path,
			file : _file,
			synchronous : _synchronous,
		}
	
	if !_session.synchronous {
		go editSessionRun (_session)
		return nil
	}
	
	return editSessionRun (_session)
}


func editSessionRun (_session *editSession) (*Error) {
	
	_globals := _session.globals
	
	if ! _globals.TerminalMutexTryLock () {
		return errorw (0x5fcbecde, nil)
	}
	defer _globals.TerminalMutexUnlock ()
	
	logf ('d', 0x0edfabbf, "[editor-session]  launching editor for `%s`...", _session.path)
	
	_command, _error := EditorResolveEditCommand (_session.editor, _session.path)
	if _error != nil {
		_session.error = _error
		return editSessionClose (_session)
	}
	
	if _error := _command.Start (); _error != nil {
		_session.error = errorw (0x4b48b0bc, _error)
		return editSessionClose (_session)
	}
	
	_session.command = _command
	
	logf ('d', 0xff9ec344, "[editor-session]  waiting editor for `%s`...", _session.path)
	
	if _error := _session.command.Wait (); _error != nil {
		_session.error = errorw (0xe877f161, _error)
		return editSessionClose (_session)
	}
	
	return editSessionFinalize (_session)
}


func editSessionFinalize (_session *editSession) (*Error) {
	
	if _session.documentOld != nil {
		
		logf ('d', 0x48f7d5f5, "[editor-session]  reloading document for `%s`...", _session.path)
		
		if _document_0, _error := DocumentReload (_session.documentOld); _error == nil {
			_session.documentNew = _document_0
		} else {
			_session.error = _error
			return editSessionClose (_session)
		}
		
	} else {
		
		logf ('d', 0xff9b1d5d, "[editor-session]  loading document for `%s`...", _session.path)
		
		if _document_0, _error := DocumentLoadFromPath (_session.path); _error == nil {
			_session.documentNew = _document_0
		} else {
			_session.error = _error
			return editSessionClose (_session)
		}
	}
	
	if _session.documentOld != nil {
		_session.documentNew.Library = _session.documentOld.Library
		_session.documentNew.PathInLibrary = _session.documentOld.PathInLibrary
		_session.documentNew.EditEnabled = _session.documentOld.EditEnabled
	} else {
		if _session.library != nil {
			_session.documentNew.Library = _session.library.Identifier
		}
	}
	
	if _error := DocumentInitializeIdentifier (_session.documentNew, _session.library); _error != nil {
		_session.error = _error
		return editSessionClose (_session)
	}
	if _error := DocumentInitializeFormat (_session.documentNew, _session.library); _error != nil {
		_session.error = _error
		return editSessionClose (_session)
	}
	
	if _session.editor.index == nil {
		return editSessionClose (_session)
	}
	
	_session.globals.MutexLock ()
	defer _session.globals.MutexUnlock ()
	
	if _session.documentOld != nil {
		
		logf ('d', 0x44c67acc, "[editor-session]  reindexing document for `%s`...", _session.path)
		
		if _error := IndexDocumentUpdate (_session.editor.index, _session.documentNew, _session.documentOld); _error != nil {
			_session.error = _error
			return editSessionClose (_session)
		}
		
	} else {
		
		logf ('d', 0x5ee2c034, "[editor-session]  indexing document for `%s`...", _session.path)
		
		if _error := IndexDocumentInclude (_session.editor.index, _session.documentNew); _error != nil {
			_session.error = _error
			return editSessionClose (_session)
		}
	}
	
	return editSessionClose (_session)
}


func editSessionClose (_session *editSession) (*Error) {
	
	if _session.error != nil {
		if !_session.synchronous {
			logErrorf ('e', 0x35a898d8, _session.error, "[editor-session]  failed for `%s`!", _session.path)
		}
		return _session.error
	}
	
	logf ('d', 0x42c39bbe, "[editor-session]  succeeded for `%s`;", _session.path)
	
	return nil
}




func EditorSelect (_editor *Editor, _options []string) ([]string, *Error) {
	
	_globals := _editor.globals
	
	if !_globals.TerminalEnabled && !_globals.XorgEnabled {
		return nil, errorw (0xdafc150d, nil)
	}
	
	if ! _globals.TerminalMutexTryLock () {
		return nil, errorw (0xcba65bc9, nil)
	}
	defer _globals.TerminalMutexUnlock ()
	
	_command := (*exec.Cmd) (nil)
	if _command_0, _error := EditorResolveSelectCommand (_editor); _error == nil {
		_command = _command_0
	} else {
		return nil, _error
	}
	
	_stdin, _error := _command.StdinPipe ()
	if _error != nil {
		return nil, errorw (0x5acfc6bd, _error)
	}
	// NOTE:  Due to race conditions within the goroutine, we leave this to be closed by the garbage collector.
	// defer _stdin.Close ()
	
	_stdout, _error := _command.StdoutPipe ()
	if _error != nil {
		return nil, errorw (0x351240e9, _error)
	}
	// NOTE:  Due to race conditions within the goroutine, we leave this to be closed by the garbage collector.
	// defer _stdout.Close ()
	
	if _error := _command.Start (); _error != nil {
		return nil, errorw (0xd3c76e67, _error)
	}
	
	_selection := make ([]string, 0, 16)
	_waiter := & sync.WaitGroup {}
	
	_waiter.Add (1)
	_stdinError := (*Error) (nil)
	go func () () {
			_buffer := bytes.NewBuffer (nil)
			for _, _option := range _options {
				_buffer.WriteString (_option)
				_buffer.WriteByte ('\n')
			}
			if _, _error := _buffer.WriteTo (_stdin); _error != nil {
				_stdinError = errorw (0x5cf9fd4b, _error)
			}
			if _error := _stdin.Close (); _error != nil {
				_stdinError = errorw (0x610bd07e, _error)
			}
			_waiter.Done ()
		} ()
	
	_waiter.Add (1)
	_stdoutError := (*Error) (nil)
	go func () () {
			_buffer := bufio.NewReaderSize (_stdout, 4096)
			for {
				if _line, _error := _buffer.ReadString ('\n'); _error == nil {
					_lineLen := len (_line)
					if (_lineLen > 0) && (_line[_lineLen - 1] == '\n') {
						_line = _line[: _lineLen - 1]
					}
					_selection = append (_selection, _line)
				} else if _error == io.EOF {
					if _line != "" {
						_selection = append (_selection, _line)
					}
					break
				} else {
					_stdoutError = errorw (0x66b5573e, _error)
					break
				}
			}
			if _error := _stdout.Close (); _error != nil {
				_stdoutError = errorw (0x39e45fb4, _error)
			}
			_waiter.Done ()
		} ()
	
	_waiter.Wait ()
	
	if _error := _command.Wait (); _error != nil {
		return nil, errorw (0xe7d64749, _error)
	}
	
	if _stdinError != nil {
		return nil, _stdinError
	}
	if _stdoutError != nil {
		return nil, _stdoutError
	}
	
	return _selection, nil
}




func EditorResolveEditCommand (_editor *Editor, _path string) (*exec.Cmd, *Error) {
	
	_globals := _editor.globals
	
	if _globals.TerminalEnabled {
		
		_executable := ""
		_executableName := ""
		if _executableName_0, _ := _globals.Environment["EDITOR"]; _executableName_0 != "" {
			if _executable_0, _error := exec.LookPath (_executableName_0); _error == nil {
				_executable = _executable_0
				_executableName = _executableName_0
			} else {
				return nil, errorw (0xccba26a3, _error)
			}
		}
		if _executable == "" {
			for _, _executableName_0 := range []string { "z-scratchpad--edit", "x-edit", "nano", "vim", "emacs" } {
				if _executable_0, _error := exec.LookPath (_executableName_0); _error == nil {
					_executable = _executable_0
					_executableName = _executableName_0
					break
				}
			}
		}
		if _executable == "" {
			return nil, errorw (0x2eebed4d, nil)
		}
		
		_arguments := make ([]string, 0, 16)
		_arguments = append (_arguments, _executable)
		switch _executableName {
			case "z-scratchpad--edit", "x-edit" :
				_arguments = append (_arguments, _path)
			case "nano", "vim", "emacs" :
				_arguments = append (_arguments, "--", _path)
			default :
				_arguments = append (_arguments, _path)
		}
		
		_command := & exec.Cmd {
				Path : _executable,
				Args : _arguments,
				Env : _globals.EnvironmentList,
				Stdin : _globals.TerminalTty,
				Stdout : _globals.TerminalTty,
				Stderr : _globals.TerminalTty,
			}
		
		return _command, nil
		
	} else if _globals.XorgEnabled {
		
		_executable := ""
		_executableName := ""
		if _executable == "" {
			for _, _executableName_0 := range []string { "z-scratchpad--edit", "x-edit", "howl", "sublime_text", "gvim", "emacs-gtk", "emacs-x11" } {
				if _executable_0, _error := exec.LookPath (_executableName_0); _error == nil {
					_executable = _executable_0
					_executableName = _executableName_0
					break
				}
			}
		}
		if _executable == "" {
			return nil, errorw (0x5a7c2f6b, nil)
		}
		
		_arguments := make ([]string, 0, 16)
		_arguments = append (_arguments, _executable)
		switch _executableName {
			case "z-scratchpad--edit", "x-edit" :
				_arguments = append (_arguments, _path)
			case "howl", "gvim", "emacs-gtk", "emacs-x11" :
				_arguments = append (_arguments, "--", _path)
			case "sublime_text" :
				_arguments = append (_arguments, "--new-window", "--wait", "--", _path)
			default :
				_arguments = append (_arguments, _path)
		}
		
		_command := & exec.Cmd {
				Path : _executable,
				Args : _arguments,
				Env : _globals.EnvironmentList,
				Stdin : _globals.DevNull,
				Stdout : _globals.DevNull,
				Stderr : _globals.DevNull,
			}
		
		return _command, nil
		
	} else {
		
		return nil, errorw (0xfe957df1, nil)
	}
}




func EditorResolveSelectCommand (_editor *Editor) (*exec.Cmd, *Error) {
	
	_globals := _editor.globals
	
	if _globals.TerminalEnabled {
		
		_executable := ""
		_executableName := ""
		if _executable == "" {
			for _, _executableName_0 := range []string { "z-scratchpad--select", "x-select", "fzf" } {
				if _executable_0, _error := exec.LookPath (_executableName_0); _error == nil {
					_executable = _executable_0
					_executableName = _executableName_0
					break
				}
			}
		}
		if _executable == "" {
			return nil, errorw (0x10e4bef3, nil)
		}
		
		_arguments := make ([]string, 0, 32)
		_arguments = append (_arguments, _executable)
		switch _executableName {
			case "z-scratchpad--select", "x-select" :
				// NOP
			case "fzf" :
				_arguments = append (_arguments,
						"--prompt", ": ",
						"-e", "-x", "-i",
						"--tiebreak", "begin,length,index",
						"--no-mouse", "--no-color", "--no-bold",
					)
			default :
				// NOP
		}
		
		_command := & exec.Cmd {
				Path : _executable,
				Args : _arguments,
				Env : _globals.EnvironmentList,
				Stderr : _globals.TerminalTty,
			}
		
		return _command, nil
		
	} else if _globals.XorgEnabled {
		
		_executable := ""
		_executableName := ""
		if _executable == "" {
			for _, _executableName_0 := range []string { "z-scratchpad--select", "x-select", "rofi", "dmenu" } {
				if _executable_0, _error := exec.LookPath (_executableName_0); _error == nil {
					_executable = _executable_0
					_executableName = _executableName_0
					break
				}
			}
		}
		if _executable == "" {
			return nil, errorw (0xcdb975c1, nil)
		}
		
		_arguments := make ([]string, 0, 32)
		_arguments = append (_arguments, _executable)
		switch _executableName {
			case "z-scratchpad--select", "x-select" :
				// NOP
			case "rofi" :
				_arguments = append (_arguments, "-dmenu", "-p", "", "-i", "-no-custom", "-matching-negate-char", "\\x0")
			case "dmenu" :
				_arguments = append (_arguments, "-p", "", "-l", "16", "-i")
			default :
				// NOP
		}
		
		_command := & exec.Cmd {
				Path : _executable,
				Args : _arguments,
				Env : _globals.EnvironmentList,
				Stderr : _globals.DevNull,
			}
		
		return _command, nil
		
	} else {
		
		return nil, errorw (0xdced1bf6, nil)
	}
}

