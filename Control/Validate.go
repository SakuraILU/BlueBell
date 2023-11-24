package control

import (
	config "bluebell/Config"
	model "bluebell/Model"
)

func validateSignUp(param *model.ParamSignUp) bool {
	usr_range := config.Cfg.Logic.UsernameRange
	pwd_range := config.Cfg.Logic.PasswordRange
	if len(param.Username) <= usr_range[0] || len(param.Username) >= usr_range[1] {
		return false
	}

	if len(param.Password) <= pwd_range[0] || len(param.Password) >= pwd_range[1] {
		return false
	}

	if param.Password != param.RePassword {
		return false
	}

	return true
}

func validateLogin(param *model.ParamLogin) bool {
	usr_range := config.Cfg.Logic.UsernameRange
	if len(param.Username) <= usr_range[0] || len(param.Username) >= usr_range[1] {
		return false
	}

	return true
}

func validateVote(param *model.ParamVote) bool {
	if param.Choice != model.POSITIVE && param.Choice != model.NEGATIVE && param.Choice != model.CANCEL {
		return false
	}

	return true
}

func validatePostsQuary(param *model.ParamPostsQuary) bool {
	if param.Page < 1 {
		return false
	}

	max_pagesize := config.Cfg.Logic.MaxPageSize
	if param.Size < 1 || param.Size > max_pagesize {
		return false
	}

	if (param.Order != model.SCORE) && (param.Order != model.TIME) {
		return false
	}

	return true
}
