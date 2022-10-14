package users

import (
	engine "odibet/Engine"
	"odibet/graph/model"
	"odibet/models"
	"odibet/utils"
)

func AddUser(input model.Creds) (*model.UserData, error) {
	userdata, err := engine.Login(input.Phone, input.Password)
	if err != nil {
		return nil, err
	}

	err = utils.DB.Save(&models.User{
		ProfileID: userdata.User.ProfileID,
		Name:      userdata.User.FirstName,
		Token:     userdata.Token,
		Balance:   userdata.User.Balance,
		Bonus:     userdata.User.BonusBalance,
		Password:  input.Password,
		Phone:     userdata.User.Msisdn,
	}).Error

	if err != nil {
		return nil, err
	}

	accountuserdata := model.UserData{
		ID:      userdata.User.ProfileID,
		Name:    userdata.User.FirstName,
		Token:   userdata.Token,
		Balance: userdata.User.Balance,
		Bonus:   userdata.User.BonusBalance,
		Phone:   userdata.User.Msisdn,
	}

	return &accountuserdata, nil
}

func FetchAllUsers() ([]*model.UserData, error) {
	var users []models.User
	err := utils.DB.Find(&users).Error
	if err != nil {
		return nil, err
	}
	var gqlusers = make([]*model.UserData, len(users))
	for i, user := range users {
		gqlusers[i] = &model.UserData{
			ID:      user.ProfileID,
			Name:    user.Name,
			Token:   user.Token,
			Balance: user.Balance,
			Bonus:   user.Bonus,
			Phone:   user.Phone,
		}
	}
	return gqlusers,nil
}
