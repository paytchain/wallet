// Copyright (c) 2013-2015 The btcsuite developers
// Copyright (c) 2019 The paytia DAG developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package legacyrpc

import "github.com/paytia-dag/paytd/paytlog"

var log = paytlog.Disabled

// UseLogger sets the package-wide logger.  Any calls to this function must be
// made before a server is created and used (it is not concurrent safe).
func UseLogger(logger paytlog.Logger) {
	log = logger
}
