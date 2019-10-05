package main

type HandlingResult struct {
	RequestId string            `json:"request_id"`
	Success   bool              `json:"success"`
	Comment   string            `json:"comment"`
	Variables map[string]string `json:"variables"`
}

type AddPrimePhoneType struct {
	RequestId  string `json:"request_id"`
	Email      string `json:"email"`
	Location   string `json:"location"`
	Phone_type string `json:"phone_type"`
	Submitter  string `json:"submitter"`
}

type AddVoiceMailType struct {
	RequestId    string `json:"request_id"`
	Email        string `json:"email"`
	Sp_name      string `json:"sp_name"`
	Sp_email     string `json:"sp_email"`
	Sp_number    string `json:"sp_number"`
	Opened_by    string `json:"opened_by"`
	Phone_number string `json:"phone_number"`
	Evm          string `json:"evm"`
	Location     string `json:"location"`
}

type ModVoiceMailType struct {
	RequestId    string `json:"request_id"`
	Email        string `json:"email"`
	Sp_name      string `json:"sp_name"`
	Sp_email     string `json:"sp_email"`
	Sp_number    string `json:"sp_number"`
	Opened_by    string `json:"opened_by"`
	Phone_number string `json:"phone_number"`
	Evm          string `json:"evm"`
	Location     string `json:"location"`
	Voicemail    string `json:"voicemail"`
}
