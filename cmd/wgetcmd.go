package cmd

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

var output string

var wgetcmd = &cobra.Command{
	Use:     "wget",
	Example: "{tool} wget iqsing.github.io/download.tar.gz -o /tmp/download.tar.gz",
	Args:    cobra.ExactArgs(1),
	Short:   "wget download cli",
	Long:    "use wget can download the resoure from internet",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("---------wget downloading------------")
		download(args[0], output)
	},
}

func init() {
	rootCmd.AddCommand(wgetcmd)
	//define flag and config setting
	wgetcmd.Flags().StringVarP(&output, "output", "o", "", "output file")
	wgetcmd.MarkFlagRequired("output")
}

func download(url, path string) {
	out, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
		return
	}

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer res.Body.Close()

	_, err = io.Copy(out, res.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("save the dowload file as " + path)
}
