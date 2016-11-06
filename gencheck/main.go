package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/abice/gencheck/generator"
	"github.com/mkideal/cli"
)

type rootT struct {
	cli.Helper
	FileNames       []string `cli:"*f,file" usage:"The file(s) to generate validations for.  Use more than one flag for more files."`
	PointerReceiver bool     `cli:"ptr" usage:"turn on pointer receiver generation for validation method." dft:"false"`
	CustomTemplates []string `cli:"t,template" usage:"custom template files"`
	TemplateDirs    []string `cli:"d,template-dir" usage:"custom template folders"`
	FailFast        bool     `cli:"failfast" usage:"Tell the generator to fail all structs fast"`
}

func main() {
	cli.Run(new(rootT), func(ctx *cli.Context) error {
		argv := ctx.Argv().(*rootT)

		for _, fileName := range argv.FileNames {

			g := generator.NewGenerator()

			if argv.PointerReceiver {
				g.WithPointerMethod()
			}

			if argv.FailFast {
				g.WithFailFast()
			}

			if len(argv.CustomTemplates) > 0 {
				ctx.String("Adding templates=%s\n", ctx.Color().Grey(argv.CustomTemplates))

				// Add them one by one, because it's really just a warning if we can't load one of their templates.
				for _, templ := range argv.CustomTemplates {
					err := g.AddTemplateFiles(argv.CustomTemplates...)
					if err != nil {
						ctx.String("Unable to add template '%s': %s\n", ctx.Color().Cyan(templ), ctx.Color().Yellow(err))
					}
				}
			}

			for _, dir := range argv.TemplateDirs {
				ctx.String("Scanning templates in %s\n", ctx.Color().Grey(dir))

				// Walk the directory and add each file one by one.
				filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
					if !info.IsDir() {
						ctx.String("Adding template %s\n", ctx.Color().Grey(path))
						err = g.AddTemplateFiles(path)
						if err != nil {
							ctx.String("Unable to add template '%s': %s\n", ctx.Color().Cyan(path), ctx.Color().Yellow(err))
						}
					}
					return err
				})
			}
			originalName := fileName

			ctx.String("gencheck started. file: %s\n", ctx.Color().Cyan(originalName))
			fileName, _ = filepath.Abs(fileName)
			outFilePath := fmt.Sprintf("%s_validators.go", strings.TrimSuffix(fileName, filepath.Ext(fileName)))

			// Parse the file given in arguments
			raw, err := g.GenerateFromFile(fileName)
			if err != nil {
				return fmt.Errorf("Error while generating validators\nInputFile=%s\nError=%s\n", ctx.Color().Cyan(fileName), ctx.Color().RedBg(err))
			}

			err = ioutil.WriteFile(outFilePath, raw, os.ModePerm)
			if err != nil {
				return fmt.Errorf("Error while writing to file %s: %s\n", ctx.Color().Cyan(outFilePath), ctx.Color().Red(err))
			}
			ctx.String("gencheck finished. file: %s\n", ctx.Color().Cyan(originalName))
		}

		return nil
	})
	return
}
