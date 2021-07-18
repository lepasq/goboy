package gb

import (
	"io/ioutil"
	"log"
	"os"
)

type Cart struct { // [0x0 - 0x8000]
	Cart  []byte
	Title string
}

func Logo() []Byte { // [0x104 - 0x133]
	return []Byte{0xCE, 0xED, 0x66, 0x66, 0xCC, 0x0D, 0x00, 0x0B, 0x03, 0x73, 0x00, 0x83, 0x00, 0x0C, 0x00, 0x0D,
		0x00, 0x08, 0x11, 0x1F, 0x88, 0x89, 0x00, 0x0E, 0xDC, 0xCC, 0x6E, 0xE6, 0xDD, 0xDD, 0xD9, 0x99,
		0xBB, 0xBB, 0x67, 0x63, 0x6E, 0x0E, 0xEC, 0xCC, 0xDD, 0xDC, 0x99, 0x9F, 0xBB, 0xB9, 0x33, 0x3E,
	}
}

func (cart *Cart) SetTitle() { // [0x134 - 0x143]
	cart.Title = string(cart.Cart[0x134:0x143])
}

func ResetCart() {

}

func (cart *Cart) Read(address Word) byte {
	return cart.Cart[address]
}

func (cart *Cart) Write(address Word, value byte) {
	cart.Cart[address] = value
}

func (cart *Cart) LoadGame() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	cart.Cart = data
}

/*
* External Ram: can optionally come from cartridge as well
* MBC for games larger than 32KB


* Boot rom: animation, checksum
last instruction of bootrom launches game rom

* maybe split the [0x8000] array into multiple arrays (boot rom, game rom, ..)
*/
