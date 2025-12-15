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

package chkget

import (
	"github.com/J-Siu/go-chktag/global"
	"github.com/J-Siu/go-helper/v2/errs"
	"github.com/J-Siu/go-helper/v2/ezlog"
)

type IChkGet interface {
	Chk(tag string) IChkGet
	Err() error
	FilePath() *string
	New(workPath string) IChkGet
	Tags() *[]string
}

func ChkTag(chkget IChkGet) {
	errs.Queue("", chkget.Chk(global.Flag.Tag).Err())
}

func GetTag(chkget IChkGet) {
	errs.Queue("", chkget.Err())
	if chkget.Err() == nil {
		ezlog.Log().N(chkget.FilePath())
		tags := *chkget.Tags()
		if len(tags) > 0 {
			if global.Flag.Verbose {
				ezlog.Lm(tags)
			} else {
				ezlog.M(tags[len(tags)-1])
			}
		}
		ezlog.Out()
	}
}
