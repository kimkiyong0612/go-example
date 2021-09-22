/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

const imgDir string = "img"

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "This command will get the desired Gopher",
	Long:  `This get command will call Github repository in order to return the desired Gopher.`,
	Run: func(cmd *cobra.Command, args []string) {
		gopherName := "xxxx.png"

		if len(args) >= 1 && args[0] != "" {
			gopherName = args[0]
		}

		URL := fmt.Sprintf("https://github.com/scraly/gophers/raw/main/%v.png", gopherName)
		fmt.Println("get called")
		fmt.Println("Try to get '" + gopherName + "' Gopher...")

		// Get the data
		resp, err := http.Get(URL)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		switch resp.StatusCode {
		case 200:
			// create the dir and file
			if _, err := os.Stat(imgDir); os.IsNotExist(err) {
				os.Mkdir(imgDir, 0777)
			}

			out, err := os.Create(filepath.Join(imgDir, gopherName+".png"))
			if err != nil {
				log.Fatal(err)
			}
			defer out.Close()

			// Write the body to file
			_, err = io.Copy(out, resp.Body)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("Perfect! Just save in " + out.Name() + "!")
		default:
			log.Fatal("Error: " + gopherName + " not exists! :-(")
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
