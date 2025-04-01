package basic

import (
	"log"
)

func LogDemo() {
	log.SetFlags(log.Lshortfile | log.Lmicroseconds | log.Lmsgprefix | log.Ldate | log.Ltime)
	//log.SetFlags(log.Lmsgprefix) // 无日志前缀
	log.Println("message")
	//log.New();
	//log.SetOutput()
	// TODO Writer ?
}
