package helper

func IfError(err error) error {

	if err != nil {
		panic(err)

	}
	return err

}
