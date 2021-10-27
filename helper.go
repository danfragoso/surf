package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		exec.Command("surf").Start()
		os.Exit(0)
	}

	xpropOut, err := exec.Command("xprop", "-root", "32x", "'\t$0'", "_NET_ACTIVE_WINDOW").Output()
	if err != nil {
		log.Fatal(err)
	}

	winId := strings.Split(string(xpropOut), "0x")[1]
	winId = "0x" + strings.Trim(winId, "'")

	parsedId, err := strconv.ParseInt(winId, 0, 64)
	if err != nil {
		os.Exit(0)
	}

	location := os.Args[1]
	if os.Args[0] == "s" {
		location = "https://google.com.br/search?q=" + strings.Join(os.Args[1:], " ")
	}

	_, err = exec.Command("xprop",
		"-id", strconv.FormatInt(parsedId, 10), "-f", "_SURF_GO", "8u", "-set", "_SURF_GO", location,
	).Output()

	fmt.Printf("The date is %d %s\n", parsedId, location)
}
