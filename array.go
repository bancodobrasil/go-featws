// Copyright 2014 Unknwon
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package featws

import (
	"strconv"
	"strings"
)

// Array represents a config section.
type Array struct {
	f        *File
	Comment  string
	name     string
	sections []*Section

	isRawArray bool
	rawBody    string
}

func newArray(f *File, name string) *Array {
	return &Array{
		f:        f,
		name:     name,
		sections: make([]*Section, 0),
	}
}

// Name returns name of Array.
func (a *Array) Name() string {
	return a.name
}

// Body returns rawBody of Array if the section was marked as unparseable.
// It still follows the other rules of the INI format surrounding leading/trailing whitespace.
func (a *Array) Body() string {
	return strings.TrimSpace(a.rawBody)
}

// SetBody updates body content only if section is raw.
func (a *Array) SetBody(body string) {
	if !a.isRawArray {
		return
	}
	a.rawBody = body
}

func (a *Array) NewItem() (*Section, error) {
	section, err := a.f.NewSection(a.name + strconv.Itoa(len(a.sections)))
	if err != nil {
		return nil, err
	}
	a.sections = append(a.sections, section)
	return section, nil
}

func (a *Array) Length() int {
	return len(a.sections)
}

// ChildSections returns a list of child sections of current section.
// For example, "[parent.child1]" and "[parent.child12]" are child sections
// of section "[parent]".
func (a *Array) ChildSections() []*Section {
	prefix := a.name + a.f.options.ChildSectionDelimiter
	children := make([]*Section, 0, 3)
	for _, name := range a.f.sectionList {
		if strings.HasPrefix(name, prefix) {
			children = append(children, a.f.sections[name]...)
		}
	}
	return children
}

// ArrayStrings assumes named section exists and returns a zero-value when not.
func (f *File) ArrayStrings() []string {
	list := make([]string, 0)

	for k := range f.arrays {
		list = append(list, k)
	}

	return list
}
