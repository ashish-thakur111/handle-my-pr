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
	"github.com/spf13/cobra"
)

// credsCmd represents the creds command
var credsCmd = &cobra.Command{
	Use:   "creds",
	Short: "Command for adding and list current credentials stored in native store",
	Long: `Command for adding and list current credentials stored in native store.
	Secure way of storing credentials rather then storing it in plain files`,
	//Run: func(cmd *cobra.Command, args []string) {},
}

func init() {
	rootCmd.AddCommand(credsCmd)
}
