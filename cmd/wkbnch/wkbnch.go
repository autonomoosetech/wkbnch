package main

import (
	"flag"
	_ "github.com/autonomoosetech/schemacan/api/v1"
	"github.com/autonomoosetech/wkbnch/pkg/codegen"
	"log"
	_ "os"
)

type Config struct {
	In     *string
	Out    *string
	Lang   *string
	Verify *bool
}

func main() {
	// configure
	config := Config{
		In:     flag.String("in", "", "input file or directory"),
		Out:    flag.String("out", "", "output directory"),
		Lang:   flag.String("lang", "c", "language to generate"),
		Verify: flag.Bool("verify", false, "don't output, only validate input files"),
	}

	flag.Parse()

	if *config.In == "" {
		log.Fatalln("input file or directory must be defined with the -in flag")
	}

	if *config.Out == "" {
		log.Fatalln("output dirctory must be defined with the -out flag")
	}

	// find files
	files, err := filesFromFlag(*config.In)
	if err != nil {
		log.Fatalf("failed getting files from input given in flag -in: %v", err)
	}

	objects, err := objectsFromFilenames(files)
	if err != nil {
		log.Fatalf("failed parsing in files: %v", err)
	}

	for _, obj := range objects {
		err = obj.Validate()
		if err != nil {
			log.Fatalf("object failed validation: %v", err)
		}
	}

	log.Printf("parsed in %d objects from %d files", len(objects), len(files))

	// generate code

	var c codegen.LangC
	buf, err := codegen.GenerateFiles(c, objects)
	if err != nil {
		log.Fatalf("%v", err)
	}

	log.Println(buf.String())

	// output
}
