/* ========================================================================
   $File: $
   $Date: $
   $Revision: $
   $Creator: Tristan Magniez $
   $Github: Doc0160 $
   $Notice: (C) Copyright 2017 by Tristan Magniez. All Rights Reserved. $
   ======================================================================== */

package main

import (
    "github.com/bwmarrin/discordgo"
    "gopkg.in/resty.v0"
)

var FuncLogger = GlobalLogger.Sublog("Func")

func cats(s *discordgo.Session, m *discordgo.MessageCreate) {
    Pool.SendWorkAsync(func(){
        type randomCat struct {
            File string `json:"file"`
        }
        resp, err := resty.R().
            SetResult(&randomCat{}).
            Get("http://random.cat/meow")
        if err != nil {
            FuncLogger.Fatal(err)
        }
        result := resp.Result().(*randomCat)
        s.ChannelMessageSend(m.ChannelID, result.File)
    }, nil)
}
