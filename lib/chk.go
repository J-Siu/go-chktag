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

	"github.com/J-Siu/go-helper/v2/ezlog"
	"github.com/J-Siu/go-helper/v2/str"
	"github.com/charlievieth/strcase"
)

func ChkGitTag(workPath, tag string) (e error) {
	prefix := "GhkGitTag"
	var (
		vers *[]string
	)
	vers, e = GetGitTag(workPath)
	if e == nil {
		if str.ArrayContains(vers, &tag, false) {
			e = errors.New(prefix + ": " + tag + " already exist")
			e = errors.New(ezlog.Log().N(prefix).N(tag).M("already exist").String())

		}
	}
	return e
}

// Check if tag is the last tag in CHANGELOG.md
func ChkVerChangelog(workPath, tag string) (e error) {
	prefix := "ChkChangeLog"

	var (
		filePath string
		vers     *[]string
	)

	vers, filePath, e = GetVerChangeLog(workPath)
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

func ChkVerVersion(workPath, tag string) (e error) {
	prefix := "ChkVersion"

	var (
		filePath string
		ver      string
	)

	ver, filePath, e = GetVerVersion(workPath)
	if e == nil {
		if !strcase.EqualFold(ver, tag) {
			e = errors.New(ezlog.Log().N(prefix).N(filePath).N(tag).M("not found").String())
		}
	}

	return e
}
