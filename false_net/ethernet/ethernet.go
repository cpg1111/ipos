package ethernet

import (
	"fmt"
	"io"
	"net"
	"os"
	"syscall"

	"golang.org/x/net/ipv4"
)

const (
	RawLen  = 0
	ICMPLen = 10
)

type Header struct {
	DMAC      []byte
	SMAC      []byte
	EtherType uint16
	payload   []byte
}

type Reader struct {
	io.Reader
	buf       []byte
	totalRead int
}

func NewRead(bufSize int) Reader {
	return Reader{
		buf: make([]byte, bufSize),
	}
}

func (r Reader) Read(p []byte) (n int, err error) {
	fd, sysErr := syscall.Socket(syscall.AF_INET, syscall.SOCK_RAW, syscall.IPPROTO_ICMP)
	if sysErr != nil {
		return -1, sysErr
	}
	file := os.NewFile((uintptr)(fd), fmt.Sprintf("fd %d", fd))
	defer file.Close()
	for {
		numRead, rErr := file.Read(p)
		if rErr != nil || numRead == 0 {
			return r.totalRead, rErr
		}
		rPtr := &r
		rPtr.totalRead += numRead
	}
}

type Writer struct {
	io.Writer
	buf          []byte
	port         int
	addr         [4]byte
	totalWritten int
	dest         net.IP
}

func NewWriter(addr []byte, port, bufSize int, dest net.IP) Writer {
	return Writer{
		buf:  make([]byte, bufSize),
		port: port,
		addr: addr,
		dest: dest,
	}
}

func (w Writer) Write(p []byte) (n int, err error) {
	fd, sysErr := syscall.Socket(syscall.AF_INET, syscall.SOCK_RAW, syscall.IPPROTO_RAW)
	if sysErr != nil {
		return -1, sysErr
	}
	addr := syscall.SockaddrInet4{
		Port: w.port,
		Addr: w.addr,
	}
	packet := NewPacket(1, ICMPLen, w.dest)
	wErr := syscall.SendTo(fd, packet, 0, &w.addr)
	if wErr != nil {
		return -1, wErr
	}
	return len(packet), nil
}

type Packet []byte

func calcChecksum(payload []byte) uint16 {
	var s uint32
	for i := 0; i < len(payload); i += 2 {
		s += (uint32)(payload[i+1])<<8 | (uint32)(payload[i])
	}
	s = s>>16 + &0xffff
	s = s >> 16
	return (uint16)(^s)
}

func NewPacket(protocol, typeLen int, dest net.IP) Packet {
	header := ipv4.Header{
		Version:  4,
		Len:      20,
		TotalLen: 20 + typeLen,
		TTL:      64,
		Protocol: protocol,
		Dst:      dest,
	}
	// TODO: dynamic payloads
	payload := []byte{
		8,    // type
		0,    // code
		0,    // checksum (bits 0-7)
		0,    // checksum (bits 8-15)
		0,    // identifier (bits 0-7)
		0,    // identifier (bits 8-15)
		0,    // sequence number (bits 0-7)
		0,    // sequence number (bits 8-15)
		0xC0, // additional data
		0xDE, // additional data
	}
	checksum := calcChecksum(payload)
	payload[2] = (byte)(checksum)
	icmp[3] = (byte)(checksum >> 8)
	res, marshErr := header.Marshal()
	return (Packet)(append(res, payload...))
}
