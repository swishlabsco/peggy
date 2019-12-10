package main

// 	Main (ebrelayer) : Implements CLI commands for the Relayer
//		service, such as initialization and event relay.

import (
	"errors"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/ethereum/go-ethereum/common"

	amino "github.com/tendermint/go-amino"
	"github.com/tendermint/tendermint/libs/cli"

	"github.com/cosmos/cosmos-sdk/client"
	sdkContext "github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/client/rpc"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtxb "github.com/cosmos/cosmos-sdk/x/auth/types"

	app "github.com/cosmos/peggy/app"
	contract "github.com/cosmos/peggy/cmd/ebrelayer/contract"
	relayer "github.com/cosmos/peggy/cmd/ebrelayer/relayer"
)

var appCodec *amino.Codec

// FlagRPCURL : flag for the RPC URL of the tendermint node
const FlagRPCURL = "rpc-url"

// FlagMakeClaims : optional flag for the ethereum relayer to automatically make OracleClaims upon every ProphecyClaim
const FlagMakeClaims = "make-claims"

func init() {

	// Read in the configuration file for the sdk
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(sdk.Bech32PrefixAccAddr, sdk.Bech32PrefixAccPub)
	config.SetBech32PrefixForValidator(sdk.Bech32PrefixValAddr, sdk.Bech32PrefixValPub)
	config.SetBech32PrefixForConsensusNode(sdk.Bech32PrefixConsAddr, sdk.Bech32PrefixConsPub)
	config.Seal()

	appCodec = app.MakeCodec()

	DefaultCLIHome := os.ExpandEnv("$HOME/.ebcli")

	// Add --chain-id to persistent flags and mark it required
	rootCmd.PersistentFlags().String(client.FlagChainID, "", "Chain ID of tendermint node")
	rootCmd.PersistentFlags().String(FlagRPCURL, "", "RPC URL of tendermint node")
	rootCmd.PersistentPreRunE = func(_ *cobra.Command, _ []string) error {
		return initConfig(rootCmd)
	}

	// Add --make-claims to init cmd as optional flag
	initCmd.PersistentFlags().String(FlagMakeClaims, "", "Make oracle claims everytime a prophecy claim is witnessed")

	// Construct Initialization Commands
	initCmd.AddCommand(
		ethereumRelayerCmd(),
		client.LineBreak,
		cosmosRelayerCmd(),
	)

	// Construct Root Command
	rootCmd.AddCommand(
		rpc.StatusCommand(),
		initCmd,
		generateBindingsCmd(),
	)

	executor := cli.PrepareMainCmd(rootCmd, "EBRELAYER", DefaultCLIHome)
	err := executor.Execute()
	if err != nil {
		log.Fatal("failed executing CLI command", err)
	}
}

var rootCmd = &cobra.Command{
	Use:          "ebrelayer",
	Short:        "Relayer service which listens for and relays ethereum smart contract events",
	SilenceUsage: true,
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialization subcommands",
}

//	ethereumRelayerCmd : Initializes a relayer service run by individual
//		validators which streams live events from an Ethereum smart contract.
//		The service automatically signs messages containing the event
//		data and relays them to tendermint for handling by the
//		EthBridge module.
//
func ethereumRelayerCmd() *cobra.Command {
	ethereumRelayerCmd := &cobra.Command{
		Use:     "ethereum [web3Provider] [bridgeContractAddress] [validatorFromName] --make-claims [make-claims] --chain-id [chain-id]",
		Short:   "Initializes a web socket which streams live events from a smart contract and relays them to the Cosmos network",
		Args:    cobra.ExactArgs(3),
		Example: "ebrelayer init ethereum wss://ropsten.infura.io/ws 05d9758cb6b9d9761ecb8b2b48be7873efae15c0 validator --make-claims=false --chain-id=testing",
		RunE:    RunEthereumRelayerCmd,
	}

	return ethereumRelayerCmd
}

//	cosmosRelayerCmd : Initializes a Cosmos relayer service run by individual
//		validators which streams live events from the Cosmos network and then
//		relaying them to an Ethereum smart contract
//
func cosmosRelayerCmd() *cobra.Command {
	cosmosRelayerCmd := &cobra.Command{
		Use:     "cosmos [tendermintNode] [web3Provider] [bridgeContractAddress]",
		Short:   "Initializes a web socket which streams live events from the Cosmos network and relays them to the Ethereum network",
		Args:    cobra.ExactArgs(3),
		Example: "ebrelayer init cosmos tcp://localhost:26657 http://localhost:7545 0xd88159878c50e4B2b03BB701DD436e4A98D6fBe2",
		RunE:    RunCosmosRelayerCmd,
	}

	return cosmosRelayerCmd
}

//	generateBindingsCmd : Generates ABIs and bindings for Bridge smart contracts which facilitate contract interaction
func generateBindingsCmd() *cobra.Command {
	generateBindingsCmd := &cobra.Command{
		Use:     "generate",
		Short:   "Generates Bridge smart contracts ABIs and bindings",
		Args:    cobra.ExactArgs(0),
		Example: "generate",
		RunE:    RunGenerateBindingsCmd,
	}

	return generateBindingsCmd
}

// RunEthereumRelayerCmd executes the initEthereumRelayerCmd with the provided parameters
func RunEthereumRelayerCmd(cmd *cobra.Command, args []string) error {
	// Parse chain's ID
	chainID := viper.GetString(client.FlagChainID)
	if strings.TrimSpace(chainID) == "" {
		return errors.New("Must specify a 'chain-id'")
	}

	// Parse make claims boolean
	var makeClaims bool

	makeClaimsString := viper.GetString(FlagMakeClaims)
	if strings.TrimSpace(makeClaimsString) == "true" {
		makeClaims = true
	} else {
		makeClaims = false
	}

	// Parse ethereum provider
	ethereumProvider := args[0]
	if !relayer.IsWebsocketURL(ethereumProvider) {
		return fmt.Errorf("invalid [web3-provider]: %s", ethereumProvider)
	}

	// Parse the address of the deployed contract
	if !common.IsHexAddress(args[1]) {
		return fmt.Errorf("invalid [bridge-contract-address]: %s", args[1])
	}
	contractAddress := common.HexToAddress(args[1])

	// Parse the validator's moniker
	validatorFrom := args[2]

	// Parse Tendermint RPC URL
	rpcURL := viper.GetString(FlagRPCURL)

	if rpcURL != "" {
		_, err := url.Parse(rpcURL)
		if rpcURL != "" && err != nil {
			return fmt.Errorf("invalid RPC URL: %v", rpcURL)
		}
	}

	// Get the validator's name and account address using their moniker
	validatorAccAddress, validatorName, err := sdkContext.GetFromFields(validatorFrom, false)
	if err != nil {
		return err
	}
	// Convert the validator's account address into type ValAddress
	validatorAddress := sdk.ValAddress(validatorAccAddress)

	// Get the validator's passphrase using their moniker
	passphrase, err := keys.GetPassphrase(validatorFrom)
	if err != nil {
		return err
	}

	// Test passphrase is correct
	_, err = authtxb.MakeSignature(nil, validatorName, passphrase, authtxb.StdSignMsg{})
	if err != nil {
		return err
	}

	// Initialize the relayer
	return relayer.InitEthereumRelayer(
		appCodec,
		chainID,
		ethereumProvider,
		contractAddress,
		makeClaims,
		validatorName,
		passphrase,
		validatorAddress,
		rpcURL,
	)
}

// RunCosmosRelayerCmd executes the initCosmosRelayerCmd with the provided parameters
func RunCosmosRelayerCmd(cmd *cobra.Command, args []string) error {
	// Load config file containing environment variables
	err := godotenv.Load()
	if err != nil {
		return errors.New("Error loading .env file")
	}

	// Private key for validator's Ethereum address must be set as an environment variable
	privateKey := os.Getenv("ETHEREUM_PRIVATE_KEY")
	if strings.TrimSpace(privateKey) == "" {
		return errors.New("Error loading validator's private key from .env file")
	}

	// Tendermint node
	tendermintNode := args[0]

	// Ethereum websocket provider
	ethereumProvider := args[1]

	// Deployed contract address
	if !common.IsHexAddress(args[2]) {
		return fmt.Errorf("invalid [bridge-contract-address]: %s", args[2])
	}
	contractAddress := common.HexToAddress(args[2])

	// Initialize the relayer
	return relayer.InitCosmosRelayer(
		tendermintNode,
		ethereumProvider,
		contractAddress,
		privateKey,
	)
}

// RunGenerateBindingsCmd : executes the generateBindingsCmd
func RunGenerateBindingsCmd(cmd *cobra.Command, args []string) error {
	contracts := []contract.Type{
		contract.BridgeRegistry,
		contract.Valset,
		contract.Oracle,
		contract.CosmosBridge,
		contract.BridgeBank,
	}

	// Compile contracts, generating contract bin and abi
	err := contract.CompileContracts(contracts)
	if err != nil {
		return err
	}

	// Generate contract bindings
	return contract.GenerateBindings(contracts)
}

func initConfig(cmd *cobra.Command) error {
	return viper.BindPFlag(client.FlagChainID, cmd.PersistentFlags().Lookup(client.FlagChainID))
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
