package main

import "time"

func main() {
	Sleep(5)
}

func Sleep(d int) {
	<-time.After(time.Second * time.Duration(d))
}
