// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package metaplex

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// If the auction manager is in Validated state, it can invoke the start command via calling this command here.
type StartAuction struct {

	// [0] = [WRITE] auctionManager
	// ··········· Auction manager
	//
	// [1] = [WRITE] auction
	// ··········· Auction
	//
	// [2] = [SIGNER] auctionManagerAuthority
	// ··········· Auction manager authority
	//
	// [3] = [] storeKey
	// ··········· Store key
	//
	// [4] = [] auctionProgram
	// ··········· Auction program
	//
	// [5] = [] clockSysvar
	// ··········· Clock sysvar
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewStartAuctionInstructionBuilder creates a new `StartAuction` instruction builder.
func NewStartAuctionInstructionBuilder() *StartAuction {
	nd := &StartAuction{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 6),
	}
	return nd
}

// SetAuctionManagerAccount sets the "auctionManager" account.
// Auction manager
func (inst *StartAuction) SetAuctionManagerAccount(auctionManager ag_solanago.PublicKey) *StartAuction {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(auctionManager).WRITE()
	return inst
}

// GetAuctionManagerAccount gets the "auctionManager" account.
// Auction manager
func (inst *StartAuction) GetAuctionManagerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}

// SetAuctionAccount sets the "auction" account.
// Auction
func (inst *StartAuction) SetAuctionAccount(auction ag_solanago.PublicKey) *StartAuction {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(auction).WRITE()
	return inst
}

// GetAuctionAccount gets the "auction" account.
// Auction
func (inst *StartAuction) GetAuctionAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[1]
}

// SetAuctionManagerAuthorityAccount sets the "auctionManagerAuthority" account.
// Auction manager authority
func (inst *StartAuction) SetAuctionManagerAuthorityAccount(auctionManagerAuthority ag_solanago.PublicKey) *StartAuction {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(auctionManagerAuthority).SIGNER()
	return inst
}

// GetAuctionManagerAuthorityAccount gets the "auctionManagerAuthority" account.
// Auction manager authority
func (inst *StartAuction) GetAuctionManagerAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[2]
}

// SetStoreKeyAccount sets the "storeKey" account.
// Store key
func (inst *StartAuction) SetStoreKeyAccount(storeKey ag_solanago.PublicKey) *StartAuction {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(storeKey)
	return inst
}

// GetStoreKeyAccount gets the "storeKey" account.
// Store key
func (inst *StartAuction) GetStoreKeyAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[3]
}

// SetAuctionProgramAccount sets the "auctionProgram" account.
// Auction program
func (inst *StartAuction) SetAuctionProgramAccount(auctionProgram ag_solanago.PublicKey) *StartAuction {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(auctionProgram)
	return inst
}

// GetAuctionProgramAccount gets the "auctionProgram" account.
// Auction program
func (inst *StartAuction) GetAuctionProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[4]
}

// SetClockSysvarAccount sets the "clockSysvar" account.
// Clock sysvar
func (inst *StartAuction) SetClockSysvarAccount(clockSysvar ag_solanago.PublicKey) *StartAuction {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(clockSysvar)
	return inst
}

// GetClockSysvarAccount gets the "clockSysvar" account.
// Clock sysvar
func (inst *StartAuction) GetClockSysvarAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[5]
}

func (inst StartAuction) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: ag_binary.TypeIDFromUint8(Instruction_StartAuction),
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst StartAuction) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *StartAuction) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.AuctionManager is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Auction is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.AuctionManagerAuthority is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.StoreKey is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.AuctionProgram is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.ClockSysvar is not set")
		}
	}
	return nil
}

func (inst *StartAuction) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("StartAuction")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=6]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("         auctionManager", inst.AccountMetaSlice[0]))
						accountsBranch.Child(ag_format.Meta("                auction", inst.AccountMetaSlice[1]))
						accountsBranch.Child(ag_format.Meta("auctionManagerAuthority", inst.AccountMetaSlice[2]))
						accountsBranch.Child(ag_format.Meta("               storeKey", inst.AccountMetaSlice[3]))
						accountsBranch.Child(ag_format.Meta("         auctionProgram", inst.AccountMetaSlice[4]))
						accountsBranch.Child(ag_format.Meta("            clockSysvar", inst.AccountMetaSlice[5]))
					})
				})
		})
}

func (obj StartAuction) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *StartAuction) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewStartAuctionInstruction declares a new StartAuction instruction with the provided parameters and accounts.
func NewStartAuctionInstruction(
	// Accounts:
	auctionManager ag_solanago.PublicKey,
	auction ag_solanago.PublicKey,
	auctionManagerAuthority ag_solanago.PublicKey,
	storeKey ag_solanago.PublicKey,
	auctionProgram ag_solanago.PublicKey,
	clockSysvar ag_solanago.PublicKey) *StartAuction {
	return NewStartAuctionInstructionBuilder().
		SetAuctionManagerAccount(auctionManager).
		SetAuctionAccount(auction).
		SetAuctionManagerAuthorityAccount(auctionManagerAuthority).
		SetStoreKeyAccount(storeKey).
		SetAuctionProgramAccount(auctionProgram).
		SetClockSysvarAccount(clockSysvar)
}
