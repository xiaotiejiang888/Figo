package Figo

import (
	"log"
	"runtime/debug"
)

func Catch() {
	err := recover()
	log.Println(string(debug.Stack()))
	log.Println(err, " (recover)")
}
