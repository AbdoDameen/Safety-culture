package folders

import (
	"github.com/gofrs/uuid"
)

// the GetAllFolders function returns all folders for the the organization.
func GetAllFolders(req *FetchFolderRequest) (*FetchFolderResponse, error) {
	r, err := FetchAllFoldersByOrgID(req.OrgID)
	if err != nil {
		return nil, err // interducing a proper error handling
	}

	// here i'm trying to directly work with pointers to Folder
	var fp []*Folder
	for _, v := range r {
		fp = append(fp, v) // hence no need to re-reference
	}

	ffr := &FetchFolderResponse{Folders: fp}
	return ffr, nil
}

// the FetchAllFoldersByOrgID function fetches all folders associated with a given OrgID.
func FetchAllFoldersByOrgID(orgID uuid.UUID) ([]*Folder, error) {
	folders := GetSampleData() // Consider renaming or refactoring to indicate the source of data!!

	var resFolder []*Folder
	for _, folder := range folders {
		if folder.OrgId == orgID {
			resFolder = append(resFolder, folder)
		}
	}
	return resFolder, nil
}