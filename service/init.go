package service

func Init() (err error) {
	err = InitTelegramBot()
	if err != nil {
		return
	}

	return
}
