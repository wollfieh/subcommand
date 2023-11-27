package main

import (
	"github.com/golang-module/carbon"
	"modarak.net/subcommand/internal/commands"
	"modarak.net/subcommand/internal/med"
)

func main() {

	var meds [2]*med.Med = [2]*med.Med{
		med.NewMed("prami", 293+100, carbon.Parse("2023-10-18"), 2, 14),
		med.NewMed("levo", 47+5+200, carbon.Parse("2023-10-18"), 3, 14),
	}
	for _, v := range meds {
		v.Dump()

	}
	// meds[0].Dump()
	// m := med.NewMed("prami", 293+100, carbon.Parse("2023-10-18"), 2, 14)
	// m.Dump()

	commands.Initialize()

	// TODO:
	// m (typ Med) hat verschiedene Methoden . diese sollten nun mit subcommand
	// registriert werden
	// die sollten vom type subcommand.Command sein
	// m sollte ein member []commands haben und commands.Execute sollte ein ref zur methode sein

	// TODO:
	// m sollte nicht type Med sein sondern möglicherweise ein Container der mehrere Meds
	// enthält?

	commands.DoExecute()

}
