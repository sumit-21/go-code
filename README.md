# go-code
Basic Primary concepts of golang for studying.

#GoLang#
- a statically-typed, open source language by google.
- faster than python, coding is slightly similar to C, not exactly object oriented but can be coded like that.

Syntaxes and basics:
1) Print: 
	- ```fmt.Println("hello world")```
2) Data types and variables:
	- https://www.geeksforgeeks.org/data-types-in-go/
	- https://gobyexample.com/variables
	- ```var name = "Sumit"```
	- ```var a, b = 14, 15```
	- ```var c int```			----> default value 0
	- ```d := "apple" ```			----> var d string = "apple"  (similar)
3) Constants:
	- ```const test = "testString"```
	- ```const s string = "constant"```
	- ```const n = 500000000```
	- ```const d = 3e20 / n```
	- ```int64(d)```                	----> converting numeric type to int
	- as per the context the variables type is taken into consideration
4) For Loop:
	- There is NO while loop in Go. Example:
	 ```
	 i := 1
	 for i <= 3 {
	 	fmt.Println(i)
          	i = i + 1
	 }
	 for j := 7; j <= 9; j++ {
	 	fmt.Println(j)
	 }
	 for {
	 	// infinite for
	 }
5) If/Else:
	```
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
          fmt.Println("7 is odd")
      	}
	if num := 9; num < 0 {
          fmt.Println(num, "is negative")
	}
6) Switch:
	- switch statement works on variables also.
	- switch works without condition also, we need to provide the condition in case statement. This is similar to if statement.
	- A type switch is for comparing the type of the values. Example:
	 ```whatAmI := func(i interface{}) {
        switch t := i.(type) {
        case bool:
            fmt.Println("I'm a bool")
        case int:
            fmt.Println("I'm an int")
        default:
            fmt.Printf("Don't know type %T\n", t)
        }
    }
    whatAmI(true)
	whatAmI(1)
    whatAmI("hey")
7) Arrays:
	- ```var a [5]int```
	- ```b := [5]int{1, 2, 3, 4, 5}```
	- ```nums := []int{2, 3, 4}```
	- ```var twoD [2][3]int```
8) Slices:
	- ```s := make([]string, 3)```			----> length can be found out like this: ```len(s)```
	- ```s = append(s, "e", "f")```
	- Copy the slice
	  ```
	  c := make([]string, len(s))
	  copy(c, s)
	- Slice operator
	  ```
	  l := s[2:5]
	  l = s[:5]
	  l = s[2:]
	- Two dimensional slices:
	  ```twoD := make([][]int, 3)
      for i := 0; i < 3; i++ {
          innerLen := i + 1
          twoD[i] = make([]int, innerLen)
          for j := 0; j < innerLen; j++ {
              twoD[i][j] = i + j
          }
      }
9) Maps:
	- ```m := make(map[string]int)```
	- Set key value : ```m["k1"] = 7```
	- Get key value : ```m["k1"]```     	----> returns two values, 7 and true
	- Find keys count 					----> len(m)
	- Delete key from					----> delete(m, "k2")
	- ```_, isPresent := m["k2"]```			----> isPresent would be equal to false
	- ```n := map[string]int{"foo": 1, "bar": 2}```
10) Range:
	Uses are like below:
	- ```nums := []int{2, 3, 4}
	  for _, num := range nums {
	  }
	  for i, num := range nums {
	  }
	- ```kvs := map[string]string{"a": "apple", "b": "banana"}
	  for k, v := range kvs {
		  fmt.Printf("%s -> %s\n", k, v)
	  }
	  for k := range kvs {
          fmt.Println("key:", k)
      }
	- Strings are iterated on the basis of their unicode:
	  ```
	  for i, c := range "go" {
          fmt.Println(i, c)				----> Prints 0 103 and 1 111
      }
11) Functions and multiple return values:
	- Function returning multiple values is like:
	  ```
	  func vals() (int, int) {
		  return 3, 7
	  }
	  a, b := vals()
	  _, c := vals()
	- One of the best examples could be swap function:
	  ```
	  func swap(x, y string) (string, string) {
		return y, x
	  }
	  func main() {
		a, b := swap("hello", "world")
		fmt.Println(a, b)
	  }
12) Variadic Functions:
	- Functions taking n number of inputs:
	  ```
	  func sum(nums ...int){}
	  sum(1,2)
	  sum(1,2,3)
	- Slices can also be passes as arguement to variadic functions
	  ```
	  nums := []int{1, 2, 3, 4}
	  sum(nums...)
13) Closures:
	- Functions returning another functions.
	  ```
	  func intSeq() func() int {
		i := 0
		return func() int {
				i++
				return i
		}
	  }
	  nextInt := intSeq()
	  fmt.Println(nextInt())
14) Recursion: same as in Java. Function calling itself.
15) Pointers:
	- Go uses pointers to pass variables by reference i.e. their memory locations.
	- Pass values with * and access them with &.
	  ```
	  func zeroptr(iptr *int) {
		  *iptr = 0
	  }
	  zeroptr(&i)
16) Structs:
	- Define a struct
	  ```
	  type person struct {
		  name string
		  age  int
	  }
	  fmt.Println(person{"Bob", 20})
	  fmt.Println(person{name: "Alice", age: 30})```
	- You can use a pointer for struct also.
	  ```
	  testP := person{name: "Tester"}
	  testPP := &testP
	  testPP.name = "Tester1"
	  testPP.age = 21
17) Methods:
	- On structs types, we can define methods in Go.
	  ```
	  type rect struct {
		  width, height int
	  }
	  func (r *rect) area() int {//mutable, as pointers
		return r.width * r.height
	  }
	  func (r rect) perim() int {// immuatable called directly
		return 2*r.width + 2*r.height
	  }
	  r.area()
	  r.perim()
18) Interfaces:
	- We can define interfaces with methods in Go.
	- In order for stuct to implement interface we need to implement all the methods on the struct.
	  ```
	  type geometry interface {
		  area() float64
		  perimeter() float64
	  }
19) Errors:
	- There are no exceptions in Go.
	- We can create new error in following way:
	  ```errors.New(fmt.Sprintf("%s %d %s", "can't divide", num, "by 0"))```
	- stucts can be used as custom errors. The struct should implement (define) the method Error().
	  ```
	  func (custErr *customError) Error() string {
		  return fmt.Sprintf("%d - %s", custErr.arg, custErr.err)
	  }
	- Custom errors can be taken into variable and used like below:
	 ```
	  _, err1 := DivideWithCustomError(60, 2)
	  custErr, boolVal := err1.(*customError)
	  fmt.Print(custErr, " ")
	  fmt.Println(boolVal)
20) Routines:
	- Goroutines are lightweight threads (asynchronous) of execution.
	- To run a function as routine specify it with "go".				----> ```go printData("test data")```
21) Channels:
	- Channels are used for data communication in between two or more routines.
	- chan is type provided by go for channel.
	- A channel is created with make keyword.						----> ```channel := make(chan string)```
	- Produce/send data with "<-" on the channel					----> ```channel <- "test data"```
	- Read data with with "<- channel" from the channel.			----> ```data := <- channel```
	- All "fatal error: all goroutines are asleep - deadlock!" errors occur, when all the the goroutines are finished and chennel is closed.
	- Channels can be used to make a routine to wait until other routines are finished. This is nothing but channel Synchronization.
22) Buffered Channels:
	- By default the channels are NOT buffered in Go. This means if data would be sent on channel only when there is concurrent receiver available to receive.
	- In Buffered channels, the above mentioned condition is not mandatory. So we can send the data and then receive it synchronously.
	- We would need to specify count of the values the channel is going to hold.
	  ```myBufferedChannel := make(chan string, 3)```
	- We would get error "fatal error: all goroutines are asleep - deadlock!", if we tried to send more values on the channel than specified count.
