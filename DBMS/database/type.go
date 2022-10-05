package database

type DBType interface {
	Value() interface{}
}

type TypeInteger struct {
	val int64
}

func (t TypeInteger) Value() interface{} {
	return t.val
}

type TypeReal struct {
	val float64
}

func (t TypeReal) Value() interface{} {
	return t.val
}

type TypeChar struct {
	val rune
}

func (t TypeChar) Value() interface{} {
	return t.val
}

type TypeString struct {
	val string
}

func (t TypeString) Value() interface{} {
	return t.val
}

type TypeHTML struct {
	val string
}

func (t TypeHTML) Value() interface{} {
	return t.val
}

type TypeStringRange struct {
	val []string
}

func (t TypeStringRange) Value() interface{} {
	return t.val
}
