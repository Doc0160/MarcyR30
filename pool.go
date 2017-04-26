/* ========================================================================
   $File: $
   $Date: $
   $Revision: $
   $Creator: Tristan Magniez $
   $Notice: (C) Copyright 2016 by Tristan Magniez. All Rights Reserved. $
   ======================================================================== */

package main

import (
    "github.com/jeffail/tunny"
)

var PoolLogger = GlobalLogger.Sublog("Pool")
var Pool *tunny.WorkPool

func init() {
    var err error
    Pool, err = tunny.CreatePoolGeneric(2).Open()
    if err != nil {
        PoolLogger.Fatalln(err)
    }
}
