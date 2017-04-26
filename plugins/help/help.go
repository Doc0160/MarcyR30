/* ========================================================================
   $File: $
   $Date: $
   $Revision: $
   $Creator: Tristan Magniez $
   $Notice: (C) Copyright 2016 by Tristan Magniez. All Rights Reserved. $
   ======================================================================== */

package main

import (
    "../../imports/PluginSystem"
    "strconv"
)

func main(){
    var s string
    l, e := PluginSystem.GetList()
    s += "" + strconv.Itoa(len(l)) + " commndes:\n"
    for _, p := range l {
        s += "\t- " +p + "\n"
    }
    PluginSystem.SendResult(s, e)
}
