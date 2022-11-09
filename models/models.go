package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)


type User struct {
	ID       primitive.ObjectID     `json:"_id"  bson:"_id"`
	Name      *string               `json:"first_name"`  
	LastName  *string               `json:"last_name"`
	Password  *string               `json:"password"`
	Email     *string               `json:"email"`
	Phone     *string               `json:"phone"`
	Token     *string               `json:"token"`
	RefrishToken *string            `json:"refrish_token"`
	Created_at    time.Time         `json:"created_at"`
	Update_at    time.Time          `json:"updated_at"`
	User_ID   string                `json:"user_id"`
	UserCart   []ProductUser        `json:"user_cart"`
	Address       []Address         `json:"address"`
	OrderStatus []Order             `json:"order_status"`
} 

type Product struct {
	Product_ID          primitive.ObjectID `json:"_id"`
	ProductName  *string                   `json:"product_name"`
	Price        *uint64                   `json:"price"`
	Rating       *int                      `json:"rating"`
	Image        *string                   `json:"image"`

}

type ProductUser struct {
	Product_ID          primitive.ObjectID      `json:"product_id"`
	ProductName *string                         `json:"product_name"`
	Price        *uint64                        `json:"price"`
	Rating       *int                           `json:"rating"`
	Image        *string                        `json:"image"`

}

type Address struct {
	Address_ID      primitive.ObjectID           `json:"address_id"`
	House          *string                       `json:"house"`
	Street         *string                       `json:"street"`
	City           *string                       `json:"city"`
	Pincode        *string                       `json:"pincode"`

}

type Order struct {

	Order_ID          primitive.ObjectID          `json:"order_id"`
	Order_Cart        []ProductUser               `json:"order_cart"`
	Ordered_At        time.Time                   `json:"ordered_at"`
	Price             *uint64                     `json:"price"`
	Discount          *int                        `json:"discount"`
	Payment_Method    Payment                     `json:"payment_method"`
 
}

type Payment struct {
	Digital bool
	COD bool
}