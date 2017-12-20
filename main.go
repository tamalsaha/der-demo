package main

import (
	lib  "github.com/go-asn1-ber/asn1-ber"
	"log"
	"fmt"
)

func main() {
	p, err := lib.DecodePacketErr([]byte{0x88, 0x05, 0x2A, 0x03, 0x04, 0x05, 0x05})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(p.ClassType)
	fmt.Println(p.Tag)
	fmt.Println(p.TagType)

	lib.PrintPacket(p)
}
