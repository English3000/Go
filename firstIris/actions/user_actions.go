package actions

import (
  "errors"

  "../datamodels"
  "../sql"
)

type UserActions interface { //declaration of interface functions
  GetAll() []datamodels.User
  GetByID(id int64) (datamodels.User, bool)
  GetByUsernameAndPassword(username, password string) (datamodels.User, bool) //in user.rb
  DeleteByID(id int64) bool

  Update(id int64, user datamodels.User) (datamodels.User, error)
  UpdateUsername(id int64, newUsername string) (datamodels.User, error)
  UpdatePassword(id int64, newPW string) (datamodels.User, error)

  //needs to hash & assign given PW
  Create(password string, user datamodels.User) (datamodels.User, error)
}

func NewUserActions(sql sql.UserQueries) UserActions {
  return &userActions{
    queries: sql,
  }
}

type userActions struct {
  queries sql.UserQueries
}

func (u *userActions) GetAll() []datamodels.User {
  return u.queries.SelectMany(func(_ datamodels.User) bool {
    return true
  }, -1)
}

func (u *userActions) GetByID(id int64) (datamodels.User, bool) {
  return u.queries.Select(func(m datamodels.User) bool {
    return m.ID == id
  })
}

func (u *userActions) GetByUsernameAndPassword(username, password string) (datamodels.User, bool) {
  if username == "" || password = "" {
    return datamodels.User{}, false
  }

  return u.queries.Select(func(m datamodels.User) bool {
    if m.Username == username {
      hashed := m.HashedPassword
      if ok, _ := datamodels.ValidatePassword(password, hashed); ok {
        return true
      }
    }
    return false
  })
}

//Update

//UpdateUsername

//UpdatePassword

//Create
