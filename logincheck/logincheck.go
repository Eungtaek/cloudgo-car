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

type Car struct {
	User_id  string `json:"user_id" xml:"user_id" form:"user_id" query:"user_id"`
	Membertype  string `json:"membertype" xml:"membertype" form:"membertype" query:"membertype"`
	Pwd  string `json:"pwd" xml:"pwd" form:"pwd" query:"pwd"`
	User_name  string `json:"user_name" xml:"user_name" form:"user_name" query:"user_name"`
	Email  string `json:"email" xml:"email" form:"email" query:"email"`
	Address  string `json:"address" xml:"address" form:"address" query:"address"`
   }
   
type Cartype struct {
	User_id  string `json:"user_id" xml:"user_id" form:"user_id" query:"user_id"`
	Membertype  string `json:"membertype" xml:"membertype" form:"membertype" query:"membertype"`
	Pwd  string `json:"pwd" xml:"pwd" form:"pwd" query:"pwd"`
	User_name  string `json:"user_name" xml:"user_name" form:"user_name" query:"user_name"`
	Email  string `json:"email" xml:"email" form:"email" query:"email"`
	Address  string `json:"address" xml:"address" form:"address" query:"address"`
   }
   
type Loginresult struct {
	Result  string 
   }
   
   
func (Account) TableName() string {
	return "account"
}

func (Car) TableName() string {
	return "car"
}

func (Cartype) TableName() string {
	return "cartype"
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
	
	
    e.POST("/login", func(c echo.Context) error {
	        u := new(Account)
	        if err := c.Bind(u); err != nil {
		    return err
	        }
			
			a := new(Account)
		    db.Table("account").Select("user_id, pwd").Where("user_id = ?", u.User_id).Scan(&a)			
			
			rst := new(Loginresult)
			
			
			if(a.User_id == ""){
				fmt.Println("Join please")
			}else if (u.Pwd == a.Pwd ){
				fmt.Println("login success")
				rst.Result = "O"
//				fmt.Println(rst.Result)
			}else {
				fmt.Println("please check your password ")
			}
			
		
		
	return c.String(http.StatusOK, "login tr")
	 })
	 
	


	e.Logger.Fatal(e.Start(":8080"))
 }



