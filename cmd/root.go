/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bfish/src"
	"fmt"
	"os"
	"strconv"
	"strings"

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
	getCmd = &cobra.Command{
		Use:   "get [strings to commands]",
		Short: "prints Sboxes, Pkeys, Nrounds, Key",
		Args:  cobra.MaximumNArgs(1),
	}

	getSboxCmd = &cobra.Command{
		Use:   "sbox [sbox number] [sbox column]",
		Short: "prints Sbox",
		Args:  cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			i, _ := strconv.Atoi(args[0])
			j, _ := strconv.Atoi(args[1])
			fmt.Printf("%#x\n", src.ORIG_S[i][j])
		},
	}
	getPkeyCmd = &cobra.Command{
		Use:   "pkey [pkey number]",
		Short: "prints pkey",
		Args:  cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			i, _ := strconv.Atoi(args[0])
			fmt.Printf("%#x\n", src.ORIG_P[i])
		},
	}
	getAllCmd = &cobra.Command{
		Use:   "all [string]",
		Short: "prints all elements of a specific parameter",
		Long: `outputs all elements of a specific parameter
example: get all pkeys; get all sboxes`,
		Args: cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			switch strings.Join(args, " ") {
			case "s", "sboxes":
				for i := 0; i < 4; i++ {
					for j := 0; j < 256; j++ {
						fmt.Printf("%#x\n", src.ORIG_S[i][j])
					}
				}
			case "p", "pkeys":
				for i := 0; i < 18; i++ {
					fmt.Printf("%#x\n", src.ORIG_P[i])
				}
			}
		},
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
	rootCmd.AddCommand(getCmd)

	// getCmd commands
	getCmd.AddCommand(getSboxCmd)
	getCmd.AddCommand(getPkeyCmd)
	getCmd.AddCommand(getAllCmd)
}
