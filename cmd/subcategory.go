package cmd

import (
	"github.com/spf13/cobra"
)

var subcategoryCmd = &cobra.Command{
	Use:   "subcategory",
	Short: "Manage documentation subcategories",
	Long:  `The subcommands of subcategory are used to manage documentation subcategories`,
}

func init() {
	rootCmd.AddCommand(subcategoryCmd)

	subcategoryCmd.PersistentFlags().String("subcategories-file", "", "Path to newline separated file of allowed data source and resource frontmatter subcategories")
}
