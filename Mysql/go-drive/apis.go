package main

func select(db){
   // 执行查询
   rows, err := db.Query("SELECT id, username, email FROM users")
   if err != nil {
	   log.Fatal(err)
   }
   defer rows.Close()

   // 处理查询结果
   for rows.Next() {
	   var id int
	   var username, email string
	   err = rows.Scan(&id, &username, &email)
	   if err != nil {
		   log.Fatal(err)
	   }
	   fmt.Printf("ID: %d, Username: %s, Email: %s\n", id, username, email)
   }

   // 检查是否有查询错误
   if err = rows.Err(); err != nil {
	   log.Fatal(err)
}
}