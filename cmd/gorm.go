/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

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
	"github.com/blawhi2435/shanjuku-backend/database/postegre"
	"github.com/blawhi2435/shanjuku-backend/internal/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// gormCmd represents the gorm command
var gormCmd = &cobra.Command{
	Use:   "gorm",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		tables := []interface{}{postegre.User{}, postegre.Group{}, postegre.Schedule{}, postegre.Activity{}}

		s, err := service.InitService()
		if err != nil {
			logrus.Panicf("start dbService failed - %s", err)
		}

		var n int
		count := 0
		for _, tbl := range tables {
			err = s.PostgreService.DB.AutoMigrate(tbl)
			if err != nil {
				panic(err)
			}
			count++
		}
		n = count
		logrus.Infof("migrate %d tables", n)
	},
}

func init() {
	rootCmd.AddCommand(gormCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// gormCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// gormCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
