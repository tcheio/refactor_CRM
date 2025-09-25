package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lister les contacts",
	RunE: func(cmd *cobra.Command, args []string) error {
		contacts, err := Store.GetAll()
		if err != nil {
			return err
		}
		if len(contacts) == 0 {
			fmt.Println("Aucun contact.")
			return nil
		}
		fmt.Println("📒 Contacts :")
		for _, c := range contacts {
			fmt.Printf("- ID:%d | Nom:%s | Email:%s\n", c.ID, c.Name, c.Email)
		}
		return nil
	},
}
