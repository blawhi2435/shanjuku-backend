/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

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
	"github.com/blawhi2435/shanjuku-backend/database"
	"github.com/blawhi2435/shanjuku-backend/internal/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var version string

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate [flags]",
	Short: "migrate db schema to specify version",
	Example: `db migrate --version 2
db migrate --version latest`,
	Run: func(cmd *cobra.Command, args []string) {
		version := cmd.Flag("version").Value.String()
		svc, err := service.InitService()
		if err != nil {
			logrus.Printf(err.Error())
			return
		}

		database.Migrate(svc.PostgresService, version)
	},
}

func init() {
	dbCmd.AddCommand(migrateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// migrateCmd.PersistentFlags().String("foo", "", "A help for foo")
	migrateCmd.Flags().StringVarP(&version, "version", "v", "", "version of migration, latest or number of version")
	migrateCmd.MarkFlagRequired("version")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// migrateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
