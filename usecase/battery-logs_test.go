package usecase

import (
	"context"
	"github/Shimaa-Ibrahim/drones/iofile"
	"github/Shimaa-Ibrahim/drones/repository"
	repoEntity "github/Shimaa-Ibrahim/drones/repository/entity"
	"github/Shimaa-Ibrahim/drones/repository/mock"
	"os"
	"path"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "../")
	err := os.Chdir(dir)
	if err != nil {
		panic(err)
	}
}

func TestCreateBatteryLogs(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	type mocks struct {
		droneRepository       repository.DroneRepo
		BatteryLogsRepository repository.BatteryLogRepo
		iofile                iofile.IOFiles
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr string
		mocks   mocks
	}{
		{
			name: "successful log file creation",
			args: args{ctx: context.Background()},
			mocks: mocks{

				droneRepository:       mock.NewSuccessMockedDrone(),
				BatteryLogsRepository: mock.NewMockedBatteryLogsRepository(),
				iofile:                iofile.NewIOFile(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			batteryLogsUsecase := NewBatterylogUsecase(tt.mocks.droneRepository, tt.mocks.BatteryLogsRepository, tt.mocks.iofile)
			err := batteryLogsUsecase.Create(tt.args.ctx)
			if err != nil {
				assert.Equal(t, tt.wantErr, err.Error())
				return
			}
		})
	}
}

func TestStringfyBatteryLogs(t *testing.T) {
	type args struct {
		drone repoEntity.Drone
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "stringfy battery logs",
			args: args{drone: repoEntity.Drone{SerialNumber: "123", BatteryCapacity: 100}},
			want: `Drone "123" : battery Level 100`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := stringfyDroneBatteryLog(tt.args.drone)
			assert.Equal(t, tt.want, got)
		})
	}
}
