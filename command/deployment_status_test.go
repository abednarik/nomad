package command

import (
	"strings"
	"testing"

	"github.com/mitchellh/cli"
)

func TestDeploymentStatusCommand_Implements(t *testing.T) {
	var _ cli.Command = &DeploymentStatusCommand{}
}

func TestDeploymentStatusCommand_Fails(t *testing.T) {
	ui := new(cli.MockUi)
	cmd := &DeploymentStatusCommand{Meta: Meta{Ui: ui}}

	// Fails on misuse
	if code := cmd.Run([]string{"some", "bad", "args"}); code != 1 {
		t.Fatalf("expected exit code 1, got: %d", code)
	}
	if out := ui.ErrorWriter.String(); !strings.Contains(out, cmd.Help()) {
		t.Fatalf("expected help output, got: %s", out)
	}
	ui.ErrorWriter.Reset()

	if code := cmd.Run([]string{"-address=nope", "12"}); code != 1 {
		t.Fatalf("expected exit code 1, got: %d", code)
	}
	if out := ui.ErrorWriter.String(); !strings.Contains(out, "Error retrieving deployment") {
		t.Fatalf("expected failed query error, got: %s", out)
	}
	ui.ErrorWriter.Reset()
}
