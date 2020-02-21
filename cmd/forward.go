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
	"net"
	"os"

	"github.com/spf13/cobra"
)

// forwardCmd represents the forward command
var forwardCmd = &cobra.Command{
	Use:   "forward",
	Short: "Forwards a unix domain socket to a tcp port",
	Long: `Forwards a unix domain socket to a tcp port

examples:
    forward --path .\socket --address :8080`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("forward called")

		err := cleanupSocket(Path)
		if err != nil {
			panic(err)
		}

		domain, err := net.ResolveUnixAddr("unix", Path)
		if err != nil {
			panic(err)
		}
		ln, err := net.ListenUnix("unix", domain)
		if err != nil {
			panic(err)
		}
		for {
			conn, err := ln.Accept()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
			go fwdConnection(conn)
		}
	},
}

func cleanupSocket(path string) error {
	pipe, err := os.Stat(path)
	if err == nil {
		mode := pipe.Mode()
		isSocket := os.ModeSocket&mode != 0

		// hard coding socket to true due to bug
		// https://github.com/golang/go/issues/33357
		isSocket = true
		if isSocket {
			err = os.Remove(path)
			if err != nil {
				return err
			}
		} else {
			return fmt.Errorf("Path [%v] exists but is not a socket", path)
		}
	} else if !os.IsNotExist(err) {
		return err
	}
	return nil
}
func fwdConnection(conn net.Conn) {
	defer conn.Close()

	remote, err := net.Dial("tcp", Address)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	defer remote.Close()
	proxy(remote, conn)
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// forwardCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// forwardCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
