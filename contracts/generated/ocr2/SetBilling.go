// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package ocr_2

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// SetBilling is the `setBilling` instruction.
type SetBilling struct {
	ObservationPaymentGjuels  *uint32
	TransmissionPaymentGjuels *uint32

	// [0] = [WRITE] state
	//
	// [1] = [SIGNER] authority
	//
	// [2] = [] accessController
	ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

// NewSetBillingInstructionBuilder creates a new `SetBilling` instruction builder.
func NewSetBillingInstructionBuilder() *SetBilling {
	nd := &SetBilling{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 3),
	}
	return nd
}

// SetObservationPaymentGjuels sets the "observationPaymentGjuels" parameter.
func (inst *SetBilling) SetObservationPaymentGjuels(observationPaymentGjuels uint32) *SetBilling {
	inst.ObservationPaymentGjuels = &observationPaymentGjuels
	return inst
}

// SetTransmissionPaymentGjuels sets the "transmissionPaymentGjuels" parameter.
func (inst *SetBilling) SetTransmissionPaymentGjuels(transmissionPaymentGjuels uint32) *SetBilling {
	inst.TransmissionPaymentGjuels = &transmissionPaymentGjuels
	return inst
}

// SetStateAccount sets the "state" account.
func (inst *SetBilling) SetStateAccount(state ag_solanago.PublicKey) *SetBilling {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(state).WRITE()
	return inst
}

// GetStateAccount gets the "state" account.
func (inst *SetBilling) GetStateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}

// SetAuthorityAccount sets the "authority" account.
func (inst *SetBilling) SetAuthorityAccount(authority ag_solanago.PublicKey) *SetBilling {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(authority).SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *SetBilling) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[1]
}

// SetAccessControllerAccount sets the "accessController" account.
func (inst *SetBilling) SetAccessControllerAccount(accessController ag_solanago.PublicKey) *SetBilling {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(accessController)
	return inst
}

// GetAccessControllerAccount gets the "accessController" account.
func (inst *SetBilling) GetAccessControllerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[2]
}

func (inst SetBilling) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_SetBilling,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst SetBilling) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *SetBilling) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.ObservationPaymentGjuels == nil {
			return errors.New("ObservationPaymentGjuels parameter is not set")
		}
		if inst.TransmissionPaymentGjuels == nil {
			return errors.New("TransmissionPaymentGjuels parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.State is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Authority is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.AccessController is not set")
		}
	}
	return nil
}

func (inst *SetBilling) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("SetBilling")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=2]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param(" ObservationPaymentGjuels", *inst.ObservationPaymentGjuels))
						paramsBranch.Child(ag_format.Param("TransmissionPaymentGjuels", *inst.TransmissionPaymentGjuels))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=3]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("           state", inst.AccountMetaSlice[0]))
						accountsBranch.Child(ag_format.Meta("       authority", inst.AccountMetaSlice[1]))
						accountsBranch.Child(ag_format.Meta("accessController", inst.AccountMetaSlice[2]))
					})
				})
		})
}

func (obj SetBilling) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
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
func (obj *SetBilling) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
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

// NewSetBillingInstruction declares a new SetBilling instruction with the provided parameters and accounts.
func NewSetBillingInstruction(
	// Parameters:
	observationPaymentGjuels uint32,
	transmissionPaymentGjuels uint32,
	// Accounts:
	state ag_solanago.PublicKey,
	authority ag_solanago.PublicKey,
	accessController ag_solanago.PublicKey) *SetBilling {
	return NewSetBillingInstructionBuilder().
		SetObservationPaymentGjuels(observationPaymentGjuels).
		SetTransmissionPaymentGjuels(transmissionPaymentGjuels).
		SetStateAccount(state).
		SetAuthorityAccount(authority).
		SetAccessControllerAccount(accessController)
}
