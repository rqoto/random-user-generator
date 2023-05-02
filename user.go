package main

import (
	"fmt"
	"math/rand"
	"time"
)

var Names = [...]string{"James", "Robert", "John", "Michael", "David", "William", "Richard", "Joseph", "Thomas", "Charles", "Christopher", "Daniel", "Matthew", "Anthony", "Mark", "Donald", "Steven", "Paul", "Andrew", "Joshua", "Kenneth", "Kevin", "Brian", "George", "Timothy", "Ronald", "Edward", "Jason", "Jeffrey", "Ryan", "Jacob", "Gary", "Nicholas", "Eric", "Jonathan", "Stephen", "Larry", "Justin", "Scott", "Brandon", "Benjamin", "Samuel", "Gregory", "Alexander", "Frank", "Patrick", "Raymond", "Jack", "Dennis", "Jerry", "Tyler", "Aaron", "Jose", "Adam", "Nathan", "Henry", "Douglas", "Zachary", "Peter", "Kyle", "Ethan", "Walter", "Noah", "Jeremy", "Christian", "Keith", "Roger", "Terry", "Gerald", "Harold", "Sean", "Austin", "Carl", "Arthur", "Lawrence", "Dylan", "Jesse", "Jordan", "Bryan", "Billy", "Joe", "Bruce", "Gabriel", "Logan", "Albert", "Willie", "Alan", "Juan", "Wayne", "Elijah", "Randy", "Roy", "Vincent", "Ralph", "Eugene", "Russell", "Bobby", "Mason", "Philip", "Louis", "Mary", "Patricia", "Jennifer", "Linda", "Elizabeth", "Barbara", "Susan", "Jessica", "Sarah", "Karen", "Lisa", "Nancy", "Betty", "Margaret", "Sandra", "Ashley", "Kimberly", "Emily", "Donna", "Michelle", "Carol", "Amanda", "Dorothy", "Melissa", "Deborah", "Stephanie", "Rebecca", "Sharon", "Laura", "Cynthia", "Kathleen", "Amy", "Angela", "Shirley", "Anna", "Brenda", "Pamela", "Emma", "Nicole", "Helen", "Samantha", "Katherine", "Christine", "Debra", "Rachel", "Carolyn", "Janet", "Catherine", "Maria", "Heather", "Diane", "Ruth", "Julie", "Olivia", "Joyce", "Virginia", "Victoria", "Kelly", "Lauren", "Christina", "Joan", "Evelyn", "Judith", "Megan", "Andrea", "Cheryl", "Hannah", "Jacqueline", "Martha", "Gloria", "Teresa", "Ann", "Sara", "Madison", "Frances", "Kathryn", "Janice", "Jean", "Abigail", "Alice", "Julia", "Judy", "Sophia", "Grace", "Denise", "Amber", "Doris", "Marilyn", "Danielle", "Beverly", "Isabella", "Theresa", "Diana", "Natalie", "Brittany", "Charlotte", "Marie", "Kayla", "Alexis", "Lori"}

type User struct {
	id     string `json:"ID"`
	Name   string `json:"Name"`
	Age    int    `json:"Age"`
	Phone  string `json:"Phone"`
	Gender string `json:"Gender"`
}

func CreateUser() *User {
	rand := rand.Intn(len(Names))

	U := User{
		id:     RandomID(),
		Name:   Names[rand],
		Age:    Age(),
		Gender: checkGender(rand),
		Phone:  fmt.Sprintf("(000) 000-%d", Number()),
	}
	return &U
}

func RandomID() string {
	return fmt.Sprintf("id-%d-%d", rand.Intn(99999-10000)+10000, rand.Uint64()+uint64(time.Now().Weekday()))
}

func Age() int {
	return rand.Intn(75-18) + 18
}

func checkGender(number int) string {
	gender := "Male"

	if number >= 100 {
		gender = "Female"
	}

	return gender
}

func Number() int {
	return rand.Intn(9999)
}
