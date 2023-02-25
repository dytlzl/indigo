package cmdutil

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func FlagSetBuilder() *flagSetBuilder {
	return &flagSetBuilder{
		flags: pflag.NewFlagSet("", 1),
	}
}

type flagSetBuilder struct {
	flags *pflag.FlagSet
}

func MarkFlagFilename(extensions ...string) func(flags *pflag.FlagSet, name string) error {
	return func(flags *pflag.FlagSet, name string) error {
		return cobra.MarkFlagFilename(flags, name, extensions...)
	}
}

func (d *flagSetBuilder) String(p *string, name, shorthand string, value string, usage string, markFns ...func(flags *pflag.FlagSet, name string) error) *flagSetBuilder {
	d.flags.StringVarP(p, name, shorthand, value, usage)
	for _, mf := range markFns {
		if err := mf(d.flags, name); err != nil {
			log.Fatalln(err)
		}
	}
	return d
}

func (d *flagSetBuilder) Int(p *int, name, shorthand string, value int, usage string, markFns ...func(flags *pflag.FlagSet, name string) error) *flagSetBuilder {
	d.flags.IntVarP(p, name, shorthand, value, usage)
	for _, mf := range markFns {
		if err := mf(d.flags, name); err != nil {
			log.Fatalln(err)
		}
	}
	return d
}

func (d *flagSetBuilder) Build() *pflag.FlagSet {
	return d.flags
}
