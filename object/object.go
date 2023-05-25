package object

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"

	"github.com/jcbbb/monkey/ast"
)

type ObjectType int

const (
	IntegerObj ObjectType = iota
	BooleanObj
	NullObj
	ReturnObj
	ErrorObj
	FunctionObj
	StringObj
	BuiltinObj
	ArrayObj
)

var objectTypes = [...]string{
	IntegerObj: "INTEGER",
	BooleanObj: "BOOLEAN",
	NullObj:    "NULL",
	ReturnObj:  "RETURN",
	ErrorObj:   "ERROR",
	StringObj:  "STRING",
	BuiltinObj: "BUILTIN",
	ArrayObj:   "ARRAY",
}

func (ot ObjectType) String() string {
	s := ""
	if 0 <= ot && ot < ObjectType(len(objectTypes)) {
		s = objectTypes[ot]
	}

	if s == "" {
		s = "type(" + strconv.Itoa(int(ot)) + ")"
	}

	return s

}

type Object interface {
	Type() ObjectType
	Inspect() string
}

type Environment struct {
	store map[string]Object
	outer *Environment
}

type Integer struct {
	Value int64
}

type Boolean struct {
	Value bool
}

type ReturnValue struct {
	Value Object
}

type Null struct{}

type Error struct {
	Message string
}

type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

type String struct {
	Value string
}

type Array struct {
	Elements []Object
}

type BuiltinFunction func(args ...Object) Object
type Builtin struct {
	Fn BuiltinFunction
}

func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s}
}

func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}

func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}

func (i *Integer) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}
func (i *Integer) Type() ObjectType { return IntegerObj }

func (b *Boolean) Inspect() string {
	return fmt.Sprintf("%t", b.Value)
}
func (b *Boolean) Type() ObjectType { return BooleanObj }

func (n *Null) Inspect() string {
	return "null"
}
func (n *Null) Type() ObjectType { return NullObj }

func (rv *ReturnValue) Inspect() string {
	return rv.Value.Inspect()
}
func (rv *ReturnValue) Type() ObjectType { return ReturnObj }

func (e *Error) Inspect() string {
	return "ERROR: " + e.Message
}
func (e *Error) Type() ObjectType { return ErrorObj }

func (f *Function) Type() ObjectType { return FunctionObj }
func (f *Function) Inspect() string {
	var out bytes.Buffer

	params := []string{}

	for _, p := range f.Parameters {
		params = append(params, p.String())
	}

	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")

	return out.String()
}

func (s *String) Inspect() string {
	return s.Value
}
func (s *String) Type() ObjectType { return StringObj }

func (bf *Builtin) Inspect() string {
	return "builtin function"
}
func (bf *Builtin) Type() ObjectType {
	return BuiltinObj
}

func (a *Array) Inspect() string {
	var out bytes.Buffer

	elements := []string{}

	for _, e := range a.Elements {
		elements = append(elements, e.Inspect())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	return out.String()
}
func (a *Array) Type() ObjectType {
	return ArrayObj
}
