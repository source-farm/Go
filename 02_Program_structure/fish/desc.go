package fish

type description struct {
	kind       FishKind
	mainColor  string
	endangered bool
}

var fishDesc map[string]description = map[string]description{
	"Salmon": {Sea, "grey, white", false},
	"Trout":  {FreshWater, "green, white, red", false},
}
