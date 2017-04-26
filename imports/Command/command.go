/* ========================================================================
   $File: $
   $Date: $
   $Revision: $
   $Creator: Tristan Magniez $
   $Notice: (C) Copyright 2016 by Tristan Magniez. All Rights Reserved. $
   ======================================================================== */

package Command

import (

)

type Parser struct {
    Command string
    Paramaters []string
    i int
}

func (c * Parser) eatSpace(commandString string){
    for c.i < len(commandString) && commandString[c.i] == ' ' {
        c.i++
    }
}

func (c * Parser) eatText(commandString string) string {
    var tmp []byte
    for c.i < len(commandString) && commandString[c.i] != ' ' {
        tmp = append(tmp, commandString[c.i])
        c.i++
    }
    return string(tmp)
}

func (c * Parser) Parse(commandString string) bool {
    *c = Parser{}
    if commandString[c.i] == '$' {
        c.i++
        c.eatSpace(commandString)
        c.Command = c.eatText(commandString)
        for c.i < len(commandString) && commandString[c.i] != ' ' {
            c.Command += string(commandString[c.i])
            c.i++
        }
        for c.i < len(commandString) {
            c.eatSpace(commandString)
            if c.i < len(commandString) {
                c.Paramaters = append(c.Paramaters, c.eatText(commandString))
            }
        }
        return true
    }
    return false
}
