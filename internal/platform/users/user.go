package users

import (
	"fmt"
	"time"

	"github.com/dmedina2150dev/learn-go/modelos"
)

func AltaUsuario() {
	u := new(modelos.User)

	u.AddUser(10, "Pablo", time.Now(), true)

	fmt.Println(u)
}
