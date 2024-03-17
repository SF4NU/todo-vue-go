package models

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	ID         uint       `gorm:"primaryKey"`
	Username   string     `json:"username"`
	Password   string     `json:"password"`
	Categories []Category `gorm:"foreignKey:UserID"`
}
type Category struct {
	ID           uint   `gorm:"primaryKey"`
	Title        string `json:"title"`
	LastModified string `json:"last_modified"`
	UserID       uint   `json:"user_id"`
	Todos        []Task `gorm:"foreignKey:CategoryID"`
}
type Task struct {
	ID           uint   `gorm:"primaryKey"`
	Description  string `json:"description"`
	Completed    bool   `json:"completed"`
	LastModified string `json:"last_modified"`
	CategoryID   uint   `json:"category_id"`
} //la relazione è one-to-many cioè uno a molti perché la tabella categorie è assocciata a più tabelle task
