// package main

// import (
// 	"embed"
// 	_ "embed"
// 	"fmt"
// 	"io/fs"
// 	"io/ioutil"
// 	"testing"
// )

// //go:embed version.txt
// var version string

// func TestString(t *testing.T) {
// 	fmt.Println(version)
// }

// //go:embed Capture.PNG
// var image []byte

// func TestEmbed(t *testing.T) {
// 	err := ioutil.WriteFile("Capture1.PNG", image, fs.ModePerm)
// 	if err != nil {
// 		panic(err)
// 	}
// }
// //go:embed files/a.txt
// //go:embed files/b.txt
// //go:embed files/c.txt
// var files embed.FS
// func TestEmbedMulti(t *testing.T) {
// 	a, _ := files.ReadFile("files/a.txt")
// 	fmt.Println(string(a))
// 	b, _ := files.ReadFile("files/b.txt")
// 	fmt.Println(string(b))
// 	c, _ := files.ReadFile("files/c.txt")
// 	fmt.Println(string(c))

// }

// //go:embed files/*.txt
// var path embed.FS

// func TestPatchmatcher(t *testing.T){
// 	dir, _ := path.ReadDir("files")
// 	for _, entry := range dir {
// 		if !entry.IsDir() {
// 			fmt.Println(entry.Name())
// 			file, _ := path.ReadFile("files/"+ entry.Name())
// 			fmt.Println(string(file))
// 		}
// 	}
// }