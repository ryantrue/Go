package client

type Avatar struct {
	URL string
	Size int64
}

type Client struct {
	id int64
	name string
	age int
	img Avatar
}

func (c Client) HasAvatar() bool {
	return false
}

func  (c *Client) UpdateAvatar()  {
	c.img.URL = "new_url"
	
}

func NewClient(name string, age int, img Avatar) Client  {
	return Client{
		id: 7,
		name: name,
		age: age,
		img: img ,
	}
	
}