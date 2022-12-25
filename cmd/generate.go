package cmd

import (
	"bytes"
	"io"
	"log"
	"os"
	"strings"

	"github.com/autonomoosetech/schemacan/api/v1"
	"github.com/autonomoosetech/wkbnch/pkg/codegen"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate source files",
	Long:  `Use SchemaCAN definitions to generate source files`,
	Run: func(cmd *cobra.Command, args []string) {
		// find files
		files, err := filesFromFlag(inputPath)
		if err != nil {
			log.Fatalf("failed getting files from input given in flag -in: %v", err)
		}

		// read in objects from files
		objects, err := objectsFromFilenames(files)
		if err != nil {
			log.Fatalf("failed parsing in files: %v", err)
		}

		// validate objects
		for _, obj := range objects {
			err = obj.Validate()
			if err != nil {
				log.Fatalf("object failed validation: %v", err)
			}
		}

		// show object summary
		log.Printf("parsed in %d objects from %d files", len(objects), len(files))

		// generate code
		var c codegen.LangC
		buf, err := codegen.GenerateFiles(c, objects)
		if err != nil {
			log.Fatalf("%v", err)
		}

		// output to console
		log.Println(buf.String())
	},
}

var (
	inputPath      string
	outputPath     string
	targetLanguage string
)

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.Flags().StringVarP(&inputPath, "input", "i", "schemacan.yaml", "input file or directory")
	generateCmd.Flags().StringVarP(&outputPath, "output", "o", "./src/schemacan", "output directory")
	generateCmd.Flags().StringVarP(&targetLanguage, "lang", "l", "", "target language")

}

func validFileExt(name string) bool {
	return strings.HasSuffix(name, ".yaml") || strings.HasSuffix(name, ".yml")
}

func filesInDir(dir string) (files []string, err error) {
	// ensure trailing slash this may cause a double slash at the end but we don't care
	dir += "/"

	fileList, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, file := range fileList {
		name := file.Name()
		if validFileExt(name) {
			files = append(files, dir+name)
		}
	}

	return
}

func filesFromFlag(input string) (files []string, err error) {
	if validFileExt(input) {
		// treat as file
		files = append(files, input)
	} else {
		// treat as directory
		files, err = filesInDir(input)
	}

	return
}

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
