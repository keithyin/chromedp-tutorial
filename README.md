
* 获取selector所选定的所有节点的 Text

```golang
texts := make([]string, 0)
chromedp.Evaluate(`[...document.querySelectorAll('#tags li')].map((e) => e.innerText)`, &texts)
// https://github.com/chromedp/chromedp/issues/87
```