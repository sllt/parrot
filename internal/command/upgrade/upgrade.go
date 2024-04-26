package upgrade

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/sllt/parrot/internal/config"
	"github.com/spf13/cobra"
)

var CmdUpgrade = &cobra.Command{
	Use:     "upgrade",
	Short:   "Upgrade the parrot command.",
	Long:    "Upgrade the parrot command.",
	Example: "parrot upgrade",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("go install %s\n", config.ParrotCmd)
		cmd := exec.Command("go", "install", config.ParrotCmd)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			log.Fatalf("go install %s error\n", err)
		}
		fmt.Printf("\n🎉 parrot upgrade successfully!\n\n")
	},
}
