package control

import model "bluebell/Model"

func validateSignUp(param *model.ParamSignUp) bool {
	if len(param.Username) <= 0 || len(param.Username) > 15 {
		return false
	}

	if len(param.Password) < 6 || len(param.RePassword) < 6 || len(param.Password) > 25 || len(param.RePassword) > 25 {
		return false
	}

	if param.Password != param.RePassword {
		return false
	}

	return true
}

func validateLogin(param *model.ParamLogin) bool {
	if len(param.Username) <= 0 || len(param.Password) < 6 {
		return false
	}

	return true
}
