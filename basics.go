package main

import "fmt"

func main() {

	// data types
	var intVar int = 42                 // Integer type
	var floatVar float64 = 3.14         // Floating-point type
	var stringVar string = "Hello, Go!" // String type
	var boolVar bool = true             // Boolean type
	fmt.Println("Integer:", intVar)
	fmt.Println("Float:", floatVar)
	fmt.Println("String:", stringVar)
	fmt.Println("Boolean:", boolVar)

	// naming rules
	// Valid variable names
	var validName int = 100
	var _validName int = 200 // Underscore is allowed
	var validName2 int = 300 // Can have numbers but not at the start
	fmt.Println("camelCase, PascalCase, and snake_case are common naming conventions in Go.")

	// Invalid variable names (uncommenting these will cause errors)
	// var 2invalidName int = 400 // Cannot start with a number
	// var invalid-name int = 500 // Hyphens are not allowed
	// var invalid name int = 600 // Spaces are not allowed
	// var int int = 700 // Cannot use reserved keywords as variable names
	fmt.Println("Valid variable names:")
	fmt.Println("validName:", validName)
	fmt.Println("_validName:", _validName)
	fmt.Println("validName2:", validName2)

	// formating verbs - Some common formatting verbs in Go
	// %d - integer, %f - float, %s - string, %t
	// %T - type, %x - hexadecimal, %e - scientific notation, %c - character, %p - pointer
	// %q - quoted string, %b - binary representation, %v - default format
	// %10s - width of 10 for strings, %.2f - precision of 2 for floats
	fmt.Println("Formatting verbs:")
	fmt.Printf("Integer: %d\n", intVar)                                     // %d for integers
	fmt.Printf("Float: %.2f\n", floatVar)                                   // %.2f for
	fmt.Printf("String: %s\n", stringVar)                                   // %s for strings
	fmt.Printf("Boolean: %t\n", boolVar)                                    // %t for boolean
	fmt.Printf("Type: %T\n", intVar)                                        // %T for type of variable
	fmt.Printf("Hexadecimal: %x\n", intVar)                                 // %x for hexadecimal representation
	fmt.Printf("Scientific notation: %e\n", floatVar)                       // %e for scientific notation
	fmt.Printf("Character: %c\n", 'A')                                      // %c for characters
	fmt.Printf("Pointer: %p\n", &intVar)                                    // %p for pointer addresses
	fmt.Printf("Width and precision: |%10s|%10.2f|\n", stringVar, floatVar) // Width and precision formatting
	// fmt.Printf("Invalid format: %q\n", intVar) // %q is for quoted strings, not integers (uncommenting will cause an error)
	fmt.Printf("Invalid format: %q\n", stringVar) // %q for quoted strings
	// fmt.Printf("Invalid format: %b\n", floatVar) // %b is for binary representation, not floats (uncommenting will cause an error)
	fmt.Printf("Invalid format: %b\n", intVar) // %b for binary representation of integers
	// fmt.Printf("Invalid format: %f\n", boolVar) // %f is for floats, not booleans (uncommenting will cause an error)

	// variables
	var a int = 10 // works in/out of a function
	b := 20        // shorthand notation, only works inside a function

	fmt.Println("Value of a:", a)
	fmt.Println("Value of b:", b)

	// constants / multiple constants
	const (
		c int    = 30
		d string = "Hello, World!"
	)
	fmt.Println("Value of c:", c)
	fmt.Println("Value of d:", d)

	// Arrays - declared length
	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array:", arr)

	// Arrays - shorthand notation + undeclared length
	arr2 := [...]int{1, 2, 3, 4, 5}
	fmt.Println("Array 2:", arr2)

	// Arrays - Initialize Only Specific Elements
	arr3 := [5]int{2: 3, 4: 5} // Only initializes indices 0, 2, and 4
	fmt.Println("Array 3:", arr3)

	// Length of an array
	fmt.Println("Length of arr3:", len(arr3))

	// Slices - dynamic length
	slice := []int{6, 7, 8, 9, 10}
	fmt.Println("Slice:", slice)

	// Slices - len and cap
	slice2 := []int{11, 12, 13}
	fmt.Println("Length of slice2:", len(slice2))   // Length is the number of elements in the slice
	fmt.Println("Capacity of slice2:", cap(slice2)) // Capacity is the number of elements the slice can hold before needing to allocate more memory

	// Create a slice from an array
	arr4 := [5]int{1, 2, 3, 4, 5}
	slice3 := arr4[1:4] // Slice from index 1 to 3 (4 is exclusive)
	fmt.Println("Slice from array:", slice3)

	// create slice with make
	// slice_name := make([]type, length, capacity)
	slice4 := make([]int, 3, 10) // Creates a slice of length 3 with zero values
	fmt.Println("Slice created with make:", slice4)

	// Slices - append
	slice5 := []int{1, 2, 3}
	slice5 = append(slice5, 4, 5) // Appending elements to the slice
	fmt.Println("Slice after append:", slice5)

	// Slices - copy
	slice6 := []int{6, 7, 8}
	slice7 := make([]int, len(slice6)) // Create a new slice with the same length
	copy(slice7, slice6)               // Copy elements from slice6 to slice7
	fmt.Println("Slice after copy:", slice7)

	// operators
	a = 10
	b = 20
	sum := a + b
	sub := a - b
	multiply := a * b
	divide := a / b
	modulus := a % b
	a++
	increment := a
	b--
	decrement := b
	fmt.Println("Sum of a and b:", sum)
	fmt.Println("Subtraction of a and b:", sub)
	fmt.Println("Multiplication of a and b:", multiply)
	fmt.Println("Division of a and b:", divide)
	fmt.Println("Modulus of a and b:", modulus)
	fmt.Println("Increment of a:", increment)
	fmt.Println("Decrement of b:", decrement)

	// logical operators
	isTrue := true
	isFalse := false
	fmt.Println("Logical AND (isTrue && isFalse):", isTrue && isFalse)
	fmt.Println("Logical OR (isTrue || isFalse):", isTrue || isFalse)
	fmt.Println("Logical NOT (!isTrue):", !isTrue)

	// bitwise operators
	bitA := 5                                            // 0101 in binary
	bitB := 3                                            // 0011 in binary
	fmt.Println("Bitwise AND (bitA & bitB):", bitA&bitB) // 0001 in binary, which is 1
	fmt.Println("Bitwise OR (bitA | bitB):", bitA|bitB)  // 0111 in binary, which is 7
	fmt.Println("Bitwise XOR (bitA ^ bitB):", bitA^bitB) // 0110 in binary, which is 6
	fmt.Println("Bitwise NOT (^bitA):", ^bitA)           // Inverts
	fmt.Println("Left Shift (bitA << 1):", bitA<<1)      // 1010 in binary, which is 10
	fmt.Println("Right Shift (bitA >> 1):", bitA>>1)     // 0010 in binary, which is 2

	// if-else, else if, nested if
	if a > b {
		if a > 15 {
			fmt.Println("a is greater than b and also greater than 15")
		} else {
			fmt.Println("a is greater than b but not greater than 15")
		}
		fmt.Println("a is greater than b")
	} else if a < b {
		fmt.Println("a is less than b")
	} else {
		fmt.Println("a is equal to b")
	}

	// switch-case - single case
	switch a {
	case 10:
		fmt.Println("a is 10")
	case 20:
		fmt.Println("a is 20")
		// switch-case - multiple cases
	case 30, 40:
		fmt.Println("a is either 30 or 40")
	default:
		fmt.Println("a is neither 10 nor 20")
	}

	// switch-case - expression
	switch {
	case a > b:
		fmt.Println("a is greater than b in switch-case")
	case a < b:
		fmt.Println("a is less than b in switch-case")
	default:
		fmt.Println("a is equal to b in switch-case")
	}

	// for loop - basic
	for i := 0; i < 5; i++ {
		fmt.Println("For loop iteration:", i)
	}
	// for loop - range over array
	for index, value := range arr {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}
	// for loop - range over slice
	for index, value := range slice {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}
	// for loop - infinite loop with break
	for {
		fmt.Println("This is an infinite loop. Breaking now.")
		break // Breaks the infinite loop
	}
	// for loop - continue
	for i := 0; i < 5; i++ {
		if i == 2 {
			continue // Skips the iteration when i is 2
		}
		fmt.Println("For loop with continue, iteration:", i)
	}
	// while loop - using for
	for i := 0; i < 5; i++ {
		if i == 3 {
			fmt.Println("Breaking the while loop at i =", i)
			break // Breaks the loop when i is 3
		}
		fmt.Println("While loop iteration:", i)
	}
	// do-while loop - using for
	i := 0
	for {
		fmt.Println("Do-while loop iteration:", i)
		i++
		if i >= 5 {
			break // Breaks the loop when i is 5 or more
		}
	}
	// functions - basic
	sumFunction := func(x, y int) int {
		return x + y // Returns the sum of x and y
	}
	fmt.Println("Sum of 5 and 10:", sumFunction(5, 10))
	// functions - multiple return values
	multiReturnFunction := func(x, y int) (int, int) {
		return x + y, x - y // Returns the sum and difference of x and y
	}
	sum, diff := multiReturnFunction(20, 10)
	fmt.Println("Sum:", sum, "Difference:", diff)

	// functions - named return values
	namedReturnFunction := func(x, y int) (sum int, diff int) {
		sum = x + y
		diff = x - y // Named return values, can be used without explicit return statement
		return       // Implicit return using named return values
	}
	resultSum, resultDiff := namedReturnFunction(30, 15)
	fmt.Println("Named return values - Sum:", resultSum, "Difference:", resultDiff)

	// functions - variadic function
	variadicFunction := func(nums ...int) int {
		sum := 0
		for _, num := range nums {
			sum += num // Sums all variadic arguments
		}
		return sum
	}

	fmt.Println("Sum of variadic function:", variadicFunction(1, 2, 3, 4, 5))
	// functions - anonymous function
	anonymousFunc := func(x int) int {
		return x * x // Returns the square of x
	}
	fmt.Println("Square of 4 using anonymous function:", anonymousFunc(4))

	// functions - closures
	counter := func() func() int {
		count := 0 // Closure variable
		return func() int {
			count++ // Increments count each time the returned function is called
			return count
		}
	}
	countFunc := counter()                     // Create a new counter
	fmt.Println("Counter value:", countFunc()) // 1

	// functions - defer
	deferFunc := func() {
		fmt.Println("This will be printed after the function returns, but before the program exits.")
	}
	defer deferFunc() // Defer the execution of deferFunc until the surrounding function returns
	fmt.Println("This is printed before the deferred function.")

	// panic and recover
	panicFunc := func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered from panic:", r) // Recover from panic
			}
		}()
		panic("This is a panic message!") // Trigger a panic
	}
	panicFunc() // Call the function that panics
	fmt.Println("Program continues after panic recovery.")

	// structs
	type Person struct {
		Name string
		Age  int
	}
	person := Person{Name: "Alice", Age: 30} // Create a new Person struct
	fmt.Println("Person Name:", person.Name)
	fmt.Println("Person Age:", person.Age)
	// structs - anonymous struct
	anonymousPerson := struct {
		Name string
		Age  int
	}{Name: "Bob", Age: 25} // Create an anonymous struct
	fmt.Println("Anonymous Person Name:", anonymousPerson.Name)
	fmt.Println("Anonymous Person Age:", anonymousPerson.Age)

	// type structs
	type Employee struct {
		Name          string
		Position      string
		Salary        float64
		DateOfJoining string
	}
	var emplyee1 Employee // Declare an Employee variable
	emplyee1.Name = "John Doe"
	emplyee1.Position = "Manager"
	emplyee1.Salary = 60000.00
	emplyee1.DateOfJoining = "2022-05-01"
	fmt.Println("Employee Name:", emplyee1.Name)
	fmt.Println("Employee Position:", emplyee1.Position)
	fmt.Println("Employee Salary:", emplyee1.Salary)
	fmt.Println("Employee Date of Joining:", emplyee1.DateOfJoining)

	// type structs - shorthand notation
	employee := Employee{
		Name:          "Charlie",
		Position:      "Software Engineer",
		Salary:        75000.00,
		DateOfJoining: "2023-01-15",
	}
	fmt.Println("Employee Name:", employee.Name)
	fmt.Println("Employee Position:", employee.Position)
	fmt.Println("Employee Salary:", employee.Salary)
	fmt.Println("Employee Date of Joining:", employee.DateOfJoining)

	// maps
	mapExample := map[string]int{
		"apple":  5,
		"banana": 10,
		"orange": 15,
	}
	fmt.Println("Map Example:", mapExample)
	// maps - accessing elements
	fmt.Println("Number of apples:", mapExample["apple"])
	// maps - adding elements
	mapExample["grape"] = 20
	fmt.Println("Map after adding grape:", mapExample)

	// maps - deleting elements
	delete(mapExample, "banana") // Deletes the key "banana" from the map
	fmt.Println("Map after deleting banana:", mapExample)
	// maps - checking existence of a key
	if value, exists := mapExample["orange"]; exists {
		fmt.Println("Number of oranges:", value)
	} else {
		fmt.Println("Oranges not found in the map")
	}

	// maps - iterating over a map
	for fruit, quantity := range mapExample {
		fmt.Printf("Fruit: %s, Quantity: %d\n", fruit, quantity)
	}
	// maps - length of a map
	fmt.Println("Length of mapExample:", len(mapExample))
	// maps - empty map
	emptyMap := make(map[string]int) // Creates an empty map
	fmt.Println("Empty map:", emptyMap)
	// maps - map with zero values
	zeroValueMap := map[string]int{"apple": 0, "banana": 0, "orange": 0} // Map with zero values
	fmt.Println("Map with zero values:", zeroValueMap)

	// maps - iterating over a map in a specific order
	sortedKeys := []string{"apple", "banana", "orange", "grape"}
	for _, key := range sortedKeys {
		if value, exists := mapExample[key]; exists {
			fmt.Printf("Fruit: %s, Quantity: %d\n", key, value)
		} else {
			fmt.Printf("Fruit: %s not found in the map\n", key)
		}
	}
}
