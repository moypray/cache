// Copyright 2010 Petar Maymounkov. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cache

import (
	"time"
)

type Int int

func (x Int) Less(than Item) bool {
	return x < than.(Int)
}

type String string

func (x String) Less(than Item) bool {
	return x < than.(String)
}

type Person struct {
	ID    int
	Name  string
	Birth time.Time
}

// Birth is the critical factor for the rank list.
// If Birth was the same, then Name will be considered.
func (x Person) Less(than Item) bool {

	if x.Birth.Before(than.(Person).Birth) {
		return true
	}
	if x.Birth.After(than.(Person).Birth) {
		return false
	}
	return x.Name < than.(Person).Name
}
