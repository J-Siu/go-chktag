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

	"github.com/J-Siu/go-helper/v2/ezlog"
	"github.com/J-Siu/go-helper/v2/str"
	"github.com/go-git/go-git/v6"
	"github.com/go-git/go-git/v6/plumbing"
	"github.com/go-git/go-git/v6/plumbing/storer"
	"golang.org/x/mod/semver"
)

// Get/Check tags of git tag
type GitTag struct {
	ChkGet
}

func (t *GitTag) New(workPath string) IChkGet {
	t.ChkGet.New(workPath)
	t.MyType = "GitTag"
	if t.WorkPath == "." {
		t.filePath = "git tag"
	} else {
		t.filePath = workPath + "/(git tag)"
	}
	t.Get()
	t.Initialized = true
	return t
}

// Check if git tag already exist
func (t *GitTag) Chk(tag string) IChkGet {
	if t.Base.Err == nil {
		if str.ArrayContains(&t.tags, &tag, false) {
			t.Base.Err = errors.New(ezlog.Log().N(t.filePath).N(tag).M("already exist").String())
		}
	}
	return t
}

// Get all git tag
func (t *GitTag) Get() IChkGet {
	prefix := t.MyType + ".get"
	var (
		repo *git.Repository
		tags storer.ReferenceIter
	)

	repo, t.Base.Err = git.PlainOpen(t.WorkPath)
	if t.Base.Err == nil {
		tags, t.Base.Err = repo.Tags()
	} else {
		t.Base.Err = errors.New(t.WorkPath + ": " + t.Base.Err.Error())
	}
	if t.Base.Err == nil {
		t.Base.Err = tags.ForEach(func(tr *plumbing.Reference) error {
			t.tags = append(t.tags, tr.Name().Short())
			return nil
		})
		semver.Sort(t.tags)
		ezlog.Debug().N(prefix).N("vers").Lm(t.tags).Out()
	}

	return t
}
