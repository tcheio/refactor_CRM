package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	updID    int
	updName  string
	updEmail string
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Mettre à jour un contact",
	RunE: func(cmd *cobra.Command, args []string) error {
		if updID <= 0 {
			return fmt.Errorf("veuillez fournir --id")
		}
		if err := Store.Update(updID, updName, updEmail); err != nil {
			return err
		}
		fmt.Println("✅ Contact mis à jour.")
		return nil
	},
}

func init() {
	updateCmd.Flags().IntVar(&updID, "id", 0, "ID du contact (obligatoire)")
	updateCmd.Flags().StringVar(&updName, "name", "", "Nouveau nom (optionnel)")
	updateCmd.Flags().StringVar(&updEmail, "email", "", "Nouvel email (optionnel)")
	_ = updateCmd.MarkFlagRequired("id")
}
