import type { IDataItem, IGradeClassItem, IGroupMemberItem, IUsersScoreItem, IReaorganizedGroupScore } from "./constants";

// дополняет пробелы в расписании
export function transformData(data: IDataItem): IDataItem {
  const { grade_class, group_member, users_score } = data;
  const newGradeClass: IGradeClassItem[] = [];
  const newUsersScore: IUsersScoreItem[] = [];

  // Create a map for faster lookup of grades by user ID
  const gradesByUser: { [userId: number]: { class_id: number; id: number; user_id: number; value: number }[] } = {};
  grade_class.forEach(item => {
    item.grades.forEach(grade => {
      gradesByUser[grade.user_id] = gradesByUser[grade.user_id] || [];
      gradesByUser[grade.user_id].push(grade);
    });
  });

  //Iterate over group_members and generate grades and users_score
  group_member.forEach(member => {
    const grades = gradesByUser[member.id] || [{ class_id: 0, id: 0, user_id: member.id, value: 0 }]; 
    newGradeClass.push({
      ...grade_class[0],
      grades: grades,
    });

    // Add user score (or create default if missing)
    const existingScore = users_score.find(score => score.user_id === member.id);
    newUsersScore.push(existingScore || { average_score: 0, sum_score: 0, user_id: member.id });
  });

  return {
    grade_class: newGradeClass,
    group_member: group_member,
    users_score: newUsersScore,
  };
}


// конвертирует формат расписания в удобный для отрисовки таблицы
export function reorganizeGroupScores(data: IDataItem): IReaorganizedGroupScore[] {
  const { grade_class, group_member, users_score } = data;

  // Create a map for faster lookup of group members by ID
  const groupMemberMap = new Map<number, IGroupMemberItem>();
  group_member.forEach(member => groupMemberMap.set(member.id, member));

  // Create a map for faster lookup of user scores by ID
  const userScoreMap = new Map<number, IUsersScoreItem>();
  users_score.forEach(score => userScoreMap.set(score.user_id, score));


  const reorganizedScores: IReaorganizedGroupScore[] = [];

  // Iterate through group members to create reorganized scores
  group_member.forEach(member => {
    const grades: {
      date: string;
      class_id: number;
      id: number;
      user_id: number;
      value: number;
    }[] = [];

    // Find and add matching grades for each member
    grade_class.forEach(gradeClass => {
      gradeClass.grades.forEach(grade => {
        if (grade.user_id === member.id) {
          grades.push({ ...grade, date: gradeClass.date }); //Include date from gradeClass
        }
      });
    });

    // Find and add user score for the member
    const userScore = userScoreMap.get(member.id);
    reorganizedScores.push({
      first_name: member.first_name,
      last_name: member.last_name,
      father_name: member.father_name,
      id: member.id,
      grades,
      average_score: userScore ? userScore.average_score : 0, //Handle case where userScore might be undefined
      sum_score: userScore ? userScore.sum_score : 0,       //Handle case where userScore might be undefined
      user_id: member.id,
      gpa: 0,
    });
  });

  return reorganizedScores;
}