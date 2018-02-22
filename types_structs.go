package main

// TypeField ...
type TypeField struct {
	Name string
	Type string
	Tags string
}

// TypeDefinition ...
type TypeDefinition struct {
	TypePackage            string
	TypeName               string
	TypeFields             []*TypeField
	InternalTypeName       string
	PluralTypeName         string
	PluralInternalTypeName string
}
