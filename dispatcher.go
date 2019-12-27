package dmt

import (
	"fmt"
	"time"
)

func (e *Engine) globalObserver(works []Work) {
	fmt.Println(time.Now())
	for i, w := range works {
		w.Type()
		fmt.Println(i, ":", w.IsDone())
	}
}
