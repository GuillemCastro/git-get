package main

import (
	"io"
	"os"
	"fmt"
	"strings"
	"path/filepath"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"gopkg.in/src-d/go-git.v4/storage/memory"
	"github.com/jessevdk/go-flags"
)

//The Arguments struct is where we save the program arguments
type Arguments struct {
	ReqArgs struct {
		URL   string `description:"URL to a Git repository"`
		Path  string `description:"File or directory to download"`
	} `positional-args:"yes" required:"yes"`
	Branch string `short:"b" long:"branch" description:"Branch, tag or commit hash"`
}

func parseArgs() (Arguments, error) {
	var opts Arguments
	parser := flags.NewParser(&opts, flags.Default)
	_, err := parser.Parse()
	if err != nil {
		if e, ok := err.(*flags.Error); ok && e.Type == flags.ErrRequired {
			parser.WriteHelp(os.Stdout)
		}
		os.Exit(1)
	}
	return opts, err
}

func getMatchingFiles(files *object.FileIter, name string) ([]*object.File, error) {
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

func checkError(err error) {
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	} 
}

func main() {
	args, err := parseArgs()
	checkError(err)
	url := args.ReqArgs.URL
	filePath := args.ReqArgs.Path
	r, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
		URL: url,
	})
	checkError(err)
	revision := plumbing.Revision(string(plumbing.HEAD))
	if args.Branch != "" {
		revision = plumbing.Revision(args.Branch)
	}
	hash, err := r.ResolveRevision(revision)
	commit, err := r.CommitObject(*hash)
	checkError(err)
	tree, err := commit.Tree()
	checkError(err)
	files, err := getMatchingFiles(tree.Files(), filePath)
	checkError(err)
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
			checkError(err)
			fileContents, err := file.Contents()
			checkError(err)
			_, err = osfile.WriteString(fileContents)
			checkError(err)
		}
	} else {
		file, _ := os.Create(filePath)
		fmt.Println(filePath)
		fileContents, err := files[0].Contents()
		checkError(err)
		_, err = file.WriteString(fileContents)
		checkError(err)
	}
}