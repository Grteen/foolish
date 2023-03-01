package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
)

func main() {
	grep := `netstat -anp | grep `
	ports := []string{"9877", "8080", "8081", "8082", "8083", "8084", "8085"}
	execs := []string{"main", "user", "artical", "search", "notify", "action", "comment"}

	for t, i := range ports {
		cmd := exec.Command("sh", "-c", grep+i)
		str, err := cmd.Output()
		if err != nil {
			log.Fatal(err)
		}

		reg := regexp.MustCompile(`([0-9]+)/./` + execs[t])
		pid := reg.FindStringSubmatch(string(str))

		file, err := os.OpenFile("../pid/"+execs[t]+".pid", os.O_CREATE|os.O_RDWR, os.ModeAppend)
		if err != nil {
			log.Fatal(err)
		}
		_, err = fmt.Fprintf(file, "%s", pid[1])
		if err != nil {
			log.Fatal(err)
		}
	}
}
