package partition

// Constraint is a struct for partition Constraints
type Constraint struct {
	Start      *Alignment
	End        *Alignment
	StartRange *Geometry
	EndRange   *Geometry
	MinSize    int64
	MaxSize    int64
}
