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
	"github.com/go-git/go-git/v6"
	"github.com/go-git/go-git/v6/plumbing"
	"github.com/go-git/go-git/v6/plumbing/storer"
	"golang.org/x/mod/semver"
)

// Check if tag is the last tag in git log/tag
func GetGitTag(workPath string) (*[]string, error) {
	prefix := "GetGitTag"

	var (
		e    error
		repo *git.Repository
		tags storer.ReferenceIter
		vers []string
	)

	repo, e = git.PlainOpen(workPath)
	if e == nil {
		tags, e = repo.Tags()
	} else {
		e = errors.New(workPath + ": " + e.Error())
	}
	if e == nil {
		e = tags.ForEach(func(t *plumbing.Reference) error {
			vers = append(vers, t.Name().Short())
			return nil
		})
		semver.Sort(vers)
		ezlog.Debug().N(prefix).N("vers").Lm(vers).Out()
	}

	return &vers, e
}

// Return all versions from CHANGELOG.md
func GetVerChangeLog(workPath string) (*[]string, string, error) {
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
		ezlog.Debug().N(prefix).N("vers").Lm(vers).Out()
	}

	errs.Queue(prefix, e)
	return &vers, filePath, e
}

// Return version from version.go
func GetVerVersion(workPath string) (ver, filePath string, e error) {
	prefix := "GetVerVersion"
	var (
		content *[]string
		matches [][]string
		pattern string
		re      *regexp.Regexp
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
