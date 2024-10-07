package parser

import (
	"strings"

	"github.com/Hilson-Alex/Butterfly/src/compiler/checker"
)

func saveModule(code string, result *string) {
	*result = `
package generated_code

import (
	"butterfly_embed/runtime"
)

func init() {
` +
		code +
		"\n}"
}

func eventListen(event, param, commandBlock string, scope *checker.BFScope) string {
	var callback = concat("func(", param, " runtime.BF__MessageContent)", commandBlock)
	return "runtime.BF__EventSubscribe(" + join(",", event, quote(scope.Module().Name()), callback) + ")"
}

func eventShare(event, messageContent string, scope *checker.BFScope) string {
	var message = concat(
		"runtime.BF__MessageCreate(",
		join(",", quote(scope.Module().Name()), messageContent),
		")",
	)
	return concat("runtime.BF__Dispatch(", join(",", event, message), ")")
}

func computeSize(delimiter string, args []string) int {
	n := len(delimiter) * (len(args) - 1)
	for i := 0; i < len(args); i++ {
		n += len(args[i])
	}
	return n
}

func join(delimiter string, args ...string) string {
	if len(args) == 0 {
		return ""
	}
	var builder strings.Builder
	builder.Grow(computeSize(delimiter, args))
	builder.WriteString(args[0])
	for _, arg := range args[1:] {
		if arg != "" {
			builder.WriteString(delimiter)
			builder.WriteString(arg)
		}
	}
	return builder.String()
}

func concat(args ...string) string {
	return join("", args...)
}

func wsJoin(args ...string) string {
	return join(" ", args...)
}

func blJoin(args ...string) string {
	return join("\n", args...)
}

func getType(baseType string) string {
	if baseType == "float" {
		return "float64"
	}
	return baseType
}

func createModule(scope *checker.BFScope, moduleName string) {
	updateScope(scope, checker.RootScope(moduleName))
}

func quote(s string) string {
	return "\"" + s + "\""
}

func updateScope(old *checker.BFScope, new *checker.BFScope) {
	*old = *new
}

func addVar(scope *checker.BFScope, name, varType, value string) {
	var bfType checker.BFType = checker.BFPrimitive(varType)
	if bfType == checker.UNKNOWN {
		bfType = checker.ANY
		// bfType, _ = checker.DefaultType(value)
	} else if !bfType.CanAssign(value, scope) {
		return
	}
	var _ = scope.AddSymbol(checker.CreateVar(name, bfType))
}

func addConst(scope *checker.BFScope, name, value, varType string) {
	var bfType checker.BFType = checker.BFPrimitive(varType)
	if bfType == checker.UNKNOWN {
		bfType = checker.ANY
		// bfType, _ = checker.DefaultType(value)
	} else if !bfType.CanAssign(value, scope) {
		return
	}
	_ = scope.AddSymbol(checker.CreateConst(name, bfType))
}
