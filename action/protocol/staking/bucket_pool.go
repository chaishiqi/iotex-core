// Copyright (c) 2022 IoTeX Foundation
// This source code is provided 'as is' and no warranties are given as to title or non-infringement, merchantability
// or fitness for purpose and, to the extent permitted by law, all liability for your use of the code is disclaimed.
// This source code is governed by Apache License 2.0 that can be found in the LICENSE file.

package staking

import (
	"math/big"

	"github.com/iotexproject/go-pkgs/hash"
	"google.golang.org/protobuf/proto"

	"github.com/iotexproject/iotex-core/v2/action/protocol"
	"github.com/iotexproject/iotex-core/v2/action/protocol/staking/stakingpb"
	"github.com/iotexproject/iotex-core/v2/state"
)

// const
const (
	_stakingBucketPool = "bucketPool"
)

var (
	_bucketPoolAddr    = hash.Hash160b([]byte(_stakingBucketPool))
	_bucketPoolAddrKey = append([]byte{_const}, _bucketPoolAddr[:]...)
)

// when a bucket is created, the amount of staked IOTX token is deducted from user, but does not transfer to any address
// in the same way when a bucket are withdrawn, bucket amount is added back to user, but does not come out of any address
//
// for better accounting/auditing, we take protocol's address as the 'bucket pool' address
// 1. at Greenland height we sum up all existing bucket's amount and set the total amount to bucket pool address
// 2. for future bucket creation/deposit/registration, the amount of staked IOTX token will be added to bucket pool (so
//    the pool is 'receiving' token)
// 3. for future bucket withdrawal, the bucket amount will be deducted from bucket pool (so the pool is 'releasing' token)

type (
	// BucketPool implements the bucket pool
	BucketPool struct {
		enableSMStorage bool
		dirty           bool
		total           *totalAmount
	}

	totalAmount struct {
		amount *big.Int
		count  uint64
	}
)

func (t *totalAmount) Serialize() ([]byte, error) {
	gen := stakingpb.TotalAmount{
		Amount: t.amount.String(),
		Count:  t.count,
	}
	return proto.Marshal(&gen)
}

func (t *totalAmount) Deserialize(data []byte) error {
	gen := stakingpb.TotalAmount{}
	if err := proto.Unmarshal(data, &gen); err != nil {
		return err
	}

	var ok bool
	if t.amount, ok = new(big.Int).SetString(gen.Amount, 10); !ok {
		return state.ErrStateDeserialization
	}

	if t.amount.Cmp(big.NewInt(0)) == -1 {
		return state.ErrNotEnoughBalance
	}
	t.count = gen.Count
	return nil
}

func (t *totalAmount) AddBalance(amount *big.Int, newBucket bool) {
	t.amount.Add(t.amount, amount)
	if newBucket {
		t.count++
	}
}

func (t *totalAmount) SubBalance(amount *big.Int) error {
	if amount.Cmp(t.amount) == 1 || t.count == 0 {
		return state.ErrNotEnoughBalance
	}
	t.amount.Sub(t.amount, amount)
	t.count--
	return nil
}

// IsDirty returns true if the bucket pool is dirty
func (bp *BucketPool) IsDirty() bool {
	return bp.dirty
}

// Total returns the total amount staked in bucket pool
func (bp *BucketPool) Total() *big.Int {
	return new(big.Int).Set(bp.total.amount)
}

// Count returns the total number of buckets in bucket pool
func (bp *BucketPool) Count() uint64 {
	return bp.total.count
}

// EnableSMStorage enables state manager storage
func (bp *BucketPool) EnableSMStorage() {
	bp.enableSMStorage = true
}

// Clone returns a copy of the bucket pool
func (bp *BucketPool) Clone() *BucketPool {
	pool := BucketPool{}
	pool.enableSMStorage = bp.enableSMStorage
	pool.total = &totalAmount{
		amount: new(big.Int).Set(bp.total.amount),
		count:  bp.total.count,
	}
	pool.dirty = bp.dirty
	return &pool
}

// Commit is called upon workingset commit
func (bp *BucketPool) Commit(sr protocol.StateReader) error {
	bp.dirty = false
	return nil
}

// CreditPool subtracts staked amount out of the pool
func (bp *BucketPool) CreditPool(sm protocol.StateManager, amount *big.Int) error {
	if err := bp.total.SubBalance(amount); err != nil {
		return err
	}

	if !bp.enableSMStorage {
		bp.dirty = true
		return nil
	}
	_, err := sm.PutState(bp.total, protocol.NamespaceOption(_stakingNameSpace), protocol.KeyOption(_bucketPoolAddrKey))
	return err
}

// DebitPool adds staked amount into the pool
func (bp *BucketPool) DebitPool(sm protocol.StateManager, amount *big.Int, newBucket bool) error {
	bp.total.AddBalance(amount, newBucket)
	if !bp.enableSMStorage {
		bp.dirty = true
		return nil
	}
	_, err := sm.PutState(bp.total, protocol.NamespaceOption(_stakingNameSpace), protocol.KeyOption(_bucketPoolAddrKey))
	return err
}
