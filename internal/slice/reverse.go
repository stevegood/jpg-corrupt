package slice

func ReverseBytes(bytes []byte) []byte {
	if len(bytes) == 0 {
		return bytes
	}
	return append(ReverseBytes(bytes[1:]), bytes[0])
}
