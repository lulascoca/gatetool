package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func apiAccess(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

		// the actual password thats read from a file called passwd in the same directory

		//TODO CHANGE THIS PLEASE

		accessKey, _ := ioutil.ReadFile(webconfig.PasswdLocation)
		// password from user
		passwd := r.FormValue("pass")
		// door to open
		door := r.FormValue("door")

		if strings.Trim(string(accessKey), "\n ") != passwd {
			fmt.Fprintf(w, "nope\n")
			return
		}

		// check which door to open
		if door == "out" {
			OpenOutside(mainconfig.ScriptsLocation)
			fmt.Fprintf(w, "called outside successfully.\n")
			return
		} else if door == "in" {
			OpenInside(mainconfig.ScriptsLocation)
			fmt.Fprintf(w, "called inside successfully.\n")
			return
		}
		fmt.Fprintf(w, "the door specified doesn't exist\n")
	}
	// if method isnt post
	fmt.Fprint(w, "revamp bitch\n")
}
