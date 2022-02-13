

package zscratchpad


import "bytes"
import "fmt"
import "os"
import "log"
import "path/filepath"
import "runtime"
import "runtime/debug"
import "strings"


import . "github.com/volution/z-scratchpad/embedded"




func PreMain () () {
	
	
	runtime.GOMAXPROCS (1)
	debug.SetMaxThreads (128)
	debug.SetGCPercent (-1)
	
	
	log.SetFlags (0)
	
	
	var _executable0 string
	if _executable0_0, _error := os.Executable (); _error == nil {
		_executable0 = _executable0_0
	} else {
		panic (abortErrorw (0x45343ff0, _error))
	}
	
	if _executable0 == "" {
		panic (abortErrorw (0xaa73ab81, nil))
	}
	
	var _executable string
	if _executable_0, _error := filepath.EvalSymlinks (_executable0); _error == nil {
		_executable = _executable_0
	} else {
		panic (abortErrorw (0x7c18ca77, _error))
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
		
		_executableName := "z-scratchpad"
		
		if (_argument == "--version") || (_argument == "-v") {
			
			_buffer := bytes.NewBuffer (nil)
			fmt.Fprintf (_buffer, "* tool          : %s\n", _executableName)
			fmt.Fprintf (_buffer, "* version       : %s\n", BUILD_VERSION)
			if _executable0 == _executable {
				fmt.Fprintf (_buffer, "* executable    : %s\n", _executable)
			} else {
				fmt.Fprintf (_buffer, "* executable    : %s\n", _executable)
				fmt.Fprintf (_buffer, "* executable-0  : %s\n", _executable0)
			}
			fmt.Fprintf (_buffer, "* build target  : %s, %s-%s, %s, %s\n", BUILD_TARGET, BUILD_TARGET_OS, BUILD_TARGET_ARCH, BUILD_COMPILER_VERSION, BUILD_COMPILER_TYPE)
			fmt.Fprintf (_buffer, "* build number  : %s, %s\n", BUILD_NUMBER, BUILD_TIMESTAMP)
			fmt.Fprintf (_buffer, "* code & issues : %s\n", PROJECT_URL)
			fmt.Fprintf (_buffer, "* sources git   : %s\n", BUILD_GIT_HASH)
			fmt.Fprintf (_buffer, "* sources hash  : %s\n", BUILD_SOURCES_HASH)
			fmt.Fprintf (_buffer, "* uname node    : %s\n", UNAME_NODE)
			fmt.Fprintf (_buffer, "* uname system  : %s, %s, %s\n", UNAME_SYSTEM, UNAME_RELEASE, UNAME_MACHINE)
			fmt.Fprintf (_buffer, "* uname hash    : %s\n", UNAME_FINGERPRINT)
			if _, _error := _buffer.WriteTo (os.Stdout); _error != nil {
				panic (abortErrorw (0x26567db5, _error))
			}
			os.Exit (0)
			panic (abortErrorw (0xd80435fe, nil))
		}
		
		if _argument == "--sources-md5" {
			if _, _error := os.Stdout.WriteString (BuildSourcesMd5); _error != nil {
				panic (abortErrorw (0x6c4a7eb3, _error))
			}
			os.Exit (0)
			panic (abortErrorw (0xbe9f73ad, nil))
		}
		
		if _argument == "--sources-cpio" {
			if _, _error := os.Stdout.Write (BuildSourcesCpioGz); _error != nil {
				panic (abortErrorw (0xf5b73ab8, _error))
			}
			os.Exit (0)
			panic (abortErrorw (0x97f9faf1, nil))
		}
		
		if (_argument == "-h") ||
				(_argument == "--help") ||
				(_argument == "--manual") ||
				(_argument == "--manual-text") || (_argument == "--manual-txt") ||
				(_argument == "--manual-html") ||
				(_argument == "--manual-man") ||
				(_argument == "--readme") ||
				(_argument == "--readme-text") || (_argument == "--readme-txt") ||
				(_argument == "--readme-html") {
			_replacements := map[string]string {
					"@{PROJECT_URL}" : PROJECT_URL,
					"@{BUILD_TARGET}" : BUILD_TARGET,
					"@{BUILD_TARGET_ARCH}" : BUILD_TARGET_ARCH,
					"@{BUILD_TARGET_OS}" : BUILD_TARGET_OS,
					"@{BUILD_COMPILER_TYPE}" : BUILD_COMPILER_TYPE,
					"@{BUILD_COMPILER_VERSION}" : BUILD_COMPILER_VERSION,
					"@{BUILD_DEVELOPMENT}" : fmt.Sprintf ("%s", BUILD_DEVELOPMENT),
					"@{BUILD_VERSION}" : BUILD_VERSION,
					"@{BUILD_NUMBER}" : BUILD_NUMBER,
					"@{BUILD_TIMESTAMP}" : BUILD_TIMESTAMP,
					"@{BUILD_GIT_HASH}" : BUILD_GIT_HASH,
					"@{BUILD_SOURCES_HASH}" : BUILD_SOURCES_HASH,
					"@{UNAME_NODE}" : UNAME_NODE,
					"@{UNAME_SYSTEM}" : UNAME_SYSTEM,
					"@{UNAME_RELEASE}" : UNAME_RELEASE,
					"@{UNAME_VERSION}" : UNAME_VERSION,
					"@{UNAME_MACHINE}" : UNAME_MACHINE,
					"@{UNAME_FINGERPRINT}" : UNAME_FINGERPRINT,
				}
			_manual := ""
			_useType := ""
			_useDecorations := false
			switch _argument {
				case "--help", "-h" :
					_manual = ZscratchpadHelpTxt
					_useType = "text"
					_useDecorations = true
				case "--manual", "--manual-text", "--manual-txt" :
					_manual = ZscratchpadManualTxt
					_useType = "text"
					_useDecorations = true
				case "--manual-html" :
					_manual = ZscratchpadManualHtml
					_useType = "html"
				case "--manual-man" :
					_manual = ZscratchpadManualMan
					_useType = "man"
				case "--readme", "--readme-text", "--readme-txt" :
					_manual = ReadmeTxt
					_useType = "text"
					_useDecorations = true
				case "--readme-html" :
					_manual = ReadmeHtml
					_useType = "html"
				default :
					panic (0x795bbfd9)
			}
			if _manual != "__custom__" {
				_chunks := make ([]string, 0, 8)
				if _useDecorations {
					_chunks = append (_chunks, HelpHeader, _manual, HelpFooter)
				} else {
					_chunks = append (_chunks, _manual)
				}
				for _index := range _chunks {
					for _key, _replacement := range _replacements {
						_chunks[_index] = strings.ReplaceAll (_chunks[_index], _key, _replacement)
					}
				}
				_buffer := bytes.NewBuffer (nil)
				for _, _chunk := range _chunks {
					_buffer.WriteString (_chunk)
				}
				// FIXME:  Use a pager like with `z-run`!
				_ = _useType
				if _, _error := _buffer.WriteTo (os.Stdout); _error != nil {
					panic (abortErrorw (0x5a1741cc, _error))
				}
				os.Exit (0)
				panic (0x25d33d9b)
			}
		}
	}
	
	
	
	
	if _error := Main (_executable, _arguments, _environment); _error == nil {
		os.Exit (0)
		panic (abortErrorw (0x0cd72a12, nil))
	} else {
		panic (abortError (_error))
	}
}

