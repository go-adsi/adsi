package adsi

func reverseUint16(v uint16) uint16 {
	return v>>8&0x00ff | v<<8&0xff00
}

func reverseUint32(v uint32) uint32 {
	return v>>24&0x000000ff | v>>8&0x0000ff00 | v<<8&0x00ff0000 | v<<24&0xff000000
}
