/* ========================================================================
   $File: $
   $Date: $
   $Revision: $
   $Creator: Tristan Magniez $
   $Notice: (C) Copyright 2016 by Tristan Magniez. All Rights Reserved. $
   ======================================================================== */

package main

import (
    "os"
    "log"
    "runtime/debug"

    "strconv"
)

var FileNotExist = os.ErrNotExist

func Recover() {
	if r := recover(); r != nil {
		log.Println("Recovered:", string(debug.Stack()))
	}
}

func Units(v uint64) string {
    var i = 0;
    for v > 1000 {
        v = v/1000
        i++
    }
    var r = strconv.Itoa(int(v))
    r += " "
    var a = [...]string{"", "K", "M", "G"}
    r += a[i]
    return r
}
