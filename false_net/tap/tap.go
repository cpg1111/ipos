package tap

import (
	"os"

	"github.com/cpg1111/ipos/false_net/netinterface"
)

const (
	IFF_TAP   = 0x0002
	IFF_NO_PI = 0x1000
)

func New(name string) (intfce *netinterface.Interface, err error) {
	file, err := os.OpenFile("/dev/net/tun", os.ORDWR, 0)
	if err != nil {
		return nil, err
	}
	iName, err := netinterface.NewReq(file.Fd(), name, IFF_TAP|IFF_NO_PI)
	if err != nil {
		return nil, err
	}
	intface = netinterface.New(file, true, iName)
	return
}
