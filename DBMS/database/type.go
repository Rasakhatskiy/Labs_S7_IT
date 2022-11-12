package database

import "reflect"

type DBType interface {
	Value() interface{}
	TypeName() string
}

type TypeInteger struct {
	Val int64
}

func (t TypeInteger) Value() interface{} {
	return t.Val
}

func (t TypeInteger) TypeName() string {
	return "Integer"
}

type TypeReal struct {
	Val float64
}

func (t TypeReal) Value() interface{} {
	return t.Val
}

func (t TypeReal) TypeName() string {
	return "Real"
}

type TypeChar struct {
	Val rune
}

func (t TypeChar) Value() interface{} {
	return t.Val
}

func (t TypeChar) TypeName() string {
	return "Char"
}

type TypeString struct {
	Val string
}

func (t TypeString) Value() interface{} {
	return t.Val
}

func (t TypeString) TypeName() string {
	return "String"
}

type TypeHTML struct {
	Val string
}

func (t TypeHTML) Value() interface{} {
	return t.Val
}

func (t TypeHTML) TypeName() string {
	return "HTML Document"
}

func (t TypeHTML) Validate(data string) error {
	//todo: implement
	// check and assign

	//panic("implement me")

	t.Val = data

	return nil
}

type TypeStringRange struct {
	Val []string
}

func (t TypeStringRange) Value() interface{} {
	return t.Val
}

func (t TypeStringRange) TypeName() string {
	return "String Interval"
}

func (t TypeStringRange) Validate(data string) error {
	//todo: implement
	// check and assign
	panic("implement me")
}

var (
	TypeIntegerTS     = reflect.TypeOf(TypeInteger{}).String()
	TypeRealTS        = reflect.TypeOf(TypeReal{}).String()
	TypeCharTS        = reflect.TypeOf(TypeChar{}).String()
	TypeStringTS      = reflect.TypeOf(TypeString{}).String()
	TypeHTMLTS        = reflect.TypeOf(TypeHTML{}).String()
	TypeStringRangeTS = reflect.TypeOf(TypeStringRange{}).String()
)
