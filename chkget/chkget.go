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

import "github.com/J-Siu/go-helper/v2/basestruct"

type ChkGet struct {
	basestruct.Base
	WorkPath string
	filePath string
	tags     []string
}

func (t *ChkGet) New(workPath string) IChkGet {
	t.WorkPath = workPath
	t.Base.Err = nil
	t.tags = nil
	return t
}

func (t *ChkGet) Chk(tag string) IChkGet       { return t }
func (t *ChkGet) Get() IChkGet                 { return t }
func (t *ChkGet) Err() error                   { return t.Base.Err }
func (t *ChkGet) FilePath() (filePath *string) { return &t.filePath }
func (t *ChkGet) Tags() (tags *[]string)       { return &t.tags }
