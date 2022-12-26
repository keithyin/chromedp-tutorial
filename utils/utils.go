package utils

import (
	"context"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
)

func ReadCookies(filename string) [][]string {

	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(fmt.Sprintf("%s", err))
	}

	cookies := string(bytes)
	splitted_cookies := strings.Split(cookies, "; ")
	cookies_pairs := make([][]string, 0)
	for _, cookie := range splitted_cookies {
		cookies_pairs = append(cookies_pairs, strings.SplitN(cookie, "=", 2))
	}
	return cookies_pairs
}

func SetCookie(cookies_pairs [][]string, domain string) chromedp.Tasks {
	tasks := chromedp.Tasks{

		chromedp.ActionFunc(func(ctx context.Context) error {
			expr := cdp.TimeSinceEpoch(time.Now().Add(180 * 24 * time.Hour))
			for _, cookie := range cookies_pairs {
				err := network.SetCookie(cookie[0], cookie[1]).
					WithExpires(&expr).
					WithDomain(domain).
					WithHTTPOnly(true).
					Do(ctx)
				if err != nil {
					fmt.Println("set cookies error: ", err)
					return err
				}
			}
			return nil
		}),
	}
	return tasks
}
