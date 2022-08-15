package cmd

import (
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	"github.com/JhndaCoder/CLI-Todo/todo"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the todos",
	Long:  `Lisitng the todos`,
	Run:   listRun,
}

func listRun(cmd *cobra.Command, args []string) {
	items, err :=
		todo.ReadItems(dataFile)
	if err != nil {
		log.Printf("%v", err)
	}
	fmt.Println(items)

	w := tabwriter.NewWriter(os.Stdout, 3, 0, 1, ' ', 0)
	for _, i := range items {
		fmt.Fprintln(w, i.Prettyp()+"\t"+i.Text+"\t")
	}
	w.Flush()
}

func init() {
	rootCmd.AddCommand(listCmd)
}
