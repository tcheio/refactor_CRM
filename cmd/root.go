package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"refactor_crm_interface/internal/storage"
)

var (
	cfgFile string
	Store   storage.Storer
)

var rootCmd = &cobra.Command{
	Use:   "mini-crm",
	Short: "Mini-CRM — gestion de contacts en CLI",
	Long:  "Un gestionnaire de contacts simple et efficace, avec stockage interchangeable (memory/json/gorm).",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return initStore()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	// config flag
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "fichier de configuration (ex: ./config.yaml)")
	cobra.OnInitialize(initConfig)

	// sous-commandes
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(updateCmd)
	rootCmd.AddCommand(deleteCmd)
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigType("yaml")
		viper.SetConfigName("config")
		viper.AddConfigPath(".")
	}
	// env: MINI_CRM_STORAGE_TYPE, MINI_CRM_STORAGE_DB_PATH, ...
	viper.SetEnvPrefix("MINI_CRM")
	viper.AutomaticEnv()

	_ = viper.ReadInConfig() // si absent, on garde les valeurs par défaut
}

func initStore() error {
	// valeurs par défaut
	viper.SetDefault("storage.type", "memory")
	viper.SetDefault("storage.json_path", "data/contacts.json")
	viper.SetDefault("storage.db_path", "data/contacts.db")

	typ := viper.GetString("storage.type")
	jsonPath := viper.GetString("storage.json_path")
	dbPath := viper.GetString("storage.db_path")

	switch typ {
	case "memory":
		Store = storage.NewMemoryStore()
	case "json":
		_ = os.MkdirAll(filepath.Dir(jsonPath), 0o755)
		s, err := storage.NewJSONFileStore(jsonPath)
		if err != nil {
			return fmt.Errorf("erreur init JSON store: %w", err)
		}
		Store = s
	case "gorm":
		_ = os.MkdirAll(filepath.Dir(dbPath), 0o755)
		s, err := storage.NewGORMStore(dbPath)
		if err != nil {
			return fmt.Errorf("erreur init GORM store: %w", err)
		}
		Store = s
	default:
		return fmt.Errorf("storage.type inconnu: %s (attendu: memory|json|gorm)", typ)
	}
	return nil
}
