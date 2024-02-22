package runn

import "fmt"

type ContainerStart struct {
	Name        string
	PortMapping string
	Image       string
}

func startContainer() {
	//format!("docker run \
	//    --name {name} \
	//    --rm \
	//    --detach \
	//    -p {port}:5432 \
	//    -e POSTGRES_PASSWORD={password} \
	//    postgres\
	//    ", name = args.name, port = args.port, password = args.password);
	fmt.Sprintf("docker run --name %s --rm --detach -p %s-e POSTGRES_PASSWORD=%s postgres", "test", "5432", "password")
}
