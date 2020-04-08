package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const (
	fileName = "data.json"
)

func main() {
	// Start processing the data and output result
	processData()
}

// The algorithm
func processData() {
	// 1. Read data from file
	recipients := readJSON(fileName)

	// 2. Populate map for recipients (Dictionary/HashSet in C#)
	generateAllRecipientsTagsMap(&recipients)

	// 3. Get the recipients pair with at least 2 overlapping tags
	overlappingRecipients := getTwoOrMoreOverlappingTagsRecipients(recipients)

	// 4. Get the intended output string
	formattedNames := getFormattedNameOfSimilarRecipients(overlappingRecipients)

	// 5. Print the result out
	fmt.Println(formattedNames)
}

// Populate recipient.tagsMap for the input list of recipients
func generateAllRecipientsTagsMap(recipients *[]recipient) {
	for r := range *recipients {
		(*recipients)[r].generateTagsMap()
	}
}

// Input a list of recipients and returns a list of recipients pairs with at least 2 overlapping tags
func getTwoOrMoreOverlappingTagsRecipients(recipients []recipient) (overlappingRecipients [][]recipient) {
	// Traverse through each recipient
	for k := range recipients {
		// Traverse through each recipient + 1
		for j := range recipients[k+1:] {
			// Find similar tag
			if recipients[k].hasTwoOrMoreOverlappingTags(recipients[j+k+1]) {
				// Append to the pair if they have at least 2 overlapping tags
				overlappingRecipients = append(overlappingRecipients, []recipient{recipients[k], recipients[j+k+1]})
			}
		}
	}

	return
}

// From the argument input of array, generate the intended name.
// Ex: "Name2, Name1|Name1, Name3"
func getFormattedNameOfSimilarRecipients(overlappingRecipients [][]recipient) string {
	// Use a string builder
	var sb strings.Builder

	// Iterate through all the pairs of names
	for k, v := range overlappingRecipients {
		// Separate pairs with '|'
		if k != 0 {
			sb.WriteByte('|')
		}

		// Iterate through each recipient in the pair
		for j, r := range v {
			// Separate recipients name with ", "
			if j != 0 {
				sb.WriteString(", ")
			}

			sb.WriteString(r.Name)
		}

		// Remove the extra ", " at the end with backspace's escape notation
		// sb.WriteString("\b\b")
	}

	// Remove the extra "|" with backspace's escape notation
	// '\b' as the last character doesn't work
	// sb.WriteString("\b   ")

	return sb.String()
}

// Read JSON data from the file
func readJSON(fileName string) []recipient {
	// Open the file
	file, err := os.Open(fileName)
	if err != nil {
		// Panic and quit the application if error
		panic(err)
	}

	// Create rw Reader and read all from the file
	reader := bufio.NewReader(file)
	jsonData, err := ioutil.ReadAll(reader)
	if err != nil {
		panic(err)
	}

	// Unmarshal the data to rw Recipient wrapper
	var rw recipientWrapper
	err = json.Unmarshal(jsonData, &rw)
	if err != nil {
		panic(err)
	}

	// Return the array/slice of recipients (Not the wrapper)
	return rw.JsonArrayData
}
