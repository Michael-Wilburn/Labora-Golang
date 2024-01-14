package main

func main() {
	// smell, _ := userJustFarted()
	if err := userJustShitTheirPants(); err != nil {
		changeDaipers()
	}
	if err := bar(); err != nil {
		//handler
	}
}

func bar() error {
	return nil
}

func changeDaipers() {
}

func userJustFarted() (int, error) {
	return 10, nil
}

func userJustShitTheirPants() error {
	return nil
}
