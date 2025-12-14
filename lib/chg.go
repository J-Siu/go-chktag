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

package lib

import (
	"errors"
	"regexp"

	"github.com/J-Siu/go-helper/v2/basestruct"
	"github.com/J-Siu/go-helper/v2/errs"
	"github.com/J-Siu/go-helper/v2/ezlog"
	"github.com/J-Siu/go-helper/v2/file"
	"github.com/charlievieth/strcase"
)

// Get/Check versions in CHANGELOG.md
type Chg struct {
	*basestruct.Base
	WorkPath string
}

func (t *Chg) New(workPath string) *Chg {
	t.Base = new(basestruct.Base)
	t.MyType = "ChangeLog"
	t.WorkPath = workPath
	return t
}

// Check if tag is the last tag in CHANGELOG.md
func (t *Chg) Chk(tag string) (e error) {
	prefix := t.MyType + ".Chk"

	var (
		filePath string
		vers     *[]string
	)

	vers, filePath, e = t.Get()
	if e == nil {
		ezlog.Debug().N(prefix).
			Ln("vers==nil").M(vers == nil)
		if vers != nil {
			ezlog.Ln("len(*vers)").M(len(*vers))
			if len(*vers) > 0 {
				ezlog.Ln("!strcase.EqualFold((*vers)[len(*vers)-1], tag)").M(!strcase.EqualFold((*vers)[len(*vers)-1], tag))
			}
		}
		ezlog.Out()
		if vers == nil || !(len(*vers) > 0 && strcase.EqualFold((*vers)[len(*vers)-1], tag)) {
			e = errors.New(ezlog.Log().N(prefix).N(filePath).N(tag).M("not last tag").String())
		}
	}

	return e
}

// Return all versions from CHANGELOG.md
func (t *Chg) Get() (vers *[]string, filePath string, e error) {
	prefix := "GetVerChangeLog"

	var (
		_vers   []string
		content *[]string
		matches [][]string
		pattern string
		re      *regexp.Regexp
	)
	filePath = file.FindFile(t.WorkPath, FileChangLog, false)
	if filePath == "" {
		e = errors.New(FileVersion + " not found")
	}
	if e == nil {
		ezlog.Debug().N(prefix).N("file").M(filePath).Out()
		content, e = file.ReadStrArray(filePath)
	}
	if e == nil {
		// Get last Version = "- <ver>"
		pattern = `^- (.*)`
		re = regexp.MustCompile(pattern)
		for _, line := range *content {
			// Extract <ver>
			ezlog.Debug().N(prefix).N("line").M(line).Out()
			matches = re.FindAllStringSubmatch(line, -1)
			if matches != nil && len(matches[0][1]) > 0 {
				_vers = append(_vers, matches[0][1])
			}
		}
		ezlog.Debug().N(prefix).N("vers").Lm(_vers).Out()
	}

	errs.Queue(prefix, e)
	return &_vers, filePath, e
}
