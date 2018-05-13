package main

import (
	"log"
	"os/exec"
)

// Execute OS command.
// ----------------------------
// - single command
// - joined command
func main() {
	commands := []string{"date", "whoami", "date ; whoami", "date & whoami"}

	for _, command := range commands {

		output, error := exec.Command(command).Output()

		if error != nil {
			log.Printf("[ERROR] %s", error)
			continue
		}

		log.Printf("$ %s -> %s", command, output)
	}
}
