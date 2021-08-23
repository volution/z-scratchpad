

package zscratchpad


import "crypto/hmac"
import crand "crypto/rand"
import "encoding/hex"
import "encoding/base64"
import "hash"
import "runtime"
import "time"


import "github.com/zeebo/blake3"




func fingerprintString (_value string) (string) {
	_fingerprint := fingerprintBytes (StringToBytes (_value))
	runtime.KeepAlive (_value)
	return _fingerprint
}

func fingerprintBytes (_value []byte) (string) {
	_hasher_0 := blacke3HasherPrototype
	_hasher := &_hasher_0
	_hasher.Write (_value)
	return fingerprintFinalize (_hasher)
}


func fingerprintStringLines (_values []string) (string) {
	_hasher_0 := blacke3HasherPrototype
	_hasher := &_hasher_0
	_separator := []byte { '\n' }
	for _, _value := range _values {
		_hasher.Write (StringToBytes (_value))
		runtime.KeepAlive (_value)
		_hasher.Write (_separator)
	}
	return fingerprintFinalize (_hasher)
}


func fingerprintFinalize (_hasher *blake3.Hasher) (string) {
	if _hasher.Size () != 32 {
		panic (0x4ac694be)
	}
	var _data_0 [32]byte
	_data := NoEscapeBytes (_data_0[:])
	_hasher.Sum (_data[:0])
	_fingerprint := hex.EncodeToString (_data[:16])
	runtime.KeepAlive (_data_0)
	return _fingerprint
}

var blacke3HasherPrototype blake3.Hasher = * blake3.New ()




func generateRandomToken () (string) {
	var _data_0 [128 / 8]byte
	_data := NoEscapeBytes (_data_0[:])
	if _read, _error := crand.Read (_data); _error == nil {
		if _read != (128 / 8) {
			panic (0x57a9916a)
		}
	} else {
		panic (0xb6800787)
	}
	_token := hex.EncodeToString (_data)
	runtime.KeepAlive (_data_0)
	return _token
}




func generateHmac (_secret string, _payload string) (string, *Error) {
	_timestamp := time.Now ()
	return generateHmac_0 (_secret, _payload, _timestamp)
}


func generateHmac_0 (_secret string, _payload string, _timestamp time.Time) (string, *Error) {
	
	if len (_secret) < 32 {
		return "", errorw (0x39641c54, nil)
	}
	if len (_payload) > 512 {
		return "", errorw (0x5efe0ea5, nil)
	}
	
	_secretBytes := StringToBytes (_secret)
	_payloadBytes := StringToBytes (_payload)
	
	_hasher := hmac.New (func () (hash.Hash) { return blake3.New () }, _secretBytes)
	
	_timestampBytes, _error := _timestamp.MarshalBinary ()
	if _error != nil {
		return "", errorw (0xb4c12b8f, _error)
	}
	if len (_timestampBytes) != 15 {
		return "", errorw (0x694ea80b, nil)
	}
	
	_hasher.Write (_payloadBytes)
	_hasher.Write (_timestampBytes)
	
	var _macBytes_0 [32]byte
	_macBytes := NoEscapeBytes (_macBytes_0[:])
	_macBytes = _hasher.Sum (_macBytes[:0])
	
	_buffer := BytesBufferNewSize (128)
	defer BytesBufferRelease (_buffer)
	
	_buffer.Write (_macBytes)
	_buffer.Write (_timestampBytes)
	
	runtime.KeepAlive (_secret)
	runtime.KeepAlive (_payload)
	runtime.KeepAlive (_macBytes_0)
	
//	_macText := hex.EncodeToString (_buffer.Bytes ())
	_macText := base64.RawURLEncoding.EncodeToString (_buffer.Bytes ())
	
	return _macText, nil
}


func verifyHmac (_secret string, _payload string, _mac string, _timeoutMilliseconds uint) (*Error) {
	
	_macText := _mac
	
	_macTextBytes := StringToBytes (_macText)
	
	_buffer := BytesBufferNewSize (128)
	defer BytesBufferRelease (_buffer)
	
	_bufferBytes := _buffer.Bytes () [:128]
//	if _size, _error := hex.Decode (_bufferBytes, _macTextBytes); _error == nil {
	if _size, _error := base64.RawURLEncoding.Decode (_bufferBytes, _macTextBytes); _error == nil {
		_bufferBytes = _bufferBytes[:_size]
	} else {
		return errorw (0xdf4124d0, _error)
	}
	
	if len (_bufferBytes) != (32 + 15) {
		return errorw (0xa8327721, nil)
	}
	
	_timestamp := time.Time {}
	if _error := _timestamp.UnmarshalBinary (_bufferBytes[32:]); _error != nil {
		return errorw (0x7eaae6b0, _error)
	}
	
	_elapsedMilliseconds := time.Since (_timestamp) .Milliseconds ()
	if _elapsedMilliseconds < 0 {
		return errorw (0x62fb6abb, nil)
	}
	if _elapsedMilliseconds >= int64 (_timeoutMilliseconds) {
		return errorw (0x13c6f020, nil)
	}
	
	if _macTextRecomputed, _error := generateHmac_0 (_secret, _payload, _timestamp); _error == nil {
		_macTextRecomputedBytes := StringToBytes (_macTextRecomputed)
		if ! hmac.Equal (_macTextBytes, _macTextRecomputedBytes) {
			return errorw (0x20018e92, nil)
		}
		runtime.KeepAlive (_macTextRecomputed)
	} else {
		return _error
	}
	
	runtime.KeepAlive (_macText)
	
	return nil
}

