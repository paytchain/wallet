// Copyright (c) 2015 The btcsuite developers
// Copyright (c) 2019 The payt DAG developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package votingpool

import "github.com/payt-dag/payt/paytlog"

// log is a logger that is initialized with no output filters.  This
// means the package will not perform any logging by default until the caller
// requests it.
var log paytlog.Logger

// The default amount of logging is none.
func init() {
	DisableLog()
}

// DisableLog disables all library log output.  Logging output is disabled
// by default until either UseLogger or SetLogWriter are called.
func DisableLog() {
	log = paytlog.Disabled
}

// UseLogger uses a specified Logger to output package logging info.
// This should be used in preference to SetLogWriter if the caller is also
// using paytlog.
func UseLogger(logger paytlog.Logger) {
	log = logger
}
