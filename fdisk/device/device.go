package device

import (
	"github.com/cpg1111/ipos/pedCHS"
)

type Device struct {
	ArchSpecific   *interface{}
	BIOSGeom       pedCHS.Geometry
	BootDirty      int
	DID            int
	Dirty          int
	ExternalMode   int
	Host           int
	HWGeom         pedCHS.Geometry
	Length         int64
	Model          string
	Next           *Device
	OpenCount      int
	Path           string
	PhysSectorSize int64
	ReadOnly       bool
	SectorSize     int64
	Type           Type
}

func New(devPath string) *Device {

}
