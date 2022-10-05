package database

type DBType interface {
	Value() interface{}
}

type TypeInteger struct {
	Val int64
}

func (t TypeInteger) Value() interface{} {
	return t.Val
}

type TypeReal struct {
	Val float64
}

func (t TypeReal) Value() interface{} {
	return t.Val
}

type TypeChar struct {
	Val rune
}

func (t TypeChar) Value() interface{} {
	return t.Val
}

type TypeString struct {
	Val string
}

func (t TypeString) Value() interface{} {
	return t.Val
}

type TypeHTML struct {
	Val string
}

func (t TypeHTML) Value() interface{} {
	return t.Val
}

type TypeStringRange struct {
	Val []string
}

func (t TypeStringRange) Value() interface{} {
	return t.Val
}
