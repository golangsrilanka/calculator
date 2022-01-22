package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var divCmd = &cobra.Command{
	Use:   "div",
	Short: "Divide two numbers",
	Long: `Format for Division
	
	Format : 
	calculator div x y 

	Example:
	calculator div 10 2
	Output for the Example: 
	5

	To input negative numbers use the '--' symbol infront of the numbers

	Format : 
	calculator div -- x -y 

	Example 1:
	calculator div -- 10 -2
	Output for the Example 1: 
	-5
	
	Example 2:
	calculator div -- -10 -2
	Output for the Example 2: 
	5`,
	Run: func(cmd *cobra.Command, args []string) {
		num1, _ := strconv.ParseFloat(args[0], 64)
		num2, _ := strconv.ParseFloat(args[1], 64)
		result := num1 / num2
		fmt.Println(result)
	},
}

func init() {
	rootCmd.AddCommand(divCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}