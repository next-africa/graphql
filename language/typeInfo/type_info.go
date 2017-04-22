package typeInfo

import (
	"github.com/next-africa/graphql-go/language/ast"
)

// TypeInfoI defines the interface for TypeInfo Implementation
type TypeInfoI interface {
	Enter(node ast.Node)
	Leave(node ast.Node)
}
