package cmd

/*
Copyright © 2023 Andrew Demyanenko a.demyanenko@corp.vk.com
*/

import (
	"fmt"
	"github.com/spf13/viper"
	"jira-cli/internal/jira"
	"os"

	"github.com/spf13/cobra"
)

var (
	cfgFile string

	listCmd = &cobra.Command{
		Use:   "list",
		Short: "A list of issues",
		Run: func(cmd *cobra.Command, args []string) {
			url := viper.GetString("url")
			token := viper.GetString("token")
			if url == "" || token == "" {
				cobra.CheckErr(fmt.Errorf("JIRA_URL and JIRA_TOKEN are required"))
			}

			jiraConnector := jira.NewJiraConnector(url, token)

			repo := jira.NewRepo(jiraConnector)

			project := viper.GetString("project")
			sprint := viper.GetString("sprint")
			order := viper.GetString("order")
			assignee := viper.GetString("assignee")

			query := getSearchIssueJql(project, assignee, sprint, order)

			issues, err := repo.SearchByJQL(query)
			cobra.CheckErr(err)

			for _, issue := range *issues {
				fmt.Println(issue.String())
				fmt.Println("----")
			}
		},
	}
)

func getSearchIssueJql(project string, assignee string, sprint string, order string) *jira.Jql {
	jql := jira.NewJql()

	if project != "" {
		jql = jql.ByProjects([]string{project})
	}

	if assignee != "" {
		jql = jql.ByAssignee([]string{assignee})
	}

	if sprint != "" {
		jql = jql.BySprints([]string{sprint})
	}

	if order != "" {
		jql = jql.OrderBy([]string{order}).Join("ASC", " ")
	}
	return jql
}

func Execute() {
	err := listCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	listCmd.PersistentFlags().StringVar(&cfgFile, "config", ".env", "CLI config")

	// connector settings
	listCmd.PersistentFlags().String("url", "", "A Jira server url")
	listCmd.PersistentFlags().String("token", "", "Access token")

	// filters
	listCmd.PersistentFlags().String("project", "INTDEV", "Project")
	listCmd.PersistentFlags().String("assignee", "a.demyanenko", "Assignee")
	listCmd.PersistentFlags().String("reporter", "", "Reporter")
	listCmd.PersistentFlags().String("sprint", "openSprints()", "Sprint. \n openSprints() - by default")
	listCmd.PersistentFlags().String("status", "Open", "Status. \n Open - by default")
	listCmd.PersistentFlags().String("order", "status", "Status. \n Open - by default")

	viper.BindPFlag("token", listCmd.PersistentFlags().Lookup("token"))
	viper.BindPFlag("url", listCmd.PersistentFlags().Lookup("url"))

	viper.BindPFlag("project", listCmd.PersistentFlags().Lookup("project"))
	viper.BindPFlag("assignee", listCmd.PersistentFlags().Lookup("assignee"))
	viper.BindPFlag("reporter", listCmd.PersistentFlags().Lookup("reporter"))
	viper.BindPFlag("sprint", listCmd.PersistentFlags().Lookup("sprint"))
	viper.BindPFlag("status", listCmd.PersistentFlags().Lookup("status"))
	viper.BindPFlag("order", listCmd.PersistentFlags().Lookup("order"))
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile("./.env")
		if err := viper.ReadInConfig(); err != nil {
			cobra.CheckErr(err)
		}
	}
	viper.AutomaticEnv()
	viper.SetEnvPrefix("JIRA")
}
