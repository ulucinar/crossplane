package main

import (
	"os"

	v1 "github.com/crossplane/crossplane/apis/apiextensions/v1"
)

func main() {
	buff, err := os.ReadFile("/Users/alper/Downloads/clusterfuzz-testcase-minimized-fuzz_patch_apply-4520309021081600")
	if err != nil {
		panic(err)
	}
	v1.FuzzPatchApply(buff)
}
