package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"
)

type User struct {
	Name  string
	Phone string
	Age   int
}

type newUser struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func Adder(values ...int) {
	total := 0
	for _, val := range values {
		total += val
	}
	fmt.Println(total)
}

func (u *User) GetData(s int) {
	fmt.Println(u.Age)
	fmt.Println("Parameter :", s)
}

func runAddFunctionGenerics[T int | float64 | string](a, b T) T {
	return a + b
}

func main() {

	hi := "akipiD"
	fmt.Println(hi)

	// reading input
	// reader := bufio.NewReader(os.Stdin)

	// comma , ok || ok , error
	// input, _ := reader.ReadString('\n')
	// fmt.Println(input)
	// fmt.Printf("The Type %T", input)

	var ptr *int

	var ptr1 *int

	myNumber := 50

	ptr = &myNumber

	ptr1 = ptr

	*ptr1 = *ptr1 + 5

	fmt.Println(*ptr)

	fmt.Println(*ptr1)

	fmt.Println(myNumber)

	var datas [4]string

	datas[0] = "Harshit"

	fmt.Println(datas)
	fmt.Println(len(datas))

	var sliceDataList = []string{}

	sliceDataList = append(sliceDataList, "A", "B")

	fmt.Println(sliceDataList)

	akipiD := make([]int, 4)

	akipiD[0] = 1
	akipiD[1] = 4
	akipiD[2] = 3
	akipiD[3] = 7

	sort.Ints(akipiD)

	fmt.Println(sort.IntsAreSorted(akipiD))

	akipiD = append(akipiD[:2], akipiD[2+1:]...)

	languages := make(map[int]string)

	languages[1] = "Java"
	languages[2] = "Go"

	fmt.Println(languages)

	// delete(languages, 2)

	for _, value := range languages {
		fmt.Println(value)
	}

	data := User{"Harshit", "73833", 10}
	// data.Age = 5

	fmt.Println(data)

	if data.Age == 5 {
		fmt.Println("Yes")
	} else if data.Age > 10 {
		fmt.Println("Hogaya")
	} else {
		fmt.Println("Done !!")
	}

	diceNo := rand.Intn(6) + 1

	switch diceNo {
	case 1:
		fmt.Println("1")
		fallthrough
	case 2:
		fmt.Println("2")
	default:
		fmt.Println("default")
	}

	var s int = 3

	for i := 0; i < s; i++ {
		fmt.Println(strconv.Itoa(i) + "-" + "akipiD")
	}

	result := runAddFunctionGenerics(1, 2)
	fmt.Println(result)

	for _, data := range akipiD {
		fmt.Println(data)
	}

	age := 32

	var agePointer *int

	agePointer = &age

	fmt.Println("Age:", *agePointer)

	reduceAge(agePointer)
	fmt.Println(age)

	Adder(1, 2, 4, 6)

	data.GetData(5)

	file, err := os.Create("./a.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, "Harshit")

	if err != nil {
		panic(err)
	}

	fmt.Println(length)

	defer file.Close()

	fileName("a.txt")

	var appUser newUser

	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	appUser = newUser{
		firstName: userFirstName,
		lastName:  userLastName,
		birthDate: userBirthdate,
		createdAt: time.Now(),
	}

	outputUserDetails(&appUser)

}

func reduceAge(age *int) {
	*age = *age - 18
}

func fileName(filename string) {
	databye, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(databye))
}

func outputUserDetails(u *newUser) {
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
