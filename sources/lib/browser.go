

package zscratchpad


import "os/exec"
import "strings"




type Browser struct {
	
	globals *Globals
	index *Index
	
	ServerUrlBase string
	
	TerminalOpenInternalCommand []string
	XorgOpenInternalCommand []string
	
	TerminalOpenExternalCommand []string
	XorgOpenExternalCommand []string
}




func BrowserNew (_globals *Globals, _index *Index) (*Browser, *Error) {
	_browser := & Browser {
			globals : _globals,
			index : _index,
		}
	return _browser, nil
}




func BrowserDocumentOpen (_browser *Browser, _library *Library, _document *Document, _synchronous bool) (*Error) {
	
	if _browser.ServerUrlBase == "" {
		return errorw (0x2de8cc31, nil)
	}
	
	_url := "/d/" + _document.Identifier
	_url = strings.TrimRight (_browser.ServerUrlBase, "/") + _url
	
	return browserUrlOpen (_browser, _url, true, _synchronous)
}


func BrowserLibraryOpen (_browser *Browser, _library *Library, _synchronous bool) (*Error) {
	
	if _browser.ServerUrlBase == "" {
		return errorw (0x9f457963, nil)
	}
	
	_url := "/l/" + _library.Identifier
	_url = strings.TrimRight (_browser.ServerUrlBase, "/") + _url
	
	return browserUrlOpen (_browser, _url, true, _synchronous)
}


func BrowserIndexOpen (_browser *Browser, _synchronous bool) (*Error) {
	
	if _browser.ServerUrlBase == "" {
		return errorw (0x82107e2f, nil)
	}
	
	_url := "/i/"
	_url = strings.TrimRight (_browser.ServerUrlBase, "/") + _url
	
	return browserUrlOpen (_browser, _url, true, _synchronous)
}


func BrowserUrlExternalOpen (_browser *Browser, _url string, _synchronous bool) (*Error) {
	return browserUrlOpen (_browser, _url, false, _synchronous)
}


func browserUrlOpen (_browser *Browser, _url string, _internal bool, _synchronous bool) (*Error) {
	
	if _url == "" {
		return errorw (0x61668460, nil)
	}
	
	_globals := _browser.globals
	
	_command, _terminal, _error := BrowserResolveOpenCommand (_browser, _internal)
	if _error != nil {
		return _error
	}
	
	_argumentPathReplaced := false
	for _argumentIndex, _argument := range _command.Args {
		if _argument == "{{url}}" {
			_command.Args[_argumentIndex] = _url
			_argumentPathReplaced = true
		} else if strings.Contains (_argument, "{{url}}") {
			_command.Args[_argumentIndex] = strings.ReplaceAll (_argument, "{{url}}", _url)
		}
	}
	if !_argumentPathReplaced {
		return errorw (0x0f74988a, nil)
	}
	
	_execute := func () (*Error) {
			
			if _terminal {
				if ! _globals.TerminalMutexTryLock () {
					return errorw (0x9db394bd, nil)
				}
				defer _globals.TerminalMutexUnlock ()
			}
			
			if _error := _command.Start (); _error != nil {
				return errorw (0xe733ede0, _error)
			}
			
			if _error := _command.Wait (); _error != nil {
				return errorw (0xf75e37d1, _error)
			}
			
			return nil
		}
	
	if !_synchronous {
		
		go func () () {
				if _error := _execute (); _error != nil {
					logErrorf ('e', 0x9fef399a, _error, "[browser]  failed for `%s`!", _url)
				}
			} ()
		
		return nil
		
	} else {
		
		return _execute ()
	}
}



func BrowserResolveOpenCommand (_browser *Browser, _internal bool) (*exec.Cmd, bool, *Error) {
	
	_globals := _browser.globals
	
	_executable := ""
	_executableName := ""
	_executableArguments := []string (nil)
	_executableTty := (*bool) (nil)
	_true := true
	_false := false
	
	_terminalOpenCommand := []string (nil)
	_xorgOpenCommand := []string (nil)
	
	if _internal {
		_terminalOpenCommand = _browser.TerminalOpenInternalCommand
		_xorgOpenCommand = _browser.XorgOpenInternalCommand
	} else {
		_terminalOpenCommand = _browser.TerminalOpenExternalCommand
		_xorgOpenCommand = _browser.XorgOpenExternalCommand
	}
	
	if (_executable == "") && _globals.TerminalEnabled && (len (_terminalOpenCommand) > 0) {
		_executableName_0 := _terminalOpenCommand[0]
		if _executableName_0 == "" {
			return nil, false, errorw (0x4c4f8738, nil)
		}
		if _executable_0, _error := exec.LookPath (_executableName_0); _error == nil {
			_executable = _executable_0
			_executableName = _executableName_0
			_executableArguments = _terminalOpenCommand[1:]
			if len (_executableArguments) == 0 {
				_executableArguments = nil
			}
			_executableTty = &_true
		} else {
			return nil, false, errorw (0xefb85c86, _error)
		}
	}
	
	if (_executable == "") && _globals.XorgEnabled && (len (_xorgOpenCommand) > 0) {
		_executableName_0 := _xorgOpenCommand[0]
		if _executableName_0 == "" {
			return nil, false, errorw (0x6ff9fcad, nil)
		}
		if _executable_0, _error := exec.LookPath (_executableName_0); _error == nil {
			_executable = _executable_0
			_executableName = _executableName_0
			_executableArguments = _xorgOpenCommand[1:]
			if len (_executableArguments) == 0 {
				_executableArguments = nil
			}
			_executableTty = &_false
		} else {
			return nil, false, errorw (0xa5aee6da, _error)
		}
	}
	
	if (_executable == "") && (_globals.TerminalEnabled || _globals.XorgEnabled) {
		if _executableName_0, _ := _globals.Environment["BROWSER"]; _executableName_0 != "" {
			if _executable_0, _error := exec.LookPath (_executableName_0); _error == nil {
				_executable = _executable_0
				_executableName = _executableName_0
			} else {
				return nil, false, errorw (0x45869ddb, _error)
			}
		}
	}
	
	if (_executable == "") && (_globals.TerminalEnabled || _globals.XorgEnabled) {
		_alternatives := make ([]string, 0, 128)
		_alternatives = append (_alternatives, "z-scratchpad--browser")
		if _globals.TerminalEnabled {
			_alternatives = append (_alternatives, "www-browser")
		}
		if _globals.XorgEnabled {
			_alternatives = append (_alternatives, "x-www-browser", "xdg-open")
		}
		for _, _executableName_0 := range _alternatives {
			if _executable_0, _error := exec.LookPath (_executableName_0); _error == nil {
				_executable = _executable_0
				_executableName = _executableName_0
				break
			}
		}
	}
	
	if _executable == "" {
		return nil, false, errorw (0x6a1fe47a, nil)
	}
	
	if _executableArguments == nil {
		switch _executableName {
			case "z-scratchpad--browser" :
				_executableArguments = []string { "{{url}}" }
			case "www-browser" :
				_executableArguments = []string { "{{url}}" }
			case "x-www-browser", "xdg-open" :
				_executableArguments = []string { "{{url}}" }
			case "firefox" :
				_executableArguments = []string { "--", "{{url}}" }
			case "chrome", "chromium" :
				_executableArguments = []string { "--", "{{url}}" }
			case "lynx" :
				_executableArguments = []string { "--", "{{url}}" }
			case "w3m", "links", "elinks" :
				_executableArguments = []string { "{{url}}" }
			default :
				_executableArguments = []string { "{{url}}" }
		}
	}
	
	if _executableTty == nil {
		switch _executableName {
			case "z-scratchpad--browser" :
				_executableTty = &_globals.TerminalEnabled
			case "www-browser" :
				_executableTty = &_globals.TerminalEnabled
			case "x-www-browser", "xdg-open" :
				_executableTty = &_false
			case "firefox" :
				_executableTty = &_false
			case "chrome", "chromium" :
				_executableTty = &_false
			case "lynx" :
				_executableTty = &_globals.TerminalEnabled
			case "w3m", "links", "elinks" :
				_executableTty = &_globals.TerminalEnabled
			default :
				_executableTty = &_globals.TerminalEnabled
		}
	}
	
	_arguments := make ([]string, 0, 1 + len (_executableArguments))
	_arguments = append (_arguments, _executable)
	_arguments = append (_arguments, _executableArguments ...)
	
	_command := & exec.Cmd {
			Path : _executable,
			Args : _arguments,
			Env : _globals.EnvironmentList,
		}
	
	if *_executableTty {
		_command.Stdin = _globals.TerminalTty
		_command.Stdout = _globals.TerminalTty
		_command.Stderr = _globals.TerminalTty
	} else {
		_command.Stdin = _globals.DevNull
		_command.Stdout = _globals.DevNull
		_command.Stderr = _globals.DevNull
	}
	
	return _command, *_executableTty, nil
}

