package widget

import (
	"context"
	"github.com/fluent-qa/qgops/internal/assets"
	feed2 "github.com/fluent-qa/qgops/internal/feed"
	"html/template"
	"time"
)

type Calendar struct {
	widgetBase `yaml:",inline"`
	Calendar   *feed2.Calendar
}

func (widget *Calendar) Initialize() error {
	widget.withTitle("Calendar").withCacheOnTheHour()

	return nil
}

func (widget *Calendar) Update(ctx context.Context) {
	widget.Calendar = feed2.NewCalendar(time.Now())
	widget.withError(nil).scheduleNextUpdate()
}

func (widget *Calendar) Render() template.HTML {
	return widget.render(widget, assets.CalendarTemplate)
}
