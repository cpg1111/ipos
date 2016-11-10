package disk

import (
	"errors"

	"github.com/cpg1111/ipos/shared"
	
	"github.com/cpg1111/ipos/fdisk/device"
	"github.com/cpg1111/ipos/fdisk/partition"
)

type Disk struct {
	Type         partition.Type
	Size         uint
	PartList         *partition.Partition
	Dev          *device.Device
	NeedsClobber bool
	UpdateMode   bool
	Specific     *interface{}
}

func New(dev *device.Device) *Disk {
	return nil
}

func NewExistingDisk(path string) *Disk {
	return nil
}

func (d *Disk) rawInsertBefore(old, new *partition.Partition) error {
	if old == nil {
		return errors.New("no starting point given")
	}
	if new == nil {
		return errors.New("no new partition given")
	}
	new.Prev = old.Prev
	new.New = old
	if new.Prev != nil {
		new.Prev.Next = new
	} else {
		if old.Type == partition.PED_PARTITION_LOGICAL {
			d.extendedPartition().PartList = new
		} else {
			d.PartList = new
		}
	}
	old.Prev = new
	return nil
}

func (d *Disk) rawInsertAfter(old, new *partition.Partition) error {
	if old == nil {
		return errors.New("no starting point given")
	}
	if new == nil {
		return errors.New("no new partition given")
	}
	new.Prev = old
	new.Next = old.Next
	if old.Next != nil {
		old.Next.Prev = new
	}
	old.Next = new
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
	if part.Type == partition.PED_PARTITION_LOGICAL {
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
		walk *partition.Partition
		last *partition.Partition
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
	for walk; last = walk, walk = walk.Next {
		if walk.Geom.Start > part.Geom.End {
			break
		}
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

func (d *Disk) CheckPartition(part *partition.Partition) error {
	return nil
}

func (d *Disk) AddPartition(part *partition.Partition, constr *partition.Constraint) error {
	var (
		overlapConstraint *partition.Constraint
		constraints *partition.Constraint
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
