/* ========================================================================
   $File: $
   $Date: $
   $Revision: $
   $Creator: Tristan Magniez $
   $Notice: (C) Copyright 2016 by Tristan Magniez. All Rights Reserved. $
   ======================================================================== */

package main

import (
    "github.com/allegro/bigcache"
    "time"
)

var CacheLogger = GlobalLogger.Sublog("Cache")
var cache *bigcache.BigCache = nil

func init() {
    var err error
    cache, err = bigcache.NewBigCache(bigcache.DefaultConfig(10 * time.Minute))
    if err != nil {
        PLuginsLogger.Fatal(err)
    }
    CacheLogger.Println("cache inited")
}
