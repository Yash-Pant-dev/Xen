package main

import (
	"os"
	"testing"
	"xen/xeni"
)

const testDir = "test"
const examplesDir = "examples"

func TestAllSourceFiles(t *testing.T) {

	source_files, err := os.ReadDir(examplesDir)
	if err != nil {
		t.Log("Error reading example source files directory. \n")
		t.FailNow()
	}

	for _, file := range source_files {
		defer func() {
			if r := recover(); r != nil {
				// fmt.Print("what")
				t.Log("Test failed. <E: ", r, ">")
				t.Fail()
			}
		}()
		params := [1]string{testDir + "/" + examplesDir + "/" + file.Name()}
		xeni.OrchestratePhases(params[:])
	}
}
