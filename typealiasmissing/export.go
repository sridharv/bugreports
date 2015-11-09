package typealiasmissing

type Alias string

type Op struct {
	Id int
}

func (o *Op) PrintAlias(alias Alias) {
	println(alias)
}
