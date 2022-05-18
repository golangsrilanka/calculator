package cmd

import (
	"calculator/internal"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add two numbers",
	Long:  `Add two numbers together`,
	Run: func(cmd *cobra.Command, args []string) {
		num1, _ := strconv.ParseUint(args[0], 10, 32)
		num2, _ := strconv.ParseUint(args[1], 10, 32)
		tot := internal.AddTwoNumbers(num1, num2)
		fmt.Println(tot)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
