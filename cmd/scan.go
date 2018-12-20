// Copyright © 2018 Benny Gächter <benny.gaechter@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

type packageJSON struct {
	Name    string
	Version string
	License string
}

// scanDirectory scans the directory and its subdirectory for package.json
// files and print their name, version and license information
func scanDirectory(path string, f os.FileInfo, err error) error {
	_, file := filepath.Split(path)
	if file == "package.json" {
		b, err := ioutil.ReadFile(path)
		if err != nil {
			fmt.Print(err)
		}
		var dependency packageJSON
		err = json.Unmarshal(b, &dependency)
		if err != nil {
			fmt.Print(err)
		}
		fmt.Printf("%s: Version %s, License: %s \n", dependency.Name, dependency.Version, dependency.License)
	}
	return nil
}

// scanCmd represents the scan command
var scanCmd = &cobra.Command{
	Use:   "scan [path]",
	Short: "scan for package.json files in [path] and print their licence information.",
	Long:  `scan for package.json files in [path] and print their licence information.`,
	Run: func(cmd *cobra.Command, args []string) {
		path := args[0]
		err := filepath.Walk(path, scanDirectory)
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)
}
