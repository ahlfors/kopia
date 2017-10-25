package cli

import (
	"time"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	blockRepackCommand       = blockCommands.Command("repack", "Repackage small blocks into bigger ones")
	blockRepackGroup         = blockRepackCommand.Flag("group", "Group to repack").Default("DIR").String()
	blockRepackSizeThreshold = blockRepackCommand.Flag("max-size", "Max size of block to re-pack").Default("500000").Int64()
	blockRepackMinAge        = blockRepackCommand.Flag("min-age", "Minimum age to repack").Default("24h").Duration()
)

func runBlockRepackAction(context *kingpin.ParseContext) error {
	rep := mustOpenRepository(nil)
	defer rep.Close()

	if err := rep.Blocks.Repackage(*blockRepackGroup, *blockRepackSizeThreshold, time.Now().Add(-*blockRepackMinAge)); err != nil {
		return err
	}

	return nil
}

func init() {
	blockRepackCommand.Action(runBlockRepackAction)
}