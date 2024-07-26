package folders

import (
	"errors"
	"github.com/gofrs/uuid"

	// To generate the token 
	"math/rand" 
	"time"
)

// Map to store pagination tokens and their corresponding start indices (Global)
var paginationMap = make(map[string]int) // Linked List but map is picked

// GetFoldersPaginsation retrieves all folders for the given request with pagination support.
func GetFoldersPaginsation(req *FetchFolderRequest, pageSize int, token string) (*FetchFolderResponse, string, error) {

	// Get the start index from the pagination token
	startIndex, err := getStartIndexFromToken(token)
	if err != nil {
		startIndex = 0 // default to 0 if token is not valid
	}

	// Fetch folders with pagination
	folders, nextToken, err := FetchAllFoldersByOrgIDPaginsation(req.OrgID, pageSize, startIndex)
	if err != nil {
		return nil, "", err
	}

	// Return the fetched folders and the next pagination token
	return &FetchFolderResponse{
		Folders: folders,
	}, nextToken, nil
}

// FetchAllFoldersByOrgIDPaginsation fetches folders by the given organization ID with pagination.
func FetchAllFoldersByOrgIDPaginsation(orgID uuid.UUID, pageSize int, startIndex int) ([]*Folder, string, error) {
	
	// Fetch sample data (presumably from a mocked or sample source)
	folders := GetSampleData()

	var resFolders []*Folder

	// Filter folders by organization ID
	for _, folder := range folders {
		if folder.OrgId == orgID {
			resFolders = append(resFolders, folder)
		}
	}

	// Paginate the results
	endIndex := startIndex + pageSize
	if endIndex > len(resFolders) {
		endIndex = len(resFolders)
	}

	var nextToken string
	if endIndex < len(resFolders) {
		// Generate the next pagination token if there are more folders to fetch
		nextToken = generateRandomToken()
		paginationMap[nextToken] = endIndex
	}

	// Return the paginated results and the next pagination token
	return resFolders[startIndex:endIndex], nextToken, nil
}

// getStartIndexFromToken retrieves the start index from the pagination token.
func getStartIndexFromToken(token string) (int, error) {

	//Initial handshake no token
	if token == "" {
		return 0, nil
	}


	startIndex, exists := paginationMap[token]
	if !exists {
		return 0, errors.New("invalid token")
	}
	return startIndex, nil
}

// generateRandomToken creates a random string of length 5.
func generateRandomToken() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	//Make the token length 5
	b := make([]byte, 5)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

/*
Write a short explanation of your chosen solution.

I use generate Random Token to generate token and accept empty string as the inital handshake, 
make the map data structure for the token and index storage,
hence given token we can slide to the corresponding index for the list.
*/