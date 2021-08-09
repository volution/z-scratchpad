

package zscratchpad


import "os"
import "sync"


import "github.com/mattn/go-isatty"




type Globals struct {
	
	Mutex *sync.Mutex
	
	Stdin *os.File
	Stdout *os.File
	Stderr *os.File
	
	StdinIsTty bool
	StdoutIsTty bool
	StderrIsTty bool
	StdioIsTty bool
	
	TerminalAvailable bool
	TerminalType string
	TerminalTty *os.File
	
	XorgAvailable bool
	
	Executable string
	Environment map[string]string
	EnvironmentList []string
	
	DevNull *os.File
}




func GlobalsNew (_executable string, _environment map[string]string) (*Globals, *Error) {
	
	_globals := & Globals {
			
			Mutex : & sync.Mutex {},
			
			Stdin : os.Stdin,
			Stdout : os.Stdout,
			Stderr : os.Stderr,
			
			Executable : _executable,
			Environment : _environment,
		}
	
	_globals.StdinIsTty = isTerminal (_globals.Stdin)
	_globals.StdoutIsTty = isTerminal (_globals.Stdout)
	_globals.StderrIsTty = isTerminal (_globals.Stderr)
	_globals.StdioIsTty = _globals.StdinIsTty && _globals.StdoutIsTty && _globals.StderrIsTty
	
	// NOTE:  Setting these on false would disable usage of terminal / Xorg.
	_globals.TerminalAvailable = true
	_globals.XorgAvailable = true
	
	if _globals.TerminalAvailable {
		switch _type, _ := _globals.Environment["TERM"]; _type {
			case "", "dumb" :
				_globals.TerminalAvailable = false
			default :
				_globals.TerminalAvailable = true
				_globals.TerminalType = _type
		}
		if _globals.TerminalAvailable {
			if _globals.StderrIsTty {
				_globals.TerminalTty = _globals.Stderr
			} else {
				_globals.TerminalAvailable = false
			}
		}
	}
	if !_globals.TerminalAvailable {
		_globals.TerminalType = "dumb"
		_globals.Environment["TERM"] = "dumb"
	}
	
	if _globals.XorgAvailable {
		if _display, _ := _globals.Environment["DISPLAY"]; _display != "" {
			_globals.XorgAvailable = true
		} else {
			_globals.XorgAvailable = false
		}
	}
	if !_globals.XorgAvailable {
		delete (_globals.Environment, "DISPLAY")
	}
	
	if _file, _error := os.OpenFile (os.DevNull, os.O_RDWR, 0); _error == nil {
		_globals.DevNull = _file
	} else {
		return nil, errorw (0x70754895, _error)
	}
	
	_environmentList := make ([]string, 0, len (_globals.Environment))
	for _name, _value := range _globals.Environment {
		if (_name == "") || (_value == "") {
			// FIXME:  We should issue a warning in this case!
			delete (_globals.Environment, _name)
			continue
		}
		_environmentVariable := _name + "=" + _value
		_environmentList = append (_environmentList, _environmentVariable)
	}
	_globals.EnvironmentList = _environmentList
	
	return _globals, nil
}




func isTerminal (_file *os.File) (bool) {
	_descriptor := _file.Fd ()
	return isatty.IsTerminal (_descriptor) || isatty.IsCygwinTerminal (_descriptor)
}

