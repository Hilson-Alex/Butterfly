package checker

type Symbol struct {
	name string
	// scope      BFScope
	// keyType     BFType
	valueType   BFType
	canReassign bool
}

func CreateVar(name string, bfType BFType) *Symbol {
	return &Symbol{
		name:        name,
		valueType:   bfType,
		canReassign: true,
	}
}

func CreateConst(name string, bfType BFType) *Symbol {
	return &Symbol{
		name:        name,
		valueType:   bfType,
		canReassign: false,
	}
}

func (symbol *Symbol) Name() string {
	return symbol.name
}

func (symbol *Symbol) Type() BFType {
	return symbol.valueType
}

func (symbol *Symbol) CanReassign() bool {
	return symbol.canReassign
}
