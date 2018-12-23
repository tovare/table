package table

import (
	"log"
	"testing"
)

func TestReadCSV(T *testing.T) {
	ReadCSV("hello.csv")
	log.Println("Hello")
}

func TestDetectTypes(T *testing.T) {
	t1 := `
    a,b,c
    1,2,3
    1,2,3`
	DetectTypes(t1)
}
