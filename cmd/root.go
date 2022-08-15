/*
Copyright © 2022 JhndaCoder
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

var dataFile string

var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "Todo is a todo application",
	Long: `Todo will help you get more done in less time.

It's designed to be as simple as possible to help you accomplish your goals.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {

	home, err := homedir.Dir()
	if err != nil {
		log.Println("Unable to detect home directory. Please set data file using --datafile.")
	}

	rootCmd.PersistentFlags().StringVar(&dataFile, "datafile", home+string(os.PathSeparator)+".tridos.json", "data file to store todos")

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
