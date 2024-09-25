package des

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
)

func Encrypt(dataToEncrypt []byte, key string, isPaddingZeros bool, isECBMode bool) ([]byte, error) {
	block, err := des.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}

	if isPaddingZeros {
		dataToEncrypt = padZeros(dataToEncrypt, block.BlockSize())
	}

	var blockMode cipher.BlockMode
	if isECBMode {
		blockMode = NewECBCipher(block)
	} else {
		iv := []byte(SW_DES_IV)
		blockMode = cipher.NewCBCEncrypter(block, iv)
	}

	encrypted := make([]byte, len(dataToEncrypt))
	blockMode.CryptBlocks(encrypted, dataToEncrypt)

	return encrypted, nil
}

func padZeros(data []byte, blockSize int) []byte {
	paddingSize := blockSize - (len(data) % blockSize)
	if paddingSize == blockSize {
		return data
	}
	return append(data, bytes.Repeat([]byte{0}, paddingSize)...)
}
