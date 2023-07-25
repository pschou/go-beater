// Copyright 2023 github.com/pschou
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// As a drummer beats a rhythm, the beater will call a passed in function until
// Stop() is called.  One can call for a Beat() directly either to cause an
// early fire or to sync up multiple beaters.

package beater

import "time"

type Beater struct {
	closed, skip bool
	next         time.Time
	fn           func()
}

func New(f func(), d time.Duration) *Beater {
	c := Beater{next: time.Now().Add(d), fn: f}
	d23, d43 := d*2/3, d+d/3
	go func() {
		time.Sleep(d)
		for !c.closed {
			if !c.skip {
				f()
			} else {
				c.skip = false
			}
			c.next = c.next.Add(d)
			if delta := c.next.Sub(time.Now()); delta > 0 && delta < d43 {
				time.Sleep(delta)
			} else if delta > 0 {
				c.next = c.next.Add(delta - (delta % d) + d)
				if sleep := d - (delta % d); sleep > d23 {
					time.Sleep(sleep)
				} else {
					time.Sleep(d23)
				}
			} else {
				c.next = c.next.Add(delta - (delta % d))
				if sleep := delta + (delta % d); sleep > d23 {
					time.Sleep(sleep)
				} else {
					time.Sleep(d23)
				}
			}
		}
	}()
	return &c
}

// Stop any future beats
func (t *Beater) Stop() {
	t.closed = true
}

// Like beat the drum, call the func() and all subsequent calls will be on the new schedule.
func (t *Beater) Beat() {
	t.skip, t.next = true, time.Now()
	t.fn()
}
