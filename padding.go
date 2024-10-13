package aesnet256

import (
	"bytes"
)

// PadKey will pad the key appropriately.
// keys > size will be truncated per OpenSSL's spec
func PadKey(key []byte, size int) (pk []byte) {
	if len(key) > size {
		pk = key[0:size]
		return
	}
	ps := size - len(key)
	pad := make([]byte, ps)
	pk = append(key, pad...)
	return
}

// PKCS5Unpad will unpad the byte slice.
// If the padding is invalid or zero it will return an error
// typically this means the key supplied was incorrect.
func PKCS5Unpad(i []byte) (o []byte, e error) {
	if len(i) == 0 {
		e = E_unpad_zero
		return
	}
	pad := int(i[len(i)-1])
	if pad > len(i) {
		e = E_invalid_pad
		return
	}
	o = i[:len(i)-pad]
	return
}

// PKCS5Pad will pad the byte slice according to the given block size.
func PKCS5Pad(i []byte, blockSize int) (o []byte) {
	pad := blockSize - len(i)%blockSize
	padBytes := bytes.Repeat([]byte{byte(pad)}, pad)
	o = append(i, padBytes...)
	return
}
