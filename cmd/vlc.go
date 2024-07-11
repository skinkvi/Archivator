package cmd

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var vlcCmd = &cobra.Command{
	Use:   "vlc",
	Short: "using varible-lenght code",
	Run:   pack,
}

var ErrEmptyPath = errors.New("Вы не указали путь до файла")

const packageExtengen = "vlc"

func pack(_ *cobra.Command, args []string) {
	if len(args) == 0 {
		handErr(ErrEmptyPath)
	}
	filePath := args[0]

	read, err := os.Open(filePath)
	if err != nil {
		handErr(err)
	}

	defer read.Close()

	var data []byte
	for {
		buf := make([]byte, 32)
		n, err := read.Read(buf)
		if n > 0 {
			data = append(data, buf[:n]...)
		}
		if err != nil {
			if err != io.EOF {
				panic(err)
			}
			break
		}
	}

	fmt.Println(string(data))

	err = ioutil.WriteFile(packedFileName(filePath), data, 0644)
	if err != nil {
		handErr(err)
	}

}

func packedFileName(path string) string {
	fileName := filepath.Base(path)
	return strings.TrimSuffix(fileName, filepath.Ext(fileName)) + "." + packageExtengen
}

func init() {
	packCmd.AddCommand(vlcCmd)
}
