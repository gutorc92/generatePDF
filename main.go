package main

import (
	"fmt"
	"os"
	"time"

	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)

func main() {
	dat, err := os.ReadFile("outorgante.txt")
	if err != nil {
		panic(err)
	}
	dat2, err := os.ReadFile("outogardo.txt")
	if err != nil {
		panic(err)
	}
	dat3, err := os.ReadFile("objeto.txt")
	if err != nil {
		panic(err)
	}
	dat4, err := os.ReadFile("poderes.txt")
	if err != nil {
		panic(err)
	}
	dat5, err := os.ReadFile("poderes_especificos.txt")
	if err != nil {
		panic(err)
	}
	fmt.Print(string(dat))
	begin := time.Now()

	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	m.SetPageMargins(10, 15, 10)

	m.RegisterHeader(func() {
		m.Row(22, func() {

			m.Col(0, func() {
				m.Text("PROCURAÇÃO", props.Text{
					Size:        14,
					Style:       consts.Bold,
					Align:       consts.Center,
					Extrapolate: false,
				})
			})
		})
	})

	m.RegisterFooter(func() {
		m.Row(10, func() {
			m.ColSpace(4)
			m.Text("Brasília-DF, "+begin.Format("02 de January de 2006"), props.Text{
				Top:   3,
				Align: consts.Left,
			})
		})
		m.Row(10, func() {
			m.ColSpace(4)
			m.Col(4, func() {
				m.Signature("Teste")
			})
		})
	})

	m.Row(10, func() {
		m.Col(12, func() {
			m.Text("Pelo presente instrumento particular de mandato por mim abaixo assinado:", props.Text{
				Top:   3,
				Align: consts.Left,
			})

		})
	})
	m.Row(17, func() {
		m.Col(2, func() {
			m.Text("OUTORGANTE:", props.Text{
				Top:   3,
				Style: consts.Bold,
				Align: consts.Left,
			})

		})
		m.Col(10, func() {
			m.Text(string(dat), props.Text{
				Top: 3,
			})
		})
	})

	m.Row(17, func() {
		m.Col(2, func() {
			m.Text("OUTORGADA:", props.Text{
				Top:   3,
				Style: consts.Bold,
				Align: consts.Left,
			})

		})
		m.Col(10, func() {
			m.Text(string(dat2), props.Text{
				Top: 3,
			})
		})
	})

	m.Row(15, func() {
		m.Col(2, func() {
			m.Text("OBJETO:", props.Text{
				Top:   3,
				Style: consts.Bold,
				Align: consts.Left,
			})

		})
		m.Col(10, func() {
			m.Text(string(dat3), props.Text{
				Top: 3,
			})
		})
	})

	m.Row(35, func() {
		m.Col(2, func() {
			m.Text("PODERES:", props.Text{
				Top:   3,
				Style: consts.Bold,
				Align: consts.Left,
			})

		})
		m.Col(10, func() {
			m.Text(string(dat4), props.Text{
				Top: 3,
			})
		})
	})

	m.Row(30, func() {
		m.Col(2, func() {
			m.Text("PODERES ESPECÍFICOS:", props.Text{
				Top:   3,
				Style: consts.Bold,
				Align: consts.Left,
			})

		})
		m.Col(10, func() {
			m.Text(string(dat5), props.Text{
				Top: 3,
			})
		})
	})

	err = m.OutputFileAndClose("billing.pdf")
	if err != nil {
		fmt.Println("Could not save PDF:", err)
		os.Exit(1)
	}

	end := time.Now()
	fmt.Println(end.Sub(begin))
}

func getHeader() []string {
	return []string{"", "Product", "Quantity", "Price"}
}

func getDarkGrayColor() color.Color {
	return color.Color{
		Red:   55,
		Green: 55,
		Blue:  55,
	}
}

func getGrayColor() color.Color {
	return color.Color{
		Red:   200,
		Green: 200,
		Blue:  200,
	}
}

func getBlueColor() color.Color {
	return color.Color{
		Red:   10,
		Green: 10,
		Blue:  150,
	}
}

func getRedColor() color.Color {
	return color.Color{
		Red:   150,
		Green: 10,
		Blue:  10,
	}
}
