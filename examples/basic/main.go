package main

import (
	"fmt"
	webview "github.com/webview/webview_go"
	"time"
)

const tagA = `function interceptClickEvent(e) {
    let href = '';
    const target = e.target || e.srcElement;
    console.log(target,target.tagName);
    if (target.tagName === 'A') {
        href = target.getAttribute('href');
    }
    else if(target.tagName === 'IMG') {
        href = target.parentElement.getAttribute('href');
    }
    if(href) {
		window.navigate(href);
        //window.location.href = href;
        //console.log(href);
        e.preventDefault();
		return false;
    }
}

if (document.addEventListener) {
    document.addEventListener('click', interceptClickEvent);
} else if (document.attachEvent) {
    document.attachEvent('onclick', interceptClickEvent);
}`

func main() {
	w := webview.New(true)
	w.Navigate("https://buyin.jinritemai.com/dashboard/live/control")
	w.Init(tagA)
	//w.Navigate("https://baidu.com")
	//w.Navigate("https://github.com/webview/webview/blob/master/webview.h")
	//w.Window()
	//w.Eval(tagA)
	defer w.Destroy()
	err := w.Bind("getCookie", func(cookie string) {
		fmt.Println(cookie)
	})
	if err != nil {
		fmt.Println(err)
	}
	err = w.Bind("open", func(url string) {
		w.Eval("window.location.href='" + url + "'")
		fmt.Println("window.location.href='" + url + "'")
	})
	if err != nil {
		fmt.Println(err)
	}
	err = w.Bind("navigate", func(url string) {
		fmt.Println("navigate", url)
		w.Navigate(url)
	})
	//w.SetTitle("Basic Example")
	w.SetSize(480*2, 320*2, webview.HintNone)
	go func() {
		time.Sleep(time.Second * 4)
		//w.Eval(tagA)
		//w.SetHtml("Thanks for using webview!")
		//w.Eval("window.open('http://www.baidu.com')")
		//w.Eval("window.open('https://github.com/webview/webview')")
	}()
	//w.SetHtml("Thanks for using webview!")
	w.Run()
}
