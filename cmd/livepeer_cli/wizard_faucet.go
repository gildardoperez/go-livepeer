package main

import (
	"fmt"
)

func (w *wizard) requestTokens() {
	httpPost(fmt.Sprintf("http://%v:%v/requestTokens", w.host, w.httpPort))
}
