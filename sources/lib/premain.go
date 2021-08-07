

package zscratchpad


import "bytes"
import "fmt"
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
	
	
	if len (_arguments) == 1 {
		_argument := _arguments[0]
		
		if (_argument == "--version") || (_argument == "-v") {
			
			_buffer := bytes.NewBuffer (nil)
			fmt.Fprintf (_buffer, "* version       : %s\n", BUILD_VERSION)
			fmt.Fprintf (_buffer, "* executable    : %s\n", os.Args[0])
			fmt.Fprintf (_buffer, "* build target  : %s, %s-%s, %s, %s\n", BUILD_TARGET, BUILD_TARGET_OS, BUILD_TARGET_ARCH, BUILD_COMPILER_VERSION, BUILD_COMPILER_TYPE)
			fmt.Fprintf (_buffer, "* build number  : %s, %s\n", BUILD_NUMBER, BUILD_TIMESTAMP)
			fmt.Fprintf (_buffer, "* sources md5   : %s\n", BUILD_SOURCES_MD5)
			fmt.Fprintf (_buffer, "* sources git   : %s\n", BUILD_GIT_HASH)
			fmt.Fprintf (_buffer, "* code & issues : %s\n", PROJECT_URL)
			fmt.Fprintf (_buffer, "* uname node    : %s\n", UNAME_NODE)
			fmt.Fprintf (_buffer, "* uname system  : %s, %s, %s\n", UNAME_SYSTEM, UNAME_RELEASE, UNAME_MACHINE)
			if _, _error := _buffer.WriteTo (os.Stdout); _error != nil {
				panic (abortErrorw (0x26567db5, _error))
			}
			
			os.Exit (0)
			panic (abortErrorw (0xd80435fe, nil))
		}
		
		if _argument == "--sources" {
			
			if _, _error := os.Stdout.Write (BUILD_SOURCES_CPIO); _error != nil {
				panic (abortErrorw (0xf5b73ab8, _error))
			}
			
			os.Exit (0)
			panic (abortErrorw (0x97f9faf1, nil))
		}
	}
	
	
	if _error := Main (_executable, _arguments, _environment); _error == nil {
		os.Exit (0)
		panic (abortErrorw (0x0cd72a12, nil))
	} else {
		panic (abortError (_error))
	}
}

