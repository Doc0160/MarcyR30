/* ========================================================================
   $File: $
   $Date: $
   $Revision: $
   $Creator: Tristan Magniez $
   $Notice: (C) Copyright 2016 by Tristan Magniez. All Rights Reserved. $
   ======================================================================== */

package main

import (
    "encoding/json"
    "../../imports/httpclient"
    "../../imports/PluginSystem"
) 

type randomCat struct {
	File string `json:"file"`
}

var err error
var t randomCat
func main(){
    res, err := httpclient.
        Get("http://random.cat/meow", nil)
    if err != nil {
        PluginSystem.SendError(err)
    }
    err = json.NewDecoder(res.Body).Decode(&t)
    PluginSystem.SendResult(t.File, err)
}

