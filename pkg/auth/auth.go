package auth

import (
	"fmt"
	"log"
	"sync"
)

type UserContext struct {
	userName string
	signedIn bool
}

var userContexts sync.Map // Map to store connHash -> *UserContext

func HandleSignIn(connHash string, userName string) error {
	_, loaded := userContexts.LoadOrStore(connHash, &UserContext{userName: userName, signedIn: true})
	if loaded {
		// Idempotent operation.
		log.Printf("User %s is already signed in", userName)
	}
	return nil
}

func HandleSignOut(connHash string) error {
	_, ok := userContexts.Load(connHash)
	if !ok {
		return fmt.Errorf("no user is signed in from this connection")
	}
	userContexts.Delete(connHash)
	return nil
}

func HandleWhoAmI(connHash string) (string, error) {
	user, ok := userContexts.Load(connHash)
	if !ok {
		return "", fmt.Errorf("no user is signed in from this connection")
	}
	userCtx := user.(*UserContext)
	return userCtx.userName, nil
}
