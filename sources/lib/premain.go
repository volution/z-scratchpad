

package zscratchpad


import "os"
import "log"
import "strings"




func PreMain () () {
	
	
	log.SetFlags (0)
	
	
	var _executable string
	if _executable_0, _error := os.Executable (); _error == nil {
		_executable = _executable_0
	} else {
		panic (abortErrorw (0x45343ff0, _error))
	}
	
	_arguments := append ([]string (nil), os.Args[1:] ...)
	
	_environment := make (map[string]string, 128)
	for _, _variable := range os.Environ () {
		if _splitIndex := strings.IndexByte (_variable, '='); _splitIndex >= 0 {
			_name := _variable[:_splitIndex]
			_value := _variable[_splitIndex + 1:]
			_environment[_name] = _value
		} else {
			logf ('w', 0x7bb25433, "invalid environment variable (missing `=`):  `%s`", _variable)
		}
	}
	
	
	if _error := Main (_executable, _arguments, _environment); _error == nil {
		os.Exit (0)
		panic (abortErrorw (0x0cd72a12, nil))
	} else {
		panic (abortError (_error))
	}
}

