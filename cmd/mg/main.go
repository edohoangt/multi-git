package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/edohoangt/multi-git/pkg/repo_manager"
)

func main() {
	command 	:= flag.String("command", "", "The git command")
	ignoreErros := flag.Bool("ignore-errors", false, "Keep running after error if true")
	flag.Parse()

	// Get managed repos from environment variables
	root := os.Getenv("MG_ROOT")
	if root[len(root) - 1] != '/' {
		root += "/"
	}

	repo_names := strings.Split(os.Getenv("MG_REPOS"), ",")
	repoManager, err := repo_manager.NewRepoManager(root, repo_names, *ignoreErros)
	if err != nil {
		log.Fatal(err)
	}

	output, _ := repoManager.Exec(*command)
	for repo, out := range output {
		fmt.Printf("[%s]: git %s\n", path.Base(repo), *command)
		fmt.Printf(out)
	}

	fmt.Println("Done.")
}

