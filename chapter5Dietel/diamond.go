package main

func diamond() {
	print5Asterisk()

}

func print5Asterisk() {
	for asterisk := 0; asterisk < 5; asterisk++ {
		for count := asterisk; count < 5; count++ {
			print(" * ")
		}
		println("")
	}
}
