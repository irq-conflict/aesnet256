package aesnet256

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
)

const Aes_128_length = 16
const Aes_192_length = 24
const Aes_256_length = 32

// This is the default initilization vector from aesencryption.net
// It is the string 1234567890123456 with the following alterations:
// byte[8] ^= 0x5b
// byte[10] ^= 0x4b
// byte[15] ^= 0x58
const Default_aesnet_iv = "12345678b0z2345n"

var E_blocksize = "Source requires a block size evenly divisible by %d"
var E_unpad_zero = errors.New("Tried to unpad data of zero length.")
var E_invalid_pad = errors.New("Invalid padding.")

// Aes256, despite it's name, will allow for 128, 192 and 256.
// This is set by the length of the key, following the documentation:
// https://pkg.go.dev/crypto/aes#NewCipher
type Aes256 struct {
	Source     []byte
	IV         []byte
	Dest       []byte
	EncodeDest bool
}

// String implements stringer for Aes256
// Base64 encoded values will be displayed as such
func (a Aes256) String() (s string) {
	if a.EncodeDest {
		s = base64.StdEncoding.EncodeToString(a.Dest)
		return
	}
	s = string(a.Dest)
	return
}

// Encrypt will encrypt the Source in a to Dest
func (a *Aes256) Encrypt(key []byte) (e error) {
	a.Source = PKCS5Pad(a.Source, aes.BlockSize)
	block, e := aes.NewCipher(key)
	if e != nil {
		return
	}
	a.Dest = make([]byte, len(a.Source))
	mode := cipher.NewCBCEncrypter(block, a.IV)
	mode.CryptBlocks(a.Dest, a.Source)
	a.EncodeDest = true
	return
}

// Decrypt will decrypt the Source in a to Dest
func (a *Aes256) Decrypt(key []byte) (e error) {
	ct, e := base64.StdEncoding.DecodeString(string(a.Source))
	if e != nil {
		return
	}
	blk, e := aes.NewCipher(key)
	if e != nil {
		return
	}
	if len(ct)%aes.BlockSize != 0 {
		e = fmt.Errorf(E_blocksize, aes.BlockSize)
		return
	}
	mode := cipher.NewCBCDecrypter(blk, a.IV)
	a.Dest = make([]byte, len(ct))
	mode.CryptBlocks(a.Dest, ct)
	a.Dest, e = PKCS5Unpad(a.Dest)
	if e != nil {
		log.Printf("Unpadding error: %s", e)
	}
	return
}
