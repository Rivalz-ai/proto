package utils
//add new var to Args if not exist,
func AppendArgs(args []interface{},key string, val interface{}) []interface{}{
	var new_args []interface{}
	if len(args)>0{
		m, err := ItoDictionary(args[0])
		if err == nil {
			if m[key]!=nil{
				m[key]=val
				new_args=append(new_args,m)
				new_args=append(new_args,args[1:])
			}
			return new_args
		}else{
			return args 
		}
	}else{
		m:=Dictionary()
		m[key]=val
		new_args=append(new_args,m)
	}
	return new_args
}