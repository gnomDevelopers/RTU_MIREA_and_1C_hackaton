from collections import defaultdict

# Данные для генерации расписания
days = ["Monday", "Tuesday", "Wednesday", "Thursday", "Friday"]
time_slots = ["9:00-10:30", "10:45-12:15", "12:45-14:15", "14:30-16:00", "16:15-17:45"]
rooms = ["101", "102"]
max_capacity = {"101": 30, "102": 50}

# Пример данных
group_subjects = {
    "Group A": ["a1"],
    "Group B": ["a1"]
}

groups = {group: size for group, size in zip(group_subjects.keys(), [25, 40, 20, 50, 40, 20, 50])}

teachers = {
    "Tutor": ["a1"],
}

# Максимальное количество пар в неделю для каждого преподавателя
teacher_max_classes = {
    "Tutor": 10,
}

# Генератор расписания
from itertools import cycle


def generate_schedule():
    schedule = defaultdict(list)
    used_slots = defaultdict(lambda: defaultdict(list))  # Учет занятых слотов
    teacher_classes_count = defaultdict(int)  # Счетчик пар для преподавателей
    group_daily_count = defaultdict(lambda: defaultdict(int))  # Количество пар у группы в день
    teacher_daily_count = defaultdict(lambda: defaultdict(int))  # Количество пар у преподавателя в день

    # Распределение дней по кругу для равномерного распределения пар
    day_cycle = cycle(days)

    for group, group_size in groups.items():
        for subject in group_subjects[group]:
            assigned = False

            # Найти подходящего преподавателя
            suitable_teachers = [t for t, t_subjects in teachers.items() if subject in t_subjects]

            for teacher in suitable_teachers:
                if teacher_classes_count[teacher] >= teacher_max_classes[teacher]:
                    continue  # Преподаватель превысил лимит пар

                # Попытаться назначить пару, проходя дни в цикле
                for _ in range(len(days)):
                    day = next(day_cycle)
                    if group_daily_count[group][day] >= 3:  # Ограничение: максимум 2 пары в день для группы
                        continue
                    if teacher_daily_count[teacher][day] >= 2:  # Ограничение: максимум 2 пары в день для преподавателя
                        continue

                    for slot in time_slots:
                        for room in rooms:
                            # Проверить ограничения
                            if room in used_slots[day][slot]:
                                continue
                            if group in used_slots[day][slot]:
                                continue
                            if teacher in used_slots[day][slot]:
                                continue
                            if max_capacity[room] < group_size:
                                continue

                            # Назначить пару
                            schedule[group].append({
                                "subject": subject,
                                "teacher": teacher,
                                "day": day,
                                "time": slot,
                                "room": room,
                            })
                            used_slots[day][slot].extend([room, group, teacher])
                            teacher_classes_count[teacher] += 1
                            group_daily_count[group][day] += 1
                            teacher_daily_count[teacher][day] += 1
                            assigned = True
                            break
                        if assigned:
                            break
                    if assigned:
                        break

    return schedule


# def display_schedule(schedule):
#     day_order = {day: index for index, day in enumerate(["Monday", "Tuesday", "Wednesday", "Thursday", "Friday"])}

#     for group, lessons in schedule.items():
#         print(f"Schedule for {group}:")

#         sorted_lessons = sorted(
#             lessons,
#             key=lambda lesson: (day_order[lesson['day']], lesson['time'])
#         )

#         for lesson in sorted_lessons:
#             print(
#                 f"  {lesson['day']} {lesson['time']} - {lesson['subject']} "
#                 f"({lesson['teacher']}) in room {lesson['room']}"
#             )
#         print()


schedule = generate_schedule()
# display_schedule(schedule)
