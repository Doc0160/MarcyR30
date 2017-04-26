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
    "../../imports/httpclient"
    "os"
    "strings"
)
//http://tinyurl.com/api-create.php?url=http://www.google.fr
func shorten(url string) (string, error) {
    res, err := httpclient.
        Get("http://tinyurl.com/api-create.php?url="+url, nil)
    if err != nil {
        return url, err
    }
    shorturl, err := res.ToString()
    return shorturl, err
}

func main(){
    if strings.Contains(os.Args[1], "://"){
        r, e := shorten(os.Args[1])
        PluginSystem.SendResult(r, e)

    } else {
        PluginSystem.SendError(PluginSystem.InvalidArgument)
    }
}

