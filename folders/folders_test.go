package folders_test

import (
	"testing"
	"github.com/georgechieng-sc/interns-2022/folders"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

// Test_GetAllFolders tests the GetAllFolders function in the folders package
func Test_GetAllFolders(t *testing.T) {

	// Subtest for fetching folders with a valid OrgID
	t.Run("GetAllFoldersTest", func(t *testing.T) {

		// Creating a sample request with a default organization ID
		req := &folders.FetchFolderRequest{
			OrgID: uuid.FromStringOrNil(folders.DefaultOrgID),
		}

		// Calling the GetAllFolders function
		res, err := folders.GetAllFolders(req)

		// Asserting that the response is not nil
		assert.NotNil(t, res)

		// Asserting that the error is nil
		assert.Nil(t, err)

		// Asserting that the number of folders returned is 666
		assert.Equal(t, len(res.Folders), 666)

	})
	
	// Subtest for fetching folders with a non-existent OrgID
	t.Run("GetAllFoldersNoEntryTest", func(t *testing.T) {

		// / Creating a test UUID that doesn't exist in the sample data
		testOrgID := uuid.FromStringOrNil("33333333-3333-3333-3333-333333333333")

		// Creating a sample request with the test UUID
		req := &folders.FetchFolderRequest{
			OrgID: testOrgID,
		}

		// Calling the GetAllFolders function
		res, err := folders.GetAllFolders(req)
		
		// Asserting that the response is not nil
		assert.NotNil(t, res)

		// Asserting that the error is nil
		assert.Nil(t, err)

		// Asserting that no folders are returned
		assert.Equal(t, len(res.Folders), 0)
	})
}
