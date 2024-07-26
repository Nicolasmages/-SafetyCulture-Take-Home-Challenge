package main

import (
	"fmt"

	"github.com/georgechieng-sc/interns-2022/folders"
	"github.com/gofrs/uuid"
)

func main() {

	req := &folders.FetchFolderRequest{
		OrgID: uuid.FromStringOrNil(folders.DefaultOrgID),
	}

	var (
		nextToken string
		err       error
		res       *folders.FetchFolderResponse
	)

	for {
		// same as spec
		res, nextToken, err = folders.GetFoldersPaginsation(req, 2, nextToken)

		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		folders.PrettyPrint(res)
		fmt.Printf("Next token is %v\n", nextToken)

		if nextToken == "" {
			break
		}
	}

}
