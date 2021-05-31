package monkey

func addVal(addr []byte, byteCode []byte) []byte {
	byteCodeInt := int(byteCode[2])<<16 + int(byteCode[1])<<8 + int(byteCode[0])
	res := (int(addr[1])<<12+int(addr[0])<<4)*2 + byteCodeInt
	resByte := []byte{
		byte(res),
		byte(res >> 8),
		byte(res >> 16),
		byteCode[3],
	}
	return resByte
}

func jmpToFunctionValue(to uintptr) []byte {
	b := []byte{}
	for i := 0; i < 8; i++ {
		b = append(b, byte(to>>(i*8)))
	}

	byteCode := []byte{
		0x08, 0x00, 0x80, 0xD2, // mov x8, 0x0000
		0x08, 0x00, 0xA0, 0xF2, // movk x8, 0x0000, lsl #16
		0x08, 0x00, 0xC0, 0xF2, // movk x8, 0x0000, lsl #32
		0x08, 0x00, 0xE0, 0xF2, // movk x8, 0x0000, lsl #48

		0x08, 0x01, 0x40, 0xF9, // ldr x8, [x8]
		0x00, 0x01, 0x1F, 0xD6, // br x8
	}

	finalByteCode := []byte{}
	for i := 0; i < 4; i++ {
		finalByteCode = append(finalByteCode, addVal(b[(i*2):(i+1)*2], byteCode[(i*4):(i+1)*4])...)
	}

	finalByteCode = append(finalByteCode, byteCode[16:]...)

	return byteCode
}
