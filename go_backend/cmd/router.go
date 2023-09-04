
/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"go_project/package/serialize"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

// testCmd represents the test command
var runserverCmd = &cobra.Command{
	Use:   "runserver",
	Short: "run server",
	Long: `run web server router`,
	Run: func(cmd *cobra.Command, args []string) {
	fmt.Println("run server start")
    router := gin.Default()
	serialize.SetupDatabase() 

    router.POST("/user", serialize.CreateUser)
	router.GET("/Getall", serialize.Getall)
	router.GET("/CreateUser/:name/:age", serialize.CreateUser)
	
	router.GET("/GetUser/:name", serialize.GetUser)
	router.GET("/DeleteUser/:name", serialize.DeleteUser)

	router.GET("/UpdateUser/:name/:age", serialize.UpdateUser)

	router.GET("/Insert_server_online", serialize.Insert_server_online)
	router.POST("/Insert_server_online", serialize.Insert_server_online)
	router.POST("/Insert_ve_monitor_data", serialize.Insert_ve_monitor_data)

	
	
    // router.GET("/user", serialize.CreateUser)

	router.Run(":8089")
   


	},
}

func init() {
	rootCmd.AddCommand(runserverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// testCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// testCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
