

package zscratchpad


import "sync"




type Globals struct {
	Mutex *sync.Mutex
}




func GlobalsNew () (*Globals, *Error) {
	_globals := & Globals {
			Mutex : & sync.Mutex {},
		}
	return _globals, nil
}

