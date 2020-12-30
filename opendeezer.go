package main

import (
	"os"
	"os/exec"
)

func main() {
	site := os.Args[1]
	site = "deezer://" + site

	println(site)
	c := exec.Command("cmd.exe", "/C", "start", "", site)

	/*if err := c.Run(); err != nil {
		fmt.Println("Error: ", err)
	}*/
	c.Start()
}
