package checker

import (
	"errors"
)

type BFScope struct {
	module      *Symbol
	symbolTable map[string]*Symbol
	parentScope *BFScope
}

func RootScope(moduleName string) *BFScope {
	return &BFScope{
		module:      CreateConst(moduleName, MODULE),
		symbolTable: make(map[string]*Symbol),
	}
}

func (scope *BFScope) AddSymbol(symbol *Symbol) error {
	if _, exist := scope.symbolTable[symbol.Name()]; exist {
		return errors.New("already declared in this scope")
	}
	scope.symbolTable[symbol.Name()] = symbol
	return nil
}

func (scope *BFScope) SubScope() *BFScope {
	return &BFScope{
		module:      scope.module,
		symbolTable: make(map[string]*Symbol),
		parentScope: scope,
	}
}

func (scope *BFScope) Close() *BFScope {
	return scope.parentScope
}

func (scope *BFScope) GetSymbol(symbolName string) *Symbol {
	if item, ok := scope.symbolTable[symbolName]; ok {
		return item
	}
	if scope.parentScope == nil {
		return nil
	}
	return scope.parentScope.GetSymbol(symbolName)
}

func (scope *BFScope) Module() *Symbol {
	return scope.module
}
