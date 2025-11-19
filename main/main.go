package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Vid8866/DN5/redovalnica"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "Redovalnica",
		Usage: "Aplikacija za upravljanje z ocenami študentov",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:  "stOcen",
				Value: int(redovalnica.StOcen),
				Usage: "najmanjše število ocen potrebnih za pozitivno oceno",
			},
			&cli.Float64Flag{
				Name:  "minOcena",
				Value: redovalnica.MinOcena,
				Usage: "najmanjša možna ocena",
			},
			&cli.Float64Flag{
				Name:  "maxOcena",
				Value: redovalnica.MaxOcena,
				Usage: "največja možna ocena",
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			stOcen := cmd.Int("stOcen")
			minOc := cmd.Float64("minOcena")
			maxOc := cmd.Float64("maxOcena")

			if stOcen <= 0 {
				return fmt.Errorf("stOcen must be > 0")
			}
			if minOc > maxOc {
				return fmt.Errorf("minOcena (%.2f) nemore biti vecja od maxOcena (%.2f)", minOc, maxOc)
			}

			// Apply configuration to package
			redovalnica.StOcen = stOcen
			redovalnica.MinOcena = minOc
			redovalnica.MaxOcena = maxOc

			fmt.Printf("Uporabljena konfiguracija: minStOcen=%d, minOcena=%.2f, maxOcena=%.2f\n", stOcen, minOc, maxOc)

			// Demo: create some students and add grades
			studenti := make(map[string]redovalnica.Student)
			studenti["1001"] = redovalnica.NovStudent("Ana", "Novak")
			studenti["1002"] = redovalnica.NovStudent("Boris", "Kranjc")
			studenti["1003"] = redovalnica.NovStudent("Cilka", "Zajc")

			redovalnica.DodajOceno(studenti, "1001", 9)
			redovalnica.DodajOceno(studenti, "1001", 8)
			redovalnica.DodajOceno(studenti, "1001", 10)
			redovalnica.DodajOceno(studenti, "1002", 5)
			redovalnica.DodajOceno(studenti, "1002", 6)
			redovalnica.DodajOceno(studenti, "1003", 2)
			redovalnica.DodajOceno(studenti, "1003", 3)

			redovalnica.IzpisVsehOcen(studenti)
			redovalnica.IzpisiKoncniUspeh(studenti)

			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
