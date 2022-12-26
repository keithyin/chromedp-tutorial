package mousetools

import (
	"context"
	"math/rand"
	"time"

	"github.com/chromedp/cdproto/input"
	"github.com/chromedp/chromedp"
)

func MoveMouse(beginX, beginY, endX, endY int, button input.MouseButton) chromedp.ActionFunc {

	return chromedp.ActionFunc(func(ctx context.Context) error {

		p := &input.DispatchMouseEventParams{
			Type:   input.MousePressed,
			X:      float64(beginX),
			Y:      float64(beginY),
			Button: button,
		}

		err := p.Do(ctx)
		if err != nil {
			return err
		}

		p.Type = input.MouseMoved

		for beginX < endX && beginY < endY {
			beginX += rand.Intn(10)
			beginY += rand.Intn(10)
			p.X = float64(beginX)
			p.Y = float64(beginY)
			err := p.Do(ctx)
			if err != nil {
				return err
			}
		}

		for beginX < endX {
			beginX += rand.Intn(10)
			p.X = float64(beginX)
			err := p.Do(ctx)
			if err != nil {
				return err
			}
		}

		for beginY < endY {
			beginY += rand.Intn(10)
			p.Y = float64(beginY)
			err := p.Do(ctx)
			if err != nil {
				return err
			}
		}

		p.Type = input.MouseReleased
		return p.Do(ctx)
	})
}

func MoveMouseMiddle(totalX int, totalY int, nTimes int, interval time.Duration) chromedp.ActionFunc {

	return chromedp.ActionFunc(func(ctx context.Context) error {
		deltaY := totalY / nTimes

		for i := 0; i < nTimes; i++ {
			time.Sleep(interval)
			p := &input.DispatchMouseEventParams{
				Type:   input.MouseWheel,
				DeltaX: float64(totalX),
				DeltaY: float64(deltaY),
			}
			err := p.Do(ctx)
			if err != nil {
				return err
			}

		}

		return nil
	})
}
