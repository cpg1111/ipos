package disk

import (
	"github.com/cpg1111/ipos/fdisk/partition"
)

type Disk struct {
	Type         partition.Type
	Size         uint
	Part         *partition.Partition
	Dev          *device.Device
	NeedsClobber bool
	UpdateMode   int
	Specific     *interface{}
}

func New(devPath string, dType byte, begin, end uint, mode int) *Disk {

}

func NewExistingDisk(path string) *Disk {

}
