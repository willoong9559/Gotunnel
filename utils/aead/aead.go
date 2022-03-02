package aead

import (
	"crypto/aes"
	"crypto/cipher"

	"golang.org/x/crypto/argon2"
)

type Cipher interface {
	KeySize() int
	SaltSize() int
	Encrypter(salt []byte) (cipher.AEAD, error)
	//	Decrypter(salt []byte) (cipher.AEAD, error)
}

type aeadAes128GcmCipher struct {
	psk      []byte
	keySize  int
	makeAEAD func(key []byte) (cipher.AEAD, error)
}

func (s *aeadAes128GcmCipher) KeySize() int  { return s.keySize }
func (s *aeadAes128GcmCipher) SaltSize() int { return 16 }
func (s *aeadAes128GcmCipher) Encrypter(salt []byte) (cipher.AEAD, error) {
	return s.makeAEAD(snellKDF(s.psk, salt, s.KeySize()))
}

// func (s *aeadAes128GcmCipher) Decrypter(salt []byte) (cipher.AEAD, error) {
// 	return s.makeAEAD(aeadAes128GcmCipher(s.psk, salt, s.KeySize()))
// }

func snellKDF(psk, salt []byte, keySize int) []byte {
	return argon2.IDKey(psk, salt, 3, 8, 1, 32)[:keySize]
}

func aesGCM(key []byte) (cipher.AEAD, error) {
	blk, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	return cipher.NewGCM(blk)
}

func NewPskCipher(psk []byte) Cipher {
	return &aeadAes128GcmCipher{
		psk:      psk,
		keySize:  16,
		makeAEAD: aesGCM,
	}
}
