/* ========================================================================
   $File: $
   $Date: $
   $Revision: $
   $Creator: Tristan Magniez $
   $Notice: (C) Copyright 2016 by Tristan Magniez. All Rights Reserved. $
   ======================================================================== */

package main

import (
    "gitla.in/Doc0160/asm.go.git"
    "time"
    "./imports/PluginSystem"
    "./imports/Command"
    "gopkg.in/vmihailenco/msgpack.v2"
    "bytes"
)

var PLuginsLogger = GlobalLogger.Sublog("Plugins")
var PreloaderLogger = GlobalLogger.Sublog("Preloader")

func GetCachedConfig(cp * Command.Parser) (PluginSystem.Config, error) {
    var err error
    var config PluginSystem.Config

    if entry, err := cache.Get(cp.Command+".json"); err == nil {
        buf := bytes.NewBuffer(entry)
        err = msgpack.NewDecoder(buf).Decode(&config)
        CacheLogger.Println(cp.Command+".json loaded")
        return config, err
    }

    config, err = PluginSystem.GetConfig(cp)
    if config.CopyOf != nil {
        config, err = GetCachedConfig(&Command.Parser{
            Command: *config.CopyOf,
            Paramaters: cp.Paramaters,
        })
    }

    Pool.SendWorkTimedAsync(time.Second, func(){
        var network bytes.Buffer
        msgpack.NewEncoder(&network).Encode(config)
        cache.Set(cp.Command+".json", network.Bytes())
        CacheLogger.Println(cp.Command+".json stored")
    }, nil)
    return config, err
}

func launch(cp * Command.Parser) (string, error) {
    var s = asm.RDTSC()
    defer func(){
        PLuginsLogger.Println(cp.Command, cp.Paramaters,
            Units(asm.RDTSC()-s)+"c")
    }()
    
    config, err := GetCachedConfig(cp)
    if err!= nil {
        PLuginsLogger.Println(err.Error())
    }

    PLuginsLogger.Println(config)
    
    if err == nil && config.MemoryType == PluginSystem.Preloadable {
        var out string
        if entry, err := cache.Get(cp.Command); err == nil {
            out = string(entry)
        }

        Pool.SendWorkTimedAsync(time.Second, func() {
            out, err := PluginSystem.Exec(cp)
            if err != nil {
                PLuginsLogger.Println(err)
            }
            cache.Set(cp.Command, out)
            PreloaderLogger.Println(cp.Command + " result preloaded")
        }, nil)

        if out != "" {
            return out, nil
        }
    }

    if err == nil && config.MemoryType == PluginSystem.Cacheable {
        if entry, err := cache.Get(cp.Command); err == nil {
            cache.Set(cp.Command, entry)
            return string(entry), nil
        }
    }
    cp.Command = config.ExeName
    out, err := PluginSystem.Exec(cp)

    if err == nil && config.MemoryType == PluginSystem.Cacheable {
        cache.Set(cp.Command, out)
        CacheLogger.Println(cp.Command + " result cached")
    }
    
    return string(out), err
}

