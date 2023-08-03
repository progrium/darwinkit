package data

type KindType struct {
	Kind TypeKind `json:"kind,omitempty"`
}

type TypeInfo interface {
	GetName() string
	GetModule() string
	FullName() string
}

type TypeKind string

const (
	KindClass    TypeKind = "class"
	KindProtocol TypeKind = "protocol"
	KindAlias    TypeKind = "alias"
	KindStruct   TypeKind = "struct"
)

type Class struct {
	Kind       TypeKind    `json:"kind,omitempty"`
	Name       string      `json:"name,omitempty"`       // class type name
	Module     string      `json:"module,omitempty"`     // module name
	Parent     string      `json:"parent,omitempty"`     // parent class type name
	Properties []*Property `json:"properties,omitempty"` // properties
	Methods    []*Method   `json:"methods,omitempty"`    // methods
}

func (c *Class) GetName() string {
	return c.Name
}

func (c *Class) GetModule() string {
	return c.Module
}

func (c *Class) FullName() string {
	return c.Module + "." + c.Name
}

type Protocol struct {
	Kind       TypeKind    `json:"kind,omitempty"`
	Name       string      `json:"name,omitempty"`       // the type name
	Module     string      `json:"module,omitempty"`     // module name
	Parents    []string    `json:"parents,omitempty"`    // the parent protocols names
	Methods    []*Method   `json:"methods,omitempty"`    // methods
	Properties []*Property `json:"properties,omitempty"` // properties
}

func (p *Protocol) GetName() string {
	return p.Name
}

func (p *Protocol) GetModule() string {
	return p.Module
}

func (p *Protocol) FullName() string {
	return p.Module + "." + p.Name
}

var _ TypeInfo = (*Struct)(nil)

// Struct c struct type
type Struct struct {
	Kind   TypeKind `json:"kind,omitempty"`
	Name   string   `json:"name,omitempty"`   // class type name
	Module string   `json:"module,omitempty"` // module name
}

func (s *Struct) GetName() string {
	return s.Name
}

func (s *Struct) GetModule() string {
	return s.Module
}

func (s *Struct) FullName() string {
	return s.Module + "." + s.Name
}

// Alias for typedef or enums
type Alias struct {
	Kind       TypeKind     `json:"kind,omitempty"`
	Name       string       `json:"name,omitempty"`        // class type name
	Module     string       `json:"module,omitempty"`      // module name
	TargetType string       `json:"target_type,omitempty"` // the target type name
	Values     []*EnumValue `json:"values,omitempty"`
}

func (a *Alias) GetName() string {
	return a.Name
}

func (a *Alias) GetModule() string {
	return a.Module
}

func (a *Alias) FullName() string {
	return a.Module + "." + a.Name
}

// EnumValue the enum name and value
type EnumValue struct {
	Name       string `json:"name,omitempty"`        // the objc enum name
	Value      string `json:"value,omitempty"`       // the enum value. if has different values in different arch, this value is for amd64.
	Arm64Value string `json:"arm64_value,omitempty"` // some may have different enum value in arm64
	Module     string `json:"module,omitempty"`      // the module enum value defined in
}

type Property struct {
	Name          string `json:"name,omitempty"`           // the property name
	GoName        string `json:"go_name,omitempty"`        // optional go property name
	ReadOnly      bool   `json:"read_only,omitempty"`      // if is readonly property
	GetterName    string `json:"getter_name,omitempty"`    // the getter name for property
	Type          string `json:"type,omitempty"`           // the type name
	ClassProperty bool   `json:"class_property,omitempty"` // if is class property
	Weak          bool   `json:"weak,omitempty"`           // weak property
	Deprecated    bool   `json:"deprecated,omitempty"`     // if is deprecated
	Required      bool   `json:"required,omitempty"`       // for protocol, if require implement
}

// Function is a objective-c Function
type Function struct {
	Name       string   `json:"name,omitempty"`        // the first part of objc function name
	Params     []*Param `json:"params,omitempty"`      // the params
	ReturnType string   `json:"return_type,omitempty"` // the return type name
	Deprecated bool     `json:"deprecated,omitempty"`  // if has been deprecated
}

// Method is a protocol/class method
type Method struct {
	Name        string   `json:"name,omitempty"`         // the first part of objc method name
	GoName      string   `json:"go_name,omitempty"`      // may be empty.if empty, use Name
	Params      []*Param `json:"params,omitempty"`       // the params
	ReturnType  string   `json:"return_type,omitempty"`  // the return type name
	ClassMethod bool     `json:"class_method,omitempty"` // true if is class method
	Deprecated  bool     `json:"deprecated,omitempty"`   // if has been deprecated
	Required    bool     `json:"required,omitempty"`     // If this method is required. only for protocol method.
}

// Param is method param
type Param struct {
	Name         string `json:"name,omitempty"`          // the param name
	VariableName string `json:"variable_name,omitempty"` // the variable name
	Type         string `json:"type,omitempty"`          // type
}
