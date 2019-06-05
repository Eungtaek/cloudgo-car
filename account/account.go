package main

import (
    "log"
	"net/http"
	"github.com/labstack/echo"
	"github.com/jinzhu/gorm"
	"fmt"
     _ "github.com/jinzhu/gorm/dialects/mysql"
)

type Account struct {
	User_id  string `json:"user_id" xml:"user_id" form:"user_id" query:"user_id"`
	Membertype  string `json:"membertype" xml:"membertype" form:"membertype" query:"membertype"`
	Pwd  string `json:"pwd" xml:"pwd" form:"pwd" query:"pwd"`
	User_name  string `json:"user_name" xml:"user_name" form:"user_name" query:"user_name"`
	Email  string `json:"email" xml:"email" form:"email" query:"email"`
	Address  string `json:"address" xml:"address" form:"address" query:"address"`
   }
   
func (Account) TableName() string {
	return "account"
}

   
func main() {
	e := echo.New()
	
	//e.Static("/static", "assets")

	
	db, err  := gorm.Open("mysql", "clgo:12345678@tcp(clgo2.ca28h3ng55rn.us-east-2.rds.amazonaws.com)/innodb?charset=utf8&parseTime=True&loc=Local")
 	defer db.Close()
	if err!=nil{
	log.Println("x")	
	}

	
	//e.File("/login", "assets/login.html")
	
	
    e.POST("/join", func(c echo.Context) error {
	        u := new(Account)
	        if err := c.Bind(u); err != nil {
		    return err
	        }
						
			ar :=&Account{
				User_id: u.User_id,
				Membertype: u.Membertype,
				Pwd: u.Pwd,
				User_name: u.User_name,
				Email: u.Email,
				Address: u.Address,
			}
				
			
			db.Create(&ar)
			
	return c.String(http.StatusOK, "join success")
	 })
	 
	 
	e.POST("/update", func(c echo.Context) error {
			u := new(Account)
			if err := c.Bind(u); err != nil {
			return err
			}
	       
		   a := new(Account)
		   
		   db.Table("account").Select("user_id, membertype, pwd, user_name, email, address").Where("user_id = ?", u.User_id).Scan(&a)
		   fmt.Println("<Before>")
		   fmt.Println("user_id: ", a.User_id)
		   fmt.Println("membertype: ", a.Membertype)
		   fmt.Println("pwd: ", a.Pwd)
		   fmt.Println("user_name: ", a.User_name)
		   fmt.Println("email: ", a.Email)
		   fmt.Println("address: ", a.Address)
		   
		   
		   if (u.Membertype== ""){
			   
		   }else{
			 db.Table("account").Where("user_id = ?", u.User_id).Update("membertype", u.Membertype)  
		   }
		   
		   if (u.Pwd== ""){
			   
		   }else{
			 db.Table("account").Where("user_id = ?", u.User_id).Update("pwd", u.Pwd)  
		   }
		   
		   if (u.User_name== ""){
			   
		   }else{
			 db.Table("account").Where("user_id = ?", u.User_id).Update("user_name", u.User_name)  
		   }
		   
		   if (u.Email== ""){
			   
		   }else{
			 db.Table("account").Where("user_id = ?", u.User_id).Update("email", u.Email)  
		   }
		   
		   if (u.Address== ""){
		  
		   }else{
			 db.Table("account").Where("user_id = ?", u.User_id).Update("address", u.Address)  
		   }
		   
		   db.Table("account").Select("user_id, membertype, pwd, user_name, email, address").Where("user_id = ?", u.User_id).Scan(&a)
		   fmt.Println("<After>")
		   fmt.Println("user_id: ", a.User_id)
		   fmt.Println("membertype: ", a.Membertype)
		   fmt.Println("pwd: ", a.Pwd)
		   fmt.Println("user_name: ", a.User_name)
		   fmt.Println("email: ", a.Email)
		   fmt.Println("address: ", a.Address)		   

		   
		
			
	return c.String(http.StatusOK, "update success")		
	 })
	 
	
	e.POST("/delete", func(c echo.Context) error {
	        u := new(Account)
	        if err := c.Bind(u); err != nil {
		    return err
	        }
	       
		  db.Table("account").Where("user_id= ?", u.User_id).Delete("account")
	      	if err!=nil{
	        log.Println("x")	
        	}
			
	return c.String(http.StatusOK, "delete success")		
	 })

	

	e.Logger.Fatal(e.Start(":8080"))
}



