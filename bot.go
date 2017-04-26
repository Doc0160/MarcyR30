/* ========================================================================
   $File: $
   $Date: $
   $Revision: $
   $Creator: Tristan Magniez $
   $Notice: (C) Copyright 2016 by Tristan Magniez. All Rights Reserved. $
   ======================================================================== */

package main

import (
    "fmt"
    "./imports/Command"
    "github.com/bwmarrin/discordgo"
    "os"

    "github.com/boltdb/bolt"
    "math/rand"
)

type ChatInfo struct {
    Token string
    ID string
    OwnerID string
}

type Bot struct {
    Discord ChatInfo

    commandParser *Command.Parser
}

var Marcy = Bot {
    Discord: ChatInfo{
        Token: "MjUyMTYxMjQzNzQ0MzA1MTUy.C2qXbw.0YlGKkWvz1MFRSp_NnGQhFEmOlo",
        ID: "252161243744305152",
        OwnerID: "252159849108865024",
    },
    commandParser: &Command.Parser{},
}

func (bot * Bot)messageCreatePlugins(s *discordgo.Session, m *discordgo.MessageCreate) {
    r, err := launch(bot.commandParser)
    if err != nil {
        PLuginsLogger.Println(m.ChannelID, err.Error() + " " +
            string(bot.commandParser.Command) + " " +
            fmt.Sprintf("%v", bot.commandParser.Paramaters))
    }else{
        s.ChannelMessageSend(m.ChannelID, r)
    }
}

func (bot * Bot)messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
    //var m discordgo.MessageCreate = *_m
    if m.Author.ID == bot.Discord.ID {
        return
    }
    defer Recover()
    PLuginsLogger.Println(m.Author.Username + " " + m.Content)
 
    if bot.commandParser.Parse(m.Content) {
        switch bot.commandParser.Command {
            
        case "ver":
            s.ChannelMessageSend(m.ChannelID, "Marcy R30 V" + Version + " b" + Build)
            
        case "marcy":
            err := database.View(func(tx *bolt.Tx) error {
                b := tx.Bucket([]byte("greetings"))
                i := btoi(b.Get([]byte("i")))
                r := uint64(rand.Intn(int(i)))
                s.ChannelMessageSend(m.ChannelID,
                    string(b.Get(itob(r))))
                return nil
            })
            if err != nil {
                s.ChannelMessageSend(m.ChannelID, err.Error())
            }

        case "myid":
            s.ChannelMessageSend(m.ChannelID, m.Author.ID)
            
        default:
            if m.Author.ID == bot.Discord.OwnerID {
                switch bot.commandParser.Command {

                case "debug":
                    DebugInfo.LoadAll()
                    s.ChannelMessageSend(m.ChannelID, DebugInfo.Text())

                case "freememory":
                    DebugInfo.FreeOSMemory()
                    
                case "restart":
                    GlobalLogger.Println("Bye")
                    s.ChannelMessageSend(m.ChannelID, "I'm gonna restart")
                    err := restart(m.ChannelID)
                    if err != nil {
                        s.ChannelMessageSend(m.ChannelID, err.Error())
                    } else {
                        s.UpdateStatus(0, "")
                        //Pool.Close()
                        os.Exit(1)
                    }

                case "stop":
                    GlobalLogger.Println("Bye")
                    s.ChannelMessageSend(m.ChannelID, "I'm gonna stop")
                    s.UpdateStatus(0, "")
                    os.Exit(1)
                    
                default:
                    bot.messageCreatePlugins(s, m)
                    
                }
            } else {
                bot.messageCreatePlugins(s, m)
            }
            
        }
    }
}

