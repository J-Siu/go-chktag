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
	"strings"

	"github.com/J-Siu/go-gitcmd/v3/gitcmd"
	"github.com/J-Siu/go-helper/v2/ezlog"
)

// Get branch of git repo
// - branch name is stored in tags[0]
type GitBranch struct {
	ChkGet
}

func (t *GitBranch) New(workPath string) IChkGet {
	t.ChkGet.New(workPath)
	t.MyType = "GitBranch"
	if t.WorkPath == "." {
		t.filePath = "git branch"
	} else {
		t.filePath = workPath + "/(git branch)"
	}
	t.Get()
	t.Initialized = true
	return t
}

// Place holder only
func (t *GitBranch) Chk(tag string) IChkGet { return t }

// Get all git tag
func (t *GitBranch) Get() IChkGet {
	prefix := t.MyType + ".Get"
	var (
		gitCmd = new(gitcmd.GitCmd).New(t.WorkPath)
		branch = gitCmd.BranchCurrent().Run().Stdout.String()
	)
	if gitCmd.Err == nil {
		t.tags = []string{branch}
		ezlog.Debug().N(prefix).N("branch").Lm(branch).Out()
	} else {
		t.Base.Err = errors.New(t.WorkPath + ": " + strings.Trim(gitCmd.Err.Error(), "\n"))
	}
	return t
}
