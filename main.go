package main

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"log"
	"os"
)

func main() {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var res string
	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://www.gaokao.cn/school/140`),
		// 等待特定元素加载完成，确保页面已经渲染
		chromedp.WaitVisible(`html`, chromedp.ByQuery),
		chromedp.OuterHTML(`body`, &res),
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
	file, err := os.OpenFile("example.html", os.O_RDWR|os.O_CREATE, 0666)
	fmt.Fprintln(file, res)
}
