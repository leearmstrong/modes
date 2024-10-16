package reader

import (
	"fmt"
	"github.com/twuillemin/modes/pkg/bitutils"
	"github.com/twuillemin/modes/pkg/modes/common"
	"github.com/twuillemin/modes/pkg/modes/messages"
)

var crcTable = [256]uint32{
	0x00000000, 0x00FFF409, 0x00001C1B, 0x00FFE812, 0x00003836, 0x00FFCC3F, 0x0000242D, 0x00FFD024, 0x0000706C, 0x00FF8465, 0x00006C77, 0x00FF987E,
	0x0000485A, 0x00FFBC53, 0x00005441, 0x00FFA048, 0x0000E0D8, 0x00FF14D1, 0x0000FCC3, 0x00FF08CA, 0x0000D8EE, 0x00FF2CE7, 0x0000C4F5, 0x00FF30FC,
	0x000090B4, 0x00FF64BD, 0x00008CAF, 0x00FF78A6, 0x0000A882, 0x00FF5C8B, 0x0000B499, 0x00FF4090, 0x0001C1B0, 0x00FE35B9, 0x0001DDAB, 0x00FE29A2,
	0x0001F986, 0x00FE0D8F, 0x0001E59D, 0x00FE1194, 0x0001B1DC, 0x00FE45D5, 0x0001ADC7, 0x00FE59CE, 0x000189EA, 0x00FE7DE3, 0x000195F1, 0x00FE61F8,
	0x00012168, 0x00FED561, 0x00013D73, 0x00FEC97A, 0x0001195E, 0x00FEED57, 0x00010545, 0x00FEF14C, 0x00015104, 0x00FEA50D, 0x00014D1F, 0x00FEB916,
	0x00016932, 0x00FE9D3B, 0x00017529, 0x00FE8120, 0x00038360, 0x00FC7769, 0x00039F7B, 0x00FC6B72, 0x0003BB56, 0x00FC4F5F, 0x0003A74D, 0x00FC5344,
	0x0003F30C, 0x00FC0705, 0x0003EF17, 0x00FC1B1E, 0x0003CB3A, 0x00FC3F33, 0x0003D721, 0x00FC2328, 0x000363B8, 0x00FC97B1, 0x00037FA3, 0x00FC8BAA,
	0x00035B8E, 0x00FCAF87, 0x00034795, 0x00FCB39C, 0x000313D4, 0x00FCE7DD, 0x00030FCF, 0x00FCFBC6, 0x00032BE2, 0x00FCDFEB, 0x000337F9, 0x00FCC3F0,
	0x000242D0, 0x00FDB6D9, 0x00025ECB, 0x00FDAAC2, 0x00027AE6, 0x00FD8EEF, 0x000266FD, 0x00FD92F4, 0x000232BC, 0x00FDC6B5, 0x00022EA7, 0x00FDDAAE,
	0x00020A8A, 0x00FDFE83, 0x00021691, 0x00FDE298, 0x0002A208, 0x00FD5601, 0x0002BE13, 0x00FD4A1A, 0x00029A3E, 0x00FD6E37, 0x00028625, 0x00FD722C,
	0x0002D264, 0x00FD266D, 0x0002CE7F, 0x00FD3A76, 0x0002EA52, 0x00FD1E5B, 0x0002F649, 0x00FD0240, 0x000706C0, 0x00F8F2C9, 0x00071ADB, 0x00F8EED2,
	0x00073EF6, 0x00F8CAFF, 0x000722ED, 0x00F8D6E4, 0x000776AC, 0x00F882A5, 0x00076AB7, 0x00F89EBE, 0x00074E9A, 0x00F8BA93, 0x00075281, 0x00F8A688,
	0x0007E618, 0x00F81211, 0x0007FA03, 0x00F80E0A, 0x0007DE2E, 0x00F82A27, 0x0007C235, 0x00F8363C, 0x00079674, 0x00F8627D, 0x00078A6F, 0x00F87E66,
	0x0007AE42, 0x00F85A4B, 0x0007B259, 0x00F84650, 0x0006C770, 0x00F93379, 0x0006DB6B, 0x00F92F62, 0x0006FF46, 0x00F90B4F, 0x0006E35D, 0x00F91754,
	0x0006B71C, 0x00F94315, 0x0006AB07, 0x00F95F0E, 0x00068F2A, 0x00F97B23, 0x00069331, 0x00F96738, 0x000627A8, 0x00F9D3A1, 0x00063BB3, 0x00F9CFBA,
	0x00061F9E, 0x00F9EB97, 0x00060385, 0x00F9F78C, 0x000657C4, 0x00F9A3CD, 0x00064BDF, 0x00F9BFD6, 0x00066FF2, 0x00F99BFB, 0x000673E9, 0x00F987E0,
	0x000485A0, 0x00FB71A9, 0x000499BB, 0x00FB6DB2, 0x0004BD96, 0x00FB499F, 0x0004A18D, 0x00FB5584, 0x0004F5CC, 0x00FB01C5, 0x0004E9D7, 0x00FB1DDE,
	0x0004CDFA, 0x00FB39F3, 0x0004D1E1, 0x00FB25E8, 0x00046578, 0x00FB9171, 0x00047963, 0x00FB8D6A, 0x00045D4E, 0x00FBA947, 0x00044155, 0x00FBB55C,
	0x00041514, 0x00FBE11D, 0x0004090F, 0x00FBFD06, 0x00042D22, 0x00FBD92B, 0x00043139, 0x00FBC530, 0x00054410, 0x00FAB019, 0x0005580B, 0x00FAAC02,
	0x00057C26, 0x00FA882F, 0x0005603D, 0x00FA9434, 0x0005347C, 0x00FAC075, 0x00052867, 0x00FADC6E, 0x00050C4A, 0x00FAF843, 0x00051051, 0x00FAE458,
	0x0005A4C8, 0x00FA50C1, 0x0005B8D3, 0x00FA4CDA, 0x00059CFE, 0x00FA68F7, 0x000580E5, 0x00FA74EC, 0x0005D4A4, 0x00FA20AD, 0x0005C8BF, 0x00FA3CB6,
	0x0005EC92, 0x00FA189B, 0x0005F089, 0x00FA0480,
}

// CheckCRC checks that the CRC of a message is valid and/or return the ICAO address / Interrogator Identifier of the
// message. As the message parity is a XOR of the CRC and the Address or Interrogator Identifier (except for DF17 and
// DF18), it is not possible to ensure that a message is correct without previously known valid Address or Interrogator
// Identifier. Only the messages DF18 and DF17 always give a valid Address.
//
// Params:
//   - message: The message to check
//   - data: The raw data of the message
//   - allowedAddresses: For the messages that have uncertainty when computing the Address, allows to reject the
//     messages having an unknown Address. Leave to nil to ignore.
//   - allowedInterrogatorIdentifiers: For the messages that have uncertainty when computing Interrogator Identifiers,
//     allows to reject the messages having an unknown Interrogator Identifier.
//
// Notes:
//   - the allowedAddresses is not used for messages DF17 and DF18 are always giving a valid address.
//
// Returns the ICAO Interrogator Identifiers for messages DF11 and Address for all others.
func CheckCRC(
	message messages.ModeSMessage,
	data []byte,
	allowedAddresses map[common.ICAOAddress]bool,
	allowedInterrogatorIdentifiers map[common.ICAOAddress]bool) (common.ICAOAddress, error) {

	switch message.GetDownLinkFormat() {
	case 11:
		return checkCRCDF11(data, allowedInterrogatorIdentifiers)
	case 17, 18:
		return checkCRCDF17And18(data)
	default:
		return checkCRCOther(data, allowedAddresses)
	}
}

func checkCRCDF11(
	data []byte,
	allowedInterrogatorIdentifiers map[common.ICAOAddress]bool) (common.ICAOAddress, error) {

	// For DF11, the ICAO code is not returned (as it is the base of the message). Instead, the interrogator id (II)
	// is used to XOR the parity. So, an interrogator can detect if a message is a reply to its interrogation.
	contentParity := computeParity(data[:4])
	messageParity := bitutils.Pack3Bytes(data[4], data[5], data[6])

	interrogatorIdentifier := common.ICAOAddress(contentParity ^ messageParity)

	// If the interrogator is not valid
	if len(allowedInterrogatorIdentifiers) > 0 {
		if _, ok := allowedInterrogatorIdentifiers[interrogatorIdentifier]; !ok {
			return 0, fmt.Errorf("the message parity resolves to an unknown Interrogator Identifier")
		}
	}

	return interrogatorIdentifier, nil
}

func checkCRCDF17And18(data []byte) (common.ICAOAddress, error) {

	// For DF17 and DF18 (extended squitter), the ICAO address is returned as the first 3 bytes of the payload.
	messageICAO := common.ICAOAddress(bitutils.Pack3Bytes(data[1], data[2], data[3]))

	// The parity is XORed against an Interrogator ID equals to 0
	contentParity := computeParity(data[:11])
	messageParity := bitutils.Pack3Bytes(data[11], data[12], data[13])

	if contentParity != messageParity {
		return 0, fmt.Errorf("the message does not have a valid CRC")
	}

	return messageICAO, nil
}

func checkCRCOther(
	data []byte,
	allowedAddresses map[common.ICAOAddress]bool) (common.ICAOAddress, error) {

	messageLength := len(data)

	// Compute parity on the whole message, except the 3 last bytes
	contentParity := computeParity(data[:messageLength-3])
	messageParity := bitutils.Pack3Bytes(data[messageLength-3], data[messageLength-2], data[messageLength-1])

	address := common.ICAOAddress(contentParity ^ messageParity)

	// If the address is not valid
	if len(allowedAddresses) > 0 {
		if _, ok := allowedAddresses[address]; !ok {
			return 0, fmt.Errorf("the message parity resolves to an unknown Address")
		}
	}

	return address, nil
}

// computeParity computes the parity of a slice of byte as 3-byte array. We used the implementation from:
// http://www.eurocontrol.int/eec/gallery/content/public/document/eec/report/1994/022_CRC_calculations_for_Mode_S.pdf
//
// params:
//   - data: The data for which to compute parity
//
// Returns the CRC (3 bytes)
func computeParity(data []byte) uint32 {

	crc := uint32(0x00000000) // Starting CRC value

	for _, byteVal := range data {
		tableIndex := byte((crc >> 16) ^ uint32(byteVal))
		crc = (crc << 8) ^ crcTable[tableIndex]
	}

	return crc & 0xFFFFFF // We are only interested in the least significant 24-bits
}
