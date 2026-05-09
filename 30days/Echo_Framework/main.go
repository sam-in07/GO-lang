package main 

type User struct{
    ID int `json:"id"`  	
	Name string `json:"name"`
	Email string `json:"email"`
	Age   int   `json:"age"`
}

func main {
  
	e := echo.new()
	e.Use(middleware)

}