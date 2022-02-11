package main

import (
	"github.com/kataras/iris/v12"
	"github.com/sideneckx/template-go-server/json_type"
)

// type JSONExample struct {
// 	body string
// }

// func main() {
// 	reps, err := http.Get("https://jsonplaceholder.typicode.com/posts")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	defer reps.Body.Close()
// 	var jsonBody *[]map[string]interface{} = &[]map[string]interface{}{}
// 	json.NewDecoder(reps.Body).Decode(jsonBody)
// 	fmt.Println((*jsonBody)[0])
// }

func main() {
	app := iris.New()
	app.Handle("GET", "/home", func(ctx iris.Context) {
		ctx.JSON([]json_type.RawJson{
			{
				"message": "1",
			},
			{
				"message": "ahuhuhu",
			},
		})
	})
	app.Listen(":8080")
}
