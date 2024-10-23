package fields

import "github.com/twuillemin/modes/pkg/bitutils"

func ReadTrackAngleRate(data []byte) (bool, float32) {
	status := (data[4] & 0x20) != 0
	track_sign := (data[4] & 0x10) >> 4

	byte1 := data[4] & 0x0F
	byte2 := data[5] & 0xF8
	allBits := float32(bitutils.Pack2Bytes(byte1, byte2) >> 3)

	if track_sign == 1 {
		allBits = allBits - 512
	}

	angleRate := allBits * 8 / 256

	return status, angleRate
}
