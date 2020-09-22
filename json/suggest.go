package json

/*
 * json for suggest
 */

//suggest opt json
type SuggestOptJson struct {
	Key string `json:"key"`
	Kind int `json:"kind"`
	Size int `json:"size"`
	BaseJson
}

 //suggest doc json
 type SuggestJson struct {
 	Kind int `json:"kind"`
 	Key string `json:"key"`
 	Count string `json:"count"`
 	BaseJson
 }

 type SuggestsJson struct {
 	Total int `json:"total"`
 	List []*SuggestJson `json:"list"`
 	BaseJson
 }


///////////////////////////
//construct for SuggestsJson
//////////////////////////

func NewSuggestsJson() *SuggestsJson {
	this := &SuggestsJson{
		List: make([]*SuggestJson, 0),
	}
	return this
}

 //add obj
func (j *SuggestsJson) AddObj(obj *SuggestJson) bool {
	if obj == nil {
		return false
	}
	j.List = append(j.List, obj)
	return true
}

//encode json data
func (j *SuggestsJson) Encode() []byte {
	return j.BaseJson.Encode(j)
}

//decode json data
func (j *SuggestsJson) Decode(data []byte) bool {
	return j.BaseJson.Decode(data, j)
}

///////////////////////////
//construct for SuggestJson
//////////////////////////

func NewSuggestJson() *SuggestJson {
	this := &SuggestJson{
	}
	return this
}

//encode json data
func (j *SuggestJson) Encode() []byte {
	return j.BaseJson.Encode(j)
}

//decode json data
func (j *SuggestJson) Decode(data []byte) bool {
	return j.BaseJson.Decode(data, j)
}


///////////////////////////
//construct for SuggestOptJson
//////////////////////////

func NewSuggestOptJson() *SuggestOptJson {
	this := &SuggestOptJson{
	}
	return this
}

//encode json data
func (j *SuggestOptJson) Encode() []byte {
	return j.BaseJson.Encode(j)
}

//decode json data
func (j *SuggestOptJson) Decode(data []byte) bool {
	return j.BaseJson.Decode(data, j)
}