package lumecmd

import (
	"fmt"
	"os"
	"time"

	"git.kill0.net/chill9/lifx-go"
	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
)

type Printer interface {
	Results(results []lifx.Result)
	Lights(lights []lifx.Light)
}

type defaultPrinter struct{}

type tablePrinter struct{}

func NewPrinter(format string) Printer {
	switch format {
	case "table":
		return &tablePrinter{}
	default:
		return &defaultPrinter{}
	}
}

func (dp *defaultPrinter) Results(results []lifx.Result) {
	sortResults(results)

	table := tablewriter.NewWriter(os.Stdout)
	_, rows := makeResultsTable(results)

	for _, v := range rows {
		table.Append(v)
	}

	fmt.Printf("total %d\n", len(results))
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetAutoWrapText(false)
	table.SetBorder(false)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetHeaderLine(false)
	table.SetNoWhiteSpace(true)
	table.SetRowSeparator("")
	table.SetTablePadding(" ")
	table.Render()
}

func (tp *tablePrinter) Results(results []lifx.Result) {
	sortResults(results)

	table := tablewriter.NewWriter(os.Stdout)
	hdr, rows := makeResultsTable(results)

	for _, v := range rows {
		table.Append(v)
	}

	table.SetHeader(hdr)
	table.Render()
}

func (dp *defaultPrinter) Lights(lights []lifx.Light) {
	sortLights(lights)

	table := tablewriter.NewWriter(os.Stdout)
	_, rows := makeLightsTable(lights)

	for _, v := range rows {
		table.Append(v)
	}

	fmt.Printf("total %d\n", len(lights))
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetAutoWrapText(false)
	table.SetBorder(false)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetHeaderLine(false)
	table.SetNoWhiteSpace(true)
	table.SetRowSeparator("")
	table.SetTablePadding(" ")
	table.Render()
}

func (tp *tablePrinter) Lights(lights []lifx.Light) {
	sortLights(lights)

	table := tablewriter.NewWriter(os.Stdout)
	hdr, rows := makeLightsTable(lights)

	for _, v := range rows {
		table.Append(v)
	}

	table.SetHeader(hdr)
	table.Render()
}

func ColorizeIndicator(s string) string {
	c := color.New(color.FgRed)
	if s == "on" {
		c = color.New(color.FgGreen)
	}

	return c.Sprint(GetConfig().Indicator)
}

func ColorizePower(s string) string {
	c := color.New(color.FgRed)
	if s == "on" {
		c = color.New(color.FgGreen)
	}

	return c.Sprint(s)
}

func ColorizeStatus(s lifx.Status) string {
	c := color.New(color.FgRed)
	if s == "ok" {
		c = color.New(color.FgGreen)
	}

	return c.Sprint(s)
}

func PrintWithIndent(indent int, s string) {
	fmt.Printf("%*s%s", indent, "", s)
}

func PrintfWithIndent(indent int, format string, a ...interface{}) (n int, err error) {
	format = fmt.Sprintf("%*s%s", indent, "", format)
	return fmt.Printf(format, a...)
}

func makeLightsTable(lights []lifx.Light) (hdr []string, rows [][]string) {
	hdr = []string{"", "ID", "Location", "Group", "Label", "Last Seen", "Power"}

	for _, l := range lights {
		rows = append(rows, []string{
			fmt.Sprint(ColorizeIndicator(l.Power)),
			fmt.Sprint(l.Id),
			fmt.Sprint(l.Location.Name),
			fmt.Sprint(l.Group.Name),
			fmt.Sprint(l.Label),
			fmt.Sprint(l.LastSeen.Local().Format(time.RFC3339)),
			fmt.Sprint(ColorizePower(l.Power)),
		})

	}

	return
}

func makeResultsTable(results []lifx.Result) (hdr []string, rows [][]string) {
	hdr = []string{"ID", "Label", "Status"}

	for _, r := range results {
		rows = append(rows, []string{
			fmt.Sprint(r.Id),
			fmt.Sprint(r.Label),
			fmt.Sprint(ColorizeStatus(r.Status)),
		})

	}

	return
}
