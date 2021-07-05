// Copyright Turing Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rpc

// RetrieveBackupTx construction
type RetrieveBackupTx struct {
	BackupAddr  string `json:"backupAddr"`
	DefaultAddr string `json:"defaultAddr"`
	DelayPeriod int64  `json:"delayPeriod"`
	Fee         int64  `json:"fee"`
}

// RetrievePrepareTx construction
type RetrievePrepareTx struct {
	BackupAddr  string `json:"backupAddr"`
	DefaultAddr string `json:"defaultAddr"`
	Fee         int64  `json:"fee"`
}

// Asset Asset
type Asset struct {
	Exec   string `json:"exec"`
	Symbol string `json:"symbol"`
}

// RetrievePerformTx construction
type RetrievePerformTx struct {
	BackupAddr  string  `json:"backupAddr"`
	DefaultAddr string  `json:"defaultAddr"`
	Assets      []Asset `json:"assets"`
	Fee         int64   `json:"fee"`
}

// RetrieveCancelTx construction
type RetrieveCancelTx struct {
	BackupAddr  string `json:"backupAddr"`
	DefaultAddr string `json:"defaultAddr"`
	Fee         int64  `json:"fee"`
}
