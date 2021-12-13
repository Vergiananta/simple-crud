package main

import (
	_ "github.com/lib/pq"
	"learn-go/routes"
)

//type Customer struct {
//	id      int
//	nama    string
//	umur    int
//	address string
//}

func main() {
	routes.NewPerpusApp().Run()
	//db, err := connect()
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}

	// Create
	// _, err = db.Query("insert into customers (nama,umur,address) values ('egi', 21, 'kemang')")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	// Update
	// _, err = db.Query("UPDATE customers SET nama = 'doni' where id = 1")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	// Read
	// var result []Customer
	// rows, err := db.Query("SELECT c.id,c.nama,c.umur,c.address FROM customers as c")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	// for rows.Next() {
	// 	var cust = Customer{}
	// 	err = rows.Scan(&cust.id,&cust.nama,&cust.umur,&cust.address)
	// 	if err != nil {
	// 		fmt.Println(err.Error())
	//	 return
	// 	}
	// 	result = append(result,cust)
	// }
	// fmt.Println(result)

	// Delete
	// _, err = db.Query("Delete from customers where id = 2")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	// select a.id, a.username, c.nama as profile_name, c.address from accounts as a join customers as c on a.customer_id = c.id;

	//var result []DataProfileResponse
	//rows, err := db.Query("select a.id, a.username, c.nama as profile_name, c.address from accounts as a join customers as c on a.customer_id = c.id;")
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}
	//
	//for rows.Next() {
	//	var cust = DataProfileResponse{}
	//	err = rows.Scan(&cust.id, &cust.username, &cust.ProfileName, &cust.Address)
	//	if err != nil {
	//		fmt.Println(err.Error())
	//		return
	//	}
	//	result = append(result, cust)
	//}
	//fmt.Println(result)
	//
	//defer db.Close()
}
