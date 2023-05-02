package main

import (
	"ECE9393/tools"
	//"bufio"
	//"context"
	"fmt"
	//"os"

	//"strings"
	"time"

	openai "github.com/sashabaranov/go-openai"
)

func main() {

	client := openai.NewClient("OPENAI'S API KEY")
	// 0.002USD / 1k tokens


	/*this part gives initial results*/
	file, err := os.Open("pp_google play.txt") //open file, presumably contains 1000 lines
	if err != nil {
		fmt.Println("error opening file", err)
		return
	}

	defer file.Close()

	fiftyLines := make([]string, 10) //let gpt analyze 5 lines each time, this is an array

	scanner := bufio.NewScanner(file)

	for batch := 0; batch < 20; batch++ {

		fiftyLines = make([]string, 10) //re-initialize array

		for i := 0; i < 5 && scanner.Scan(); i++ { //read file 5 URLs each time
			URL := scanner.Text()
			fiftyLines = append(fiftyLines, URL)
			fmt.Println(URL)
		}

		inOneLine := strings.Join(fiftyLines, "\n") //5 lines in one string

		//make a request
		resp, err := client.CreateChatCompletion(
			context.Background(),
			openai.ChatCompletionRequest{
				Model: openai.GPT3Dot5Turbo,
				Messages: []openai.ChatCompletionMessage{
					{
						Role: openai.ChatMessageRoleUser,

						/*HIPPA √ */
						// Content: "please summarize privacy policies of 5 apps here: " + inOneLine +
						// 	"\ntell me how many of them violates any HIPPA clauses?",

						/*how many √ */
						// Content: "please summarize privacy policies of 5 apps here: " + inOneLine +
						// 	"\ntell me how many of them mentioned sharing users' data with third parties?",

						/*sharing with third parties, for what √ */
						Content: "please summarize privacy policies of 5 apps here: " + inOneLine +
							"\nif any of them claims sharing user data with third parties, " +
							"please summarize the purpose of doing so, and what kind of third parties" +
							"they are sending to",

						/*how long √ */
						// Content: `please summarize privacy policies of 5 apps here: ` + inOneLine +
						// 	`If it mentions the collection of user data, ` +
						// 	`how long will they retain the data?`,

						/*when to collect what kind of data √ */
						// Content: `please summarize privacy policies of several apps here: ` + inOneLine +
						// 	`If it mentions the collection of user data, ` +
						// 	`under what circumstances will they do so, ` +
						// 	`and what types of data will be collected?`,
						/*high refuse rate: when collect*/
						// Content: `summarize privacy policies here: ` + inOneLine +
						// 	`If it mentions the collection of user data, ` +
						// 	`under what circumstances will data collection be triggered ` +
						// 	`and what types of data will be collected?`,
						/*refuse to respond: when collect*/
						// `according to privacy policies here: ` + inOneLine +
						// `If it mentions the collection of user data, ` +
						// `under what circumstances will data collection be triggered ` +
						// `and what types of data will be collected?`,

					},
				},
			},
		)
		if err != nil {
			fmt.Printf("error occured in %d iterations", batch)
			panic(err)
		}

		fmt.Println(resp.Choices[0].Message.Content)

		tools.WriteFile(`output`+fmt.Sprintf("%d", batch)+".txt",
			resp.Choices[0].Message.Content, batch)

		fmt.Println("waiting...")
		time.Sleep(time.Second * 20)
	}
	/*call this to tally initial results*/
	count(client)
}

func count(client *openai.Client) { //analyze results
	counts := make([]int, 10)

	//count each files
	for i := 0; i <= 19; i++ {
		content := tools.ReadWholeFile("Goutput" + fmt.Sprintf("%d", i) + ".txt")
		tools.CountNumbers(counts, client, content)
		fmt.Printf("counts[0] = %d, counts[1] =  %d, counts[2] = %d,"+
			"counts[3] = %d, counts[4] = %d, counts[5] = %d, counts[6] = %d, "+
			"counts[7] = %d, counts[8] = %d, counts[9] = %d \nwaiting...\n", counts[0], counts[1], counts[2],
			counts[3], counts[4], counts[5], counts[6], counts[7], counts[8], counts[9])
		time.Sleep(time.Second * 20)
	}
	/*finally dump results of 20 files to a single file*/
	wholeContent := fmt.Sprintf("improve services: %d, personalize content: %d, conduct promotions: %d,"+
		"advertising: %d, payment processing: %d, customer support: %d, other: %d\nwaiting...\n", counts[0], counts[1], counts[2], counts[3], counts[4],
		counts[5], counts[6])
	tools.WriteFile("counts.txt", wholeContent, 0)

}
