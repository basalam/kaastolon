// Copyright 2018 Sorint.lab
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"log"
	"os"

	keepercmd "stolon/cmd/keeper/cmd"
	proxycmd "stolon/cmd/proxy/cmd"
	sentinelcmd "stolon/cmd/sentinel/cmd"
	stolonctlcmd "stolon/cmd/stolonctl/cmd"

	"github.com/spf13/cobra/doc"
)

func main() {
	// use os.Args instead of "flags" because "flags" will mess up the man pages!
	var outDir string
	if len(os.Args) == 2 {
		outDir = os.Args[1]
	} else {
		fmt.Fprintf(os.Stderr, "usage: %s [output directory]", os.Args[0])
		os.Exit(1)
	}

	if err := doc.GenMarkdownTree(keepercmd.CmdKeeper, outDir); err != nil {
		log.Fatal(err)
	}
	if err := doc.GenMarkdownTree(sentinelcmd.CmdSentinel, outDir); err != nil {
		log.Fatal(err)
	}
	if err := doc.GenMarkdownTree(proxycmd.CmdProxy, outDir); err != nil {
		log.Fatal(err)
	}
	if err := doc.GenMarkdownTree(stolonctlcmd.CmdStolonCtl, outDir); err != nil {
		log.Fatal(err)
	}
}
