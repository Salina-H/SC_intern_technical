package folder
import (
	"errors"
	"strings"
)
// preconditions/assumptions:
// - 	The file structure is accurate. E.g a folder with path: "alpha.bravo",
//		will not exist without a folder "alpha" in the same organisation,
//		all paths are valid, all names are valid (not empty)
// - 	A valid input will be given, i.e not an empty string or a nil UUID
// - 	every folder has a unique name
// - 	inputs are case sensitive so "Folder" is different from "folder"
func (f *driver) MoveFolder(name string, dst string) ([]Folder, error) {
	if name == dst {
		return nil, errors.New("cannot move a folder to itself")
	}
	
	folders := f.folders

	// retrieve the source folder
	var sourceFolder Folder
	var destinationFolder Folder

	for _, folder := range folders {
		if folder.Name == name {
			sourceFolder = folder
			if destinationFolder.Name != "" {
				break;
			}
		} else if folder.Name == dst {
			destinationFolder = folder
			if sourceFolder.Name != "" {
				break;
			}
		}
	}
	if sourceFolder.Name == "" {
		return nil, errors.New("source folder does not exist")
	} else if destinationFolder.Name == "" {
		return nil, errors.New("destination folder does not exist")
	} else if sourceFolder.OrgId != destinationFolder.OrgId {
		return nil, errors.New("cannot move a folder to a different organisation")
	} else if strings.HasPrefix(destinationFolder.Paths, sourceFolder.Paths + ".") {
		return nil, errors.New("cannot move a folder to a child of itself")
	} else if sourceFolder.Paths == destinationFolder.Paths + "." + name {
		return nil, errors.New("source folder is already in destination folder")
	}

	for i, folder := range folders {
		if strings.HasPrefix(folder.Paths, sourceFolder.Paths + ".") || folder.Paths == sourceFolder.Paths {
			folders[i].Paths = strings.Replace(folder.Paths, sourceFolder.Paths, destinationFolder.Paths + "." + name, 1)
		}
	}
	return folders, nil
}