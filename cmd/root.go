package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/streamsets/sdk-cli/models"
)

var rootCmd = &cobra.Command{
	Use:   "sx",
	Short: "Streamsets SDK init",
	Long:  `Create Streamsets SDK library and connection base packages for development`,
	Run: func(cmd *cobra.Command, args []string) {
		flags := validateFlags(cmd)

		switch flags.Engine {
		case "datacollector":
			
		case "spark":
			panic("Spark implementation not supported at this time")
		case "snowflake":
			panic("Snowflake implementation not supported at this time")
		}
	},
	Version: "0.1.0",
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.PersistentFlags().String("project", "", "Name of project")
	rootCmd.PersistentFlags().Bool("createLibrary", true, "Create library stub")
	rootCmd.PersistentFlags().Bool("createConnection", false, "Create connection stub")
	rootCmd.PersistentFlags().String("engine", "datacollector", "Engine developed for [datacollector, spark, snowflake]")
}

func validateFlags(cmd *cobra.Command) *models.CliFlags {
	p, err := cmd.Flags().GetString("project")

	if len(p) == 0 || err != nil {
		panic("Error with flag [project]")
	}

	l, err := cmd.Flags().GetBool("createLibrary")
	if err != nil {
		panic(err)
	}

	c, err := cmd.Flags().GetBool("createConnection")
	if err != nil {
		panic(err)
	}

	if !l && !c {
		panic("Both [createConnection] and [createLibrary] flags cannot be false")
	}

	e, err := cmd.Flags().GetString("engine")
	if err != nil {
		panic(err)
	}

	if !validateEngine(e) {
		panic(fmt.Sprintf("Engine %s is not one of valid arguements of %v", e, getSupportedEngines()))
	}

	return &models.CliFlags{
		Project:          p,
		CreateLibrary:    l,
		CreateConnection: c,
		Engine:           e,
	}
}

func getSupportedEngines() []string {
	return []string{"datacollector", "spark", "snowflake"}
}

func validateEngine(engine string) bool {
	for _, e := range getSupportedEngines() {
		if strings.Compare(strings.ToLower(engine), e) == 0 {
			return true
		}
	}

	return false
}
