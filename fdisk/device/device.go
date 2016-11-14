package device

import (
	"github.com/cpg1111/ipos/fdisk/pedCHS"
)

// Device is a struct representing a device on the machine
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

// New takes a path to a device and returns a pointer to the Device struct
func New(devPath string) *Device {
	return nil
}
