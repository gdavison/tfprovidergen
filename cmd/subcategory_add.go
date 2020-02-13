package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a documentation subcategory",
	Long: `Adds a subcategory to the provider documentation.

If a subcategories-file is passed, the new subcategory is added to the list of allowed subcategories.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			errorExit("subcategory needs a subcategory name")
		}

		subcategoryName := args[0]

		allowedSubcategoriesFile, err := cmd.Flags().GetString("subcategories-file")
		if err != nil {
			errorExitf("error reading \"subcategories-file\": %s", err)
		}

		if allowedSubcategoriesFile != "" {
			subcategories, err := readAllowedSubcategoriesFile(allowedSubcategoriesFile)
			if err != nil {
				errorExit(err)
			}
			index := sort.SearchStrings(subcategories, subcategoryName)
			if subcategories[index] != subcategoryName {
				subcategories = append(subcategories, "")
				copy(subcategories[index+1:], subcategories[index:])
				subcategories[index] = subcategoryName
				err = writeAllowedSubcategoriesFile(allowedSubcategoriesFile, subcategories)
				if err != nil {
					errorExit(err)
				}
			} else {
				log.Printf("Subcategory \"%s\" already found in %s", subcategoryName, allowedSubcategoriesFile)
			}
		}

		fmt.Println("done.")
	},
}

func init() {
	subcategoryCmd.AddCommand(addCmd)
}

func readAllowedSubcategoriesFile(path string) ([]string, error) {
	log.Printf("[DEBUG] Loading allowed subcategories file: %s", path)

	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("error opening allowed subcategories file (%s): %w", path, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var allowedSubcategories []string

	for scanner.Scan() {
		allowedSubcategories = append(allowedSubcategories, scanner.Text())
	}

	if err != nil {
		return nil, fmt.Errorf("error reading allowed subcategories file (%s): %w", path, err)
	}

	return allowedSubcategories, nil
}

func writeAllowedSubcategoriesFile(path string, allowedSubcategories []string) error {
	log.Printf("[DEBUG] Saving allowed subcategories file: %s", path)

	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("error opening allowed subcategories file (%s): %w", path, err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range allowedSubcategories {
		_, err = writer.WriteString(line + "\n")
		if err != nil {
			return fmt.Errorf("error writing allowed subcategories file (%s): %w", path, err)
		}
	}
	err = writer.Flush()
	if err != nil {
		return fmt.Errorf("error flushing allowed subcategories file (%s): %w", path, err)
	}

	return nil
}
