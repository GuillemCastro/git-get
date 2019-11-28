package main

import (
	"io"
	"os"
	"fmt"
	"strings"
	"path/filepath"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"gopkg.in/src-d/go-git.v4/storage/memory"
)

func GetMatchingFiles(files *object.FileIter, name string) ([]*object.File, error) {
	var result []*object.File;
	for file, err := files.Next(); err != io.EOF; file, err = files.Next() {
		if err != nil {
			return nil, err
		}
		fileName := file.Name
		if strings.HasPrefix(fileName, name) {
			result = append(result, file)
		}
	}
	return result, nil
}

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("Usage: %s <remote> <file/path>", os.Args[0])
		os.Exit(1)
	}
	url := os.Args[1]
	filePath := os.Args[2]

	r, _ := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
		URL: url,
	})
	ref, _ := r.Head()
	commit, _ := r.CommitObject(ref.Hash())
	tree, _ := commit.Tree()
	files, _ := GetMatchingFiles(tree.Files(), filePath)
	if len(files) > 1 {
		pathTokens := strings.Split(filePath, "/")
		baseDir := pathTokens[len(pathTokens)-1]
		for _, file := range files {
			relFilePath := filepath.Join(baseDir, strings.TrimPrefix(file.Name, filePath))
			fmt.Println(relFilePath)
			pathToCreate := filepath.Dir(relFilePath)
			if _, err := os.Stat(pathToCreate); os.IsNotExist(err) {
				os.MkdirAll(pathToCreate, os.ModePerm)
			}
			osfile, err := os.Create(relFilePath)
			if (err != nil) {
				fmt.Println(err)
			}
			fileContents, _ := file.Contents()
			osfile.WriteString(fileContents)
		}
	} else {
		file, _ := os.Create(filePath)
		fmt.Println(filePath)
		fileContents, _ := files[0].Contents()
		file.WriteString(fileContents)
	}
}