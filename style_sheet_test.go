package xlsx

import (
	"github.com/stretchr/testify/require"
	"testing"

	"github.com/plandem/ooxml"
	sharedML "github.com/plandem/ooxml/ml"
	"github.com/plandem/xlsx/format"
	"github.com/plandem/xlsx/internal/ml"
)

func TestStyleSheets(t *testing.T) {
	pkg := shared.NewPackage(nil)
	doc := &Spreadsheet{
		pkg:           pkg,
		Package:       pkg,
		relationships: shared.NewRelationships("not matter the name", pkg),
	}

	ss := newStyleSheet("xl/styles.xml", doc)

	require.NotNil(t, pkg)
	require.NotNil(t, ss)
	require.Equal(t, 0, len(ss.xfIndex))
	require.Nil(t, ss.ml.NumberFormats)
	require.Nil(t, ss.ml.Borders)
	require.Nil(t, ss.ml.Fills)
	require.Nil(t, ss.ml.Fonts)
	require.Nil(t, ss.ml.CellXfs)

	style := format.New(
		format.Font.Name("Calibri"),
		format.Font.Size(12),
		format.Font.Color("#FF0000"),
		format.Font.Scheme(format.FontSchemeMinor),
		format.Font.Family(format.FontFamilySwiss),

		format.Fill.Type(format.PatternTypeNone),

		format.Alignment.VAlign(format.VAlignBottom),
		format.Alignment.HAlign(format.HAlignFill),
		format.Border.Color("#ff00ff"),
		format.Border.Type(format.BorderStyleDashDot),
		format.Protection.Hidden,
		format.Protection.Locked,

		format.Fill.Type(format.PatternTypeDarkDown),
		format.Fill.Color("#FFFFFF"),
		format.Fill.Background("#FF0000"),
	)

	styleRef := ss.addXF(style)
	require.Equal(t, format.StyleRefID(0), styleRef)
	require.Nil(t, ss.ml.NumberFormats)

	indexedColor := 2
	require.Equal(t, &[]*ml.Font{{
		Name:   sharedML.Property("Calibri"),
		Size:   sharedML.Property("12"),
		Color:  &ml.Color{Indexed: &indexedColor},
		Scheme: sharedML.Property("minor"),
		Family: sharedML.Property("2"),
	}}, ss.ml.Fonts)

	indexedColor2 := 1
	require.Equal(t, &[]*ml.Fill{{
		Pattern: &ml.PatternFill{
			Type:       8,
			Color:      &ml.Color{Indexed: &indexedColor2},
			Background: &ml.Color{Indexed: &indexedColor},
		},
	}}, ss.ml.Fills)

	indexedColor = 6
	require.Equal(t, &[]*ml.Border{{
		Left:   &ml.BorderSegment{Type: 10, Color: &ml.Color{Indexed: &indexedColor}},
		Right:  &ml.BorderSegment{Type: 10, Color: &ml.Color{Indexed: &indexedColor}},
		Top:    &ml.BorderSegment{Type: 10, Color: &ml.Color{Indexed: &indexedColor}},
		Bottom: &ml.BorderSegment{Type: 10, Color: &ml.Color{Indexed: &indexedColor}},
	}}, ss.ml.Borders)

	//TODO: refactor test after of added 'default items' for a new XLSX, because NumFmtId,FontId,... equals 0 when it's default settings
	require.Equal(t, &[]*ml.StyleRef{{
		NumFmtId: 0,
		FontId:   0,
		FillId:   0,
		BorderId: 0,
		XfId:     0,
		Protection: &ml.CellProtection{
			Hidden: true,
			Locked: true,
		},
		Alignment: &ml.CellAlignment{
			Horizontal: 5,
			Vertical:   3,
		},
		//ApplyFill:       true,
		//ApplyBorder:     true,
		//ApplyFont:       true,
		ApplyAlignment:  true,
		ApplyProtection: true,
	}}, ss.ml.CellXfs)
}