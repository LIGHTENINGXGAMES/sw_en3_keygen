package main

import (
	"encoding/base64"
	"fmt"
	"lstudio/pkg/buf"
	"lstudio/pkg/des"
	"lstudio/pkg/md5"
	"os"
	"time"
)

const SW_DES_KEY = "LKbcInie"

// active info
const (
	ACTIVE_DAYS = 365
)

func main() {
	var buf buf.Buffer
	v, h := value()
	buf.WriteInt32(1)
	buf.WriteString("02")
	buf.WriteString(" ********* BEGIN *********")
	buf.WriteBytes(v)
	buf.WriteString("********** END **********")
	buf.WriteString(h)

	enc, _ := des.Encrypt(buf.Bytes(), SW_DES_KEY, true, true)
	enc, _ = des.Encrypt(enc, SW_DES_KEY, true, true)
	enc, _ = des.Encrypt(enc, SW_DES_KEY, true, true)

	os.WriteFile("SWAF1501.swaf", enc, 0666)
}

func value() ([]byte, string) {
	var buf buf.Buffer
	// version name
	buf.WriteString("15")
	buf.WriteString("01")
	// active info, origin: 02150102290919114801500000000
	now := time.Now()
	nowStr := fmt.Sprintf("%d%d%d", now.Year()%100, now.Month(), now.Day())
	buf.WriteString(fmt.Sprintf("02150102%s114801500000000", nowStr))
	// active day
	buf.WriteRaw([]byte(nowStr))
	// times (day)
	buf.WriteInt32(365)
	// active day
	buf.WriteRaw([]byte(nowStr))
	return buf.Buffer().Bytes(), md5.Encrypt32(base64.StdEncoding.EncodeToString(buf.Buffer().Bytes()))
}
