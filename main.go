/*
 Program Name: Lab5
 Student Names: Duncan Scott Martinson & John Payton
 Semester: Spring 2017
 Class: COSC 30403
 Instructor: Dr. James Comer
 Due Date: 4/11/2017

 Program Overview:
   This GO program uses a text file to input data into a linked list of
   Employees and execute commands in order to manipulate the list. After every
   command, the program writes information about the execution to a separate
   text file.
   After all the data from the file has been read, the program terminates.

 Input:
   The program requires formatted data in a text file entitled "Lab5Data.txt"

 Output:
   The program outputs a data file entitled "Lab5Ans.txt" containing information
   about the program's execution.

 Program Limitations:
   1) The program does not allow for real time user interaction, and the output
   file is overwritten after every execution.


 Significant Program Variables:
   input - the identifier for the input file
   output - the identifier for the output file
   reader - the Scanner that reads data from input
   writer - the Writer that writes data to output
   elements - a string array that stores all of the words in a line from input
   in each element
   cmd - a string that stores the first two characters read from a line of
   input. This is used to determine the appropriate function to execute.
   head- an Employee node that serves as the list head. Used to reference the
   linked list.
   current- an Employee node pointer used in almost every procedure. Is used to
   traverse the list and analyze data, manipulating it if need be.
*/

package main

//Calling import signals to the program to use the specified packages in this
//program
import (
	"bufio"
	"os"
	"strings"
)

/*Defining the Employee data type. In this lab, we represented id and pay as
strings instead of numeric values to preserve the leading zeros for formatting
reasons. Go'scomparison operators (==, <, >) are overloaded to compare strings,
so this is possible.*/
type Employee struct {
	id, name, department, title, pay string
	next                             *Employee
}

//Header node
var head = Employee{"0", "", "", "", "", nil}

//io setup
var output, _ = os.Create("Lab5Ans.txt")
var writer = bufio.NewWriter(output)
var input, _ = os.Open("Lab5Data.txt")
var reader = bufio.NewScanner(bufio.NewReader(input))

//The main function
func main() {
	//Set up the reader to read line by line
	reader.Split(bufio.ScanLines)
	writer.WriteString("START PROGRAM\n")

	//Close the files when all the other functions stop executing (at the end of the program)
	defer input.Close()
	defer output.Close()

	//Check for end of file
	for reader.Scan() {
		//Read in a line, put each word in the line into an element of a string array
		inStr := reader.Text()
		elements := strings.Fields(inStr)
		cmd := elements[0]
		//Check the command and execute the aproppriate function
		switch cmd {
		case "IN":
			insert(Employee{elements[1], addSpace(elements[2], 11), addSpace(elements[3], 15), addSpace(elements[4], 15), elements[5], nil})
		case "DE":
			delete(elements[1])
		case "PA":
			printAll()
		case "PI":
			printID(elements[1])
		case "PD":
			printDept(elements[1])
		case "UN":
			uName(elements[1], addSpace(elements[4], 11))
		case "UD":
			uDept(elements[1], addSpace(elements[4], 15))
		case "UT":
			uTitle(elements[1], addSpace(elements[4], 15))
		case "UR":
			uRate(elements[1], elements[4])
		}
	}
	writer.WriteString("END PROGRAM")
	writer.Flush()
}

/*The insert function takes an Employee and inserts it into the linked list.
It uses a pointer to the header node called "current" to access information from
the list.It looks down the list one when finding the correct spot to insert in order to
conserve memory rather than using a doubly linked list. First, it checks to see if current is
at the end of the list. If it is, then it inserts at the end of the list. Otherwise, it compares the
id's of the current node and the new one, and inserts if the new id is a lower number. Otherwise,
it goes to the next node and repeats. */
func insert(newEmp Employee) {
	//Point to the head
	current := &head
	for {
		//Check to see if current is at the last node
		if current.next == nil {
			//insert
			current.next = &newEmp
			writer.WriteString("IN: Employee #" + newEmp.id + " inserted.\n")
			break
		} else if current.next.id > newEmp.id { //Check the id's
			//insert
			newEmp.next = current.next
			current.next = &newEmp
			writer.WriteString("IN: Employee #" + newEmp.id + " inserted.\n")
			break
		} else {
			//Go to the next node
			current = current.next
		}
	}
}

/*The delete function deletes the node with the id matching "target" from
the list. It uses the same traversal as insert, looking one down the line to
make it easier to connect around the deleted node. Since Go has garbage collection,
the node is left to be collected. It uses a pointer to the head node called "current"
to access data from the nodes. If current is at the end of the list, then we didn't
find the specified id, so we notify the user and exit. Otherwise, if the id of current.next
matches target, we delete it from the list and exit. Otherwise, we go to the next node and
repeat*/
func delete(target string) {
	//Point to current
	current := &head
	for {
		//Check if current is at the end, if so, break out of the loop
		if current.next == nil {
			break
		} else if current.next.id == target { //Check if the id matches target
			//Delete, notify, and return
			current.next = current.next.next
			writer.WriteString("DE: Employee #" + target + " deleted.\n")
			return
		} else {
			//Go to the next node
			current = current.next
		}
	}
	//Error message only happens if we don't find the target because if we do,
	//we return instead of breaking.
	writer.WriteString("DE ERROR: Employee #" + target + " not found.\n")
}

/*The printAll function traverses the list and prints out the information of each
node. It uses a pointer to the head node "current" to access data from the nodes.
If current is null, we're at the end of the list, and we break out of the loop.
Otherwise, we print the information out of current and then go to the next node, repeating
until we hit the end.*/
func printAll() {
	//Point to the first employee
	current := head.next
	writer.WriteString("PA: BEGIN PRINT ALL\n")
	for {
		//Check to see if were at the end, breaking if we are
		if current == nil {
			break
		} else {
			//Print out the info and iterate
			printEmp(*current)
			current = current.next
		}
	}
	writer.WriteString("END PRINT ALL\n")
}

//The printEmp function is a helper function that prints out all of the information
//from the given employee
func printEmp(emp Employee) {
	writer.WriteString(emp.id + " " + emp.name + " " + emp.department + " " + emp.title + " " + emp.pay + "\n")
}

/*The printID function looks for the employee with the id that matches "target" and
prints out all of its information. It uses a pointer to the head node "current" to access
the information in the nodes. It checks to see if current is at the end of the list
and exits the loop if it is, notifying the user before the function terminates. Otherwise
it checks to see if current's id matches target, printing it out and returning
if it does and iterating if it doesnt*/
func printID(target string) {
	//Point to the head
	current := head.next
	for {
		//Are we at the end
		if current == nil {
			break
		} else if current.id == target { //Is it the right id
			//Print out the employee
			writer.WriteString("PI: ")
			printEmp(*current)
			return
		} else {
			//Go to next node
			current = current.next
		}
	}
	//Error message only writes if we didn't find the employee
	writer.WriteString("PI ERROR: Employee #" + target + " not found.\n")
}

/*The printDept function works similarly to the printAll function, except it only
prints out the employees from a specified department.*/
func printDept(dept string) {
	//This boolean trips if we find at least one employee in the department.
	//Used for the error message
	found := false
	current := head.next
	writer.WriteString("PD: BEGIN PRINT DEPARTMENT:\n")
	for {
		if current == nil {
			break
		} else if strings.TrimSpace(current.department) == dept {
			//Trip the boolean and print the employee
			found = true
			printEmp(*current)
			current = current.next
		} else {
			current = current.next
		}
	}
	//Check to see if we didn't find an employee
	if !found {
		//Error message
		writer.WriteString("PD ERROR: " + strings.TrimSpace(dept) + " department not found.\n")
	}
	writer.WriteString("END PRINT DEPARTMENT.\n")
}

/*The following functions uName, uDept, uTitle, and uRate all work almost exactly the same
way except that they update different values of an employee. It takes the id of the target
employee and the new information and uses a pointer to the head node* "current" to access the data
from the nodes. Each function checks if current is at the end of the list. If it is, it means
we did not find the target id, so we break out of the loop and notify the user before the function
terminates. Otherwise, if the target id matches the id of current. If so, we update the appropriate information,
notify the user, and return. If not, we iterate down the list and repeat*/
func uName(target, newName string) {
	var tempName string
	//Point to first node
	current := head.next
	for {
		//Check if we're at the end
		if current == nil {
			break
		} else if current.id == target { //Check if it's the right node
			//Update and notify
			tempName = current.name
			current.name = newName
			writer.WriteString("UN: Employee #" + target + " name changed from " + strings.TrimSpace(tempName) + " to " + strings.TrimSpace(newName) + ".\n")
			return
		} else {
			//Next node
			current = current.next
		}
	}
	//Error message
	writer.WriteString("UN ERROR: Employee #" + target + " not found.\n")
}

func uDept(target, newDept string) {
	var tempDept string
	current := &head
	for {
		if current == nil {
			break
		} else if current.id == target {
			tempDept = current.department
			current.department = newDept
			writer.WriteString("UD: Employee #" + target + " department changed from " + strings.TrimSpace(tempDept) + " to " + strings.TrimSpace(newDept) + ".\n")
			return
		} else {
			current = current.next
		}
	}
	writer.WriteString("UD ERROR: Employee #" + target + " not found.\n")
}

func uTitle(target, newTitle string) {
	var tempTitle string
	current := &head
	for {
		if current == nil {
			break
		} else if current.id == target {
			tempTitle = current.title
			current.title = newTitle
			writer.WriteString("UT: Employee #" + target + " title changed from " + strings.TrimSpace(tempTitle) + " to " + strings.TrimSpace(newTitle) + ".\n")
			return
		} else {
			current = current.next
		}
	}
	writer.WriteString("UT ERROR: Employee #" + target + " not found.\n")
}

func uRate(target, newPay string) {
	var tempPay string
	current := &head
	for {
		if current == nil {
			break
		} else if current.id == target {
			tempPay = current.pay
			current.pay = newPay
			writer.WriteString("UR: Employee #" + target + " payrate changed from " + strings.TrimSpace(tempPay) + " to " + strings.TrimSpace(newPay) + ".\n")
			return
		} else {
			current = current.next
		}
	}
	writer.WriteString("UR ERROR: Employee #" + target + " not found.\n")
}

/*The addSpace function is a helper function that adds spaces to a string str to make it
num characters long. It's used in formatting the output so that the data is aligned and
easy to read*/
func addSpace(str string, num int) string {
	for i := len(str); i < num; i++ {
		str = str + " "
	}
	return str
}
