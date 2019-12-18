package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"

	flag "github.com/ogier/pflag"
)

var (
	limit     int
	isQuiet   bool
	searchVal string
)

func main() {
	flag.Parse()

	if len(os.Args) < 2 {
		fmt.Print("Search value must be provided!\n")
		return
	}

	searchVal = os.Args[1]

	if strings.HasPrefix(searchVal, "-") {
		fmt.Print("Search value must be the first argument!\n")
		return
	}

	var command = "kubectl get pods"

	if isQuiet {
		command += " --template '{{range .items}}{{.metadata.name}}{{\"\\n\"}}{{end}}'"
	}

	command += " | grep " + searchVal

	if limit > 0 {
		command += " | head -n " + strconv.Itoa(limit)
	}

	result, err := exec.Command("bash", "-c", command).Output()

	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Printf("%s", result)
}

func init() {
	flag.IntVarP(&limit, "limit", "l", 0, "Sets limit of the displayed row")
	flag.BoolVarP(&isQuiet, "quiet", "q", false, "Indicates if only names of pods are needed")
}
