package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"sort"
	"strings"
)

var Handler string = os.Getenv("REDIRECT_HANDLER")
var Source string = os.Getenv("PATH_TRANSLATED")
var Query string = os.Getenv("QUERY_STRING")

func main() {
	conf := make([]string, 0)
	confArray, err := readLines("/etc/pandoc-apache.conf")
	if err == nil {
		conf = filterComments(confArray)
	}
	_ = conf
	switch Query {
	case "raw":
		Raw(Source)
		return
	case "env":
		Env()
		return
	case "plain":
		Plain("pandoc", append(conf, "--from=markdown", "--to=plain", Source))
	case "html4", "xhtml":
		HTML("pandoc", append(conf, "--from=markdown", "--to=html", Source))
	default:
		HTML("pandoc", append(conf, "--from=markdown", "--to=html5", Source))
	}
}

func HTML(name string, args []string){
	cmd := exec.Command(name, args ...)
	result, _ := cmd.Output()
	fmt.Println("Content-type: text/html\n")
	fmt.Println(string(result))
}

func Plain(name string, args []string){
	cmd := exec.Command(name, args ...)
	result, _ := cmd.Output()
	fmt.Println("Content-type: text/plain\n")
	fmt.Println(string(result))
}

func Raw(path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Status: 500 Internal Server Error\n")
		return
	}
	defer file.Close()

	fmt.Println("Content-type: text/plain\n")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

}

func Env() {
	fmt.Println("Content-type: text/html\n")
	fmt.Println("<doctype HTML>")
	fmt.Println("<link rel=\"stylesheet\" href=\"/css/style.css\">")
	fmt.Println("<style>body {max-width: 100%; margin: 1em}</style>")
	fmt.Printf("<p>Handler: %v\n", Handler)
	fmt.Printf("<p>Source:  %v\n", Source)
	fmt.Printf("<p>Query:   %v\n", Query)
	fmt.Println("<p>")
	
	var env []string = os.Environ()
	sort.Strings(env)
	fmt.Println("<p>")
	for _, s := range env {
		fmt.Printf("<br>%v\n", s)
	}
	
	//~ fmt.Printf("<p>%v", os.Environ())
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func filterComments (ss []string) (fss []string) {
	for _, s := range ss {
		if strings.HasPrefix(s, "#") {
			continue
		}
		fss = append (fss, s)
	}
	return
}
