package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/chromedp/cdproto/input"
	"github.com/chromedp/chromedp"
	"github.com/keithyin/chromedp-tutorial/mousetools"
	"github.com/keithyin/chromedp-tutorial/utils"
)

func main() {

	cookies := utils.ReadCookies("cookies.txt")
	fmt.Println(cookies)

	ctx, cancel := chromedp.NewExecAllocator(context.Background(),
		append(chromedp.DefaultExecAllocatorOptions[:], chromedp.Flag("headless", false))...)
	defer cancel()
	ctx, cancel = chromedp.NewContext(ctx)
	defer cancel()
	err := chromedp.Run(ctx,
		// utils.SetCookie(cookies, "github.com"),
		chromedp.Navigate("https://news.baidu.com/"),
		chromedp.Sleep(2*time.Second),
		mousetools.MoveMouseMiddle(0, 4096, 10, 500*time.Millisecond),
		mousetools.MoveMouse(0, 0, 1024, 0, input.Left),
	)

	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(time.Second * 10)

	fmt.Println("hello")
}
