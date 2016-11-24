package main

import (
	"fmt"
	"log"
)

func logE(format string, v ...interface{}) {
	log.Printf("\033[1;31mERROR:\033[0m " + fmt.Sprintf(format, v...));
}

func logW(format string, v ...interface{}) {
        log.Printf("\033[1;31mWARN:\033[0m  " + fmt.Sprintf(format, v...));
}

func logI(format string, v ...interface{}) {
        log.Printf("\033[1;36mINFO:\033[0m  " + fmt.Sprintf(format, v...));
}

func logD(format string, v ...interface{}) {
        log.Printf("\033[1;35mDEBUG:\033[0m " + fmt.Sprintf(format, v...));
}
