package partition

import (
	"github.com/cpg1111/ipos/fdisk/device"
)

// Geometry is a struct for the geometry of the partition
type Geometry struct {
	Dev    *device.Device
	Start  int64
	Length int64
	End    int64
}

func geometryErr(geom *Geometry, t *SysType) int64 {
	return 0
}
