package repository

import "context"

func (r *Repository) GetTestById(ctx context.Context, input GetTestByIdInput) (output GetTestByIdOutput, err error) {
	err = r.Db.QueryRowContext(ctx, "SELECT name FROM test WHERE id = $1", input.Id).Scan(&output.Name)
	if err != nil {
		return
	}
	return
}

func (r *Repository) InsertUser(ctx context.Context, input UserData) (output UserId, err error) {
	err = r.Db.QueryRowContext(ctx, "INSERT INTO users (full_name,phone_number,password,login_count) VALUES ($1, $2, $3, 0)  RETURNING user_id;",
		input.FullName, input.PhoneNumber, input.Password).Scan(&output.UserID)
	if err != nil {
		return
	}
	return
}

func (r *Repository) GetUserDataByPhoneNumber(ctx context.Context, input UserData) (userData UserData, userId UserId, err error) {
	err = r.Db.QueryRowContext(ctx, "SELECT user_id, password FROM users WHERE phone_number = $1", input.PhoneNumber).Scan(&userId.UserID, &userData.Password)
	if err != nil {
		return
	}
	return
}

func (r *Repository) GetUserByPhoneNumber(ctx context.Context, input UserData) (userId UserId, err error) {
	row, err := r.Db.QueryContext(ctx, "SELECT user_id FROM users WHERE phone_number = $1", input.PhoneNumber)
	if row.Next() {
		err = row.Scan(&userId.UserID)
	}
	if err != nil {
		return
	}
	return
}

func (r *Repository) GetUserDataById(ctx context.Context, input UserId) (output UserData, err error) {
	err = r.Db.QueryRowContext(ctx, "SELECT full_name, phone_number FROM users WHERE user_id = $1", input.UserID).Scan(&output.FullName, &output.PhoneNumber)
	if err != nil {
		return
	}
	return
}

func (r *Repository) UpdatePhoneNumber(ctx context.Context, userData UserData, userId UserId) (err error) {
	_, err = r.Db.ExecContext(ctx, "UPDATE users SET phone_number = $1 WHERE user_id = $2", userData.PhoneNumber, userId.UserID)
	if err != nil {
		return
	}
	return
}

func (r *Repository) UpdateFullName(ctx context.Context, userData UserData, userId UserId) (err error) {
	_, err = r.Db.ExecContext(ctx, "UPDATE users SET full_name = $1 WHERE user_id = $2", userData.FullName, userId.UserID)
	if err != nil {
		return
	}
	return
}

func (r *Repository) UpdateUserData(ctx context.Context, userData UserData, userId UserId) (err error) {
	_, err = r.Db.ExecContext(ctx, "UPDATE users SET full_name = $1, phone_number = $2 WHERE user_id = $3", userData.FullName, userData.PhoneNumber, userId.UserID)
	if err != nil {
		return
	}
	return
}

func (r *Repository) UpdateLoginCount(ctx context.Context, userId UserId) (err error) {
	_, err = r.Db.ExecContext(ctx, "UPDATE users SET login_count = login_count + 1 WHERE user_id = $1", userId.UserID)
	if err != nil {
		return
	}
	return
}
