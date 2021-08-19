

package zscratchpad


import "bytes"
import "unsafe"
import "reflect"
import "sync"




//go:nosplit
func NoEscapePointer (p unsafe.Pointer) (unsafe.Pointer) {
	x := uintptr (p)
	return unsafe.Pointer (x ^ 0)
}


func NoEscapeBytesPointer (p *[]byte) (*[]byte) {
	return (*[]byte) (NoEscapePointer (unsafe.Pointer (p)))
}

func NoEscapeStringPointer (p *string) (*string) {
	return (*string) (NoEscapePointer (unsafe.Pointer (p)))
}


func NoEscapeBytes (p []byte) ([]byte) {
	return * NoEscapeBytesPointer (&p)
}

func NoEscapeString (p string) (string) {
	return * NoEscapeStringPointer (&p)
}




func BytesToString (_bytes []byte) (string) {
	return *(*string) (unsafe.Pointer (&_bytes))
}

func StringToBytes (_string string) ([]byte) {
	
	// NOTE:  The following is broken!
	// return *(*[]byte) (unsafe.Pointer (&_string))
	
	// NOTE:  Based on `https://github.com/valyala/fasthttp/blob/2a6f7db5bbc4d7c11f1ccc0cb827e145b9b7d7ea/bytesconv.go#L342`
	_bytes := []byte (nil)
	_bytesHeader := (*reflect.SliceHeader) (unsafe.Pointer (&_bytes))
	_stringHeader := (*reflect.StringHeader) (unsafe.Pointer (&_string))
	_bytesHeader.Data = _stringHeader.Data
	_bytesHeader.Len = _stringHeader.Len
	_bytesHeader.Cap = _stringHeader.Len
	return _bytes
}




func BytesBufferNewSize (_size int) (*bytes.Buffer) {
	if _size > (128 * (1 << 15)) {
		return bytes.NewBuffer (make ([]byte, 0, _size))
	}
	_poolIndex, _poolSize := bufferPoolIndex (_size)
//	logf ('d', 0x3fd6e577, "%d %d %d", _poolIndex, _poolSize, _size)
	_buffer_0 := bufferPools[_poolIndex].Get ()
	if _buffer_0 == nil {
		return bytes.NewBuffer (make ([]byte, 0, _poolSize))
	}
	_buffer := _buffer_0.(*bytes.Buffer)
	return _buffer
}

func BytesBufferRelease (_buffer *bytes.Buffer) () {
	_buffer.Reset ()
	_size := _buffer.Cap ()
	if _size > (128 * (1 << 15)) {
		return
	}
	_poolIndex, _poolSize := bufferPoolIndex (_size)
//	logf ('d', 0x7708b8b2, "%d %d %d", _poolIndex, _poolSize, _size)
	if _size < _poolSize {
		_poolIndex -= 1
	}
	if _poolIndex < 0 {
		return
	}
	bufferPools[_poolIndex].Put (_buffer)
}

func bufferPoolIndex (_size int) (int, int) {
	_poolIndex := -1
	_poolSize := 0
	for _index := 0; _index <= 15; _index += 1 {
		_poolSize = (128 * (1 << _index))
		if _size <= _poolSize {
			_poolIndex = _index
			break
		}
	}
	return _poolIndex, _poolSize
}

func init () () {
	bufferPools = make ([]*sync.Pool, 17)
	for _index := range bufferPools {
		bufferPools[_index] = & sync.Pool {}
	}
}

var bufferPools []*sync.Pool

