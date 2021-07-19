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
	"context"
	"fmt"
	"os"
	"strconv"

	pkgGithub "github.com/ashish-thakur111/handle-my-pr/pkg/vendors/github"
	"github.com/google/go-github/v37/github"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the PRs in a project",
	Long: `List all the Pull Requests in the project. Result will be 
	displayed in ascii table form which includes PR number, title and status`,
	Run: func(cmd *cobra.Command, args []string) {
		DoList()
	},
}

func DoList() {
	owner, _ := rootCmd.PersistentFlags().GetString("owner")
	repo, _ := rootCmd.PersistentFlags().GetString("repo")
	githubClient := pkgGithub.NewGithubClientWithoutAuth()
	prOpts := &github.PullRequestListOptions{
		State: "open",
	}

	prs, _, err := githubClient.PullRequests.List(context.Background(), owner, repo, prOpts)

	if err != nil {
		panic(err)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"number", "description", "status"})
	table.SetRowLine(true)
	fmt.Println("Displaying Pull Request Information")
	data := [][]string{}
	for _, pr := range prs {
		data = append(data, []string{strconv.Itoa(*pr.Number), *pr.Title, *pr.State})
	}
	for _, v := range data {
		table.Append(v)
	}
	table.Render()
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
