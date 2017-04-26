/* ========================================================================
   $File: $
   $Date: $
   $Revision: $
   $Creator: Tristan Magniez $
   $Notice: (C) Copyright 2016 by Tristan Magniez. All Rights Reserved. $
   ======================================================================== */

package main

import (
    "github.com/bwmarrin/discordgo"
    "./imports/TheLog"
    "os"
    "time"
    "image"
    "github.com/nfnt/resize"
    "gitla.in/Doc0160/platform.go.git"
)

var Version string = "(undefined)"
var Build string = "(undefined)"
var GlobalLogger = TheLog.New("Global")

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
        Pool.SendWorkAsync(func(){
            Marcy.messageCreate(s, m)
        }, nil)
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
    
    var FrameBuffer = image.NewNRGBA(image.Rect(0,0,1000,1000))
    MarcyAvatar, _ := discord.UserAvatar(Marcy.Discord.ID)
    tmp := resize.Resize(1000, 0, MarcyAvatar, resize.Lanczos3)
    for x := 0; x < tmp.Bounds().Dx(); x++ {
        for y := 0; y < tmp.Bounds().Dy(); y++ {
            FrameBuffer.Set(x, y, tmp.At(x, y))
        }
    }
    
    platform.InitDisplayLoop(100*2+40, 100*2+40,
        FrameBuffer.Bounds().Dx(), FrameBuffer.Bounds().Dy(),
        func(window *platform.WindowState) {
            var redraw = true
            var r = resize.Lanczos3
            var tmp = MarcyAvatar
            var x int
            var y int
            lastVBlankTime := time.Now()
            spent := time.Now().Sub(lastVBlankTime)
            toWait := 66 * time.Millisecond - spent
            for {

                window.Mutex.Lock()
                if window.CharIsDown('1') {
                    r = resize.NearestNeighbor
                    redraw = true
                } else if window.CharIsDown('2') {
                    r = resize.Bilinear
                    redraw = true
                } else if window.CharIsDown('3') {
                    r = resize.Bicubic
                    redraw = true
                } else if window.CharIsDown('4') {
                    r = resize.MitchellNetravali
                    redraw = true
                } else if window.CharIsDown('5') {
                    r = resize.Lanczos2
                    redraw = true
                } else if window.CharIsDown('6') {
                    r = resize.Lanczos3
                    redraw = true
                }
                window.Mutex.Unlock()

                if redraw {
                    redraw = false
                    tmp = resize.Resize(1000, 0, MarcyAvatar, r)
                    for x = 0; x < tmp.Bounds().Dx(); x++ {
                        for y = 0; y < tmp.Bounds().Dy(); y++ {
                            FrameBuffer.Set(x, y, tmp.At(x, y))
                        }
                    }
                    window.Mutex.Lock()
                    copy(window.Pix, FrameBuffer.Pix)
                    window.RequestDraw()
                    window.Mutex.Unlock()
                }
                
                spent = time.Now().Sub(lastVBlankTime)
                toWait = 17 * time.Millisecond - spent
                if toWait > time.Duration(0) {
                    <-time.NewTimer(toWait).C
                }
                lastVBlankTime = time.Now()
            }
        })
    return
}
