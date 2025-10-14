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

	"github.com/J-Siu/go-helper/v2/errs"
	"github.com/J-Siu/go-helper/v2/file"
	"github.com/J-Siu/go-helper/v2/str"
	"github.com/go-git/go-git/v6"
	"github.com/go-git/go-git/v6/plumbing"
	"github.com/go-git/go-git/v6/plumbing/storer"
)

const (
	FileVersion  = "version.go"
	FileChangLog = "changelog.md"
)

func ChkGitTag(workPath, tag string) (e error) {
	prefix := "GhkGitTag"
	var (
		r       *git.Repository
		tagRefs storer.ReferenceIter
		tags    []string
	)
	r, e = git.PlainOpen(workPath)
	if e == nil {
		tagRefs, e = r.Tags()
	}
	if e == nil {
		e = tagRefs.ForEach(func(t *plumbing.Reference) error {
			tags = append(tags, t.Name().Short())
			return nil
		})
	}
	if str.ArrayContains(&tags, &tag, false) {
		e = errors.New(tag + " already exist")
	}
	errs.Queue(prefix, e)
	return e
}

func ChkChangelog(workPath, tag string) (e error) {
	prefix := "ChkChangeLog"
	var (
		filePath string
	)
	// check version.go
	filePath = file.FindFile(workPath, FileChangLog, false)
	if filePath == "" {
		e = errors.New(FileChangLog + " not found")
	}
	if e == nil && !findInFile(filePath, tag) {
		e = errors.New(tag + " not found in " + filePath)
	}
	errs.Queue(prefix, e)
	return e
}

func ChkVersion(workPath, tag string) (e error) {
	prefix := "ChkVersion"
	var (
		filePath string
	)
	// check version.go
	filePath = file.FindFile(workPath, FileVersion, false)
	if filePath == "" {
		e = errors.New(FileVersion + " not found")
	}
	if e == nil && !findInFile(filePath, tag) {
		e = errors.New(tag + " not found in " + filePath)
	}
	errs.Queue(prefix, e)
	return e
}

func findInFile(filePath, strIn string) (b bool) {
	var (
		e        error
		sP       *string
		strArrIn = []string{strIn}
	)
	sP, e = file.ReadStr(filePath)
	if e == nil {
		if str.ContainsAnySubStringsBool(sP, &strArrIn, true) {
			return true
		}
	}
	return b
}
