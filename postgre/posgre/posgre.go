package posgre

import (
	"bitbucket.org/bosyak/go-operators44/runn/runn"
	"fmt"
	"os/user"
	"strings"
)

type PgContainerStart struct {
	Name           string
	PortMapping    string
	Image          string
	Password       string
	DbFolder       string
	DbFolderBackup string
	CurrentUser    user.User
}

//inline func for PgContainerStart

func (pgc *PgContainerStart) ShellCmd() string {
	var mountVolumes []string
	if pgc.DbFolder != "" {
		mountVolumes = append(mountVolumes, fmt.Sprintf("--volume %s:/var/lib/postgresql/data", pgc.DbFolder))
	}
	if pgc.DbFolderBackup != "" {
		mountVolumes = append(mountVolumes, fmt.Sprintf("--volume %s:/var/lib/postgresql/data_backup", pgc.DbFolderBackup))
	}

	user := "--user " + pgc.CurrentUser.Uid + ":" + pgc.CurrentUser.Gid
	user = ""

	return fmt.Sprintf("docker run --rm %s --name %s -e POSTGRES_PASSWORD=%s %s -p %s -d %s",
		user, pgc.Name, pgc.Password, strings.Join(mountVolumes, " "), pgc.PortMapping, pgc.Image)
}

func StartPgContainer(pcs PgContainerStart) error {

	out, err := runn.RunCommandVerbose(pcs.ShellCmd())
	if err != nil {
		return err
	}
	fmt.Println(out)
	return nil
}
