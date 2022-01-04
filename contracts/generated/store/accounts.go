// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package store

import (
	"fmt"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
)

type Store struct {
	Owner                    ag_solanago.PublicKey
	ProposedOwner            ag_solanago.PublicKey
	RaisingAccessController  ag_solanago.PublicKey
	LoweringAccessController ag_solanago.PublicKey
	Flags                    Flags
}

var StoreDiscriminator = [8]byte{130, 48, 247, 244, 182, 191, 30, 26}

func (obj Store) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(StoreDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `Owner` param:
	err = encoder.Encode(obj.Owner)
	if err != nil {
		return err
	}
	// Serialize `ProposedOwner` param:
	err = encoder.Encode(obj.ProposedOwner)
	if err != nil {
		return err
	}
	// Serialize `RaisingAccessController` param:
	err = encoder.Encode(obj.RaisingAccessController)
	if err != nil {
		return err
	}
	// Serialize `LoweringAccessController` param:
	err = encoder.Encode(obj.LoweringAccessController)
	if err != nil {
		return err
	}
	// Serialize `Flags` param:
	err = encoder.Encode(obj.Flags)
	if err != nil {
		return err
	}
	return nil
}

func (obj *Store) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(StoreDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[130 48 247 244 182 191 30 26]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Owner`:
	err = decoder.Decode(&obj.Owner)
	if err != nil {
		return err
	}
	// Deserialize `ProposedOwner`:
	err = decoder.Decode(&obj.ProposedOwner)
	if err != nil {
		return err
	}
	// Deserialize `RaisingAccessController`:
	err = decoder.Decode(&obj.RaisingAccessController)
	if err != nil {
		return err
	}
	// Deserialize `LoweringAccessController`:
	err = decoder.Decode(&obj.LoweringAccessController)
	if err != nil {
		return err
	}
	// Deserialize `Flags`:
	err = decoder.Decode(&obj.Flags)
	if err != nil {
		return err
	}
	return nil
}

type Transmissions struct {
	Version           uint8
	Store             ag_solanago.PublicKey
	Writer            ag_solanago.PublicKey
	FlaggingThreshold uint32
	LatestRoundId     uint32
	Granularity       uint8
	LiveLength        uint32
	LiveCursor        uint32
	HistoricalCursor  uint32
}

var TransmissionsDiscriminator = [8]byte{96, 179, 69, 66, 128, 129, 73, 117}

func (obj Transmissions) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(TransmissionsDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `Version` param:
	err = encoder.Encode(obj.Version)
	if err != nil {
		return err
	}
	// Serialize `Store` param:
	err = encoder.Encode(obj.Store)
	if err != nil {
		return err
	}
	// Serialize `Writer` param:
	err = encoder.Encode(obj.Writer)
	if err != nil {
		return err
	}
	// Serialize `FlaggingThreshold` param:
	err = encoder.Encode(obj.FlaggingThreshold)
	if err != nil {
		return err
	}
	// Serialize `LatestRoundId` param:
	err = encoder.Encode(obj.LatestRoundId)
	if err != nil {
		return err
	}
	// Serialize `Granularity` param:
	err = encoder.Encode(obj.Granularity)
	if err != nil {
		return err
	}
	// Serialize `LiveLength` param:
	err = encoder.Encode(obj.LiveLength)
	if err != nil {
		return err
	}
	// Serialize `LiveCursor` param:
	err = encoder.Encode(obj.LiveCursor)
	if err != nil {
		return err
	}
	// Serialize `HistoricalCursor` param:
	err = encoder.Encode(obj.HistoricalCursor)
	if err != nil {
		return err
	}
	return nil
}

func (obj *Transmissions) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(TransmissionsDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[96 179 69 66 128 129 73 117]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Version`:
	err = decoder.Decode(&obj.Version)
	if err != nil {
		return err
	}
	// Deserialize `Store`:
	err = decoder.Decode(&obj.Store)
	if err != nil {
		return err
	}
	// Deserialize `Writer`:
	err = decoder.Decode(&obj.Writer)
	if err != nil {
		return err
	}
	// Deserialize `FlaggingThreshold`:
	err = decoder.Decode(&obj.FlaggingThreshold)
	if err != nil {
		return err
	}
	// Deserialize `LatestRoundId`:
	err = decoder.Decode(&obj.LatestRoundId)
	if err != nil {
		return err
	}
	// Deserialize `Granularity`:
	err = decoder.Decode(&obj.Granularity)
	if err != nil {
		return err
	}
	// Deserialize `LiveLength`:
	err = decoder.Decode(&obj.LiveLength)
	if err != nil {
		return err
	}
	// Deserialize `LiveCursor`:
	err = decoder.Decode(&obj.LiveCursor)
	if err != nil {
		return err
	}
	// Deserialize `HistoricalCursor`:
	err = decoder.Decode(&obj.HistoricalCursor)
	if err != nil {
		return err
	}
	return nil
}
