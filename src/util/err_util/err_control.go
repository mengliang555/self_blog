package err_util

func PanicError(err error) {
	if err != nil {
		panic(err)
	}
}
