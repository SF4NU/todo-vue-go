package models

type LoginRequest struct { // questa struttura viene utilizzata solo in fase di login
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct { //questa struttura definisce la tabella degli user in postgresql
	ID         uint       `gorm:"primaryKey"`
	Username   string     `json:"username"`
	Password   string     `json:"password"`
	Categories []Category `gorm:"foreignKey:UserID"` //la relazione è one-to-many cioè uno a molti, perché la tabella users è assocciata a più  categorie
}
type Category struct { //questa struttura definisce la tabella delle categorie in postgresql
	ID           uint   `gorm:"primaryKey"`
	Title        string `json:"title"`
	LastModified string `json:"last_modified"`
	UserID       uint   `json:"user_id"`
	Tasks        []Task `gorm:"foreignKey:CategoryID"` //la relazione è one-to-many cioè uno a molti, perché la tabella categorie è assocciata a più  task
}
type Task struct { //questa struttura definisce la tabella delle task in postgresql
	ID           uint   `gorm:"primaryKey"`
	Description  string `json:"description"`
	Completed    bool   `json:"completed"`
	LastModified string `json:"last_modified"`
	CategoryID   uint   `json:"category_id"`
}
