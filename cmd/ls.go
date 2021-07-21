/*
Copyright © 2021 Ashish Thakur <ashish.thakur1110@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List the credentials stored in native credentials store",
	Long:  `List the credentials stored in native credentials store`,
	Run: func(cmd *cobra.Command, args []string) {
		DoLs()
	},
}

func DoLs() {
	entries, err := passStore.List()
	if err != nil {
		panic("getting list of credentials from native credentails manager failed")
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"name/server", "username"})
	table.SetRowLine(true)
	data := [][]string{}
	for k, e := range entries {
		data = append(data, []string{k, e})
	}
	for _, v := range data {
		table.Append(v)
	}
	table.Render()
}

func init() {
	credsCmd.AddCommand(lsCmd)
}
