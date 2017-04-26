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
    "math/rand"
    "time"
)

var classe = []string{
    "Maxence(tentacules)",
    "Maxence(pas tentacules)",
    "Thomas",
    "Joshua",
    "Gaelle",
    "Guillaume",
    "Steeven",
    "Antoine",
    "Etienne",
    "Greg",
}

func getSomeone() string {
    rand.Seed(time.Now().UnixNano())
    return classe[rand.Int() % len(classe)]
}

func main(){
    PluginSystem.SendText(getSomeone())
}

