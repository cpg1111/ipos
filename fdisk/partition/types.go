package disk

// partition types

type Type byte

const (
	PED_PARTITION_NORMAL    = 0x00
	PED_PARTITION_LOGICAL   = 0x01
	PED_PARTITION_EXTENDED  = 0x02
	PED_PARTITION_FREESPACE = 0x04
	PED_PARTITION_METADATA  = 0x08
	PED_PARTITION_PROTECTED = 0x10
)

// partition flags

type Flag int

const (
	PED_PARTITION_BOOT          = 1
	PED_PARTITION_ROOT          = 2
	PED_PARTITION_SWAP          = 3
	PED_PARTITION_HIDDEN        = 4
	PED_PARTITION_RAID          = 5
	PED_PARTITION_LVM           = 6
	PED_PARTITION_LBA           = 7
	PED_PARTITION_HPSERVICE     = 8
	PED_PARTITION_PALO          = 9
	PED_PARTITION_PREP          = 10
	PED_PARTITION_MSFT_RESERVED = 11
)

// partition type feature

type TypeFeature int

const (
	PED_DISK_TYPE_EXTENDED       = 1
	PED_DISK_TYPE_PARTITION_NAME = 2
)
