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
	"fmt"

	"github.com/ashish-thakur111/handle-my-pr/pkg/osHelper"
	"github.com/docker/docker-credential-helpers/credentials"
	"github.com/docker/docker-credential-helpers/pass"
	"github.com/docker/docker-credential-helpers/wincred"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add remote server credentials",
	Long:  `Add remote server credentials for github, gitlab, gittea etc.`,
	Run: func(cmd *cobra.Command, args []string) {
		DoAdd()
	},
}

type cred struct {
	server   string
	username string
	secret   string
}

var opts cred

var passStore = pass.Pass{}

var winStore = wincred.Wincred{}

func DoAdd() {
	c := &credentials.Credentials{
		ServerURL: opts.server,
		Username:  opts.username,
		Secret:    opts.secret,
	}
	fmt.Println("Adding credentials to native store")
	var err error
	if osHelper.GetOperatingSystem() == osHelper.Windows {
		winStore.Add(c)
	} else {
		err = passStore.Add(c)
	}
	if err != nil {
		fmt.Println("Adding credentials to pass credential manager failed")
		panic(err)
	}
	fmt.Println("Adding credentials finished")
}

func init() {
	credsCmd.AddCommand(addCmd)

	addCmd.Flags().StringVarP(&opts.server, "server", "s", "github", "name of server github- github, gitlab- gitlab")
	addCmd.Flags().StringVarP(&opts.username, "username", "u", "", "user name or owner")
	addCmd.Flags().StringVarP(&opts.secret, "token", "t", "", "API token for authentication")
	addCmd.MarkFlagRequired("username")
	addCmd.MarkFlagRequired("token")
}
