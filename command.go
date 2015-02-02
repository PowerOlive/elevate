// package escalate provides support for executing commands with escalated
// privileges.
package escalate

import (
	"os/exec"
)

// Command is like exec.Command, except that it runs the given command with
// escalated privileges.
func Command(name string, args ...string) *exec.Cmd {
	cmd, err := buildCommand("", "", name, args...)
	if err != nil {
		panic(err)
	}
	return cmd
}

// Prompt is like Command, except the escalation prompt contains the custom
// prompt string.
func Prompt(prompt string, name string, args ...string) *exec.Cmd {
	cmd, err := buildCommand(prompt, "", name, args...)
	if err != nil {
		panic(err)
	}
	return cmd
}

// PromptWithIcon is like Prompt, except that hte escalation prompt also
// includes an icon loaded from the given path.
func PromptWithIcon(prompt string, icon string, name string, args ...string) *exec.Cmd {
	cmd, err := buildCommand(prompt, icon, name, args...)
	if err != nil {
		panic(err)
	}
	return cmd
}
