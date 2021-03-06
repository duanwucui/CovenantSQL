package types

// Code generated by github.com/CovenantSQL/HashStablePack DO NOT EDIT.

import (
	hsp "github.com/CovenantSQL/HashStablePack/marshalhash"
)

// MarshalHash marshals for hash
func (z *TxBilling) MarshalHash() (o []byte, err error) {
	var b []byte
	o = hsp.Require(b, z.Msgsize())
	// map header, size 7
	o = append(o, 0x87, 0x87)
	if z.Signee == nil {
		o = hsp.AppendNil(o)
	} else {
		if oTemp, err := z.Signee.MarshalHash(); err != nil {
			return nil, err
		} else {
			o = hsp.AppendBytes(o, oTemp)
		}
	}
	o = append(o, 0x87)
	if z.Signature == nil {
		o = hsp.AppendNil(o)
	} else {
		if oTemp, err := z.Signature.MarshalHash(); err != nil {
			return nil, err
		} else {
			o = hsp.AppendBytes(o, oTemp)
		}
	}
	o = append(o, 0x87)
	if z.SignedBlock == nil {
		o = hsp.AppendNil(o)
	} else {
		if oTemp, err := z.SignedBlock.MarshalHash(); err != nil {
			return nil, err
		} else {
			o = hsp.AppendBytes(o, oTemp)
		}
	}
	o = append(o, 0x87)
	if z.TxHash == nil {
		o = hsp.AppendNil(o)
	} else {
		if oTemp, err := z.TxHash.MarshalHash(); err != nil {
			return nil, err
		} else {
			o = hsp.AppendBytes(o, oTemp)
		}
	}
	o = append(o, 0x87)
	if z.AccountAddress == nil {
		o = hsp.AppendNil(o)
	} else {
		if oTemp, err := z.AccountAddress.MarshalHash(); err != nil {
			return nil, err
		} else {
			o = hsp.AppendBytes(o, oTemp)
		}
	}
	o = append(o, 0x87)
	if oTemp, err := z.TxContent.MarshalHash(); err != nil {
		return nil, err
	} else {
		o = hsp.AppendBytes(o, oTemp)
	}
	o = append(o, 0x87)
	o = hsp.AppendByte(o, z.TxType)
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *TxBilling) Msgsize() (s int) {
	s = 1 + 7
	if z.Signee == nil {
		s += hsp.NilSize
	} else {
		s += z.Signee.Msgsize()
	}
	s += 10
	if z.Signature == nil {
		s += hsp.NilSize
	} else {
		s += z.Signature.Msgsize()
	}
	s += 12
	if z.SignedBlock == nil {
		s += hsp.NilSize
	} else {
		s += z.SignedBlock.Msgsize()
	}
	s += 7
	if z.TxHash == nil {
		s += hsp.NilSize
	} else {
		s += z.TxHash.Msgsize()
	}
	s += 15
	if z.AccountAddress == nil {
		s += hsp.NilSize
	} else {
		s += z.AccountAddress.Msgsize()
	}
	s += 10 + z.TxContent.Msgsize() + 7 + hsp.ByteSize
	return
}

// MarshalHash marshals for hash
func (z *TxContent) MarshalHash() (o []byte, err error) {
	var b []byte
	o = hsp.Require(b, z.Msgsize())
	// map header, size 6
	o = append(o, 0x86, 0x86)
	if oTemp, err := z.BillingRequest.MarshalHash(); err != nil {
		return nil, err
	} else {
		o = hsp.AppendBytes(o, oTemp)
	}
	o = append(o, 0x86)
	if oTemp, err := z.BillingResponse.MarshalHash(); err != nil {
		return nil, err
	} else {
		o = hsp.AppendBytes(o, oTemp)
	}
	o = append(o, 0x86)
	o = hsp.AppendArrayHeader(o, uint32(len(z.Receivers)))
	for za0001 := range z.Receivers {
		if z.Receivers[za0001] == nil {
			o = hsp.AppendNil(o)
		} else {
			if oTemp, err := z.Receivers[za0001].MarshalHash(); err != nil {
				return nil, err
			} else {
				o = hsp.AppendBytes(o, oTemp)
			}
		}
	}
	o = append(o, 0x86)
	o = hsp.AppendArrayHeader(o, uint32(len(z.Fees)))
	for za0002 := range z.Fees {
		o = hsp.AppendUint64(o, z.Fees[za0002])
	}
	o = append(o, 0x86)
	o = hsp.AppendArrayHeader(o, uint32(len(z.Rewards)))
	for za0003 := range z.Rewards {
		o = hsp.AppendUint64(o, z.Rewards[za0003])
	}
	o = append(o, 0x86)
	o = hsp.AppendUint32(o, z.SequenceID)
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *TxContent) Msgsize() (s int) {
	s = 1 + 15 + z.BillingRequest.Msgsize() + 16 + z.BillingResponse.Msgsize() + 10 + hsp.ArrayHeaderSize
	for za0001 := range z.Receivers {
		if z.Receivers[za0001] == nil {
			s += hsp.NilSize
		} else {
			s += z.Receivers[za0001].Msgsize()
		}
	}
	s += 5 + hsp.ArrayHeaderSize + (len(z.Fees) * (hsp.Uint64Size)) + 8 + hsp.ArrayHeaderSize + (len(z.Rewards) * (hsp.Uint64Size)) + 11 + hsp.Uint32Size
	return
}
