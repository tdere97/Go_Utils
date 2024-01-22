package main

import (
	"errors"
	"fmt"
	"sort"
)

func mapFunc() {
	fmt.Println("===============================map.go=================================")
	// MAPS
	// Maps are similar to JavaScript objects, Python dictionaries, and Ruby hashes.
	// Maps are a data structure that provides key->value mapping.
	// The zero value of a map is nil.
	// We can create a map by using a literal or by using the make() function:
	//make(map[key]value)

	ages := make(map[string]int)
	ages["tej"] = 37
	ages["swap"] = 24
	ages["swap"] = 21 // overwrites 24

	salary := map[string]int{
		"tej":  200,
		"swap": 400,
	}
	fmt.Printf("Length of salary map : %v\n", len(salary))
	//fmt.Printf("Key is : %v Value is : %v\n", salary.key, salary.value)

	// INSERT AN ELEMENT
	// m[key] = elem
	salary["t"] = 500
	fmt.Printf("Length of salary map : %v\n", len(salary))

	// GET AN ELEMENT
	// elem = m[key]
	ele := salary["tej"]
	fmt.Printf("ele is : %v\n", ele)

	// DELETE AN ELEMENT
	// delete(m, key)
	delete(salary, "t")

	// CHECK IF A KEY EXISTS
	// elem, ok := m[key]
	// If key is in m, then ok is true. If not, ok is false.
	// If key is not in the map, then elem is the zero value for the map's element type.
	elem, ok := salary["t"]
	if !ok {
		fmt.Printf("Key not present\n")
	} else {
		fmt.Printf("Key is present :%v\n", elem)
	}
	fmt.Printf("Length of salary map : %v\n", len(salary))

	bootdev()

	fmt.Printf("Key Type :\nAny type can be used as the value in a map, but keys are more restrictive.\nmap keys may be of any type that is comparable\n")
	fmt.Printf("comparable types are boolean\n numeric\n string\n pointer\n channel\n interface types\n structs\n arrays\n that contain only those types.\n")
	fmt.Printf("absent from the list are slices, maps, and functions; these types cannot be compared using ==, and may not be used as map keys.\n")
}

func deleteIfNecessary(users map[string]user, name string) (deleted bool, err error) {
	existingUser, ok := users[name]
	if !ok {
		return false, errors.New("not found")
	}

	if !existingUser.scheduledForDeletion {
		return false, nil
	}

	delete(users, name)
	return true, nil
}

type user struct {
	name                 string
	number               int
	scheduledForDeletion bool
}

func test(users map[string]user, name string) {
	fmt.Printf("Attempting to delete %s...\n", name)
	defer fmt.Println("====================================")
	deleted, err := deleteIfNecessary(users, name)
	if err != nil {
		fmt.Println(err)
		return
	}
	if deleted {
		fmt.Println("Deleted:", name)
		return
	}
	fmt.Println("Did not delete:", name)
}

func bootdev() {
	users := map[string]user{
		"john": {
			name:                 "john",
			number:               18965554631,
			scheduledForDeletion: true,
		},
		"elon": {
			name:                 "elon",
			number:               19875556452,
			scheduledForDeletion: true,
		},
		"breanna": {
			name:                 "breanna",
			number:               98575554231,
			scheduledForDeletion: false,
		},
		"kade": {
			name:                 "kade",
			number:               10765557221,
			scheduledForDeletion: false,
		},
	}
	fmt.Printf("-------------------- Inside boot dev function --------------------\n")
	test(users, "john")
	test(users, "musk")
	test(users, "santa")
	test(users, "kade")

	keys := []string{}
	for name := range users {
		keys = append(keys, name)
	}
	sort.Strings(keys)

	fmt.Println("Final map keys:")
	for _, name := range keys {
		fmt.Println(" - ", name)
	}
}
