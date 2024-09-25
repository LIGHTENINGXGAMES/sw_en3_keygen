package buf

import (
	"bytes"
	"encoding/binary"
)

type Buffer bytes.Buffer

func (b *Buffer) Buffer() *bytes.Buffer { return (*bytes.Buffer)(b) }

func (b *Buffer) WriteInt32(num int32) {
	binary.Write(b.Buffer(), binary.LittleEndian, &num)
}

func (b *Buffer) WriteBytes(bytes []byte) {
	length := int32(len(bytes))
	binary.Write(b.Buffer(), binary.LittleEndian, &length)
	binary.Write(b.Buffer(), binary.LittleEndian, bytes)
}

func (b *Buffer) WriteString(str string) {
	length := int32(len(str))
	binary.Write(b.Buffer(), binary.LittleEndian, &length)
	binary.Write(b.Buffer(), binary.LittleEndian, []byte(str))
}

func (b *Buffer) WriteRaw(bytes []byte) {
	binary.Write(b.Buffer(), binary.LittleEndian, bytes)
}

func (b *Buffer) Bytes() []byte { return b.Buffer().Bytes() }
