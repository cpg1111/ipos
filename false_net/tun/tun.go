package tun

import (
	"os"

	"github.com/cpg1111/ipos/false_net/netinterface"
)

const (
	IFF_TUN   = 0x0001
	IFF_NO_PI = 0x1000
)

func New(name string) (intfce *netinterface.Interface, err error) {
	file, err := os.OpenFile("/dev/net/tun", os.O_RDWR, 0)
	if err != nil {
		return nil, err
	}
	iName, err := netinterface.NewReq(file.Fd(), name, IFF_TUN|IFF_NO_PI)
	if err != nil {
		return nil, err
	}
	intfce = netinterface.New(file, false, iName)
	return
}
