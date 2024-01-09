package cmd

import (
	"github.com/spf13/cobra"
)

var dbCmd = &cobra.Command{
	Use:   "db",
	Short: "migrate tables",
	Example: `db migrate --version 2`,
}

func init() {
	rootCmd.AddCommand(dbCmd)

	// Here you will define your flags and configuration settings.
	//dbCmd.PersistentFlags().BoolVar(&isTestEnv, "test", false, "is test env or not")
	//adminCmd.PersistentFlags().StringVar(&sqlInsert, "sql", "serviceConfig.json", "insert sql statement")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dbCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dbCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
