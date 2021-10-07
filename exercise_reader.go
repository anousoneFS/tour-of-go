package main

// import "golang.org/x/tour/reader"

type MyReader struct{}

// todo: return ASCII OF A
//* question: why b arguement is not poiter
func (r MyReader) Read(b []byte) (int, error) {
	for i := range b {
		b[i] = 65
	}
	return len(b), nil
}

func ExerciseReaders() {
	// reader.Validate(MyReader{})
}
