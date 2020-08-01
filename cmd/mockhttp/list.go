package main

import (
	"fmt"
	"os"

	"github.com/nicklarsennz/mockhttp/responders"
	"github.com/olekukonko/tablewriter"
)

func listResponderMappings(config *responders.ResponderConfig) error {

	table := tablewriter.NewWriter(os.Stdout)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetAutoMergeCells(true)
	table.SetAutoWrapText(false)
	table.SetHeader([]string{"Responder", "Request", "Response"})

	for i, responder := range config.Responders {
		when := responder.When
		then := responder.Then

		table.Append([]string{
			fmt.Sprintf("%d", i),
			fmt.Sprintf("%s %s", when.Http.Method, when.Http.Path),
			fmt.Sprintf("%d %s", then.Http.Status, then.Http.Message),
		})

		table.Append([]string{fmt.Sprintf("%d", i), "", ""})

		var whenHeaders string
		for k, v := range when.Headers {
			whenHeaders = fmt.Sprintf("%s%s: %s\n", whenHeaders, k, v)
		}

		var thenHeaders string
		for k, v := range then.Headers {
			thenHeaders = fmt.Sprintf("%s%s: %s\n", thenHeaders, k, v)
		}

		table.Append([]string{
			fmt.Sprintf("%d", i),
			fmt.Sprintf("%s", whenHeaders),
			fmt.Sprintf("%s", thenHeaders),
		})

		table.Append([]string{
			fmt.Sprintf("%d", i),
			fmt.Sprintf("%s", when.Body),
			fmt.Sprintf("%s", then.Body),
		})

		table.Append([]string{fmt.Sprintf("%d", i), "", ""})
	}

	table.Render()
	return nil
}
