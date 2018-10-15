package demo

import (
	"fmt"

	"github.com/pkg/errors"

	"gopkg.in/urfave/cli.v2"
)

func Enroll() *cli.Command {
	return &cli.Command{
		Name:      "enroll",
		Aliases:   []string{"e"},
		ArgsUsage: "password > record",
		Usage:     "Gets enrollment record for a password",
		Action: func(context *cli.Context) error {
			return enrollFunc(context)
		},
	}
}
func enrollFunc(context *cli.Context) error {

	if context.NArg() < 1 {
		return errors.New("invalid number of arguments")
	}

	token := context.String("token")
	appId := context.String("appid")
	pub := context.String("pk")
	priv := context.String("sk")
	pwd := context.Args().First()

	fmt.Println(token, appId, pub, priv, pwd)
	return nil
}
