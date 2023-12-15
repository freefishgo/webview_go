package main

import (
	"fmt"
	webview "github.com/webview/webview_go"
	"time"
)

const html = `<button id="increment">Tap me</button>
<div>You tapped <span id="count">0</span> time(s).</div>
<a href="https://blog.csdn.net/gisdtcn/article/details/131409006" target="_blank">webview</a>
<script>
  const [incrementElement, countElement] =
    document.querySelectorAll("#increment, #count");
  document.addEventListener("DOMContentLoaded", () => {
    incrementElement.addEventListener("click", () => {
      window.increment(countElement.textContent).then(result => {
        countElement.textContent = result.count;
      });
    });
  });
</script>`

type IncrementResult struct {
	Count uint `json:"count"`
}

func main() {
	var count uint = 0
	w := webview.New(true)
	w.Init(``)
	defer w.Destroy()
	w.SetTitle("Bind Example")
	w.SetSize(480, 320, webview.HintNone)

	// A binding that increments a value and immediately returns the new value.
	w.Bind("increment", func(abs string) IncrementResult {
		fmt.Println(count)
		count++
		w.Eval("console.log('increment')")
		return IncrementResult{Count: count}
	})
	w.Bind("open", func(url string) {
		w.Eval("window.location.href='" + url + "'")
	})
	go func() {
		time.Sleep(time.Second * 3)
		//		w.Eval(`
		//function interceptClickEvent(e) {
		//    let href = '';
		//    const target = e.target || e.srcElement;
		//    console.log(target,target.tagName);
		//    if (target.tagName === 'A') {
		//        href = target.getAttribute('href');
		//    }
		//    else if(target.tagName === 'IMG') {
		//        href = target.parentElement.getAttribute('href');
		//    }
		//    if(href) {
		//        window.location.href = href;
		//        console.log(href);
		//        e.preventDefault();
		//    }
		//}
		//
		//if (document.addEventListener) {
		//    document.addEventListener('click', interceptClickEvent);
		//} else if (document.attachEvent) {
		//    document.attachEvent('onclick', interceptClickEvent);
		//}
		//`)
		//w.Eval("window.increment('sdfsdfds')")
	}()

	w.SetHtml(html)
	w.Run()
}
