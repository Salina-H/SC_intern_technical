package folder
import (
	"strings"
	"errors"
	"github.com/gofrs/uuid"
)
func GetAllFolders(data ...string) []Folder {
	return GetSampleData(data...)
}

func (f *driver) GetFoldersByOrgID(orgID uuid.UUID) []Folder {
	folders := f.folders

	res := []Folder{}
	for _, f := range folders {
		if f.OrgId == orgID {
			res = append(res, f)
		}
	}

	return res
}

// preconditions/assumptions:
// - 	The file structure is accurate. E.g a folder with path: "alpha.bravo",
// 		will not exist without a folder "alpha" in the same organisation
// - 	A valid input will be given, i.e not an empty string or a nil UUID
// - 	every folder has a unique orgID and name combo. E.g a folder can't have the same name
//	 	and be in the same org (Based off component 2, assuming that no 2 folders will have the same name)
// - 	inputs are case sensitive so "Folder" is different from "folder"
func (f *driver) GetAllChildFolders(orgID uuid.UUID, name string) ([]Folder, error) {
	var wrongOrg, orgExists, folderFound bool
	folders := f.folders

	res := []Folder{}
	for _, folder := range folders {
		if (strings.Contains(folder.Paths, "." + name + ".") || strings.HasPrefix(folder.Paths, name + ".")) && folder.OrgId == orgID {
			res = append(res, folder)
		}
		if folder.OrgId == orgID {
			orgExists = true
		}

		if folder.Name == name {
			folderFound = true
			if folder.OrgId != orgID {
				wrongOrg = true
			}
		}
	}
	if len(res) == 0 {
		if !orgExists {
			return nil, errors.New("organisation does not exist")
		} else if wrongOrg {
			return nil, errors.New("folder does not exist in the specified organisation")
		} else if !folderFound {
			return nil, errors.New("folder does not exist")
		}
	}
	return res, nil
}