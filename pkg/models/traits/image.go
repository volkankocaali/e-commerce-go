package traits

type Image struct {
	ImagePath string `json:"image_path" sql:"type:text;"`
}
