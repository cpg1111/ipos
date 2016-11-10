package partition

import (
	"github.com/cpg1111/ipos/fdisk/device"
)

type Geometry struct {
	Dev    *device.Device
	Start  int64
	Length int64
	End    int64
}
