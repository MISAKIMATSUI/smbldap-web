package main

import (
	"fmt"
	"net/http"
	"os/exec"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, guestbookform)
	usr := r.FormValue("uid")
	pass := r.FormValue("passwd")
	fmt.Fprint(w, usr)
	fmt.Fprint(w, pass)
	inputunixcommand(usr, pass)
}

func inputunixcommand(usr string, pass string) {
	out, err := exec.Command(usr, pass).Output()
	if err != nil {
		fmt.Println("ERROR")
	}
	fmt.Printf("%s", out)
}

const guestbookform = `
<html>
<body>
make LDAP account
<form action="/make_account" method="post">
<label> User name: <input type="text" name="uid"> </label>
<label> Password: <input type="text" name="passwd"> </label>
<input type="submit" name="post">
</form>
</body>
</html>

`
