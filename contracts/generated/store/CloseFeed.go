// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package store

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// CloseFeed is the `closeFeed` instruction.
type CloseFeed struct {

	// [0] = [] store
	//
	// [1] = [WRITE] feed
	//
	// [2] = [WRITE] receiver
	//
	// [3] = [SIGNER] authority
	ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

// NewCloseFeedInstructionBuilder creates a new `CloseFeed` instruction builder.
func NewCloseFeedInstructionBuilder() *CloseFeed {
	nd := &CloseFeed{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 4),
	}
	return nd
}

// SetStoreAccount sets the "store" account.
func (inst *CloseFeed) SetStoreAccount(store ag_solanago.PublicKey) *CloseFeed {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(store)
	return inst
}

// GetStoreAccount gets the "store" account.
func (inst *CloseFeed) GetStoreAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}

// SetFeedAccount sets the "feed" account.
func (inst *CloseFeed) SetFeedAccount(feed ag_solanago.PublicKey) *CloseFeed {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(feed).WRITE()
	return inst
}

// GetFeedAccount gets the "feed" account.
func (inst *CloseFeed) GetFeedAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[1]
}

// SetReceiverAccount sets the "receiver" account.
func (inst *CloseFeed) SetReceiverAccount(receiver ag_solanago.PublicKey) *CloseFeed {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(receiver).WRITE()
	return inst
}

// GetReceiverAccount gets the "receiver" account.
func (inst *CloseFeed) GetReceiverAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[2]
}

// SetAuthorityAccount sets the "authority" account.
func (inst *CloseFeed) SetAuthorityAccount(authority ag_solanago.PublicKey) *CloseFeed {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(authority).SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *CloseFeed) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[3]
}

func (inst CloseFeed) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_CloseFeed,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst CloseFeed) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *CloseFeed) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Store is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Feed is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Receiver is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.Authority is not set")
		}
	}
	return nil
}

func (inst *CloseFeed) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("CloseFeed")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=4]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("    store", inst.AccountMetaSlice[0]))
						accountsBranch.Child(ag_format.Meta("     feed", inst.AccountMetaSlice[1]))
						accountsBranch.Child(ag_format.Meta(" receiver", inst.AccountMetaSlice[2]))
						accountsBranch.Child(ag_format.Meta("authority", inst.AccountMetaSlice[3]))
					})
				})
		})
}

func (obj CloseFeed) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *CloseFeed) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewCloseFeedInstruction declares a new CloseFeed instruction with the provided parameters and accounts.
func NewCloseFeedInstruction(
	// Accounts:
	store ag_solanago.PublicKey,
	feed ag_solanago.PublicKey,
	receiver ag_solanago.PublicKey,
	authority ag_solanago.PublicKey) *CloseFeed {
	return NewCloseFeedInstructionBuilder().
		SetStoreAccount(store).
		SetFeedAccount(feed).
		SetReceiverAccount(receiver).
		SetAuthorityAccount(authority)
}
