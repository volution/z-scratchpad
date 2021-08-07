

package zscratchpad


import "crypto/sha256"
import crand "crypto/rand"
import "encoding/hex"




func fingerprintString (_value string) (string) {
	return fingerprintBytes ([]byte (_value))
}

func fingerprintBytes (_value []byte) (string) {
	_hasher := sha256.New ()
	_hasher.Write (_value)
	return hex.EncodeToString (_hasher.Sum (nil))
}


func generateRandomToken () (string) {
	var _data [128 / 8]byte
	if _read, _error := crand.Read (_data[:]); _error == nil {
		if _read != (128 / 8) {
			panic (0x57a9916a)
		}
	} else {
		panic (0xb6800787)
	}
	_token := hex.EncodeToString (_data[:])
	return _token
}

