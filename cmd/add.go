/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

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
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add two numbers",
	Long: `Format for Addition
	
	Format : 
	calculator add x y 

	Example:
	calculator add 2 3
	Output for the Example: 
	5

	To input negative numbers use the '--' symbol infront of the numbers

	Format : 
	calculator add -- x -y 
	Example 1:

	calculator add -- 2 -3
	Output for the Example 1: 
	-1
	
	Example 2:
	calculator add -- -2 -3
	Output for the Example 2: 
	-5`,
	Run: func(cmd *cobra.Command, args []string) {
		num1, _ := strconv.Atoi(args[0])
		num2, _ := strconv.Atoi(args[1])
		result := num1 + num2
		fmt.Println(result)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
