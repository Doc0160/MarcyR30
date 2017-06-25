/* ========================================================================
   $File: $
   $Date: $
   $Revision: $
   $Creator: Tristan Magniez $
   $Notice: (C) Copyright 2016 by Tristan Magniez. All Rights Reserved. $
   ======================================================================== */

package main

import (
    "gopkg.in/resty.v0"
    "github.com/bwmarrin/discordgo"
    "./imports/TheLog"
    "os"
)

type WitResp struct {
    MsgID string `json:"msg_id"`
    Text string `json:"_text"`
    Entities struct {
        Intent []struct {
            Confidence float64 `json:"confidence"`
            Value string `json:"value"`
        } `json:"intent"`
    } `json:"entities"`
}

var Version string = "(undefined)"
var Build string = "(undefined)"
var GlobalLogger = TheLog.New("Global")
var BotID = "252161243744305152"

func init() {
    GlobalLogger.Println("Marcy R30 V" + Version + " b" + Build)
}

func main(){
    var err error
    
    discord, err := discordgo.New("Bot "+Marcy.Discord.Token)
    if err != nil {
        GlobalLogger.Fatal(err)
    }
    
	discord.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate){
        if m.Author.ID == BotID {
            return
        }
        entry, err := cache.Get(m.Content)
        var result string
        if(err != nil) {
            resp, err := resty.R().
                SetQueryParams(map[string]string{
                "q": m.Content,
            }).
                SetHeader("Accept", "application/json").
                SetAuthToken("V5DLKYOR6PP2QDJSZ4ZFVMNWM27EJAAS").
                SetResult(&WitResp{}).
                Get("https://api.wit.ai/message")
            if err != nil {
                GlobalLogger.Fatal(err)
            }
            _result := resp.Result().(*WitResp)
            if len(_result.Entities.Intent) > 0 {
                cache.Set(m.Content, []byte(_result.Entities.Intent[0].Value))
                result = _result.Entities.Intent[0].Value
            }
        } else {
            result = string(entry)
        }
        switch(result) {
        case "cats":
            cats(s, m)
        default:
            s.ChannelMessageSend(m.ChannelID, result)
        }
        
    })
    
	err = discord.Open()
	if err != nil {
        GlobalLogger.Fatal(err)
	}

    if len(os.Args) > 1 {
        discord.ChannelMessageSend(os.Args[1], "I restarted")
        discord.UpdateStatus(0, "Restarted V" + Version + " b" + Build)
    } else {
        discord.UpdateStatus(0, "Started V" + Version + " b" + Build)
    }
    
	GlobalLogger.Println("Bot is now running.")

    select {}
    return
}
