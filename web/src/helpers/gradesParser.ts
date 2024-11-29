import type { IDataItem, IGradeClassItem, IGroupMemberItem, IUsersScoreItem, IReaorganizedGroupScore } from "./constants";

// дополняет пробелы в расписании
export function transformData(data: IDataItem): IDataItem {
  const { grade_class, group_member, users_score } = data;
  const numMembers = group_member.length;

  // Create a map of existing grades for quick access.
  const existingGrades: { [userId: number]: { class_id: number; id: number; user_id: number; value: number }[] } = {};
  grade_class.forEach(gc => {
    gc.grades.forEach(grade => {
      existingGrades[grade.user_id] = existingGrades[grade.user_id] || [];
      existingGrades[grade.user_id].push(grade);
    });
  });

  // Create a new grade_class array with all grades filled in.
  const newGradeClass: IGradeClassItem[] = [];
  grade_class.forEach(gradeClass => {
    const newGrades: { class_id: number; id: number; user_id: number; value: number }[] = [];
    //Add existing grades for this gradeClass
    gradeClass.grades.forEach(grade => newGrades.push(grade));
    //Fill in missing grades with default values.
    for (let i = 0; i < numMembers; i++) {
      const userId = group_member[i].id;
      if (!newGrades.find(g => g.user_id === userId)) {
        newGrades.push({ class_id: 0, id: 0, user_id: userId, value: 0 });
      }
    }
    newGradeClass.push({ ...gradeClass, grades: newGrades });
  });

  const newScoresClass: IUsersScoreItem[] = [];
  for(let groupMember of group_member){
    const newScore = { average_score: 0, sum_score: 0, user_id: 0 };
    for(let item of users_score){
      if(item.user_id === groupMember.id){
        newScore.average_score = item.average_score;
        newScore.sum_score = item.sum_score;
        newScore.user_id = item.user_id;
        break;
      }
    }
    newScoresClass.push(newScore);
  }

  return {
    grade_class: newGradeClass,
    group_member: group_member,
    users_score: newScoresClass,
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