

package zscratchpad


import "io"


import "github.com/mdp/qrterminal"




func QrcodeTerminalDisplay (_data string, _stream io.Writer) (*Error) {
	
	_buffer := BytesBufferNewSize (16 * 1024)
	defer BytesBufferRelease (_buffer)
	
	qrterminal.GenerateHalfBlock (_data, qrterminal.L, _buffer)
	
	if _, _error := _buffer.WriteTo (_stream); _error != nil {
		return errorw (0xb64c335f, _error)
	}
	
	return nil
}

