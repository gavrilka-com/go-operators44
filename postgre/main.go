package main

import (
	"bitbucket.org/bosyak/go-operators44/posgre/posgre"
	"fmt"
	"os"
	"os/user"
	"path"
)

func main() {
	fmt.Println("Hello Postgre!!")

	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Errorf("error: %s", err)
	}
	pgDataFolder := path.Join(homeDir, "db", "pgdata")

	currentUser, err := user.Current()
	if err != nil {
		fmt.Errorf("error: %s", err)
		return
	}

	pcs := posgre.PgContainerStart{
		Name:           "pgsql",
		PortMapping:    "5432:5432",
		Image:          "postgres",
		Password:       "123",
		DbFolder:       pgDataFolder,
		DbFolderBackup: path.Join(homeDir, "db", "pgdata_backup"),
		CurrentUser:    *currentUser,
	}

	err = posgre.StartPgContainer(pcs)
	if err != nil {
		fmt.Errorf("error: %s", err)
	}
}
