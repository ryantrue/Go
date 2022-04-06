package main

import "fmt"
import "slurm/client"

func main() {
	avatar := client.Avatar{
		URL: "my_url",
	}

	c := client.NewClient("Antony", 27, avatar)
	fmt.Printf("%#v\n", c)
	c.UpdateAvatar()
	fmt.Printf("%#v\n", c)
}