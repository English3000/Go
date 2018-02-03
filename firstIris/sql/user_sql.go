package sql

import (
	"errors"
	"sync"

	"../datamodels"
)

type Query func(datamodels.User) bool

type UserQueries interface {
	Exec(query Query, action Query, limit int, mode int) (ok bool)

	Select(query Query) (user datamodels.User, found bool)
	SelectMany(query Query, limit int) (results []datamodels.User)

	InsertOrUpdate(user datamodels.User) (updatedUser datamodels.User, err error)
	Delete(query Query, limit int) (deleted bool)
}

func UsersTable(table map[int]datamodels.User) UserQueries {
	return &userMemoryRepo{table: table}
}

type userMemoryRepo struct {
	table map[int64]datamodels.User
	mu    sync.RWMutex
}

const (
	//RLock(read)
	ReadOnlyMode = iota
	//Lock(read/write)
	ReadWriteMode
)

//why query && action??
func (r *userMemoryRepo) Exec(query Query, action Query, limit int, mode int) (ok bool) {
	selections := 0

	if mode == ReadOnlyMode {
		r.mu.RLock()
		defer r.mu.RUnlock()
	} else {
		r.mu.Lock()
		defer r.mu.Unlock()
	}

	for _, user := range r.table {
		ok = query(user)
		if ok {
			if action(user) {
				selections++
				if limit >= selections {
					break
				}
			}
		}
	}

	return
}

func (r *userMemoryRepo) Select(query Query) (user datamodels.User, found bool) {
	found = r.Exec(query, func(m datamodels.User) bool {
		user = m
		return true
	}, 1, ReadOnlyMode)

	if !found {
		user = datamodels.User{}
	}

	return
}

func (r *userMemoryRepo) SelectMany(query Query, limit int) (results []datamodels.User) {
	r.Exec(query, func(m datamodels.User) bool {
		results = append(results, m)
		return true
	}, limit, ReadOnlyMode)

	return
}

func (r *userMemoryRepo) InsertOrUpdate(user datamodels.User) (datamodels.User, error) {
	id := user.ID

	//insert/create
	if id == 0 {
		var lastID int64 //0

		r.mu.RLock()
		for _, item := range r.table { //go thru UsersTable
			if item.ID > lastID {
				lastID = item.ID
			}
		}
		r.mu.RUnlock()

		id = lastID + 1
		user.ID = id

		r.mu.Lock()
		r.table[id] = user
		r.mu.Unlock()

		return user, nil
	}

	_, exists := r.Select(func(user datamodels.User) bool {
		return user.ID == id
	})

	if !exists {
		return datamodels.User{}, errors.New("user does not exist")
	}

	r.mu.Lock()
	r.table[id] = user
	r.mu.Unlock()

	return user, nil
}

func (r *userMemoryRepo) Delete(query Query, limit int) bool {
	return r.Exec(query, func(user datamodels.User) bool {
		delete(r.table, user.ID)
		return true
	}, limit, ReadWriteMode)
}
