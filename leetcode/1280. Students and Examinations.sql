SELECT st.student_id,
       st.student_name,
       su.subject_name,
       COUNT(e.student_id) as attended_exams
FROM Students st
         CROSS JOIN Subjects su
         LEFT JOIN Examinations e ON st.student_id = e.student_id AND su.subject_name = e.subject_name
GROUP BY st.student_id, st.student_name, su.subject_name
ORDER BY st.student_id, su.subject_name;