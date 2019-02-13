package main

// Must raises an error if it not nil
func PanicIf(e error) {
	if e != nil {
		panic(e)
	}
}
