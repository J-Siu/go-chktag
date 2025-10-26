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

	"github.com/J-Siu/go-helper/v2/errs"
	"github.com/J-Siu/go-helper/v2/ezlog"
	"github.com/J-Siu/go-helper/v2/file"
	"github.com/J-Siu/go-helper/v2/str"
	"github.com/charlievieth/strcase"
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
		// tagObjs *object.TagIter
		tags []string
	)
	r, e = git.PlainOpen(workPath)
	if e == nil {
		tagRefs, e = r.Tags()
		// tagObjs, e = r.TagObjects()
	}
	if e == nil {
		// e = tagObjs.ForEach(func(t *object.Tag) error {
		// 	ezlog.Log().N("t").M(t.Name).Out()
		// 	// fmt.Println(t)
		// 	tags = append(tags, t.Name)
		// 	return nil
		// })

		e = tagRefs.ForEach(func(t *plumbing.Reference) error {
			ezlog.Debug().N(prefix).Nn("t").M(t.Strings()).Out()
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

// Check if tag is the last tag in CHANGELOG.md
func ChkVerChangelog(workPath, tag string) (e error) {
	prefix := "ChkChangeLog"

	vers := GetVerChangeLog(workPath)
	if vers == nil || !strcase.EqualFold((*vers)[len(*vers)-1], tag) {
		e = errors.New(tag + " not found or not last tag in " + FileChangLog)
	}

	errs.Queue(prefix, e)
	return e
}

func ChkVerVersion(workPath, tag string) (e error) {
	prefix := "ChkVersion"

	ver := GetVerVersion(workPath)
	if !strcase.EqualFold(ver, tag) {
		e = errors.New(tag + " not found")
	}

	errs.Queue(prefix, e)
	return e
}

// Check if tag is the last tag in git log/tag
func GetGitTag(workPath string) (ver string) { return ver }

// Return all versions from CHANGELOG.md
func GetVerChangeLog(workPath string) *[]string {
	prefix := "GetVerChangeLog"

	var (
		content  *[]string
		e        error
		filePath string
		matches  [][]string
		pattern  string
		re       *regexp.Regexp
		vers     []string
	)
	filePath = file.FindFile(workPath, FileChangLog, false)
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
				vers = append(vers, matches[0][1])
			}
		}
		ezlog.Debug().N(prefix).Nn("vers").M(vers).Out()
	}

	errs.Queue(prefix, e)
	return &vers
}

// Return version from version.go
func GetVerVersion(workPath string) (ver string) {
	prefix := "GetVerVersion"
	var (
		content  *[]string
		e        error
		filePath string
		matches  [][]string
		pattern  string
		re       *regexp.Regexp
	)
	// check version.go
	filePath = file.FindFile(workPath, FileVersion, false)
	if filePath == "" {
		e = errors.New(FileVersion + " not found")
	}
	if e == nil {
		ezlog.Debug().N(prefix).N("file").M(filePath).Out()
		content, e = file.ReadStrArray(filePath)
	}
	if e == nil {
		// Get line: Version = "<ver>"
		pattern = `\s*Version\s*=\s*\"(.*)\"`
		re = regexp.MustCompile(pattern)
		for _, line := range *content {
			matches = re.FindAllStringSubmatch(line, -1)
			ezlog.Debug().N(prefix).N("line").M(line).Out()
			if matches != nil && len(matches[0][1]) != 0 {
				// Extract <ver>
				ver = matches[0][1]
				ezlog.Debug().N(prefix).N("ver").M(ver).Out()
				break
			}
		}
	}

	errs.Queue(prefix, e)
	return ver
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
