package main

func main() {
	x, err := Foo()
	if err != nil {
		//handle it
		return
	}
	//x is free to use, no errors at all
	_ = x

	y, err := Bar()
	if err != nil {
		//handle it
		return
	}
	//y is free to use, no errors at all
	_ = y
}

func Foo() (int, error) {
	return 666, nil
}
func Bar() (int, error) {
	return 666, nil
}
func Baz() (int, error) {
	return 666, nil
}
