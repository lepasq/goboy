package gb

var Data [0x10000]int32

func ResetMemory() {
	for i := range Data {
		Data[i] = 0
	}
}

func Operator(address int32) int32 {
	return Data[address]
}
