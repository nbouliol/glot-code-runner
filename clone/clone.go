package clone

import (
	"gopkg.in/src-d/go-git.v4"
	"fmt"
	"os"
	"io/ioutil"
	"log"
)

func Clone(repoUrl string) string {

	// Clone options : https://godoc.org/gopkg.in/src-d/go-git.v4#CloneOptions

	// Tempdir to clone the repository
	dir, err := ioutil.TempDir("", "clone-example")
	if err != nil {
		log.Fatal(err)
	}

	//defer os.RemoveAll(dir)

	//fmt.Println(dir)
	_,err = git.PlainClone(dir, false, &git.CloneOptions{
		URL: repoUrl,
		Progress: os.Stdout,
	})

	if err != nil {
		fmt.Println(err)
	}
	return dir

}

