package cmd

import (
	"calculator/internal"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// mulCmd represents the mul command
var divCmd = &cobra.Command{
	Use:   "div",
	Short: "Divide two numbers",
	Long:  `Divide one number by another`,
	Run: func(cmd *cobra.Command, args []string) {
		num1, _ := strconv.ParseUint(args[0], 10, 32)
		num2, _ := strconv.ParseUint(args[1], 10, 32)
		tot := internal.DivideTwoNumbers(num1, num2)
		fmt.Println(tot)
	},
}

func init() {
	rootCmd.AddCommand(divCmd)
}
