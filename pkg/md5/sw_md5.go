package md5

import (
	"crypto/md5"
	"encoding/hex"
)

func Encrypt32(v string) string {
	hash := md5.New()
	hash.Write([]byte(v))
	byts := hash.Sum(nil)
	return hex.EncodeToString(byts)
}
