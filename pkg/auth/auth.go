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

var userContexts sync.Map // Map to store uuid -> *UserContext

func HandleSignIn(uuid string, userName string) error {
	_, loaded := userContexts.LoadOrStore(uuid, &UserContext{userName: userName, signedIn: true})
	if loaded {
		// Idempotent operation.
		log.Printf("User %s is already signed in", userName)
	}
	return nil
}

func HandleSignOut(uuid string) error {
	_, ok := userContexts.Load(uuid)
	if !ok {
		return fmt.Errorf("no user is signed in from this connection")
	}
	userContexts.Delete(uuid)
	return nil
}

func HandleWhoAmI(uuid string) (string, error) {
	user, ok := userContexts.Load(uuid)
	if !ok {
		return "", fmt.Errorf("no user is signed in from this connection")
	}
	userCtx := user.(*UserContext)
	return userCtx.userName, nil
}
