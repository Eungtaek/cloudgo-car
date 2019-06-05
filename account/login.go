package main

import (
    "log"
	"net/http"
	"github.com/labstack/echo"
	"github.com/jinzhu/gorm"
	//"fmt"
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
   
			type Result struct {
                User_id string
                Pwd  string
            }			
	var result Result
   
func (Account) TableName() string {
	return "account"
}
   
func main() {
	e := echo.New()
	
	e.Static("/static", "assets")

	
	type AccountLogin struct {
	User_id  string `json:"user_id" xml:"user_id" form:"user_id" query:"user_id"`
	Pwd  string `json:"pwd" xml:"pwd" form:"pwd" query:"pwd"`
	
    }
	
//	type AccountLoginResult struct {
//    User_id  string `json:"user_id" xml:"user_id" form:"user_id" query:"user_id"`
//	Pwd  string `json:"pwd" xml:"pwd" form:"pwd" query:"pwd"`
//	Result string `json:"result" xml:"result" form:"result" query:"result"`
//    }
	
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
	//return c.JSON(http.StatusCreated, r)
	// or
	// return c.XML(http.StatusCreated, u)
    
	
//	e.POST("/update", func(c echo.Context) error {
//	        u := new(Account)
//	        if err := c.Bind(u); err != nil {
//		    return err
//	        }
//	  var user_name String 		
//	
//	  ur := new(Account)		
//      db.Select(ur.User_name).Table(account).Where("User_id = ?", "helciger").Scan(&user_name)
//			
//	return c.String(http.StatusOK, "update success")		
//	 })
//	
//	
//	
//	
//	e.POST("/login", func(c echo.Context) error {
//	        u := new(AccountLogin)
//	        if err := c.Bind(u); err != nil {
//		    return err
//	        }
//			
//			
//			
//			
//			db.Raw("SELECT user_id, pwd FROM account WHERE user_id = ? ", u.User_id ).Scan(&result)
//			fmt.Println(result.User_id)

			
//			ar :=&Account{
//				User_id: u.User_id,
//				Pwd: u.Pwd,
//
//			}
			
			//fmt.Println(db.Table("account").Select("User_name").Where("User_id = ?", u.User_id))
			//if ( u.User_name == db.Select("User_name").Where("User_id = ?", u.User_id).Find(&ar) ) (
			//    fmt.Println("login:", u.User_id)
			//)
				
			
//			r := &AccountLoginResult{
//				   User_id: u.User_id,
//			       Pwd: u.Pwd,
////				Result: "반갑습니다",
//			}
//			
//			if(u.User_id == oriId && u.Pwd == oriPwd) {
//				r.Result = "success"
//			}else{
//				r.Result = "fail"
//			}
//	return c.String(http.StatusOK, "login success")		
//	})
//	return c.String(http.StatusOK, "login success")
	
	// or
	// return c.XML(http.StatusCreated, u)
    

	

	e.Logger.Fatal(e.Start(":8080"))
}



