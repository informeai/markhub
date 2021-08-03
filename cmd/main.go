package main

import (
	"fmt"
	mkd "github.com/informeai/markhub"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("markhubserve:\nServe files .md the localhost:4000.\nUsage: markhubserve [file.md]")
		os.Exit(0)
	} else {
		filePath := os.Args[1]
		m := mkd.NewMarkHub()
		err := m.ParseFile(filePath)
		if err != nil {
			log.Fatalln(err)
		}
		err = m.Serve(":4000")
		if err != nil {
			log.Fatalln(err)
		}
	}
}
