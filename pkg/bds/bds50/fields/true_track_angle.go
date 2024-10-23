package fields

import (
	"github.com/twuillemin/modes/pkg/bitutils"
)

func ReadTrueTrackAngle(data []byte) (bool, float32) {
	status := (data[1] & 0x10) != 0
	angle_sign := data[1] & 0x08 >> 3

	byte1 := data[1] & 0x07
	byte2 := data[2] & 0xFE
	allBits := float32(bitutils.Pack2Bytes(byte1, byte2) >> 1)

	if angle_sign == 1 {
		allBits = allBits - 1024
	}

	angle := allBits * 90 / 512.0

	if angle < 0 {
		angle = 360 + angle
	}

	return status, angle
}
