package contract

import (
	"fmt"
	"os/exec"
	"strings"
)

const (
	// _standardABIBINGenCmd : standard command for contract ABI and BIN generation
	_standardABIBINGenCmd = "solc " +
		"--[SOLC_CMD] ./testnet-contracts/contracts/[DIRECTORY][CONTRACT].sol " +
		"-o ./cmd/ebrelayer/generated/[SOLC_CMD]/[CONTRACT] " +
		"--overwrite " +
		"--allow-paths *,"
	// _standardBindingGenCmd : standard command for contract binding generation
	_standardBindingGenCmd = "abigen " +
		"--bin ./cmd/ebrelayer/generated/bin/[CONTRACT]/[CONTRACT].bin " +
		"--abi ./cmd/ebrelayer/generated/abi/[CONTRACT]/[CONTRACT].abi " +
		"--pkg [CONTRACT] " +
		"--type [CONTRACT] " +
		"--out ./cmd/ebrelayer/generated/bindings/[CONTRACT]/[CONTRACT].go"
)

const (
	_solcCmd   = "[SOLC_CMD]"
	_directory = "[DIRECTORY]"
	_contract  = "[CONTRACT]"
)

// CompileContracts : compiles named contracts and generates ABI, BIN
func CompileContracts(contracts []Type) error {
	for _, contract := range contracts {
		// Validate contract name
		err := contract.IsValid()
		if err != nil {
			return err
		}

		err = compileContract(contract)
		if err != nil {
			return err
		}
	}
	fmt.Println("Successfully generated all bins and abis.")
	return nil
}

// compileContract : compiles a contract, generating both BIN and BIN files
func compileContract(contract Type) error {
	contractBINGenCmd, contractABIGenCmd := prepareContractABICmdBINCmd(contract)

	// Generate Contract.bin file
	err := execCmd(contractBINGenCmd)
	if err != nil {
		return err
	}

	// Generate Contract.abi file
	return execCmd(contractABIGenCmd)
}

// prepareContractABICmdBINCmd : prepares the BIN and ABI generation cmds for a contract via text replacement
func prepareContractABICmdBINCmd(contract Type) (string, string) {
	// Replace directory with appropriate directory path
	directoryReplacementText := ""
	if contract == BridgeBank {
		directoryReplacementText = contract.String() + "/"
	}
	directoryABIBINGenCmd := replaceText(_standardABIBINGenCmd, _directory, directoryReplacementText)

	// Populate contract name
	contractABIBINGenCmd := replaceText(directoryABIBINGenCmd, _contract, contract.String())

	// Prepare BIN and ABI generation cmds
	contractBINGenCmd := replaceText(contractABIBINGenCmd, _solcCmd, "bin")
	contractABIGenCmd := replaceText(contractABIBINGenCmd, _solcCmd, "abi")

	return contractBINGenCmd, contractABIGenCmd
}

// GenerateBindings : generates bindings for the named contracts
func GenerateBindings(contracts []Type) error {
	for _, contract := range contracts {
		// Validate contract name
		err := contract.IsValid()
		if err != nil {
			return err
		}

		err = generateBinding(contract)
		if err != nil {
			return err
		}
	}
	fmt.Println("Successfully generated all bindings.")
	return nil
}

// generateBinding : generates bindings for a single contract
func generateBinding(contract Type) error {
	bindingGenerationCmd := replaceText(_standardBindingGenCmd, _contract, contract.String())
	// Generate Contract.go bindings file
	return execCmd(bindingGenerationCmd)
}

// replaceText : replaces specific current text with new text
func replaceText(cmd string, current string, new string) string {
	return strings.Replace(cmd, current, new, -1)
}

// execCmd : executes the given cmd
func execCmd(cmd string) error {
	_, err := exec.Command("sh", "-c", cmd).Output()
	return err
}
