package commands

import (
	"fmt"
	"uhppote"
)

type SetDoorDelayCommand struct {
}

func (c *SetDoorDelayCommand) Execute(u *uhppote.UHPPOTE) error {
	serialNumber, err := getUint32(1, "Missing serial number", "Invalid serial number: %v")
	if err != nil {
		return err
	}

	door, err := getDoor(2, "Missing door", "Invalid door: %v")
	if err != nil {
		return err
	}

	delay, err := getUint8(3, "Missing delay", "Invalid delay: %v")
	if err != nil {
		return err
	}

	record, err := u.SetDoorDelay(serialNumber, door, delay)
	if err != nil {
		return err
	}

	fmt.Printf("%v\n", record)

	return nil
}

func (c *SetDoorDelayCommand) CLI() string {
	return "set-door-delay"
}

func (c *SetDoorDelayCommand) Help() {
	fmt.Println("Usage: uhppote-cli [options] set-door-delay <serial number> <door> <delay>")
	fmt.Println()
	fmt.Println(" Sets the door open delay (in seconds)")
	fmt.Println()
	fmt.Println("  serial-number  (required) controller serial number")
	fmt.Println("  door           (required) door (1,2,3 or 4")
	fmt.Println("  delay          (required) delay in seconds")
	fmt.Println()
	fmt.Println("  Examples:")
	fmt.Println()
	fmt.Println("    uhppote-cli set-door-delay 12345678 3 15")
	fmt.Println()
}