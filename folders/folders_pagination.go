package folders

import (
	
	"fmt"
	"strconv"
)

const PageSize = 10 // Define the page size

type PaginatedFoldersResponse struct {
	Folders []*Folder
	Token   string
}

func GetAllFoldersPaginated(req *FetchFolderRequest, pageToken string) (*PaginatedFoldersResponse, error) {
	allFolders, err := FetchAllFoldersByOrgID(req.OrgID)
	if err != nil {
		return nil, err // proper error handling
	}

	// Calculate the pagination boundaries
	page, err := decodePageToken(pageToken)
	if err != nil {
		return nil, err // handle token decode error
	}
	start := page * PageSize
	end := start + PageSize
	if end > len(allFolders) {
		end = len(allFolders)
	}

	// Slice the full array to get just the page we're interested in
	foldersPage := allFolders[start:end]

	// Generate the next page token if there are more results
	var nextPageToken string
	if end < len(allFolders) {
		nextPageToken = encodePageToken(page + 1)
	}

	paginatedResponse := &PaginatedFoldersResponse{
		Folders: foldersPage,
		Token:   nextPageToken,
	}

	return paginatedResponse, nil
}

// Helper functions to encode/decode the page number into/from a page token.
func encodePageToken(page int) string {
	// This is where you'd implement your token encoding logic,
	// which might include encryption or signing to ensure its validity and integrity.
	return fmt.Sprintf("%d", page) // Simplest possible implementation for example purposes
}

func decodePageToken(token string) (int, error) {
	// Decode the token back into a page number.
	// Again, you'd include error checking and more complex logic in a real-world app.
	page, err := strconv.Atoi(token)
	if err != nil {
		return 0, err // token was not an integer
	}
	return page, nil
}
