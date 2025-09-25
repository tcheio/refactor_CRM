package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var delID int

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Supprimer un contact",
	RunE: func(cmd *cobra.Command, args []string) error {
		if delID <= 0 {
			return fmt.Errorf("veuillez fournir --id")
		}
		if err := Store.Delete(delID); err != nil {
			return err
		}
		fmt.Println("🗑️ Contact supprimé.")
		return nil
	},
}

func init() {
	deleteCmd.Flags().IntVar(&delID, "id", 0, "ID du contact (obligatoire)")
	_ = deleteCmd.MarkFlagRequired("id")
}
