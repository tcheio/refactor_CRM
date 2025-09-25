package cmd

import (
	"fmt"

	"refactor_crm_interface/internal/storage"

	"github.com/spf13/cobra"
)

var (
	addName  string
	addEmail string
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Ajouter un contact",
	RunE: func(cmd *cobra.Command, args []string) error {
		if addName == "" || addEmail == "" {
			return fmt.Errorf("veuillez fournir --name et --email")
		}
		c := &storage.Contact{Name: addName, Email: addEmail}
		if err := Store.Add(c); err != nil {
			return err
		}
		fmt.Printf("✅ Contact ajouté: ID=%d, Nom=%s, Email=%s\n", c.ID, c.Name, c.Email)
		return nil
	},
}

func init() {
	addCmd.Flags().StringVar(&addName, "name", "", "Nom du contact (obligatoire)")
	addCmd.Flags().StringVar(&addEmail, "email", "", "Email du contact (obligatoire)")
	_ = addCmd.MarkFlagRequired("name")
	_ = addCmd.MarkFlagRequired("email")
}
