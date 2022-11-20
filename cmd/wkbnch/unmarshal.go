package main

import (
	"bytes"
	"github.com/autonomoosetech/schemacan/api/v1"
	"gopkg.in/yaml.v3"
	"io"
	"log"
	"os"
)

func objectsFromFilenames(files []string) (objects []api.Object, err error) {
	log.Println("Parsing in:")
	for _, file := range files {
		log.Println("\t" + file)

		dat, err := os.ReadFile(file)
		if err != nil {
			log.Fatalln(err)
		}

		decoder := yaml.NewDecoder(bytes.NewBuffer(dat))

		for {
			var obj api.Object

			err := decoder.Decode(&obj)
			if err != nil {
				if err == io.EOF {
					break
				}
				log.Fatalf("file decode failed: %v", err)
			}

			objects = append(objects, obj)
		}
	}

	return
}
