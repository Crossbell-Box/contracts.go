// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// DataTypesCharacter is an auto generated low-level Go binding around an user-defined struct.
type DataTypesCharacter struct {
	CharacterId *big.Int
	Handle      string
	Uri         string
	NoteCount   *big.Int
	SocialToken common.Address
	LinkModule  common.Address
}

// DataTypesCreateCharacterData is an auto generated low-level Go binding around an user-defined struct.
type DataTypesCreateCharacterData struct {
	To                 common.Address
	Handle             string
	Uri                string
	LinkModule         common.Address
	LinkModuleInitData []byte
}

// DataTypesERC721Struct is an auto generated low-level Go binding around an user-defined struct.
type DataTypesERC721Struct struct {
	TokenAddress  common.Address
	Erc721TokenId *big.Int
}

// DataTypesMintNoteData is an auto generated low-level Go binding around an user-defined struct.
type DataTypesMintNoteData struct {
	CharacterId    *big.Int
	NoteId         *big.Int
	To             common.Address
	MintModuleData []byte
}

// DataTypesNote is an auto generated low-level Go binding around an user-defined struct.
type DataTypesNote struct {
	LinkItemType [32]byte
	LinkKey      [32]byte
	ContentUri   string
	LinkModule   common.Address
	MintModule   common.Address
	MintNFT      common.Address
	Deleted      bool
	Locked       bool
}

// DataTypesNoteStruct is an auto generated low-level Go binding around an user-defined struct.
type DataTypesNoteStruct struct {
	CharacterId *big.Int
	NoteId      *big.Int
}

// DataTypesPostNoteData is an auto generated low-level Go binding around an user-defined struct.
type DataTypesPostNoteData struct {
	CharacterId        *big.Int
	ContentUri         string
	LinkModule         common.Address
	LinkModuleInitData []byte
	MintModule         common.Address
	MintModuleInitData []byte
	Locked             bool
}

// DataTypescreateThenLinkCharacterData is an auto generated low-level Go binding around an user-defined struct.
type DataTypescreateThenLinkCharacterData struct {
	FromCharacterId *big.Int
	To              common.Address
	LinkType        [32]byte
}

// DataTypeslinkAddressData is an auto generated low-level Go binding around an user-defined struct.
type DataTypeslinkAddressData struct {
	FromCharacterId *big.Int
	EthAddress      common.Address
	LinkType        [32]byte
	Data            []byte
}

// DataTypeslinkAnyUriData is an auto generated low-level Go binding around an user-defined struct.
type DataTypeslinkAnyUriData struct {
	FromCharacterId *big.Int
	ToUri           string
	LinkType        [32]byte
	Data            []byte
}

// DataTypeslinkCharacterData is an auto generated low-level Go binding around an user-defined struct.
type DataTypeslinkCharacterData struct {
	FromCharacterId *big.Int
	ToCharacterId   *big.Int
	LinkType        [32]byte
	Data            []byte
}

// DataTypeslinkERC721Data is an auto generated low-level Go binding around an user-defined struct.
type DataTypeslinkERC721Data struct {
	FromCharacterId *big.Int
	TokenAddress    common.Address
	TokenId         *big.Int
	LinkType        [32]byte
	Data            []byte
}

// DataTypeslinkLinklistData is an auto generated low-level Go binding around an user-defined struct.
type DataTypeslinkLinklistData struct {
	FromCharacterId *big.Int
	ToLinkListId    *big.Int
	LinkType        [32]byte
	Data            []byte
}

// DataTypeslinkNoteData is an auto generated low-level Go binding around an user-defined struct.
type DataTypeslinkNoteData struct {
	FromCharacterId *big.Int
	ToCharacterId   *big.Int
	ToNoteId        *big.Int
	LinkType        [32]byte
	Data            []byte
}

// DataTypessetLinkModule4AddressData is an auto generated low-level Go binding around an user-defined struct.
type DataTypessetLinkModule4AddressData struct {
	Account            common.Address
	LinkModule         common.Address
	LinkModuleInitData []byte
}

// DataTypessetLinkModule4LinklistData is an auto generated low-level Go binding around an user-defined struct.
type DataTypessetLinkModule4LinklistData struct {
	LinklistId         *big.Int
	LinkModule         common.Address
	LinkModuleInitData []byte
}

// DataTypessetMintModule4NoteData is an auto generated low-level Go binding around an user-defined struct.
type DataTypessetMintModule4NoteData struct {
	CharacterId        *big.Int
	NoteId             *big.Int
	MintModule         common.Address
	MintModuleInitData []byte
}

// DataTypesunlinkAddressData is an auto generated low-level Go binding around an user-defined struct.
type DataTypesunlinkAddressData struct {
	FromCharacterId *big.Int
	EthAddress      common.Address
	LinkType        [32]byte
}

// DataTypesunlinkAnyUriData is an auto generated low-level Go binding around an user-defined struct.
type DataTypesunlinkAnyUriData struct {
	FromCharacterId *big.Int
	ToUri           string
	LinkType        [32]byte
}

// DataTypesunlinkCharacterData is an auto generated low-level Go binding around an user-defined struct.
type DataTypesunlinkCharacterData struct {
	FromCharacterId *big.Int
	ToCharacterId   *big.Int
	LinkType        [32]byte
}

// DataTypesunlinkERC721Data is an auto generated low-level Go binding around an user-defined struct.
type DataTypesunlinkERC721Data struct {
	FromCharacterId *big.Int
	TokenAddress    common.Address
	TokenId         *big.Int
	LinkType        [32]byte
}

// DataTypesunlinkLinklistData is an auto generated low-level Go binding around an user-defined struct.
type DataTypesunlinkLinklistData struct {
	FromCharacterId *big.Int
	ToLinkListId    *big.Int
	LinkType        [32]byte
}

// DataTypesunlinkNoteData is an auto generated low-level Go binding around an user-defined struct.
type DataTypesunlinkNoteData struct {
	FromCharacterId *big.Int
	ToCharacterId   *big.Int
	ToNoteId        *big.Int
	LinkType        [32]byte
}

// ContractsMetaData contains all meta data concerning the Contracts contract.
var ContractsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"}],\"name\":\"ErrCharacterNotExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrHandleContainsInvalidCharacters\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrHandleExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrHandleLengthInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrNotAddressOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrNotCharacterOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrNotEnoughPermission\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrNotEnoughPermissionForThisNote\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrNoteIsDeleted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrNoteLocked\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrNoteNotExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrSocialTokenExists\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"handle\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"linkModuleInitData\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.CreateCharacterData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"createCharacter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"}],\"internalType\":\"structDataTypes.createThenLinkCharacterData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"createThenLinkCharacter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"noteId\",\"type\":\"uint256\"}],\"name\":\"deleteNote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"}],\"name\":\"getCharacter\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"handle\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"noteCount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"socialToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"}],\"internalType\":\"structDataTypes.Character\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"handle\",\"type\":\"string\"}],\"name\":\"getCharacterByHandle\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"handle\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"noteCount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"socialToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"}],\"internalType\":\"structDataTypes.Character\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"}],\"name\":\"getCharacterUri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"}],\"name\":\"getHandle\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getLinkModule4Address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getLinkModule4ERC721\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getLinkModule4Linklist\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLinklistContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"}],\"name\":\"getLinklistId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"linkListId\",\"type\":\"uint256\"}],\"name\":\"getLinklistType\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getLinklistUri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"noteId\",\"type\":\"uint256\"}],\"name\":\"getNote\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"linkItemType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"linkKey\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"contentUri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"mintModule\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"mintNFT\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"deleted\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"locked\",\"type\":\"bool\"}],\"internalType\":\"structDataTypes.Note\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"getOperatorPermissions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"}],\"name\":\"getOperators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"noteId\",\"type\":\"uint256\"}],\"name\":\"getOperators4Note\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"blocklist\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"allowlist\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getPrimaryCharacterId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRevision\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"permissionBitMap\",\"type\":\"uint256\"}],\"name\":\"grantOperatorPermissions\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"noteId\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"blocklist\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"allowlist\",\"type\":\"address[]\"}],\"name\":\"grantOperators4Note\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"linklist_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"mintNFTImpl_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"periphery_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newbieVilla_\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"noteId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isOperatorAllowedForNote\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"}],\"name\":\"isPrimaryCharacter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"ethAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.linkAddressData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"linkAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"toUri\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.linkAnyUriData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"linkAnyUri\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.linkCharacterData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"linkCharacter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.linkERC721Data\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"linkERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toLinkListId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.linkLinklistData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"linkLinklist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toNoteId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.linkNoteData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"linkNote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"noteId\",\"type\":\"uint256\"}],\"name\":\"lockNote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"noteId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"mintModuleData\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.MintNoteData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"mintNote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"contentUri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"linkModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"mintModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"mintModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"locked\",\"type\":\"bool\"}],\"internalType\":\"structDataTypes.PostNoteData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"postNote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"noteId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"contentUri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"linkModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"mintModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"mintModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"locked\",\"type\":\"bool\"}],\"internalType\":\"structDataTypes.PostNoteData\",\"name\":\"vars\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"ethAddress\",\"type\":\"address\"}],\"name\":\"postNote4Address\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"contentUri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"linkModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"mintModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"mintModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"locked\",\"type\":\"bool\"}],\"internalType\":\"structDataTypes.PostNoteData\",\"name\":\"vars\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"postNote4AnyUri\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"contentUri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"linkModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"mintModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"mintModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"locked\",\"type\":\"bool\"}],\"internalType\":\"structDataTypes.PostNoteData\",\"name\":\"vars\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"toCharacterId\",\"type\":\"uint256\"}],\"name\":\"postNote4Character\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"contentUri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"linkModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"mintModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"mintModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"locked\",\"type\":\"bool\"}],\"internalType\":\"structDataTypes.PostNoteData\",\"name\":\"vars\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc721TokenId\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.ERC721Struct\",\"name\":\"erc721\",\"type\":\"tuple\"}],\"name\":\"postNote4ERC721\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"contentUri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"linkModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"mintModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"mintModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"locked\",\"type\":\"bool\"}],\"internalType\":\"structDataTypes.PostNoteData\",\"name\":\"vars\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"toLinklistId\",\"type\":\"uint256\"}],\"name\":\"postNote4Linklist\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"contentUri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"linkModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"mintModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"mintModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"locked\",\"type\":\"bool\"}],\"internalType\":\"structDataTypes.PostNoteData\",\"name\":\"vars\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"noteId\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.NoteStruct\",\"name\":\"note\",\"type\":\"tuple\"}],\"name\":\"postNote4Note\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resolver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"newUri\",\"type\":\"string\"}],\"name\":\"setCharacterUri\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"newHandle\",\"type\":\"string\"}],\"name\":\"setHandle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"linkModuleInitData\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.setLinkModule4AddressData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"setLinkModule4Address\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"linklistId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"linkModuleInitData\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.setLinkModule4LinklistData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"setLinkModule4Linklist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"linklistId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"setLinklistUri\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"noteId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"mintModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"mintModuleInitData\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.setMintModule4NoteData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"setMintModule4Note\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"noteId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"newUri\",\"type\":\"string\"}],\"name\":\"setNoteUri\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"}],\"name\":\"setPrimaryCharacterId\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"setSocialToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"ethAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"}],\"internalType\":\"structDataTypes.unlinkAddressData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"unlinkAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"toUri\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"}],\"internalType\":\"structDataTypes.unlinkAnyUriData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"unlinkAnyUri\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"}],\"internalType\":\"structDataTypes.unlinkCharacterData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"unlinkCharacter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"}],\"internalType\":\"structDataTypes.unlinkERC721Data\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"unlinkERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toLinkListId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"}],\"internalType\":\"structDataTypes.unlinkLinklistData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"unlinkLinklist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toNoteId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"}],\"internalType\":\"structDataTypes.unlinkNoteData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"unlinkNote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50615b5980620000216000396000f3fe608060405234801561001057600080fd5b506004361061043e5760003560e01c8063867884e611610236578063c053f6b81161013b578063dc17b6de116100c3578063ef0828ab11610087578063ef0828ab14610a71578063f2ad807514610a84578063f316bacd14610a97578063f6479d7714610aaa578063fe9299fb14610abd57600080fd5b8063dc17b6de146109e9578063dca27135146109fc578063e56f2fe414610a0f578063e985e9c514610a22578063ec81d19414610a5e57600080fd5b8063cd69fe611161010a578063cd69fe6114610965578063d23b320b14610978578063d70e10c61461098b578063dabb0531146109b6578063db491e80146109c957600080fd5b8063c053f6b81461091b578063c2a6fe3b1461092c578063c87b56dd1461093f578063cb8e757e1461095257600080fd5b80639a4dec18116101be578063a7ccb4bf1161018d578063a7ccb4bf146108af578063ac9650d8146108c2578063af90b112146108e2578063b88d4fde146108f5578063b9d328451461090857600080fd5b80639a4dec18146108565780639a50248d14610869578063a22cb46514610889578063a6e6178d1461089c57600080fd5b806393f057e51161020557806393f057e5146107df578063952be0ef146107f257806395d89b411461082857806395d9fa7d146108305780639864c3071461084357600080fd5b8063867884e6146107935780638734bbfc146107a65780638b4ca06a146107b957806392f7070b146107cc57600080fd5b80632f745c591161034757806347f94de7116102cf578063628b644a11610293578063628b644a146107275780636352211e1461073a5780636bf55d5f1461074d57806370a082311461076d57806374f345cf1461078057600080fd5b806347f94de7146106ba57806349186953146106cd5780634f6ccce7146106ee5780635a936d10146107015780635fb881831461071457600080fd5b8063388f508311610316578063388f50831461065b57806340ad34d81461066e57806342842e0e1461068157806342966c681461069457806344b82a24146106a757600080fd5b80632f745c59146105f657806331b9d08c14610609578063327b2a031461063557806333f06ee61461064857600080fd5b8063144a3e83116103ca5780632209d145116103995780632209d1451461055e57806323b872dd1461059457806328fbb805146105a757806329c301c2146105ba5780632abc6bf6146105cd57600080fd5b8063144a3e831461051d57806318160ddd14610530578063188b04b314610538578063206657f21461054b57600080fd5b806308cb68ff1161041157806308cb68ff146104be578063095ea7b3146104d35780630c4dd5f2146104e65780630ff98244146104f95780631316529d1461050c57600080fd5b806301ffc9a71461044357806304f3bcec1461046b57806306fdde0314610496578063081812fc146104ab575b600080fd5b6104566104513660046145e1565b610ae6565b60405190151581526020015b60405180910390f35b60175461047e906001600160a01b031681565b6040516001600160a01b039091168152602001610462565b61049e610b11565b604051610462919061464e565b61047e6104b9366004614661565b610ba3565b6104d16104cc366004614692565b610c3d565b005b6104d16104e13660046146e2565b610d14565b6104d16104f4366004614692565b610e29565b6104d161050736600461470c565b610f33565b60045b604051908152602001610462565b61049e61052b366004614661565b610fac565b60085461050f565b6104d161054636600461473a565b610fb7565b6104d161055936600461476e565b611055565b61047e61056c3660046146e2565b6001600160a01b03918216600090815260106020908152604080832093835292905220541690565b6104d16105a23660046147a3565b6110ed565b6104566105b53660046147cf565b61111e565b61050f6105c8366004614816565b611135565b61050f6105db36600461484a565b6001600160a01b03166000908152600c602052604090205490565b61050f6106043660046146e2565b6111c6565b61047e61061736600461484a565b6001600160a01b039081166000908152601160205260409020541690565b61050f610643366004614877565b61125c565b6104d1610656366004614904565b6113eb565b6104d161066936600461473a565b6114d2565b6104d161067c36600461494f565b611557565b6104d161068f3660046147a3565b6115d3565b6104d16106a2366004614661565b6115ee565b61050f6106b536600461496b565b61168b565b6104d16106c8366004614904565b6116ec565b6106e06106db3660046149af565b611753565b604051610462929190614a15565b61050f6106fc366004614661565b6117ab565b6104d161070f36600461470c565b61183e565b6104d161072236600461473a565b6118b7565b6104d1610735366004614a3a565b611925565b61047e610748366004614661565b6119b2565b61076061075b366004614661565b611a29565b6040516104629190614a8c565b61050f61077b36600461484a565b611a43565b6104d161078e3660046149af565b611aca565b6104d16107a136600461494f565b611b44565b6104566107b4366004614661565b611bee565b61050f6107c7366004614661565b611c1c565b61050f6107da366004614a9f565b611c88565b6104d16107ed36600461470c565b611d03565b61050f610800366004614ae3565b60009182526019602090815260408084206001600160a01b0393909316845291905290205490565b61049e611da3565b6104d161083e366004614ae3565b611db2565b6104d161085136600461473a565b611e44565b61050f610864366004614877565b611ea5565b61087c610877366004614b06565b611f90565b6040516104629190614b47565b6104d1610897366004614bd3565b61213f565b6104d16108aa366004614904565b61214a565b61050f6108bd36600461473a565b6121f5565b6108d56108d0366004614c41565b612296565b6040516104629190614c76565b61050f6108f036600461496b565b61238a565b6104d1610903366004614dc3565b6123b1565b6104d1610916366004614e3c565b6123e9565b6013546001600160a01b031661047e565b6104d161093a3660046149af565b61249c565b61049e61094d366004614661565b61250e565b6104d1610960366004614e3c565b6125b3565b61050f610973366004614e3c565b612642565b6104d161098636600461473a565b612657565b61050f6109993660046149af565b6000918252600d6020908152604080842092845291905290205490565b61087c6109c4366004614661565b6126e2565b6109dc6109d73660046149af565b612869565b6040516104629190614e70565b6104d16109f7366004614f06565b6129ca565b61049e610a0a366004614661565b612a5b565b6104d1610a1d366004614f88565b612acd565b610456610a30366004615039565b6001600160a01b03918216600090815260056020908152604080832093909416825291909152205460ff1690565b61049e610a6c366004614661565b612c5a565b6104d1610a7f366004614692565b612c7a565b6104d1610a92366004614661565b612d04565b61050f610aa5366004615063565b612d53565b61050f610ab836600461470c565b612e94565b61047e610acb366004614661565b6000908152600f60205260409020546001600160a01b031690565b60006001600160e01b0319821663780e9d6360e01b1480610b0b5750610b0b82612faf565b92915050565b606060008054610b20906150be565b80601f0160208091040260200160405190810160405280929190818152602001828054610b4c906150be565b8015610b995780601f10610b6e57610100808354040283529160200191610b99565b820191906000526020600020905b815481529060010190602001808311610b7c57829003601f168201915b5050505050905090565b6000818152600260205260408120546001600160a01b0316610c215760405162461bcd60e51b815260206004820152602c60248201527f4552433732313a20617070726f76656420717565727920666f72206e6f6e657860448201526b34b9ba32b73a103a37b5b2b760a11b60648201526084015b60405180910390fd5b506000908152600460205260409020546001600160a01b031690565b610c4a602082018261484a565b6001600160a01b0316336001600160a01b031614610c7b5760405163ca67421960e01b815260040160405180910390fd5b73__$69db18565ecbefece480e92cc5f8fb1274$__63dfc34f25610ca2602084018461484a565b610cb2604085016020860161484a565b610cbf60408601866150f2565b60116040518663ffffffff1660e01b8152600401610ce1959493929190615161565b60006040518083038186803b158015610cf957600080fd5b505af4158015610d0d573d6000803e3d6000fd5b5050505050565b6000610d1f826119b2565b9050806001600160a01b0316836001600160a01b031603610d8c5760405162461bcd60e51b815260206004820152602160248201527f4552433732313a20617070726f76616c20746f2063757272656e74206f776e656044820152603960f91b6064820152608401610c18565b336001600160a01b0382161480610da85750610da88133610a30565b610e1a5760405162461bcd60e51b815260206004820152603860248201527f4552433732313a20617070726f76652063616c6c6572206973206e6f74206f7760448201527f6e6572206e6f7220617070726f76656420666f7220616c6c00000000000000006064820152608401610c18565b610e248383612fff565b505050565b6013546040516367880d6160e11b8152823560048201526000916001600160a01b03169063cf101ac290602401602060405180830381865afa158015610e73573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e9791906151a0565b9050610ea48160c161306d565b73__$69db18565ecbefece480e92cc5f8fb1274$__636252159e8335610ed0604086016020870161484a565b610edd60408701876150f2565b600f6040518663ffffffff1660e01b8152600401610eff9594939291906151b9565b60006040518083038186803b158015610f1757600080fd5b505af4158015610f2b573d6000803e3d6000fd5b505050505050565b610f3f813560b261306d565b60135481356000818152600d6020908152604080832081870135808552908352928190205490516337fb824760e11b815273__$e2c9a8f399964af47bb173600dd9e5f662$__95636ff7048e95610ce19590948901359390926001600160a01b03909216916004016151e4565b6060610b0b8261250e565b610fc3813560b261306d565b610fd081602001356130fe565b73__$e2c9a8f399964af47bb173600dd9e5f662$__639ec52a2382356020840135604085013561100360608701876150f2565b6013546020808a01356000908152600a9091526040908190206005015490516001600160e01b031960e08a901b168152610ce19796959493926001600160a01b03908116921690600d90600401615210565b61106083600261306d565b604051631f8c0b6760e11b8152600481018490526001600160a01b038316602482015260448101829052601860648201526019608482015273__$bfdcc7011a2e80c167e18378382b8de19c$__90633f1816ce9060a4015b60006040518083038186803b1580156110d057600080fd5b505af41580156110e4573d6000803e3d6000fd5b50505050505050565b6110f73382613139565b6111135760405162461bcd60e51b8152600401610c189061525d565b610e24838383613230565b600061112b8484846133dd565b90505b9392505050565b6000611143823560ec61306d565b61114d8235613467565b6040516342a34a5360e01b815290915073__$2873850e5b3f5697aafc2f5cd9f27cc884$__906342a34a539061119190859085906000908190600e906004016153b5565b60006040518083038186803b1580156111a957600080fd5b505af41580156111bd573d6000803e3d6000fd5b50505050919050565b60006111d183611a43565b82106112335760405162461bcd60e51b815260206004820152602b60248201527f455243373231456e756d657261626c653a206f776e657220696e646578206f7560448201526a74206f6620626f756e647360a81b6064820152608401610c18565b506001600160a01b03919091166000908152600660209081526040808320938352929052205490565b600061126a833560ca61306d565b6545524337323160d01b60006112808535613467565b6013549091506000906001600160a01b0316632ea24efc826112a5602089018961484a565b6040516001600160e01b031960e085901b16815260048101929092526001600160a01b03166024820152602088013560448201526064016020604051808303816000875af11580156112fb573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061131f91906151a0565b905073__$2873850e5b3f5697aafc2f5cd9f27cc884$__6342a34a538784868561134c60208c018c61484a565b8b6020013560405160200161137f92919060609290921b6bffffffffffffffffffffffff19168252601482015260340190565b604051602081830303815290604052600e6040518763ffffffff1660e01b81526004016113b1969594939291906153fa565b60006040518083038186803b1580156113c957600080fd5b505af41580156113dd573d6000803e3d6000fd5b509398975050505050505050565b6013546040516367880d6160e11b8152600481018590526000916001600160a01b03169063cf101ac290602401602060405180830381865afa158015611435573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061145991906151a0565b90506114668160b161306d565b601354604051633c17845760e11b81526001600160a01b039091169063782f08ae9061149a90879087908790600401615445565b600060405180830381600087803b1580156114b457600080fd5b505af11580156114c8573d6000803e3d6000fd5b5050505050505050565b6114de813560b961306d565b73__$e2c9a8f399964af47bb173600dd9e5f662$__63a4159c6b823561150a604085016020860161484a565b6013546040805160e086901b6001600160e01b031916815260048101949094526001600160a01b0392831660248501528601356044840152166064820152600d608482015260a401610ce1565b611563813560b661306d565b60135460408051631d4deabf60e01b81528335600482015260208401356024820152908301356044820152606083013560648201526001600160a01b039091166084820152600d60a482015273__$e2c9a8f399964af47bb173600dd9e5f662$__90631d4deabf9060c401610ce1565b610e24838383604051806020016040528060008152506123b1565b6000818152600a6020526040808220905161160c916001019061545f565b60408051918290039091206000818152600b6020908152838220829055858252600a9052918220828155909250906116476001830182614530565b611655600283016000614530565b50600060038201556004810180546001600160a01b031990811690915560059091018054909116905561168782613490565b5050565b6000611699833560c861306d565b67131a5b9adb1a5cdd60c21b60006116b18535613467565b905060008460001b905073__$2873850e5b3f5697aafc2f5cd9f27cc884$__6342a34a53878486858a60405160200161137f91815260200190565b6116f78360b061306d565b6000838152600a6020526040902060020161171382848361551b565b50827f17d7c9f69270ba135480ef16837f38b9d37d3ab291cbd3ba03982290c663199783836040516117469291906155da565b60405180910390a2505050565b6000828152601a602090815260408083208484529091529020606090819061177a906134ef565b6000858152601a6020908152604080832087845290915290209092506117a2906002016134ef565b90509250929050565b60006117b660085490565b82106118195760405162461bcd60e51b815260206004820152602c60248201527f455243373231456e756d657261626c653a20676c6f62616c20696e646578206f60448201526b7574206f6620626f756e647360a01b6064820152608401610c18565b6008828154811061182c5761182c6155ee565b90600052602060002001549050919050565b61184a813560be61306d565b60135481356000818152600d602090815260408083208187013580855290835292819020549051633fe4fe3960e11b815273__$e2c9a8f399964af47bb173600dd9e5f662$__95637fc9fc7295610ce19590948901359390926001600160a01b03909216916004016151e4565b6118c3813560bb61306d565b73__$e2c9a8f399964af47bb173600dd9e5f662$__6348391dcb82356118ec60208501856150f2565b601354604080516001600160e01b031960e088901b168152610ce19594939291890135916001600160a01b031690600d90600401615604565b61192f84846134fc565b6119398484613561565b61194384846135d8565b6040516001626802bf60e01b0319815273__$2873850e5b3f5697aafc2f5cd9f27cc884$__9063ff97fd4190611986908790879087908790600e90600401615645565b60006040518083038186803b15801561199e57600080fd5b505af41580156114c8573d6000803e3d6000fd5b6000818152600260205260408120546001600160a01b031680610b0b5760405162461bcd60e51b815260206004820152602960248201527f4552433732313a206f776e657220717565727920666f72206e6f6e657869737460448201526832b73a103a37b5b2b760b91b6064820152608401610c18565b6000818152601860205260409020606090610b0b906134ef565b60006001600160a01b038216611aae5760405162461bcd60e51b815260206004820152602a60248201527f4552433732313a2062616c616e636520717565727920666f7220746865207a65604482015269726f206164647265737360b01b6064820152608401610c18565b506001600160a01b031660009081526003602052604090205490565b611ad58260c461306d565b611adf8282613561565b6000828152600e60209081526040808320848452825291829020600501805460ff60a81b1916600160a81b179055905182815283917f036469f3e73c83520cdefa197d7a9c854c2f8bc0164b82e9f2bd4aa7e150fd3091015b60405180910390a25050565b611b50813560b861306d565b73__$e2c9a8f399964af47bb173600dd9e5f662$__63154246368235611b7c604085016020860161484a565b60135485356000908152600d6020908152604080832060608a013580855292529182902054825160e088901b6001600160e01b031916815260048101969096526001600160a01b03948516602487015291880135604486015260648501529116608483015260a482015260c401610ce1565b600080611bfa836119b2565b6001600160a01b03166000908152600c60205260409020549290921492915050565b60135460405162fba02760e01b8152600481018390526000916001600160a01b03169062fba02790602401602060405180830381865afa158015611c64573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b0b91906151a0565b6000611c96833560c761306d565b664164647265737360c81b6000611cad8535613467565b6040516bffffffffffffffffffffffff19606087901b1660208201529091506001600160a01b0385169073__$2873850e5b3f5697aafc2f5cd9f27cc884$__906342a34a5390889085908790869060340161137f565b611d0f813560ba61306d565b73__$e2c9a8f399964af47bb173600dd9e5f662$__6393c96a528235611d3b604085016020860161484a565b60135485356000908152600d60209081526040808320818a013580855292529182902054915160e087901b6001600160e01b031916815260048101959095526001600160a01b039384166024860152604485015291166064830152608482015260a401610ce1565b606060018054610b20906150be565b611dbd82600161306d565b6000828152600a60205260409020600401546001600160a01b031615611df65760405163fe6f50e560e01b815260040160405180910390fd5b6040516384b44a2f60e01b8152600481018390526001600160a01b0382166024820152600a604482015273__$d8529f0216812a2dec6b96adf1943c3537$__906384b44a2f90606401610eff565b611e50813560bd61306d565b60135460408051632ca904df60e01b815273__$e2c9a8f399964af47bb173600dd9e5f662$__92632ca904df92610ce192863592602088013592880135916001600160a01b0390911690600d906004016151e4565b6000611eb3833560c961306d565b634e6f746560e01b6000611ec78535613467565b601354604051635cb46be760e01b815260006004820181905287356024830152602088013560448301529293506001600160a01b0390911690635cb46be7906064016020604051808303816000875af1158015611f28573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611f4c91906151a0565b905073__$2873850e5b3f5697aafc2f5cd9f27cc884$__6342a34a53878486858a600001358b6020013560405160200161137f929190918252602082015260400190565b611f9861456a565b60008383604051611faa929190615665565b604080519182900382206000818152600b602090815283822054808352600a82529184902060c0860190945283548552600184018054939650919493929084019190611ff5906150be565b80601f0160208091040260200160405190810160405280929190818152602001828054612021906150be565b801561206e5780601f106120435761010080835404028352916020019161206e565b820191906000526020600020905b81548152906001019060200180831161205157829003601f168201915b50505050508152602001600282018054612087906150be565b80601f01602080910402602001604051908101604052809291908181526020018280546120b3906150be565b80156121005780601f106120d557610100808354040283529160200191612100565b820191906000526020600020905b8154815290600101906020018083116120e357829003601f168201915b50505091835250506003820154602082015260048201546001600160a01b03908116604083015260059092015490911660609091015295945050505050565b61168733838361361d565b61215583600061306d565b6121758282604051612168929190615665565b60405180910390206136eb565b6121b482828080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061371892505050565b60405163130f361d60e01b815273__$d8529f0216812a2dec6b96adf1943c3537$__9063130f361d906110b890869086908690600b90600a90600401615675565b600061220682356020840135613561565b73__$2873850e5b3f5697aafc2f5cd9f27cc884$__639d2e06f083356020850135612237606087016040880161484a565b61224460608801886150f2565b6014546040516001600160e01b031960e089901b1681526122799695949392916001600160a01b031690600e906004016156a3565b602060405180830381865af4158015611c64573d6000803e3d6000fd5b6060816001600160401b038111156122b0576122b0614cd8565b6040519080825280602002602001820160405280156122e357816020015b60608152602001906001900390816122ce5790505b50905060005b828110156123835761235330858584818110612307576123076155ee565b905060200281019061231991906150f2565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061378592505050565b828281518110612365576123656155ee565b6020026020010181905250808061237b906156fd565b9150506122e9565b5092915050565b6000612398833560c661306d565b6821b430b930b1ba32b960b91b60006116b18535613467565b6123bb3383613139565b6123d75760405162461bcd60e51b8152600401610c189061525d565b6123e3848484846137aa565b50505050565b6123f5813560b561306d565b61240781602001358260400135613561565b73__$e2c9a8f399964af47bb173600dd9e5f662$__6371bd9b06823560208401356040850135606086013561243f60808801886150f2565b6013546020808b01356000908152600e82526040808220818e013583529092528190206003015490516001600160e01b031960e08b901b168152610ce1989796959493926001600160a01b03908116921690600d90600401615716565b6124a78260c561306d565b6124b18282613561565b6000828152600e60209081526040808320848452825291829020600501805460ff60a01b1916600160a01b179055905182815283917f4f1db9708b537c1d26a7af4b235fd079bf2342d92a276e27eb6c8717e8bbcf939101611b38565b6000818152600a6020526040902060020180546060919061252e906150be565b80601f016020809104026020016040519081016040528092919081815260200182805461255a906150be565b80156125a75780601f1061257c576101008083540402835291602001916125a7565b820191906000526020600020905b81548152906001019060200180831161258a57829003601f168201915b50505050509050919050565b6125bf813560b761306d565b73__$e2c9a8f399964af47bb173600dd9e5f662$__63f35deae182356125eb604085016020860161484a565b6013546040805160e086901b6001600160e01b031916815260048101949094526001600160a01b039283166024850152860135604484015260608601356064840152166084820152600d60a482015260c401610ce1565b6000610b0b6126508361576c565b60016137dd565b612663813560c261306d565b61267281356020830135613561565b612681813560208301356135d8565b73__$69db18565ecbefece480e92cc5f8fb1274$__6320828a02823560208401356126b2606086016040870161484a565b6126bf60608701876150f2565b600e6040518763ffffffff1660e01b8152600401610ce19695949392919061581f565b6126ea61456a565b600a60008381526020019081526020016000206040518060c001604052908160008201548152602001600182018054612722906150be565b80601f016020809104026020016040519081016040528092919081815260200182805461274e906150be565b801561279b5780601f106127705761010080835404028352916020019161279b565b820191906000526020600020905b81548152906001019060200180831161277e57829003601f168201915b505050505081526020016002820180546127b4906150be565b80601f01602080910402602001604051908101604052809291908181526020018280546127e0906150be565b801561282d5780601f106128025761010080835404028352916020019161282d565b820191906000526020600020905b81548152906001019060200180831161281057829003601f168201915b50505091835250506003820154602082015260048201546001600160a01b03908116604083015260059092015490911660609091015292915050565b60408051610100808201835260008083526020808401829052606084860181905284018290526080840182905260a0840182905260c0840182905260e08401829052868252600e815284822086835281529084902084519283018552805483526001810154918301919091526002810180549394929391928401916128ed906150be565b80601f0160208091040260200160405190810160405280929190818152602001828054612919906150be565b80156129665780601f1061293b57610100808354040283529160200191612966565b820191906000526020600020905b81548152906001019060200180831161294957829003601f168201915b505050918352505060038201546001600160a01b039081166020830152600483015481166040830152600590920154918216606082015260ff600160a01b8304811615156080830152600160a81b909204909116151560a090910152905092915050565b6129d586600361306d565b6129df8686613561565b604051630afb883f60e41b815273__$bfdcc7011a2e80c167e18378382b8de19c$__9063afb883f090612a2390899089908990899089908990601a9060040161589c565b60006040518083038186803b158015612a3b57600080fd5b505af4158015612a4f573d6000803e3d6000fd5b50505050505050505050565b601354604051632b05429560e21b8152600481018390526060916001600160a01b03169063ac150a5490602401600060405180830381865afa158015612aa5573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610b0b91908101906158e4565b601454600390600160a81b900460ff16158015612af8575060145460ff808316600160a01b90920416105b612b5b5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610c18565b6014805460ff60a81b1960ff8416600160a01b021661ffff60a01b1990911617600160a81b179055612b8f898989896138ba565b601380546001600160a01b03199081166001600160a01b0388811691909117909255601480548216878416179055601580548216868416179055601b80549091169184169190911790556040514281527f400175a56dd3710794078f7b9dbe8296ac94c5a248dfd51bb22ed4ab9eaa9fbf9060200160405180910390a16014805460ff60a81b1916905560405160ff821681527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a1505050505050505050565b6000818152600a6020526040902060010180546060919061252e906150be565b612c86813560bc61306d565b73__$e2c9a8f399964af47bb173600dd9e5f662$__631873e2188235612caf60208501856150f2565b60135486356000908152600d60209081526040808320818b01358085529252918290205491516001600160e01b031960e089901b168152610ce19695949391926001600160a01b039092169190600401615604565b612d0d8161390b565b336000818152600c602052604080822080549085905590519092839285927fce95332e6082aebeb8058a7b56d1a109f67d6550552ed04d36aca4a6acd4d7de9190a45050565b6000612d61843560cb61306d565b65416e7955726960d01b6000612d778635613467565b601354604051633610bf0960e11b81529192506000916001600160a01b0390911690636c217e1290612db19084908a908a90600401615445565b6020604051808303816000875af1158015612dd0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612df491906151a0565b905073__$2873850e5b3f5697aafc2f5cd9f27cc884$__6342a34a53888486858b8b604051602001612e27929190615665565b604051602081830303815290604052600e6040518763ffffffff1660e01b8152600401612e59969594939291906153fa565b60006040518083038186803b158015612e7157600080fd5b505af4158015612e85573d6000803e3d6000fd5b50939998505050505050505050565b6000612ea2823560b461306d565b612f306040518060a00160405280846020016020810190612ec3919061484a565b6001600160a01b03168152602001612eec856020016020810190612ee7919061484a565b613971565b815260200160405180602001604052806000815250815260200160006001600160a01b031681526020016040518060200160405280600081525081525060006137dd565b60135460408051639ec52a2360e01b8152853560048201526024810184905290850135604482015260e06064820152600060e482018190526001600160a01b03909216608482015260a4810191909152600d60c482015290915073__$e2c9a8f399964af47bb173600dd9e5f662$__90639ec52a239061010401611191565b60006001600160e01b031982166380ac58cd60e01b1480612fe057506001600160e01b03198216635b5e139f60e01b145b80610b0b57506301ffc9a760e01b6001600160e01b0319831614610b0b565b600081815260046020526040902080546001600160a01b0319166001600160a01b0384169081179091558190613034826119b2565b6001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45050565b61307682613a7a565b1561307f575050565b6015546001600160a01b031633036130bb576000828152601960209081526040808320328452909152902054600190821c8116036130bb575050565b6000828152601960209081526040808320338452909152902054600190821c8116036130e5575050565b604051632c4bc2b960e21b815260040160405180910390fd5b6000818152600260205260409020546001600160a01b0316613136576040516375af0fc960e11b815260048101829052602401610c18565b50565b6000818152600260205260408120546001600160a01b03166131b25760405162461bcd60e51b815260206004820152602c60248201527f4552433732313a206f70657261746f7220717565727920666f72206e6f6e657860448201526b34b9ba32b73a103a37b5b2b760a11b6064820152608401610c18565b60006131bd836119b2565b9050806001600160a01b0316846001600160a01b031614806131f85750836001600160a01b03166131ed84610ba3565b6001600160a01b0316145b8061322857506001600160a01b0380821660009081526005602090815260408083209388168352929052205460ff165b949350505050565b826001600160a01b0316613243826119b2565b6001600160a01b0316146132a75760405162461bcd60e51b815260206004820152602560248201527f4552433732313a207472616e736665722066726f6d20696e636f72726563742060448201526437bbb732b960d91b6064820152608401610c18565b6001600160a01b0382166133095760405162461bcd60e51b8152602060048201526024808201527f4552433732313a207472616e7366657220746f20746865207a65726f206164646044820152637265737360e01b6064820152608401610c18565b613314838383613ada565b61331f600082612fff565b6001600160a01b0383166000908152600360205260408120805460019290613348908490615951565b90915550506001600160a01b0382166000908152600360205260408120805460019290613376908490615964565b909155505060008181526002602052604080822080546001600160a01b0319166001600160a01b0386811691821790925591518493918716917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91a4610e24838383613bac565b6000838152601a6020908152604080832085845290915281206134008184613be9565b1561340f57600091505061112e565b61341c6002820184613be9565b1561342b57600191505061112e565b60008581526019602090815260408083206001600160a01b038716845290915290205461345e9060c31c60019081161490565b95945050505050565b6000818152600a6020526040812060030180548290613485906156fd565b918290555092915050565b61349a3382613139565b6134e65760405162461bcd60e51b815260206004820152601b60248201527f4e4654426173653a204e6f744f776e65724f72417070726f76656400000000006044820152606401610c18565b61313681613c0b565b6060600061112e83613cba565b61350582613a7a565b1561350e575050565b6015546001600160a01b031633036135345761352b8282326133dd565b15613534575050565b61353f8282336133dd565b15613548575050565b604051631a1d1d4760e11b815260040160405180910390fd5b6000828152600e60209081526040808320848452909152902060050154600160a01b900460ff16156135a657604051631f0fc8f560e11b815260040160405180910390fd5b6000828152600a6020526040902060030154811115611687576040516364783acb60e01b815260040160405180910390fd5b6000828152600e60209081526040808320848452909152902060050154600160a81b900460ff161561168757604051630bc06a0f60e21b815260040160405180910390fd5b816001600160a01b0316836001600160a01b03160361367e5760405162461bcd60e51b815260206004820152601960248201527f4552433732313a20617070726f766520746f2063616c6c6572000000000000006044820152606401610c18565b6001600160a01b03838116600081815260056020908152604080832094871680845294825291829020805460ff191686151590811790915591519182527f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31910160405180910390a3505050565b6000818152600b60205260409020541561313657604051631b659b9f60e21b815260040160405180910390fd5b80518190601f81118061372b5750600381105b1561374957604051636f819c2160e11b815260040160405180910390fd5b60005b818110156123e35761377d838281518110613769576137696155ee565b01602001516001600160f81b031916613d15565b60010161374c565b606061112e8383604051806060016040528060278152602001615afd60279139613dc3565b6137b5848484613230565b6137c184848484613e3b565b6123e35760405162461bcd60e51b8152600401610c1890615977565b60006137f38360200151805190602001206136eb565b8115613806576138068360200151613718565b601260008154613815906156fd565b918290555083519091506138299082613f3c565b73__$d8529f0216812a2dec6b96adf1943c3537$__634daae5688460000151856020015186604001518760600151886080015187600b600a6040518963ffffffff1660e01b81526004016138849897969594939291906159c9565b60006040518083038186803b15801561389c57600080fd5b505af41580156138b0573d6000803e3d6000fd5b5050505092915050565b6138c684848484613f56565b7f414cd0b34676984f09a5f76ce9718d4062e50283abe0e7e274a9a5b4e0c99c3084848484426040516138fd959493929190615a40565b60405180910390a150505050565b6000613916826119b2565b6015549091506001600160a01b03163314801561393b5750326001600160a01b038216145b15613944575050565b6001600160a01b0381163303613958575050565b604051631b0c476f60e11b815260040160405180910390fd5b60408051602a80825260608281019093526f181899199a1a9b1b9c1cb0b131b232b360811b916001600160a01b0385169160009190602082018180368337019050509050600360fc1b816000815181106139cd576139cd6155ee565b60200101906001600160f81b031916908160001a905350600f60fb1b816001815181106139fc576139fc6155ee565b60200101906001600160f81b031916908160001a90535060295b6001811115613a71578383600f1660108110613a3457613a346155ee565b1a60f81b828281518110613a4a57613a4a6155ee565b60200101906001600160f81b031916908160001a90535060049290921c9160001901613a16565b50949350505050565b600080613a86836119b2565b90506001600160a01b0381163303613aa15750600192915050565b6015546001600160a01b031633148015613ac35750326001600160a01b038216145b15613ad15750600192915050565b50600092915050565b601b546001600160a01b03848116911614613ba1576000818152601860205260408120613b0690613f71565b600083815260186020526040812091925090613b21906134ef565b905060005b82811015613b6357613b5184838381518110613b4457613b446155ee565b6020026020010151613f7b565b80613b5b816156fd565b915050613b26565b506001600160a01b0385166000908152600c6020526040902054839003613b9e576001600160a01b0385166000908152600c60205260408120555b50505b610e24838383613fb4565b6001600160a01b0382166000908152600c60205260408120549003610e24576001600160a01b03919091166000908152600c602052604090205550565b6001600160a01b0381166000908152600183016020526040812054151561112e565b6000613c16826119b2565b9050613c2481600084613ada565b613c2f600083612fff565b6001600160a01b0381166000908152600360205260408120805460019290613c58908490615951565b909155505060008281526002602052604080822080546001600160a01b0319169055518391906001600160a01b038416907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef908390a461168781600084613bac565b6060816000018054806020026020016040519081016040528092919081815260200182805480156125a757602002820191906000526020600020905b815481526020019060010190808311613cf65750505050509050919050565b600360fc1b6001600160f81b031982161080613d3e5750603d60f91b6001600160f81b03198216115b80613d6e5750603960f81b6001600160f81b03198216118015613d6e5750606160f81b6001600160f81b03198216105b8015613d885750602d60f81b6001600160f81b0319821614155b8015613da25750605f60f81b6001600160f81b0319821614155b15613136576040516001621693dd60e01b0319815260040160405180910390fd5b6060600080856001600160a01b031685604051613de09190615a7a565b600060405180830381855af49150503d8060008114613e1b576040519150601f19603f3d011682016040523d82523d6000602084013e613e20565b606091505b5091509150613e318683838761406c565b9695505050505050565b60006001600160a01b0384163b15613f3157604051630a85bd0160e11b81526001600160a01b0385169063150b7a0290613e7f903390899088908890600401615a96565b6020604051808303816000875af1925050508015613eba575060408051601f3d908101601f19168201909252613eb791810190615ac9565b60015b613f17573d808015613ee8576040519150601f19603f3d011682016040523d82523d6000602084013e613eed565b606091505b508051600003613f0f5760405162461bcd60e51b8152600401610c1890615977565b805181602001fd5b6001600160e01b031916630a85bd0160e11b149050613228565b506001949350505050565b6116878282604051806020016040528060008152506140e5565b6000613f6384868361551b565b506001610d0d82848361551b565b6000610b0b825490565b60008281526019602090815260408083206001600160a01b0385168452825280832083905584835260189091529020610e249082614118565b6001600160a01b03831661400f5761400a81600880546000838152600960205260408120829055600182018355919091527ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee30155565b614032565b816001600160a01b0316836001600160a01b03161461403257614032838261412d565b6001600160a01b03821661404957610e24816141ca565b826001600160a01b0316826001600160a01b031614610e2457610e248282614279565b606083156140db5782516000036140d4576001600160a01b0385163b6140d45760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610c18565b5081613228565b61322883836142bd565b6140ef83836142e7565b6140fc6000848484613e3b565b610e245760405162461bcd60e51b8152600401610c1890615977565b600061112e836001600160a01b03841661443d565b6000600161413a84611a43565b6141449190615951565b600083815260076020526040902054909150808214614197576001600160a01b03841660009081526006602090815260408083208584528252808320548484528184208190558352600790915290208190555b5060009182526007602090815260408084208490556001600160a01b039094168352600681528383209183525290812055565b6008546000906141dc90600190615951565b60008381526009602052604081205460088054939450909284908110614204576142046155ee565b906000526020600020015490508060088381548110614225576142256155ee565b600091825260208083209091019290925582815260099091526040808220849055858252812055600880548061425d5761425d615ae6565b6001900381819060005260206000200160009055905550505050565b600061428483611a43565b6001600160a01b039093166000908152600660209081526040808320868452825280832085905593825260079052919091209190915550565b8151156142cd5781518083602001fd5b8060405162461bcd60e51b8152600401610c18919061464e565b6001600160a01b03821661433d5760405162461bcd60e51b815260206004820181905260248201527f4552433732313a206d696e7420746f20746865207a65726f20616464726573736044820152606401610c18565b6000818152600260205260409020546001600160a01b0316156143a25760405162461bcd60e51b815260206004820152601c60248201527f4552433732313a20746f6b656e20616c7265616479206d696e746564000000006044820152606401610c18565b6143ae60008383613ada565b6001600160a01b03821660009081526003602052604081208054600192906143d7908490615964565b909155505060008181526002602052604080822080546001600160a01b0319166001600160a01b03861690811790915590518392907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef908290a461168760008383613bac565b60008181526001830160205260408120548015614526576000614461600183615951565b855490915060009061447590600190615951565b90508181146144da576000866000018281548110614495576144956155ee565b90600052602060002001549050808760000184815481106144b8576144b86155ee565b6000918252602080832090910192909255918252600188019052604090208390555b85548690806144eb576144eb615ae6565b600190038181906000526020600020016000905590558560010160008681526020019081526020016000206000905560019350505050610b0b565b6000915050610b0b565b50805461453c906150be565b6000825580601f1061454c575050565b601f01602090049060005260206000209081019061313691906145b2565b6040518060c001604052806000815260200160608152602001606081526020016000815260200160006001600160a01b0316815260200160006001600160a01b031681525090565b5b808211156145c757600081556001016145b3565b5090565b6001600160e01b03198116811461313657600080fd5b6000602082840312156145f357600080fd5b813561112e816145cb565b60005b83811015614619578181015183820152602001614601565b50506000910152565b6000815180845261463a8160208601602086016145fe565b601f01601f19169290920160200192915050565b60208152600061112e6020830184614622565b60006020828403121561467357600080fd5b5035919050565b60006060828403121561468c57600080fd5b50919050565b6000602082840312156146a457600080fd5b81356001600160401b038111156146ba57600080fd5b6132288482850161467a565b80356001600160a01b03811681146146dd57600080fd5b919050565b600080604083850312156146f557600080fd5b6146fe836146c6565b946020939093013593505050565b60006060828403121561471e57600080fd5b61112e838361467a565b60006080828403121561468c57600080fd5b60006020828403121561474c57600080fd5b81356001600160401b0381111561476257600080fd5b61322884828501614728565b60008060006060848603121561478357600080fd5b83359250614793602085016146c6565b9150604084013590509250925092565b6000806000606084860312156147b857600080fd5b6147c1846146c6565b9250614793602085016146c6565b6000806000606084860312156147e457600080fd5b83359250602084013591506147fb604085016146c6565b90509250925092565b600060e0828403121561468c57600080fd5b60006020828403121561482857600080fd5b81356001600160401b0381111561483e57600080fd5b61322884828501614804565b60006020828403121561485c57600080fd5b61112e826146c6565b60006040828403121561468c57600080fd5b6000806060838503121561488a57600080fd5b82356001600160401b038111156148a057600080fd5b6148ac85828601614804565b9250506117a28460208501614865565b60008083601f8401126148ce57600080fd5b5081356001600160401b038111156148e557600080fd5b6020830191508360208285010111156148fd57600080fd5b9250929050565b60008060006040848603121561491957600080fd5b8335925060208401356001600160401b0381111561493657600080fd5b614942868287016148bc565b9497909650939450505050565b60006080828403121561496157600080fd5b61112e8383614728565b6000806040838503121561497e57600080fd5b82356001600160401b0381111561499457600080fd5b6149a085828601614804565b95602094909401359450505050565b600080604083850312156149c257600080fd5b50508035926020909101359150565b600081518084526020808501945080840160005b83811015614a0a5781516001600160a01b0316875295820195908201906001016149e5565b509495945050505050565b604081526000614a2860408301856149d1565b828103602084015261345e81856149d1565b60008060008060608587031215614a5057600080fd5b843593506020850135925060408501356001600160401b03811115614a7457600080fd5b614a80878288016148bc565b95989497509550505050565b60208152600061112e60208301846149d1565b60008060408385031215614ab257600080fd5b82356001600160401b03811115614ac857600080fd5b614ad485828601614804565b9250506117a2602084016146c6565b60008060408385031215614af657600080fd5b823591506117a2602084016146c6565b60008060208385031215614b1957600080fd5b82356001600160401b03811115614b2f57600080fd5b614b3b858286016148bc565b90969095509350505050565b60208152815160208201526000602083015160c06040840152614b6d60e0840182614622565b90506040840151601f19848303016060850152614b8a8282614622565b91505060608401516080840152608084015160018060a01b0380821660a08601528060a08701511660c086015250508091505092915050565b803580151581146146dd57600080fd5b60008060408385031215614be657600080fd5b614bef836146c6565b91506117a260208401614bc3565b60008083601f840112614c0f57600080fd5b5081356001600160401b03811115614c2657600080fd5b6020830191508360208260051b85010111156148fd57600080fd5b60008060208385031215614c5457600080fd5b82356001600160401b03811115614c6a57600080fd5b614b3b85828601614bfd565b6000602080830181845280855180835260408601915060408160051b870101925083870160005b82811015614ccb57603f19888603018452614cb9858351614622565b94509285019290850190600101614c9d565b5092979650505050505050565b634e487b7160e01b600052604160045260246000fd5b60405160a081016001600160401b0381118282101715614d1057614d10614cd8565b60405290565b604051601f8201601f191681016001600160401b0381118282101715614d3e57614d3e614cd8565b604052919050565b60006001600160401b03821115614d5f57614d5f614cd8565b50601f01601f191660200190565b600082601f830112614d7e57600080fd5b8135614d91614d8c82614d46565b614d16565b818152846020838601011115614da657600080fd5b816020850160208301376000918101602001919091529392505050565b60008060008060808587031215614dd957600080fd5b614de2856146c6565b9350614df0602086016146c6565b92506040850135915060608501356001600160401b03811115614e1257600080fd5b614e1e87828801614d6d565b91505092959194509250565b600060a0828403121561468c57600080fd5b600060208284031215614e4e57600080fd5b81356001600160401b03811115614e6457600080fd5b61322884828501614e2a565b60208152815160208201526020820151604082015260006040830151610100806060850152614ea3610120850183614622565b9150606085015160018060a01b0380821660808701528060808801511660a0870152505060a0850151614ee160c08601826001600160a01b03169052565b5060c085015180151560e08601525060e0850151801515858301525090949350505050565b60008060008060008060808789031215614f1f57600080fd5b863595506020870135945060408701356001600160401b0380821115614f4457600080fd5b614f508a838b01614bfd565b90965094506060890135915080821115614f6957600080fd5b50614f7689828a01614bfd565b979a9699509497509295939492505050565b60008060008060008060008060c0898b031215614fa457600080fd5b88356001600160401b0380821115614fbb57600080fd5b614fc78c838d016148bc565b909a50985060208b0135915080821115614fe057600080fd5b50614fed8b828c016148bc565b9097509550615000905060408a016146c6565b935061500e60608a016146c6565b925061501c60808a016146c6565b915061502a60a08a016146c6565b90509295985092959890939650565b6000806040838503121561504c57600080fd5b615055836146c6565b91506117a2602084016146c6565b60008060006040848603121561507857600080fd5b83356001600160401b038082111561508f57600080fd5b61509b87838801614804565b945060208601359150808211156150b157600080fd5b50614942868287016148bc565b600181811c908216806150d257607f821691505b60208210810361468c57634e487b7160e01b600052602260045260246000fd5b6000808335601e1984360301811261510957600080fd5b8301803591506001600160401b0382111561512357600080fd5b6020019150368190038213156148fd57600080fd5b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b6001600160a01b0386811682528516602082015260806040820181905260009061518e9083018587615138565b90508260608301529695505050505050565b6000602082840312156151b257600080fd5b5051919050565b8581526001600160a01b038516602082015260806040820181905260009061518e9083018587615138565b948552602085019390935260408401919091526001600160a01b03166060830152608082015260a00190565b88815287602082015286604082015260e06060820152600061523660e083018789615138565b6001600160a01b0395861660808401529390941660a082015260c001529695505050505050565b60208082526031908201527f4552433732313a207472616e736665722063616c6c6572206973206e6f74206f6040820152701ddb995c881b9bdc88185c1c1c9bdd9959607a1b606082015260800190565b6000808335601e198436030181126152c557600080fd5b83016020810192503590506001600160401b038111156152e457600080fd5b8036038213156148fd57600080fd5b80358252600061530660208301836152ae565b60e0602086015261531b60e086018284615138565b91505061532a604084016146c6565b6001600160a01b03818116604087015261534760608601866152ae565b9250868403606088015261535c848483615138565b9350508061536c608087016146c6565b166080870152505061538160a08401846152ae565b85830360a0870152615394838284615138565b925050506153a460c08401614bc3565b151560c08501528091505092915050565b60c0815260006153c860c08301886152f3565b602083810197909752604083019590955250606081019290925281830360808301526000835260a09091015201919050565b60c08152600061540d60c08301896152f3565b87602084015286604084015285606084015282810360808401526154318186614622565b9150508260a0830152979650505050505050565b83815260406020820152600061345e604083018486615138565b600080835461546d816150be565b60018281168015615485576001811461549a576154c9565b60ff19841687528215158302870194506154c9565b8760005260208060002060005b858110156154c05781548a8201529084019082016154a7565b50505082870194505b50929695505050505050565b601f821115610e2457600081815260208120601f850160051c810160208610156154fc5750805b601f850160051c820191505b81811015610f2b57828155600101615508565b6001600160401b0383111561553257615532614cd8565b6155468361554083546150be565b836154d5565b6000601f84116001811461557a57600085156155625750838201355b600019600387901b1c1916600186901b178355610d0d565b600083815260209020601f19861690835b828110156155ab578685013582556020948501946001909201910161558b565b50868210156155c85760001960f88860031b161c19848701351681555b505060018560011b0183555050505050565b60208152600061112b602083018486615138565b634e487b7160e01b600052603260045260246000fd5b86815260a06020820152600061561e60a083018789615138565b6040830195909552506001600160a01b039290921660608301526080909101529392505050565b85815284602082015260806040820152600061518e608083018587615138565b8183823760009101908152919050565b85815260806020820152600061568f608083018688615138565b604083019490945250606001529392505050565b878152866020820152600060018060a01b03808816604084015260c060608401526156d260c084018789615138565b941660808301525060a0015295945050505050565b634e487b7160e01b600052601160045260246000fd5b60006001820161570f5761570f6156e7565b5060010190565b60006101008b83528a6020840152896040840152886060840152806080840152615743818401888a615138565b6001600160a01b0396871660a08501529490951660c08301525060e00152979650505050505050565b600060a0823603121561577e57600080fd5b615786614cee565b61578f836146c6565b815260208301356001600160401b03808211156157ab57600080fd5b6157b736838701614d6d565b602084015260408501359150808211156157d057600080fd5b6157dc36838701614d6d565b60408401526157ed606086016146c6565b6060840152608085013591508082111561580657600080fd5b5061581336828601614d6d565b60808301525092915050565b86815285602082015260018060a01b038516604082015260a06060820152600061584d60a083018587615138565b9050826080830152979650505050505050565b8183526000602080850194508260005b85811015614a0a576001600160a01b03615889836146c6565b1687529582019590820190600101615870565b87815286602082015260a0604082015260006158bc60a083018789615860565b82810360608401526158cf818688615860565b91505082608083015298975050505050505050565b6000602082840312156158f657600080fd5b81516001600160401b0381111561590c57600080fd5b8201601f8101841361591d57600080fd5b805161592b614d8c82614d46565b81815285602083850101111561594057600080fd5b61345e8260208301602086016145fe565b81810381811115610b0b57610b0b6156e7565b80820180821115610b0b57610b0b6156e7565b60208082526032908201527f4552433732313a207472616e7366657220746f206e6f6e20455243373231526560408201527131b2b4bb32b91034b6b83632b6b2b73a32b960711b606082015260800190565b6001600160a01b038981168252610100602083018190526000916159ef8483018c614622565b91508382036040850152615a03828b614622565b908916606085015283810360808501529050615a1f8188614622565b60a0840196909652505060c081019290925260e09091015295945050505050565b606081526000615a54606083018789615138565b8281036020840152615a67818688615138565b9150508260408301529695505050505050565b60008251615a8c8184602087016145fe565b9190910192915050565b6001600160a01b0385811682528416602082015260408101839052608060608201819052600090613e3190830184614622565b600060208284031215615adb57600080fd5b815161112e816145cb565b634e487b7160e01b600052603160045260246000fdfe416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a26469706673582212202919b220b9f3f760cd7f6f1d0ca99c488da60af16e51055ad0dbb7c1c8ad775264736f6c63430008100033",
}

// ContractsABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractsMetaData.ABI instead.
var ContractsABI = ContractsMetaData.ABI

// ContractsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractsMetaData.Bin instead.
var ContractsBin = ContractsMetaData.Bin

// DeployContracts deploys a new Ethereum contract, binding an instance of Contracts to it.
func DeployContracts(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Contracts, error) {
	parsed, err := ContractsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contracts{ContractsCaller: ContractsCaller{contract: contract}, ContractsTransactor: ContractsTransactor{contract: contract}, ContractsFilterer: ContractsFilterer{contract: contract}}, nil
}

// Contracts is an auto generated Go binding around an Ethereum contract.
type Contracts struct {
	ContractsCaller     // Read-only binding to the contract
	ContractsTransactor // Write-only binding to the contract
	ContractsFilterer   // Log filterer for contract events
}

// ContractsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractsSession struct {
	Contract     *Contracts        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractsCallerSession struct {
	Contract *ContractsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ContractsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractsTransactorSession struct {
	Contract     *ContractsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ContractsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractsRaw struct {
	Contract *Contracts // Generic contract binding to access the raw methods on
}

// ContractsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractsCallerRaw struct {
	Contract *ContractsCaller // Generic read-only contract binding to access the raw methods on
}

// ContractsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractsTransactorRaw struct {
	Contract *ContractsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContracts creates a new instance of Contracts, bound to a specific deployed contract.
func NewContracts(address common.Address, backend bind.ContractBackend) (*Contracts, error) {
	contract, err := bindContracts(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contracts{ContractsCaller: ContractsCaller{contract: contract}, ContractsTransactor: ContractsTransactor{contract: contract}, ContractsFilterer: ContractsFilterer{contract: contract}}, nil
}

// NewContractsCaller creates a new read-only instance of Contracts, bound to a specific deployed contract.
func NewContractsCaller(address common.Address, caller bind.ContractCaller) (*ContractsCaller, error) {
	contract, err := bindContracts(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsCaller{contract: contract}, nil
}

// NewContractsTransactor creates a new write-only instance of Contracts, bound to a specific deployed contract.
func NewContractsTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractsTransactor, error) {
	contract, err := bindContracts(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsTransactor{contract: contract}, nil
}

// NewContractsFilterer creates a new log filterer instance of Contracts, bound to a specific deployed contract.
func NewContractsFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractsFilterer, error) {
	contract, err := bindContracts(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractsFilterer{contract: contract}, nil
}

// bindContracts binds a generic wrapper to an already deployed contract.
func bindContracts(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.ContractsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Contracts *ContractsCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Contracts *ContractsSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Contracts.Contract.BalanceOf(&_Contracts.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Contracts *ContractsCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Contracts.Contract.BalanceOf(&_Contracts.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Contracts *ContractsCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Contracts *ContractsSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Contracts.Contract.GetApproved(&_Contracts.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Contracts *ContractsCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Contracts.Contract.GetApproved(&_Contracts.CallOpts, tokenId)
}

// GetCharacter is a free data retrieval call binding the contract method 0xdabb0531.
//
// Solidity: function getCharacter(uint256 characterId) view returns((uint256,string,string,uint256,address,address))
func (_Contracts *ContractsCaller) GetCharacter(opts *bind.CallOpts, characterId *big.Int) (DataTypesCharacter, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getCharacter", characterId)

	if err != nil {
		return *new(DataTypesCharacter), err
	}

	out0 := *abi.ConvertType(out[0], new(DataTypesCharacter)).(*DataTypesCharacter)

	return out0, err

}

// GetCharacter is a free data retrieval call binding the contract method 0xdabb0531.
//
// Solidity: function getCharacter(uint256 characterId) view returns((uint256,string,string,uint256,address,address))
func (_Contracts *ContractsSession) GetCharacter(characterId *big.Int) (DataTypesCharacter, error) {
	return _Contracts.Contract.GetCharacter(&_Contracts.CallOpts, characterId)
}

// GetCharacter is a free data retrieval call binding the contract method 0xdabb0531.
//
// Solidity: function getCharacter(uint256 characterId) view returns((uint256,string,string,uint256,address,address))
func (_Contracts *ContractsCallerSession) GetCharacter(characterId *big.Int) (DataTypesCharacter, error) {
	return _Contracts.Contract.GetCharacter(&_Contracts.CallOpts, characterId)
}

// GetCharacterByHandle is a free data retrieval call binding the contract method 0x9a50248d.
//
// Solidity: function getCharacterByHandle(string handle) view returns((uint256,string,string,uint256,address,address))
func (_Contracts *ContractsCaller) GetCharacterByHandle(opts *bind.CallOpts, handle string) (DataTypesCharacter, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getCharacterByHandle", handle)

	if err != nil {
		return *new(DataTypesCharacter), err
	}

	out0 := *abi.ConvertType(out[0], new(DataTypesCharacter)).(*DataTypesCharacter)

	return out0, err

}

// GetCharacterByHandle is a free data retrieval call binding the contract method 0x9a50248d.
//
// Solidity: function getCharacterByHandle(string handle) view returns((uint256,string,string,uint256,address,address))
func (_Contracts *ContractsSession) GetCharacterByHandle(handle string) (DataTypesCharacter, error) {
	return _Contracts.Contract.GetCharacterByHandle(&_Contracts.CallOpts, handle)
}

// GetCharacterByHandle is a free data retrieval call binding the contract method 0x9a50248d.
//
// Solidity: function getCharacterByHandle(string handle) view returns((uint256,string,string,uint256,address,address))
func (_Contracts *ContractsCallerSession) GetCharacterByHandle(handle string) (DataTypesCharacter, error) {
	return _Contracts.Contract.GetCharacterByHandle(&_Contracts.CallOpts, handle)
}

// GetCharacterUri is a free data retrieval call binding the contract method 0x144a3e83.
//
// Solidity: function getCharacterUri(uint256 characterId) view returns(string)
func (_Contracts *ContractsCaller) GetCharacterUri(opts *bind.CallOpts, characterId *big.Int) (string, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getCharacterUri", characterId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetCharacterUri is a free data retrieval call binding the contract method 0x144a3e83.
//
// Solidity: function getCharacterUri(uint256 characterId) view returns(string)
func (_Contracts *ContractsSession) GetCharacterUri(characterId *big.Int) (string, error) {
	return _Contracts.Contract.GetCharacterUri(&_Contracts.CallOpts, characterId)
}

// GetCharacterUri is a free data retrieval call binding the contract method 0x144a3e83.
//
// Solidity: function getCharacterUri(uint256 characterId) view returns(string)
func (_Contracts *ContractsCallerSession) GetCharacterUri(characterId *big.Int) (string, error) {
	return _Contracts.Contract.GetCharacterUri(&_Contracts.CallOpts, characterId)
}

// GetHandle is a free data retrieval call binding the contract method 0xec81d194.
//
// Solidity: function getHandle(uint256 characterId) view returns(string)
func (_Contracts *ContractsCaller) GetHandle(opts *bind.CallOpts, characterId *big.Int) (string, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getHandle", characterId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetHandle is a free data retrieval call binding the contract method 0xec81d194.
//
// Solidity: function getHandle(uint256 characterId) view returns(string)
func (_Contracts *ContractsSession) GetHandle(characterId *big.Int) (string, error) {
	return _Contracts.Contract.GetHandle(&_Contracts.CallOpts, characterId)
}

// GetHandle is a free data retrieval call binding the contract method 0xec81d194.
//
// Solidity: function getHandle(uint256 characterId) view returns(string)
func (_Contracts *ContractsCallerSession) GetHandle(characterId *big.Int) (string, error) {
	return _Contracts.Contract.GetHandle(&_Contracts.CallOpts, characterId)
}

// GetLinkModule4Address is a free data retrieval call binding the contract method 0x31b9d08c.
//
// Solidity: function getLinkModule4Address(address account) view returns(address)
func (_Contracts *ContractsCaller) GetLinkModule4Address(opts *bind.CallOpts, account common.Address) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getLinkModule4Address", account)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLinkModule4Address is a free data retrieval call binding the contract method 0x31b9d08c.
//
// Solidity: function getLinkModule4Address(address account) view returns(address)
func (_Contracts *ContractsSession) GetLinkModule4Address(account common.Address) (common.Address, error) {
	return _Contracts.Contract.GetLinkModule4Address(&_Contracts.CallOpts, account)
}

// GetLinkModule4Address is a free data retrieval call binding the contract method 0x31b9d08c.
//
// Solidity: function getLinkModule4Address(address account) view returns(address)
func (_Contracts *ContractsCallerSession) GetLinkModule4Address(account common.Address) (common.Address, error) {
	return _Contracts.Contract.GetLinkModule4Address(&_Contracts.CallOpts, account)
}

// GetLinkModule4ERC721 is a free data retrieval call binding the contract method 0x2209d145.
//
// Solidity: function getLinkModule4ERC721(address tokenAddress, uint256 tokenId) view returns(address)
func (_Contracts *ContractsCaller) GetLinkModule4ERC721(opts *bind.CallOpts, tokenAddress common.Address, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getLinkModule4ERC721", tokenAddress, tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLinkModule4ERC721 is a free data retrieval call binding the contract method 0x2209d145.
//
// Solidity: function getLinkModule4ERC721(address tokenAddress, uint256 tokenId) view returns(address)
func (_Contracts *ContractsSession) GetLinkModule4ERC721(tokenAddress common.Address, tokenId *big.Int) (common.Address, error) {
	return _Contracts.Contract.GetLinkModule4ERC721(&_Contracts.CallOpts, tokenAddress, tokenId)
}

// GetLinkModule4ERC721 is a free data retrieval call binding the contract method 0x2209d145.
//
// Solidity: function getLinkModule4ERC721(address tokenAddress, uint256 tokenId) view returns(address)
func (_Contracts *ContractsCallerSession) GetLinkModule4ERC721(tokenAddress common.Address, tokenId *big.Int) (common.Address, error) {
	return _Contracts.Contract.GetLinkModule4ERC721(&_Contracts.CallOpts, tokenAddress, tokenId)
}

// GetLinkModule4Linklist is a free data retrieval call binding the contract method 0xfe9299fb.
//
// Solidity: function getLinkModule4Linklist(uint256 tokenId) view returns(address)
func (_Contracts *ContractsCaller) GetLinkModule4Linklist(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getLinkModule4Linklist", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLinkModule4Linklist is a free data retrieval call binding the contract method 0xfe9299fb.
//
// Solidity: function getLinkModule4Linklist(uint256 tokenId) view returns(address)
func (_Contracts *ContractsSession) GetLinkModule4Linklist(tokenId *big.Int) (common.Address, error) {
	return _Contracts.Contract.GetLinkModule4Linklist(&_Contracts.CallOpts, tokenId)
}

// GetLinkModule4Linklist is a free data retrieval call binding the contract method 0xfe9299fb.
//
// Solidity: function getLinkModule4Linklist(uint256 tokenId) view returns(address)
func (_Contracts *ContractsCallerSession) GetLinkModule4Linklist(tokenId *big.Int) (common.Address, error) {
	return _Contracts.Contract.GetLinkModule4Linklist(&_Contracts.CallOpts, tokenId)
}

// GetLinklistContract is a free data retrieval call binding the contract method 0xc053f6b8.
//
// Solidity: function getLinklistContract() view returns(address)
func (_Contracts *ContractsCaller) GetLinklistContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getLinklistContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLinklistContract is a free data retrieval call binding the contract method 0xc053f6b8.
//
// Solidity: function getLinklistContract() view returns(address)
func (_Contracts *ContractsSession) GetLinklistContract() (common.Address, error) {
	return _Contracts.Contract.GetLinklistContract(&_Contracts.CallOpts)
}

// GetLinklistContract is a free data retrieval call binding the contract method 0xc053f6b8.
//
// Solidity: function getLinklistContract() view returns(address)
func (_Contracts *ContractsCallerSession) GetLinklistContract() (common.Address, error) {
	return _Contracts.Contract.GetLinklistContract(&_Contracts.CallOpts)
}

// GetLinklistId is a free data retrieval call binding the contract method 0xd70e10c6.
//
// Solidity: function getLinklistId(uint256 characterId, bytes32 linkType) view returns(uint256)
func (_Contracts *ContractsCaller) GetLinklistId(opts *bind.CallOpts, characterId *big.Int, linkType [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getLinklistId", characterId, linkType)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLinklistId is a free data retrieval call binding the contract method 0xd70e10c6.
//
// Solidity: function getLinklistId(uint256 characterId, bytes32 linkType) view returns(uint256)
func (_Contracts *ContractsSession) GetLinklistId(characterId *big.Int, linkType [32]byte) (*big.Int, error) {
	return _Contracts.Contract.GetLinklistId(&_Contracts.CallOpts, characterId, linkType)
}

// GetLinklistId is a free data retrieval call binding the contract method 0xd70e10c6.
//
// Solidity: function getLinklistId(uint256 characterId, bytes32 linkType) view returns(uint256)
func (_Contracts *ContractsCallerSession) GetLinklistId(characterId *big.Int, linkType [32]byte) (*big.Int, error) {
	return _Contracts.Contract.GetLinklistId(&_Contracts.CallOpts, characterId, linkType)
}

// GetLinklistType is a free data retrieval call binding the contract method 0x8b4ca06a.
//
// Solidity: function getLinklistType(uint256 linkListId) view returns(bytes32)
func (_Contracts *ContractsCaller) GetLinklistType(opts *bind.CallOpts, linkListId *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getLinklistType", linkListId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetLinklistType is a free data retrieval call binding the contract method 0x8b4ca06a.
//
// Solidity: function getLinklistType(uint256 linkListId) view returns(bytes32)
func (_Contracts *ContractsSession) GetLinklistType(linkListId *big.Int) ([32]byte, error) {
	return _Contracts.Contract.GetLinklistType(&_Contracts.CallOpts, linkListId)
}

// GetLinklistType is a free data retrieval call binding the contract method 0x8b4ca06a.
//
// Solidity: function getLinklistType(uint256 linkListId) view returns(bytes32)
func (_Contracts *ContractsCallerSession) GetLinklistType(linkListId *big.Int) ([32]byte, error) {
	return _Contracts.Contract.GetLinklistType(&_Contracts.CallOpts, linkListId)
}

// GetLinklistUri is a free data retrieval call binding the contract method 0xdca27135.
//
// Solidity: function getLinklistUri(uint256 tokenId) view returns(string)
func (_Contracts *ContractsCaller) GetLinklistUri(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getLinklistUri", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetLinklistUri is a free data retrieval call binding the contract method 0xdca27135.
//
// Solidity: function getLinklistUri(uint256 tokenId) view returns(string)
func (_Contracts *ContractsSession) GetLinklistUri(tokenId *big.Int) (string, error) {
	return _Contracts.Contract.GetLinklistUri(&_Contracts.CallOpts, tokenId)
}

// GetLinklistUri is a free data retrieval call binding the contract method 0xdca27135.
//
// Solidity: function getLinklistUri(uint256 tokenId) view returns(string)
func (_Contracts *ContractsCallerSession) GetLinklistUri(tokenId *big.Int) (string, error) {
	return _Contracts.Contract.GetLinklistUri(&_Contracts.CallOpts, tokenId)
}

// GetNote is a free data retrieval call binding the contract method 0xdb491e80.
//
// Solidity: function getNote(uint256 characterId, uint256 noteId) view returns((bytes32,bytes32,string,address,address,address,bool,bool))
func (_Contracts *ContractsCaller) GetNote(opts *bind.CallOpts, characterId *big.Int, noteId *big.Int) (DataTypesNote, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getNote", characterId, noteId)

	if err != nil {
		return *new(DataTypesNote), err
	}

	out0 := *abi.ConvertType(out[0], new(DataTypesNote)).(*DataTypesNote)

	return out0, err

}

// GetNote is a free data retrieval call binding the contract method 0xdb491e80.
//
// Solidity: function getNote(uint256 characterId, uint256 noteId) view returns((bytes32,bytes32,string,address,address,address,bool,bool))
func (_Contracts *ContractsSession) GetNote(characterId *big.Int, noteId *big.Int) (DataTypesNote, error) {
	return _Contracts.Contract.GetNote(&_Contracts.CallOpts, characterId, noteId)
}

// GetNote is a free data retrieval call binding the contract method 0xdb491e80.
//
// Solidity: function getNote(uint256 characterId, uint256 noteId) view returns((bytes32,bytes32,string,address,address,address,bool,bool))
func (_Contracts *ContractsCallerSession) GetNote(characterId *big.Int, noteId *big.Int) (DataTypesNote, error) {
	return _Contracts.Contract.GetNote(&_Contracts.CallOpts, characterId, noteId)
}

// GetOperatorPermissions is a free data retrieval call binding the contract method 0x952be0ef.
//
// Solidity: function getOperatorPermissions(uint256 characterId, address operator) view returns(uint256)
func (_Contracts *ContractsCaller) GetOperatorPermissions(opts *bind.CallOpts, characterId *big.Int, operator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getOperatorPermissions", characterId, operator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOperatorPermissions is a free data retrieval call binding the contract method 0x952be0ef.
//
// Solidity: function getOperatorPermissions(uint256 characterId, address operator) view returns(uint256)
func (_Contracts *ContractsSession) GetOperatorPermissions(characterId *big.Int, operator common.Address) (*big.Int, error) {
	return _Contracts.Contract.GetOperatorPermissions(&_Contracts.CallOpts, characterId, operator)
}

// GetOperatorPermissions is a free data retrieval call binding the contract method 0x952be0ef.
//
// Solidity: function getOperatorPermissions(uint256 characterId, address operator) view returns(uint256)
func (_Contracts *ContractsCallerSession) GetOperatorPermissions(characterId *big.Int, operator common.Address) (*big.Int, error) {
	return _Contracts.Contract.GetOperatorPermissions(&_Contracts.CallOpts, characterId, operator)
}

// GetOperators is a free data retrieval call binding the contract method 0x6bf55d5f.
//
// Solidity: function getOperators(uint256 characterId) view returns(address[])
func (_Contracts *ContractsCaller) GetOperators(opts *bind.CallOpts, characterId *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getOperators", characterId)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetOperators is a free data retrieval call binding the contract method 0x6bf55d5f.
//
// Solidity: function getOperators(uint256 characterId) view returns(address[])
func (_Contracts *ContractsSession) GetOperators(characterId *big.Int) ([]common.Address, error) {
	return _Contracts.Contract.GetOperators(&_Contracts.CallOpts, characterId)
}

// GetOperators is a free data retrieval call binding the contract method 0x6bf55d5f.
//
// Solidity: function getOperators(uint256 characterId) view returns(address[])
func (_Contracts *ContractsCallerSession) GetOperators(characterId *big.Int) ([]common.Address, error) {
	return _Contracts.Contract.GetOperators(&_Contracts.CallOpts, characterId)
}

// GetOperators4Note is a free data retrieval call binding the contract method 0x49186953.
//
// Solidity: function getOperators4Note(uint256 characterId, uint256 noteId) view returns(address[] blocklist, address[] allowlist)
func (_Contracts *ContractsCaller) GetOperators4Note(opts *bind.CallOpts, characterId *big.Int, noteId *big.Int) (struct {
	Blocklist []common.Address
	Allowlist []common.Address
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getOperators4Note", characterId, noteId)

	outstruct := new(struct {
		Blocklist []common.Address
		Allowlist []common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Blocklist = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.Allowlist = *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)

	return *outstruct, err

}

// GetOperators4Note is a free data retrieval call binding the contract method 0x49186953.
//
// Solidity: function getOperators4Note(uint256 characterId, uint256 noteId) view returns(address[] blocklist, address[] allowlist)
func (_Contracts *ContractsSession) GetOperators4Note(characterId *big.Int, noteId *big.Int) (struct {
	Blocklist []common.Address
	Allowlist []common.Address
}, error) {
	return _Contracts.Contract.GetOperators4Note(&_Contracts.CallOpts, characterId, noteId)
}

// GetOperators4Note is a free data retrieval call binding the contract method 0x49186953.
//
// Solidity: function getOperators4Note(uint256 characterId, uint256 noteId) view returns(address[] blocklist, address[] allowlist)
func (_Contracts *ContractsCallerSession) GetOperators4Note(characterId *big.Int, noteId *big.Int) (struct {
	Blocklist []common.Address
	Allowlist []common.Address
}, error) {
	return _Contracts.Contract.GetOperators4Note(&_Contracts.CallOpts, characterId, noteId)
}

// GetPrimaryCharacterId is a free data retrieval call binding the contract method 0x2abc6bf6.
//
// Solidity: function getPrimaryCharacterId(address account) view returns(uint256)
func (_Contracts *ContractsCaller) GetPrimaryCharacterId(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getPrimaryCharacterId", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPrimaryCharacterId is a free data retrieval call binding the contract method 0x2abc6bf6.
//
// Solidity: function getPrimaryCharacterId(address account) view returns(uint256)
func (_Contracts *ContractsSession) GetPrimaryCharacterId(account common.Address) (*big.Int, error) {
	return _Contracts.Contract.GetPrimaryCharacterId(&_Contracts.CallOpts, account)
}

// GetPrimaryCharacterId is a free data retrieval call binding the contract method 0x2abc6bf6.
//
// Solidity: function getPrimaryCharacterId(address account) view returns(uint256)
func (_Contracts *ContractsCallerSession) GetPrimaryCharacterId(account common.Address) (*big.Int, error) {
	return _Contracts.Contract.GetPrimaryCharacterId(&_Contracts.CallOpts, account)
}

// GetRevision is a free data retrieval call binding the contract method 0x1316529d.
//
// Solidity: function getRevision() pure returns(uint256)
func (_Contracts *ContractsCaller) GetRevision(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getRevision")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRevision is a free data retrieval call binding the contract method 0x1316529d.
//
// Solidity: function getRevision() pure returns(uint256)
func (_Contracts *ContractsSession) GetRevision() (*big.Int, error) {
	return _Contracts.Contract.GetRevision(&_Contracts.CallOpts)
}

// GetRevision is a free data retrieval call binding the contract method 0x1316529d.
//
// Solidity: function getRevision() pure returns(uint256)
func (_Contracts *ContractsCallerSession) GetRevision() (*big.Int, error) {
	return _Contracts.Contract.GetRevision(&_Contracts.CallOpts)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Contracts *ContractsCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Contracts *ContractsSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Contracts.Contract.IsApprovedForAll(&_Contracts.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Contracts *ContractsCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Contracts.Contract.IsApprovedForAll(&_Contracts.CallOpts, owner, operator)
}

// IsOperatorAllowedForNote is a free data retrieval call binding the contract method 0x28fbb805.
//
// Solidity: function isOperatorAllowedForNote(uint256 characterId, uint256 noteId, address operator) view returns(bool)
func (_Contracts *ContractsCaller) IsOperatorAllowedForNote(opts *bind.CallOpts, characterId *big.Int, noteId *big.Int, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "isOperatorAllowedForNote", characterId, noteId, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperatorAllowedForNote is a free data retrieval call binding the contract method 0x28fbb805.
//
// Solidity: function isOperatorAllowedForNote(uint256 characterId, uint256 noteId, address operator) view returns(bool)
func (_Contracts *ContractsSession) IsOperatorAllowedForNote(characterId *big.Int, noteId *big.Int, operator common.Address) (bool, error) {
	return _Contracts.Contract.IsOperatorAllowedForNote(&_Contracts.CallOpts, characterId, noteId, operator)
}

// IsOperatorAllowedForNote is a free data retrieval call binding the contract method 0x28fbb805.
//
// Solidity: function isOperatorAllowedForNote(uint256 characterId, uint256 noteId, address operator) view returns(bool)
func (_Contracts *ContractsCallerSession) IsOperatorAllowedForNote(characterId *big.Int, noteId *big.Int, operator common.Address) (bool, error) {
	return _Contracts.Contract.IsOperatorAllowedForNote(&_Contracts.CallOpts, characterId, noteId, operator)
}

// IsPrimaryCharacter is a free data retrieval call binding the contract method 0x8734bbfc.
//
// Solidity: function isPrimaryCharacter(uint256 characterId) view returns(bool)
func (_Contracts *ContractsCaller) IsPrimaryCharacter(opts *bind.CallOpts, characterId *big.Int) (bool, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "isPrimaryCharacter", characterId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPrimaryCharacter is a free data retrieval call binding the contract method 0x8734bbfc.
//
// Solidity: function isPrimaryCharacter(uint256 characterId) view returns(bool)
func (_Contracts *ContractsSession) IsPrimaryCharacter(characterId *big.Int) (bool, error) {
	return _Contracts.Contract.IsPrimaryCharacter(&_Contracts.CallOpts, characterId)
}

// IsPrimaryCharacter is a free data retrieval call binding the contract method 0x8734bbfc.
//
// Solidity: function isPrimaryCharacter(uint256 characterId) view returns(bool)
func (_Contracts *ContractsCallerSession) IsPrimaryCharacter(characterId *big.Int) (bool, error) {
	return _Contracts.Contract.IsPrimaryCharacter(&_Contracts.CallOpts, characterId)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Contracts *ContractsCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Contracts *ContractsSession) Name() (string, error) {
	return _Contracts.Contract.Name(&_Contracts.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Contracts *ContractsCallerSession) Name() (string, error) {
	return _Contracts.Contract.Name(&_Contracts.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Contracts *ContractsCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Contracts *ContractsSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Contracts.Contract.OwnerOf(&_Contracts.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Contracts *ContractsCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Contracts.Contract.OwnerOf(&_Contracts.CallOpts, tokenId)
}

// Resolver is a free data retrieval call binding the contract method 0x04f3bcec.
//
// Solidity: function resolver() view returns(address)
func (_Contracts *ContractsCaller) Resolver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "resolver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Resolver is a free data retrieval call binding the contract method 0x04f3bcec.
//
// Solidity: function resolver() view returns(address)
func (_Contracts *ContractsSession) Resolver() (common.Address, error) {
	return _Contracts.Contract.Resolver(&_Contracts.CallOpts)
}

// Resolver is a free data retrieval call binding the contract method 0x04f3bcec.
//
// Solidity: function resolver() view returns(address)
func (_Contracts *ContractsCallerSession) Resolver() (common.Address, error) {
	return _Contracts.Contract.Resolver(&_Contracts.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Contracts *ContractsCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Contracts *ContractsSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Contracts.Contract.SupportsInterface(&_Contracts.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Contracts *ContractsCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Contracts.Contract.SupportsInterface(&_Contracts.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Contracts *ContractsCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Contracts *ContractsSession) Symbol() (string, error) {
	return _Contracts.Contract.Symbol(&_Contracts.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Contracts *ContractsCallerSession) Symbol() (string, error) {
	return _Contracts.Contract.Symbol(&_Contracts.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Contracts *ContractsCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Contracts *ContractsSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Contracts.Contract.TokenByIndex(&_Contracts.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Contracts *ContractsCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Contracts.Contract.TokenByIndex(&_Contracts.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Contracts *ContractsCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Contracts *ContractsSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Contracts.Contract.TokenOfOwnerByIndex(&_Contracts.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Contracts *ContractsCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Contracts.Contract.TokenOfOwnerByIndex(&_Contracts.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 characterId) view returns(string)
func (_Contracts *ContractsCaller) TokenURI(opts *bind.CallOpts, characterId *big.Int) (string, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "tokenURI", characterId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 characterId) view returns(string)
func (_Contracts *ContractsSession) TokenURI(characterId *big.Int) (string, error) {
	return _Contracts.Contract.TokenURI(&_Contracts.CallOpts, characterId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 characterId) view returns(string)
func (_Contracts *ContractsCallerSession) TokenURI(characterId *big.Int) (string, error) {
	return _Contracts.Contract.TokenURI(&_Contracts.CallOpts, characterId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Contracts *ContractsCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Contracts *ContractsSession) TotalSupply() (*big.Int, error) {
	return _Contracts.Contract.TotalSupply(&_Contracts.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Contracts *ContractsCallerSession) TotalSupply() (*big.Int, error) {
	return _Contracts.Contract.TotalSupply(&_Contracts.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Contracts *ContractsTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Contracts *ContractsSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Approve(&_Contracts.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Contracts *ContractsTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Approve(&_Contracts.TransactOpts, to, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_Contracts *ContractsTransactor) Burn(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "burn", tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_Contracts *ContractsSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Burn(&_Contracts.TransactOpts, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_Contracts *ContractsTransactorSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Burn(&_Contracts.TransactOpts, tokenId)
}

// CreateCharacter is a paid mutator transaction binding the contract method 0xcd69fe61.
//
// Solidity: function createCharacter((address,string,string,address,bytes) vars) returns(uint256 characterId)
func (_Contracts *ContractsTransactor) CreateCharacter(opts *bind.TransactOpts, vars DataTypesCreateCharacterData) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "createCharacter", vars)
}

// CreateCharacter is a paid mutator transaction binding the contract method 0xcd69fe61.
//
// Solidity: function createCharacter((address,string,string,address,bytes) vars) returns(uint256 characterId)
func (_Contracts *ContractsSession) CreateCharacter(vars DataTypesCreateCharacterData) (*types.Transaction, error) {
	return _Contracts.Contract.CreateCharacter(&_Contracts.TransactOpts, vars)
}

// CreateCharacter is a paid mutator transaction binding the contract method 0xcd69fe61.
//
// Solidity: function createCharacter((address,string,string,address,bytes) vars) returns(uint256 characterId)
func (_Contracts *ContractsTransactorSession) CreateCharacter(vars DataTypesCreateCharacterData) (*types.Transaction, error) {
	return _Contracts.Contract.CreateCharacter(&_Contracts.TransactOpts, vars)
}

// CreateThenLinkCharacter is a paid mutator transaction binding the contract method 0xf6479d77.
//
// Solidity: function createThenLinkCharacter((uint256,address,bytes32) vars) returns(uint256 characterId)
func (_Contracts *ContractsTransactor) CreateThenLinkCharacter(opts *bind.TransactOpts, vars DataTypescreateThenLinkCharacterData) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "createThenLinkCharacter", vars)
}

// CreateThenLinkCharacter is a paid mutator transaction binding the contract method 0xf6479d77.
//
// Solidity: function createThenLinkCharacter((uint256,address,bytes32) vars) returns(uint256 characterId)
func (_Contracts *ContractsSession) CreateThenLinkCharacter(vars DataTypescreateThenLinkCharacterData) (*types.Transaction, error) {
	return _Contracts.Contract.CreateThenLinkCharacter(&_Contracts.TransactOpts, vars)
}

// CreateThenLinkCharacter is a paid mutator transaction binding the contract method 0xf6479d77.
//
// Solidity: function createThenLinkCharacter((uint256,address,bytes32) vars) returns(uint256 characterId)
func (_Contracts *ContractsTransactorSession) CreateThenLinkCharacter(vars DataTypescreateThenLinkCharacterData) (*types.Transaction, error) {
	return _Contracts.Contract.CreateThenLinkCharacter(&_Contracts.TransactOpts, vars)
}

// DeleteNote is a paid mutator transaction binding the contract method 0xc2a6fe3b.
//
// Solidity: function deleteNote(uint256 characterId, uint256 noteId) returns()
func (_Contracts *ContractsTransactor) DeleteNote(opts *bind.TransactOpts, characterId *big.Int, noteId *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "deleteNote", characterId, noteId)
}

// DeleteNote is a paid mutator transaction binding the contract method 0xc2a6fe3b.
//
// Solidity: function deleteNote(uint256 characterId, uint256 noteId) returns()
func (_Contracts *ContractsSession) DeleteNote(characterId *big.Int, noteId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.DeleteNote(&_Contracts.TransactOpts, characterId, noteId)
}

// DeleteNote is a paid mutator transaction binding the contract method 0xc2a6fe3b.
//
// Solidity: function deleteNote(uint256 characterId, uint256 noteId) returns()
func (_Contracts *ContractsTransactorSession) DeleteNote(characterId *big.Int, noteId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.DeleteNote(&_Contracts.TransactOpts, characterId, noteId)
}

// GrantOperatorPermissions is a paid mutator transaction binding the contract method 0x206657f2.
//
// Solidity: function grantOperatorPermissions(uint256 characterId, address operator, uint256 permissionBitMap) returns()
func (_Contracts *ContractsTransactor) GrantOperatorPermissions(opts *bind.TransactOpts, characterId *big.Int, operator common.Address, permissionBitMap *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "grantOperatorPermissions", characterId, operator, permissionBitMap)
}

// GrantOperatorPermissions is a paid mutator transaction binding the contract method 0x206657f2.
//
// Solidity: function grantOperatorPermissions(uint256 characterId, address operator, uint256 permissionBitMap) returns()
func (_Contracts *ContractsSession) GrantOperatorPermissions(characterId *big.Int, operator common.Address, permissionBitMap *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.GrantOperatorPermissions(&_Contracts.TransactOpts, characterId, operator, permissionBitMap)
}

// GrantOperatorPermissions is a paid mutator transaction binding the contract method 0x206657f2.
//
// Solidity: function grantOperatorPermissions(uint256 characterId, address operator, uint256 permissionBitMap) returns()
func (_Contracts *ContractsTransactorSession) GrantOperatorPermissions(characterId *big.Int, operator common.Address, permissionBitMap *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.GrantOperatorPermissions(&_Contracts.TransactOpts, characterId, operator, permissionBitMap)
}

// GrantOperators4Note is a paid mutator transaction binding the contract method 0xdc17b6de.
//
// Solidity: function grantOperators4Note(uint256 characterId, uint256 noteId, address[] blocklist, address[] allowlist) returns()
func (_Contracts *ContractsTransactor) GrantOperators4Note(opts *bind.TransactOpts, characterId *big.Int, noteId *big.Int, blocklist []common.Address, allowlist []common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "grantOperators4Note", characterId, noteId, blocklist, allowlist)
}

// GrantOperators4Note is a paid mutator transaction binding the contract method 0xdc17b6de.
//
// Solidity: function grantOperators4Note(uint256 characterId, uint256 noteId, address[] blocklist, address[] allowlist) returns()
func (_Contracts *ContractsSession) GrantOperators4Note(characterId *big.Int, noteId *big.Int, blocklist []common.Address, allowlist []common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.GrantOperators4Note(&_Contracts.TransactOpts, characterId, noteId, blocklist, allowlist)
}

// GrantOperators4Note is a paid mutator transaction binding the contract method 0xdc17b6de.
//
// Solidity: function grantOperators4Note(uint256 characterId, uint256 noteId, address[] blocklist, address[] allowlist) returns()
func (_Contracts *ContractsTransactorSession) GrantOperators4Note(characterId *big.Int, noteId *big.Int, blocklist []common.Address, allowlist []common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.GrantOperators4Note(&_Contracts.TransactOpts, characterId, noteId, blocklist, allowlist)
}

// Initialize is a paid mutator transaction binding the contract method 0xe56f2fe4.
//
// Solidity: function initialize(string name_, string symbol_, address linklist_, address mintNFTImpl_, address periphery_, address newbieVilla_) returns()
func (_Contracts *ContractsTransactor) Initialize(opts *bind.TransactOpts, name_ string, symbol_ string, linklist_ common.Address, mintNFTImpl_ common.Address, periphery_ common.Address, newbieVilla_ common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "initialize", name_, symbol_, linklist_, mintNFTImpl_, periphery_, newbieVilla_)
}

// Initialize is a paid mutator transaction binding the contract method 0xe56f2fe4.
//
// Solidity: function initialize(string name_, string symbol_, address linklist_, address mintNFTImpl_, address periphery_, address newbieVilla_) returns()
func (_Contracts *ContractsSession) Initialize(name_ string, symbol_ string, linklist_ common.Address, mintNFTImpl_ common.Address, periphery_ common.Address, newbieVilla_ common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.Initialize(&_Contracts.TransactOpts, name_, symbol_, linklist_, mintNFTImpl_, periphery_, newbieVilla_)
}

// Initialize is a paid mutator transaction binding the contract method 0xe56f2fe4.
//
// Solidity: function initialize(string name_, string symbol_, address linklist_, address mintNFTImpl_, address periphery_, address newbieVilla_) returns()
func (_Contracts *ContractsTransactorSession) Initialize(name_ string, symbol_ string, linklist_ common.Address, mintNFTImpl_ common.Address, periphery_ common.Address, newbieVilla_ common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.Initialize(&_Contracts.TransactOpts, name_, symbol_, linklist_, mintNFTImpl_, periphery_, newbieVilla_)
}

// LinkAddress is a paid mutator transaction binding the contract method 0x388f5083.
//
// Solidity: function linkAddress((uint256,address,bytes32,bytes) vars) returns()
func (_Contracts *ContractsTransactor) LinkAddress(opts *bind.TransactOpts, vars DataTypeslinkAddressData) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "linkAddress", vars)
}

// LinkAddress is a paid mutator transaction binding the contract method 0x388f5083.
//
// Solidity: function linkAddress((uint256,address,bytes32,bytes) vars) returns()
func (_Contracts *ContractsSession) LinkAddress(vars DataTypeslinkAddressData) (*types.Transaction, error) {
	return _Contracts.Contract.LinkAddress(&_Contracts.TransactOpts, vars)
}

// LinkAddress is a paid mutator transaction binding the contract method 0x388f5083.
//
// Solidity: function linkAddress((uint256,address,bytes32,bytes) vars) returns()
func (_Contracts *ContractsTransactorSession) LinkAddress(vars DataTypeslinkAddressData) (*types.Transaction, error) {
	return _Contracts.Contract.LinkAddress(&_Contracts.TransactOpts, vars)
}

// LinkAnyUri is a paid mutator transaction binding the contract method 0x5fb88183.
//
// Solidity: function linkAnyUri((uint256,string,bytes32,bytes) vars) returns()
func (_Contracts *ContractsTransactor) LinkAnyUri(opts *bind.TransactOpts, vars DataTypeslinkAnyUriData) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "linkAnyUri", vars)
}

// LinkAnyUri is a paid mutator transaction binding the contract method 0x5fb88183.
//
// Solidity: function linkAnyUri((uint256,string,bytes32,bytes) vars) returns()
func (_Contracts *ContractsSession) LinkAnyUri(vars DataTypeslinkAnyUriData) (*types.Transaction, error) {
	return _Contracts.Contract.LinkAnyUri(&_Contracts.TransactOpts, vars)
}

// LinkAnyUri is a paid mutator transaction binding the contract method 0x5fb88183.
//
// Solidity: function linkAnyUri((uint256,string,bytes32,bytes) vars) returns()
func (_Contracts *ContractsTransactorSession) LinkAnyUri(vars DataTypeslinkAnyUriData) (*types.Transaction, error) {
	return _Contracts.Contract.LinkAnyUri(&_Contracts.TransactOpts, vars)
}

// LinkCharacter is a paid mutator transaction binding the contract method 0x188b04b3.
//
// Solidity: function linkCharacter((uint256,uint256,bytes32,bytes) vars) returns()
func (_Contracts *ContractsTransactor) LinkCharacter(opts *bind.TransactOpts, vars DataTypeslinkCharacterData) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "linkCharacter", vars)
}

// LinkCharacter is a paid mutator transaction binding the contract method 0x188b04b3.
//
// Solidity: function linkCharacter((uint256,uint256,bytes32,bytes) vars) returns()
func (_Contracts *ContractsSession) LinkCharacter(vars DataTypeslinkCharacterData) (*types.Transaction, error) {
	return _Contracts.Contract.LinkCharacter(&_Contracts.TransactOpts, vars)
}

// LinkCharacter is a paid mutator transaction binding the contract method 0x188b04b3.
//
// Solidity: function linkCharacter((uint256,uint256,bytes32,bytes) vars) returns()
func (_Contracts *ContractsTransactorSession) LinkCharacter(vars DataTypeslinkCharacterData) (*types.Transaction, error) {
	return _Contracts.Contract.LinkCharacter(&_Contracts.TransactOpts, vars)
}

// LinkERC721 is a paid mutator transaction binding the contract method 0xcb8e757e.
//
// Solidity: function linkERC721((uint256,address,uint256,bytes32,bytes) vars) returns()
func (_Contracts *ContractsTransactor) LinkERC721(opts *bind.TransactOpts, vars DataTypeslinkERC721Data) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "linkERC721", vars)
}

// LinkERC721 is a paid mutator transaction binding the contract method 0xcb8e757e.
//
// Solidity: function linkERC721((uint256,address,uint256,bytes32,bytes) vars) returns()
func (_Contracts *ContractsSession) LinkERC721(vars DataTypeslinkERC721Data) (*types.Transaction, error) {
	return _Contracts.Contract.LinkERC721(&_Contracts.TransactOpts, vars)
}

// LinkERC721 is a paid mutator transaction binding the contract method 0xcb8e757e.
//
// Solidity: function linkERC721((uint256,address,uint256,bytes32,bytes) vars) returns()
func (_Contracts *ContractsTransactorSession) LinkERC721(vars DataTypeslinkERC721Data) (*types.Transaction, error) {
	return _Contracts.Contract.LinkERC721(&_Contracts.TransactOpts, vars)
}

// LinkLinklist is a paid mutator transaction binding the contract method 0x9864c307.
//
// Solidity: function linkLinklist((uint256,uint256,bytes32,bytes) vars) returns()
func (_Contracts *ContractsTransactor) LinkLinklist(opts *bind.TransactOpts, vars DataTypeslinkLinklistData) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "linkLinklist", vars)
}

// LinkLinklist is a paid mutator transaction binding the contract method 0x9864c307.
//
// Solidity: function linkLinklist((uint256,uint256,bytes32,bytes) vars) returns()
func (_Contracts *ContractsSession) LinkLinklist(vars DataTypeslinkLinklistData) (*types.Transaction, error) {
	return _Contracts.Contract.LinkLinklist(&_Contracts.TransactOpts, vars)
}

// LinkLinklist is a paid mutator transaction binding the contract method 0x9864c307.
//
// Solidity: function linkLinklist((uint256,uint256,bytes32,bytes) vars) returns()
func (_Contracts *ContractsTransactorSession) LinkLinklist(vars DataTypeslinkLinklistData) (*types.Transaction, error) {
	return _Contracts.Contract.LinkLinklist(&_Contracts.TransactOpts, vars)
}

// LinkNote is a paid mutator transaction binding the contract method 0xb9d32845.
//
// Solidity: function linkNote((uint256,uint256,uint256,bytes32,bytes) vars) returns()
func (_Contracts *ContractsTransactor) LinkNote(opts *bind.TransactOpts, vars DataTypeslinkNoteData) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "linkNote", vars)
}

// LinkNote is a paid mutator transaction binding the contract method 0xb9d32845.
//
// Solidity: function linkNote((uint256,uint256,uint256,bytes32,bytes) vars) returns()
func (_Contracts *ContractsSession) LinkNote(vars DataTypeslinkNoteData) (*types.Transaction, error) {
	return _Contracts.Contract.LinkNote(&_Contracts.TransactOpts, vars)
}

// LinkNote is a paid mutator transaction binding the contract method 0xb9d32845.
//
// Solidity: function linkNote((uint256,uint256,uint256,bytes32,bytes) vars) returns()
func (_Contracts *ContractsTransactorSession) LinkNote(vars DataTypeslinkNoteData) (*types.Transaction, error) {
	return _Contracts.Contract.LinkNote(&_Contracts.TransactOpts, vars)
}

// LockNote is a paid mutator transaction binding the contract method 0x74f345cf.
//
// Solidity: function lockNote(uint256 characterId, uint256 noteId) returns()
func (_Contracts *ContractsTransactor) LockNote(opts *bind.TransactOpts, characterId *big.Int, noteId *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "lockNote", characterId, noteId)
}

// LockNote is a paid mutator transaction binding the contract method 0x74f345cf.
//
// Solidity: function lockNote(uint256 characterId, uint256 noteId) returns()
func (_Contracts *ContractsSession) LockNote(characterId *big.Int, noteId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.LockNote(&_Contracts.TransactOpts, characterId, noteId)
}

// LockNote is a paid mutator transaction binding the contract method 0x74f345cf.
//
// Solidity: function lockNote(uint256 characterId, uint256 noteId) returns()
func (_Contracts *ContractsTransactorSession) LockNote(characterId *big.Int, noteId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.LockNote(&_Contracts.TransactOpts, characterId, noteId)
}

// MintNote is a paid mutator transaction binding the contract method 0xa7ccb4bf.
//
// Solidity: function mintNote((uint256,uint256,address,bytes) vars) returns(uint256 tokenId)
func (_Contracts *ContractsTransactor) MintNote(opts *bind.TransactOpts, vars DataTypesMintNoteData) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "mintNote", vars)
}

// MintNote is a paid mutator transaction binding the contract method 0xa7ccb4bf.
//
// Solidity: function mintNote((uint256,uint256,address,bytes) vars) returns(uint256 tokenId)
func (_Contracts *ContractsSession) MintNote(vars DataTypesMintNoteData) (*types.Transaction, error) {
	return _Contracts.Contract.MintNote(&_Contracts.TransactOpts, vars)
}

// MintNote is a paid mutator transaction binding the contract method 0xa7ccb4bf.
//
// Solidity: function mintNote((uint256,uint256,address,bytes) vars) returns(uint256 tokenId)
func (_Contracts *ContractsTransactorSession) MintNote(vars DataTypesMintNoteData) (*types.Transaction, error) {
	return _Contracts.Contract.MintNote(&_Contracts.TransactOpts, vars)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_Contracts *ContractsTransactor) Multicall(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "multicall", data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_Contracts *ContractsSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _Contracts.Contract.Multicall(&_Contracts.TransactOpts, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_Contracts *ContractsTransactorSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _Contracts.Contract.Multicall(&_Contracts.TransactOpts, data)
}

// PostNote is a paid mutator transaction binding the contract method 0x29c301c2.
//
// Solidity: function postNote((uint256,string,address,bytes,address,bytes,bool) vars) returns(uint256 noteId)
func (_Contracts *ContractsTransactor) PostNote(opts *bind.TransactOpts, vars DataTypesPostNoteData) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "postNote", vars)
}

// PostNote is a paid mutator transaction binding the contract method 0x29c301c2.
//
// Solidity: function postNote((uint256,string,address,bytes,address,bytes,bool) vars) returns(uint256 noteId)
func (_Contracts *ContractsSession) PostNote(vars DataTypesPostNoteData) (*types.Transaction, error) {
	return _Contracts.Contract.PostNote(&_Contracts.TransactOpts, vars)
}

// PostNote is a paid mutator transaction binding the contract method 0x29c301c2.
//
// Solidity: function postNote((uint256,string,address,bytes,address,bytes,bool) vars) returns(uint256 noteId)
func (_Contracts *ContractsTransactorSession) PostNote(vars DataTypesPostNoteData) (*types.Transaction, error) {
	return _Contracts.Contract.PostNote(&_Contracts.TransactOpts, vars)
}

// PostNote4Address is a paid mutator transaction binding the contract method 0x92f7070b.
//
// Solidity: function postNote4Address((uint256,string,address,bytes,address,bytes,bool) vars, address ethAddress) returns(uint256)
func (_Contracts *ContractsTransactor) PostNote4Address(opts *bind.TransactOpts, vars DataTypesPostNoteData, ethAddress common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "postNote4Address", vars, ethAddress)
}

// PostNote4Address is a paid mutator transaction binding the contract method 0x92f7070b.
//
// Solidity: function postNote4Address((uint256,string,address,bytes,address,bytes,bool) vars, address ethAddress) returns(uint256)
func (_Contracts *ContractsSession) PostNote4Address(vars DataTypesPostNoteData, ethAddress common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.PostNote4Address(&_Contracts.TransactOpts, vars, ethAddress)
}

// PostNote4Address is a paid mutator transaction binding the contract method 0x92f7070b.
//
// Solidity: function postNote4Address((uint256,string,address,bytes,address,bytes,bool) vars, address ethAddress) returns(uint256)
func (_Contracts *ContractsTransactorSession) PostNote4Address(vars DataTypesPostNoteData, ethAddress common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.PostNote4Address(&_Contracts.TransactOpts, vars, ethAddress)
}

// PostNote4AnyUri is a paid mutator transaction binding the contract method 0xf316bacd.
//
// Solidity: function postNote4AnyUri((uint256,string,address,bytes,address,bytes,bool) vars, string uri) returns(uint256)
func (_Contracts *ContractsTransactor) PostNote4AnyUri(opts *bind.TransactOpts, vars DataTypesPostNoteData, uri string) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "postNote4AnyUri", vars, uri)
}

// PostNote4AnyUri is a paid mutator transaction binding the contract method 0xf316bacd.
//
// Solidity: function postNote4AnyUri((uint256,string,address,bytes,address,bytes,bool) vars, string uri) returns(uint256)
func (_Contracts *ContractsSession) PostNote4AnyUri(vars DataTypesPostNoteData, uri string) (*types.Transaction, error) {
	return _Contracts.Contract.PostNote4AnyUri(&_Contracts.TransactOpts, vars, uri)
}

// PostNote4AnyUri is a paid mutator transaction binding the contract method 0xf316bacd.
//
// Solidity: function postNote4AnyUri((uint256,string,address,bytes,address,bytes,bool) vars, string uri) returns(uint256)
func (_Contracts *ContractsTransactorSession) PostNote4AnyUri(vars DataTypesPostNoteData, uri string) (*types.Transaction, error) {
	return _Contracts.Contract.PostNote4AnyUri(&_Contracts.TransactOpts, vars, uri)
}

// PostNote4Character is a paid mutator transaction binding the contract method 0xaf90b112.
//
// Solidity: function postNote4Character((uint256,string,address,bytes,address,bytes,bool) vars, uint256 toCharacterId) returns(uint256)
func (_Contracts *ContractsTransactor) PostNote4Character(opts *bind.TransactOpts, vars DataTypesPostNoteData, toCharacterId *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "postNote4Character", vars, toCharacterId)
}

// PostNote4Character is a paid mutator transaction binding the contract method 0xaf90b112.
//
// Solidity: function postNote4Character((uint256,string,address,bytes,address,bytes,bool) vars, uint256 toCharacterId) returns(uint256)
func (_Contracts *ContractsSession) PostNote4Character(vars DataTypesPostNoteData, toCharacterId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.PostNote4Character(&_Contracts.TransactOpts, vars, toCharacterId)
}

// PostNote4Character is a paid mutator transaction binding the contract method 0xaf90b112.
//
// Solidity: function postNote4Character((uint256,string,address,bytes,address,bytes,bool) vars, uint256 toCharacterId) returns(uint256)
func (_Contracts *ContractsTransactorSession) PostNote4Character(vars DataTypesPostNoteData, toCharacterId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.PostNote4Character(&_Contracts.TransactOpts, vars, toCharacterId)
}

// PostNote4ERC721 is a paid mutator transaction binding the contract method 0x327b2a03.
//
// Solidity: function postNote4ERC721((uint256,string,address,bytes,address,bytes,bool) vars, (address,uint256) erc721) returns(uint256)
func (_Contracts *ContractsTransactor) PostNote4ERC721(opts *bind.TransactOpts, vars DataTypesPostNoteData, erc721 DataTypesERC721Struct) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "postNote4ERC721", vars, erc721)
}

// PostNote4ERC721 is a paid mutator transaction binding the contract method 0x327b2a03.
//
// Solidity: function postNote4ERC721((uint256,string,address,bytes,address,bytes,bool) vars, (address,uint256) erc721) returns(uint256)
func (_Contracts *ContractsSession) PostNote4ERC721(vars DataTypesPostNoteData, erc721 DataTypesERC721Struct) (*types.Transaction, error) {
	return _Contracts.Contract.PostNote4ERC721(&_Contracts.TransactOpts, vars, erc721)
}

// PostNote4ERC721 is a paid mutator transaction binding the contract method 0x327b2a03.
//
// Solidity: function postNote4ERC721((uint256,string,address,bytes,address,bytes,bool) vars, (address,uint256) erc721) returns(uint256)
func (_Contracts *ContractsTransactorSession) PostNote4ERC721(vars DataTypesPostNoteData, erc721 DataTypesERC721Struct) (*types.Transaction, error) {
	return _Contracts.Contract.PostNote4ERC721(&_Contracts.TransactOpts, vars, erc721)
}

// PostNote4Linklist is a paid mutator transaction binding the contract method 0x44b82a24.
//
// Solidity: function postNote4Linklist((uint256,string,address,bytes,address,bytes,bool) vars, uint256 toLinklistId) returns(uint256)
func (_Contracts *ContractsTransactor) PostNote4Linklist(opts *bind.TransactOpts, vars DataTypesPostNoteData, toLinklistId *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "postNote4Linklist", vars, toLinklistId)
}

// PostNote4Linklist is a paid mutator transaction binding the contract method 0x44b82a24.
//
// Solidity: function postNote4Linklist((uint256,string,address,bytes,address,bytes,bool) vars, uint256 toLinklistId) returns(uint256)
func (_Contracts *ContractsSession) PostNote4Linklist(vars DataTypesPostNoteData, toLinklistId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.PostNote4Linklist(&_Contracts.TransactOpts, vars, toLinklistId)
}

// PostNote4Linklist is a paid mutator transaction binding the contract method 0x44b82a24.
//
// Solidity: function postNote4Linklist((uint256,string,address,bytes,address,bytes,bool) vars, uint256 toLinklistId) returns(uint256)
func (_Contracts *ContractsTransactorSession) PostNote4Linklist(vars DataTypesPostNoteData, toLinklistId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.PostNote4Linklist(&_Contracts.TransactOpts, vars, toLinklistId)
}

// PostNote4Note is a paid mutator transaction binding the contract method 0x9a4dec18.
//
// Solidity: function postNote4Note((uint256,string,address,bytes,address,bytes,bool) vars, (uint256,uint256) note) returns(uint256)
func (_Contracts *ContractsTransactor) PostNote4Note(opts *bind.TransactOpts, vars DataTypesPostNoteData, note DataTypesNoteStruct) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "postNote4Note", vars, note)
}

// PostNote4Note is a paid mutator transaction binding the contract method 0x9a4dec18.
//
// Solidity: function postNote4Note((uint256,string,address,bytes,address,bytes,bool) vars, (uint256,uint256) note) returns(uint256)
func (_Contracts *ContractsSession) PostNote4Note(vars DataTypesPostNoteData, note DataTypesNoteStruct) (*types.Transaction, error) {
	return _Contracts.Contract.PostNote4Note(&_Contracts.TransactOpts, vars, note)
}

// PostNote4Note is a paid mutator transaction binding the contract method 0x9a4dec18.
//
// Solidity: function postNote4Note((uint256,string,address,bytes,address,bytes,bool) vars, (uint256,uint256) note) returns(uint256)
func (_Contracts *ContractsTransactorSession) PostNote4Note(vars DataTypesPostNoteData, note DataTypesNoteStruct) (*types.Transaction, error) {
	return _Contracts.Contract.PostNote4Note(&_Contracts.TransactOpts, vars, note)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Contracts *ContractsTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Contracts *ContractsSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.SafeTransferFrom(&_Contracts.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Contracts *ContractsTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.SafeTransferFrom(&_Contracts.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Contracts *ContractsTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Contracts *ContractsSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Contracts.Contract.SafeTransferFrom0(&_Contracts.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Contracts *ContractsTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Contracts.Contract.SafeTransferFrom0(&_Contracts.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Contracts *ContractsTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Contracts *ContractsSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Contracts.Contract.SetApprovalForAll(&_Contracts.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Contracts *ContractsTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Contracts.Contract.SetApprovalForAll(&_Contracts.TransactOpts, operator, approved)
}

// SetCharacterUri is a paid mutator transaction binding the contract method 0x47f94de7.
//
// Solidity: function setCharacterUri(uint256 characterId, string newUri) returns()
func (_Contracts *ContractsTransactor) SetCharacterUri(opts *bind.TransactOpts, characterId *big.Int, newUri string) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "setCharacterUri", characterId, newUri)
}

// SetCharacterUri is a paid mutator transaction binding the contract method 0x47f94de7.
//
// Solidity: function setCharacterUri(uint256 characterId, string newUri) returns()
func (_Contracts *ContractsSession) SetCharacterUri(characterId *big.Int, newUri string) (*types.Transaction, error) {
	return _Contracts.Contract.SetCharacterUri(&_Contracts.TransactOpts, characterId, newUri)
}

// SetCharacterUri is a paid mutator transaction binding the contract method 0x47f94de7.
//
// Solidity: function setCharacterUri(uint256 characterId, string newUri) returns()
func (_Contracts *ContractsTransactorSession) SetCharacterUri(characterId *big.Int, newUri string) (*types.Transaction, error) {
	return _Contracts.Contract.SetCharacterUri(&_Contracts.TransactOpts, characterId, newUri)
}

// SetHandle is a paid mutator transaction binding the contract method 0xa6e6178d.
//
// Solidity: function setHandle(uint256 characterId, string newHandle) returns()
func (_Contracts *ContractsTransactor) SetHandle(opts *bind.TransactOpts, characterId *big.Int, newHandle string) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "setHandle", characterId, newHandle)
}

// SetHandle is a paid mutator transaction binding the contract method 0xa6e6178d.
//
// Solidity: function setHandle(uint256 characterId, string newHandle) returns()
func (_Contracts *ContractsSession) SetHandle(characterId *big.Int, newHandle string) (*types.Transaction, error) {
	return _Contracts.Contract.SetHandle(&_Contracts.TransactOpts, characterId, newHandle)
}

// SetHandle is a paid mutator transaction binding the contract method 0xa6e6178d.
//
// Solidity: function setHandle(uint256 characterId, string newHandle) returns()
func (_Contracts *ContractsTransactorSession) SetHandle(characterId *big.Int, newHandle string) (*types.Transaction, error) {
	return _Contracts.Contract.SetHandle(&_Contracts.TransactOpts, characterId, newHandle)
}

// SetLinkModule4Address is a paid mutator transaction binding the contract method 0x08cb68ff.
//
// Solidity: function setLinkModule4Address((address,address,bytes) vars) returns()
func (_Contracts *ContractsTransactor) SetLinkModule4Address(opts *bind.TransactOpts, vars DataTypessetLinkModule4AddressData) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "setLinkModule4Address", vars)
}

// SetLinkModule4Address is a paid mutator transaction binding the contract method 0x08cb68ff.
//
// Solidity: function setLinkModule4Address((address,address,bytes) vars) returns()
func (_Contracts *ContractsSession) SetLinkModule4Address(vars DataTypessetLinkModule4AddressData) (*types.Transaction, error) {
	return _Contracts.Contract.SetLinkModule4Address(&_Contracts.TransactOpts, vars)
}

// SetLinkModule4Address is a paid mutator transaction binding the contract method 0x08cb68ff.
//
// Solidity: function setLinkModule4Address((address,address,bytes) vars) returns()
func (_Contracts *ContractsTransactorSession) SetLinkModule4Address(vars DataTypessetLinkModule4AddressData) (*types.Transaction, error) {
	return _Contracts.Contract.SetLinkModule4Address(&_Contracts.TransactOpts, vars)
}

// SetLinkModule4Linklist is a paid mutator transaction binding the contract method 0x0c4dd5f2.
//
// Solidity: function setLinkModule4Linklist((uint256,address,bytes) vars) returns()
func (_Contracts *ContractsTransactor) SetLinkModule4Linklist(opts *bind.TransactOpts, vars DataTypessetLinkModule4LinklistData) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "setLinkModule4Linklist", vars)
}

// SetLinkModule4Linklist is a paid mutator transaction binding the contract method 0x0c4dd5f2.
//
// Solidity: function setLinkModule4Linklist((uint256,address,bytes) vars) returns()
func (_Contracts *ContractsSession) SetLinkModule4Linklist(vars DataTypessetLinkModule4LinklistData) (*types.Transaction, error) {
	return _Contracts.Contract.SetLinkModule4Linklist(&_Contracts.TransactOpts, vars)
}

// SetLinkModule4Linklist is a paid mutator transaction binding the contract method 0x0c4dd5f2.
//
// Solidity: function setLinkModule4Linklist((uint256,address,bytes) vars) returns()
func (_Contracts *ContractsTransactorSession) SetLinkModule4Linklist(vars DataTypessetLinkModule4LinklistData) (*types.Transaction, error) {
	return _Contracts.Contract.SetLinkModule4Linklist(&_Contracts.TransactOpts, vars)
}

// SetLinklistUri is a paid mutator transaction binding the contract method 0x33f06ee6.
//
// Solidity: function setLinklistUri(uint256 linklistId, string uri) returns()
func (_Contracts *ContractsTransactor) SetLinklistUri(opts *bind.TransactOpts, linklistId *big.Int, uri string) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "setLinklistUri", linklistId, uri)
}

// SetLinklistUri is a paid mutator transaction binding the contract method 0x33f06ee6.
//
// Solidity: function setLinklistUri(uint256 linklistId, string uri) returns()
func (_Contracts *ContractsSession) SetLinklistUri(linklistId *big.Int, uri string) (*types.Transaction, error) {
	return _Contracts.Contract.SetLinklistUri(&_Contracts.TransactOpts, linklistId, uri)
}

// SetLinklistUri is a paid mutator transaction binding the contract method 0x33f06ee6.
//
// Solidity: function setLinklistUri(uint256 linklistId, string uri) returns()
func (_Contracts *ContractsTransactorSession) SetLinklistUri(linklistId *big.Int, uri string) (*types.Transaction, error) {
	return _Contracts.Contract.SetLinklistUri(&_Contracts.TransactOpts, linklistId, uri)
}

// SetMintModule4Note is a paid mutator transaction binding the contract method 0xd23b320b.
//
// Solidity: function setMintModule4Note((uint256,uint256,address,bytes) vars) returns()
func (_Contracts *ContractsTransactor) SetMintModule4Note(opts *bind.TransactOpts, vars DataTypessetMintModule4NoteData) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "setMintModule4Note", vars)
}

// SetMintModule4Note is a paid mutator transaction binding the contract method 0xd23b320b.
//
// Solidity: function setMintModule4Note((uint256,uint256,address,bytes) vars) returns()
func (_Contracts *ContractsSession) SetMintModule4Note(vars DataTypessetMintModule4NoteData) (*types.Transaction, error) {
	return _Contracts.Contract.SetMintModule4Note(&_Contracts.TransactOpts, vars)
}

// SetMintModule4Note is a paid mutator transaction binding the contract method 0xd23b320b.
//
// Solidity: function setMintModule4Note((uint256,uint256,address,bytes) vars) returns()
func (_Contracts *ContractsTransactorSession) SetMintModule4Note(vars DataTypessetMintModule4NoteData) (*types.Transaction, error) {
	return _Contracts.Contract.SetMintModule4Note(&_Contracts.TransactOpts, vars)
}

// SetNoteUri is a paid mutator transaction binding the contract method 0x628b644a.
//
// Solidity: function setNoteUri(uint256 characterId, uint256 noteId, string newUri) returns()
func (_Contracts *ContractsTransactor) SetNoteUri(opts *bind.TransactOpts, characterId *big.Int, noteId *big.Int, newUri string) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "setNoteUri", characterId, noteId, newUri)
}

// SetNoteUri is a paid mutator transaction binding the contract method 0x628b644a.
//
// Solidity: function setNoteUri(uint256 characterId, uint256 noteId, string newUri) returns()
func (_Contracts *ContractsSession) SetNoteUri(characterId *big.Int, noteId *big.Int, newUri string) (*types.Transaction, error) {
	return _Contracts.Contract.SetNoteUri(&_Contracts.TransactOpts, characterId, noteId, newUri)
}

// SetNoteUri is a paid mutator transaction binding the contract method 0x628b644a.
//
// Solidity: function setNoteUri(uint256 characterId, uint256 noteId, string newUri) returns()
func (_Contracts *ContractsTransactorSession) SetNoteUri(characterId *big.Int, noteId *big.Int, newUri string) (*types.Transaction, error) {
	return _Contracts.Contract.SetNoteUri(&_Contracts.TransactOpts, characterId, noteId, newUri)
}

// SetPrimaryCharacterId is a paid mutator transaction binding the contract method 0xf2ad8075.
//
// Solidity: function setPrimaryCharacterId(uint256 characterId) returns()
func (_Contracts *ContractsTransactor) SetPrimaryCharacterId(opts *bind.TransactOpts, characterId *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "setPrimaryCharacterId", characterId)
}

// SetPrimaryCharacterId is a paid mutator transaction binding the contract method 0xf2ad8075.
//
// Solidity: function setPrimaryCharacterId(uint256 characterId) returns()
func (_Contracts *ContractsSession) SetPrimaryCharacterId(characterId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.SetPrimaryCharacterId(&_Contracts.TransactOpts, characterId)
}

// SetPrimaryCharacterId is a paid mutator transaction binding the contract method 0xf2ad8075.
//
// Solidity: function setPrimaryCharacterId(uint256 characterId) returns()
func (_Contracts *ContractsTransactorSession) SetPrimaryCharacterId(characterId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.SetPrimaryCharacterId(&_Contracts.TransactOpts, characterId)
}

// SetSocialToken is a paid mutator transaction binding the contract method 0x95d9fa7d.
//
// Solidity: function setSocialToken(uint256 characterId, address tokenAddress) returns()
func (_Contracts *ContractsTransactor) SetSocialToken(opts *bind.TransactOpts, characterId *big.Int, tokenAddress common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "setSocialToken", characterId, tokenAddress)
}

// SetSocialToken is a paid mutator transaction binding the contract method 0x95d9fa7d.
//
// Solidity: function setSocialToken(uint256 characterId, address tokenAddress) returns()
func (_Contracts *ContractsSession) SetSocialToken(characterId *big.Int, tokenAddress common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.SetSocialToken(&_Contracts.TransactOpts, characterId, tokenAddress)
}

// SetSocialToken is a paid mutator transaction binding the contract method 0x95d9fa7d.
//
// Solidity: function setSocialToken(uint256 characterId, address tokenAddress) returns()
func (_Contracts *ContractsTransactorSession) SetSocialToken(characterId *big.Int, tokenAddress common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.SetSocialToken(&_Contracts.TransactOpts, characterId, tokenAddress)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Contracts *ContractsTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Contracts *ContractsSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.TransferFrom(&_Contracts.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Contracts *ContractsTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.TransferFrom(&_Contracts.TransactOpts, from, to, tokenId)
}

// UnlinkAddress is a paid mutator transaction binding the contract method 0x93f057e5.
//
// Solidity: function unlinkAddress((uint256,address,bytes32) vars) returns()
func (_Contracts *ContractsTransactor) UnlinkAddress(opts *bind.TransactOpts, vars DataTypesunlinkAddressData) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "unlinkAddress", vars)
}

// UnlinkAddress is a paid mutator transaction binding the contract method 0x93f057e5.
//
// Solidity: function unlinkAddress((uint256,address,bytes32) vars) returns()
func (_Contracts *ContractsSession) UnlinkAddress(vars DataTypesunlinkAddressData) (*types.Transaction, error) {
	return _Contracts.Contract.UnlinkAddress(&_Contracts.TransactOpts, vars)
}

// UnlinkAddress is a paid mutator transaction binding the contract method 0x93f057e5.
//
// Solidity: function unlinkAddress((uint256,address,bytes32) vars) returns()
func (_Contracts *ContractsTransactorSession) UnlinkAddress(vars DataTypesunlinkAddressData) (*types.Transaction, error) {
	return _Contracts.Contract.UnlinkAddress(&_Contracts.TransactOpts, vars)
}

// UnlinkAnyUri is a paid mutator transaction binding the contract method 0xef0828ab.
//
// Solidity: function unlinkAnyUri((uint256,string,bytes32) vars) returns()
func (_Contracts *ContractsTransactor) UnlinkAnyUri(opts *bind.TransactOpts, vars DataTypesunlinkAnyUriData) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "unlinkAnyUri", vars)
}

// UnlinkAnyUri is a paid mutator transaction binding the contract method 0xef0828ab.
//
// Solidity: function unlinkAnyUri((uint256,string,bytes32) vars) returns()
func (_Contracts *ContractsSession) UnlinkAnyUri(vars DataTypesunlinkAnyUriData) (*types.Transaction, error) {
	return _Contracts.Contract.UnlinkAnyUri(&_Contracts.TransactOpts, vars)
}

// UnlinkAnyUri is a paid mutator transaction binding the contract method 0xef0828ab.
//
// Solidity: function unlinkAnyUri((uint256,string,bytes32) vars) returns()
func (_Contracts *ContractsTransactorSession) UnlinkAnyUri(vars DataTypesunlinkAnyUriData) (*types.Transaction, error) {
	return _Contracts.Contract.UnlinkAnyUri(&_Contracts.TransactOpts, vars)
}

// UnlinkCharacter is a paid mutator transaction binding the contract method 0x0ff98244.
//
// Solidity: function unlinkCharacter((uint256,uint256,bytes32) vars) returns()
func (_Contracts *ContractsTransactor) UnlinkCharacter(opts *bind.TransactOpts, vars DataTypesunlinkCharacterData) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "unlinkCharacter", vars)
}

// UnlinkCharacter is a paid mutator transaction binding the contract method 0x0ff98244.
//
// Solidity: function unlinkCharacter((uint256,uint256,bytes32) vars) returns()
func (_Contracts *ContractsSession) UnlinkCharacter(vars DataTypesunlinkCharacterData) (*types.Transaction, error) {
	return _Contracts.Contract.UnlinkCharacter(&_Contracts.TransactOpts, vars)
}

// UnlinkCharacter is a paid mutator transaction binding the contract method 0x0ff98244.
//
// Solidity: function unlinkCharacter((uint256,uint256,bytes32) vars) returns()
func (_Contracts *ContractsTransactorSession) UnlinkCharacter(vars DataTypesunlinkCharacterData) (*types.Transaction, error) {
	return _Contracts.Contract.UnlinkCharacter(&_Contracts.TransactOpts, vars)
}

// UnlinkERC721 is a paid mutator transaction binding the contract method 0x867884e6.
//
// Solidity: function unlinkERC721((uint256,address,uint256,bytes32) vars) returns()
func (_Contracts *ContractsTransactor) UnlinkERC721(opts *bind.TransactOpts, vars DataTypesunlinkERC721Data) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "unlinkERC721", vars)
}

// UnlinkERC721 is a paid mutator transaction binding the contract method 0x867884e6.
//
// Solidity: function unlinkERC721((uint256,address,uint256,bytes32) vars) returns()
func (_Contracts *ContractsSession) UnlinkERC721(vars DataTypesunlinkERC721Data) (*types.Transaction, error) {
	return _Contracts.Contract.UnlinkERC721(&_Contracts.TransactOpts, vars)
}

// UnlinkERC721 is a paid mutator transaction binding the contract method 0x867884e6.
//
// Solidity: function unlinkERC721((uint256,address,uint256,bytes32) vars) returns()
func (_Contracts *ContractsTransactorSession) UnlinkERC721(vars DataTypesunlinkERC721Data) (*types.Transaction, error) {
	return _Contracts.Contract.UnlinkERC721(&_Contracts.TransactOpts, vars)
}

// UnlinkLinklist is a paid mutator transaction binding the contract method 0x5a936d10.
//
// Solidity: function unlinkLinklist((uint256,uint256,bytes32) vars) returns()
func (_Contracts *ContractsTransactor) UnlinkLinklist(opts *bind.TransactOpts, vars DataTypesunlinkLinklistData) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "unlinkLinklist", vars)
}

// UnlinkLinklist is a paid mutator transaction binding the contract method 0x5a936d10.
//
// Solidity: function unlinkLinklist((uint256,uint256,bytes32) vars) returns()
func (_Contracts *ContractsSession) UnlinkLinklist(vars DataTypesunlinkLinklistData) (*types.Transaction, error) {
	return _Contracts.Contract.UnlinkLinklist(&_Contracts.TransactOpts, vars)
}

// UnlinkLinklist is a paid mutator transaction binding the contract method 0x5a936d10.
//
// Solidity: function unlinkLinklist((uint256,uint256,bytes32) vars) returns()
func (_Contracts *ContractsTransactorSession) UnlinkLinklist(vars DataTypesunlinkLinklistData) (*types.Transaction, error) {
	return _Contracts.Contract.UnlinkLinklist(&_Contracts.TransactOpts, vars)
}

// UnlinkNote is a paid mutator transaction binding the contract method 0x40ad34d8.
//
// Solidity: function unlinkNote((uint256,uint256,uint256,bytes32) vars) returns()
func (_Contracts *ContractsTransactor) UnlinkNote(opts *bind.TransactOpts, vars DataTypesunlinkNoteData) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "unlinkNote", vars)
}

// UnlinkNote is a paid mutator transaction binding the contract method 0x40ad34d8.
//
// Solidity: function unlinkNote((uint256,uint256,uint256,bytes32) vars) returns()
func (_Contracts *ContractsSession) UnlinkNote(vars DataTypesunlinkNoteData) (*types.Transaction, error) {
	return _Contracts.Contract.UnlinkNote(&_Contracts.TransactOpts, vars)
}

// UnlinkNote is a paid mutator transaction binding the contract method 0x40ad34d8.
//
// Solidity: function unlinkNote((uint256,uint256,uint256,bytes32) vars) returns()
func (_Contracts *ContractsTransactorSession) UnlinkNote(vars DataTypesunlinkNoteData) (*types.Transaction, error) {
	return _Contracts.Contract.UnlinkNote(&_Contracts.TransactOpts, vars)
}

// ContractsApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Contracts contract.
type ContractsApprovalIterator struct {
	Event *ContractsApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsApproval represents a Approval event raised by the Contracts contract.
type ContractsApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Contracts *ContractsFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*ContractsApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractsApprovalIterator{contract: _Contracts.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Contracts *ContractsFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ContractsApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsApproval)
				if err := _Contracts.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Contracts *ContractsFilterer) ParseApproval(log types.Log) (*ContractsApproval, error) {
	event := new(ContractsApproval)
	if err := _Contracts.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Contracts contract.
type ContractsApprovalForAllIterator struct {
	Event *ContractsApprovalForAll // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsApprovalForAll)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsApprovalForAll)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsApprovalForAll represents a ApprovalForAll event raised by the Contracts contract.
type ContractsApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Contracts *ContractsFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*ContractsApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractsApprovalForAllIterator{contract: _Contracts.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Contracts *ContractsFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ContractsApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsApprovalForAll)
				if err := _Contracts.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Contracts *ContractsFilterer) ParseApprovalForAll(log types.Log) (*ContractsApprovalForAll, error) {
	event := new(ContractsApprovalForAll)
	if err := _Contracts.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Contracts contract.
type ContractsInitializedIterator struct {
	Event *ContractsInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsInitialized represents a Initialized event raised by the Contracts contract.
type ContractsInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Contracts *ContractsFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractsInitializedIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractsInitializedIterator{contract: _Contracts.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Contracts *ContractsFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractsInitialized) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsInitialized)
				if err := _Contracts.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Contracts *ContractsFilterer) ParseInitialized(log types.Log) (*ContractsInitialized, error) {
	event := new(ContractsInitialized)
	if err := _Contracts.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Contracts contract.
type ContractsTransferIterator struct {
	Event *ContractsTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsTransfer represents a Transfer event raised by the Contracts contract.
type ContractsTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Contracts *ContractsFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*ContractsTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractsTransferIterator{contract: _Contracts.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Contracts *ContractsFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ContractsTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsTransfer)
				if err := _Contracts.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Contracts *ContractsFilterer) ParseTransfer(log types.Log) (*ContractsTransfer, error) {
	event := new(ContractsTransfer)
	if err := _Contracts.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
