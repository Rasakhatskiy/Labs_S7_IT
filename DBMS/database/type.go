package database

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

func (t TypeHTML) Validate() bool {
	//todo: implement
	panic("implement me")
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

func (t TypeStringRange) Validate() bool {
	//todo: implement
	panic("implement me")
}
