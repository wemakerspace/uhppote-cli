package commands

import (
	"fmt"
	"uhppote"
)

type RevokeAllCommand struct {
	SerialNumber uint32
}

func NewRevokeAllCommand() (*RevokeAllCommand, error) {
	serialNumber, err := getUint32(1, "Missing serial number", "Invalid serial number: %v")
	if err != nil {
		return nil, err
	}

	return &RevokeAllCommand{serialNumber}, nil
}

func (c *RevokeAllCommand) Execute(u *uhppote.UHPPOTE) error {
	deleted, err := u.DeleteAll(c.SerialNumber)

	if err == nil {
		fmt.Printf("%v\n", deleted)
	}

	return err
}

func (c *RevokeAllCommand) CLI() string {
	return "revoke-all"
}

func (c *RevokeAllCommand) Help() {
	fmt.Println("Usage: uhppote-cli [options] revoke-all <serial number>")
	fmt.Println()
	fmt.Println(" Removes all cards from the authorised list")
	fmt.Println()
	fmt.Println("  <serial number>  (required) controller serial number")
	fmt.Println()
	fmt.Println("  Examples:")
	fmt.Println()
	fmt.Println("    uhppote-cli revoke-all 12345678")
	fmt.Println()
}
