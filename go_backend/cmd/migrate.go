/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"go_project/package/serialize"

 
  
    // "github.com/google/uuid"
)
// type User struct {
//     ID   uuid.UUID `gorm:"type:UUID;primaryKey"`
//     Name string `gorm:"column:name;not null"`
//     Age  int64    `gorm:"column:age;not null"`
// }
type Server_online struct {
    Time   time.Time `gorm:"type:datetime;not null"`
    Data string `gorm:"column:data;not null"`
    // Age  int64    `gorm:"column:age;not null"`
}


type Ve_monitor struct {
    Time             time.Time `gorm:"type:datetime;not null"`
    DA_status        string    `gorm:"column:da_status;not null"`
    Data_reserved    string    `gorm:"column:data_reserved;not null"`
    IPC_status       string    `gorm:"column:ipc_status"`
	
	Plate_number     string    `gorm:"column:plate_number"`
    Source           string    `gorm:"column:source"`

    Type             string     `gorm:"column:type"`
	
    Schema_version   string      `gorm:"column:schema_version"`

    Vehicle          string     `gorm:"column:vehicle"`
    Fdi_type         string      `gorm:"column:fdi_type"`
	Timestamp        time.Time       `gorm:"column:timestamp"`

    SH_status        string    `gorm:"column:sh_status"`
    SW_version       string    `gorm:"column:sw_version"`
    Vehicletype      string    `gorm:"column:vehicletype"`
    Cpe_status       bool      `gorm:"column:cpe_status"`
    Isp_lock         string    `gorm:"column:isp_lock"`
    Isp_vsync        string    `gorm:"column:isp_vsync"`
    Lidar_check      string    `gorm:"column:lidar_check"`
    Monitor_pcie     string    `gorm:"column:monitor_pcie"`
    Net_interface    string    `gorm:"column:net_interface"`
    Net_status       string    `gorm:"column:net_status"`
    Trunk_temperature float64      `gorm:"column:trunk_temperature"`
    Vechicles_status string    `gorm:"column:vechicles_status"`
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
		serialize.Db.AutoMigrate(&Server_online{})
		serialize.Db.AutoMigrate(&Ve_monitor{})

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
