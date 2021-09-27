package overhaul

import "fmt"

func Check_msg(name string) (string, error) {
	// Return a greeting that embeds the name in a msg.
	msg := fmt.Sprintf("Hi, %v. Welcome!", name)
	return msg, nil
}
