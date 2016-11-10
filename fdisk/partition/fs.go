package partition

// FS is a struct for the file system
type FS struct {
	Type         *SysType
	Geom         *Geometry
	Checked      bool
	TypeSpecific *interface{}
}
