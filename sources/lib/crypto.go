

package zscratchpad


import "crypto/sha256"
import "encoding/hex"




func fingerprintString (_value string) (string) {
	return fingerprintBytes ([]byte (_value))
}

func fingerprintBytes (_value []byte) (string) {
	_hasher := sha256.New ()
	_hasher.Write (_value)
	return hex.EncodeToString (_hasher.Sum (nil))
}

