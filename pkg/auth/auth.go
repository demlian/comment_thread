package auth

import (
	"fmt"
	"log"
	"sync"
)

type UserContext struct {
	clientId string
	signedIn bool
}

var userContexts sync.Map // Map to store uuid -> *UserContext

func HandleSignIn(uuid string, clientId string) error {
	_, loaded := userContexts.LoadOrStore(uuid, &UserContext{clientId: clientId, signedIn: true})
	if loaded {
		// Idempotent operation.
		log.Printf("User %s is already signed in", clientId)
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
	return userCtx.clientId, nil
}
