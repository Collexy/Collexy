package models

type ShoppingCartViewModel struct {
	CookieId string
	Items []ShoppingCartItemViewModel
	SubTotalPrice float64
	TotalPrice float64
	ShippingMethod *ShippingMethodViewModel
	PaymentMethod *PaymentMethodViewModel
	ShippingPrice float64
	TaxAmount float64
	BillingAddress *AddressViewModel
	ShippingAddress *AddressViewModel
	ApiModel *ShoppingCart
}

func (this *ShoppingCartViewModel) GetItems() ([]ShoppingCartItemViewModel){

}

func (this *ShoppingCartViewModel) SetBillingAddress(a AddressViewModel){
	this.BillingAddress = a
	ApiModel.SetBillingAddress(a)
}

func (this *ShoppingCartViewModel) SetShippingAddress(a AddressViewModel){
	this.ShippingAddress = a
	ApiModel.SetShippingAddress(a)
}

func (this *ShoppingCartItemViewModel) AddItem(item ShoppingCartItemViewModel){
	for _, i := range this.Items {
		if i.Id == item.Id{
			if i.Attributes == item.Attributes {
				i.Quantity = i.Quantity + 1
			} else {
				this.Items = append(this.Items, item)
			}
			ApiModel.Update(this)
			break
		}
	}
	
	
}

func (this *ShoppingCartItemViewModel) RemoveItem(id int){
	for _, item := range this.Items {
		if item.Id == id {
			// remove item from slice
			shoppingCartApiModel := ShoppingCart{}
			ApiModel.Update(this)
			break
		}
	}
}