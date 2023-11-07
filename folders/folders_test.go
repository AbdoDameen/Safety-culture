package folders_test

import (
    "testing"
    "github.com/georgechieng-sc/interns-2022/folders"
    "github.com/gofrs/uuid"
    "github.com/stretchr/testify/assert"
)

func TestGetAllFolders(t *testing.T) {
    t.Run("should return folders without error", func(t *testing.T) {
        // Mock the data or use some sample data that matches the expected structure
        orgID := uuid.Must(uuid.NewV4())
        req := folders.FetchFolderRequest{OrgID: orgID}
        
        // Call the function under test
        resp, err := folders.GetAllFolders(&req)
        
        // Assert the expectations
        assert.NoError(t, err)
        assert.NotNil(t, resp)
        // Add more assertions based on the expected behavior of the function
    })
}
