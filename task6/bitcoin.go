package main 

import ( 
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
) 

func main() { 

	// os.Open() opens specific file in 
	// read-only mode and this return 
	// a pointer of type os. 
	file, err := os.Open("test.txt") 

	if err != nil { 
		log.Fatalf("failed to open") 

	} 

	// The bufio.NewScanner() function is called in which the 
	// object os.File passed as its parameter and this returns a 
	// object bufio.Scanner which is further used on the 
	// bufio.Scanner.Split() method. 
	scanner := bufio.NewScanner(file) 

	// The bufio.ScanLines is used as an 
	// input to the method bufio.Scanner.Split() 
	// and then the scanning forwards to each 
	// new line using the bufio.Scanner.Scan() 
	// method. 
	scanner.Split(bufio.ScanLines) 
	var text []string 

	for scanner.Scan() { 
		text = append(text, scanner.Text()) 
	} 

	// The method os.File.Close() is called 
	// on the os.File object to close the file 
	file.Close() 
	languageCode := "en"
	searchTerm := "bitcoin"
	occurence := 0
	// and then a loop iterates through 

	for _, each_ln := range text { 
		if string(each_ln[0]) + string(each_ln[1]) == languageCode && strings.Contains(string(each_ln), searchTerm) {
			count := 0
			for i := 0; i < len(each_ln); i++ {
				if string(each_ln[i]) == " " {
					count++
				}
				if count == 2 {
					occurence += int(each_ln[i+1]) - 48
					break
				}
			}
		}
	}
	
	fmt.Println("The word", searchTerm, "is searched for", occurence,"times") 
}
