package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
	"time"
)

//for reach test
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Starfish Request API - Welcome Page")

}

//LHM Tips testing
func BoeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("In GetBOELocationData handler")
	var keys []string
	for k, _ := range r.Header {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		for _, hval := range r.Header[k] {
			fmt.Println(k, ":", hval)
		}
		//fmt.Println(k, ":", r.Header[k])
	}
	fmt.Println()

	auth_val := r.Header["Authorization"]
	indx := strings.Index(auth_val[0], " ")
	encoded := auth_val[0][indx+1:]
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("auth decode error:", err)

	}
	fmt.Println("auth decoded: " + string(decoded))

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	logger.Printf("BoeHandler - %v\n", http.StatusOK)

	//pStr := string("{\"A\":\"B\"}")
	pStr := "{\"ValidationResult\":{\"errorCode\":\"200\",\"validationMessage\":\"ok valid\",\"wasSuccessful\":\"true\"},\"BOEDataItemList\":[{\"State\":null,\"Country\":\"US\",\"Address1\":\"500 First St.\",\"LocationId\":\"1000\",\"Address2\":\"\",\"Address3\":\"\",\"City\":\"Wichita\",\"PostalCode\":\"10101\"},{\"State\":null,\"Country\":\"US\",\"Address1\":\"304 Second St.\",\"LocationId\":\"1001\",\"Address2\":\"\",\"Address3\":\"\",\"City\":\"Topeka\",\"PostalCode\":\"10303\"}]}"
	w.Write([]byte(pStr))

}

//for amex cookie testing
func PatchIt(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	bodyStr := string(body)
	fmt.Println("got post: " + bodyStr)

	cookies := r.Cookies()
	fmt.Printf("number of cookies: %d\n", len(cookies))
	for _, cookie := range cookies {
		fmt.Printf("cookie %s = %s\n", cookie.Name, cookie.Value)
	}

	w.WriteHeader(http.StatusOK)

}

//ShowAxl
func ShowAxl(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	bodyStr := string(body)
	fmt.Println("got ShowAxl: " + bodyStr)

	w.WriteHeader(http.StatusOK)
}


func GetTelephoneNumbersByGUID(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	bodyStr := string(body)
	fmt.Println("got ShowAxl: " + bodyStr)

	fmt.Println("returning canned response from file: GetTelephoneNumbersByGUID_Response.txt")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	b, _ := ioutil.ReadFile("GetTelephoneNumbersByGUID_Response.txt")
	w.Write(b)
}

//also for amex cookie testing
func Authenticate(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	bodyStr := string(body)
	fmt.Println("got authenticate: " + bodyStr)
	r.ParseForm()
	//u := r.PostFormValue("userid")
	//p := r.PostFormValue("pwd")
	for k, v := range r.Form {
		fmt.Printf("%s = %s\n", k, v)
	}

	expire := time.Now().AddDate(0, 0, 1)
	cookie := http.Cookie{Name: "csrftoken", Value: "abcd", Expires: expire}
	//http.SetCookie(w, &cookie)
	http.SetCookie(w, &cookie)
	w.WriteHeader(http.StatusOK)
	reply := "<soapenv:Envelope xmlns:soapenv=\"http://schemas.xmlsoap.org/soap/envelope/\"> <soapenv:Body>      <ns:updateRemoteDestinationResponse xmlns:ns=\"http://www.cisco.com/AXL/API/10.5\">         <return>{54B0F5A1-955D-490F-981B-081EF9B7BD80}</return>      </ns:updateRemoteDestinationResponse>   </soapenv:Body></soapenv:Envelope>"
	w.Write([]byte(reply))
}

//LHM
func AddVoiceMail(w http.ResponseWriter, r *http.Request) {
	logger.Printf("START HANDLING ADDVOICEMAIL\n")
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	bodyStr := string(body)
	fmt.Println("got: " + bodyStr)
	logger.Printf("in: %s\n", bodyStr)
	var addVM AddVoiceMailType

	err = json.Unmarshal(body, &addVM)
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		logger.Printf("FAIL: AddVoiceMail - %v\n", http.StatusUnprocessableEntity)
		res := HandlingResult{"?", false, err.Error(), make(map[string]string)}
		if err := json.NewEncoder(w).Encode(res); err != nil {
			panic(err)
		}
		logger.Printf("END HANDLING ADDVOICEMAIL\n")
	} else {

		fmt.Printf("servicing user: %s on request %s\n", addVM.Email, addVM.RequestId)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		logger.Printf("AddVoiceMail - %v\n", http.StatusOK)
		res := HandlingResult{addVM.RequestId, true, "addvoicemail request queued", make(map[string]string)}
		b, err := json.Marshal(res)
		if err == nil {
			pStr := string(b)
			w.Write([]byte(pStr))
		}
		logger.Printf("END HANDLING ADDVOICEMAIL\n")
	}
}

//LHM
func ModVoiceMail(w http.ResponseWriter, r *http.Request) {

	logger.Printf("START HANDLING MODVOICEMAIL\n")
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	bodyStr := string(body)
	fmt.Println("got: " + bodyStr)
	logger.Printf("in: %s\n", bodyStr)
	var modVM ModVoiceMailType

	err = json.Unmarshal(body, &modVM)
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		logger.Printf("FAIL: ModVoiceMail - %v\n", http.StatusUnprocessableEntity)
		w.WriteHeader(422) // unprocessable entity
		res := HandlingResult{"?", false, err.Error(), make(map[string]string)}
		if err := json.NewEncoder(w).Encode(res); err != nil {
			panic(err)
		}
		logger.Printf("END HANDLING MODVOICEMAIL\n")
	} else {

		fmt.Printf("servicing user: %s\n", modVM.Email)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		logger.Printf("ModVoiceMail - %v\n", http.StatusOK)
		res := HandlingResult{modVM.RequestId, true, "modvoicemail request queued", make(map[string]string)}
		b, err := json.Marshal(res)
		if err == nil {
			pStr := string(b)
			w.Write([]byte(pStr))
		}
		logger.Printf("END HANDLING MODVOICEMAIL\n")
	}
}

//LHM
func ReportIncident(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	bodyStr := string(body)
	fmt.Println("got incident POST: " + bodyStr)
	var request []string
	for name, headers := range r.Header {
		name = strings.ToLower(name)
		for _, h := range headers {
			request = append(request, fmt.Sprintf("%v: %v", name, h))
		}
	}
	hlist := strings.Join(request, "\n")
	fmt.Println("hearders: " + hlist)

	w.WriteHeader(http.StatusOK)

}

//LHM
func Voicemail(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	bodyStr := string(body)
	fmt.Println("got Voicemail POST: " + bodyStr)

	// Loop through headers
	var request []string
	for name, headers := range r.Header {
		name = strings.ToLower(name)
		for _, h := range headers {
			request = append(request, fmt.Sprintf("%v: %v", name, h))
		}
	}
	hlist := strings.Join(request, "\n")
	fmt.Println("hearders: " + hlist)
	w.WriteHeader(http.StatusOK)

}

//LHM
func Sponsored(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	bodyStr := string(body)
	fmt.Println("got Sponsored POST: " + bodyStr)

	// Loop through headers
	var request []string
	for name, headers := range r.Header {
		name = strings.ToLower(name)
		for _, h := range headers {
			request = append(request, fmt.Sprintf("%v: %v", name, h))
		}
	}
	hlist := strings.Join(request, "\n")
	fmt.Println("hearders: " + hlist)
	w.WriteHeader(http.StatusOK)

}


func DellBulkHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	bodyStr := string(body)
	fmt.Println("got BULK REPORT:\n" + bodyStr)

	w.WriteHeader(http.StatusOK)
}

func ChangePhoneResult(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	bodyStr := string(body)
	fmt.Println("got change_phone_post POST: " + bodyStr)

	w.WriteHeader(http.StatusOK)
}


func PrimePhoneResult(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	bodyStr := string(body)
	fmt.Println("got prim_phone_post.do POST: " + bodyStr)

	w.WriteHeader(http.StatusOK)

}

//TerminationCB
func TerminationCB(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	bodyStr := string(body)
	fmt.Println("got NewHireCB POST: " + bodyStr)
	logger.Printf("NewHireCB - %v\n", http.StatusOK)
	w.WriteHeader(http.StatusOK)
	res := HandlingResult{"dummyId", true, "Term response", make(map[string]string)}
	b, err := json.Marshal(res)
	if err == nil {
		pStr := string(b)
		fmt.Println("replying with: " + pStr)
		w.Write([]byte(pStr))
	}

}

func NewHireCB(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	bodyStr := string(body)
	fmt.Println("got NewHireCB POST: " + bodyStr)
	logger.Printf("NewHireCB - %v\n", http.StatusOK)
	w.WriteHeader(http.StatusOK)
	res := HandlingResult{"dummyId", true, "dummy response", make(map[string]string)}
	b, err := json.Marshal(res)
	if err == nil {
		pStr := string(b)
		fmt.Println("replying with: " + pStr)
		w.Write([]byte(pStr))
	}

}


//LHM
func AddPrimePhone(w http.ResponseWriter, r *http.Request) {
	logger.Printf("START HANDLING ADDPRIMEPHONE\n")
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	bodyStr := string(body)
	logger.Printf("in: %s\n", bodyStr)
	fmt.Println("got: " + bodyStr)
	var addP AddPrimePhoneType

	err = json.Unmarshal(body, &addP)
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		w.WriteHeader(422) // unprocessable entity
		logger.Printf("FAIL: AddPrimePhone - %v\n", http.StatusUnprocessableEntity)
		res := HandlingResult{"?", false, err.Error(), make(map[string]string)}
		if err := json.NewEncoder(w).Encode(res); err != nil {
			panic(err)
		}
		logger.Printf("END HANDLING ADDPRIMEPHONE\n")
	} else {

		fmt.Printf("servicing user: %s\n", addP.Email)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		logger.Printf("AddPrimePhone - %v\n", http.StatusOK)
		res := HandlingResult{addP.RequestId, true, "addprimephone request queued", make(map[string]string)}
		b, err := json.Marshal(res)
		if err == nil {
			pStr := string(b)
			w.Write([]byte(pStr))
		}
		logger.Printf("END HANDLING ADDPRIMEPHONE\n")
	}

}
