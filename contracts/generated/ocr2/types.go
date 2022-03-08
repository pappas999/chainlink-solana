// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package ocr_2

import (
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
)

type NewOracle struct {
	Signer      [20]uint8
	Transmitter ag_solanago.PublicKey
}

func (obj NewOracle) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Signer` param:
	err = encoder.Encode(obj.Signer)
	if err != nil {
		return err
	}
	// Serialize `Transmitter` param:
	err = encoder.Encode(obj.Transmitter)
	if err != nil {
		return err
	}
	return nil
}

func (obj *NewOracle) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Signer`:
	err = decoder.Decode(&obj.Signer)
	if err != nil {
		return err
	}
	// Deserialize `Transmitter`:
	err = decoder.Decode(&obj.Transmitter)
	if err != nil {
		return err
	}
	return nil
}

type Billing struct {
	ObservationPaymentGjuels  uint32
	TransmissionPaymentGjuels uint32
}

func (obj Billing) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `ObservationPaymentGjuels` param:
	err = encoder.Encode(obj.ObservationPaymentGjuels)
	if err != nil {
		return err
	}
	// Serialize `TransmissionPaymentGjuels` param:
	err = encoder.Encode(obj.TransmissionPaymentGjuels)
	if err != nil {
		return err
	}
	return nil
}

func (obj *Billing) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `ObservationPaymentGjuels`:
	err = decoder.Decode(&obj.ObservationPaymentGjuels)
	if err != nil {
		return err
	}
	// Deserialize `TransmissionPaymentGjuels`:
	err = decoder.Decode(&obj.TransmissionPaymentGjuels)
	if err != nil {
		return err
	}
	return nil
}

type Oracles struct {
	Xs  [19]Oracle
	Len uint64
}

func (obj Oracles) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
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

func (obj *Oracles) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
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

type ProposedOracle struct {
	Transmitter ag_solanago.PublicKey
	Signer      SigningKey
	Padding     uint32
	Payee       ag_solanago.PublicKey
}

func (obj ProposedOracle) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Transmitter` param:
	err = encoder.Encode(obj.Transmitter)
	if err != nil {
		return err
	}
	// Serialize `Signer` param:
	err = encoder.Encode(obj.Signer)
	if err != nil {
		return err
	}
	// Serialize `Padding` param:
	err = encoder.Encode(obj.Padding)
	if err != nil {
		return err
	}
	// Serialize `Payee` param:
	err = encoder.Encode(obj.Payee)
	if err != nil {
		return err
	}
	return nil
}

func (obj *ProposedOracle) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Transmitter`:
	err = decoder.Decode(&obj.Transmitter)
	if err != nil {
		return err
	}
	// Deserialize `Signer`:
	err = decoder.Decode(&obj.Signer)
	if err != nil {
		return err
	}
	// Deserialize `Padding`:
	err = decoder.Decode(&obj.Padding)
	if err != nil {
		return err
	}
	// Deserialize `Payee`:
	err = decoder.Decode(&obj.Payee)
	if err != nil {
		return err
	}
	return nil
}

type ProposedOracles struct {
	Xs  [19]ProposedOracle
	Len uint64
}

func (obj ProposedOracles) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
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

func (obj *ProposedOracles) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
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

type OffchainConfig struct {
	Version uint64
	Xs      [4096]uint8
	Len     uint64
}

func (obj OffchainConfig) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Version` param:
	err = encoder.Encode(obj.Version)
	if err != nil {
		return err
	}
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

func (obj *OffchainConfig) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Version`:
	err = decoder.Decode(&obj.Version)
	if err != nil {
		return err
	}
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

type Config struct {
	Owner                     ag_solanago.PublicKey
	ProposedOwner             ag_solanago.PublicKey
	TokenMint                 ag_solanago.PublicKey
	TokenVault                ag_solanago.PublicKey
	RequesterAccessController ag_solanago.PublicKey
	BillingAccessController   ag_solanago.PublicKey
	MinAnswer                 ag_binary.Int128
	MaxAnswer                 ag_binary.Int128
	F                         uint8
	Round                     uint8
	Padding0                  uint16
	Epoch                     uint32
	LatestAggregatorRoundId   uint32
	LatestTransmitter         ag_solanago.PublicKey
	ConfigCount               uint32
	LatestConfigDigest        [32]uint8
	LatestConfigBlockNumber   uint64
	Billing                   Billing
}

func (obj Config) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
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
	// Serialize `TokenMint` param:
	err = encoder.Encode(obj.TokenMint)
	if err != nil {
		return err
	}
	// Serialize `TokenVault` param:
	err = encoder.Encode(obj.TokenVault)
	if err != nil {
		return err
	}
	// Serialize `RequesterAccessController` param:
	err = encoder.Encode(obj.RequesterAccessController)
	if err != nil {
		return err
	}
	// Serialize `BillingAccessController` param:
	err = encoder.Encode(obj.BillingAccessController)
	if err != nil {
		return err
	}
	// Serialize `MinAnswer` param:
	err = encoder.Encode(obj.MinAnswer)
	if err != nil {
		return err
	}
	// Serialize `MaxAnswer` param:
	err = encoder.Encode(obj.MaxAnswer)
	if err != nil {
		return err
	}
	// Serialize `F` param:
	err = encoder.Encode(obj.F)
	if err != nil {
		return err
	}
	// Serialize `Round` param:
	err = encoder.Encode(obj.Round)
	if err != nil {
		return err
	}
	// Serialize `Padding0` param:
	err = encoder.Encode(obj.Padding0)
	if err != nil {
		return err
	}
	// Serialize `Epoch` param:
	err = encoder.Encode(obj.Epoch)
	if err != nil {
		return err
	}
	// Serialize `LatestAggregatorRoundId` param:
	err = encoder.Encode(obj.LatestAggregatorRoundId)
	if err != nil {
		return err
	}
	// Serialize `LatestTransmitter` param:
	err = encoder.Encode(obj.LatestTransmitter)
	if err != nil {
		return err
	}
	// Serialize `ConfigCount` param:
	err = encoder.Encode(obj.ConfigCount)
	if err != nil {
		return err
	}
	// Serialize `LatestConfigDigest` param:
	err = encoder.Encode(obj.LatestConfigDigest)
	if err != nil {
		return err
	}
	// Serialize `LatestConfigBlockNumber` param:
	err = encoder.Encode(obj.LatestConfigBlockNumber)
	if err != nil {
		return err
	}
	// Serialize `Billing` param:
	err = encoder.Encode(obj.Billing)
	if err != nil {
		return err
	}
	return nil
}

func (obj *Config) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
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
	// Deserialize `TokenMint`:
	err = decoder.Decode(&obj.TokenMint)
	if err != nil {
		return err
	}
	// Deserialize `TokenVault`:
	err = decoder.Decode(&obj.TokenVault)
	if err != nil {
		return err
	}
	// Deserialize `RequesterAccessController`:
	err = decoder.Decode(&obj.RequesterAccessController)
	if err != nil {
		return err
	}
	// Deserialize `BillingAccessController`:
	err = decoder.Decode(&obj.BillingAccessController)
	if err != nil {
		return err
	}
	// Deserialize `MinAnswer`:
	err = decoder.Decode(&obj.MinAnswer)
	if err != nil {
		return err
	}
	// Deserialize `MaxAnswer`:
	err = decoder.Decode(&obj.MaxAnswer)
	if err != nil {
		return err
	}
	// Deserialize `F`:
	err = decoder.Decode(&obj.F)
	if err != nil {
		return err
	}
	// Deserialize `Round`:
	err = decoder.Decode(&obj.Round)
	if err != nil {
		return err
	}
	// Deserialize `Padding0`:
	err = decoder.Decode(&obj.Padding0)
	if err != nil {
		return err
	}
	// Deserialize `Epoch`:
	err = decoder.Decode(&obj.Epoch)
	if err != nil {
		return err
	}
	// Deserialize `LatestAggregatorRoundId`:
	err = decoder.Decode(&obj.LatestAggregatorRoundId)
	if err != nil {
		return err
	}
	// Deserialize `LatestTransmitter`:
	err = decoder.Decode(&obj.LatestTransmitter)
	if err != nil {
		return err
	}
	// Deserialize `ConfigCount`:
	err = decoder.Decode(&obj.ConfigCount)
	if err != nil {
		return err
	}
	// Deserialize `LatestConfigDigest`:
	err = decoder.Decode(&obj.LatestConfigDigest)
	if err != nil {
		return err
	}
	// Deserialize `LatestConfigBlockNumber`:
	err = decoder.Decode(&obj.LatestConfigBlockNumber)
	if err != nil {
		return err
	}
	// Deserialize `Billing`:
	err = decoder.Decode(&obj.Billing)
	if err != nil {
		return err
	}
	return nil
}

type SigningKey struct {
	Key [20]uint8
}

func (obj SigningKey) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Key` param:
	err = encoder.Encode(obj.Key)
	if err != nil {
		return err
	}
	return nil
}

func (obj *SigningKey) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Key`:
	err = decoder.Decode(&obj.Key)
	if err != nil {
		return err
	}
	return nil
}

type Oracle struct {
	Transmitter   ag_solanago.PublicKey
	Signer        SigningKey
	Payee         ag_solanago.PublicKey
	ProposedPayee ag_solanago.PublicKey
	FromRoundId   uint32
	PaymentGjuels uint64
}

func (obj Oracle) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Transmitter` param:
	err = encoder.Encode(obj.Transmitter)
	if err != nil {
		return err
	}
	// Serialize `Signer` param:
	err = encoder.Encode(obj.Signer)
	if err != nil {
		return err
	}
	// Serialize `Payee` param:
	err = encoder.Encode(obj.Payee)
	if err != nil {
		return err
	}
	// Serialize `ProposedPayee` param:
	err = encoder.Encode(obj.ProposedPayee)
	if err != nil {
		return err
	}
	// Serialize `FromRoundId` param:
	err = encoder.Encode(obj.FromRoundId)
	if err != nil {
		return err
	}
	// Serialize `PaymentGjuels` param:
	err = encoder.Encode(obj.PaymentGjuels)
	if err != nil {
		return err
	}
	return nil
}

func (obj *Oracle) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Transmitter`:
	err = decoder.Decode(&obj.Transmitter)
	if err != nil {
		return err
	}
	// Deserialize `Signer`:
	err = decoder.Decode(&obj.Signer)
	if err != nil {
		return err
	}
	// Deserialize `Payee`:
	err = decoder.Decode(&obj.Payee)
	if err != nil {
		return err
	}
	// Deserialize `ProposedPayee`:
	err = decoder.Decode(&obj.ProposedPayee)
	if err != nil {
		return err
	}
	// Deserialize `FromRoundId`:
	err = decoder.Decode(&obj.FromRoundId)
	if err != nil {
		return err
	}
	// Deserialize `PaymentGjuels`:
	err = decoder.Decode(&obj.PaymentGjuels)
	if err != nil {
		return err
	}
	return nil
}
