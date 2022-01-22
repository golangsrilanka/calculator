package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var subCmd = &cobra.Command{
	Use:   "sub",
	Short: "Substract two numbers",
	Long: `Format for Substraction
	
	Format : 
	calculator sub x y 

	Example:
	calculator sub 2 3
	Output for the Example: 
	6

	To input negative numbers use the '--' symbol infront of the numbers

	Format : 
	calculator sub -- x -y 

	Example 1:
	calculator sub -- 2 -3
	Output for the Example 1: 
	5
	
	Example 2:
	calculator sub -- -2 -3
	Output for the Example 2: 
	1`,
	Run: func(cmd *cobra.Command, args []string) {
		num1, _ := strconv.Atoi(args[0])
		num2, _ := strconv.Atoi(args[1])
		result := num1 - num2
		fmt.Println(result)
	},
}

func init() {
	rootCmd.AddCommand(subCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}