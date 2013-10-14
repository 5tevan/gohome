package main

import(
	"crypto/md5"
	"fmt"
	"io"
)

const salt = "bananamillion"

type User struct {
	id int32
	name string
	password string
}


/** Create a new user entry in the database
* information comes from form post data
*
*	@return error if there is an error
*/
func (u *User) createUser() error {
	return nil
}

//turn a raw sring password into a hash and save it to the user
func (u *User) cipherPass(raw_pass string) string {
	p := md5.New()
	io.WriteString(p, raw_pass)
	io.WriteString(p, salt)
	io.WriteString(p, u.name)

	return fmt.Sprintf("%x", p.Sum(nil))
}