// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package fx

import "go.uber.org/fx/internal/lifecycle"

// Lifecycle allows constructors to register callbacks that are executed on
// application start and stop.
type Lifecycle interface {
	Append(Hook)
}

// A Hook is a pair of start and stop callbacks, either of which can be nil.
// If a Hook's OnStart callback isn't executed (because a previous OnStart
// failure short-circuited application start), its OnStop callback won't be
// executed.
type Hook struct {
	OnStart func() error
	OnStop  func() error
}

type lifecycleWrapper struct{ *lifecycle.Lifecycle }

func (l *lifecycleWrapper) Append(h Hook) {
	l.Lifecycle.Append(lifecycle.Hook{
		OnStart: h.OnStart,
		OnStop:  h.OnStop,
	})
}
