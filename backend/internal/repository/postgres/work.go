package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/lib/pq"
	"github.com/rs/zerolog/log"
	"server/internal/entities"
)

type WorkRepository struct {
	db *sql.DB
}

func NewWorkRepository(db *sql.DB) *WorkRepository {
	return &WorkRepository{
		db: db,
	}
}

func (r *WorkRepository) Exists(ctx context.Context, id int) (bool, error) {
	var exists int
	query := `SELECT 1 FROM work_user WHERE id = $1`
	err := r.db.QueryRowContext(ctx, query, id).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return exists > 0, nil
}

func (r *WorkRepository) ExistsHR(ctx context.Context, university string) (bool, error) {
	var exists int
	query := `SELECT 1 FROM hr WHERE email = $1`
	err := r.db.QueryRowContext(ctx, query, university).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return exists > 0, nil
}

func (r *WorkRepository) Create(ctx context.Context, workUser *entities.WorkUser) error {
	if check, _ := r.Exists(ctx, workUser.Id); check != true {
		query := `INSERT INTO work_user (id, speciality, work_experience, additional_experience, useful_links, phone_number, telegram, skills, cv_path, hide_profile) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`
		_, err := r.db.ExecContext(ctx, query, workUser.Id, workUser.Speciality, workUser.WorkExperience, workUser.AdditionalExperience, pq.Array(workUser.UsefulLinks), workUser.PhoneNumber, workUser.Telegram, pq.Array(workUser.Skills), workUser.CVPath, workUser.HideProfile)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *WorkRepository) CreateHR(ctx context.Context, hr *entities.HR) error {
	if check, _ := r.ExistsHR(ctx, hr.Email); check != true {
		query := `INSERT INTO hr (email, password, last_name, first_name, father_name, company, image_link) VALUES ($1, $2, $3, $4, $5, $6, $7)`
		_, err := r.db.ExecContext(ctx, query, hr.Email, hr.Password, hr.LastName, hr.FirstName, hr.FatherName, hr.Company, hr.ImageLink)
		if err != nil {
			return err
		}

	} else {
		log.Info().Msg(fmt.Sprintf("hr %v already exists", hr))
	}
	return nil
}

func (r *WorkRepository) GetByIdWorkUser(ctx context.Context, id int) (*entities.FullWorkUser, error) {
	var fullWorkUser entities.FullWorkUser
	query := `SELECT work_user.id, speciality, work_experience, additional_experience, useful_links, phone_number, telegram, skills, last_name, first_name, father_name, university.name, gpa.value FROM work_user JOIN users ON work_user.id = users.id JOIN university ON users.university_id = university.id JOIN gpa ON users.id = gpa.user_id WHERE work_user.id = $1;`
	err := r.db.QueryRowContext(ctx, query, id).Scan(&fullWorkUser.Id, &fullWorkUser.Speciality, &fullWorkUser.WorkExperience, &fullWorkUser.AdditionalExperience, pq.Array(&fullWorkUser.UsefulLinks), &fullWorkUser.PhoneNumber, &fullWorkUser.Telegram, pq.Array(&fullWorkUser.Skills), &fullWorkUser.LastName, &fullWorkUser.FirstName, &fullWorkUser.FatherName, &fullWorkUser.University)
	if err != nil {
		return nil, err
	}
	return &fullWorkUser, nil
}

func (r *WorkRepository) GetByIdWorkUserCVPath(ctx context.Context, id int) (string, error) {
	var cvPath string
	query := `SELECT cv_path FROM work_user WHERE id = $1`
	err := r.db.QueryRowContext(ctx, query, id).Scan(&cvPath)
	if err != nil {
		return "", err
	}
	return cvPath, nil
}

func (r *WorkRepository) GetByEmailHR(ctx context.Context, email string) (*entities.HR, error) {
	var hr entities.HR
	query := "SELECT id, email, password, last_name, first_name, father_name, company, image_link FROM hr WHERE email = $1"
	err := r.db.QueryRowContext(ctx, query, email).Scan(&hr.Id, &hr.Email, &hr.Password, &hr.LastName, &hr.FirstName, &hr.FatherName, &hr.Company, &hr.ImageLink)
	if err != nil {
		return nil, err
	}
	return &hr, nil
}

func (r *WorkRepository) UpdateWorkUser(ctx context.Context, workUser *entities.WorkUserUpdateRequest) error {
	err := r.Create(ctx, &entities.WorkUser{Id: workUser.Id, Speciality: workUser.Speciality, WorkExperience: workUser.WorkExperience, AdditionalExperience: workUser.AdditionalExperience, UsefulLinks: workUser.UsefulLinks, PhoneNumber: workUser.PhoneNumber, Telegram: workUser.Telegram, Skills: workUser.Skills})
	if err != nil {
		return err
	}

	query := `UPDATE work_user SET phone_number = $1, telegram = $2, skills = $3, speciality = $4, work_experience = $5, additional_experience = $6, useful_links = $7 WHERE id = $8`
	_, err = r.db.ExecContext(ctx, query, workUser.PhoneNumber, workUser.Telegram, pq.Array(workUser.Skills), workUser.Speciality, workUser.WorkExperience, workUser.AdditionalExperience, pq.Array(workUser.UsefulLinks), workUser.Id)
	if err != nil {
		return err
	}
	return nil
}

func (r *WorkRepository) UpdateCVPath(ctx context.Context, workUser *entities.UpdateCV) error {
	query := `UPDATE work_user SET cv_path = $1 WHERE id = $2`
	_, err := r.db.ExecContext(ctx, query, workUser.CVPath, workUser.Id)
	if err != nil {
		return err
	}
	return nil
}

func (r *WorkRepository) ExistsResponse(ctx context.Context, response *entities.Response) (bool, error) {
	var exists int
	query := `SELECT 1 FROM response WHERE hr_id = $1 AND response.work_user_id = $2`
	err := r.db.QueryRowContext(ctx, query, response.HRId, response.WorkUserId).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists > 0, nil
}

func (r *WorkRepository) CreateResponse(ctx context.Context, response *entities.Response) error {
	if check, _ := r.ExistsResponse(ctx, response); check != true {
		query := `INSERT INTO response (hr_id, work_user_id) VALUES ($1, $2)`
		_, err := r.db.ExecContext(ctx, query, response.HRId, response.WorkUserId)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *WorkRepository) GetWorkUserResponses(ctx context.Context, workUserId int) (*[]entities.Response, error) {
	var responses []entities.Response
	query := `SELECT id, hr_id, work_user_id FROM response WHERE work_user_id = $1`
	rows, err := r.db.QueryContext(ctx, query, workUserId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var response entities.Response
		err = rows.Scan(&response.Id, &response.HRId, &response.WorkUserId)
		if err != nil {
			return nil, err
		}
		responses = append(responses, response)
	}

	return &responses, nil
}

func (r *WorkRepository) GetHRResponses(ctx context.Context, hrId int) (*[]entities.CandidateResponse, error) {
	var responses []entities.CandidateResponse
	query := `
        SELECT r.id, r.hr_id, u.id, u.last_name, u.first_name, u.father_name, w.speciality 
        FROM response AS r
        JOIN work_user AS w ON r.work_user_id = w.id
        JOIN users AS u ON w.id = u.id
        WHERE r.hr_id = $1
    `
	rows, err := r.db.QueryContext(ctx, query, hrId)
	if err != nil {
		return nil, err
	}
	defer rows.Close() // Закрываем rows после использования

	for rows.Next() {
		var response entities.CandidateResponse

		err = rows.Scan(&response.Id, &response.HRId, &response.WorkUserId, &response.LastName, &response.FirstName, &response.FatherName, &response.Speciality)
		if err != nil {
			return nil, err
		}

		responses = append(responses, response)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &responses, nil
}

func (r *WorkRepository) GetAllWorkUser(ctx context.Context) (*[]entities.FullWorkUser, error) {
	var fullWorkUsers []entities.FullWorkUser
	query := `SELECT work_user.id, speciality, work_experience, additional_experience, useful_links, phone_number, telegram, skills, last_name, first_name, father_name, university.name, gpa.value FROM work_user JOIN users ON work_user.id = users.id JOIN university ON users.university_id = university.id JOIN gpa ON users.id = gpa.user_id;`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var fullWorkUser entities.FullWorkUser
		err = rows.Scan(&fullWorkUser.Id, &fullWorkUser.Speciality, &fullWorkUser.WorkExperience, &fullWorkUser.AdditionalExperience, pq.Array(&fullWorkUser.UsefulLinks), &fullWorkUser.PhoneNumber, &fullWorkUser.Telegram, pq.Array(&fullWorkUser.Skills), &fullWorkUser.LastName, &fullWorkUser.FirstName, &fullWorkUser.FatherName, &fullWorkUser.University, &fullWorkUser.Gpa)
		if err != nil {
			return nil, err
		}
		fullWorkUsers = append(fullWorkUsers, fullWorkUser)
	}

	return &fullWorkUsers, nil
}
