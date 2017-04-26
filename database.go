/* ========================================================================
   $File: $
   $Date: $
   $Revision: $
   $Creator: Tristan Magniez $
   $Notice: (C) Copyright 2016 by Tristan Magniez. All Rights Reserved. $
   ======================================================================== */

package main

import (
    "github.com/boltdb/bolt"
    "time"
    "strconv"
    "fmt"
    "encoding/binary"
    "runtime"
)
var DatabaseLogger = GlobalLogger.Sublog("Database")
var database *bolt.DB
func init() {
    var err error
    database, err = bolt.Open("marcy.db", 0600,
        &bolt.Options{Timeout: 1 * time.Second})
    if err != nil {
        DatabaseLogger.Fatal(err)
    }
    DatabaseLogger.Println("marcy.db opened")
    database.Update(func(tx *bolt.Tx) error {
        
        _, err := tx.CreateBucketIfNotExists([]byte("Options"))
        if err != nil {
            return fmt.Errorf("create bucket: %s", err)
        }
        
        _, err = tx.CreateBucketIfNotExists([]byte("Cache"))
        if err != nil {
            return fmt.Errorf("create bucket: %s", err)
        }

        err = tx.DeleteBucket([]byte("Store"))
        if err != nil {
            return fmt.Errorf("create bucket: %s", err)
        }
        
        b, err := tx.CreateBucketIfNotExists([]byte("greetings"))
        if err != nil {
            return fmt.Errorf("create bucket: %s", err)
        }
        
        var i uint64 = 0
        err = b.Put(itob(i), []byte("Rex!"))
        if err != nil {
            return fmt.Errorf("create bucket: %s", err)
        }
        i++
        err = b.Put(itob(i), []byte("Rex !"))
        if err != nil {
            return fmt.Errorf("create bucket: %s", err)
        }
        i++
        err = b.Put(itob(i), []byte("Rex !!!"))
        if err != nil {
            return fmt.Errorf("create bucket: %s", err)
        }
        i++
        err = b.Put(itob(i), []byte("I <3 REX !"))
        if err != nil {
            return fmt.Errorf("create bucket: %s", err)
        }
        i++
        err = b.Put([]byte("i"), itob(i))
        if err != nil {
            return fmt.Errorf("create bucket: %s", err)
        }
        
        return nil
    })
    runtime.SetFinalizer(database, func(p *bolt.DB) {
        DatabaseLogger.Println("Closing")
        p.Close()
    })
}

func itob(v uint64) []byte {
    b := make([]byte, 8)
    binary.BigEndian.PutUint64(b, v)
    return b
}

func btoi(b []byte) uint64 {
    n := binary.BigEndian.Uint64(b)
    return n
}

type FizzBuzzThing struct {
    Word string
    Number int
}

func FizzBuzz(min int, max int, things []FizzBuzzThing){
    if min > max {
        min, max = max, min
    }
    var s = ""
    var tmp = ""

    for i := min; i < max; i++ {
        s = ""
        for _, v := range things {
            if i % v.Number == 0 {
                s += v.Word
            }
            tmp = strconv.Itoa(i)
            for _, v2 := range tmp {
                if string(v2) == strconv.Itoa(v.Number) {
                    s += v.Word
                }
            }
        }
        if s == "" {
            s = strconv.Itoa(i)
        }
        println(s)
    }
}
