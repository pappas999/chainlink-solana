// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package store

import (
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
)

type Flags struct {
	Xs  [128]ag_solanago.PublicKey
	Len uint64
}

func (obj Flags) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Xs` param:
	err = encoder.Encode(obj.Xs)
	if err != nil {
		return err
	}
	// Serialize `Len` param:
	err = encoder.Encode(obj.Len)
	if err != nil {
		return err
	}
	return nil
}

func (obj *Flags) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Xs`:
	err = decoder.Decode(&obj.Xs)
	if err != nil {
		return err
	}
	// Deserialize `Len`:
	err = decoder.Decode(&obj.Len)
	if err != nil {
		return err
	}
	return nil
}