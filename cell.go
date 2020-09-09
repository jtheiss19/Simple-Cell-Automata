package automata

//Cell remain mostly unchanged, they carry the Value Type which can carry type such as "air" or "wood" for tree simulation
type Cell struct {
	Type  string
	XPos  int
	YPos  int
	Value float64
}
