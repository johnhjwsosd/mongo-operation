mongodb Operation
===
基于mgo封装了读写操作


Installation
-------
<pre><code>
go get github.com/johnhjwsosd/mongo-operation/logoperation
</code></pre>


Example
----------
写
----
<pre><code>
	logser, err1 := m.NewLogServer("mongodb://127.0.0.1:27017")

	content := m.M{
		"info": "3",
		"msg":  18,
	}

	logmodel := &m.Log{"xxx", content, "333", time.Now().String()}

	if err1 == nil {
		err := logser.Writelog("log", "test", logmodel)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("OK")
		}
	} else {
		fmt.Println(err1)
	}
</code></pre>

读
----
<pre><code>
	logser, err1 := m.NewLogServer("mongodb://127.0.0.1:27017")
	where := bson.M{
		"content.info": "3",
		"content.msg":  18,
	}
	log, err2 := logser.ReadlogALL("log", "test", where)
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println(log)
	}
</code></pre>