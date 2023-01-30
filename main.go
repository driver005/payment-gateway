package main

import (
	"context"

	"github.com/driver005/gateway/config"
	"github.com/driver005/gateway/registry"
	// "github.com/gofiber/fiber/v2"
	// "github.com/gofiber/fiber/v2/middleware/cors"
	// "github.com/gofiber/fiber/v2/middleware/logger"
)

var ctx = context.Background()

func main() {
	o := make([]registry.OptionsModifier, 0)
	o = append(o, registry.DisableMigration())

	r := registry.New(ctx, o)

	config.Init()

	// m := r.Driver().Boleto(&intent.PaymentIntent{
	// 	PaymentMethod: method.PaymentMethod{
	// 		BillingDetails: &address.BillingDetails{
	// 			Name:  "1",
	// 			Email: "1",
	// 			Address: &address.Address{
	// 				City: "Tü",
	// 			},
	// 		},
	// 	},
	// 	Amount:      1000,
	// 	Description: "131313",
	// })

	// fmt.Println(m.NextAction.BoletoDisplayDetails)

	r.Setup()

	// var all = []spin.Name{spin.Dots5, spin.Dots11, spin.Line2, spin.Arc, spin.Toggle2, spin.Star, spin.Circle, spin.Earth, spin.Toggle4, spin.Arrow2, spin.BouncingBar, spin.Monkey, spin.Dots8, spin.BoxBounce2, spin.SquareCorners, spin.Toggle11, spin.Moon, spin.Triangle, spin.Toggle10, spin.Grenade, spin.Dots2, spin.Balloon, spin.CircleQuarters, spin.BouncingBall, spin.Point, spin.Smiley, spin.Noise, spin.Arrow, spin.Arrow3, spin.Line, spin.Pipe, spin.SimpleDots, spin.SimpleDotsScrolling, spin.Star2, spin.Dqpb, spin.Layer, spin.Dots12, spin.Hamburger, spin.Bounce, spin.Weather, spin.Dots9, spin.GrowHorizontal, spin.Toggle3, spin.Hearts, spin.Shark, spin.Dots6, spin.BoxBounce, spin.CircleHalves, spin.Toggle, spin.Toggle5, spin.Dots7, spin.GrowVertical, spin.Toggle13, spin.Christmas, spin.Balloon2, spin.Toggle7, spin.Toggle9, spin.Clock, spin.Runner, spin.Dots10, spin.Flip, spin.Toggle6, spin.Dots, spin.Squish, spin.Toggle12, spin.Dots3, spin.Dots4, spin.Toggle8, spin.Pong}

	// for _, v := range all {
	// 	w := wow.New(os.Stdout, spin.Get(v), " "+v.String())
	// 	w.Start()
	// 	time.Sleep(3 * time.Second)
	// 	w.PersistWith(spin.Spinner{Frames: []string{"✅"}}, " Finished!")
	// }
}
