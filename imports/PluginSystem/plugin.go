/* ========================================================================
   $File: $
   $Date: $
   $Revision: $
   $Creator: Tristan Magniez $
   $Notice: (C) Copyright 2016 by Tristan Magniez. All Rights Reserved. $
   ======================================================================== */

package PluginSystem

import (
    "../Command"
    "os/exec"
    "encoding/json"
    msgpack "gopkg.in/vmihailenco/msgpack.v2"
    "io/ioutil"
    "bytes"
    "errors"
)

var InvalidArgument = errors.New("Invalide argument")

type MemoryType int
const (
    None MemoryType = iota // 0
    Cacheable = iota
    Preloadable = iota
)

type Config struct {
	Name string `json:"name"`
    ExeName string `json:"exe_name"`
	Description string `json:"description"`
    MemoryType int `json:"memory"`
    CopyOf *string `json:"copy_of"`
}

type Result struct {
    Text []byte
    Error error
}

func GetList() ([]string, error) {
    var l []string
    files, err := ioutil.ReadDir("./plugins/")
    for _, f := range files {
        if f.IsDir() {
            l = append(l, f.Name())
        }
    }
    return l, err
}

func Exec(cp * Command.Parser) ([]byte, error) {
    cmd := exec.Command("./plugins/"+cp.Command, cp.Paramaters...)
    out, err := cmd.CombinedOutput()
    buf := bytes.NewBuffer(out)
    r := Result{}
    err = msgpack.NewDecoder(buf).Decode(&r)
    return r.Text, err
}

func GetConfig(cp * Command.Parser) (Config, error) {
    var c = Config{}
    file, err := ioutil.ReadFile("./plugins/"+cp.Command+".json")
    if err != nil {
        return c, err
    }
    err = json.Unmarshal(file, &c)
    return c, err
}

func SendResult(text string, err error) {
    var network bytes.Buffer
    var r Result = Result{
        Text: []byte(text),
        Error: err,
    }
    msgpack.NewEncoder(&network).Encode(r)
    println(string(network.Bytes()))
}

func SendError(err error) {
    SendResult("", err)
}

func SendText(text string) {
    SendResult(text, nil)
}

