// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package token_metadata

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Create Metadata object.
type CreateMetadataAccount struct {
	Args *CreateMetadataAccountArgs

	// [0] = [WRITE] metadataKeyPDA
	// ··········· Metadata key (pda of ['metadata', program id, mint id])
	//
	// [1] = [] mintOfToken
	// ··········· Mint of token asset
	//
	// [2] = [SIGNER] mintAuthority
	// ··········· Mint authority
	//
	// [3] = [SIGNER] payer
	// ··········· payer
	//
	// [4] = [] updateAuthorityInfo
	// ··········· update authority info
	//
	// [5] = [] system
	// ··········· System program
	//
	// [6] = [] rent
	// ··········· Rent info
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewCreateMetadataAccountInstructionBuilder creates a new `CreateMetadataAccount` instruction builder.
func NewCreateMetadataAccountInstructionBuilder() *CreateMetadataAccount {
	nd := &CreateMetadataAccount{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 7),
	}
	return nd
}

// SetArgs sets the "args" parameter.
func (inst *CreateMetadataAccount) SetArgs(args CreateMetadataAccountArgs) *CreateMetadataAccount {
	inst.Args = &args
	return inst
}

// SetMetadataKeyPDAAccount sets the "metadataKeyPDA" account.
// Metadata key (pda of ['metadata', program id, mint id])
func (inst *CreateMetadataAccount) SetMetadataKeyPDAAccount(metadataKeyPDA ag_solanago.PublicKey) *CreateMetadataAccount {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(metadataKeyPDA).WRITE()
	return inst
}

// GetMetadataKeyPDAAccount gets the "metadataKeyPDA" account.
// Metadata key (pda of ['metadata', program id, mint id])
func (inst *CreateMetadataAccount) GetMetadataKeyPDAAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}

// SetMintOfTokenAccount sets the "mintOfToken" account.
// Mint of token asset
func (inst *CreateMetadataAccount) SetMintOfTokenAccount(mintOfToken ag_solanago.PublicKey) *CreateMetadataAccount {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(mintOfToken)
	return inst
}

// GetMintOfTokenAccount gets the "mintOfToken" account.
// Mint of token asset
func (inst *CreateMetadataAccount) GetMintOfTokenAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[1]
}

// SetMintAuthorityAccount sets the "mintAuthority" account.
// Mint authority
func (inst *CreateMetadataAccount) SetMintAuthorityAccount(mintAuthority ag_solanago.PublicKey) *CreateMetadataAccount {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(mintAuthority).SIGNER()
	return inst
}

// GetMintAuthorityAccount gets the "mintAuthority" account.
// Mint authority
func (inst *CreateMetadataAccount) GetMintAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[2]
}

// SetPayerAccount sets the "payer" account.
// payer
func (inst *CreateMetadataAccount) SetPayerAccount(payer ag_solanago.PublicKey) *CreateMetadataAccount {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(payer).SIGNER()
	return inst
}

// GetPayerAccount gets the "payer" account.
// payer
func (inst *CreateMetadataAccount) GetPayerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[3]
}

// SetUpdateAuthorityInfoAccount sets the "updateAuthorityInfo" account.
// update authority info
func (inst *CreateMetadataAccount) SetUpdateAuthorityInfoAccount(updateAuthorityInfo ag_solanago.PublicKey) *CreateMetadataAccount {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(updateAuthorityInfo)
	return inst
}

// GetUpdateAuthorityInfoAccount gets the "updateAuthorityInfo" account.
// update authority info
func (inst *CreateMetadataAccount) GetUpdateAuthorityInfoAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[4]
}

// SetSystemAccount sets the "system" account.
// System program
func (inst *CreateMetadataAccount) SetSystemAccount(system ag_solanago.PublicKey) *CreateMetadataAccount {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(system)
	return inst
}

// GetSystemAccount gets the "system" account.
// System program
func (inst *CreateMetadataAccount) GetSystemAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[5]
}

// SetRentAccount sets the "rent" account.
// Rent info
func (inst *CreateMetadataAccount) SetRentAccount(rent ag_solanago.PublicKey) *CreateMetadataAccount {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(rent)
	return inst
}

// GetRentAccount gets the "rent" account.
// Rent info
func (inst *CreateMetadataAccount) GetRentAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[6]
}

func (inst CreateMetadataAccount) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: ag_binary.TypeIDFromUint8(Instruction_CreateMetadataAccount),
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst CreateMetadataAccount) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *CreateMetadataAccount) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Args == nil {
			return errors.New("Args parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.MetadataKeyPDA is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.MintOfToken is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.MintAuthority is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.Payer is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.UpdateAuthorityInfo is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.System is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.Rent is not set")
		}
	}
	return nil
}

func (inst *CreateMetadataAccount) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("CreateMetadataAccount")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Args", *inst.Args))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=7]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("     metadataKeyPDA", inst.AccountMetaSlice[0]))
						accountsBranch.Child(ag_format.Meta("        mintOfToken", inst.AccountMetaSlice[1]))
						accountsBranch.Child(ag_format.Meta("      mintAuthority", inst.AccountMetaSlice[2]))
						accountsBranch.Child(ag_format.Meta("              payer", inst.AccountMetaSlice[3]))
						accountsBranch.Child(ag_format.Meta("updateAuthorityInfo", inst.AccountMetaSlice[4]))
						accountsBranch.Child(ag_format.Meta("             system", inst.AccountMetaSlice[5]))
						accountsBranch.Child(ag_format.Meta("               rent", inst.AccountMetaSlice[6]))
					})
				})
		})
}

func (obj CreateMetadataAccount) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Args` param:
	err = encoder.Encode(obj.Args)
	if err != nil {
		return err
	}
	return nil
}
func (obj *CreateMetadataAccount) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Args`:
	err = decoder.Decode(&obj.Args)
	if err != nil {
		return err
	}
	return nil
}

// NewCreateMetadataAccountInstruction declares a new CreateMetadataAccount instruction with the provided parameters and accounts.
func NewCreateMetadataAccountInstruction(
	// Parameters:
	args CreateMetadataAccountArgs,
	// Accounts:
	metadataKeyPDA ag_solanago.PublicKey,
	mintOfToken ag_solanago.PublicKey,
	mintAuthority ag_solanago.PublicKey,
	payer ag_solanago.PublicKey,
	updateAuthorityInfo ag_solanago.PublicKey,
	system ag_solanago.PublicKey,
	rent ag_solanago.PublicKey) *CreateMetadataAccount {
	return NewCreateMetadataAccountInstructionBuilder().
		SetArgs(args).
		SetMetadataKeyPDAAccount(metadataKeyPDA).
		SetMintOfTokenAccount(mintOfToken).
		SetMintAuthorityAccount(mintAuthority).
		SetPayerAccount(payer).
		SetUpdateAuthorityInfoAccount(updateAuthorityInfo).
		SetSystemAccount(system).
		SetRentAccount(rent)
}
