# 1. How to run Go

1. in terminal, cd to current file directory.

2. type: go run filename.go   (you can type 'go' in terminal to check whether go is properly installed)

3. Files in the same package do not have to be imported into each other. Can simply run by: 
    ```go 
    go run main.go lib1.go
    ```

<br/><br/>




# 2. Code template

## A. declare variable:
```go
// types: bool / string / int / float64 ...
foo := "First string"  
// this is same as : 
var foo string = "First string"
// := only used in the first var declare. Next assign can directly use "="

// can also first init and later assign (purly initialize a variable outside of a function is valid)
var foo int
foo = 52
```
<br/>


## B. function & OO in Go:
1. create and call function:
    ```go
    import "fmt"

    func main() {
        foo := newFunc()
        fmt.Println(foo)  //or jsut:  fmt.Println(newFunc())
    }

    // need to declare return type of the function
    func newFunc() string {
        return "this is string"
    }
    ```
    ```go
    // can also assign variables to function
    import "strconv"
    func newFunc2(text string, num int) string {
        return text + ": " + strconv.Itoa(num)
    }
    ```

    ```go
    // can also return multiple results from funcion
    func newFunc2(text string, num int) (string, int) {
        return text, num+1
    }
    ```

2. Go is not an OOP. To achieve similar concept in Go, you can create a type and assign relevant methods into it:
    ```go
    // Example 1:
    type customSliceType []string

    // receiver of new type:
    /*
    with define 'customSliceType' in function:
    every variable inside the package with type customSliceType can get 'print' method.

    c here is the receiver argument. i.e. customSlice in this case.const
    this similar to self in python, this in js
    usually it is just one character which represents the the value of the defined type
    */
    func (c customSliceType) print() {
        for i, n := range c {
            fmt.Println(i, n)
        }
    }

    func main() {
        customSlice := customSliceType{"string1", "string2"}
        customSlice.print()
    }

    ```

    ```go
    // Example 2:
    type book string

    func (b book) toSentence() book {
        return "this is my book: " + b  
    }

    func main() {
        var b book = "Harry Potter"
        fmt.Println(b.toSentence())
    }
    ```
<br/>


## C. Data Structure in Go:
1. Array: fixed length list of things
2. Slice: An array that can grow or shrink. Every element in a slice must be of same type

    ```go
    func main() {
        my_slice := []string{"string1", "string2"}
        my_slice = append(my_slice, "string3") 
        // !! append does not modify existed slice, but create a new one.
    }
    ```

    slice index starts from 0.
    ```go
    my_slice := []string{"string1", "string2", "string3", "string4"}
    my_slice[0:2] // this return ["string1", "string2"]
    my_slice[:2] // this return ["string1", "string2"]
    my_slice[2:] // this return ["string3", "string4"]

    ```

    ```go
    // create empty map:
    // way 1:
    var colors []byte{}  // replace byte by any type you want
    // way 2:
    colors := make([]byte, 99999) // 99999 means create 99999 empty space (i.e. default value) in the empty slice
    ```


    When slice is created, it contains two parts:
    
    a. slice: this part includes: pointer to head item of the array, capacity, length
    
    b. array with items in this slice
    
    ** a and b will be in different containers of the memory
    
    ** When slice is assigned to a function, only "a" part will be copy to anotehr memory container but it is still point to same "b".
    
    Thus, we can directly modify the item is slice (Reference Type)

    ```go
    // Value Type: int, float, string, bool, structs

    // Reference Type: slices, maps, channels, pointers, functions
    ```
    <br/>




3. Struct: a collection in go. similar to dictionary in python
    ```go
    type contactInfo struct {
        email   string
        zipCode int
    }

    type person struct {
        firstName string
        lastName  string
        contactInfo
    }

    func main() {
        // can also do: jim := person{"Jim", "Party", contactInfo{"jim@gmail.com", 94000}}

        /* can also do: 
           var jim person
           jim.firstName = "Jim"
           jim.lastName = "Party"*/
        
        jim := person{
            firstName: "Jim",
            lastName:  "Party",
            contactInfo: contactInfo{
                email:   "jim@gmail.com",
                zipCode: 94000,  //*** in go, each item in struct should have comma in the end
            },
        }

        fmt.Printf("%+v", jim)
    }

    // reciever in struct
    /* note that in go, it will copy the variable to another container in the memory 
        when you assign the variable to the function. (pass by value language)
        so here we need to get the memory address of the variable first (by &), pass it
        into receiver, and in receiver change this address back to value (by *)

        &variable: this means giving me the memory address (pointer) of the value this variable is pointing at
        *pointer: this means giving me the value this memory address is pointing at
    */


    jimPointer := &jim
    jimPointer.updateName("jimmy")

    // above two lines are same as:
    jim.updateName("jimmy")

    /* jimPointer is type of *person (pointer to a person)
       jim is type of person
       if reciever is type of pointer to x, then Go also accept you pass type x into it
       and it will auto transform type x to pointer to x
    */

    
    /* note: *person: *type is different concept than *pointer
       it means we are working with a pointer to a person 
    */
    func (pointerToPerson *person) updateName(newFirstName string) {
        (*pointerToPerson).firstName = newFirstName
    }


    //Similar example:
    func main() {
        name := "Bill"
        updateValue(&name)
    }

    func updateValue(n *string) {
        *n = "Alex"
    }

    ```

4. Map: also a collection in go. also similar to dictionary in python
   
   In map, all keys should be in the same type; all values should also be in the same type.

   But value and key does not necessary to be the same type

    ```go
    //map[keyType]valueType{}
    colors := map[string]string{
    "red":   "#ff0000",
    "green": "#4bf745",
    "white": "#ffffff",
	}


    // create empty map:
    // way 1:
    var colors map[string]string
    // way 2:
    colors := make(map[string]string)

    // Add items:
    colors["yellow"] = "laksjdf"
    // Delete items:
    delete(colors, "yellow")

    // iterate map:
    func printMap(c map[string]string) {
        for color, hex := range c {
            fmt.Println("Hex code for", color, "is", hex)
        }
    }

    printMap(colors)

    ```

    Map vs Struct:
    
    1. all keys/valus must bbe the same type ; values can be different type

    2. keys are indexed (can iterate) ; keys don't support indexing

    3. reference type ; value type

    4. don't need to know all keys at compile time ; need to know all the different fields at compile time

    5. represent a collection of related properties ; represent a "thing" with a lot of differnt properties


 5. interfaces:
    
    this is used when you need a general funcion which needs to accept variables with different type.

    For example, sometimes you need to feed int variable to function A, but sometimes need to feed string to function A

    ```go
    type cusType interface {
        cusReceiverFunction1(string, int) (string, error)  // input types & pit[ut types
        cusReceiverFunction2() string
    }

    func mainFunction(c cusType) {
        fmt.Println(c.cusReceiverFunction1())
    }
    // see details in the codes in interfaces folder 
    ```

<br/>


## D. Coding Operation:
1. itteration:
    ```go
    func main() {
        my_slice := []string{"string1", "string2"}

        for index, item := range my_slice {
            fmt.Println(index, item)
        }

    // each declared variable must be user. So, if index is unused, you must change code to:
        for _, item := range my_slice {
            fmt.Println(item)
        }
    }
    ```

    ```go
    for i := 0; i < (links); i++{
        fmt.Println("Hi")
    }
    ```

2. if statement:
    ```go
    import os
    if err != nil {
        fmt.Println("Error:", err)
        os.Exit(1)
    }
    ```






3. data structure manipulation:

    ```go
    // a. convert data type:  type(variabke)
    my_string := "Hi there"
    fmt.Println([]byte(my_string))
    ```

    ```go
    // b. join list of string:
    import strings
    strings.Join([]string("a1", "a2", "a3"), ",")
    ```

     ```go
    // b. split string:
    import strings
    strings.Split(string(bs), ",")  //string(bs) : just convert byte slice to string

    ```

4. random:
    ```go
    import (
	"math/rand"
	"time"
    )
    source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
    newNumber := r.Intn(len(foo) - 1)  //number between 0 - len(foo)-1

    ```

5. Concurrency:

    Concurrency: the program can execute (schedule) multiple threads. If one thread blocks, another one is picked up and worked on. (it does not matter whether you can 1 cpu core or multiple)

    ** Parallelism: Multiple threads executed at the exactly same time (Require multiple CPU)

    Channel: it used to communicate between Main Go rutine and its child rutines
    ```go
    c := make(chan string)
    ```

    ```go
    for _, link := range links {
		go checkLink(link, c) // 'go' will create and handel Go rutine
	}
    ```

    ```go
    // send data to channel
    func checkLink(link string, c chan string) {  // c channelType CommunicateType
        _, err := http.Get(link)
        if err != nil {
            fmt.Println(link, "might be down!")

            // c <- link: send x (e.g. link here) to channel
            c <- link  
            return
        }

        fmt.Println(link, "is up!")
        c <- link
    }
    ```

    ```go
    // receive data from channel

    // variable <- c: wait value to be sent to channel, then assign it to a variable
    // can also do fmt.Println(<- c)
    fmt.Println(<- c)

    /*
    <- c: is also a blocker code, the main go routine will sleep when waits for data from channel.
    it only wake up when the channel receive the data from first child routine.
    Thus, if you have n child routine, you have to do fmt.Println(<- c) n times as well in order to receive all data from channel
    */


    // scenario 1: print all in links 
    for i := 0; i < (links); i++{
        fmt.Println(<- c)
    }


    // scenario 2: continual loop whenever c send back something:
		for l := range c {
			go checkLink(link, c)
		}

        //this loops equal to:
        for {
			go checkLink(<- c, c)
		}

    // scenario 3: temp pause function:
    for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}    
    ```
    



## E. Testing:
1. create file ending in _test.go
2. to run all tests in a package, run: go test
3. function name should start with Test and pass variable 't *testing.T'
   ```go
    import "testing"

    func TestNewDeck(t *testing.T) {
        d := newDeck()
        if len(d) != 16 {
            t.Errorf("Expected deck length of 16, but got %v", len(d))
        }
    }
    ```






