package tools

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	openai "github.com/sashabaranov/go-openai"
)

func CountNumbers(counts []int, client *openai.Client, txt string) {
	//txt is analysiz result from 5 apps
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role: openai.ChatMessageRoleUser,
					/*how many: third parties and HIPPA √ */
					// Content: "\nthis is information about 5 apps: \n" + txt +
					// 	"tell me how many apps were mentioned " +
					// 	//"may share data with third parties " +
					// 	"may violate HIPPA clauses." +
					// 	"return the result to me. please strictly follow the format:\n" +
					// 	//"Sharing data to third parties: \n" +
					// 	"Violating HIPPA clauses: \n" +
					// 	"if you can't determine based on the information, just take it as 0" +
					// 	"NOTHING ELSE IS NEEDED, DO NOT ADD a single extra character\n",

					/*for what reason √ */
					Content: "I will give you information about 5 apps: " + txt +
						"which may mention sharing user data with third parties for certain purposes. " +
						"I need you to tally the number of each purpose is mentioned. " +
						"give me only numbers, your response should only be a single list like this: " +
						"\nimprove services:" +
						"\npersonalize content:" +
						"\nconduct promotions:" +
						"\nadvertising:" +
						"\npayment processing:" +
						"\ncustomer support:" +
						"\nother:" +
						//"\nif you can't determine based on the information, just take it as other" +
						"no more expanations needed in any field, only give me numbers.\n",

					/*collect what √ */
					// Content: "I will give you information about 5 apps: " + txt +
					// 	"which may mention collecting specific kinds of user data" +
					// 	"I need you to tally the number of each kind of information mentioned above. " +
					// 	"give me only numbers, your response should only be a single list like this:" +
					// 	"\nlocations:" +
					// 	"\ndevice information:" +
					// 	"\nsleeping habits:" +
					// 	"\nuser activity:" +
					// 	"\nemail:" +
					// 	"\nname:" +
					// 	"\nage:" +
					// 	"\nbiometric:" +
					// 	"\npayment information:" +
					// 	"\nother: " +
					// 	// "\nif you can't determine based on the information, just take it as other" +
					// 	"no more expanations needed in any field, only give me numbers",
				},
			},
		},
	)
	response := resp.Choices[0].Message.Content
	fmt.Println(response)

	//parse counts from responses
	curCounts := make([]int, 10)
	lines := strings.Split(response, "\n")
	for _, line := range lines {
		if strings.Contains(line, ":") {
			parts := strings.Split(line, ":")
			if strings.Contains(parts[0], "improve") {
				curCounts[0], err = strconv.Atoi(strings.TrimSpace(parts[1]))
			}
			if strings.Contains(parts[0], "personalize") {
				curCounts[1], err = strconv.Atoi(strings.TrimSpace(parts[1]))
			}
			if strings.Contains(parts[0], "promotions") {
				curCounts[2], err = strconv.Atoi(strings.TrimSpace(parts[1]))
			}
			if strings.Contains(parts[0], "advertising") {
				curCounts[3], err = strconv.Atoi(strings.TrimSpace(parts[1]))
			}
			if strings.Contains(parts[0], "payment") {
				curCounts[4], err = strconv.Atoi(strings.TrimSpace(parts[1]))
			}
			if strings.Contains(parts[0], "support") {
				curCounts[5], err = strconv.Atoi(strings.TrimSpace(parts[1]))
			}
			if strings.Contains(parts[0], "other") {
				curCounts[6], err = strconv.Atoi(strings.TrimSpace(parts[1]))
			}
			// if strings.Contains(parts[0], "biometric") {
			// 	curCounts[7], err = strconv.Atoi(strings.TrimSpace(parts[1]))
			// }
			// if strings.Contains(parts[0], "payment") {
			// 	curCounts[8], err = strconv.Atoi(strings.TrimSpace(parts[1]))
			// }
			// if strings.Contains(parts[0], "other") {
			// 	curCounts[9], err = strconv.Atoi(strings.TrimSpace(parts[1]))
			// }
			if err != nil {
				panic(err)
			}
		}
	}
	counts[0] += curCounts[0]
	counts[1] += curCounts[1]
	counts[2] += curCounts[2]
	counts[3] += curCounts[3]
	counts[4] += curCounts[4]
	counts[5] += curCounts[5]
	counts[6] += curCounts[6]
	counts[7] += curCounts[7]
	counts[8] += curCounts[8]
	counts[9] += curCounts[9]
}
