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
	if _, ok := scope.symbolTable[symbol.Name()]; ok {
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

func (scope *BFScope) HasSymbol(symbolName string) bool {
	if _, ok := scope.symbolTable[symbolName]; ok {
		return true
	}
	if scope.parentScope == nil {
		return false
	}
	return scope.parentScope.HasSymbol(symbolName)
}

func (scope *BFScope) Module() *Symbol {
	return scope.module
}
