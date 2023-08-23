package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"io/ioutil"
)

const templateContent = `package main

import (
	"fmt"
)

func main() {
	
}
`

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run script.go <directory_name>")
		return
	}

	newDirName := os.Args[1]

	// Get a list of all directories in the current directory
	dirs, err := getDirectories(".")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Find the highest number prefix among existing directories
	highestNumber := findHighestNumber(dirs)

	// Increment the highest number and create the new directory
	newNumber := highestNumber + 1
	newDirPrefix := fmt.Sprintf("%02d", newNumber) // Pad with leading zero if necessary
	newDir := newDirPrefix + "_" + newDirName

	err = os.Mkdir(newDir, 0755)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Create the Go file inside the new directory
	goFilePath := filepath.Join(newDir, newDir+".go")
	err = ioutil.WriteFile(goFilePath, []byte(templateContent), 0644)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("New submodule created:", newDir)
}

func getDirectories(rootDir string) ([]string, error) {
	var dirs []string
	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && isNumberedDir(info.Name()) {
			dirs = append(dirs, info.Name())
		}
		return nil
	})
	return dirs, err
}

func isNumberedDir(name string) bool {
	match, _ := regexp.MatchString(`^\d+_.+`, name)
	return match
}

func findHighestNumber(dirs []string) int {
	var numbers []int
	for _, dir := range dirs {
		parts := regexp.MustCompile(`(\d+)_.*`).FindStringSubmatch(dir)
		if len(parts) == 2 {
			number, _ := strconv.Atoi(parts[1])
			numbers = append(numbers, number)
		}
	}

	if len(numbers) == 0 {
		return 0
	}

	sort.Ints(numbers)
	return numbers[len(numbers)-1]
}
