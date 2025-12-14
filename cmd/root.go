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
		ezlog.Debug().N("Version").M(global.Version).Ln("Flag").M(&global.Flag).Out()
	},
	Run: func(cmd *cobra.Command, args []string) {
		var (
			argc = len(args)
		)
		if argc == 0 {
			args = []string{"."}
		}
		for _, path := range args {
			if argc > 1 {
				ezlog.Log().N(path).Out()
			}
			if global.Flag.Tag == "" {
				getTag(path)
			} else {
				chkTag(path)
				if errs.IsEmpty() {
					ezlog.Log().M("Passed").Out()
				}
			}
		}
		if errs.NotEmpty() {
			ezlog.Err().L().M(errs.Errs).Out()
			errs.Clear()
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
	cmd.PersistentFlags().BoolVarP(&global.Flag.Verbose, "verbose", "v", false, "Enable verbose")
	cmd.PersistentFlags().StringVarP(&global.Flag.Tag, "tag", "t", "", "check specific tag")
}

func chkTag(path string) {
	errs.Queue("", new(lib.Ver).New(path).Chk(global.Flag.Tag))
	errs.Queue("", new(lib.Chg).New(path).Chk(global.Flag.Tag))
	errs.Queue("", new(lib.Tag).New(path).Chk(global.Flag.Tag))
}

func getTag(path string) {
	var (
		e        error
		filePath string
		ver      string
		vers     *[]string
	)

	ver, filePath, e = new(lib.Ver).New(path).Get()
	errs.Queue("", e)
	if e == nil {
		ezlog.Log().N(filePath).M(ver).Out()
	}

	vers, filePath, e = new(lib.Chg).New(path).Get()
	errs.Queue("", e)
	if e == nil {
		ezlog.Log().N(filePath)
		if vers != nil && len(*vers) > 0 {
			if global.Flag.Verbose {
				ezlog.Lm(vers)
			} else {
				ezlog.M((*vers)[len(*vers)-1])
			}
		}
		ezlog.Out()
	}

	vers, e = new(lib.Tag).New(path).Get()
	errs.Queue("", e)
	if e == nil {
		ezlog.Log().N("Git Tag")
		if vers != nil && len(*vers) > 0 {
			if global.Flag.Verbose {
				ezlog.Lm(vers)
			} else {
				ezlog.M((*vers)[len(*vers)-1])
			}
		}
		ezlog.Out()
	}

}
