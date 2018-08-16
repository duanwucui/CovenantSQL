/*
 * Copyright 2018 The ThunderDB Authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package chain

import (
	bolt "github.com/coreos/bbolt"
	ci "gitlab.com/thunderdb/ThunderDB/chain/interfaces"
)

var (
	metaBucket        = [4]byte{0x0, 0x0, 0x0, 0x0}
	metaTxIndexBucket = []byte("covenantsql-tx-index-bucket")
)

type TxPersistence struct {
	db *bolt.DB
}

func NewTxPersistence(db *bolt.DB) (ins *TxPersistence, err error) {
	// Initialize buckets
	if err = db.Update(func(tx *bolt.Tx) (err error) {
		meta, err := tx.CreateBucketIfNotExists(metaBucket[:])
		if err != nil {
			return
		}
		_, err = meta.CreateBucketIfNotExists(metaTxIndexBucket)
		return
	}); err != nil {
		// Create instance if succeed
		ins = &TxPersistence{db: db}
	}
	return
}

func (p *TxPersistence) PutTransaction(tx ci.Transaction) (err error) {
	var key, value []byte
	key = tx.GetPersistenceKey()
	if value, err = tx.MarshalBinary(); err != nil {
		return
	}
	return p.db.Update(func(tx *bolt.Tx) error {
		return tx.Bucket(metaBucket[:]).Bucket(metaTxIndexBucket).Put(key, value)
	})
}

func (p *TxPersistence) GetTransaction(key []byte, tx ci.Transaction) (err error) {
	var value []byte
	if err = p.db.View(func(tx *bolt.Tx) error {
		value = tx.Bucket(metaBucket[:]).Bucket(metaTxIndexBucket).Get(key)
		return nil
	}); err != nil {
		return
	}
	return tx.UnmarshalBinary(value)
}

func (p *TxPersistence) DelTransaction(key []byte) (err error) {
	return p.db.Update(func(tx *bolt.Tx) error {
		return tx.Bucket(metaBucket[:]).Bucket(metaTxIndexBucket).Delete(key)
	})
}