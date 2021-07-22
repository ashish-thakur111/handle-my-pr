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

	pkgGithub "github.com/ashish-thakur111/handle-my-pr/pkg/vendors/github"
	"github.com/google/go-github/v37/github"
	"github.com/spf13/cobra"
)

// mergeCmd represents the merge command
var mergeCmd = &cobra.Command{
	Use:   "merge",
	Short: "merge the opened pull request",
	Long:  `merge the opened pull request to the target branch`,
	Run: func(cmd *cobra.Command, args []string) {
		prID, _ := cmd.Flags().GetInt("pr-id")
		mergeMethod, _ := cmd.Flags().GetString("merge-method")
		if mergeMethod == "" || prID == 0 {
			cmd.Help()
			return
		}
		DoMerge(prID, mergeMethod)
	},
}

func DoMerge(prID int, mergeMethod string) {
	owner, _ := rootCmd.PersistentFlags().GetString("owner")
	repo, _ := rootCmd.PersistentFlags().GetString("repo")
	if repo == "" {
		panic("repo flag is required")
	}
	username, secret, err := passStore.Get("github")
	if err != nil {
		panic(err)
	}
	if owner == "" {
		owner = username
	}

	githubClient := pkgGithub.NewGithubClientWithAuth(secret)
	opts := &github.PullRequestOptions{
		CommitTitle:        "Merge pull request #" + fmt.Sprint(prID),
		MergeMethod:        mergeMethod,
		DontDefaultIfBlank: false,
	}

	mergeMessage, response, err := githubClient.PullRequests.Merge(context.Background(),
		owner, repo, prID, "merged", opts)

	if err != nil {
		panic(err)
	}
	if (response.StatusCode != 200) && (response.StatusCode != 201) {
		fmt.Println(response.Status)
		panic(response.Status)
	}

	fmt.Println(*mergeMessage.Merged)
	fmt.Println(*mergeMessage.SHA)
	fmt.Println(*mergeMessage.Message)
}

func init() {
	rootCmd.AddCommand(mergeCmd)
	mergeCmd.Flags().IntP("pr-id", "i", 0, "unique PR number")
	mergeCmd.Flags().StringP("merge-method", "m", "squash",
		"merge method options are merge, squash and rebase. Default is squash")
	mergeCmd.MarkFlagRequired("pr-id")
}
