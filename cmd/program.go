package cmd

import (
	"github.com/urfave/cli"
)

const Version = "1.0.0"

func NewApp() *cli.App {
	app := cli.NewApp()

	app.Version = Version
	app.Description =
		"Command-line utility which allows to encrypt and decrypt files by one-time pad (Vernam) algorithm"

	app.Commands = []cli.Command{
		{
			Name: "keygen",
			Aliases: []string{"k"},
			Description: "Generate key",
			ArgsUsage: "<length> <filename>",
			Action: actionKeygen,
		},
		{
			Name: "encrypt",
			Aliases: []string{"e"},
			Description: "Encrypt file by key",
			ArgsUsage: "<filename> <keyfile> <target_file>",
			Action: actionEncrypt,
		},
		{
			Name: "decrypt",
			Aliases: []string{"d"},
			Description: "Decrypt file by key",
			ArgsUsage: "<filename> <keyfile> <target_file>",
			Action: actionDecrypt,
		},
	}

	return app
}
