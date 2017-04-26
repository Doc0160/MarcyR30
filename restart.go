/* ========================================================================
   $File: $
   $Date: $
   $Revision: $
   $Creator: Tristan Magniez $
   $Notice: (C) Copyright 2016 by Tristan Magniez. All Rights Reserved. $
   ======================================================================== */

package main

import (
    "os"
    "os/exec"
    "syscall"
)

func restart(arg string) error {
    argv0, err := lookPath()
    if nil != err {
        return err
    }
    wd, err := os.Getwd()
    if nil != err {
        return err
    }
    
    args := []string{os.Args[0], arg}
    
    _, err = os.StartProcess(argv0, args, &os.ProcAttr{
        Dir:   wd,
        Env:   os.Environ(),
        Files: []*os.File{
            os.Stdin,
            os.Stdout,
            os.Stderr,
        },
        Sys:   &syscall.SysProcAttr{},
    })
    return err
}

func lookPath() (argv0 string, err error) {
	argv0, err = exec.LookPath(os.Args[0])
	if nil != err {
		return
	}
	if _, err = os.Stat(argv0); nil != err {
		return
	}
	return
}
