// Copyright (c) 2016 The btcsuite developers
// Copyright (c) 2019 The payt DAG developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

// Package helpers provides convenience functions to simplify wallet code.  This
// package is intended for internal wallet use only.
package helpers

import (
	"github.com/payt-dag/payt/paytutil"
	"github.com/payt-dag/payt/wire"
)

// SumOutputValues sums up the list of TxOuts and returns an Amount.
func SumOutputValues(outputs []*wire.TxOut) (totalOutput paytutil.Amount) {
	for _, txOut := range outputs {
		totalOutput += paytutil.Amount(txOut.Value)
	}
	return totalOutput
}

// SumOutputSerializeSizes sums up the serialized size of the supplied outputs.
func SumOutputSerializeSizes(outputs []*wire.TxOut) (serializeSize int) {
	for _, txOut := range outputs {
		serializeSize += txOut.SerializeSize()
	}
	return serializeSize
}
