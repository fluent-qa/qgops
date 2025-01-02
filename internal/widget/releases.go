package widget

import (
	"context"
	"github.com/fluent-qa/qgops/internal/assets"
	feed2 "github.com/fluent-qa/qgops/internal/feed"
	"html/template"
	"time"
)

type Releases struct {
	widgetBase    `yaml:",inline"`
	Releases      feed2.AppReleases `yaml:"-"`
	Repositories  []string          `yaml:"repositories"`
	Token         OptionalEnvString `yaml:"token"`
	Limit         int               `yaml:"limit"`
	CollapseAfter int               `yaml:"collapse-after"`
}

func (widget *Releases) Initialize() error {
	widget.withTitle("Releases").withCacheDuration(2 * time.Hour)

	if widget.Limit <= 0 {
		widget.Limit = 10
	}

	if widget.CollapseAfter == 0 || widget.CollapseAfter < -1 {
		widget.CollapseAfter = 5
	}

	return nil
}

func (widget *Releases) Update(ctx context.Context) {
	releases, err := feed2.FetchLatestReleasesFromGithub(widget.Repositories, string(widget.Token))

	if !widget.canContinueUpdateAfterHandlingErr(err) {
		return
	}

	if len(releases) > widget.Limit {
		releases = releases[:widget.Limit]
	}

	widget.Releases = releases
}

func (widget *Releases) Render() template.HTML {
	return widget.render(widget, assets.ReleasesTemplate)
}
