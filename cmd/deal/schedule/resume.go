package schedule

import (
	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler/deal/schedule"
	"github.com/urfave/cli/v2"
)

var ResumeCmd = &cli.Command{
	Name:      "resume",
	Usage:     "Resume a specific schedule",
	ArgsUsage: "SCHEDULE_ID",
	Action: func(c *cli.Context) error {
		db, err := database.OpenFromCLI(c)
		if err != nil {
			return err
		}
		schedule, err := schedule.ResumeHandler(db, c.Args().Get(0))
		if err != nil {
			return err
		}
		cliutil.PrintToConsole(schedule, c.Bool("json"), nil)
		return nil
	},
}
