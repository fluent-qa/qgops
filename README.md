# README

## Build and Deploy to Local

```shell
sh deploy-local.sh
```

## fluent cli help

```shell
❯ fluent -h       
fluent cli is a collection of daily scripts

Usage:
  fluent [flags]
  fluent [command]

Examples:
fluent <category> <action> <--options>

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  issue       Create a new issue
  starter     fluent start java,go,ts-lib,fe-next,mono,python
  util        fluent  util echo,mac-cleanup

Flags:
  -h, --help      help for fluent
  -v, --version   version for fluent

Use "fluent [command] --help" for more information about a command.

```

## fluent starer cli

create template project: ***fluent starter <name>***

```shell
fluent starter -h 


Usage:
  fluent starter [flags]

Examples:
fluent start java

Flags:
  -h, --help   help for starter

```

```shell
fluent starter -h
fluent starter java
fluent starter python
fluent starter go
fluent starter ts-lib
fluent starter mono
```

## fluent util cli

***fluent util <cmd>***, to run ,and add more commands in utils.json file

```shell
fluent  util echo,mac-cleanup

Usage:
  fluent util [flags]

Examples:
fluent  util echo

Flags:
  -h, --help   help for util

```

## action json file

action structure:
- name: command name
- desc: command usage
- commands: a list of shell commands or something else which is runnable 

```json
{
  "actions": [
    {
      "name": "echo",
      "desc": "try cli",
      "commands": [
        "echo \"hello fluent cli\""
      ]
    },    {
      "name": "mac-cleanup",
      "desc": "cleanup mac pc to free disk space",
      "commands": [
        "sh mac-cleanup.sh"
      ]
    }
  ]
}
```




## Pocketbase Server

```shell
./pocketbase serve
```

All Url:

- http://127.0.0.1:8090 - if pb_public directory exists, serves the static content from it (html, css, images, etc.)
- http://127.0.0.1:8090/_/ - Admin dashboard UI
- http://127.0.0.1:8090/api/ - REST API



### References 

- [github-cli](https://github.com/cli/cli.git)

### Project Layout-GITHUB-CLI
Project Layout:

```shell
.
├── LICENSE
├── Makefile
├── README.md
├── api
│   ├── client.go
│   ├── client_test.go
│   ├── export_pr.go
│   ├── export_pr_test.go
│   ├── export_repo.go
│   ├── http_client.go
│   ├── http_client_test.go
│   ├── pull_request_test.go
│   ├── queries_branch_issue_reference.go
│   ├── queries_comments.go
│   ├── queries_issue.go
│   ├── queries_org.go
│   ├── queries_pr.go
│   ├── queries_pr_review.go
│   ├── queries_pr_test.go
│   ├── queries_projects_v2.go
│   ├── queries_projects_v2_test.go
│   ├── queries_repo.go
│   ├── queries_repo_test.go
│   ├── queries_user.go
│   ├── query_builder.go
│   ├── query_builder_test.go
│   ├── reaction_groups.go
│   └── reaction_groups_test.go
├── build
│   └── windows
│       ├── ci.wixproj
│       ├── ci.wxs
│       └── ui.wxs
├── cmd
│   ├── gen-docs
│   │   ├── main.go
│   │   └── main_test.go
│   └── ci
│       ├── main.go
│       └── main_test.go
├── context
│   ├── context.go
│   ├── remote.go
│   └── remote_test.go
├── docs
│   ├── README.md
│   ├── command-line-syntax.md
│   ├── ci-vs-hub.md
│   ├── install_linux.md
│   ├── multiple-accounts.md
│   ├── project-layout.md
│   ├── releasing.md
│   ├── source.md
│   ├── triage.md
│   └── working-with-us.md
├── git
│   ├── client.go
│   ├── client_test.go
│   ├── command.go
│   ├── errors.go
│   ├── fixtures
│   │   └── simple.git
│   │       ├── HEAD
│   │       ├── config
│   │       ├── index
│   │       ├── logs
│   │       │   ├── HEAD
│   │       │   └── refs
│   │       │       └── heads
│   │       │           └── main
│   │       ├── objects
│   │       │   ├── 4b
│   │       │   │   └── 825dc642cb6eb9a060e54bf8d69288fbee4904
│   │       │   ├── 6f
│   │       │   │   └── 1a2405cace1633d89a79c74c65f22fe78f9659
│   │       │   └── d1
│   │       │       └── e0abfb7d158ed544a202a6958c62d4fc22e12f
│   │       └── refs
│   │           └── heads
│   │               └── main
│   ├── objects.go
│   ├── url.go
│   └── url_test.go
├── go.mod
├── go.sum
├── internal
│   ├── authflow
│   │   ├── flow.go
│   │   └── success.go
│   ├── browser
│   │   ├── browser.go
│   │   └── stub.go
│   ├── build
│   │   └── build.go
│   ├── codespaces
│   │   ├── api
│   │   │   ├── api.go
│   │   │   └── api_test.go
│   │   ├── codespaces.go
│   │   ├── connection
│   │   │   ├── connection.go
│   │   │   ├── connection_test.go
│   │   │   └── tunnels_api_server_mock.go
│   │   ├── portforwarder
│   │   │   ├── port_forwarder.go
│   │   │   └── port_forwarder_test.go
│   │   ├── rpc
│   │   │   ├── codespace
│   │   │   │   ├── codespace_host_service.v1.pb.go
│   │   │   │   ├── codespace_host_service.v1.proto
│   │   │   │   ├── codespace_host_service.v1.proto.mock.go
│   │   │   │   └── codespace_host_service.v1_grpc.pb.go
│   │   │   ├── generate.md
│   │   │   ├── generate.sh
│   │   │   ├── invoker.go
│   │   │   ├── invoker_test.go
│   │   │   ├── jupyter
│   │   │   │   ├── jupyter_server_host_service.v1.pb.go
│   │   │   │   ├── jupyter_server_host_service.v1.proto
│   │   │   │   ├── jupyter_server_host_service.v1.proto.mock.go
│   │   │   │   └── jupyter_server_host_service.v1_grpc.pb.go
│   │   │   ├── ssh
│   │   │   │   ├── ssh_server_host_service.v1.pb.go
│   │   │   │   ├── ssh_server_host_service.v1.proto
│   │   │   │   ├── ssh_server_host_service.v1.proto.mock.go
│   │   │   │   └── ssh_server_host_service.v1_grpc.pb.go
│   │   │   └── test
│   │   │       └── port_forwarder.go
│   │   ├── ssh.go
│   │   ├── ssh_test.go
│   │   └── states.go
│   ├── config
│   │   ├── auth_config_test.go
│   │   ├── config.go
│   │   ├── config_mock.go
│   │   ├── config_test.go
│   │   ├── migrate_test.go
│   │   ├── migration
│   │   │   ├── multi_account.go
│   │   │   └── multi_account_test.go
│   │   ├── migration_mock.go
│   │   └── stub.go
│   ├── docs
│   │   ├── docs_test.go
│   │   ├── man.go
│   │   ├── man_test.go
│   │   ├── markdown.go
│   │   └── markdown_test.go
│   ├── featuredetection
│   │   ├── detector_mock.go
│   │   ├── feature_detection.go
│   │   └── feature_detection_test.go
│   ├── ghinstance
│   │   ├── host.go
│   │   └── host_test.go
│   ├── ghrepo
│   │   ├── repo.go
│   │   └── repo_test.go
│   ├── keyring
│   │   └── keyring.go
│   ├── prompter
│   │   ├── prompter.go
│   │   ├── prompter_mock.go
│   │   └── test.go
│   ├── run
│   │   ├── run.go
│   │   └── stub.go
│   ├── tableprinter
│   │   └── table_printer.go
│   ├── text
│   │   ├── text.go
│   │   └── text_test.go
│   └── update
│       ├── update.go
│       └── update_test.go
├── pkg
│   ├── cmd
│   │   ├── actions
│   │   │   └── actions.go
│   │   ├── alias
│   │   │   ├── alias.go
│   │   │   ├── delete
│   │   │   │   ├── delete.go
│   │   │   │   └── delete_test.go
│   │   │   ├── imports
│   │   │   │   ├── import.go
│   │   │   │   └── import_test.go
│   │   │   ├── list
│   │   │   │   ├── list.go
│   │   │   │   └── list_test.go
│   │   │   ├── set
│   │   │   │   ├── set.go
│   │   │   │   └── set_test.go
│   │   │   └── shared
│   │   │       ├── validations.go
│   │   │       └── validations_test.go
│   │   ├── api
│   │   │   ├── api.go
│   │   │   ├── api_test.go
│   │   │   ├── fields.go
│   │   │   ├── fields_test.go
│   │   │   ├── http.go
│   │   │   ├── http_test.go
│   │   │   ├── pagination.go
│   │   │   └── pagination_test.go
│   │   ├── auth
│   │   │   ├── auth.go
│   │   │   ├── gitcredential
│   │   │   │   ├── helper.go
│   │   │   │   └── helper_test.go
│   │   │   ├── login
│   │   │   │   ├── login.go
│   │   │   │   └── login_test.go
│   │   │   ├── logout
│   │   │   │   ├── logout.go
│   │   │   │   └── logout_test.go
│   │   │   ├── refresh
│   │   │   │   ├── refresh.go
│   │   │   │   └── refresh_test.go
│   │   │   ├── setupgit
│   │   │   │   ├── setupgit.go
│   │   │   │   └── setupgit_test.go
│   │   │   ├── shared
│   │   │   │   ├── git_credential.go
│   │   │   │   ├── git_credential_test.go
│   │   │   │   ├── login_flow.go
│   │   │   │   ├── login_flow_test.go
│   │   │   │   ├── oauth_scopes.go
│   │   │   │   ├── oauth_scopes_test.go
│   │   │   │   ├── prompt.go
│   │   │   │   └── writeable.go
│   │   │   ├── status
│   │   │   │   ├── status.go
│   │   │   │   └── status_test.go
│   │   │   ├── switch
│   │   │   │   ├── switch.go
│   │   │   │   └── switch_test.go
│   │   │   └── token
│   │   │       ├── token.go
│   │   │       └── token_test.go
│   │   ├── browse
│   │   │   ├── browse.go
│   │   │   └── browse_test.go
│   │   ├── cache
│   │   │   ├── cache.go
│   │   │   ├── delete
│   │   │   │   ├── delete.go
│   │   │   │   └── delete_test.go
│   │   │   ├── list
│   │   │   │   ├── list.go
│   │   │   │   └── list_test.go
│   │   │   └── shared
│   │   │       ├── shared.go
│   │   │       └── shared_test.go
│   │   ├── codespace
│   │   │   ├── code.go
│   │   │   ├── code_test.go
│   │   │   ├── codespace_selector.go
│   │   │   ├── codespace_selector_test.go
│   │   │   ├── common.go
│   │   │   ├── common_test.go
│   │   │   ├── create.go
│   │   │   ├── create_test.go
│   │   │   ├── delete.go
│   │   │   ├── delete_test.go
│   │   │   ├── edit.go
│   │   │   ├── edit_test.go
│   │   │   ├── jupyter.go
│   │   │   ├── list.go
│   │   │   ├── list_test.go
│   │   │   ├── logs.go
│   │   │   ├── logs_test.go
│   │   │   ├── mock_api.go
│   │   │   ├── mock_prompter.go
│   │   │   ├── ports.go
│   │   │   ├── ports_test.go
│   │   │   ├── rebuild.go
│   │   │   ├── rebuild_test.go
│   │   │   ├── root.go
│   │   │   ├── select.go
│   │   │   ├── select_test.go
│   │   │   ├── ssh.go
│   │   │   ├── ssh_test.go
│   │   │   ├── stop.go
│   │   │   ├── stop_test.go
│   │   │   ├── view.go
│   │   │   └── view_test.go
│   │   ├── completion
│   │   │   ├── completion.go
│   │   │   └── completion_test.go
│   │   ├── config
│   │   │   ├── clear-cache
│   │   │   │   ├── clear_cache.go
│   │   │   │   └── clear_cache_test.go
│   │   │   ├── config.go
│   │   │   ├── get
│   │   │   │   ├── get.go
│   │   │   │   └── get_test.go
│   │   │   ├── list
│   │   │   │   ├── list.go
│   │   │   │   └── list_test.go
│   │   │   └── set
│   │   │       ├── set.go
│   │   │       └── set_test.go
│   │   ├── extension
│   │   │   ├── browse
│   │   │   │   ├── browse.go
│   │   │   │   ├── browse_test.go
│   │   │   │   └── rg.go
│   │   │   ├── command.go
│   │   │   ├── command_test.go
│   │   │   ├── ext_tmpls
│   │   │   │   ├── buildScript.sh
│   │   │   │   ├── goBinMain.go.txt
│   │   │   │   ├── goBinWorkflow.yml
│   │   │   │   ├── otherBinWorkflow.yml
│   │   │   │   └── script.sh
│   │   │   ├── extension.go
│   │   │   ├── extension_test.go
│   │   │   ├── git.go
│   │   │   ├── http.go
│   │   │   ├── manager.go
│   │   │   ├── manager_test.go
│   │   │   ├── mocks.go
│   │   │   ├── symlink_other.go
│   │   │   └── symlink_windows.go
│   │   ├── factory
│   │   │   ├── default.go
│   │   │   ├── default_test.go
│   │   │   ├── remote_resolver.go
│   │   │   └── remote_resolver_test.go
│   │   ├── gist
│   │   │   ├── clone
│   │   │   │   ├── clone.go
│   │   │   │   └── clone_test.go
│   │   │   ├── create
│   │   │   │   ├── create.go
│   │   │   │   └── create_test.go
│   │   │   ├── delete
│   │   │   │   ├── delete.go
│   │   │   │   └── delete_test.go
│   │   │   ├── edit
│   │   │   │   ├── edit.go
│   │   │   │   └── edit_test.go
│   │   │   ├── gist.go
│   │   │   ├── list
│   │   │   │   ├── list.go
│   │   │   │   └── list_test.go
│   │   │   ├── rename
│   │   │   │   ├── rename.go
│   │   │   │   └── rename_test.go
│   │   │   ├── shared
│   │   │   │   ├── shared.go
│   │   │   │   └── shared_test.go
│   │   │   └── view
│   │   │       ├── view.go
│   │   │       └── view_test.go
│   │   ├── gpg-key
│   │   │   ├── add
│   │   │   │   ├── add.go
│   │   │   │   ├── add_test.go
│   │   │   │   └── http.go
│   │   │   ├── delete
│   │   │   │   ├── delete.go
│   │   │   │   ├── delete_test.go
│   │   │   │   └── http.go
│   │   │   ├── gpg_key.go
│   │   │   └── list
│   │   │       ├── http.go
│   │   │       ├── list.go
│   │   │       └── list_test.go
│   │   ├── issue
│   │   │   ├── close
│   │   │   │   ├── close.go
│   │   │   │   └── close_test.go
│   │   │   ├── comment
│   │   │   │   ├── comment.go
│   │   │   │   └── comment_test.go
│   │   │   ├── create
│   │   │   │   ├── create.go
│   │   │   │   ├── create_test.go
│   │   │   │   └── fixtures
│   │   │   │       └── repoWithNonLegacyIssueTemplates
│   │   │   ├── delete
│   │   │   │   ├── delete.go
│   │   │   │   └── delete_test.go
│   │   │   ├── develop
│   │   │   │   ├── develop.go
│   │   │   │   └── develop_test.go
│   │   │   ├── edit
│   │   │   │   ├── edit.go
│   │   │   │   └── edit_test.go
│   │   │   ├── issue.go
│   │   │   ├── list
│   │   │   │   ├── fixtures
│   │   │   │   │   ├── issueList.json
│   │   │   │   │   └── issueSearch.json
│   │   │   │   ├── http.go
│   │   │   │   ├── http_test.go
│   │   │   │   ├── list.go
│   │   │   │   └── list_test.go
│   │   │   ├── lock
│   │   │   │   ├── lock.go
│   │   │   │   └── lock_test.go
│   │   │   ├── pin
│   │   │   │   ├── pin.go
│   │   │   │   └── pin_test.go
│   │   │   ├── reopen
│   │   │   │   ├── reopen.go
│   │   │   │   └── reopen_test.go
│   │   │   ├── shared
│   │   │   │   ├── display.go
│   │   │   │   ├── lookup.go
│   │   │   │   └── lookup_test.go
│   │   │   ├── status
│   │   │   │   ├── fixtures
│   │   │   │   │   └── issueStatus.json
│   │   │   │   ├── status.go
│   │   │   │   └── status_test.go
│   │   │   ├── transfer
│   │   │   │   ├── transfer.go
│   │   │   │   └── transfer_test.go
│   │   │   ├── unpin
│   │   │   │   ├── unpin.go
│   │   │   │   └── unpin_test.go
│   │   │   └── view
│   │   │       ├── fixtures
│   │   │       │   ├── issueView_preview.json
│   │   │       │   ├── issueView_previewClosedState.json
│   │   │       │   ├── issueView_previewFullComments.json
│   │   │       │   ├── issueView_previewSingleComment.json
│   │   │       │   ├── issueView_previewWithEmptyBody.json
│   │   │       │   └── issueView_previewWithMetadata.json
│   │   │       ├── http.go
│   │   │       ├── view.go
│   │   │       └── view_test.go
│   │   ├── label
│   │   │   ├── clone.go
│   │   │   ├── clone_test.go
│   │   │   ├── create.go
│   │   │   ├── create_test.go
│   │   │   ├── delete.go
│   │   │   ├── delete_test.go
│   │   │   ├── edit.go
│   │   │   ├── edit_test.go
│   │   │   ├── http.go
│   │   │   ├── http_test.go
│   │   │   ├── label.go
│   │   │   ├── list.go
│   │   │   ├── list_test.go
│   │   │   └── shared.go
│   │   ├── org
│   │   │   ├── list
│   │   │   │   ├── http.go
│   │   │   │   ├── http_test.go
│   │   │   │   ├── list.go
│   │   │   │   └── list_test.go
│   │   │   └── org.go
│   │   ├── pr
│   │   │   ├── checkout
│   │   │   │   ├── checkout.go
│   │   │   │   └── checkout_test.go
│   │   │   ├── checks
│   │   │   │   ├── aggregate.go
│   │   │   │   ├── checks.go
│   │   │   │   ├── checks_test.go
│   │   │   │   ├── fixtures
│   │   │   │   │   ├── allPassing.json
│   │   │   │   │   ├── onlyRequired.json
│   │   │   │   │   ├── someCancelled.json
│   │   │   │   │   ├── someFailing.json
│   │   │   │   │   ├── somePending.json
│   │   │   │   │   ├── someSkipping.json
│   │   │   │   │   ├── withDescriptions.json
│   │   │   │   │   ├── withEvents.json
│   │   │   │   │   ├── withStatuses.json
│   │   │   │   │   └── withoutEvents.json
│   │   │   │   └── output.go
│   │   │   ├── close
│   │   │   │   ├── close.go
│   │   │   │   └── close_test.go
│   │   │   ├── comment
│   │   │   │   ├── comment.go
│   │   │   │   └── comment_test.go
│   │   │   ├── create
│   │   │   │   ├── create.go
│   │   │   │   ├── create_test.go
│   │   │   │   ├── fixtures
│   │   │   │   │   └── repoWithNonLegacyPRTemplates
│   │   │   │   ├── regexp_writer.go
│   │   │   │   └── regexp_writer_test.go
│   │   │   ├── diff
│   │   │   │   ├── diff.go
│   │   │   │   └── diff_test.go
│   │   │   ├── edit
│   │   │   │   ├── edit.go
│   │   │   │   └── edit_test.go
│   │   │   ├── list
│   │   │   │   ├── fixtures
│   │   │   │   │   ├── prList.json
│   │   │   │   │   └── prListWithDuplicates.json
│   │   │   │   ├── http.go
│   │   │   │   ├── http_test.go
│   │   │   │   ├── list.go
│   │   │   │   └── list_test.go
│   │   │   ├── merge
│   │   │   │   ├── http.go
│   │   │   │   ├── merge.go
│   │   │   │   └── merge_test.go
│   │   │   ├── pr.go
│   │   │   ├── ready
│   │   │   │   ├── ready.go
│   │   │   │   └── ready_test.go
│   │   │   ├── reopen
│   │   │   │   ├── reopen.go
│   │   │   │   └── reopen_test.go
│   │   │   ├── review
│   │   │   │   ├── review.go
│   │   │   │   └── review_test.go
│   │   │   ├── shared
│   │   │   │   ├── commentable.go
│   │   │   │   ├── comments.go
│   │   │   │   ├── completion.go
│   │   │   │   ├── display.go
│   │   │   │   ├── display_test.go
│   │   │   │   ├── editable.go
│   │   │   │   ├── editable_http.go
│   │   │   │   ├── finder.go
│   │   │   │   ├── finder_test.go
│   │   │   │   ├── params.go
│   │   │   │   ├── params_test.go
│   │   │   │   ├── preserve.go
│   │   │   │   ├── preserve_test.go
│   │   │   │   ├── reaction_groups.go
│   │   │   │   ├── state.go
│   │   │   │   ├── survey.go
│   │   │   │   ├── survey_test.go
│   │   │   │   ├── templates.go
│   │   │   │   └── templates_test.go
│   │   │   ├── status
│   │   │   │   ├── fixtures
│   │   │   │   │   ├── prStatus.json
│   │   │   │   │   ├── prStatusChecks.json
│   │   │   │   │   ├── prStatusChecksWithStatesByCount.json
│   │   │   │   │   ├── prStatusCurrentBranch.json
│   │   │   │   │   ├── prStatusCurrentBranchClosed.json
│   │   │   │   │   ├── prStatusCurrentBranchClosedOnDefaultBranch.json
│   │   │   │   │   ├── prStatusCurrentBranchMerged.json
│   │   │   │   │   └── prStatusCurrentBranchMergedOnDefaultBranch.json
│   │   │   │   ├── http.go
│   │   │   │   ├── status.go
│   │   │   │   └── status_test.go
│   │   │   └── view
│   │   │       ├── fixtures
│   │   │       │   ├── prViewPreview.json
│   │   │       │   ├── prViewPreviewClosedState.json
│   │   │       │   ├── prViewPreviewDraftState.json
│   │   │       │   ├── prViewPreviewFullComments.json
│   │   │       │   ├── prViewPreviewManyReviews.json
│   │   │       │   ├── prViewPreviewMergedState.json
│   │   │       │   ├── prViewPreviewReviews.json
│   │   │       │   ├── prViewPreviewSingleComment.json
│   │   │       │   ├── prViewPreviewWithAllChecksFailing.json
│   │   │       │   ├── prViewPreviewWithAllChecksPassing.json
│   │   │       │   ├── prViewPreviewWithAutoMergeEnabled.json
│   │   │       │   ├── prViewPreviewWithMetadataByNumber.json
│   │   │       │   ├── prViewPreviewWithNilProject.json
│   │   │       │   ├── prViewPreviewWithNoChecks.json
│   │   │       │   ├── prViewPreviewWithReviewersByNumber.json
│   │   │       │   ├── prViewPreviewWithSomeChecksFailing.json
│   │   │       │   └── prViewPreviewWithSomeChecksPending.json
│   │   │       ├── view.go
│   │   │       └── view_test.go
│   │   ├── project
│   │   │   ├── close
│   │   │   │   ├── close.go
│   │   │   │   └── close_test.go
│   │   │   ├── copy
│   │   │   │   ├── copy.go
│   │   │   │   └── copy_test.go
│   │   │   ├── create
│   │   │   │   ├── create.go
│   │   │   │   └── create_test.go
│   │   │   ├── delete
│   │   │   │   ├── delete.go
│   │   │   │   └── delete_test.go
│   │   │   ├── edit
│   │   │   │   ├── edit.go
│   │   │   │   └── edit_test.go
│   │   │   ├── field-create
│   │   │   │   ├── field_create.go
│   │   │   │   └── field_create_test.go
│   │   │   ├── field-delete
│   │   │   │   ├── field_delete.go
│   │   │   │   └── field_delete_test.go
│   │   │   ├── field-list
│   │   │   │   ├── field_list.go
│   │   │   │   └── field_list_test.go
│   │   │   ├── item-add
│   │   │   │   ├── item_add.go
│   │   │   │   └── item_add_test.go
│   │   │   ├── item-archive
│   │   │   │   ├── item_archive.go
│   │   │   │   └── item_archive_test.go
│   │   │   ├── item-create
│   │   │   │   ├── item_create.go
│   │   │   │   └── item_create_test.go
│   │   │   ├── item-delete
│   │   │   │   ├── item_delete.go
│   │   │   │   └── item_delete_test.go
│   │   │   ├── item-edit
│   │   │   │   ├── item_edit.go
│   │   │   │   └── item_edit_test.go
│   │   │   ├── item-list
│   │   │   │   ├── item_list.go
│   │   │   │   └── item_list_test.go
│   │   │   ├── list
│   │   │   │   ├── list.go
│   │   │   │   └── list_test.go
│   │   │   ├── mark-template
│   │   │   │   ├── mark_template.go
│   │   │   │   └── mark_template_test.go
│   │   │   ├── project.go
│   │   │   ├── shared
│   │   │   │   ├── client
│   │   │   │   │   └── client.go
│   │   │   │   ├── format
│   │   │   │   │   ├── display.go
│   │   │   │   │   └── display_test.go
│   │   │   │   └── queries
│   │   │   │       ├── export_data_test.go
│   │   │   │       ├── queries.go
│   │   │   │       └── queries_test.go
│   │   │   └── view
│   │   │       ├── view.go
│   │   │       └── view_test.go
│   │   ├── release
│   │   │   ├── create
│   │   │   │   ├── create.go
│   │   │   │   ├── create_test.go
│   │   │   │   └── http.go
│   │   │   ├── delete
│   │   │   │   ├── delete.go
│   │   │   │   └── delete_test.go
│   │   │   ├── delete-asset
│   │   │   │   ├── delete_asset.go
│   │   │   │   └── delete_asset_test.go
│   │   │   ├── download
│   │   │   │   ├── download.go
│   │   │   │   └── download_test.go
│   │   │   ├── edit
│   │   │   │   ├── edit.go
│   │   │   │   ├── edit_test.go
│   │   │   │   └── http.go
│   │   │   ├── list
│   │   │   │   ├── http.go
│   │   │   │   ├── list.go
│   │   │   │   └── list_test.go
│   │   │   ├── release.go
│   │   │   ├── shared
│   │   │   │   ├── fetch.go
│   │   │   │   ├── upload.go
│   │   │   │   └── upload_test.go
│   │   │   ├── upload
│   │   │   │   ├── upload.go
│   │   │   │   └── upload_test.go
│   │   │   └── view
│   │   │       ├── view.go
│   │   │       └── view_test.go
│   │   ├── repo
│   │   │   ├── archive
│   │   │   │   ├── archive.go
│   │   │   │   ├── archive_test.go
│   │   │   │   └── http.go
│   │   │   ├── clone
│   │   │   │   ├── clone.go
│   │   │   │   └── clone_test.go
│   │   │   ├── create
│   │   │   │   ├── create.go
│   │   │   │   ├── create_test.go
│   │   │   │   ├── fixtures
│   │   │   │   │   └── repoTempList.json
│   │   │   │   ├── http.go
│   │   │   │   └── http_test.go
│   │   │   ├── credits
│   │   │   │   └── credits.go
│   │   │   ├── delete
│   │   │   │   ├── delete.go
│   │   │   │   ├── delete_test.go
│   │   │   │   └── http.go
│   │   │   ├── deploy-key
│   │   │   │   ├── add
│   │   │   │   │   ├── add.go
│   │   │   │   │   ├── add_test.go
│   │   │   │   │   └── http.go
│   │   │   │   ├── delete
│   │   │   │   │   ├── delete.go
│   │   │   │   │   ├── delete_test.go
│   │   │   │   │   └── http.go
│   │   │   │   ├── deploy-key.go
│   │   │   │   └── list
│   │   │   │       ├── http.go
│   │   │   │       ├── list.go
│   │   │   │       └── list_test.go
│   │   │   ├── edit
│   │   │   │   ├── edit.go
│   │   │   │   └── edit_test.go
│   │   │   ├── fork
│   │   │   │   ├── fork.go
│   │   │   │   ├── forkResult.json
│   │   │   │   └── fork_test.go
│   │   │   ├── garden
│   │   │   │   ├── garden.go
│   │   │   │   └── http.go
│   │   │   ├── list
│   │   │   │   ├── fixtures
│   │   │   │   │   ├── repoList.json
│   │   │   │   │   └── repoSearch.json
│   │   │   │   ├── http.go
│   │   │   │   ├── http_test.go
│   │   │   │   ├── list.go
│   │   │   │   └── list_test.go
│   │   │   ├── rename
│   │   │   │   ├── rename.go
│   │   │   │   └── rename_test.go
│   │   │   ├── repo.go
│   │   │   ├── setdefault
│   │   │   │   ├── setdefault.go
│   │   │   │   └── setdefault_test.go
│   │   │   ├── shared
│   │   │   │   ├── repo.go
│   │   │   │   └── repo_test.go
│   │   │   ├── sync
│   │   │   │   ├── git.go
│   │   │   │   ├── http.go
│   │   │   │   ├── mocks.go
│   │   │   │   ├── sync.go
│   │   │   │   └── sync_test.go
│   │   │   ├── unarchive
│   │   │   │   ├── http.go
│   │   │   │   ├── unarchive.go
│   │   │   │   └── unarchive_test.go
│   │   │   └── view
│   │   │       ├── http.go
│   │   │       ├── view.go
│   │   │       └── view_test.go
│   │   ├── root
│   │   │   ├── alias.go
│   │   │   ├── alias_test.go
│   │   │   ├── extension.go
│   │   │   ├── help.go
│   │   │   ├── help_reference.go
│   │   │   ├── help_test.go
│   │   │   ├── help_topic.go
│   │   │   ├── help_topic_test.go
│   │   │   └── root.go
│   │   ├── ruleset
│   │   │   ├── check
│   │   │   │   ├── check.go
│   │   │   │   ├── check_test.go
│   │   │   │   └── fixtures
│   │   │   │       └── rulesetCheck.json
│   │   │   ├── list
│   │   │   │   ├── fixtures
│   │   │   │   │   └── rulesetList.json
│   │   │   │   ├── list.go
│   │   │   │   └── list_test.go
│   │   │   ├── ruleset.go
│   │   │   ├── shared
│   │   │   │   ├── http.go
│   │   │   │   └── shared.go
│   │   │   └── view
│   │   │       ├── fixtures
│   │   │       │   ├── rulesetViewMultiple.json
│   │   │       │   ├── rulesetViewOrg.json
│   │   │       │   └── rulesetViewRepo.json
│   │   │       ├── http.go
│   │   │       ├── view.go
│   │   │       └── view_test.go
│   │   ├── run
│   │   │   ├── cancel
│   │   │   │   ├── cancel.go
│   │   │   │   └── cancel_test.go
│   │   │   ├── delete
│   │   │   │   ├── delete.go
│   │   │   │   └── delete_test.go
│   │   │   ├── download
│   │   │   │   ├── download.go
│   │   │   │   ├── download_test.go
│   │   │   │   ├── fixtures
│   │   │   │   │   └── myproject.zip
│   │   │   │   ├── http.go
│   │   │   │   ├── http_test.go
│   │   │   │   ├── zip.go
│   │   │   │   └── zip_test.go
│   │   │   ├── list
│   │   │   │   ├── list.go
│   │   │   │   └── list_test.go
│   │   │   ├── rerun
│   │   │   │   ├── rerun.go
│   │   │   │   └── rerun_test.go
│   │   │   ├── run.go
│   │   │   ├── shared
│   │   │   │   ├── artifacts.go
│   │   │   │   ├── artifacts_test.go
│   │   │   │   ├── presentation.go
│   │   │   │   ├── shared.go
│   │   │   │   ├── shared_test.go
│   │   │   │   └── test.go
│   │   │   ├── view
│   │   │   │   ├── fixtures
│   │   │   │   │   └── run_log.zip
│   │   │   │   ├── view.go
│   │   │   │   └── view_test.go
│   │   │   └── watch
│   │   │       ├── watch.go
│   │   │       └── watch_test.go
│   │   ├── search
│   │   │   ├── code
│   │   │   │   ├── code.go
│   │   │   │   └── code_test.go
│   │   │   ├── commits
│   │   │   │   ├── commits.go
│   │   │   │   └── commits_test.go
│   │   │   ├── issues
│   │   │   │   ├── issues.go
│   │   │   │   └── issues_test.go
│   │   │   ├── prs
│   │   │   │   ├── prs.go
│   │   │   │   └── prs_test.go
│   │   │   ├── repos
│   │   │   │   ├── repos.go
│   │   │   │   └── repos_test.go
│   │   │   ├── search.go
│   │   │   └── shared
│   │   │       ├── shared.go
│   │   │       └── shared_test.go
│   │   ├── secret
│   │   │   ├── delete
│   │   │   │   ├── delete.go
│   │   │   │   └── delete_test.go
│   │   │   ├── list
│   │   │   │   ├── list.go
│   │   │   │   └── list_test.go
│   │   │   ├── secret.go
│   │   │   ├── set
│   │   │   │   ├── http.go
│   │   │   │   ├── set.go
│   │   │   │   └── set_test.go
│   │   │   └── shared
│   │   │       ├── shared.go
│   │   │       └── shared_test.go
│   │   ├── ssh-key
│   │   │   ├── add
│   │   │   │   ├── add.go
│   │   │   │   ├── add_test.go
│   │   │   │   └── http.go
│   │   │   ├── delete
│   │   │   │   ├── delete.go
│   │   │   │   ├── delete_test.go
│   │   │   │   └── http.go
│   │   │   ├── list
│   │   │   │   ├── list.go
│   │   │   │   └── list_test.go
│   │   │   ├── shared
│   │   │   │   └── user_keys.go
│   │   │   └── ssh_key.go
│   │   ├── status
│   │   │   ├── fixtures
│   │   │   │   ├── events.json
│   │   │   │   ├── notifications.json
│   │   │   │   ├── search.json
│   │   │   │   └── search_forbidden.json
│   │   │   ├── status.go
│   │   │   └── status_test.go
│   │   ├── variable
│   │   │   ├── delete
│   │   │   │   ├── delete.go
│   │   │   │   └── delete_test.go
│   │   │   ├── list
│   │   │   │   ├── list.go
│   │   │   │   └── list_test.go
│   │   │   ├── set
│   │   │   │   ├── http.go
│   │   │   │   ├── set.go
│   │   │   │   └── set_test.go
│   │   │   ├── shared
│   │   │   │   ├── shared.go
│   │   │   │   └── shared_test.go
│   │   │   └── variable.go
│   │   ├── version
│   │   │   ├── version.go
│   │   │   └── version_test.go
│   │   └── workflow
│   │       ├── disable
│   │       │   ├── disable.go
│   │       │   └── disable_test.go
│   │       ├── enable
│   │       │   ├── enable.go
│   │       │   └── enable_test.go
│   │       ├── list
│   │       │   ├── list.go
│   │       │   └── list_test.go
│   │       ├── run
│   │       │   ├── run.go
│   │       │   └── run_test.go
│   │       ├── shared
│   │       │   ├── shared.go
│   │       │   └── test.go
│   │       ├── view
│   │       │   ├── view.go
│   │       │   └── view_test.go
│   │       └── workflow.go
│   ├── cmdutil
│   │   ├── args.go
│   │   ├── args_test.go
│   │   ├── auth_check.go
│   │   ├── auth_check_test.go
│   │   ├── cmdgroup.go
│   │   ├── errors.go
│   │   ├── factory.go
│   │   ├── factory_test.go
│   │   ├── file_input.go
│   │   ├── flags.go
│   │   ├── json_flags.go
│   │   ├── json_flags_test.go
│   │   ├── legacy.go
│   │   └── repo_override.go
│   ├── extensions
│   │   ├── extension.go
│   │   ├── extension_mock.go
│   │   └── manager_mock.go
│   ├── findsh
│   │   ├── find.go
│   │   └── find_windows.go
│   ├── githubtemplate
│   │   ├── github_template.go
│   │   └── github_template_test.go
│   ├── httpmock
│   │   ├── legacy.go
│   │   ├── registry.go
│   │   └── stub.go
│   ├── iostreams
│   │   ├── color.go
│   │   ├── color_test.go
│   │   ├── console.go
│   │   ├── console_windows.go
│   │   ├── epipe_other.go
│   │   ├── epipe_windows.go
│   │   ├── iostreams.go
│   │   └── iostreams_test.go
│   ├── jsoncolor
│   │   ├── jsoncolor.go
│   │   └── jsoncolor_test.go
│   ├── markdown
│   │   └── markdown.go
│   ├── search
│   │   ├── query.go
│   │   ├── query_test.go
│   │   ├── result.go
│   │   ├── result_test.go
│   │   ├── searcher.go
│   │   ├── searcher_mock.go
│   │   └── searcher_test.go
│   ├── set
│   │   ├── string_set.go
│   │   └── string_set_test.go
│   ├── ssh
│   │   └── ssh_keys.go
│   └── surveyext
│       ├── editor.go
│       ├── editor_manual.go
│       └── editor_test.go
├── projectlayout.txt
├── script
│   ├── build.bat
│   ├── build.go
│   ├── createrepo.sh
│   ├── distributions
│   ├── label-assets
│   ├── nolint-insert
│   ├── override.ubuntu
│   ├── release
│   ├── rpmmacros
│   ├── sign
│   └── sign.ps1
├── test
│   └── helpers.go
└── utils
    └── utils.go

262 directories, 729 files

```