package arp

type Header struct {
	HWType  uint16
	ProType uint16
	HWSize  byte
	ProSize byte
	OPCode  uint16
	Data    []byte
}

type IPV4 struct {
	SMAC []byte
	SIP  uint32
	DMAC []byte
	DIP  uint32
}
