package fields

import (
	"github.com/twuillemin/modes/pkg/bitutils"
)

func ReadMagneticHeading(data []byte) (bool, float32) {
	status := (data[0] & 0x80) != 0
	heading_sign := data[0] & 0x40 >> 6

	byte1 := data[0] & 0x3F
	byte2 := data[1] & 0xF0
	allBits := float32(bitutils.Pack2Bytes(byte1, byte2) >> 4)

	if heading_sign == 1 {
		allBits = allBits - 1024
	}

	magnetic_heading := allBits * 90 / 512

	if magnetic_heading < 0 {
		magnetic_heading = 360 + magnetic_heading
	}

	return status, magnetic_heading
}
