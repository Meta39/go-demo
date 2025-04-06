package controller

import (
	"database/sql"
	"encoding/json"
	"mysql/model"
	"mysql/router"
	"net/http"
)

// UserHandler 实现 ResourceHandler 接口
type UserHandler struct {
	DB *sql.DB
}

// Create 创建用户（包含事务控制）
func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	//开启事务
	tx, err := h.DB.Begin()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//报错回滚事务
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := user.Create(tx) //返回新增的id值和错误
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//提交事务
	if err := tx.Commit(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user.Id = int(id) //赋值新增用户的Id值返回给前端
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// FindByID 查找用户
func (h *UserHandler) FindByID(w http.ResponseWriter, r *http.Request) {
	id, err := router.ExtractIDFromPath(r.URL.Path, "users")
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := model.FindByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// Update 更新用户（包含事务控制）
func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	tx, err := h.DB.Begin()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err = user.Update(tx); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tx.Commit(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// Delete 删除用户（包含事务控制）
func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	tx, err := h.DB.Begin()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	id, err := router.ExtractIDFromPath(r.URL.Path, "users")
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	if err = model.Delete(tx, id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tx.Commit(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
