package models

import
(
	coreglobals "collexy/core/globals"
	"encoding/json"
	"log"
)

type ShoppingCart struct {
	Id int `json:"id,omitempty"`
	CookieId string `json:"cookie_id,omitempty"`
	Items []ShoppingCartItem `json:"items,omitempty"`
}

func (this *ShoppingCart) Post(cookieId string){
	db := coreglobals.Db

	sqlStr := `INSERT INTO colecom_cart (cookie_id, items) VALUES ($1, $2)`

	itemsStr, err := json.Marshal(this.Items)

	if err != nil{
		log.Println(err)
	}

	_, err1 := db.Exec(sqlStr,cookieId, itemsStr)

	if err1 != nil{
		log.Println(err1)
	}
}

func GetShoppingCart(cookieId string) (shoppingCart *ShoppingCart){
	var id int
	var items []byte

	db := coreglobals.Db

	sqlStr := `SELECT id, items FROM colecom_cart WHERE cookie_id=$1`

	row := db.QueryRow(sqlStr, cookieId)

	err := row.Scan(&id, &items)

	if err != nil {
		log.Println("Err in GetShoppingCart")
		return nil
	}

	var itemsSlice []ShoppingCartItem
	err1 := json.Unmarshal(items, &itemsSlice)

	if err1 != nil {
		log.Println("Err unmarshalling shopping cart items")
		return nil
	}

	shoppingCart = &ShoppingCart{id, cookieId, itemsSlice}

	return
}

func (this *ShoppingCart) Put(){
	db := coreglobals.Db

	sqlStr := `UPDATE colecom_cart SET items=$1 WHERE id=$2`

	itemsStr, err := json.Marshal(this.Items)

	if err != nil{
		log.Println(err)
	}

	_, err1 := db.Exec(sqlStr,itemsStr, this.Id)

	if err1 != nil{
		log.Println(err1)
	}
}