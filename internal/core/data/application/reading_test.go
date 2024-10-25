//
// Copyright (C) 2023 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package application

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/edgexfoundry/edgex-go/internal/core/data/config"
	"github.com/edgexfoundry/edgex-go/internal/core/data/container"
	dbMock "github.com/edgexfoundry/edgex-go/internal/core/data/infrastructure/interfaces/mocks"
	"github.com/edgexfoundry/edgex-go/internal/core/data/mocks"
	"github.com/edgexfoundry/go-mod-bootstrap/v4/di"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/errors"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/models"
)

func TestAllReadings(t *testing.T) {
	readings := buildReadings()
	totalCount := uint32(len(readings))

	dic := mocks.NewMockDIC()
	dbClientMock := &dbMock.DBClient{}
	dbClientMock.On("AllReadings", 0, 20).Return(readings, nil)
	dbClientMock.On("ReadingTotalCount").Return(totalCount, nil)
	dbClientMock.On("AllReadings", 3, 10).Return([]models.Reading{}, errors.NewCommonEdgeX(errors.KindRangeNotSatisfiable, "query objects bounds out of range.", nil))
	dic.Update(di.ServiceConstructorMap{
		container.DBClientInterfaceName: func(get di.Get) interface{} {
			return dbClientMock
		},
	})

	tests := []struct {
		name               string
		offset             int
		limit              int
		errorExpected      bool
		ExpectedErrKind    errors.ErrKind
		expectedCount      int
		expectedTotalCount uint32
		expectedStatusCode int
	}{
		{"Valid - all readings", 0, 20, false, "", len(readings), totalCount, http.StatusOK},
		{"Invalid - bounds out of range", 3, 10, true, errors.KindRangeNotSatisfiable, 0, 0, http.StatusRequestedRangeNotSatisfiable},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			readings, total, err := AllReadings(testCase.offset, testCase.limit, dic)
			if testCase.errorExpected {
				require.Error(t, err)
				assert.NotEmpty(t, err.Error(), "Error message is empty")
				assert.Equal(t, testCase.ExpectedErrKind, errors.Kind(err), "Error kind not as expected")
				assert.Equal(t, testCase.expectedStatusCode, err.Code(), "Status code not as expected")
			} else {
				require.NoError(t, err)
				assert.Equal(t, testCase.expectedCount, len(readings), "Reading count is not expected")
				assert.Equal(t, testCase.expectedTotalCount, total, "Total count is not expected")
			}
		})
	}
}

func TestReadingsByTimeRange(t *testing.T) {
	readings := buildReadings()
	totalCount5 := uint32(5)
	totalCount3 := uint32(3)

	dic := mocks.NewMockDIC()
	dbClientMock := &dbMock.DBClient{}
	dbClientMock.On("ReadingCountByTimeRange", readings[0].GetBaseReading().Origin, readings[4].GetBaseReading().Origin).Return(totalCount5, nil)
	dbClientMock.On("ReadingsByTimeRange", readings[0].GetBaseReading().Origin, readings[4].GetBaseReading().Origin, 0, 10).Return(readings, nil)
	dbClientMock.On("ReadingCountByTimeRange", readings[1].GetBaseReading().Origin, readings[3].GetBaseReading().Origin).Return(totalCount3, nil)
	dbClientMock.On("ReadingsByTimeRange", readings[1].GetBaseReading().Origin, readings[3].GetBaseReading().Origin, 0, 10).Return([]models.Reading{readings[3], readings[2], readings[1]}, nil)
	dbClientMock.On("ReadingsByTimeRange", readings[1].GetBaseReading().Origin, readings[3].GetBaseReading().Origin, 1, 2).Return([]models.Reading{readings[2], readings[1]}, nil)
	dbClientMock.On("ReadingsByTimeRange", readings[1].GetBaseReading().Origin, readings[3].GetBaseReading().Origin, 4, 2).Return(nil, errors.NewCommonEdgeX(errors.KindRangeNotSatisfiable, "query objects bounds out of range", nil))
	dic.Update(di.ServiceConstructorMap{
		container.DBClientInterfaceName: func(get di.Get) interface{} {
			return dbClientMock
		},
	})

	tests := []struct {
		name               string
		start              int64
		end                int64
		offset             int
		limit              int
		errorExpected      bool
		ExpectedErrKind    errors.ErrKind
		expectedCount      int
		expectedTotalCount uint32
		expectedStatusCode int
	}{
		{"Valid - all readings", readings[0].GetBaseReading().Origin, readings[4].GetBaseReading().Origin, 0, 10, false, "", 5, totalCount5, http.StatusOK},
		{"Valid - readings trimmed by latest and oldest", readings[1].GetBaseReading().Origin, readings[3].GetBaseReading().Origin, 0, 10, false, "", 3, totalCount3, http.StatusOK},
		{"Valid - readings trimmed by latest and oldest and skipped first", readings[1].GetBaseReading().Origin, readings[3].GetBaseReading().Origin, 1, 2, false, "", 2, totalCount3, http.StatusOK},
		{"Invalid - bounds out of range", readings[1].GetBaseReading().Origin, readings[3].GetBaseReading().Origin, 4, 2, true, errors.KindRangeNotSatisfiable, 0, uint32(0), http.StatusRequestedRangeNotSatisfiable},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			readings, totalCount, err := ReadingsByTimeRange(testCase.start, testCase.end, testCase.offset, testCase.limit, dic)
			if testCase.errorExpected {
				require.Error(t, err)
				assert.NotEmpty(t, err.Error(), "Error message is empty")
				assert.Equal(t, testCase.ExpectedErrKind, errors.Kind(err), "Error kind not as expected")
				assert.Equal(t, testCase.expectedStatusCode, err.Code(), "Status code not as expected")
			} else {
				require.NoError(t, err)
				assert.Equal(t, testCase.expectedCount, len(readings), "Reading count is not expected")
				assert.Equal(t, testCase.expectedTotalCount, totalCount, "Total count is not expected")
			}
		})
	}
}

func TestReadingsByResourceName(t *testing.T) {
	readings := buildReadings()
	totalCount := uint32(len(readings))

	dic := mocks.NewMockDIC()
	dbClientMock := &dbMock.DBClient{}
	dbClientMock.On("ReadingCountByResourceName", testDeviceResourceName).Return(totalCount, nil)
	dbClientMock.On("ReadingsByResourceName", 0, 20, testDeviceResourceName).Return(readings, nil)
	dbClientMock.On("ReadingsByResourceName", len(readings)+1, 10, testDeviceResourceName).Return([]models.Reading{}, errors.NewCommonEdgeX(errors.KindRangeNotSatisfiable, "query objects bounds out of range.", nil))
	dic.Update(di.ServiceConstructorMap{
		container.DBClientInterfaceName: func(get di.Get) interface{} {
			return dbClientMock
		},
	})

	tests := []struct {
		name               string
		offset             int
		limit              int
		resourceName       string
		errorExpected      bool
		ExpectedErrKind    errors.ErrKind
		expectedCount      int
		expectedTotalCount uint32
		expectedStatusCode int
	}{
		{"Valid - all readings", 0, 20, testDeviceResourceName, false, "", len(readings), totalCount, http.StatusOK},
		{"Invalid - bounds out of range", len(readings) + 1, 10, testDeviceResourceName, true, errors.KindRangeNotSatisfiable, 0, 0, http.StatusRequestedRangeNotSatisfiable},
		{"Invalid - empty resource name", len(readings) + 1, 10, "", true, errors.KindContractInvalid, 0, 0, http.StatusBadRequest},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			readings, total, err := ReadingsByResourceName(testCase.offset, testCase.limit, testCase.resourceName, dic)
			if testCase.errorExpected {
				require.Error(t, err)
				assert.NotEmpty(t, err.Error(), "Error message is empty")
				assert.Equal(t, testCase.ExpectedErrKind, errors.Kind(err), "Error kind not as expected")
				assert.Equal(t, testCase.expectedStatusCode, err.Code(), "Status code not as expected")
			} else {
				require.NoError(t, err)
				assert.Equal(t, testCase.expectedCount, len(readings), "Reading count is not expected")
				assert.Equal(t, testCase.expectedTotalCount, total, "Total count is not expected")
			}
		})
	}
}

func TestReadingsByDeviceName(t *testing.T) {
	readings := buildReadings()
	totalCount := uint32(len(readings))

	dic := mocks.NewMockDIC()
	dbClientMock := &dbMock.DBClient{}
	dbClientMock.On("ReadingCountByDeviceName", testDeviceName).Return(totalCount, nil)
	dbClientMock.On("ReadingsByDeviceName", 0, 20, testDeviceName).Return(readings, nil)
	dbClientMock.On("ReadingsByDeviceName", 3, 10, testDeviceName).Return([]models.Reading{}, errors.NewCommonEdgeX(errors.KindRangeNotSatisfiable, "query objects bounds out of range.", nil))
	dic.Update(di.ServiceConstructorMap{
		container.DBClientInterfaceName: func(get di.Get) interface{} {
			return dbClientMock
		},
	})

	tests := []struct {
		name               string
		offset             int
		limit              int
		deviceName         string
		errorExpected      bool
		ExpectedErrKind    errors.ErrKind
		expectedCount      int
		expectedTotalCount uint32
		expectedStatusCode int
	}{
		{"Valid - all readings", 0, 20, testDeviceName, false, "", len(readings), totalCount, http.StatusOK},
		{"Invalid - bounds out of range", 3, 10, testDeviceName, true, errors.KindRangeNotSatisfiable, 0, 0, http.StatusRequestedRangeNotSatisfiable},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			readings, total, err := ReadingsByDeviceName(testCase.offset, testCase.limit, testCase.deviceName, dic)
			if testCase.errorExpected {
				require.Error(t, err)
				assert.NotEmpty(t, err.Error(), "Error message is empty")
				assert.Equal(t, testCase.ExpectedErrKind, errors.Kind(err), "Error kind not as expected")
				assert.Equal(t, testCase.expectedStatusCode, err.Code(), "Status code not as expected")
			} else {
				require.NoError(t, err)
				assert.Equal(t, testCase.expectedCount, len(readings), "Reading count is not expected")
				assert.Equal(t, testCase.expectedTotalCount, total, "Total count is not expected")
			}
		})
	}
}

func TestReadingCountByDeviceName(t *testing.T) {
	expectedReadingCount := uint32(656672)
	dic := mocks.NewMockDIC()
	dbClientMock := &dbMock.DBClient{}
	dbClientMock.On("ReadingCountByDeviceName", testDeviceName).Return(expectedReadingCount, nil)
	dic.Update(di.ServiceConstructorMap{
		container.DBClientInterfaceName: func(get di.Get) interface{} {
			return dbClientMock
		},
	})

	tests := []struct {
		name               string
		deviceName         string
		errorExpected      bool
		ExpectedErrKind    errors.ErrKind
		expectedCount      uint32
		expectedStatusCode int
	}{
		{"Valid - all readings", testDeviceName, false, "", expectedReadingCount, http.StatusOK},
		{"Invalid - empty device name", "", true, errors.KindContractInvalid, 0, http.StatusBadRequest},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			count, err := ReadingCountByDeviceName(testCase.deviceName, dic)
			if testCase.errorExpected {
				require.Error(t, err)
				assert.NotEmpty(t, err.Error(), "Error message is empty")
				assert.Equal(t, testCase.ExpectedErrKind, errors.Kind(err), "Error kind not as expected")
				assert.Equal(t, testCase.expectedStatusCode, err.Code(), "Status code not as expected")
			} else {
				require.NoError(t, err)
				assert.Equal(t, expectedReadingCount, count, "Reading total count is not expected")
			}
		})
	}
}

func TestPurgeReading(t *testing.T) {
	dic := mocks.NewMockDIC()
	coreDataConfig := container.ConfigurationFrom(dic.Get)
	coreDataConfig.Retention = config.ReadingRetention{
		Enabled:  true,
		Interval: "1s",
		MaxCap:   5,
		MinCap:   3,
	}
	dic.Update(di.ServiceConstructorMap{
		container.ConfigurationName: func(get di.Get) interface{} {
			return coreDataConfig
		},
	})

	tests := []struct {
		name         string
		readingCount uint32
	}{
		{"invoke reading purging", coreDataConfig.Retention.MaxCap},
		{"not invoke reading purging", coreDataConfig.Retention.MinCap},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			dbClientMock := &dbMock.DBClient{}
			var reading models.Reading = models.SimpleReading{}
			dbClientMock.On("LatestReadingByOffset", coreDataConfig.Retention.MinCap).Return(reading, nil)
			dbClientMock.On("ReadingTotalCount").Return(testCase.readingCount, nil)
			dbClientMock.On("DeleteEventsByAge", mock.Anything).Return(nil)
			dic.Update(di.ServiceConstructorMap{
				container.DBClientInterfaceName: func(get di.Get) interface{} {
					return dbClientMock
				},
			})
			err := purgeReading(dic)
			require.NoError(t, err)
			if testCase.readingCount >= coreDataConfig.Retention.MaxCap {
				dbClientMock.AssertCalled(t, "DeleteEventsByAge", mock.Anything)
			} else {
				dbClientMock.AssertNotCalled(t, "DeleteEventsByAge", mock.Anything)
			}
		})
	}
}
