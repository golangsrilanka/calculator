package cmd

import (
	"calculator/internal"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// mulCmd represents the mul command
var subCmd = &cobra.Command{
	Use:   "sub",
	Short: "Subtract two numbers",
	Long:  `Subtract one number from another`,
	Run: func(cmd *cobra.Command, args []string) {
		num1, _ := strconv.ParseUint(args[0], 10, 32)
		num2, _ := strconv.ParseUint(args[1], 10, 32)
		tot := internal.SubtractTwoNumbers(num1, num2)
		fmt.Println(tot)
	},
}

func init() {
	rootCmd.AddCommand(subCmd)
}
