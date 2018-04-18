package test

type Test struct {
}

func (*Test) Version() string {
	return "1.0.0"
}
