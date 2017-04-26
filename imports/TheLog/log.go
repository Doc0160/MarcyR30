/* ========================================================================
   $File: $
   $Date: $
   $Revision: $
   $Creator: Tristan Magniez $
   $Notice: (C) Copyright 2016 by Tristan Magniez. All Rights Reserved. $
   ======================================================================== */

package TheLog

import (
    "log"
    "os"
    "fmt"
    "github.com/shiena/ansicolor"
)

type Logger struct {
    Name string

    logger *log.Logger
}

var Colors = []string {
	//"\x1b[30m", 	
	"\x1b[31m", 	
	"\x1b[32m", 	
	"\x1b[33m", 	
	"\x1b[34m", 	
	"\x1b[35m", 	
	"\x1b[36m", 	
	"\x1b[37m",
}
var BackColors = []string{
	//"\x1b[40m",
	"\x1b[41m", 
	"\x1b[42m", 
	"\x1b[43m", 
	"\x1b[44m", 
	"\x1b[45m", 
	"\x1b[46m", 
	"\x1b[47m",
}
var LighterOn      = "\x1b[1m"
var LighterOff     = "\x1b[21m"
var BackLighterOn  = "\x1b[5m"
var BackLighterOff = "\x1b[25m"
var Clear = "\x1b[0m"

func New(name string) *Logger {
    l := &Logger{Name : name}
    out := ansicolor.NewAnsiColorWriter(os.Stdout)
    l.logger = log.New(out, "", log.LstdFlags)
    return l
}

func (l * Logger) Sublog(name string) *Logger {
    var nl = *l;
    nl.Name = name
    return &nl
}

func (l * Logger) output(b string) {
    var x = int(l.Name[0]) + (int(l.Name[1]) << 8)
    l.logger.SetPrefix(LighterOff +
        Colors[x % len(Colors)])
    var a = LighterOn + l.Name + ": " +
        LighterOff
    l.logger.Output(4, a + b + Clear)
}

func (l * Logger) Print(v ...interface{}) {
    l.output(fmt.Sprint(v...))
}

func (l * Logger) Println(v ...interface{}) {
    l.output(fmt.Sprintln(v...))
}

func (l * Logger) Fatalln(v ...interface{}) {
    l.output(fmt.Sprintln(v...))
	os.Exit(1)
}

func (l * Logger) Fatal(v ...interface{}) {
    l.output(fmt.Sprint(v...))
	os.Exit(1)
}
