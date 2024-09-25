package des

import "crypto/cipher"

type ecb struct {
	b         cipher.Block
	blockSize int
}

func NewECBCipher(b cipher.Block) cipher.BlockMode {
	return &ecb{b: b, blockSize: b.BlockSize()}
}

func (x *ecb) BlockSize() int { return x.blockSize }

func (x *ecb) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("input not full blocks")
	}
	for len(src) > 0 {
		x.b.Encrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}
