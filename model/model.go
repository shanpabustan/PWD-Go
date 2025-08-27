// package model

// import (
// 	"time"

// 	"github.com/google/uuid"
// )

// type Account struct {
// 	ID                     uuid.UUID  `gorm:"type:uuid;primaryKey" json:"id"`
// 	ApplicantType          string     `gorm:"type:varchar(10);check:applicant_type IN ('New','Renewal')" json:"applicant_type"`
// 	DisabilityNumber       *string    `json:"disability_number,omitempty"`
// 	CreatedAt              time.Time  `json:"created_at"`
// 	LastName               string     `json:"last_name"`
// 	FirstName              string     `json:"first_name"`
// 	MiddleName             string     `json:"middle_name"`
// 	Suffix                 *string    `json:"suffix,omitempty"`
// 	DateOfBirth            time.Time  `json:"date_of_birth"`
// 	Gender                 string     `json:"gender"`
// 	CivilStatus            string     `json:"civil_status"`
// 	FatherLastName         string     `json:"father_last_name"`
// 	FatherFirstName        string     `json:"father_first_name"`
// 	FatherMiddleName       string     `json:"father_middle_name"`
// 	MotherLastName         string     `json:"mother_last_name"`
// 	MotherFirstName        string     `json:"mother_first_name"`
// 	MotherMiddleName       string     `json:"mother_middle_name"`
// 	GuardianLastName       string     `json:"guardian_last_name"`
// 	GuardianFirstName      string     `json:"guardian_first_name"`
// 	GuardianMiddleName     string     `json:"guardian_middle_name"`
// 	TypeOfDisability       string     `json:"type_of_disability"`
// 	CauseOfDisability      string     `json:"cause_of_disability"`
// 	HouseNoAndStreet       string     `json:"house_no_and_street"`
// 	Barangay               string     `json:"barangay"`
// 	Municipality           string     `json:"municipality"`
// 	Province               string     `json:"province"`
// 	Region                 string     `json:"region"`
// 	LandlineNo             *string    `json:"landline_no,omitempty"`
// 	MobileNo               string     `json:"mobile_no"`
// 	EmailAddress           string     `json:"email_address"`
// 	EducationalAttainment  string     `json:"educational_attainment"`
// 	StatusOfEmployment     *string    `json:"status_of_employment,omitempty"`
// 	CategoryOfEmployment   *string    `json:"category_of_employment,omitempty"`
// 	TypeOfEmployment       *string    `json:"type_of_employment,omitempty"`
// 	Occupation             *string    `json:"occupation,omitempty"`
// 	OtherOccupation        *string    `json:"other_occupation,omitempty"`
// 	OrganizationAffiliated *string    `json:"organization_affiliated,omitempty"`
// 	ContactPerson          *string    `json:"contact_person,omitempty"`
// 	OfficeAddress          *string    `json:"office_address,omitempty"`
// 	OfficeTelNo            *string    `json:"office_tel_no,omitempty"`
// 	SSSNo                  *string    `json:"sss_no,omitempty"`
// 	GSISNo                 *string    `json:"gsis_no,omitempty"`
// 	PagIBIGNo              *string    `json:"pagibig_no,omitempty"`
// 	PSNNo                  *string    `json:"psn_no,omitempty"`
// 	PhilHealthNo           *string    `json:"philhealth_no,omitempty"`
// 	AccomplishByLastName   string     `json:"accomplish_by_last_name"`
// 	AccomplishByFirstName  string     `json:"accomplish_by_first_name"`
// 	AccomplishByMiddleName string     `json:"accomplish_by_middle_name"`
// 	ReferenceNumber        string     `json:"reference_number"`
// 	ValidUntil             *time.Time `json:"valid_until,omitempty"`
// 	DisapprovalDate        *time.Time `json:"disapproval_date,omitempty"`

// 	// ✅ One-to-One Relationship
// 	ApprovalStatus ApprovalStatus `gorm:"foreignKey:AccountID" json:"approval_status"`
// }

// // Helper function
// func (a *Account) IsExpired() bool {
// 	return a.ValidUntil != nil && a.ValidUntil.Before(time.Now())
// }

// type ApprovalStatus struct {
// 	ID           uint      `gorm:"primaryKey" json:"id"`
// 	AccountID    uuid.UUID `gorm:"type:uuid;not null;unique" json:"account_id"`
// 	Status       string    `gorm:"type:varchar(20);default:'Pending'" json:"status"`
// 	Requirement1 bool      `json:"requirement_1"`
// 	Requirement2 bool      `json:"requirement_2"`
// 	Requirement3 bool      `json:"requirement_3"`
// 	Requirement4 bool      `json:"requirement_4"`
// 	IsApproved   bool      `json:"is_approved"`

// 	// 🔁 Back-reference to Account
// 	Account *Account `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
// }




package model

import (
	"time"

	"github.com/google/uuid"
)

//
// =========================
// PWD User Account
// =========================
//
type Account struct {
	ID                     uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	ApplicantType          string     `gorm:"type:varchar(10);check:applicant_type IN ('New','Renewal')" json:"applicant_type"`
	DisabilityNumber       *string    `json:"disability_number,omitempty"`
	CreatedAt              time.Time  `gorm:"autoCreateTime" json:"created_at"`
	LastName               string     `json:"last_name"`
	FirstName              string     `json:"first_name"`
	MiddleName             string     `json:"middle_name"`
	Suffix                 *string    `json:"suffix,omitempty"`
	DateOfBirth            time.Time  `json:"date_of_birth"`
	Gender                 string     `json:"gender"`
	CivilStatus            string     `json:"civil_status"`
	FatherLastName         string     `json:"father_last_name"`
	FatherFirstName        string     `json:"father_first_name"`
	FatherMiddleName       string     `json:"father_middle_name"`
	MotherLastName         string     `json:"mother_last_name"`
	MotherFirstName        string     `json:"mother_first_name"`
	MotherMiddleName       string     `json:"mother_middle_name"`
	GuardianLastName       string     `json:"guardian_last_name"`
	GuardianFirstName      string     `json:"guardian_first_name"`
	GuardianMiddleName     string     `json:"guardian_middle_name"`
	TypeOfDisability       string     `json:"type_of_disability"`
	CauseOfDisability      string     `json:"cause_of_disability"`
	HouseNoAndStreet       string     `json:"house_no_and_street"`
	Barangay               string     `json:"barangay"`
	Municipality           string     `json:"municipality"`
	Province               string     `json:"province"`
	Region                 string     `json:"region"`
	LandlineNo             *string    `json:"landline_no,omitempty"`
	MobileNo               string     `json:"mobile_no"`
	EmailAddress           string     `json:"email_address"`
	EducationalAttainment  string     `json:"educational_attainment"`
	StatusOfEmployment     *string    `json:"status_of_employment,omitempty"`
	CategoryOfEmployment   *string    `json:"category_of_employment,omitempty"`
	TypeOfEmployment       *string    `json:"type_of_employment,omitempty"`
	Occupation             *string    `json:"occupation,omitempty"`
	OtherOccupation        *string    `json:"other_occupation,omitempty"`
	OrganizationAffiliated *string    `json:"organization_affiliated,omitempty"`
	ContactPerson          *string    `json:"contact_person,omitempty"`
	OfficeAddress          *string    `json:"office_address,omitempty"`
	OfficeTelNo            *string    `json:"office_tel_no,omitempty"`
	SSSNo                  *string    `json:"sss_no,omitempty"`
	GSISNo                 *string    `json:"gsis_no,omitempty"`
	PagIBIGNo              *string    `json:"pagibig_no,omitempty"`
	PSNNo                  *string    `json:"psn_no,omitempty"`
	PhilHealthNo           *string    `json:"philhealth_no,omitempty"`
	AccomplishByLastName   string     `json:"accomplish_by_last_name"`
	AccomplishByFirstName  string     `json:"accomplish_by_first_name"`
	AccomplishByMiddleName string     `json:"accomplish_by_middle_name"`
	ReferenceNumber        string     `json:"reference_number"`
	ValidUntil             *time.Time `json:"valid_until,omitempty"`
	DisapprovalDate        *time.Time `json:"disapproval_date,omitempty"`
	QRCode                 string     `json:"qr_code"`

	ApprovalStatus ApprovalStatus        `gorm:"foreignKey:AccountID" json:"approval_status"`
	Documents      []Document            `gorm:"foreignKey:AccountID" json:"documents"`
	Transactions   []DiscountTransaction `gorm:"foreignKey:AccountID" json:"transactions"`
}

func (a *Account) IsExpired() bool {
	return a.ValidUntil != nil && a.ValidUntil.Before(time.Now())
}


type ApprovalStatus struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	AccountID  uuid.UUID `gorm:"type:uuid;not null;unique" json:"account_id"`
	Status     string    `gorm:"type:varchar(20);default:'Pending'" json:"status"`
	IsApproved bool      `json:"is_approved"`

	Account *Account `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}


type Document struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	AccountID  uuid.UUID `gorm:"type:uuid;not null" json:"account_id"`
	Name       string    `gorm:"type:varchar(100)" json:"name"`
	FilePath   string    `gorm:"type:text" json:"file_path"`
	Validated  bool      `gorm:"default:false" json:"validated"`
	Remarks    *string   `json:"remarks,omitempty"`
	UploadedAt time.Time `gorm:"autoCreateTime" json:"uploaded_at"`

	Account *Account `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}

type Establishment struct {
	ID            uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Name          string    `json:"name"`
	Address       string    `json:"address"`
	ContactNumber string    `json:"contact_number"`
	OwnerName     string    `json:"owner_name"`
	Email         string    `gorm:"unique" json:"email"`
	Password      string    `json:"-"` // hashed
	CreatedAt     time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	Transactions []DiscountTransaction `gorm:"foreignKey:EstablishmentID" json:"transactions"`
}

type DiscountTransaction struct {
	ID              uint        `gorm:"primaryKey;autoIncrement" json:"id"`
	AccountID       uuid.UUID   `gorm:"type:uuid;not null" json:"account_id"`
	EstablishmentID uuid.UUID   `gorm:"type:uuid;not null" json:"establishment_id"`
	Amount          float64     `json:"amount"`
	DiscountApplied float64     `json:"discount_applied"`
	DateAvailed     time.Time   `gorm:"autoCreateTime" json:"date_availed"`

	Account       Account       `gorm:"foreignKey:AccountID" json:"account"`
	Establishment Establishment `gorm:"foreignKey:EstablishmentID" json:"establishment"`
}

type Admin struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `gorm:"unique" json:"email"`
	Password  string    `json:"-"` // hashed
	Role      string    `json:"role"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type Announcement struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedBy uuid.UUID `gorm:"type:uuid" json:"created_by"` // Admin ID
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}
