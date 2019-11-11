package contract

import (
	"fmt"
	"os/exec"
	"strings"
)

// StandardABIGenCmd : standard contract ABI generation command
const StandardABIGenCmd = "solc " +
	"--abi ./testnet-contracts/contracts/[DIRECTORY][CONTRACT].sol " +
	"-o ./cmd/ebrelayer/generated/abi/[CONTRACT] " +
	"--overwrite " +
	"--allow-paths *,"

// StandardBindingGenCmd : standard contract binding generation command
const StandardBindingGenCmd = "abigen " +
	"--abi ./cmd/ebrelayer/generated/abi/[CONTRACT]/[CONTRACT].abi " +
	"--pkg [CONTRACT] " +
	"--type [CONTRACT] " +
	"--out ./cmd/ebrelayer/generated/bindings/[CONTRACT]/[CONTRACT].go"

// GenerateContractABIs : generates ABI for the named contracts
func GenerateContractABIs(bridgeContracts []string) error {
	for i := 0; i < len(bridgeContracts); i++ {
		// Set up directory replacement text
		var directoryReplacementText string
		if bridgeContracts[i] == "BridgeBank" {
			directoryReplacementText = bridgeContracts[i] + "/"
		} else {
			directoryReplacementText = ""
		}
		// Populate the contract abi's directory field
		standardABIGenCmd := strings.Replace(StandardABIGenCmd, "[DIRECTORY]", directoryReplacementText, -1)
		// Populate contract name into cmd
		abiGenerationCmd := replaceContractName(standardABIGenCmd, bridgeContracts[i])
		_, err := exec.Command("sh", "-c", abiGenerationCmd).Output()
		if err != nil {
			return err
		}
	}
	fmt.Println("Successfully generated contract ABIs.")
	return nil
}

// GenerateBindings : generates bindings for the named contracts
func GenerateBindings(bridgeContracts []string) error {
	for i := 0; i < len(bridgeContracts); i++ {
		// Populate contract name into cmd
		bindingGenerationCmd := replaceContractName(StandardBindingGenCmd, bridgeContracts[i])
		// Execute cmd
		_, err := exec.Command("sh", "-c", bindingGenerationCmd).Output()
		if err != nil {
			return err
		}
	}
	fmt.Println("Successfully generated contract bindings.")
	return nil
}

// replaceContractName : replaces placeholder text "[CONTRACT]" with the contract's name
func replaceContractName(cmd string, name string) string {
	return strings.Replace(cmd, "[CONTRACT]", name, -1)
}
