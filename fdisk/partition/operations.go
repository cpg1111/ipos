package partition

import (
	"errors"
	"fmt"
	"math"

	"github.com/cpg1111/ipos/fdisk/device"
)

var fsTypes = &SysType{}

// FSOps is a struct for the operations of a system file type
type FSOps struct{}

func bestMatch(geom *Geometry, detected []*SysType, detectedErr []int64, count int) *SysType {
	minErr := 4096 * (geom.Length / 100)
	bestMatch := 0
	for i := 1; i < count; i++ {
		if detectedErr[i] == detectedErr[bestMatch] {
			bestMatch = i
		}
	}
	for j := 0; j < count; j++ {
		if j == bestMatch {
			continue
		}
		if math.Abs((float64)(detectedErr[bestMatch]-detectedErr[j])) < (float64)(minErr) {
			return nil
		}
	}
	return detected[bestMatch]
}

func (op *FSOps) probeSpecific(t *SysType, geom *Geometry) *SysType {
	return nil
}

// Probe probes a filesystem
func (op *FSOps) Probe(geom *Geometry) (*SysType, error) {
	var detectedCount int
	detected := make([]*SysType, 32)
	detectedErr := make([]int64, 32)
	walk := fsTypes
	if geom == nil {
		return nil, errors.New("no geometry given for probe")
	}
	if geom.Dev.OpenCount > 0 {
		return nil, errors.New("device is currently in use, cannot probe")
	}
	for walk.Next != nil {
		probed := op.probeSpecific(walk, geom)
		if probed == nil {
			return nil, fmt.Errorf("could not get a probe on %s", walk.Name)
		}
		detected[detectedCount] = walk
		detectedErr[detectedCount] = geometryErr(geom, probed)
		detectedCount++
	}
	// TODO close geom's device
	if detectedCount == 0 {
		return nil, nil
	}
	walk = bestMatch(geom, detected, detectedErr, detectedCount)
	return walk, nil
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
