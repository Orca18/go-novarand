// Copyright (C) 2019-2022 Algorand, Inc.
// This file is part of go-algorand
//
// go-algorand is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// go-algorand is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with go-algorand.  If not, see <https://www.gnu.org/licenses/>.

package ledger

import (
	"github.com/Orca18/go-novarand/crypto"
	"github.com/Orca18/go-novarand/data/basics"
	"github.com/Orca18/go-novarand/ledger/ledgercore"
)

// CatchpointFileHeader is the content we would have in the "content.msgpack" file in the catchpoint tar archive.
// we need it to be public, as it's being decoded externally by the catchpointdump utility.
type CatchpointFileHeader struct {
	_struct struct{} `codec:",omitempty,omitemptyarray"`

	Version           uint64                   `codec:"version"`
	BalancesRound     basics.Round             `codec:"balancesRound"`
	BlocksRound       basics.Round             `codec:"blocksRound"`
	Totals            ledgercore.AccountTotals `codec:"accountTotals"`
	TotalAccounts     uint64                   `codec:"accountsCount"`
	TotalChunks       uint64                   `codec:"chunksCount"`
	Catchpoint        string                   `codec:"catchpoint"`
	BlockHeaderDigest crypto.Digest            `codec:"blockHeaderDigest"`
}
