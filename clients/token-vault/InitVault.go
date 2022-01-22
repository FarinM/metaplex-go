// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package token_vault

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Initialize a token vault, starts inactivate. Add tokens in subsequent instructions, then activate.
type InitVault struct {
	Args *InitVaultArgs

	// [0] = [WRITE] initializedFractionalShareMint
	// ··········· Initialized fractional share mint with 0 tokens in supply, authority on mint must be pda of program with seed [prefix, programid]
	//
	// [1] = [WRITE] initializedRedeemTreasuryTokenAccount
	// ··········· Initialized redeem treasury token account with 0 tokens in supply, owner of account must be pda of program like above
	//
	// [2] = [WRITE] initializedFractionTreasuryTokenAccount
	// ··········· Initialized fraction treasury token account with 0 tokens in supply, owner of account must be pda of program like above
	//
	// [3] = [WRITE] uninitializedVault
	// ··········· Uninitialized vault account
	//
	// [4] = [] vaultAuthority
	// ··········· Authority on the vault
	//
	// [5] = [] pricingLookupAddress
	// ··········· Pricing Lookup Address
	//
	// [6] = [] tokenProgram
	// ··········· Token program
	//
	// [7] = [] rentSysvar
	// ··········· Rent sysvar
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewInitVaultInstructionBuilder creates a new `InitVault` instruction builder.
func NewInitVaultInstructionBuilder() *InitVault {
	nd := &InitVault{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 8),
	}
	return nd
}

// SetArgs sets the "args" parameter.
func (inst *InitVault) SetArgs(args InitVaultArgs) *InitVault {
	inst.Args = &args
	return inst
}

// SetInitializedFractionalShareMintAccount sets the "initializedFractionalShareMint" account.
// Initialized fractional share mint with 0 tokens in supply, authority on mint must be pda of program with seed [prefix, programid]
func (inst *InitVault) SetInitializedFractionalShareMintAccount(initializedFractionalShareMint ag_solanago.PublicKey) *InitVault {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(initializedFractionalShareMint).WRITE()
	return inst
}

// GetInitializedFractionalShareMintAccount gets the "initializedFractionalShareMint" account.
// Initialized fractional share mint with 0 tokens in supply, authority on mint must be pda of program with seed [prefix, programid]
func (inst *InitVault) GetInitializedFractionalShareMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}

// SetInitializedRedeemTreasuryTokenAccount sets the "initializedRedeemTreasuryTokenAccount" account.
// Initialized redeem treasury token account with 0 tokens in supply, owner of account must be pda of program like above
func (inst *InitVault) SetInitializedRedeemTreasuryTokenAccount(initializedRedeemTreasuryTokenAccount ag_solanago.PublicKey) *InitVault {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(initializedRedeemTreasuryTokenAccount).WRITE()
	return inst
}

// GetInitializedRedeemTreasuryTokenAccount gets the "initializedRedeemTreasuryTokenAccount" account.
// Initialized redeem treasury token account with 0 tokens in supply, owner of account must be pda of program like above
func (inst *InitVault) GetInitializedRedeemTreasuryTokenAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[1]
}

// SetInitializedFractionTreasuryTokenAccount sets the "initializedFractionTreasuryTokenAccount" account.
// Initialized fraction treasury token account with 0 tokens in supply, owner of account must be pda of program like above
func (inst *InitVault) SetInitializedFractionTreasuryTokenAccount(initializedFractionTreasuryTokenAccount ag_solanago.PublicKey) *InitVault {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(initializedFractionTreasuryTokenAccount).WRITE()
	return inst
}

// GetInitializedFractionTreasuryTokenAccount gets the "initializedFractionTreasuryTokenAccount" account.
// Initialized fraction treasury token account with 0 tokens in supply, owner of account must be pda of program like above
func (inst *InitVault) GetInitializedFractionTreasuryTokenAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[2]
}

// SetUninitializedVaultAccount sets the "uninitializedVault" account.
// Uninitialized vault account
func (inst *InitVault) SetUninitializedVaultAccount(uninitializedVault ag_solanago.PublicKey) *InitVault {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(uninitializedVault).WRITE()
	return inst
}

// GetUninitializedVaultAccount gets the "uninitializedVault" account.
// Uninitialized vault account
func (inst *InitVault) GetUninitializedVaultAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[3]
}

// SetVaultAuthorityAccount sets the "vaultAuthority" account.
// Authority on the vault
func (inst *InitVault) SetVaultAuthorityAccount(vaultAuthority ag_solanago.PublicKey) *InitVault {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(vaultAuthority)
	return inst
}

// GetVaultAuthorityAccount gets the "vaultAuthority" account.
// Authority on the vault
func (inst *InitVault) GetVaultAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[4]
}

// SetPricingLookupAddressAccount sets the "pricingLookupAddress" account.
// Pricing Lookup Address
func (inst *InitVault) SetPricingLookupAddressAccount(pricingLookupAddress ag_solanago.PublicKey) *InitVault {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(pricingLookupAddress)
	return inst
}

// GetPricingLookupAddressAccount gets the "pricingLookupAddress" account.
// Pricing Lookup Address
func (inst *InitVault) GetPricingLookupAddressAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[5]
}

// SetTokenProgramAccount sets the "tokenProgram" account.
// Token program
func (inst *InitVault) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *InitVault {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
// Token program
func (inst *InitVault) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[6]
}

// SetRentSysvarAccount sets the "rentSysvar" account.
// Rent sysvar
func (inst *InitVault) SetRentSysvarAccount(rentSysvar ag_solanago.PublicKey) *InitVault {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(rentSysvar)
	return inst
}

// GetRentSysvarAccount gets the "rentSysvar" account.
// Rent sysvar
func (inst *InitVault) GetRentSysvarAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[7]
}

func (inst InitVault) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: ag_binary.TypeIDFromUint8(Instruction_InitVault),
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst InitVault) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *InitVault) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Args == nil {
			return errors.New("Args parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.InitializedFractionalShareMint is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.InitializedRedeemTreasuryTokenAccount is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.InitializedFractionTreasuryTokenAccount is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.UninitializedVault is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.VaultAuthority is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.PricingLookupAddress is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.RentSysvar is not set")
		}
	}
	return nil
}

func (inst *InitVault) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("InitVault")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Args", *inst.Args))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=8]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("  initializedFractionalShareMint", inst.AccountMetaSlice[0]))
						accountsBranch.Child(ag_format.Meta("  initializedRedeemTreasuryToken", inst.AccountMetaSlice[1]))
						accountsBranch.Child(ag_format.Meta("initializedFractionTreasuryToken", inst.AccountMetaSlice[2]))
						accountsBranch.Child(ag_format.Meta("              uninitializedVault", inst.AccountMetaSlice[3]))
						accountsBranch.Child(ag_format.Meta("                  vaultAuthority", inst.AccountMetaSlice[4]))
						accountsBranch.Child(ag_format.Meta("            pricingLookupAddress", inst.AccountMetaSlice[5]))
						accountsBranch.Child(ag_format.Meta("                    tokenProgram", inst.AccountMetaSlice[6]))
						accountsBranch.Child(ag_format.Meta("                      rentSysvar", inst.AccountMetaSlice[7]))
					})
				})
		})
}

func (obj InitVault) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Args` param:
	err = encoder.Encode(obj.Args)
	if err != nil {
		return err
	}
	return nil
}
func (obj *InitVault) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Args`:
	err = decoder.Decode(&obj.Args)
	if err != nil {
		return err
	}
	return nil
}

// NewInitVaultInstruction declares a new InitVault instruction with the provided parameters and accounts.
func NewInitVaultInstruction(
	// Parameters:
	args InitVaultArgs,
	// Accounts:
	initializedFractionalShareMint ag_solanago.PublicKey,
	initializedRedeemTreasuryTokenAccount ag_solanago.PublicKey,
	initializedFractionTreasuryTokenAccount ag_solanago.PublicKey,
	uninitializedVault ag_solanago.PublicKey,
	vaultAuthority ag_solanago.PublicKey,
	pricingLookupAddress ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
	rentSysvar ag_solanago.PublicKey) *InitVault {
	return NewInitVaultInstructionBuilder().
		SetArgs(args).
		SetInitializedFractionalShareMintAccount(initializedFractionalShareMint).
		SetInitializedRedeemTreasuryTokenAccount(initializedRedeemTreasuryTokenAccount).
		SetInitializedFractionTreasuryTokenAccount(initializedFractionTreasuryTokenAccount).
		SetUninitializedVaultAccount(uninitializedVault).
		SetVaultAuthorityAccount(vaultAuthority).
		SetPricingLookupAddressAccount(pricingLookupAddress).
		SetTokenProgramAccount(tokenProgram).
		SetRentSysvarAccount(rentSysvar)
}
