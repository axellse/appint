package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Provide an appimage to integrate.")
		return
	}

	file, err := filepath.Abs(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	aa := strings.Split(file, "/")
	name := strings.Split(aa[len(aa)-1], ".")[0]
	var desktopFile = `[Desktop Entry]
Name=` + name + `
Exec=` + file + `
Type=Application
Path=` + strings.Join(aa[:len(aa)-1], "/")

	werr := os.WriteFile(path.Join("/usr/share/applications", name + ".desktop"), []byte(desktopFile), 0777)
	if werr != nil {
		log.Fatal(werr)
		fmt.Println("Couldnt create file, make sure to run as sudo.")
	}

	fmt.Println("Done, make sure to chmod +x the appimage if you haven't already.")
}