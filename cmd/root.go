/*
Copyright © 2020 Nick Maliwacki <knic.knic@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var Path string
var Address string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "domain-socket-proxy",
	Short: "Allows you to proxy one tcp port to another via a unix domain socket",
	Long: `Allows you to proxy one tcp port to another via a unix domain socket
  
  Lets say you have an existing app on port 80 and you want to expose it on port 9999

  launch the forwarder
    .\domain-socket-proxy.exe forward --address :80 --path .\unix.socket

  launch the server  
    .\domain-socket-proxy.exe serve --address localhost:9999 --path .\unix.socket
  
  The above example uses the file path .\unix.socket for the unix domain socket
  `,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {

	rootCmd.AddCommand(forwardCmd)
	rootCmd.PersistentFlags().StringVarP(&Path, "path", "p", "", "The path for the unix domain socket")
	rootCmd.MarkFlagRequired("path")
	rootCmd.PersistentFlags().StringVarP(&Address, "address", "a", "", "The tcp address ex: localhost:80 or :80")
	rootCmd.MarkFlagRequired("address")

}
