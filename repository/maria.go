package repository

// conn, err := sql.Open("mysql", "root:dkssud123@tcp(127.0.0.1:3306)/test")
// if err != nil {
// 	fmt.Println(err)
// 	//os.Exit(1)
// }

// // use your own select statement
// // this is just an example statement
// statement, err := conn.Prepare("select level from keyword")

// if err != nil {
// 	fmt.Println(err)
// 	//os.Exit(1)
// }

// rows, err := statement.Query() // execute our select statement

// if err != nil {
// 	fmt.Println(err)
// 	//os.Exit(1)
// }

// for rows.Next() {
// 	var level string
// 	rows.Scan(&level)
// 	fmt.Println("level is :", level)
// }

// defer conn.Close()
