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

	"github.com/J-Siu/go-helper/v2/basestruct"
	"github.com/J-Siu/go-helper/v2/ezlog"
	"github.com/J-Siu/go-helper/v2/str"
	"github.com/go-git/go-git/v6"
	"github.com/go-git/go-git/v6/plumbing"
	"github.com/go-git/go-git/v6/plumbing/storer"
	"golang.org/x/mod/semver"
)

// Get/Check tags of git tag
type Tag struct {
	*basestruct.Base
	WorkPath string
}

func (t *Tag) New(workPath string) *Tag {
	t.Base = new(basestruct.Base)
	t.MyType = "GitTag"
	t.WorkPath = workPath
	return t
}

func (t *Tag) Chk(tag string) (e error) {
	prefix := t.MyType + ".Chk"
	var (
		vers *[]string
	)
	vers, e = t.Get()
	if e == nil {
		if str.ArrayContains(vers, &tag, false) {
			e = errors.New(prefix + ": " + tag + " already exist")
			e = errors.New(ezlog.Log().N(prefix).N(tag).M("already exist").String())
		}
	}
	return e
}

// Check if tag is the last tag in git log/tag
func (t *Tag) Get() (vers *[]string, e error) {
	prefix := t.MyType + ".Get"

	var (
		_vers []string
		repo  *git.Repository
		tags  storer.ReferenceIter
	)

	repo, e = git.PlainOpen(t.WorkPath)
	if e == nil {
		tags, e = repo.Tags()
	} else {
		e = errors.New(t.WorkPath + ": " + e.Error())
	}
	if e == nil {
		e = tags.ForEach(func(tr *plumbing.Reference) error {
			_vers = append(_vers, tr.Name().Short())
			return nil
		})
		semver.Sort(_vers)
		ezlog.Debug().N(prefix).N("vers").Lm(_vers).Out()
	}

	return &_vers, e
}
