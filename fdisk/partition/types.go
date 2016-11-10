package partition

// partition types

// Type is an alias to byte for partition types
type Type byte

const (
	// PartitionNormal is a standard Linux/Unix partition
	PartitionNormal = 0x00
	// PartitionLogical is a lvm partition
	PartitionLogical = 0x01
	// PartitionExtended is an extended partition
	PartitionExtended = 0x02
	// PartitionFreespace is a free (empty) space partition
	PartitionFreespace = 0x04
	// PartitionMetadata is for a metadata partition
	PartitionMetadata = 0x08
	// PartitionProtected is a protected partition
	PartitionProtected = 0x10
)

// partition flags

// Flag is an alias to an int for partition flags
type Flag int

const (
	// PartitionBoot equals 1, it is for a boot partition
	PartitionBoot = 1
	// PartitionRoot equals 2, it is for a root partition
	PartitionRoot = 2
	// PartitionSwap equals 3, it is for a Swap partition
	PartitionSwap = 3
	// PartitionHidden equals 4, it is for a hidden partition
	PartitionHidden = 4
	// PartitionRaid equals 5, it is for a RAID partition
	PartitionRaid = 5
	// PartitionLVM equals 6, it is for a LVM partition
	PartitionLVM = 6
	// PartitionLBA equals 7, it is for a LBA partition
	PartitionLBA = 7
	// PartitionHPService equals 8, it is for a HPService partition
	PartitionHPService = 8
	// PartitionPalo equals 9, it is for a Palo partition
	PartitionPalo = 9
	// PartitionPrep equals 10, it is for a prep partition
	PartitionPrep = 10
	// PartitionMSFTReserved equals 11, it is for reserving a partition for MSFT
	PartitionMSFTReserved = 11
)

// partition type feature

// TypeFeature is an alias to an int for feature types
type TypeFeature int

const (
	// DiskTypeExtended equals 1, it is for an extended disk type
	DiskTypeExtended = 1
	// DiskTypePartitionName equals 2, it is for the name of the partition type
	DiskTypePartitionName = 2
)

// System file types

// SysType is a struct for the type of FS on the system
type SysType struct {
	BlockSizes []int
	Name       string
	Next       *SysType
	Ops        *FSOps
}
