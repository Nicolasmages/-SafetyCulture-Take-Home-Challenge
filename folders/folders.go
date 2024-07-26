package folders

import (
	"github.com/gofrs/uuid"
)

/*
Remove variable that is not used, remove two for loop with k and k1 as it basically doing the same thing as 
folders, err = FetchAll Folders OrgID(req.OrgID)
*/
// GetAllFolders retrieves all folders for the specified organization ID.
func GetAllFolders(req *FetchFolderRequest) (*FetchFolderResponse, error) {

	var (
		folders []*Folder
		err error
	)

	// Fetch all folders by organization ID
	folders, err = FetchAllFoldersByOrgID(req.OrgID)
	if err != nil {
		return nil, err// Return error if fetching folders fails
	}

	// Return the fetched folders in a FetchFolderResponse struct
	return &FetchFolderResponse{Folders: folders}, nil
}

/*
Declare variables before use and nothing else is changed
*/
// FetchAllFoldersByOrgID fetches all folders by the given organization ID.
func FetchAllFoldersByOrgID(orgID uuid.UUID) ([]*Folder, error) {

	var (
		folders []*Folder 
		resFolders []*Folder
	)

	// Fetch sample data (presumably from a mocked or sample source)
	folders = GetSampleData()
	
	// Iterate through the fetched folders
	for _, folder := range folders {
		// Check if the folder's organization ID matches the provided orgID
		if folder.OrgId == orgID {
			resFolders = append(resFolders, folder)
		}
	}

	// Return the filtered folders
	return resFolders, nil
}
