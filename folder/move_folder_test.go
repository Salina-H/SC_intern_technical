package folder_test

import (
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/stretchr/testify/assert"
)

func Test_folder_MoveFolder(t *testing.T) {
	const dataFile = "./testData/moveFolder_exampleScenario.json"
	t.Parallel()
	tests := [...]struct {
		name    string
		source  string
		dst     string
		folders []folder.Folder
		want    []folder.Folder
		isError   bool
		errorString string
	}{
		{
			name:    "case where destination shares a parent folder",
			source:  "bravo",
			dst:     "delta",
			folders: folder.GetSampleData(dataFile),
			want:    folder.GetSampleData("./testData/moveFolder_Test_1.json"),
			isError: false,
		},
		{
			name:    "case where destination does not have a parent folder",
			source:  "bravo",
			dst:     "golf",
			folders: folder.GetSampleData(dataFile),
			want:    folder.GetSampleData("./testData/moveFolder_Test_2.json"),
			isError: false,
		},
		{
			name:        "moving a folder to a child of itself",
			source:      "bravo",
			dst:         "charlie",
			folders:     folder.GetSampleData(dataFile),
			isError:     true,
			errorString: "cannot move a folder to a child of itself",
		},
		{
			name:        "moving a folder to itself",
			source:      "bravo",
			dst:         "bravo",
			folders:     folder.GetSampleData(dataFile),
			isError:     true,
			errorString: "cannot move a folder to itself",
		},
		{
			name:        "moving a folder to a different organisation",
			source:      "bravo",
			dst:         "foxtrot",
			folders:     folder.GetSampleData(dataFile),
			isError:     true,
			errorString: "cannot move a folder to a different organisation",
		},
		{
			name:        "source folder does not exist",
			source:      "invalid_folder",
			dst:         "delta",
			folders:     folder.GetSampleData(dataFile),
			isError:     true,
			errorString: "source folder does not exist",
		},
		{
			name:        "destination folder does not exist",
			source:      "bravo",
			dst:         "invalid_folder",
			folders:     folder.GetSampleData(dataFile),
			isError:     true,
			errorString: "destination folder does not exist",
		},
		{
			name:        "source folder is already in destination folder",
			source:      "charlie",
			dst:         "bravo",
			folders:     folder.GetSampleData(dataFile),
			isError:     true,
			errorString: "source folder is already in destination folder",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)
			res, err := f.MoveFolder(tt.source, tt.dst)
	
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
