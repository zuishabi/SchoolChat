# 网址

* http://school.zuishabi.top

# 需求分析

* 用户登录注册：
  * 用户通过输入自己的学号以及密码向服务器进行验证登录信息，登录会记录session，保存学号，名称，毕业年份信息。
  * 用户输入自己的学号，名字，密码，毕业年份创建用户，服务器会验证学号是否重复
* 主页面：
  * 用户登录后会显示主页面，分为查找班级模块，创建班级模块和显示加入的班级模块
  * 查找班级可以通过毕业的年份查找到当年毕业的所有班级，并可以请求加入，可以请求多个班级，但是一旦成功加入一个其他的请求都会失效
  * 创建班级可以输入用户的班级名称和班级简介，毕业年份会自动通过当前用户的毕业年份决定，无法修改，同时会检查是否已经加入了班级，若已加入则无法创建班级
  * 显示加入的班级可以查看班级中的用户列表，并进入对应的个人主页，同时可以选择管理班级和退出班级，管理班级需要验证管理员身份。
  * 用户可以检索其他用户的信息，比如按照学号，姓名和毕业时间来进行查询，会显示对应用户的基本信息。
  * 导航栏中可以进入个人主页和退出当前登录，清除session信息。
* 用户主页：
  * 首先会显示用户的个人信息，用户名，学号，毕业年份和个人简介信息
  * 服务器会验证访问主页的是否是用户自己，如果是，则会显示修改个人信息的栏目，可以修改用户名，毕业时间以及个人简介的信息，但是无法修改用户学号
  * 用户可以通过留言板进行留言。

# ER图

![image-20250616144759758](C:\Users\张俏悦\AppData\Roaming\Typora\typora-user-images\image-20250616144759758.png)

# 接口文档

* ```go
  POST("/", routers.LoginHandler)
  // 验证参数信息
  
  // 请求值
  type LoginInfo struct {
  	UID      string `json:"uid"`
  	Password string `json:"password"`
  }
  ```

* ```go
  POST("register", routers.RegisterHandler)
  // 验证注册信息
  
  // 请求值
  type RegisterInfo struct {
  	UID          string `json:"uid"`
  	UserName     string `json:"user_name"`
  	Password     string `json:"password"`
  	GraduateTime string `json:"graduate_time"`
  }
  ```

* ```go
  GET("/searchUser", routers.SearchUser)
  // 查询用户，url参数：uid，用户学号 user_name,用户名 graduate_time，毕业时间
  
  // 返回值
  type SearchInfo struct {
  	ID           string `json:"id"`
  	Name         string `json:"name"`
  	Description  string `json:"description"`
  	GraduateTime string `json:"graduate_time"`
  }
  ```

  
  
* ```go
  GET("/getUserInfo", routers.GetUserInfo)
  // 返回个人信息，url参数：uid，查询者的uid
  
  // 返回值
  type UserInfoReq struct {
  	Name         string `json:"name"`
  	CurrentUID   string `json:"current_uid"`
  	ID           string `json:"student_id"`
  	GraduateTime string `json:"graduate_time"`
  	Description  string `json:"description"`
  }
  ```

* ```go
  POST("/updateUserInfo", routers.UpdateUserInfo)
  // 更新一个用户的个人信息
  
  // 请求值
  type UpdateUserInfoReq struct {
  	Name         string `json:"name"`
  	GraduateTime string `json:"graduate_time"`
  	Description  string `json:"description"`
  }
  ```

* ```go
  GET("/getMessages", routers.ChatBoard)
  // 获得留言板内容，url参数：page，请求的页码  uid，查询者的uid
  
  //返回值
  type Chat struct {
  	Messages []Message `json:"messages"`
  	Total    int64     `json:"total_pages"`
  }
  
  type Message struct {
  	Name    string `json:"from_name"`
  	Time    string `json:"time"`
  	Content string `json:"content"`
  }
  ```

* ```go
  POST("/leaveMessage", routers.CreateChat)
  // 进行留言
  
  // 请求值
  type sendChat struct {
  	UID     string `json:"uid"`
  	Content string `json:"content"`
  }
  ```

* ```go
  GET("/createClass", routers.CreateClass)
  // 创建一个班级，url参数：class_name，班级的名称 class_description：班级的描述
  ```

* ```go
  GET("/getClass", routers.GetClass)
  // 返回当前用户加入的班级的信息
  
  // 返回值
  type classInfo struct {
  	CID          uint32 `json:"cid"`
  	CName        string `json:"c_name"`
  	CDescription string `json:"c_description"`
  	GraduateTime string `json:"graduate_time"`
  	OP           string `json:"op"` //管理员id
  }
  ```

* ```go
  GET("/getClassMembers", routers.GetClassMembers)
  // 返回当前房间的成员列表
  
  // 返回值
  type Member struct {
  	ID   string `json:"id"`
  	Name string `json:"name"`
  }
  ```

* ```go
  GET("/searchClass", routers.SearchClass)
  // 查询一个班级，url参数：graduate_time，毕业时间
  
  // 返回值
  type ClassList struct {
  	Name        string `json:"name"`
  	ID          uint32 `json:"id"`
  	Description string `json:"description"`
  }
  ```

* ```go
  GET("/addClass", routers.AddClass)
  // 请求加入一个班级，url参数：cid，班级id
  ```

* ```go
  GET("/deleteClass", routers.DeleteClass)
  // 解散一个班级，url参数：cid，班级id
  ```

* ```go
  GET("/getEnterRequest", routers.GetEnterRequest)
  // 获得用户申请加入列表
  
  // 返回值
  type Member struct {
  	ID   string `json:"id"`
  	Name string `json:"name"`
  }
  ```

* ```go
  GET("/approveEnter", routers.ApproveEnter)
  // 同意加入班级，url参数：uid，用户id cid，班级id
  
  ```

* ```go
  GET("/rejectEnter", routers.RejectEnter)
  // 拒绝加入班级，url参数：uid，用户id cid，班级id
  ```

* ```go
  GET("/kickMember", routers.RemoveUser)
  // 剔除用户，url参数：uid，用户id cid，班级id
  ```

* ```go
  GET("/quitClass", routers.QuitClass)
  // 退出班级
  ```

  

