package tui

import (
	"fmt"
	"math"

	"github.com/dundee/gdu/v4/analyze"
)

func (ui *UI) formatFileRow(item analyze.Item) string {
	var part int

	if ui.showApparentSize {
		part = int(float64(item.GetSize()) / float64(item.GetParent().GetSize()) * 10.0)
	} else {
		part = int(float64(item.GetUsage()) / float64(item.GetParent().GetUsage()) * 10.0)
	}

	row := string(item.GetFlag())

	if ui.useColors {
		row += "[#e67100::b]"
	} else {
		row += "[::b]"
	}

	if ui.showApparentSize {
		row += fmt.Sprintf("%15s", ui.formatSize(item.GetSize(), false, true))
	} else {
		row += fmt.Sprintf("%15s", ui.formatSize(item.GetUsage(), false, true))
	}

	row += getUsageGraph(part)

	if item.IsDir() {
		if ui.useColors {
			row += "[#3498db::b]/"
		} else {
			row += "[::b]/"
		}
	}
	row += item.GetName()
	return row
}

func (ui *UI) formatSize(size int64, reverseColor bool, transparentBg bool) string {
	var color string
	if reverseColor {
		if ui.useColors {
			color = "[black:#2479d0:-]"
		} else {
			color = "[black:white:-]"
		}
	} else {
		if transparentBg {
			color = "[-::]"
		} else {
			color = "[white:black:-]"
		}
	}

	switch {
	case size > 1e12:
		return fmt.Sprintf("%.1f%s TiB", float64(size)/math.Pow(2, 40), color)
	case size > 1e9:
		return fmt.Sprintf("%.1f%s GiB", float64(size)/math.Pow(2, 30), color)
	case size > 1e6:
		return fmt.Sprintf("%.1f%s MiB", float64(size)/math.Pow(2, 20), color)
	case size > 1e3:
		return fmt.Sprintf("%.1f%s KiB", float64(size)/math.Pow(2, 10), color)
	default:
		return fmt.Sprintf("%d%s B", size, color)
	}
}
