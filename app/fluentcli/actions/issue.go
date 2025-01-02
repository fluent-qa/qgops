package actions

import (
	"github.com/fluent-qa/qgops/pkg/utils"
	"github.com/fluent-qa/qgops/pkg/utils/shell"
	_ "log/slog"

	"github.com/spf13/cobra"
)

func runGhIssueCreate(title string, body string, doc string) (int, error) {
	template := "gh issue create --label enhancement -t \"{{.Title}}\" -b \"{{.Body}}\""
	if doc == "" {
		if body == "" {
			body = title
		}
	} else {
		body, _ = utils.ReadStringFromFile(doc)
	}

	commandStr := utils.RenderTemplate(template, map[string]string{
		"Title": title,
		"Body":  body,
	})
	return shell.ExecShellCommand(commandStr)
}

func init() {
	IssueCmd.Flags().StringVarP(&issueTitle, "title", "t", "", "Title of the issue")
	IssueCmd.Flags().StringVarP(&issueBody, "body", "b", "", "Body of the issue")
	IssueCmd.Flags().StringVarP(&issueDoc, "doc", "d", "", "Body do of the issue")
	_ = IssueCmd.MarkFlagRequired("title")
	//IssueCmd.MarkFlagRequired("body")
}

var (
	issueTitle string
	issueBody  string
	issueDoc   string
	IssueCmd   = &cobra.Command{
		Use:   "issue",
		Short: "Create a new issue",
		Long:  "Create a new issue with a title and body",
		Run: func(cmd *cobra.Command, args []string) {
			_, err := runGhIssueCreate(issueTitle, issueBody, issueDoc)
			if err != nil {
				cmd.PrintErrf("Error creating issue: %v\n", err)
			}
		},
	}
)
