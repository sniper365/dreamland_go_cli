package kill

import (
	"github.com/taubyte/dreamland/cli/command"
	"github.com/taubyte/dreamland/cli/common"
	client "github.com/taubyte/dreamland/service"
	"github.com/urfave/cli/v2"
)

func universe(multiverse *client.Client) *cli.Command {
	c := &cli.Command{
		Name:   "universe",
		Action: killUniverse(multiverse),
	}

	command.NameWithDefault(c, common.DefaultUniverseName)

	return c
}

func killUniverse(multiverse *client.Client) cli.ActionFunc {
	/**
	@hope to have call
	common days 7pm-10pm
	sunday 4pm-10pm
	let me know your phone number
	*/
	return func(c *cli.Context) (err error) {
		return multiverse.Universe(c.String("name")).Kill()
	}
}
