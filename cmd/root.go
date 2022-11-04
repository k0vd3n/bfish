/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/common-nighthawk/go-figure"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var (
	rootCmd = &cobra.Command{
		Use:   "bfish",
		Short: "A blowfish application",
		Long: `bfish is an application for symmetric encryption and decryption of 
	text, hex, decimal data. Blowfish is a cryptographic algorithm that implements 
	block symmetric encryption with a variable key length. 
	
	Designed by Bruce Schneierin 1993. It is a Feistel network. Performed on simple 
	and fast operations: XOR, substitution, addition. It is non-proprietary and 
	freely distributed. You can readmore about how the algorithm works on the Internet`,
		Run: func(cmd *cobra.Command, args []string) {
			myFigure := figure.NewColorFigure("k0vd3n_app", "", "yellow", true)
			myFigure.Print()
			fmt.Println("  use flag -h or --help for more information")
		},
		// Uncomment the following line if your bare application
		// has an action associated with it:
		// Run: func(cmd *cobra.Command, args []string) { },
	}
	getIdCmd = &cobra.Command{
		Use:   "get [strings to commands]",
		Short: "prints Sboxes, Pkeys, Nrounds, Key",
		Args:  cobra.MaximumNArgs(1),
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.bfish.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.AddCommand(getIdCmd)
}
