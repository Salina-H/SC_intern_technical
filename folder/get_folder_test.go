package folder_test

import (
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

// feel free to change how the unit test is structured
func Test_folder_GetAllFolders(t *testing.T) {
	t.Parallel()
	tests := [...]struct {
		name        string
		dataFile    string
		want        []folder.Folder
	}{
		{
			name: "check all folders retrieved in example scenario",
 			dataFile: "./testData/moveFolder_exampleScenario.json",
	 		want: folder.GetSampleData("./testData/moveFolder_exampleScenario.json"),
 		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := folder.GetAllFolders(tt.dataFile)
			assert.Equal(t, tt.want, output, "not equal")
		})
	}
}
func Test_folder_GetFoldersByOrgID(t *testing.T) {
	const dataFile = "./testData/moveFolder_exampleScenario.json"
	t.Parallel()
	tests := [...]struct {
		name    string
		orgID   uuid.UUID
		folders []folder.Folder
		want    []folder.Folder
	}{
		{
			name: "check one folder with orgid retrieved",
			orgID: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
			folders: folder.GetSampleData(dataFile),
	 		want: folder.GetSampleData("./testData/getFolderByOrgId_Test_1.json"),
 		},
		 {
			name: "check multiple folders with orgid folder retrieved",
			orgID: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
			folders: folder.GetSampleData(dataFile),
	 		want: folder.GetSampleData("./testData/getFolderByOrgId_Test_2.json"),
 		},
		 {
			name: "check no folders with orgid folder retrieved",
			orgID: uuid.FromStringOrNil("18c9879b-f7eb-4b0e-b9d9-4fc4c23643a6"),
			folders: folder.GetSampleData(dataFile),
	 		want: folder.GetSampleData("./testData/getFolderByOrgId_Test_3.json"),
 		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)
			get := f.GetFoldersByOrgID(tt.orgID)
			assert.ElementsMatch(t, get, tt.want)
		})
	}
}
func Test_folder_GetAllChildFolders(t *testing.T) {
	const dataFile = "./testData/getAllChildFolders_exampleScenario.json"
	t.Parallel()
	tests := [...]struct {
		name          string
		orgID         uuid.UUID
		folder        string
		folders       []folder.Folder
		want          []folder.Folder
		isError       bool
		errorString   string
	}{
		{
			name: "folder has multiple levels of children",
			orgID: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
			folder: "alpha",
			folders: folder.GetSampleData(dataFile),
	 		want: folder.GetSampleData("./testData/getAllChildFolders_Test_1.json"),
			isError: false,
 		},
		{
			name: "folder has one direct child",
			orgID: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
			folder: "bravo",
			folders: folder.GetSampleData(dataFile),
	 		want: folder.GetSampleData("./testData/getAllChildFolders_Test_2.json"),
			isError: false,
 		},
		{
			name: "root folder with no children",
			orgID: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
			folder: "charlie",
			folders: folder.GetSampleData(dataFile),
	 		want: []folder.Folder{},
			isError: false,
 		},
		{
			name: "sub folder with no children",
			orgID: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
			folder: "echo",
			folders: folder.GetSampleData(dataFile),
	 		want: []folder.Folder{},
			isError: false,
 		},
		{
			name: "folder does not exist",
			orgID: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
			folder: "invalid",
			folders: folder.GetSampleData(dataFile),
			isError: true,
			errorString: "folder does not exist",
 		},
		{
			name: "organisation does not exist",
			orgID: uuid.FromStringOrNil("c1576e17-b7c9-45a3-a6ae-9546248fb17a"),
			folder: "alpha",
			folders: folder.GetSampleData(dataFile),
			isError: true,
			errorString: "organisation does not exist",
 		},
		{
			name: "folder does not exist in the specified organisation",
			orgID: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
			folder: "foxtrot",
			folders: folder.GetSampleData(dataFile),
			isError: true,
			errorString: "folder does not exist in the specified organisation",
 		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)
			res, err := f.GetAllChildFolders(tt.orgID, tt.folder)
	
			if tt.isError {
				assert.Nil(t, res)
				assert.EqualError(t, err, tt.errorString)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, res)
			}
		})
	}
}
