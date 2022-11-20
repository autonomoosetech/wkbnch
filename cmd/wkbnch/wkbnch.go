package main

import (
	"flag"
	"github.com/autonomoosetech/schemacan/api/v1"
	"log"
	"os"
)

type Config struct {
	In       *string
	Out      *string
	Lang     *string
	Validate *bool
}

func main() {
	// configure
	config := Config{
		In:       flag.String("in", "", "input file or directory"),
		Out:      flag.String("out", "", "output directory"),
		Lang:     flag.String("lang", "c", "language to generate"),
		Validate: flag.Bool("validate", false, "don't output, only validate input files"),
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

	slot := api.Slot{
		Max:    100.0,
		Min:    0.0,
		Offset: 0,
		Size:   8,
		Unit:   "mV",
	}

	slots := []api.Slot{slot}

	out, err := templateSlots(slots)
	if err != nil {
		panic(err)
	}

	log.Println(*config.Out + "/slot.h")

	f, err := os.Create(*config.Out + "/slot.h")
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer f.Close()
	_, err = f.Write(out.Bytes())
	if err != nil {
		log.Fatalf("%v", err)
	}

	log.Println(out)

	// output
}
