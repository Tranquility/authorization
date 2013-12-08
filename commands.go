package main

import (
        "fmt"
        "net/http"
        "log"
        "io/ioutil"
        "os"
        "regexp"
        "os/user"
)

func authorizedKeysPath() string {
    usr, err := user.Current()
    if err != nil {
        log.Fatal( err )
    }
    return fmt.Sprintf("%s/.ssh/authorized_keys", usr.HomeDir)
}

func Add(username string) {
        url := fmt.Sprintf("https://github.com/%s.keys", username)
        res, err := http.Get(url)
        if err != nil {
                log.Fatal(err)
        }
        keys, err := ioutil.ReadAll(res.Body)
        res.Body.Close()
        if err != nil {
                log.Fatal(err)
        }
        keyText := string(keys[:])
        toFile(fmt.Sprintf("#---%s---\n%s\n#------\n",username, keyText))
}

func Remove(username string) {
        data, err := ioutil.ReadFile(authorizedKeysPath())
        if err != nil {
                panic(err)
        }
        keyText := string(data[:])
        regex := fmt.Sprintf("(?Us)#---%s.*#------\n", username)
        re := regexp.MustCompile(regex)
        byteArray := []byte(re.ReplaceAllString(keyText, ""))
        ioutil.WriteFile(authorizedKeysPath(), byteArray, 0644)
}


func toFile(text string) error {
        f, err := os.OpenFile(authorizedKeysPath(), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
        if err != nil {
                panic(err)
        }

        defer f.Close()


        if _, err = f.WriteString(text); err != nil {
                panic(err)
        }
        fmt.Printf("The keys have been written to %s\n", authorizedKeysPath())

        return nil
}
