package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/streamingfast/substreams/manifest"
	"google.golang.org/protobuf/proto"
)

var codegenCmd = &cobra.Command{
	Use:          "codeden <package>",
	Short:        "Generate / update your subsreams code ",
	RunE:         runProtogen,
	Args:         cobra.ExactArgs(1),
	SilenceUsage: true,
}

func init() {
	rootCmd.AddCommand(codegenCmd)
}

func runCodeGen(cmd *cobra.Command, args []string) error {
	manifestPath := args[0]
	manifestReader := manifest.NewReader(manifestPath, manifest.SkipSourceCodeReader())

	pkg, err := manifestReader.Read()
	if err != nil {
		return fmt.Errorf("reading manifest %q: %w", manifestPath, err)
	}

	// write the manifest to temp location
	// write buf.gen.yaml with custom stuff
	// run `buf generate`
	// remove if we wrote buf.gen.yaml (--keep-buf-gen-yaml)
	if _, err = manifest.NewModuleGraph(pkg.Modules.Modules); err != nil {
		return fmt.Errorf("processing module graph %w", err)
	}

	defaultFilename := filepath.Join(os.TempDir(), "tmp.spkg")
	cnt, err := proto.Marshal(pkg)
	if err != nil {
		return fmt.Errorf("marshalling package: %w", err)
	}

	if err := ioutil.WriteFile(defaultFilename, cnt, 0644); err != nil {
		fmt.Println("")
		return fmt.Errorf("writing %q: %w", defaultFilename, err)
	}

	_, err = os.Stat("buf.gen.yaml")
	bufFileNotFound := errors.Is(err, os.ErrNotExist)

	if bufFileNotFound {
		content := `
version: v1
plugins:
  - remote: buf.build/prost/plugins/prost:v0.1.3-2
    out: ` + outputPath + `
    opt:
`
		fmt.Println(`Writing to temporary 'buf.gen.yaml':
---
` + content + `
---`)
		if err := ioutil.WriteFile("buf.gen.yaml", []byte(content), 0644); err != nil {
			return fmt.Errorf("error writing buf.gen.yaml: %w", err)
		}
	}

	filepath := filepath.Join(os.TempDir(), "tmp.spkg#format=bin")
	cmdArgs := []string{
		"generate", filepath,
	}
	for _, excludePath := range excludePaths {
		cmdArgs = append(cmdArgs, "--exclude-path", excludePath)
	}
	fmt.Printf("Running: buf %s\n", strings.Join(cmdArgs, " "))
	c := exec.Command("buf", cmdArgs...)
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	if err := c.Run(); err != nil {
		return fmt.Errorf("error executing 'buf':: %w", err)
	}

	if bufFileNotFound {
		fmt.Println("Removing temporary 'buf.gen.yaml'")
		if err := os.Remove("buf.gen.yaml"); err != nil {
			fmt.Errorf("error delefing buf.gen.yaml: %w", err)
		}
	}

	fmt.Println("Done")

	return nil
}
