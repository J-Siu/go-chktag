/*
Copyright © 2025 John, Sing Dao, Siu <john.sd.siu@gmail.com>

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
	"os"

	"github.com/J-Siu/go-chktag/global"
	"github.com/J-Siu/go-chktag/lib"
	"github.com/J-Siu/go-helper/v2/errs"
	"github.com/J-Siu/go-helper/v2/ezlog"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-chktag",
	Short: `Check tag information`,
	Long: `Show ` + lib.FileChangLog + `, ` + lib.FileVersion + ` and git tag.
Use -t to specify tag version.`,
	Version: global.Version,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if global.Flag.Debug {
			ezlog.SetLogLevel(ezlog.DEBUG)
		}
		ezlog.Debug().N("Version").Mn(global.Version).Nn("Flag").M(&global.Flag).Out()
	},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			args = []string{"."}
		}
		for _, path := range args {
			if len(args) > 1 {
				ezlog.Log().N(path).Out()
			}
			if global.Flag.Tag == "" {
				getTag(path)
			} else {
				chkTag(path)
				if errs.IsEmpty() {
					ezlog.Log().M("Passed").Out()
				} else {
					ezlog.Err().L().M(errs.Errs).Out()
					errs.Clear()
				}
			}
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cmd := rootCmd
	cmd.PersistentFlags().BoolVarP(&global.Flag.Debug, "debug", "d", false, "Enable debug")
	cmd.PersistentFlags().StringVarP(&global.Flag.Tag, "tag", "t", "", "check specific tag")
}

func chkTag(path string) {
	lib.ChkGitTag(path, global.Flag.Tag)
	lib.ChkVerVersion(path, global.Flag.Tag)
	lib.ChkVerChangelog(path, global.Flag.Tag)
}

func getTag(path string) {
	var (
		filePath string
		tag      string
		tags     *[]string
	)

	tag, filePath = lib.GetVerVersion(path)
	if filePath == "" {
		ezlog.Err().M(lib.FileVersion + " not found").Out()
	} else {
		ezlog.Log().N(filePath).M(tag).Out()
	}

	tags, filePath = lib.GetVerChangeLog(path)
	if filePath == "" {
		ezlog.Err().M(lib.FileChangLog + " not found").Out()
	} else if tags != nil && len(*tags) > 0 {
		ezlog.Log().Nn(filePath).M(tags).Out()
	} else {
		ezlog.Log().N("ChangeLog").Out()
	}

	tags = lib.GetGitTag(path)
	if tags != nil && len(*tags) > 0 {
		ezlog.Log().Nn("Git Tag").M(tags).Out()
	} else {
		ezlog.Log().N("Git Tag").Out()
	}
}
