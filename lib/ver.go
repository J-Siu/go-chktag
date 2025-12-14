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
	"github.com/J-Siu/go-helper/v2/ezlog"
	"github.com/J-Siu/go-helper/v2/file"
	"github.com/charlievieth/strcase"
)

// Get/Check version in version.go
type Ver struct {
	*basestruct.Base
	WorkPath string
}

func (t *Ver) New(workPath string) *Ver {
	t.Base = new(basestruct.Base)
	t.MyType = "Ver"
	t.WorkPath = workPath
	return t
}

func (t *Ver) Chk(tag string) (e error) {
	prefix := "ChkVersion"

	var (
		filePath string
		ver      string
	)

	ver, filePath, e = t.Get()
	if e == nil {
		if !strcase.EqualFold(ver, tag) {
			e = errors.New(ezlog.Log().N(prefix).N(filePath).N(tag).M("not found").String())
		}
	}

	return e
}

// Return version from version.go
func (t *Ver) Get() (ver, filePath string, e error) {
	prefix := "GetVerVersion"
	var (
		content *[]string
		matches [][]string
		pattern string
		re      *regexp.Regexp
	)
	// check version.go
	filePath = file.FindFile(t.WorkPath, FileVersion, false)
	if filePath == "" {
		e = errors.New(FileVersion + " not found")
	}
	if e == nil {
		ezlog.Debug().N(prefix).N("file").M(filePath).Out()
		content, e = file.ReadStrArray(filePath)
	}
	if e == nil {
		// Get line: Version = "<ver>"
		pattern = `\s*Version\s*(string)?\s*=\s*\"(.*)\"`
		re = regexp.MustCompile(pattern)
		for _, line := range *content {
			matches = re.FindAllStringSubmatch(line, -1)
			ezlog.Debug().N(prefix).N("line").M(line).Out()
			ezlog.Debug().N(prefix).N("matches").M(matches).Out()
			if matches != nil && len(matches[0][2]) != 0 {
				// Extract <ver>
				ver = matches[0][2]
				ezlog.Debug().N(prefix).N("ver").M(ver).Out()
				break
			}
		}
	}

	return ver, filePath, e
}
