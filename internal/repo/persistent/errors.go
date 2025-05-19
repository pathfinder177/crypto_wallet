package persistent

import "fmt"

var (
	errLoginNoUser          = fmt.Errorf("no such user, please register")
	errLoginNoMatchPassword = fmt.Errorf("password does not match")

	errRegUserExists            = fmt.Errorf("user exists")
	errRegBcryptGenFromPassword = fmt.Errorf("err: bcrypt GenerateFromPassword")
)
