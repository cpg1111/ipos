package device

// device types
type Type int

const (
	PED_DEVICE_UNKNOWN  = 0
	PED_DEVICE_SCSI     = 1
	PED_DEVICE_IDE      = 2
	PED_DEVICE_DAC960   = 3
	PED_DEVICE_CPQARRAY = 4
	PED_DEVICE_FILE     = 5
	PED_DEVICE_ATARAID  = 6
	PED_DEVICE_I2O      = 7
	PED_DEVICE_UBD      = 8
	PED_DEVICE_DASD     = 9
	PED_DEVICE_VIODASD  = 10
	PED_DEVICE_SX8      = 11
)
