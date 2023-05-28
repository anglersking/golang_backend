/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

    

	"github.com/spf13/cobra"
	"go_project/package/serialize"

 
  
    "github.com/google/uuid"
)
type User struct {
    ID   uuid.UUID `gorm:"type:UUID;primaryKey"`
    Name string `gorm:"column:name;not null"`
    Age  int64    `gorm:"column:age;not null"`
}


// testCmd represents the test command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		serialize.SetupDatabase() 
		serialize.Db.AutoMigrate(&User{})

		fmt.Println("database migrate succuess")


	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// testCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// testCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
