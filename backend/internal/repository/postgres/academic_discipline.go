package postgres

import (
	"context"
	"database/sql"
	"errors"
	"server/internal/entities"
)

type AcademicDisciplineRepository struct {
	db *sql.DB
}

func NewAcademicDisciplineRepository(db *sql.DB) *AcademicDisciplineRepository {
	return &AcademicDisciplineRepository{
		db: db,
	}
}

func (r *AcademicDisciplineRepository) Exists(ctx context.Context, academicDiscipline *entities.AcademicDiscipline) (bool, error) {
	var exists int
	query := `SELECT 1 FROM academic_discipline WHERE name = $1 AND semester = $2 AND educational_direction = $3`
	err := r.db.QueryRowContext(ctx, query, academicDiscipline.Name, academicDiscipline.Semester, academicDiscipline.EducationalDirection).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists > 0, nil
}

func (r *AcademicDisciplineRepository) Create(ctx context.Context, academicDiscipline *entities.AcademicDiscipline) (int, error) {
	if academicDiscipline.Name == "" {
		return 0, errors.New("empty academicDiscipline name ")
	}

	if check, _ := r.Exists(ctx, academicDiscipline); check == true {
		return 0, errors.New("academicDiscipline already exists")
	}

	var id int
	query := `INSERT INTO academic_discipline (name, educational_direction, semester, lecture_hours, practice_hours, lab_hours, individual_hours, type_of_assessment) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`

	err := r.db.QueryRowContext(ctx, query, academicDiscipline.Name, academicDiscipline.EducationalDirection, academicDiscipline.Semester, academicDiscipline.LectureHours, academicDiscipline.PracticeHours, academicDiscipline.LabHours, academicDiscipline.IndividualHours, academicDiscipline.TypeOfAssessment).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AcademicDisciplineRepository) GetByEducationalDirectionAndSemester(ctx context.Context, educationalDirection string, semester int) (*[]entities.AcademicDiscipline, error) {
	if educationalDirection == "" {
		return nil, errors.New("educationalDirection is empty")
	}

	var academicDisciplines []entities.AcademicDiscipline

	query := `SELECT * FROM academic_discipline WHERE educational_direction=$1 AND semester=$2`
	rows, err := r.db.QueryContext(ctx, query, educationalDirection, semester)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var academicDiscipline entities.AcademicDiscipline
		err = rows.Scan(&academicDiscipline.Id, &academicDiscipline.Name, &academicDiscipline.EducationalDirection, &academicDiscipline.Semester, &academicDiscipline.LectureHours, &academicDiscipline.PracticeHours, &academicDiscipline.LabHours, &academicDiscipline.IndividualHours, &academicDiscipline.TypeOfAssessment)
		if err != nil {
			return nil, err
		}
		academicDisciplines = append(academicDisciplines, academicDiscipline)
	}

	if len(academicDisciplines) == 0 {
		return nil, errors.New("there are no academicDisciplines with such parameters")
	}

	return &academicDisciplines, nil
}

func (r *AcademicDisciplineRepository) GetByEducationalDirectionAndName(ctx context.Context, name, educationalDirection string) (*entities.AcademicDiscipline, error) {
	if educationalDirection == "" {
		return nil, errors.New("educationalDirection is empty")
	}

	var academicDiscipline entities.AcademicDiscipline
	query := `SELECT * FROM academic_discipline WHERE educational_direction=$1 AND name=$2`
	err := r.db.QueryRowContext(ctx, query, educationalDirection, name).Scan(&academicDiscipline.Id, &academicDiscipline.Name, &academicDiscipline.EducationalDirection, &academicDiscipline.Semester, &academicDiscipline.LectureHours, &academicDiscipline.PracticeHours, &academicDiscipline.LabHours, &academicDiscipline.IndividualHours, &academicDiscipline.TypeOfAssessment)
	if err != nil {
		return nil, err
	}
	return &academicDiscipline, nil
}

func (r *AcademicDisciplineRepository) Update(ctx context.Context, academicDiscipline *entities.AcademicDiscipline) error {
	if academicDiscipline.EducationalDirection == "" {
		return errors.New("educationalDirection is empty")
	}

	query := `UPDATE academic_discipline SET name=$1, educational_direction=$2, semester=$3, lecture_hours=$4, practice_hours=$5, lab_hours=$6, individual_hours=$7, type_of_assessment=$8 WHERE id=$9`

	_, err := r.db.ExecContext(ctx, query, academicDiscipline.Name, academicDiscipline.EducationalDirection, academicDiscipline.Semester, academicDiscipline.LectureHours, academicDiscipline.PracticeHours, academicDiscipline.LabHours, academicDiscipline.IndividualHours, academicDiscipline.TypeOfAssessment, academicDiscipline.Id)
	if err != nil {
		return err
	}

	return nil
}

func (r *AcademicDisciplineRepository) Delete(ctx context.Context, id int) error {
	if id == 0 {
		return errors.New("")
	}

	query := `DELETE FROM academic_discipline WHERE id = $1`

	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}
