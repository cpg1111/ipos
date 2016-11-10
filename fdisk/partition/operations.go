package partition

import (
	"errors"

	"github.com/cpg1111/ipos/fdisk/device"
)

// FSOps is a struct for the operations of a system file type
type FSOps struct{}

// Probe probes a filesystem
func (op *FSOps) Probe(geom *Geometry) (*SysType, error) {
	var walk *SysType
	var detectedCount int
	detected := make([]*SysType, 32)
	detectedErr := make([]error, 32)
	if geom == nil {
		return nil, errors.New("no geometry given for probe")
	}
	return nil
}

// Clobber clobbers a file system
func (op *FSOps) Clobber(geom *Geometry) {}

// Open opens a block in the file system
func (op *FSOps) Open(geom *Geometry) {}

// Create creates on the file system
func (op *FSOps) Create(geom *Geometry) {}

// Close closes on the file system
func (op *FSOps) Close(fs *FS) {}

// Check checks the file system
func (op *FSOps) Check(fs *FS, timer *Timer) {}

// Copy copies on the file system
func (op *FSOps) Copy(fs *FS, geom *Geometry, timer *Timer) {}

// Resize resizes the file system
func (op *FSOps) Resize(fs *FS, geom *Geometry, timer *Timer) {}

// GetCreateConstraint gets the constraints for creating
func (op *FSOps) GetCreateConstraint(dev *device.Device) {}

// GetResizeConstraint gets the constraints for resizing
func (op *FSOps) GetResizeConstraint(fs *FS) {}

// GetCopyConstraint gets the constraints for copying
func (op *FSOps) GetCopyConstraint(fs *FS, dev *device.Device) {}
