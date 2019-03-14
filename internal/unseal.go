package unsealer

import (
    "encoding/json"
    "fmt"
    "http"
    "time"
) 


type unsealparams struct {
    Keys *[]string `json:"key"`
    reset bool
    migrate bool
}


type Unsealer struct {
    ID int
    UnsealQueue chan UnsealRequest
    ManageChan chan int
    LogChan <- chan string
    params *unsealparams
}

func NewUnsealer(id int, unsealQueue chan UnsealRequest, logChan chan string, *up unsealparams) Unsealer {
    unsealer := Unsealer {
        ID: id
        UnsealQueue : unsealQueue
        LogChan logchan
        params: up
    }

    return unsealer
}

func (u *Unsealer) Start() {
    go func {
        UnsealRequestSuccesful := chan bool  // Track unseal request
        // Wait until there is an unseal request
        for unsealRequest := range u.UnsealQueue {
            // Performing the unsealing request

            



func unsealCall

func checkStatus(url) (bool, error) {
    // Checks the vault status
    var unsealed bool
    resp, err := http.Head(url)
    if err != nil {
        fmt.Println(err)
        return (unsealed, err)
    }
    if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
        unsealed = true
    } else {
        // Status not in the 200, something is wrong
        unsealed = false
        err = error('Status code returned: ' + resp.StatusCode)
    }
    return unsealed, err
}

func unsealCall(url string, payload unsealparams) (status int, err error) {
    /* Making an unseal call for an instance */
    // Converting the payload to a byte array

    req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonBytesPayload)
    
    // Cleaning key from memory
    jsonBytesPayload = nil

    client := &http.Client{}
    resp, err := client.Do(req)
    // Making sure body is closed
    defer resp.Body.Close()
    if err != nil {
        return -1, err
    }
    return resp.StatusCode, err
}

func instanceMonitor(vault *Vault, instance *Instance, status chan int, keys *[]string) {
    // Using status to send status update to main process
    // Creating the urls to use
    statusUrl := *vault.protocl + "://" + *instance.Domain + *vault.StatusPath
    unsealUrl := *vault.protocol + "://" + *instance.Domain + *vault.UnsealPath

    for {
        status, err := checkStatus(statusUrl)
        if err != nil {
            fmt.Println(err)
        } else if !unsealed {
            // This instance is sealed and needs to be unsealed
            for i := 0;i < len(keys); i++ {
                payload := unsealparams {
                    Key: &(*keys[i])
                    reset: false
                    migrate: false
                }
                go unsealCall(unsealUrl, payload)
                // Create payload
                // Do a put (post?) request
            }
        }
        time.Sleep(*vault.CheckInterval * time.Millisecond)            
    }
}

