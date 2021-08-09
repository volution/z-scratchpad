

package zscratchpad


import "os"


import "github.com/mattn/go-isatty"
import "github.com/subchen/go-trylock/v2"




type Globals struct {
	
	mutex trylock.TryLocker
	
	Stdin *os.File
	Stdout *os.File
	Stderr *os.File
	
	StdinIsTty bool
	StdoutIsTty bool
	StderrIsTty bool
	StdioIsTty bool
	
	TerminalAvailable bool
	TerminalEnabled bool
	TerminalType string
	TerminalTty *os.File
	terminalMutex trylock.TryLocker
	
	XorgAvailable bool
	XorgEnabled bool
	
	Executable string
	Environment map[string]string
	EnvironmentList []string
	
	DevNull *os.File
}




func GlobalsNew (_executable string, _environment map[string]string) (*Globals, *Error) {
	
	_globals := & Globals {
			
			mutex : trylock.New (),
			terminalMutex : trylock.New (),
			
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
	
	_globals.TerminalEnabled = true
	_globals.XorgEnabled = true
	
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
		_globals.TerminalEnabled = false
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
		_globals.XorgEnabled = false
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




func (_globals *Globals) MutexLock () () {
	_globals.mutex.Lock ()
}

func (_globals *Globals) MutexUnlock () () {
	_globals.mutex.Unlock ()
}

func (_globals *Globals) MutexTryLock () (bool) {
	return _globals.mutex.TryLock (nil)
}




func (_globals *Globals) TerminalMutexLock () () {
	_globals.terminalMutex.Lock ()
}

func (_globals *Globals) TerminalMutexUnlock () () {
	_globals.terminalMutex.Unlock ()
}

func (_globals *Globals) TerminalMutexTryLock () (bool) {
	return _globals.terminalMutex.TryLock (nil)
}




func isTerminal (_file *os.File) (bool) {
	_descriptor := _file.Fd ()
	return isatty.IsTerminal (_descriptor) || isatty.IsCygwinTerminal (_descriptor)
}

