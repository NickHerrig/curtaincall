package entity

type Theater struct {
    ID          int    `json:"id"`
    Name        string `json:"name"`
    Address     string `json:"address"`
    Description string `json:"description"`
}

func (t *Theater) Validate() map[string]string {
  // TODO: Data validation and cleansing here. 
}
