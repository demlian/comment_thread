package auth

import (
	"errors"
	"log"
	"sync"
)

type UserContext struct {
	Username      string
	TCP5TupleHash string
}

var userContexts sync.Map // Map to store requestID -> UserContext

func HandleSignIn(connHash string, userName string) error {
	if userName == "" {
		return errors.New("cannot sign in; missing user name")
	}
	if connHash == "" {
		return errors.New("cannot sign in; missing connection information")
	}
	_, ok := userContexts.Load(string(connHash))
	if ok {
		return errors.New("a user is already signed in with this connection")
	}

	userContexts.Store(connHash, &UserContext{Username: userName, TCP5TupleHash: connHash})
	return nil
}
func HandleSignOut(connHash string) error {
	_, ok := userContexts.Load(connHash)
	if !ok {
		return errors.New("no user is signed in with this connection")
	}

	userContexts.Delete(connHash)
	return nil
}
func HandleWhoAmI(connHash string) string {
	ctx, ok := userContexts.Load(connHash)
	if !ok {
		log.Println("no user is signed in from this client")
		return ""
	}
	userCtx := ctx.(*UserContext)
	return "|" + userCtx.Username
}
