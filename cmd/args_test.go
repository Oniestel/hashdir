package cmd

import "testing"

func TestValidateError_Error(t *testing.T) {
	tests := []struct {
		name string
		r    ValidateError
		want string
	}{
		{
			name:"Test ValidateError",
			r: ValidateError{Msg:"Test"},
			want:"Test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.Error(); got != tt.want {
				t.Errorf("ValidateError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getPath(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name     string
		args     args
		wantPath string
	}{
		{
			name:"Not empty",
			args:args{
				args: []string{"arg1","arg2"},
			},
			wantPath:"arg1",
		},
		{
			name:"Empty",
			args:args{
				args: []string{},
			},
			wantPath:"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotPath := getPath(tt.args.args); gotPath != tt.wantPath {
				t.Errorf("getPath() = %v, want %v", gotPath, tt.wantPath)
			}
		})
	}
}

func Test_checkPath(t *testing.T) {
	defer removeTestDir()
	dirPath:=createTestDir(t)
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:"Empty",
			args:args{
				path:"",
			},
			wantErr:true,
		},
		{
			name:"Not exist",
			args:args{
				path:"error",
			},
			wantErr:true,
		},
		{
			name:"Exist",
			args:args{
				path:dirPath,
			},
			wantErr:false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := checkPath(tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("checkPath() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_checkStorage(t *testing.T) {
	type args struct {
		storage string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:"Empty",
			args:args{
				storage:"",
			},
			wantErr:true,
		},
		{
			name:"Not empty",
			args:args{
				storage:"test",
			},
			wantErr:false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := checkStorage(tt.args.storage); (err != nil) != tt.wantErr {
				t.Errorf("checkStorage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_checkDest(t *testing.T) {
	defer removeTestDir()
	dirPath:=createTestDir(t)
	type args struct {
		dest string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:"Empty",
			args:args{
				dest:"",
			},
			wantErr:true,
		},
		{
			name:"Not exist",
			args:args{
				dest:"error",
			},
			wantErr:true,
		},
		{
			name:"Exist",
			args:args{
				dest:dirPath,
			},
			wantErr:false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := checkDest(tt.args.dest); (err != nil) != tt.wantErr {
				t.Errorf("checkDest() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_dirExist(t *testing.T) {
	defer removeTestDir()
	dirPath := createTestDir(t)
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name:"Not exist",
			args:args{
				path:"error",
			},
			want:false,
		},
		{
			name:"Not dir",
			args:args{
				path:"args_test.go",
			},
			want:false,
		},
		{
			name:"Exist",
			args:args{
				path:dirPath,
			},
			want:true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dirExist(tt.args.path); got != tt.want {
				t.Errorf("dirExist() = %v, want %v", got, tt.want)
			}
		})
	}
}
