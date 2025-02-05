package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// print out all the metadata available about a file
	// Stat returns file info. It will return an error if there is no file.
	fileInfo, err := os.Stat("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("File name: ", fileInfo.Name())
	fmt.Println("Size in bytes: ", fileInfo.Size())
	fmt.Println("Permission: ", fileInfo.Mode())
	fmt.Println("Last modified: ", fileInfo.ModTime())
	fmt.Println("Is Directory: ", fileInfo.IsDir())
	fmt.Printf("System interface type: %T\n", fileInfo.Sys())
	fmt.Printf("System info: %+v\n\n", fileInfo.Sys())
}

/*
File name:  test.txt
Size in bytes:  100
Permission:  -rw-r--r--
Last modified:  2022-11-16 21:47:59.710696818 +0800 CST
Is Directory:  false
System interface type: *syscall.Stat_t
System info: &{Dev:16777220 Mode:33188 Nlink:1 Ino:4317384711 Uid:502 Gid:20 Rdev:0 Pad_cgo_0:[0 0 0 0] Atimespec:{Sec:1668611295 Nsec:772573901} Mtimespec:{Sec:1668606479 Nsec:710696818} Ctimespec:{Sec:1668606479 Nsec:710696818} Birthtimespec:{Sec:1668605439 Nsec:723992999} Size:100 Blocks:0 Blksize:4096 Flags:0 Gen:0 Lspare:0 Qspare:[0 0]}
*/
