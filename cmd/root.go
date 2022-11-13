/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bfish/blowfish"
	"bfish/src"
	"encoding/binary"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/common-nighthawk/go-figure"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var (
	mode    string
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
	encryptCmd = &cobra.Command{
		Use:   "encrypt [string] [string]",
		Short: "encrypts 64 bits",
		Long: `accepts 2 strings in hex, each of which is half of a 64-bit fragment.
encrypts data in 3 different modes: decimal, hexadecimal and as a string. 
The output gives two types of encrypted data in 3 different types: decimal, 
hexadecimal and as a string`,
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			b := make([]byte, 4)
			result := make([]byte, 8)
			var xl uint32
			var xr uint32
			// if binary.BigEndian.Uint32([]byte(args[0])) > 0xffffffff || binary.BigEndian.Uint32([]byte(args[1])) > 0xffffffff {
			// 	fmt.Println("arguments can't contain more than 64 bits")
			// 	os.Exit(1)
			// }
			switch mode {
			case "string":
				// _, err := strconv.Atoi(args[0])
				// if err != nil {
				// 	os.Exit(1)
				// }
				// _, err = strconv.Atoi(args[1])
				// if err != nil {
				// 	os.Exit(1)
				// }
				xl = binary.BigEndian.Uint32([]byte(args[0]))
				xr = binary.BigEndian.Uint32([]byte(args[1]))
				binary.BigEndian.PutUint32(b, xl)
				fmt.Println(strings.Join(args, " "))
				fmt.Println(b)
				binary.BigEndian.PutUint32(b, xr)
				fmt.Println(b)
			case "hex":
				l, _ := strconv.ParseUint(args[0], 16, 32)
				r, _ := strconv.ParseUint(args[1], 16, 32)
				xl = uint32(l)
				xr = uint32(r)
				fmt.Printf("%#x %#x \n", xl, xr)
			case "decimal":
				l, _ := strconv.ParseUint(args[0], 10, 32)
				r, _ := strconv.ParseUint(args[1], 10, 32)
				xl = uint32(l)
				xr = uint32(r)
				fmt.Printf("%d %d \n", xl, xr)
			default:
				os.Exit(1)
			}
			var bf = &blowfish.Blowfish{}
			bf = blowfish.New(blowfish.Key)
			bf.Encrypt(&xl, &xr)
			binary.BigEndian.PutUint32(b, xl)
			fmt.Println("ciphertex xl in []byte: ", b)
			result = append(result, b...)
			binary.BigEndian.PutUint32(b, xr)
			fmt.Println("ciphertex xj in []byte: ", b)

			fmt.Println("ciphertex xl in decimal: ", xl)
			fmt.Println("ciphertex xr in decimal: ", xr)
			fmt.Printf("ciphertex xl in hex: %#x\n", xl)
			fmt.Printf("ciphertex xr in hex: %#x\n", xr)

			result = append(result, b...)
			fmt.Println(string(result))
		},
	}

	fullencryptCmd = &cobra.Command{
		Use:   "full [string]",
		Short: "encrypts message",
		Args:  cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("source string = ", strings.Join(args, " "))
			fmt.Println("source string in bytes = ", []byte(strings.Join(args, " ")))
			var bf = *blowfish.New(blowfish.Key)
			bytearr := []byte(strings.Join(args, " "))
			var bytestr []byte = blowfish.EncryptLoop(bytearr /*strings.Join(args, " ")*/, bf)
			cipherstr := string(bytestr)
			fmt.Printf("ciphertext in hex = ")
			for i := 0; i < len(bytestr); i++ {
				fmt.Printf("%x ", bytestr[i])
			}
			fmt.Printf("\n")
			fmt.Printf("ciphertext in decimal = ")
			for i := 0; i < len(bytestr); i++ {
				fmt.Printf("%d ", bytestr[i])
			}
			fmt.Printf("\n")
			fmt.Println(strings.Join(args, " "))
			fmt.Println([]byte(strings.Join(args, " ")))
			fmt.Println("ciphertext in bytes = ", bytestr)
			fmt.Println("ciphertext in string = ", cipherstr)
		},
	}

	decryptCmd = &cobra.Command{
		Use:   "decrypt",
		Short: "decrypt 64 bits",
		Long: `decrypts 2 variables in 3 different modes:
decimal, hexadecimal and as a string`,
		Args: cobra.MaximumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			b := make([]byte, 4)
			result := make([]byte, 8)
			var xl uint32
			var xr uint32
			if binary.BigEndian.Uint32([]byte(args[0])) > 0xffffffff || binary.BigEndian.Uint32([]byte(args[1])) > 0xffffffff {
				fmt.Println("arguments can't contain more than 64 bits")
				os.Exit(1)
			}
			switch mode {
			case "string":
				xl = binary.BigEndian.Uint32([]byte(args[0]))
				xr = binary.BigEndian.Uint32([]byte(args[1]))
				binary.BigEndian.PutUint32(b, xl)
				fmt.Println(strings.Join(args, " "))
				fmt.Println(b)
				binary.BigEndian.PutUint32(b, xr)
				fmt.Println(b)
			case "hex":
				l, _ := strconv.ParseUint(args[0], 16, 32)
				r, _ := strconv.ParseUint(args[1], 16, 32)
				xl = uint32(l)
				xr = uint32(r)
				fmt.Printf("%#x %#x \n", xl, xr)
			case "decimal":
				l, _ := strconv.ParseUint(args[0], 10, 32)
				r, _ := strconv.ParseUint(args[1], 10, 32)
				xl = uint32(l)
				xr = uint32(r)
				fmt.Printf("%d %d \n", xl, xr)
			default:
				os.Exit(1)
			}
			var bf = &blowfish.Blowfish{}
			bf = blowfish.New(blowfish.Key)
			bf.Decrypt(&xl, &xr)
			binary.BigEndian.PutUint32(b, xl)
			fmt.Println("source text xl in []byte: ", b)
			result = append(result, b...)
			binary.BigEndian.PutUint32(b, xr)
			fmt.Println("source text xj in []byte: ", b)

			fmt.Println("source text xl in decimal: ", xl)
			fmt.Println("source text xr in decimal: ", xr)
			fmt.Printf("source text xl in hex: %#x\n", xl)
			fmt.Printf("source text xr in hex: %#x\n", xr)

			result = append(result, b...)
			fmt.Println(string(result))
		},
	}

	fulldecryptCmd = &cobra.Command{
		Use:   "full [string]",
		Short: "encrypts message",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			var intString []int
			for i := 0; i < len(args); i++ {
				x, _ := strconv.ParseInt(args[i], 10, 10)
				intString = append(intString, int(x))
			}
			newarr := make([]byte, 2)
			var sourcearr []byte
			for i := 0; i < len(intString); i++ {
				newarr[0] = 0
				binary.BigEndian.PutUint16(newarr, uint16(intString[i]))
				sourcearr = append(sourcearr, newarr[1])
			}

			fmt.Println("ciphertext in bytes = ", strings.Join(args, " "))
			var bf = *blowfish.New(blowfish.Key)
			bytestr := blowfish.DecryptLoop(sourcearr, bf)
			decryptedstr := string(bytestr)
			fmt.Printf("\n")
			fmt.Println("decrypted message in bytes = ", bytestr)
			fmt.Println("decrypted message in string = ", decryptedstr)
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
	rootCmd.AddCommand(encryptCmd)
	rootCmd.AddCommand(decryptCmd)

	// getCmd commands
	getCmd.AddCommand(getSboxCmd)
	getCmd.AddCommand(getPkeyCmd)
	getCmd.AddCommand(getAllCmd)

	// encryptCmd commands
	encryptCmd.AddCommand(fullencryptCmd)

	// decryptCmd commands
	decryptCmd.AddCommand(fulldecryptCmd)

	// flags
	encryptCmd.Flags().StringVarP(&mode, "mode", "m", "", "mode of input (string, hex, decimal)")
	decryptCmd.Flags().StringVarP(&mode, "mode", "m", "", "mode of input (string, hex, decimal)")

}
