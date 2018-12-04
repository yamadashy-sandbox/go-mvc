package main

import (
	"github.com/stivan622/kiokumushi-api/public"
	"google.golang.org/appengine"
)

func main() {
	public.Run()
	appengine.Main()
}
