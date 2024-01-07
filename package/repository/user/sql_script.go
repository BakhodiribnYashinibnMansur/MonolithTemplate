package user

const (
	CreateUserQuery                 = `INSERT INTO crm_user (full_name,birthday_date,	added_date,phone_number,role_id,password,photo) VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING id`
	UpdateUserQuery                 = `UPDATE crm_user SET full_name=$1,	birthday_date=$2 ,added_date=$3,phone_number=$4,role_id=$5,photo=$6,	updated_at = NOW() WHERE id=$7`
	UpdateUserPasswordQuery         = `UPDATE crm_user SET password=$1,updated_at = NOW() WHERE id=$2 AND deleted_at IS NULL`
	DeleteUserQuery                 = `UPDATE crm_user SET deleted_at = NOW() ,updated_at = NOW(),phone_number=concat(phone_number,' ',uuid_generate_v4()) WHERE id=$1 AND deleted_at IS NULL`
	GetUserByIDQuery                = `SELECT id,full_name,phone_number,birthday_date,role_id,added_date,photo FROM crm_user WHERE id=$1 AND deleted_at IS NULL ORDER BY created_at DESC `
	CheckUserByIDQuery              = `SELECT id FROM crm_user WHERE id=$1 AND deleted_at IS NULL`
	GetUserListByRoleQuery          = `SELECT id,full_name,phone_number,birthday_date,role_id,added_date ,photo FROM crm_user WHERE deleted_at IS	NULL AND role_id=$1  ORDER BY created_at DESC  LIMIT $2 OFFSET $3`
	GetUserListAllQuery             = `SELECT id,full_name,phone_number,birthday_date,role_id,photo,	added_date  FROM crm_user WHERE deleted_at IS NULL  ORDER BY created_at DESC  LIMIT $1	OFFSET $2`
	SignInUserQuery                 = `SELECT id,	role_id FROM crm_user WHERE  phone_number=$1  AND password=$2 AND	deleted_at IS NULL`
	GetGroupStudentListQuery        = `SELECT id,full_name,phone_number,birthday_date,role_id,added_date,photo FROM crm_user WHERE id= ANY($1) AND deleted_at IS NULL `
	GetUserCountByRoleIDQuery       = `SELECT COUNT(id) FROM crm_user WHERE deleted_at IS NULL AND role_id=$1`
	GetUserCountQuery               = `SELECT COUNT(id) FROM crm_user WHERE deleted_at IS NULL `
	CheckUserByPhoneQuery           = `SELECT id FROM crm_user WHERE phone_number=$1 AND deleted_at IS NULL`
	GetStudentFullNameByIdQuery     = `SELECT full_name FROM crm_user WHERE id=$1 AND deleted_at IS NULL `
	GetUserDataListByUserIDsList    = `SELECT id , full_name FROM crm_user WHERE id=ANY($1) AND deleted_at IS NULL `
	GetRoleUserRelationalCountQuery = `SELECT COUNT(id) FROM crm_user WHERE deleted_at IS NULL AND role_id = $1`
	GetUsersWithSearchKey           = `SELECT id, full_name, role_id FROM crm_user WHERE deleted_at IS NULL `
)
const (
	CreateManagerQuery         = `INSERT INTO crm_user (full_name,birthday_date,added_date,phone_number,role_id,password,photo) VALUES ($1,$2,$3,$4,	$5,$6,$7) RETURNING id`
	UpdateManagerQuery         = `UPDATE crm_user SET full_name=$1,	birthday_date=$2 ,added_date=$3,phone_number=$4,role_id=$5,photo=$6,	updated_at = NOW() WHERE id=$7`
	UpdateManagerPasswordQuery = `UPDATE crm_user SET password=$1,updated_at = NOW() WHERE id=$2 AND deleted_at IS NULL`
	DeleteManagerQuery         = `UPDATE crm_user SET deleted_at = NOW() ,updated_at = NOW(),phone_number=concat(phone_number,' ',uuid_generate_v4()) WHERE id=$1 AND deleted_at IS NULL`
	GetManagerByIDQuery        = `SELECT id,full_name,phone_number,birthday_date,role_id,added_date,photo FROM crm_user WHERE id=$1 AND deleted_at IS NULL ORDER BY created_at DESC `
	CheckManagerByIDQuery      = `SELECT id FROM crm_user WHERE id=$1 AND deleted_at IS NULL`
	GetManagerListByRoleQuery  = `SELECT id,full_name,phone_number,birthday_date,role_id,added_date ,photo FROM crm_user WHERE deleted_at IS	NULL AND role_id=$1  ORDER BY created_at DESC  LIMIT $2 OFFSET $3`
	GetManagerCountQuery       = `SELECT COUNT(id) FROM crm_user WHERE deleted_at IS NULL  AND role_id=$1`
	CheckManagerByPhoneQuery   = `SELECT id FROM crm_user WHERE phone_number=$1 AND deleted_at IS NULL`
)
const (
	CreateTeacherQuery           = `INSERT INTO crm_user (full_name,birthday_date,added_date,phone_number,role_id,password,photo) VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING id`
	UpdateTeacherQuery           = `UPDATE crm_user SET full_name=$1,	birthday_date=$2 ,added_date=$3,phone_number=$4,role_id=$5,photo=$6,	updated_at = NOW() WHERE id=$7`
	UpdateTeacherPasswordQuery   = `UPDATE crm_user SET password=$1,updated_at = NOW() WHERE id=$2 AND deleted_at IS NULL`
	DeleteTeacherQuery           = `UPDATE crm_user SET deleted_at = NOW() ,updated_at = NOW(),phone_number=concat(phone_number,' ',uuid_generate_v4()) WHERE id=$1 AND deleted_at IS NULL`
	GetTeacherByIDQuery          = `SELECT id,full_name,phone_number,birthday_date,role_id,added_date,photo FROM crm_user WHERE id=$1 AND deleted_at IS NULL ORDER BY created_at DESC `
	CheckTeacherByIDQuery        = `SELECT id FROM crm_user WHERE id=$1 AND deleted_at IS NULL`
	GetTeacherListByRoleQuery    = `SELECT id,full_name,phone_number,birthday_date,role_id,added_date ,photo FROM crm_user WHERE deleted_at IS	NULL AND role_id=$1  ORDER BY created_at DESC  LIMIT $2 OFFSET $3`
	GetTeacherStatisticListQuery = `SELECT id, full_name, photo  FROM crm_user WHERE deleted_at IS NULL AND role_id = $1 LIMIT 10 OFFSET 0`
	GetTeacherCountQuery         = `SELECT COUNT(id) FROM crm_user WHERE deleted_at IS NULL  AND role_id=$1`
	CheckTeacherByPhoneQuery     = `SELECT id FROM crm_user WHERE phone_number=$1 AND deleted_at IS NULL`
	GetFilterTeacherCountQuery   = `SELECT COUNT(id) as total FROM crm_user WHERE deleted_at IS NULL AND role_id=$1 `
	GetFilterTeacherListQuery    = `SELECT id,full_name FROM crm_user WHERE deleted_at IS NULL AND role_id=$1 `
)
const (
	CreateStudentQuery             = `INSERT INTO crm_user (full_name,birthday_date,	added_date,phone_number,role_id,password,photo) VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING id`
	UpdateStudentQuery             = `UPDATE crm_user SET full_name=$1,	birthday_date=$2 ,added_date=$3,phone_number=$4,role_id=$5,photo=$6,	updated_at = NOW() WHERE id=$7`
	UpdateStudentPasswordQuery     = `UPDATE crm_user SET password=$1,updated_at = NOW() WHERE id=$2 AND deleted_at IS NULL`
	DeleteStudentQuery             = `UPDATE crm_user SET deleted_at = NOW() ,updated_at = NOW(),phone_number=concat(phone_number,' ',uuid_generate_v4()) WHERE id=$1 AND deleted_at IS NULL`
	GetStudentByIDQuery            = `SELECT DISTINCT crm_user.id,crm_user.full_name,crm_user.phone_number,crm_user.birthday_date,crm_user.role_id,crm_user.added_date,crm_user.photo,student_data.student_tag, student_data.student_status ,  COALESCE(student_account.balance,0) as amount FROM crm_user LEFT JOIN student_data ON crm_user.id = student_data.student_id LEFT JOIN student_account ON crm_user.id = student_account.student_id  WHERE crm_user.deleted_at IS NULL AND student_data.deleted_at IS NULL AND crm_user.id =$1  and student_account.deleted_at is null `
	CheckStudentByIDQuery          = `SELECT id FROM crm_user WHERE id=$1 AND deleted_at IS NULL`
	GetStudentIDByNameOrPhoneQuery = `SELECT id FROM crm_user WHERE ( full_name ILIKE '%'||$1||'%' OR phone_number ILIKE '%'||$1||'%') AND deleted_at IS NULL`
	GetStudentListQuery            = `SELECT  crm_user.id,crm_user.full_name,crm_user.phone_number,crm_user.birthday_date,crm_user.role_id,crm_user.added_date,crm_user.photo,student_data.student_tag, student_data.student_status, COALESCE(student_account.balance,0) as amount  FROM crm_user LEFT JOIN student_data ON crm_user.id = student_data.student_id LEFT JOIN student_account ON crm_user.id = student_account.student_id WHERE crm_user.deleted_at IS NULL AND student_data.deleted_at IS NULL AND crm_user.role_id=$1  and student_account.deleted_at is null `
	GetStudentListCountQuery       = `SELECT COUNT(crm_user.id) FROM crm_user LEFT JOIN student_data ON crm_user.id = student_data.student_id WHERE crm_user.deleted_at IS NULL AND student_data.deleted_at IS NULL AND crm_user.role_id=$1  `
	GetStudentCountQuery           = `SELECT COUNT(id) FROM crm_user WHERE deleted_at IS NULL   AND role_id=$1 `
	CheckStudentByPhoneQuery       = `SELECT id FROM crm_user WHERE phone_number=$1 AND deleted_at IS NULL`
	GetSearchStudentListQuery      = `SELECT id, full_name, phone_number FROM crm_user WHERE deleted_at IS NULL AND role_id=$1 `
	GetSearchStudentCountQuery     = `SELECT COUNT(id) as total FROM crm_user WHERE deleted_at IS NULL AND role_id=$1 `
	GetStudentCountLineChartQuery  = `SELECT all_months.month, all_months.year, COALESCE((SELECT COUNT(id) FROM crm_user WHERE role_id = $1 AND deleted_at IS NULL AND  EXTRACT(MONTH FROM created_at) = all_months.month and EXTRACT(year FROM created_at) = all_months.year  ), 0) AS student_count FROM (SELECT extract(month from generate_series(current_date-$2::interval, current_date,'1 month')) AS month, extract(year from generate_series(current_date-$2::interval, current_date,'1 month')) AS year) AS all_months  ORDER BY all_months.year;`
	GetStudentFullDataListQuery    = `SELECT DISTINCT crm_user.id,crm_user.full_name,crm_user.phone_number,crm_user.birthday_date,crm_user.role_id,crm_user.added_date,crm_user.photo,student_data.student_tag, student_data.student_status FROM crm_user INNER JOIN student_data ON crm_user.id = student_data.student_id WHERE crm_user.deleted_at IS NULL AND student_data.deleted_at IS NULL `
)
const (
	CreateStudentDataQuery  = `INSERT INTO student_data (student_id,cellular_phone,student_given_id,parents_phone,email,telegram_nick_name,location,passport,student_tag,discount, student_status) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11) RETURNING id`
	UpdateStudentDataQuery  = `UPDATE student_data SET cellular_phone=$1,student_given_id=$2,parents_phone=$3,email=$4,telegram_nick_name=$5,location=$6,passport=$7,student_tag=$8,discount=$9,updated_at = NOW() WHERE student_id=$10 AND deleted_at IS NULL`
	DeleteStudentDataQuery  = `UPDATE student_data SET deleted_at = NOW() ,updated_at = NOW() WHERE student_id=$1 AND deleted_at IS NULL`
	GetStudentDataByIDQuery = `SELECT id,cellular_phone,student_given_id,parents_phone,email,telegram_nick_name,location,passport,student_tag,discount, student_status FROM student_data WHERE student_id=$1 AND deleted_at IS NULL `
)
const (
	CreateEmployeeQuery           = `INSERT INTO crm_user (full_name,birthday_date,added_date,phone_number,role_id,password,photo) VALUES ($1,$2,$3,$4,	$5,$6,$7) RETURNING id`
	UpdateEmployeeQuery           = `UPDATE crm_user SET full_name=$1,	birthday_date=$2 ,added_date=$3,phone_number=$4,role_id=$5,photo=$6,	updated_at = NOW() WHERE id=$7`
	UpdateEmployeePasswordQuery   = `UPDATE crm_user SET password=$1,updated_at = NOW() WHERE id=$2 AND deleted_at IS NULL`
	DeleteEmployeeQuery           = `UPDATE crm_user SET deleted_at = NOW() ,updated_at = NOW(),phone_number=concat(phone_number,' ',uuid_generate_v4()) WHERE id=$1 AND deleted_at IS NULL`
	GetEmployeeByIDQuery          = `SELECT id,full_name,phone_number,birthday_date,role_id,added_date,photo FROM crm_user WHERE id=$1 AND role_id <> ANY($2) AND deleted_at IS NULL ORDER BY created_at DESC `
	CheckEmployeeByIDQuery        = `SELECT id FROM crm_user WHERE id=$1 AND deleted_at IS NULL`
	GetEmployeeListByRoleQuery    = `SELECT id,full_name,phone_number,birthday_date,role_id,added_date ,photo FROM crm_user WHERE deleted_at IS	NULL AND role_id <> ANY($1)  ORDER BY created_at DESC  LIMIT $2 OFFSET $3`
	GetEmployeeCountByRoleQuery   = `SELECT COUNT(id) FROM crm_user WHERE deleted_at IS NULL AND  role_id <> ANY($1)`
	CheckEmployeeByPhoneQuery     = `SELECT id FROM crm_user WHERE phone_number=$1 AND deleted_at IS NULL`
	GetDropDownEmployeeListQuery  = `SELECT id , full_name, role_id FROM crm_user WHERE deleted_at IS	NULL AND role_id <> ANY($1)  `
	GetDropDownEmployeeCountQuery = `SELECT  COUNT(id) as total FROM crm_user WHERE deleted_at IS	NULL AND role_id <> ANY($1)  `
	GetUserDropDownListQuery      = `SELECT id , full_name, role_id,photo FROM crm_user WHERE deleted_at IS NULL `
	GetEmployeeCountQuery         = `SELECT COUNT(id) FROM crm_user WHERE deleted_at IS NULL AND role_id <> ANY($1)`
)
