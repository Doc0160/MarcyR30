/* ========================================================================
   $File: $
   $Date: $
   $Revision: $
   $Creator: Tristan Magniez $
   $Notice: (C) Copyright 2016 by Tristan Magniez. All Rights Reserved. $
   ======================================================================== */

package main

import (
    "runtime"
    "runtime/debug"
    "strconv"
)

var DebugLogger = GlobalLogger.Sublog("Debug")

type _DebugInfo struct {
    NumCPU int
    NumGoroutine int
    GC debug.GCStats
}

var DebugInfo _DebugInfo

func init() {
    DebugInfo.NumCPU = runtime.NumCPU()
    DebugInfo.NumGoroutine = runtime.NumGoroutine()
    runtime.GOMAXPROCS(16)
    debug.SetGCPercent(250)
    debug.SetMaxThreads(25)
    debug.ReadGCStats(&DebugInfo.GC)
}

func (d * _DebugInfo) FreeOSMemory() {
    debug.FreeOSMemory()
}

func (d * _DebugInfo) Text() (text string) {
    text = strconv.Itoa(DebugInfo.NumGoroutine) + " goroutine(s)\n" +
        strconv.Itoa(DebugInfo.NumCPU) + " CPU(s)\n" +
        "GC:\n\tTotal GC: " +
        strconv.FormatFloat(float64(DebugInfo.GC.PauseTotal)/1000000,
        'f', -1, 64) +
            " ms\n" +
            "\tNumGC: " +
            strconv.Itoa(int(DebugInfo.GC.NumGC)) + "\n" +
            "Pool goroutine:\n" +
            "\tNumWorkers: " +
            strconv.Itoa(Pool.NumWorkers()) + "\n" +
            "\tNumPendingAsyncJobs: " +
            strconv.Itoa(int(Pool.NumPendingAsyncJobs())) + "\n" +
            ""
    return
}

func (d * _DebugInfo) LoadAll() {
    d.NumCPU = runtime.NumCPU()
    d.NumGoroutine = runtime.NumGoroutine()
}

func (d * _DebugInfo) LoadGCStats() {
    debug.ReadGCStats(&DebugInfo.GC)
}

