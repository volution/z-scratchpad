

package zscratchpad


import "crypto/sha1"
import crand "crypto/rand"
import "encoding/hex"
import "hash"
import "runtime"




func fingerprintString (_value string) (string) {
	_fingerprint := fingerprintBytes (StringToBytes (_value))
	runtime.KeepAlive (_value)
	return _fingerprint
}

func fingerprintBytes (_value []byte) (string) {
	_hasher := sha1.New ()
	_hasher.Write (_value)
	return fingerprintFinalize (_hasher)
}


func fingerprintStringLines (_values []string) (string) {
	_hasher := sha1.New ()
	_separator := []byte { '\n' }
	for _, _value := range _values {
		_hasher.Write (StringToBytes (_value))
		runtime.KeepAlive (_value)
		_hasher.Write (_separator)
	}
	return fingerprintFinalize (_hasher)
}


func fingerprintFinalize (_hasher hash.Hash) (string) {
	var _data_0 [sha1.Size]byte
	_data := NoEscapeBytes (_data_0[:])
	_hasher.Sum (_data[:0])
	_fingerprint := hex.EncodeToString (_data)
	runtime.KeepAlive (_data_0)
	return _fingerprint
}




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

