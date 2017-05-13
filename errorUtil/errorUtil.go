package errorUtil

func CheckErrorExit(err error)  {
	if err != nil {
		panic(err)
	}
}

func CheckErrorContinue(err error)  {
	if err != nil{
		println(err.Error())
	}
}