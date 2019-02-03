package sampleModule

import (
	"fmt"
	"io"
	"os"
	"path"
	"strings"
)

func ArrayOne() {
	sr := strings.NewReader("hello string reader!")

	buf := make([]byte, 8)
	for {
		n, err := sr.Read(buf)
		fmt.Printf("n = %v, err = %v, buf = %v\n", n, err, buf)
		if err == io.EOF {
			break
		}
	}

	emails := make(map[string]string)
	emails["Josh"] = "jsk9260@gmail.com"

	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	filepath := path.Join(wd, "test.txt")
	w, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, 0755)
	defer w.Close()
	if err != nil {
		panic("hhh")
	}
	bb := []byte(emails["Josh"])
	wrote, _ := w.Write(bb)
	fmt.Println(filepath + " : " + string(wrote))


}