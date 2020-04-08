package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	recipients := readJSON()

	generateAllRecipientsTagsMap(&recipients)

	similarRecipients := getSimilarRecipients(recipients)

	formattedNames := getFormattedNameOfSimilarRecipients(similarRecipients)

	fmt.Println(formattedNames)
}

func generateAllRecipientsTagsMap(recipients *[]recipient) {
	for r := range *recipients {
		(*recipients)[r].generateTagsMap()
	}
}

func getSimilarRecipients(recipients []recipient) (similarRecipients [][]recipient) {
	for k := range recipients {
		for j := range recipients[k+1:] {
			if recipients[k].hasTwoOrMoreSimilarTags(recipients[j+k+1]) {
				similarRecipients = append(similarRecipients, []recipient{recipients[k], recipients[j+k+1]})
			}
		}
	}

	return
}

func getFormattedNameOfSimilarRecipients(similarRecipients [][]recipient) string {
	var sb strings.Builder

	for _, v := range similarRecipients {
		for _, r := range v {
			sb.WriteString(r.Name)
			sb.WriteString(", ")
		}
		sb.WriteString("\b\b")
		sb.WriteByte('|')
	}
	sb.WriteString("\b")

	return sb.String()
}

func readJSON() []recipient {
	file, err := os.Open("data.json")
	if err != nil {
		fmt.Println(err)
	}
	reader := bufio.NewReader(file)
	jsonStr, err := ioutil.ReadAll(reader)
	if err != nil {
		fmt.Println(err)
	}
	var a recipientWrapper
	err = json.Unmarshal(jsonStr, &a)
	if err != nil {
		fmt.Println(err)
	}

	return a.JsonArrayData
}
