package pdf

import (
	"fmt"
	"image"
	"io"

	"github.com/unidoc/unipdf/v3/creator"
)

func GenerateWithBarcode(ws io.Writer, title string, url string, bcodeImg image.Image, code string) error {
	xPos := 100.0
	yPos := 100.0
	c := creator.New()
	c.SetPageMargins(50, 50, 100, 70)

	// Draw a header on each page.
	c.DrawHeader(func(block *creator.Block, args creator.HeaderFunctionArgs) {
		p := c.NewParagraph(title)
		p.SetFontSize(28)
		p.SetPos(50, 20)
		p.SetColor(creator.ColorRGBFrom8bit(63, 68, 76))
		block.Draw(p)
	})

	img, err := c.NewImageFromGoImage(bcodeImg)
	if err != nil {
		return err
	}
	img.ScaleToWidth(410)
	img.SetPos(xPos, yPos)
	_ = c.Draw(img)

	// Add the code below.
	cd := c.NewParagraph(code)
	cd.SetWidth(410)
	cd.SetPos(xPos, yPos+10+img.Height())
	cd.SetTextAlignment(creator.TextAlignmentCenter)
	_ = c.Draw(cd)

	// Add the code below.
	ul := c.NewStyledParagraph()
	ul.AddExternalLink("Scan here or open barcode online", url)
	ul.SetWidth(410)
	ul.SetPos(xPos, yPos+40+img.Height())
	ul.SetTextAlignment(creator.TextAlignmentCenter)
	_ = c.Draw(ul)

	// Draw footer on each page.
	c.DrawFooter(func(block *creator.Block, args creator.FooterFunctionArgs) {
		// Draw the on a block for each page.
		p := c.NewParagraph("unidoc.io")
		p.SetFontSize(8)
		p.SetPos(50, 20)
		p.SetColor(creator.ColorRGBFrom8bit(63, 68, 76))
		block.Draw(p)

		strPage := fmt.Sprintf("Page %d of %d", args.PageNum, args.TotalPages)
		p = c.NewParagraph(strPage)
		p.SetFontSize(8)
		p.SetPos(300, 20)
		p.SetColor(creator.ColorRGBFrom8bit(63, 68, 76))
		block.Draw(p)
	})

	err = c.Write(ws)
	if err != nil {
		return err
	}

	return nil
}
