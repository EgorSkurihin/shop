package main

type Album struct {
	Title     string `schema:"title" json:"title"`
	Performer string `schema:"performer" json:"performer"`
	Cost      int    `schema:"cost" json:"cost"`
	Image     string `schema:"image" json:"image"`
}

type User struct {
	Id             int    `schema:"id"`
	Name           string `schema:"login"`
	OpenPassword   string `schema:"password"`
	HashedPassword string `schema:"-"`
}

type MailData struct {
	UName      string         `schema:"name"`
	UEmail     string         `schema:"email"`
	UTel       string         `schema:"tel"`
	CartString string         `schema:"cart"`
	CartObj    map[string]int `schema:"-"`
}

/* func (u *User) BeforeCreate() error {
	if len(u.OpenPassword) > 0 {
		h, err := hashString(u.OpenPassword)
		if err != nil {
			return err
		}
		u.HashedPassword = h
		u.OpenPassword = ""
	}
	return nil
}

func hashString(s string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
*/
