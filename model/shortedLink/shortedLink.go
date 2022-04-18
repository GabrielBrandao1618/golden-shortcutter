package shortedLink

import "github.com/google/uuid"

type ShortedLink struct {
	Ref      string `json:"ref"`
	HashCode string `json:"hashCode"`
}

func New(ref string) ShortedLink {
	return ShortedLink{Ref: ref, HashCode: uuid.New().String()}
}
