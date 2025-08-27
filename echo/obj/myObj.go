package obj

type MyObj struct {
	Name string
}

func (o *MyObj) Get() string {
	return o.Name
}

func (o *MyObj) Put(name string) {
	o.Name = name
}
