package main

import (
	"github.com/golang-module/carbon"
	"modarak.net/subcommand/internal/commands"
	"modarak.net/subcommand/internal/med"
)

func main() {

	// var meds [2]*med.Med = [2]*med.Med{
	// 	med.NewMed("prami", 293+100, carbon.Parse("2023-10-18"), 2, 14),
	// 	med.NewMed("levo", 47+5+200, carbon.Parse("2023-10-18"), 3, 14),
	// }
	// meds[0].Dump()
	p := med.NewMed("prami", 336, carbon.Parse("2023-11-05"), 2, 14)
	l := med.NewMed("levo", 198, carbon.Parse("2023-11-05"), 2, 14)
	p.Dump()
	l.Dump()

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
