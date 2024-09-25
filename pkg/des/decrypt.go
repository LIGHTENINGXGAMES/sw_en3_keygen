package des

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
)

func Decrypt(encryptedData []byte, key string, isPaddingZeros bool, isECBMode bool) ([]byte, error) {
	block, err := des.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}

	var blockMode cipher.BlockMode
	if isECBMode {
		blockMode = NewECBCipher(block)
	} else {
		iv := []byte(SW_DES_IV)
		blockMode = cipher.NewCBCDecrypter(block, iv)
	}

	decrypted := make([]byte, len(encryptedData))
	blockMode.CryptBlocks(decrypted, encryptedData)

	if isPaddingZeros {
		decrypted = unpadZeros(decrypted)
	}

	return decrypted, nil
}

func unpadZeros(data []byte) []byte {
	return bytes.TrimRight(data, "\x00")
}
