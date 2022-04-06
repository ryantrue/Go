package main

import "fmt"

type Avatar struct {
	URL string
	Size int64
}

type Client struct {
	Id int
	Name string
	Age int
	IMG Avatar
}

func main() {
	i := new(int64)
	_ = i

	client := Client{}
	updateAvatar(&client)
fmt.Printf("%#v\n", client)
}

func updateAvatar(client *Client) {
client.IMG.URL = "update_url"
}

func updateClient(client Client) {
	client.Name = "Den"
}