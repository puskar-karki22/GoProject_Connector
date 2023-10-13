package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type bodyData struct {
	TotalCount int `json:"total_count"`
	Entries    []struct {
		Type        string `json:"type"`
		ID          string `json:"id"`
		SequenceID  string `json:"sequence_id"`
		Etag        string `json:"etag"`
		Name        string `json:"name"`
		FileVersion struct {
			Type string `json:"type"`
			ID   string `json:"id"`
			Sha1 string `json:"sha1"`
		} `json:"file_version,omitempty"`
		Sha1 string `json:"sha1,omitempty"`
		URL  string `json:"url,omitempty"`
	} `json:"entries"`
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
	Order  []struct {
		By        string `json:"by"`
		Direction string `json:"direction"`
	} `json:"order"`
}
type FileInfo struct {
	Type              string         `json:"type"`
	ID                string         `json:"id"`
	FileVersion       FileVersion    `json:"file_version"`
	SequenceID        string         `json:"sequence_id"`
	Etag              string         `json:"etag"`
	Sha1              string         `json:"sha1"`
	Name              string         `json:"name"`
	Description       string         `json:"description"`
	Size              int            `json:"size"`
	PathCollection    PathCollection `json:"path_collection"`
	CreatedAt         string         `json:"created_at"`
	ModifiedAt        string         `json:"modified_at"`
	TrashedAt         any            `json:"trashed_at"`
	PurgedAt          any            `json:"purged_at"`
	ContentCreatedAt  string         `json:"content_created_at"`
	ContentModifiedAt string         `json:"content_modified_at"`
	CreatedBy         CreatedBy      `json:"created_by"`
	ModifiedBy        ModifiedBy     `json:"modified_by"`
	OwnedBy           OwnedBy        `json:"owned_by"`
	SharedLink        any            `json:"shared_link"`
	Parent            Parent         `json:"parent"`
	ItemStatus        string         `json:"item_status"`
}
type FileVersion struct {
	Type string `json:"type"`
	ID   string `json:"id"`
	Sha1 string `json:"sha1"`
}
type Entries struct {
	Type       string `json:"type"`
	ID         string `json:"id"`
	SequenceID any    `json:"sequence_id"`
	Etag       any    `json:"etag"`
	Name       string `json:"name"`
}
type PathCollection struct {
	TotalCount int       `json:"total_count"`
	Entries    []Entries `json:"entries"`
}
type CreatedBy struct {
	Type  string `json:"type"`
	ID    string `json:"id"`
	Name  string `json:"name"`
	Login string `json:"login"`
}
type ModifiedBy struct {
	Type  string `json:"type"`
	ID    string `json:"id"`
	Name  string `json:"name"`
	Login string `json:"login"`
}
type OwnedBy struct {
	Type  string `json:"type"`
	ID    string `json:"id"`
	Name  string `json:"name"`
	Login string `json:"login"`
}
type Parent struct {
	Type       string `json:"type"`
	ID         string `json:"id"`
	SequenceID any    `json:"sequence_id"`
	Etag       any    `json:"etag"`
	Name       string `json:"name"`
}
type Download struct {
	Type       string `json:"type"`
	ID         string `json:"id"`
	Etag       string `json:"etag"`
	SharedLink struct {
		URL                 string `json:"url"`
		DownloadURL         string `json:"download_url"`
		VanityURL           any    `json:"vanity_url"`
		VanityName          any    `json:"vanity_name"`
		EffectiveAccess     string `json:"effective_access"`
		EffectivePermission string `json:"effective_permission"`
		IsPasswordEnabled   bool   `json:"is_password_enabled"`
		UnsharedAt          any    `json:"unshared_at"`
		DownloadCount       int    `json:"download_count"`
		PreviewCount        int    `json:"preview_count"`
		Access              string `json:"access"`
		Permissions         struct {
			CanPreview  bool `json:"can_preview"`
			CanDownload bool `json:"can_download"`
			CanEdit     bool `json:"can_edit"`
		} `json:"permissions"`
	} `json:"shared_link"`
}

func main() {
	// info := bodyData{
	// 	grant_type:    "client_credentials",
	// 	client_id:     "lk0kcp7ap85pxu2qs4gva8wfzckdt76h",
	// 	client_secret: "1DSyD42okEunrm6tmzBTzJh7UDE1RaoY",
	// }
	// url_1 := "https://api.box.com/oauth2/token"
	// client_idd := "lk0kcp7ap85pxu2qs4gva8wfzckdt76h"
	// client_secret_keyy := "1DSyD42okEunrm6tmzBTzJh7UDE1RaoY"
	// body, _ := json.Marshal(info)
	// resp, err := http.Post("https://api.box.com/oauth2/token", "application/json", bytes.NewBuffer(body))
	// if err != nil {
	// 	panic(err)
	// }
	// defer resp.Body.Close()

	// if resp.StatusCode == http.StatusCreated {
	// 	body, err := io.ReadAll(resp.Body)
	// 	if err != nil {
	// 		//Failed to read response.
	// 		panic(err)
	// 	}

	// 	//Convert bytes to String and print
	// 	jsonStr := string(body)
	// 	fmt.Println("Response: ", jsonStr)

	// } else {
	// 	//The status is not Created. print the error.
	// 	fmt.Println("Get failed with error: ", resp.Status)
	// }
	url := "https://api.box.com/2.0/folders/0/items"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("authorization", "Bearer tr2RM43RmFHmqkS5PTvH9xYHNhodFoIM")
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	var responseObject bodyData
	json.Unmarshal(body, &responseObject)
	// fmt.Printf("API Response as struct %+v\n", responseObject)
	// fmt.Println(responseObject.TotalCount)
	// fmt.Println(responseObject.limits)
	for i := 0; i < 4; i++ {
		fmt.Println("Title of the entity : ", responseObject.Entries[i].Name)

		fmt.Println("ID : ", responseObject.Entries[i].ID)
		fmt.Println("Type : ", responseObject.Entries[i].Type)
		{
			var alpha string = responseObject.Entries[i].ID
			url_2 := fmt.Sprintf("https://api.box.com/2.0/files/" + alpha)
			reqq, _ := http.NewRequest("GET", url_2, nil)
			reqq.Header.Add("authorization", "Bearer tr2RM43RmFHmqkS5PTvH9xYHNhodFoIM")
			ress, _ := http.DefaultClient.Do(reqq)
			defer ress.Body.Close()
			body, _ := io.ReadAll(ress.Body)
			var responseObject FileInfo
			json.Unmarshal(body, &responseObject)
			fmt.Println("Owner : ", responseObject.CreatedBy.Name)
			fmt.Println("Last Modified Time : ", responseObject.ModifiedAt)
		}
		url_3 := fmt.Sprintf("https://api.box.com/2.0/files/" + responseObject.Entries[i].ID + "?fields=shared_link")
		reqqq, _ := http.NewRequest("GET", url_3, nil)
		reqqq.Header.Add("authorization", "Bearer tr2RM43RmFHmqkS5PTvH9xYHNhodFoIM")
		resss, _ := http.DefaultClient.Do(reqqq)
		defer resss.Body.Close()
		body_1, _ := io.ReadAll(resss.Body)
		var responseObject Download
		json.Unmarshal(body_1, &responseObject)
		fmt.Println("url ", responseObject.SharedLink.URL)
		fmt.Println("download_url ", responseObject.SharedLink.DownloadURL)
		fmt.Println()

	}

	// fmt.Println(res)
	// 	fmt.Println(string(body))
}
