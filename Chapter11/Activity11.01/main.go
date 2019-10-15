package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type customer struct {
	UserName      string  `json:"username"`
	Password      string  `json:"-"`
	Token         string  `json:"-"`
	ShipTo        address `json:"shipto"`
	PurchaseOrder order   `json:"order"`
}

type order struct {
	TotalPrice  int    `json:"total"`
	IsPaid      bool   `json:"paid"`
	Fragile     bool   `json:",omitempty"`
	OrderDetail []item `json:"orderdetail"`
}

type item struct {
	Name        string `json:"itemname"`
	Description string `json:"desc,omitempty"`
	Quantity    int    `json:"qty"`
	Price       int    `json:"price"`
}

type address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode int    `json:"zipcode"`
}

func main() {
	jsonData := []byte(`
	{
		"username" :"blackhat",
		"shipto":  
			{
		    		"street": "Sulphur Springs Rd",
	    			"city": "Park City",
	    			"state": "VA",
	    			"zipcode": 12345

			},
		"order":
			{
				"paid":false,
				"orderdetail" : 
					[{
						"itemname":"A Guide to the World of zeros and ones",
						"desc": "book",
						"qty": 3,
						"price": 50

					}]
			
			}
		
	}
	`)

	if !json.Valid(jsonData) {
		fmt.Printf("JSON is not valid: %s", jsonData)
		os.Exit(1)
	}
	var c customer
	err := json.Unmarshal(jsonData, &c)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	game := item{}
	game.Name = "Final Fantasy The Zodiac Age"
	game.Description = "Nintendo Switch Game"
	game.Quantity = 1
	game.Price = 50

	glass := item{}
	glass.Name = "Crystal Drinking Glass"
	glass.Quantity = 11
	glass.Price = 25

	c.PurchaseOrder.OrderDetail = append(c.PurchaseOrder.OrderDetail, game)
	c.PurchaseOrder.OrderDetail = append(c.PurchaseOrder.OrderDetail, glass)
	c.Total()
	c.PurchaseOrder.IsPaid = true
	c.PurchaseOrder.Fragile = true

	customerOrder, err := json.MarshalIndent(c, "", "    ")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(customerOrder))

}

func (c *customer) Total() {
	price := 0
	for _, item := range c.PurchaseOrder.OrderDetail {
		price += item.Quantity * item.Price
	}
	c.PurchaseOrder.TotalPrice = price
}
