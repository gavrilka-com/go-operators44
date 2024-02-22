package posgre

import (
	"github.com/stretchr/testify/assert"
	"os/user"
	"testing"
)

func TestPgContainerStart_ShellCmd(t *testing.T) {
	pcs := PgContainerStart{
		Name:           "test",
		PortMapping:    "5432:5432",
		Image:          "postgres",
		Password:       "p@ssw0rd",
		DbFolder:       "/home/user/database/pgdata",
		DbFolderBackup: "/home/user/database/pgdata_backup",
		CurrentUser:    user.User{Uid: "1000", Gid: "1000"},
	}

	assert.Equal(t,
		"docker run --rm --user 1000:1000 --name test -e POSTGRES_PASSWORD=p@ssw0rd --volume /home/user/database/pgdata:/var/lib/postgresql/data --volume /home/user/database/pgdata_backup:/var/lib/postgresql/data_backup -p 5432:5432 -d postgres",
		pcs.ShellCmd(),
	)
}
