package widget

import (
	"context"
	"github.com/fluent-qa/qgops/internal/assets"
	feed2 "github.com/fluent-qa/qgops/internal/feed"
	"html/template"
	"time"
)

type Markets struct {
	widgetBase     `yaml:",inline"`
	StocksRequests []feed2.MarketRequest `yaml:"stocks"`
	MarketRequests []feed2.MarketRequest `yaml:"markets"`
	Sort           string                `yaml:"sort-by"`
	Style          string                `yaml:"style"`
	Markets        feed2.Markets         `yaml:"-"`
}

func (widget *Markets) Initialize() error {
	widget.withTitle("Markets").withCacheDuration(time.Hour)

	if len(widget.MarketRequests) == 0 {
		widget.MarketRequests = widget.StocksRequests
	}

	return nil
}

func (widget *Markets) Update(ctx context.Context) {
	markets, err := feed2.FetchMarketsDataFromYahoo(widget.MarketRequests)

	if !widget.canContinueUpdateAfterHandlingErr(err) {
		return
	}

	if widget.Sort == "absolute-change" {
		markets.SortByAbsChange()
	}

	widget.Markets = markets
}

func (widget *Markets) Render() template.HTML {
	return widget.render(widget, assets.MarketsTemplate)
}
