package main

import (
	"fmt"
	"github.com/MarinX/keylogger"
	hook "github.com/robotn/gohook"
	"time"
)

func main() {

}

func skipLoopForTime() {
	skip := false
	for i := 0; i <= 10; i++ {
		time.Sleep(1 * time.Second)
		if skip {
			fmt.Printf("Skip.. %d\n", i)
			continue
		}

		fmt.Printf("Execute.. %d\n", i)
		skip = true
		go func() {
			time.Sleep(3 * time.Second)
			skip = false
		}()
	}
}

func listenEvent2() {
	fmt.Println("Looking for devices")
	fmt.Println(keylogger.FindKeyboardDevice())
	var devs = keylogger.FindAllKeyboardDevices()
	for dev := range devs {
		fmt.Printf("find device:%s\n\n", dev)
	}
	////our keyboard..on your system, it will be diffrent
	//rd := keylogger.NewKeyLogger(devs[3])
	//in, err := rd.Read()
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//for i := range in {
	//	//listen only key stroke event
	//	if i.Type == keylogger.EV_KEY {
	//		fmt.Println(i.KeyString())
	//	}
	//}
}

func listenEvent() {
	hook.Register(hook.KeyDown, []string{"enter"}, func(e hook.Event) {
		fmt.Printf("[Event](%d) Enter detected!\n", e.Rawcode)
	})

	s := hook.Start()
	defer hook.End()
	<-hook.Process(s)
}

func listenAllEvents() {
	chanHook := hook.Start()
	defer hook.End()

	for ev := range chanHook {
		fmt.Printf("hook: %v\n", ev)
	}
}
