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
	"errors"
	"regexp"

	"github.com/J-Siu/go-chktag/global"
	"github.com/J-Siu/go-helper/v2/ezlog"
	"github.com/J-Siu/go-helper/v2/file"
	"github.com/charlievieth/strcase"
)

// Get/Check version in version.go
type VerFile struct {
	ChkGet
}

func (t *VerFile) New(workPath string) IChkGet {
	t.ChkGet.New(workPath)
	t.MyType = "Ver"
	t.Get()
	t.Initialized = true
	return t
}

func (t *VerFile) Chk(tag string) IChkGet {
	if t.Base.Err == nil {
		if !strcase.EqualFold(t.tags[0], tag) {
			t.Base.Err = errors.New(ezlog.Log().N(t.filePath).N(tag).M("not found").String())
		}
	}
	return t
}

// Return version from version.go
func (t *VerFile) Get() IChkGet {
	prefix := t.MyType + ".Get"
	var (
		content *[]string
		matches [][]string
		pattern string
		re      *regexp.Regexp
	)
	// check version.go
	t.filePath = file.FindFile(t.WorkPath, global.FileVersion, false)
	if t.filePath == "" {
		t.Base.Err = errors.New(t.WorkPath + ": " + global.FileVersion + " not found")
	}
	if t.Base.Err == nil {
		ezlog.Debug().N(prefix).N("file").M(t.filePath).Out()
		content, t.Base.Err = file.ReadStrArray(t.filePath)
	}
	if t.Base.Err == nil {
		// Get line: Version = "<ver>"
		pattern = `\s*Version\s*(string)?\s*=\s*\"(.*)\"`
		re = regexp.MustCompile(pattern)
		for _, line := range *content {
			matches = re.FindAllStringSubmatch(line, -1)
			ezlog.Debug().N(prefix).N("line").M(line).Out()
			ezlog.Debug().N(prefix).N("matches").M(matches).Out()
			// Extract <ver>
			if matches != nil && len(matches[0][2]) != 0 {
				t.tags = append(t.tags, matches[0][2])
				ezlog.Debug().N(prefix).N("ver").M(t.tags[0]).Out()
				break
			}
		}
	}

	return t
}
