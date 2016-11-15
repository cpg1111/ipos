package disk

import (
	"errors"

	"github.com/cpg1111/ipos/fdisk/device"
	"github.com/cpg1111/ipos/fdisk/partition"
)

// Disk is a struct representing a disk
type Disk struct {
	Type         partition.Type
	Size         uint
	PartList     *partition.Partition
	Dev          *device.Device
	NeedsClobber bool
	UpdateMode   bool
	Specific     *interface{}
}

// New takes a *device.Device as an argument and returns a new *Disk
func New(dev *device.Device) *Disk {
	return nil
}

// Existing finds a disk based on path
func Existing(path string) *Disk {
	return nil
}

func (d *Disk) rawInsertBefore(old, newP *partition.Partition) error {
	if old == nil {
		return errors.New("no starting point given")
	}
	if newP == nil {
		return errors.New("no new partition given")
	}
	newP.Prev = old.Prev
	newP.New = old
	if newP.Prev != nil {
		newP.Prev.Next = newP
	} else {
		if old.Type == partition.PED_PARTITION_LOGICAL {
			d.extendedPartition().PartList = newP
		} else {
			d.PartList = newP
		}
	}
	old.Prev = newP
	return nil
}

func (d *Disk) rawInsertAfter(old, newP *partition.Partition) error {
	if old == nil {
		return errors.New("no starting point given")
	}
	if newP == nil {
		return errors.New("no new partition given")
	}
	newP.Prev = old
	newP.Next = old.Next
	if old.Next != nil {
		old.Next.Prev = newP
	}
	old.Next = newP
	return nil
}

func (d *Disk) rawRemove(part *partition.Partition) error {
	if part == nil {
		return errors.New("no partition given")
	}
	if part.Prev != nil {
		part.Prev.Next = part.Next
		if part.Next != nil {
			part.Next.Prev = part.prev
		}
		return nil
	}
	if part.Type == partition.PartitionLogical {
		d.ExtendedPartition().PartList = part.Next
	} else {
		d.PartList = part.Next
	}
	if part.Next != nil {
		part.Next.Prev = nil
	}
	return nil
}

func (d *Disk) rawAdd(part *partition.Partition) error {
	var (
		walk    *partition.Partition
		last    *partition.Partition
		extPart *partition.Partition
	)
	if !d.UpdateMode {
		return errors.New("the disk is not in update mode")
	}
	extPart = d.ExtendPartition()
	if part.Type == partition.PED_PARTITION_LOGICAL {
		walk = extPart.PartList
	}
	walk = d.PartList
	for walk.Next != nil {
		if walk.Geom.Start > part.Geom.End {
			break
		}
		last = walk
		walk = walk.Next
	}
	if walk != nil {
		return d.rawInsertBefore(walk, part)
	}
	if last != nil {
		return d.rawInsertAfter(last, part)
	}
	if part.Type == partition.PED_PARTITION_LOGICAL {
		extPart.PartList = part
	} else {
		d.PartList = part
	}
	return nil
}

// CheckPartition checks a partition on the disk
func (d *Disk) CheckPartition(part *partition.Partition) error {
	return nil
}

// AddPartition adds a partition to the disk
func (d *Disk) AddPartition(part *partition.Partition, constr *partition.Constraint) error {
	var (
		overlapConstraint *partition.Constraint
		constraints       *partition.Constraint
	)
	if part == nil {
		return errors.New("no partition provided for adding to disk")
	}
	if part.IsActive() && constr != nil {
		overlapConstraint = part.GetOverlapConstraint()
		constraints = constr.GetIntersect(overlapConstraint)
		if constraints == nil {
			return errors.New("can not have overlapping partitions")
		}
		constr = constraints
	}
	alignErr := part.Align(constr)
	if alignErr != nil {
		return alignErr
	}
	checkErr := d.CheckPartition(part)
	if checkErr != nil {
		return checkErr
	}
	return d.rawAdd(part)
}
