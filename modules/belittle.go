package modules

import (
	lazlo "github.com/djosephsen/lazlo/lib"
	"math/rand"
	"time"
)

var Belittle = &lazlo.Module{
    Name: `Belittle`,
    Usage: `%BOTNAME belittle <username>`,
    Run: belittleUser,
}

func belittleUser(broker *lazlo.Broker) {
    cb := broker.MessageCallback(`(?i)(belittle) (@*[\w-]+)`, true)
    for {
        select {
        case newReq := <-cb.Chan:
            go makeFunOfUser(b, newReq)
        }
    }
}

func makeFunOfUser(b *lazlo.Broker, req lazlo.PatternMatch) {
    lazlo.Logger.Info("Making fun of a user: %s", req[2])
}

func randomShitSaying() string {
    now := time.Now()
    rand.Seed(int64(now.Unix()))
    replies := []String{
        " you suck.",
    }

    return replies[rand.Intn(len(replies)-1)]
}
