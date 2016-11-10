package partition

type Constraint struct {
	Start      *Alignment
	End        *Alignment
	StartRange *Geometry
	EndRange   *Geometry
	MinSize    int64
	MaxSize    int64
}
