package fish

func Kind(fishName string) FishKind {
	desc, ok := fishDesc[fishName]
	if ok {
		return desc.kind
	}
	return Unknown
}
