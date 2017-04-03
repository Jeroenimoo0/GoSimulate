package simulate

type Resource int
const (
	None Resource = iota
	Wood
	Stone

	Axe
)

var Items = []Resource {
	Wood,
	Stone,

	Axe,
}

type enumResource struct {
	Resource
	name string
}

var EnumResources = []enumResource{
	{None, "NONE"},
	{Wood, "WOOD"},
	{Stone, "STONE"},

	{Axe, "AXE"},
}

func (i Resource) GetName() string {
	return EnumResources[i].name
}
